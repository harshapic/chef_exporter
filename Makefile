all: chef_exporter 
.PHONY: all

deps:
	go get github.com/prometheus/client_golang/prometheus
	go get github.com/prometheus/log

namenode_exporter: deps chef_exporter.go
	go build chef_exporter.go

clean:
	rm -rf chef_exporter
