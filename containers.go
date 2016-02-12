package main

import (
	"fmt"

	"sort"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	defaultHelp = "Metric autogenerated by marathon_exporter."
)

type CounterContainer struct {
	counters map[string]*prometheus.CounterVec
}

func NewCounterContainer() *CounterContainer {
	return &CounterContainer{
		counters: make(map[string]*prometheus.CounterVec),
	}
}

func (c *CounterContainer) GetOrCreate(metricName string, labels ...string) *prometheus.CounterVec {
	key := containerKey(metricName, labels)
	counter, ok := c.counters[key]

	if !ok {
		counter = prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      metricName,
			Help:      defaultHelp,
		}, labels)

		c.counters[key] = counter
	}

	return counter
}

type GaugeContainer struct {
	gauges map[string]*prometheus.GaugeVec
}

func NewGaugeContainer() *GaugeContainer {
	return &GaugeContainer{
		gauges: make(map[string]*prometheus.GaugeVec),
	}
}

func (c *GaugeContainer) GetOrCreate(metricName string, labels ...string) *prometheus.GaugeVec {
	key := containerKey(metricName, labels)
	gauge, ok := c.gauges[key]

	if !ok {
		gauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      metricName,
			Help:      defaultHelp,
		}, labels)

		c.gauges[key] = gauge
	}

	return gauge
}

func containerKey(metric string, labels []string) string {
	s := make([]string, len(labels))
	copy(s, labels)
	sort.Strings(s)
	return fmt.Sprintf("%s{%v}", metric, strings.Join(s, ","))
}
