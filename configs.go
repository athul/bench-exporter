package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

// readSiteConfig reads the site config of the 
// sitename argument. Returns the DB name and Password
func readSiteConfig(siteName string) (string,string) {
	siteDir := filepath.Join(benchDir, "sites", siteName)

	siteConfigFile, err := ioutil.ReadFile(siteDir + "/site_config.json")
	if err != nil {
		log.Println("Unable to read Site Config ", err)
	}

	var config map[string]interface{}

	json.Unmarshal(siteConfigFile, &config)

	log.Println(config)
  return config["db_name"].(string),config["db_password"].(string)
}

// readCommonSiteConfig reads the common site config and checks if the DB host is specified
func readCommonSiteConfig() string {
	siteDir := filepath.Join(benchDir, "sites")
	commonConfig, err := ioutil.ReadFile(siteDir + "/common_site_config.json")
	if err != nil {
		log.Println("Unable to read common_site_config.json ", err)
	}
	var commonConfigMap map[string]interface{}

	err = json.Unmarshal(commonConfig, &commonConfigMap)
	if err != nil {
		log.Println("Error Unmarshaling common_site_config ", err)
	}
	if commonConfigMap["db_host"] != nil {
		return commonConfigMap["db_host"].(string)
	}
	return ""
}

// generateDbURI generates the mysql URI for the specific site
func generateDbURI(siteName string)string{
  var URI string
  user,pass := readSiteConfig(siteName)
  host := readCommonSiteConfig()
  if host!=""{
    URI=fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",user,pass,host,user)
  }else{
    URI=fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s",user,pass,user)
  }
  return URI
}
