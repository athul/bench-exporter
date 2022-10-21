package main

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"
)

// Holds the Sites and Apps installed on
// the specific site
type Apps struct {
	Site string
	Apps []string
}

// getSites returns the sites in the bench
// returns the sites as a string slice
// as multitenant systems will be used
func getSites() []string {
	cmd := exec.Command("bench", "--site", "all", "list-apps", "--format", "json")
	cmd.Dir = benchDir
	out, err := cmd.Output()
	if err != nil {
		log.Println("Error getting Sites", err)
	}
	var apps map[string][]interface{}
	json.Unmarshal(out, &apps)
	var sites []string
	for k := range apps {
		sites = append(sites, k)
	}
	return sites
}

// getApps returns the Apps installed on the site
// Will return as a slice of Apps struct
func getApps() []Apps {
	cmd := exec.Command("bench", "--site", "all", "list-apps", "--format", "json")
	cmd.Dir = benchDir
	out, err := cmd.Output()
	if err != nil {
		log.Println("Error getting All apps", err)
	}
	var apps map[string][]interface{}
	json.Unmarshal(out, &apps)
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
	}
	return appstruct
}

// AppVersions hold the JSON struct
// for the Versions of Apps
type AppVersions struct {
	Commit  string `json:"commit"`
	App     string `json:"app"`
	Branch  string `json:"branch"`
	Version string `json:"version"`
}

// getAppVersions returns the Versions of the apps
// the data is fetched from bench with commit, branch and version of the app
func getAppVersions() []AppVersions {
	cmd := exec.Command("bench", "version", "--format", "json")
	cmd.Dir = benchDir
	out, err := cmd.Output()
	if err != nil {
		log.Println("Version retreival failed", err)
	}
	var appversions []AppVersions
	json.Unmarshal(out, &appversions)
	return appversions
}

// getBenchVersion returns the version of Bench CLI installed
func getBenchVersion() string {
	cmd := exec.Command("bench", "--version")
	cmd.Dir = benchDir
	out, err := cmd.Output()
	if err != nil {
		log.Println("Unable to fetch bench version ", err)
	}
	return strings.TrimSuffix(string(out), "\n")
}
