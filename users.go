package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// getUserCountonSite returns the count for allusers, activeusers and system managers 
// uses direct SQL queries.
func getUserCountonSite(siteName string)(float64,float64,float64) {
	allUserQuery := `SELECT COUNT(name) FROM tabUser WHERE enabled=1 AND user_type != 'Website User' AND name NOT IN ("Administrator","Guest");`
	activeUserQuery := `SELECT COUNT(*) FROM tabUser WHERE enabled=1 AND user_type != 'Website User' AND name NOT IN ("Administrator","Guest") AND hour(timediff(now(), last_active)) < 72;`
	systemManagerQuery := "SELECT DISTINCT COUNT(name) FROM `tabUser` AS p WHERE enabled=1 AND docstatus<2 AND name NOT IN (\"Administrator\",\"Guest\") AND EXISTS(SELECT * FROM `tabHas Role` AS ur WHERE ur.parent=p.name AND ur.role=\"System Manager\");"
	log.Println("Usercount func:",siteName)
	db, err := sql.Open("mysql", generateDbURI(siteName))
	defer db.Close()
	if err != nil {
		log.Println("Database Connection Error ", err)
	}
	allUserQuerystmt, err := db.Prepare(allUserQuery)
	defer allUserQuerystmt.Close()
	if err != nil {
		log.Println("Preparing allUserQuery failed ", err)
	}
	activeUserQuerystmt, err := db.Prepare(activeUserQuery)
	defer activeUserQuerystmt.Close()
	if err != nil {
		log.Println("Preparing activeUserQuery failed ", err)
	}
	systemManagerQuerystmt, err := db.Prepare(systemManagerQuery)
	defer systemManagerQuerystmt.Close()
	if err != nil {
		log.Println("Preparing systemManagerQuery failed ", err)
	}
	var activeUsers, systemManagers, allUsers float64

	if err := allUserQuerystmt.QueryRow().Scan(&allUsers); err != nil {
		log.Println("AllUserQuery failed ", err)
	}

	if err := activeUserQuerystmt.QueryRow().Scan(&activeUsers); err != nil {
		log.Println("activeUserQuery failed", err)
	}

	if err := systemManagerQuerystmt.QueryRow().Scan(&systemManagers); err != nil {
		log.Println("systemManagerQuery failed ", err)
	}


	return allUsers,activeUsers,systemManagers

}
