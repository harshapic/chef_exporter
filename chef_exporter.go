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
        ManageEventStatus           prometheus.Gauge
        ManageRedisStatus           prometheus.Gauge
        ManageWebStatus             prometheus.Gauge
        ManageWorkerStatus          prometheus.Gauge
        ServerBookshelfStatus           prometheus.Gauge
        ServerEcsyncStatus             prometheus.Gauge
        ServerNginxStatus          prometheus.Gauge
        ServerOcbifrostStatus           prometheus.Gauge
        ServerOcidStatus             prometheus.Gauge
        ServerErchefStatus          prometheus.Gauge
        ServerExpanderStatus           prometheus.Gauge
        ServerSolr4Status             prometheus.Gauge
        ServerPostgresrStatus          prometheus.Gauge
        ServerRabbitmqStatus           prometheus.Gauge
        ServerRedisStatus             prometheus.Gauge
}

func NewExporter() *Exporter {
        return &Exporter{
                ManageEventStatus: prometheus.NewGauge(prometheus.GaugeOpts{
                        Namespace: namespace,
                        Name:      "ManageEventStatus",
                        Help:      "ManageEventStatus",
                }),
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
                ServerBookshelfStatus: prometheus.NewGauge(prometheus.GaugeOpts{
                        Namespace: namespace,
                        Name:      "ServerBookshelfStatus",
                        Help:      "ServerBookshelfStatus",
                }),
                ServerEcsyncStatus: prometheus.NewGauge(prometheus.GaugeOpts{
                        Namespace: namespace,
                        Name:      "ServerEcsyncStatus",
                        Help:      "ServerEcsyncStatus",
                }),
                ServerNginxStatus: prometheus.NewGauge(prometheus.GaugeOpts{
                        Namespace: namespace,
                        Name:      "ServerNginxStatus",
                        Help:      "ServerNginxStatus",
                }),
                ServerOcbifrostStatus: prometheus.NewGauge(prometheus.GaugeOpts{
                        Namespace: namespace,
                        Name:      "ServerOcbifrostStatus",
                        Help:      "ServerOcbifrostStatus",
                }),
                ServerOcidStatus: prometheus.NewGauge(prometheus.GaugeOpts{
                        Namespace: namespace,
                        Name:      "ServerOcidStatus",
                        Help:      "ServerOcidStatus",
                }),
                ServerErchefStatus: prometheus.NewGauge(prometheus.GaugeOpts{
                        Namespace: namespace,
                        Name:      "ServerErchefStatus",
                        Help:      "ServerErchefStatus",
                }),
                ServerExpanderStatus: prometheus.NewGauge(prometheus.GaugeOpts{
                        Namespace: namespace,
                        Name:      "ServerExpanderStatus",
                        Help:      "ServerExpanderStatus",
                }),
                ServerSolr4Status: prometheus.NewGauge(prometheus.GaugeOpts{
                        Namespace: namespace,
                        Name:      "ServerSolr4Status",
                        Help:      "ServerSolr4Status",
                }),
                ServerPostgresrStatus: prometheus.NewGauge(prometheus.GaugeOpts{
                        Namespace: namespace,
                        Name:      "ServerPostgresrStatus",
                        Help:      "ServerPostgresrStatus",
                }),
                ServerRabbitmqStatus: prometheus.NewGauge(prometheus.GaugeOpts{
                        Namespace: namespace,
                        Name:      "ServerRabbitmqStatus",
                        Help:      "ServerRabbitmqStatus",
                }),
                ServerRedisStatus: prometheus.NewGauge(prometheus.GaugeOpts{
                        Namespace: namespace,
                        Name:      "ServerRedisStatus",
                        Help:      "ServerRedisStatus",
                }),

        }
}

// Describe implements the prometheus.Collector interface.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
        e.ManageEventStatus.Describe(ch)
        e.ManageRedisStatus.Describe(ch)
        e.ManageWebStatus.Describe(ch)
        e.ManageWorkerStatus.Describe(ch)
        e.ServerBookshelfStatus.Describe(ch)
        e.ServerEcsyncStatus.Describe(ch)
        e.ServerNginxStatus.Describe(ch)
        e.ServerOcbifrostStatus.Describe(ch)
        e.ServerOcidStatus.Describe(ch)
        e.ServerErchefStatus.Describe(ch)
        e.ServerRabbitmqStatus.Describe(ch)
        e.ServerExpanderStatus.Describe(ch)
        e.ServerSolr4Status.Describe(ch)
        e.ServerPostgresrStatus.Describe(ch)
        e.ServerRedisStatus.Describe(ch)
}

