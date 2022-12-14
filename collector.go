package main

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type benchMetrics struct {
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
		benchAppVersions: prometheus.NewDesc("bench_apps", "Shows the versions of apps on the bench", []string{"app_name", "version", "commit"}, nil),
		appsonSites:      prometheus.NewDesc("bench_siteapps", "Shows the site and number of apps installed", []string{"site_name", "site_apps"}, nil),
		benchSites:       prometheus.NewDesc("bench_sites", "Shows the sites present on bench", []string{"sites"}, nil),
		usersonSites:     prometheus.NewDesc("bench_all_users", "Shows the number of users on the site", []string{"site_name"}, nil),
		activeUsers:      prometheus.NewDesc("bench_active_users", "Shows the number of active users on the site", []string{"site_name"}, nil),
		systemUsers:      prometheus.NewDesc("bench_system_user", "Shows the number of system managers on the site", []string{"site_name"}, nil),
	}
}

// Describe implements the Describe method for the collectors
func (b *benchMetrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- b.benchAppVersions
	ch <- b.appsonSites
	ch <- b.benchSites
	ch <- b.activeUsers
	ch <- b.usersonSites
	ch <- b.systemUsers
}

// Collect implements the collector for prometheus
func (b *benchMetrics) Collect(ch chan<- prometheus.Metric) {
	var bench = getAll()
	versions := bench.Version
	// versions := getAppVersions()
	// benchVersion := getBenchVersion()
	sites := []string{}
	for i := range bench.SAPS {
		sites = append(sites, bench.SAPS[i].Site)
		//TODO: Make a chan map[string]int for concurrent execution
		all, active, smusers := getUserCountonSite(bench.SAPS[i].Site)
		ch <- prometheus.MustNewConstMetric(b.appsonSites, prometheus.CounterValue, float64(len(bench.SAPS[i].Apps)), bench.SAPS[i].Site, strings.Join(bench.SAPS[i].Apps, ","))
		if userCounts {
			ch <- prometheus.MustNewConstMetric(b.usersonSites, prometheus.CounterValue, all, bench.SAPS[i].Site)
			ch <- prometheus.MustNewConstMetric(b.activeUsers, prometheus.CounterValue, active, bench.SAPS[i].Site)
			ch <- prometheus.MustNewConstMetric(b.systemUsers, prometheus.CounterValue, smusers, bench.SAPS[i].Site)
		}
		ch <- prometheus.MustNewConstMetric(b.benchSites, prometheus.CounterValue, float64(len(bench.Sites)), sites[i])
	}
	for i := range versions {
		ch <- prometheus.MustNewConstMetric(b.benchAppVersions, prometheus.CounterValue, 1.0, versions[i].App, versions[i].Version, versions[i].Commit)
	}
}
