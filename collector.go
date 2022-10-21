package main

import (
	"log"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type benchMetrics struct {
	Sites            []string
	benchVersion     *prometheus.Desc
	benchAppVersions *prometheus.Desc
	benchSites       *prometheus.Desc
	appsonSites      *prometheus.Desc
	usersonSites     *prometheus.Desc
	activeUsers      *prometheus.Desc
	systemUsers      *prometheus.Desc
}
// newBenchMetrics is benchMetrics contructor
func newBenchMetrics() *benchMetrics {
	return &benchMetrics{
		benchVersion:     prometheus.NewDesc("bench_version", "Shows the version of Bench", []string{"version"}, nil),
		benchAppVersions: prometheus.NewDesc("bench_apps", "Shows the versions of apps on the bench", []string{"app_name", "version", "commit"}, nil),
		appsonSites:      prometheus.NewDesc("bench_siteapps", "Shows the site and number of apps installed", []string{"site_name", "site_apps"}, nil),
		benchSites:       prometheus.NewDesc("bench_sites", "Shows the sites present on bench", []string{"sites"}, nil),
		usersonSites:     prometheus.NewDesc("all_users", "Shows the number of users on the site", []string{"site_name"}, nil),
		activeUsers:      prometheus.NewDesc("active_users", "Shows the number of active users on the site", []string{"site_name"}, nil),
		systemUsers:      prometheus.NewDesc("system_user", "Shows the number of system managers on the site", []string{"site_name"}, nil),
		Sites:            getSites(),
	}
}

// Describe implements the Describe method for the collectors
func (b *benchMetrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- b.benchVersion
	ch <- b.benchAppVersions
	ch <- b.appsonSites
	ch <- b.benchSites
	ch <- b.activeUsers
	ch <- b.usersonSites
	ch <- b.systemUsers
}
// Collect implements the collector for prometheus
func (b *benchMetrics) Collect(ch chan<- prometheus.Metric) {
	apps := getApps()
	versions := getAppVersions()
	benchVersion := getBenchVersion()
	sites := []string{}
	for i := range apps {
		sites = append(sites, apps[i].Site)
		log.Println(apps[i].Site)
		log.Println(strings.Join(apps[i].Apps, ","))
		all, active, smusers := getUserCountonSite(apps[i].Site)
		ch <- prometheus.MustNewConstMetric(b.appsonSites, prometheus.CounterValue, float64(len(apps[i].Apps)), apps[i].Site, strings.Join(apps[i].Apps, ","))
		ch <- prometheus.MustNewConstMetric(b.usersonSites, prometheus.CounterValue, all, apps[i].Site)
		ch <- prometheus.MustNewConstMetric(b.activeUsers, prometheus.CounterValue, active, apps[i].Site)
		ch <- prometheus.MustNewConstMetric(b.systemUsers, prometheus.CounterValue, smusers, apps[i].Site)
	}
	ch <- prometheus.MustNewConstMetric(b.benchSites, prometheus.CounterValue, float64(len(apps)), strings.Join(sites, ","))
	for i := range versions {
		ch <- prometheus.MustNewConstMetric(b.benchAppVersions, prometheus.UntypedValue, 1.0, versions[i].App, versions[i].Version, versions[i].Commit)
	}
	ch <- prometheus.MustNewConstMetric(b.benchVersion, prometheus.CounterValue, 1.0, benchVersion)

}