// Collect implements the prometheus.Collector interface.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
        manage := strings.Split(execute("chef-manage-ctl status"), "\n")
        server := strings.Split(execute("chef-server-ctl status"), "\n")
        m := make(map[string]string)
        for _,s := range manage {
          t := strings.Split(s," ")
          m[t[1]]=t[0]
        }
        for _,s := range server {
          t := strings.Split(s," ")
          m[t[1]]=t[0]
        }
         if m["events:"] == "run:" {
           e.ManageEventStatus.Set(1)
        } else {
           e.ManageEventStatus.Set(0)
        }
        if m["redis:"] == "run:" {
           e.ManageRedisStatus.Set(1)
        } else {
           e.ManageRedisStatus.Set(0)
        }
        if m["web:"] == "run:" {
           e.ManageWebStatus.Set(1)
        } else {
           e.ManageWebStatus.Set(0)
        }
        if m["worker:"] == "run:" {
           e.ManageWorkerStatus.Set(1)
        } else {
           e.ManageWorkerStatus.Set(0)
        }
        if m["bookshelf:"] == "run:" {
           e.ServerBookshelfStatus.Set(1)
        } else {
           e.ServerBookshelfStatus.Set(0)
        }
        if m["ec_sync_client:"] == "run:" {
           e.ServerEcsyncStatus.Set(1)
        } else {
           e.ServerEcsyncStatus.Set(0)
        }
        if m["nginx:"] == "run:" {
           e.ServerNginxStatus.Set(1)
        } else {
           e.ServerNginxStatus.Set(0)
        }
        if m["oc_bifrost:"] == "run:" {
           e.ServerOcbifrostStatus.Set(1)
        } else {
           e.ServerOcbifrostStatus.Set(0)
        }
        if m["oc_id:"] == "run:" {
           e.ServerOcidStatus.Set(1)
        } else {
           e.ServerOcidStatus.Set(0)
        }
        if m["opscode-erchef:"] == "run:" {
           e.ServerErchefStatus.Set(1)
        } else {
           e.ServerErchefStatus.Set(0)
        }
        if m["opscode-expander:"] == "run:" {
           e.ServerExpanderStatus.Set(1)
        } else {
           e.ServerExpanderStatus.Set(0)
        }
        if m["opscode-solr4:"] == "run:" {
           e.ServerSolr4Status.Set(1)
        } else {
           e.ServerSolr4Status.Set(0)
        }
        if m["postgresql:"] == "run:" {
           e.ServerPostgresrStatus.Set(1)
        } else {
           e.ServerPostgresrStatus.Set(0)
        }
        if m["rabbitmq:"] == "run:" {
           e.ServerRabbitmqStatus.Set(1)
        } else {
           e.ServerRabbitmqStatus.Set(0)
        }
        if m["redis_lb:"] == "run:" {
           e.ServerRedisStatus.Set(1)
        } else {
           e.ServerRedisStatus.Set(0)
        }
        e.ManageEventStatus.Collect(ch)
        e.ManageRedisStatus.Collect(ch)
        e.ManageWebStatus.Collect(ch)
        e.ManageWorkerStatus.Collect(ch)
        e.ServerBookshelfStatus.Collect(ch)
        e.ServerEcsyncStatus.Collect(ch)
        e.ServerNginxStatus.Collect(ch)
        e.ServerOcbifrostStatus.Collect(ch)
        e.ServerOcidStatus.Collect(ch)
        e.ServerErchefStatus.Collect(ch)
        e.ServerRabbitmqStatus.Collect(ch)
        e.ServerExpanderStatus.Collect(ch)
        e.ServerSolr4Status.Collect(ch)
        e.ServerPostgresrStatus.Collect(ch)
        e.ServerRedisStatus.Collect(ch)
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
