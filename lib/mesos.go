package mpmesos

import (
	"strings"

	"fmt"
	"net/http"

	"encoding/json"

	mp "github.com/mackerelio/go-mackerel-plugin"
	"gopkg.in/alecthomas/kingpin.v2"
)

// MesosPlugin mackerel plugin
type MesosPlugin struct {
	Host   string
	Port   string
	Node   string
	Prefix string
}

// MetricKeyPrefix interface for PluginWithPrefix
func (p *MesosPlugin) MetricKeyPrefix() string {
	if p.Prefix == "" {
		p.Prefix = "mesos"
	}
	return p.Prefix
}

// GraphDefinition interface for mackerelplugin
func (p *MesosPlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := fmt.Sprintf("%s %s", strings.Title(p.MetricKeyPrefix()), strings.Title(p.Node))
	graphs := map[string]mp.Graphs{
		fmt.Sprintf("%s.cpu", p.Node): {
			Label: fmt.Sprintf("%s CPU", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "cpus_total", Label: "Total"},
				{Name: "cpus_used", Label: "Used"},
			},
		},
		fmt.Sprintf("%s.disk", p.Node): {
			Label: fmt.Sprintf("%s Disk", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "disk_total", Label: "Total"},
				{Name: "disk_used", Label: "Used"},
			},
		},
		fmt.Sprintf("%s.memory", p.Node): {
			Label: fmt.Sprintf("%s Memory", labelPrefix),
			Unit:  mp.UnitBytes,
			Metrics: []mp.Metrics{
				{Name: "mem_total", Label: "Total"},
				{Name: "mem_used", Label: "Used"},
			},
		},
		fmt.Sprintf("%s.resources", p.Node): {
			Label: fmt.Sprintf("%s Resources Usage", labelPrefix),
			Unit:  mp.UnitPercentage,
			Metrics: []mp.Metrics{
				{Name: "cpus_percent", Label: "CPU"},
				{Name: "disk_percent", Label: "Disk"},
				{Name: "mem_percent", Label: "Memory"},
			},
		},
		fmt.Sprintf("%s.tasks", p.Node): {
			Label: fmt.Sprintf("%s Tasks", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "tasks_error", Label: "Error", Diff: true},
				{Name: "tasks_failed", Label: "Failed", Diff: true},
				{Name: "tasks_finished", Label: "Finished", Diff: true},
				{Name: "tasks_killed", Label: "Killed", Diff: true},
				{Name: "tasks_lost", Label: "Lost", Diff: true},
				{Name: "tasks_running", Label: "Running"},
				{Name: "tasks_staging", Label: "Staging"},
				{Name: "tasks_starting", Label: "Starting"},
			},
		},
	}

	switch p.Node {
	case "master":
		graphs["master.agents"] = mp.Graphs{
			Label: fmt.Sprintf("%s Agents", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "slave_registrations", Label: "Registrations", Diff: true},
				{Name: "slave_removals", Label: "removals", Diff: true},
				{Name: "slave_reregistrations", Label: "Re-Registrations", Diff: true},
				{Name: "slave_shutdowns_canceled", Label: "Shutdowns Canceled", Diff: true},
				{Name: "slave_shutdowns_scheduled", Label: "Shutdowns Scheduled", Diff: true},
				{Name: "slaves_active", Label: "Active"},
				{Name: "slaves_connected", Label: "Connected"},
				{Name: "slaves_disconnected", Label: "Disconnected"},
				{Name: "slaves_inactive", Label: "Inactive"},
			},
		}
		graphs["master.event"] = mp.Graphs{
			Label: fmt.Sprintf("%s Event Queue", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "event_queue_dispatches", Label: "Dispatches"},
				{Name: "event_queue_http_requests", Label: "Http Requests"},
				{Name: "event_queue_messages", Label: "Messages"},
			},
		}
		graphs["master.frameworks"] = mp.Graphs{
			Label: fmt.Sprintf("%s Frameworks", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "frameworks_active", Label: "Active"},
				{Name: "frameworks_connected", Label: "Connected"},
				{Name: "frameworks_disconnected", Label: "Disconnected"},
				{Name: "frameworks_inactive", Label: "Inactive"},
				{Name: "outstanding_offers", Label: "Outstanding Offers"},
			},
		}
		graphs["master.messages"] = mp.Graphs{
			Label: fmt.Sprintf("%s Messages", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "dropped_messages", Label: "Dropped", Diff: true},
				{Name: "invalid_framework_to_executor_messages", Label: "Invalid Framework to Executor", Diff: true},
				{Name: "invalid_status_update_acknowledgements", Label: "Invalid Update Acknowledgements", Diff: true},
				{Name: "invalid_status_updates", Label: "Invalid Updates", Diff: true},
				{Name: "valid_framework_to_executor_messages", Label: "Valid Framework to Executor", Diff: true},
				{Name: "valid_status_update_acknowledgements", Label: "Valid Update Acknowledgements", Diff: true},
				{Name: "valid_status_updates", Label: "Valid Updates", Diff: true},
			},
		}
		graphs["master.registrar.latency"] = mp.Graphs{
			Label: fmt.Sprintf("%s Registrar Latency", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "state_fetch_ms", Label: "Read"},
				{Name: "state_store_ms", Label: "Write"},
			},
		}
		graphs["master.registrar.store.latency"] = mp.Graphs{
			Label: fmt.Sprintf("%s Registrar Write Latency", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "max", Label: "Max"},
				{Name: "min", Label: "Min"},
				{Name: "p50", Label: "50 Percentile"},
				{Name: "p90", Label: "90 Percentile"},
				{Name: "p95", Label: "95 Percentile"},
				{Name: "p99", Label: "99 Percentile"},
				{Name: "p999", Label: "99.9 Percentile"},
				{Name: "p9999", Label: "99.99 Percentile"},
			},
		}

	case "slave":
		graphs["slave.executors"] = mp.Graphs{
			Label: fmt.Sprintf("%s Executors", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "executors_registering", Label: "Registering"},
				{Name: "executors_running", Label: "Running"},
				{Name: "executors_terminated", Label: "Terminated", Diff: true},
				{Name: "executors_terminating", Label: "Terminating"},
			},
		}
		graphs["slave.frameworks"] = mp.Graphs{
			Label: fmt.Sprintf("%s Frameworks", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "frameworks_active", Label: "Active"},
			},
		}
		graphs["slave.messages"] = mp.Graphs{
			Label: fmt.Sprintf("%s Messages", labelPrefix),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "invalid_framework_messages", Label: "Invalid Framework", Diff: true},
				{Name: "valid_framework_messages", Label: "Valid Framework", Diff: true},
				{Name: "invalid_status_updates", Label: "Invalid Updates", Diff: true},
				{Name: "valid_status_updates", Label: "Valid Updates", Diff: true},
			},
		}
	}
	return graphs
}

