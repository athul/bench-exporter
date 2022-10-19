package main

import (
	"log"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type benchMetrics struct {
benchVersions    *prometheus.Desc
	benchAppVersions *prometheus.Desc
	benchSites       *prometheus.Desc
	appsonSites      *prometheus.Desc
}

func newBenchMetrics() *benchMetrics {
	return &benchMetrics{
		benchVersions:    prometheus.NewDesc("bench_version", "Shows the version of Bench", nil, nil),
		benchAppVersions: prometheus.NewDesc("bench_appversions", "Shows the versions of apps on the bench", []string{"app_name", "version", "commit"}, nil),
		appsonSites:      prometheus.NewDesc("bench_siteapps", "Shows the site and number of apps installed", []string{"site_name", "site_apps"}, nil),
	}
}

func (b *benchMetrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- b.benchVersions
	ch <- b.benchAppVersions
	ch <- b.appsonSites
}

func (b *benchMetrics) Collect(ch chan<- prometheus.Metric) {
	apps := getApps()
	for i := range apps {
		log.Println(apps[i].Site)
		log.Println(strings.Join(apps[i].Apps, ","))
		ch <- prometheus.MustNewConstMetric(b.appsonSites, prometheus.CounterValue, float64(len(apps[i].Apps)), apps[i].Site, strings.Join(apps[i].Apps, ","))
		// ch <- prometheus.MustNewConstMetric(b.benchAppVersions,prometheus.CounterValue,14)
	}
	versions := getAppVersions()

	for i:=range versions{
	  ch <- prometheus.MustNewConstMetric(b.benchAppVersions,prometheus.UntypedValue,float64(len(versions)),versions[i].App,versions[i].Version,versions[i].Commit)
	}


}
