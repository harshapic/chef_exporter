package main

import (
        "os/exec"
        "flag"
        "fmt"
        "strings"
        "net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/log"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespace = "chef"
)

var (
	listenAddress  = flag.String("web.listen-address", ":9070", "Address on which to expose metrics and web interface.")
	metricsPath    = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
)

type Exporter struct {
	ManageRedisStatus           prometheus.Gauge
	ManageWebStatus             prometheus.Gauge
	ManageWorkerStatus          prometheus.Gauge
}

func NewExporter() *Exporter {
	return &Exporter{
		ManageRedisStatus: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "ManageRedisStatus",
			Help:      "ManageRedisStatus",
		}),
		ManageWebStatus: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "ManageWebStatus",
			Help:      "ManageWebStatus",
		}),
		ManageWorkerStatus: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "ManageWorkerStatus",
			Help:      "ManageWorkerStatus",
		}),
	}
}

// Describe implements the prometheus.Collector interface.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.ManageRedisStatus.Describe(ch)
	e.ManageWebStatus.Describe(ch)
	e.ManageWorkerStatus.Describe(ch)
}

// Collect implements the prometheus.Collector interface.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	if execute("chef-manage-ctl status | grep redis | awk {'print $1'}") == "run:" {
	   e.ManageRedisStatus.Set(1)
	} else {
	   e.ManageRedisStatus.Set(0)
	}
        if execute("chef-manage-ctl status | grep web | awk {'print $1'}") == "run:" {
           e.ManageWebStatus.Set(1)
        } else {
           e.ManageWebStatus.Set(0)
        }
        if execute("chef-manage-ctl status | grep worker | awk {'print $1'}") == "run:" {
           e.ManageWorkerStatus.Set(1)
        } else {
           e.ManageWorkerStatus.Set(0)
        }
	e.ManageRedisStatus.Collect(ch)
	e.ManageWebStatus.Collect(ch)
	e.ManageWorkerStatus.Collect(ch)
}

func execute(cmd string)(s string) {
    out, err := exec.Command("bash", "-c", cmd).Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    output := string(out[:])
    s = strings.TrimSpace(output)
    return s
}

func main() {
        flag.Parse()
	exporter := NewExporter()
	prometheus.MustRegister(exporter)

	log.Printf("Starting Server: %s", *listenAddress)
	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		<head><title>Chef Exporter</title></head>
		<body>
		<h1>Chef
		<p><a href="` + *metricsPath + `">Metrics</a></p>
		</body>
		</html>`))
	})
	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