// FetchMetrics interface for mackerelplugin
func (p *MesosPlugin) FetchMetrics() (map[string]float64, error) {
	url := fmt.Sprintf("http://%s:%s/metrics/snapshot", p.Host, p.Port)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data := make(map[string]float64)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return p.parseMetrics(data), nil
}

func (p *MesosPlugin) parseMetrics(data map[string]float64) map[string]float64 {
	metrics := make(map[string]float64)
	for k, v := range data {
		arr := strings.Split(k, "/")
		if arr[0] != "system" {
			metrics[arr[len(arr)-1]] = v
		}
	}
	return metrics
}

// Do the plugin
func Do() {
	optHost := kingpin.Flag("host", "Hostname").Default("localhost").String()
	optPort := kingpin.Flag("port", "Port").Default("5050").String()
	optNode := kingpin.Flag("node", "Node").Required().String()
	optPrefix := kingpin.Flag("metric-key-prefix", "Metric key prefix").Default("").String()
	optTempfile := kingpin.Flag("tempfile", "Temp file name").Default("").String()
	kingpin.Parse()

	plugin := mp.NewMackerelPlugin(&MesosPlugin{
		Host:   *optHost,
		Port:   *optPort,
		Node:   *optNode,
		Prefix: *optPrefix,
	})
	plugin.Tempfile = *optTempfile
	plugin.Run()
}
