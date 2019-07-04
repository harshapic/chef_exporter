# Chef Exporter for Prometheus
Exports chef metrics via HTTP for Prometheus consumption.

How to build
```
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/log
go build chef_exporter.go
Running exporter on chef-server host:
noup chef_exporter &
```
