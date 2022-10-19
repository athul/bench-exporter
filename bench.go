package main

import (
	"encoding/json"
	"log"
	"os/exec"
)

var benchDir = "/home/frappe/bench-version-14"

type Apps struct {
	Site string
	Apps []string
}

func getApps() []Apps {
	cmd := exec.Command("bench", "--site", "all", "list-apps", "--format", "json")
	cmd.Dir = benchDir
	out, err := cmd.Output()
	if err != nil {
		log.Println("Error getting result")
	}
	log.Println(string(out))
	var apps map[string][]interface{}
	json.Unmarshal(out, &apps)
	log.Println(len(apps))
	var appstruct []Apps
	for k, v := range apps {
		var appnames []string
		for i := range v {
			appnames = append(appnames, v[i].(string))
		}
		appstruct = append(appstruct, Apps{
			Site: k,
			Apps: appnames,
		})
		log.Println(appstruct)
	}
	return appstruct

}
type AppVersions struct {
	Commit  string `json:"commit"`
	App     string `json:"app"`
	Branch  string `json:"branch"`
	Version string `json:"version"`
}

func getAppVersions() []AppVersions{
  cmd := exec.Command("bench", "version", "--format", "json")
	cmd.Dir = benchDir
	out, err := cmd.Output()
	if err!=nil{
	  log.Println("Version retreival failed",err)
	}

	var appversions []AppVersions

	json.Unmarshal(out,&appversions)

	log.Println(appversions)

	return appversions

}
