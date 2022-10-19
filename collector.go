package main

import (
	"log"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type benchMetrics struct {
	benchVersion     *prometheus.Desc
	benchAppVersions *prometheus.Desc
	benchSites       *prometheus.Desc
	appsonSites      *prometheus.Desc
}

func newBenchMetrics() *benchMetrics {
	return &benchMetrics{
		benchVersion:     prometheus.NewDesc("bench_version", "Shows the version of Bench", []string{"version"}, nil),
		benchAppVersions: prometheus.NewDesc("bench_appversions", "Shows the versions of apps on the bench", []string{"app_name", "version", "commit"}, nil),
		appsonSites:      prometheus.NewDesc("bench_siteapps", "Shows the site and number of apps installed", []string{"site_name", "site_apps"}, nil),
		benchSites:       prometheus.NewDesc("bench_sites", "Shows the sites present on bench", []string{"sites"}, nil),
	}
}

func (b *benchMetrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- b.benchVersion
	ch <- b.benchAppVersions
	ch <- b.appsonSites
	ch <- b.benchSites
}

func (b *benchMetrics) Collect(ch chan<- prometheus.Metric) {
	apps := getApps()
	versions := getAppVersions()
	benchVersion := getBenchVersion()
	sites := []string{}
	for i := range apps {
		sites = append(sites, apps[i].Site)
		log.Println(apps[i].Site)
		log.Println(strings.Join(apps[i].Apps, ","))
		ch <- prometheus.MustNewConstMetric(b.appsonSites, prometheus.CounterValue, float64(len(apps[i].Apps)), apps[i].Site, strings.Join(apps[i].Apps, ","))
	}
	ch <- prometheus.MustNewConstMetric(b.benchSites, prometheus.CounterValue, float64(len(apps)), strings.Join(sites, ","))
	for i := range versions {
		ch <- prometheus.MustNewConstMetric(b.benchAppVersions, prometheus.UntypedValue, float64(len(versions)), versions[i].App, versions[i].Version, versions[i].Commit)
	}
	ch <- prometheus.MustNewConstMetric(b.benchVersion, prometheus.CounterValue, 1.0, benchVersion)

}
