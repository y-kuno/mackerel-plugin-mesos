package mpmesos

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMetricsMaster(t *testing.T) {
	str := `{
  "frameworks/jenkins/messages_processed": 6141678,
  "frameworks/jenkins/messages_received": 6141678,
  "master/cpus_percent": 0.454687500000005,
  "master/cpus_total": 32,
  "master/cpus_used": 14.5500000000001,
  "master/disk_percent": 0,
  "master/disk_total": 1269540,
  "master/disk_used": 0,
  "master/dropped_messages": 0,
  "master/elected": 1,
  "master/event_queue_dispatches": 17,
  "master/event_queue_http_requests": 0,
  "master/event_queue_messages": 0,
  "master/frameworks_active": 1,
  "master/frameworks_connected": 1,
  "master/frameworks_disconnected": 0,
  "master/frameworks_inactive": 0,
  "master/invalid_framework_to_executor_messages": 0,
  "master/invalid_status_update_acknowledgements": 0,
  "master/invalid_status_updates": 0,
  "master/mem_percent": 0.130365744563718,
  "master/mem_total": 116256,
  "master/mem_used": 15155.7999999996,
  "master/messages_authenticate": 0,
  "master/messages_deactivate_framework": 0,
  "master/messages_decline_offers": 50738088,
  "master/messages_exited_executor": 0,
  "master/messages_framework_to_executor": 0,
  "master/messages_kill_task": 492369,
  "master/messages_launch_tasks": 529176,
  "master/messages_reconcile_tasks": 0,
  "master/messages_register_framework": 32,
  "master/messages_register_slave": 63,
  "master/messages_reregister_framework": 0,
  "master/messages_reregister_slave": 47,
  "master/messages_resource_request": 0,
  "master/messages_revive_offers": 0,
  "master/messages_status_update": 1035469,
  "master/messages_status_update_acknowledgement": 1035469,
  "master/messages_unregister_framework": 0,
  "master/messages_unregister_slave": 0,
  "master/outstanding_offers": 0,
  "master/recovery_slave_removals": 0,
  "master/slave_registrations": 32,
  "master/slave_removals": 31,
  "master/slave_reregistrations": 3,
  "master/slave_shutdowns_canceled": 0,
  "master/slave_shutdowns_scheduled": 21,
  "master/slaves_active": 4,
  "master/slaves_connected": 4,
  "master/slaves_disconnected": 0,
  "master/slaves_inactive": 0,
  "master/task_failed/source_slave/reason_command_executor_failed": 19687,
  "master/task_failed/source_slave/reason_memory_limit": 13343,
  "master/task_killed/source_master/reason_framework_removed": 797,
  "master/task_killed/source_slave/reason_executor_unregistered": 10,
  "master/task_lost/source_master/reason_invalid_offers": 959,
  "master/task_lost/source_master/reason_slave_disconnected": 102,
  "master/task_lost/source_master/reason_slave_removed": 136,
  "master/tasks_error": 0,
  "master/tasks_failed": 34790,
  "master/tasks_finished": 0,
  "master/tasks_killed": 493156,
  "master/tasks_lost": 1197,
  "master/tasks_running": 33,
  "master/tasks_staging": 0,
  "master/tasks_starting": 0,
  "master/uptime_secs": 79859726.82003,
  "master/valid_framework_to_executor_messages": 0,
  "master/valid_status_update_acknowledgements": 1035469,
  "master/valid_status_updates": 1035469,
  "registrar/queued_operations": 0,
  "registrar/registry_size_bytes": 909,
  "registrar/state_fetch_ms": 9.39904,
  "registrar/state_store_ms": 6.155008,
  "registrar/state_store_ms/count": 2,
  "registrar/state_store_ms/max": 7.742208,
  "registrar/state_store_ms/min": 6.155008,
  "registrar/state_store_ms/p50": 6.948608,
  "registrar/state_store_ms/p90": 7.583488,
  "registrar/state_store_ms/p95": 7.662848,
  "registrar/state_store_ms/p99": 7.726336,
  "registrar/state_store_ms/p999": 7.7406208,
  "registrar/state_store_ms/p9999": 7.74204928,
  "system/cpus_total": 2,
  "system/load_15min": 0,
  "system/load_1min": 0.08,
  "system/load_5min": 0.04,
  "system/mem_free_bytes": 4287860736,
  "system/mem_total_bytes": 7724630016
}`
	var p MesosPlugin
	data := make(map[string]float64)
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		t.Fatal(err)
	}

	metrics := p.parseMetrics(data)
	assert.Equal(t, metrics["slave_registrations"], float64(32))
	assert.Equal(t, metrics["event_queue_dispatches"], float64(17))
	assert.Equal(t, metrics["outstanding_offers"], float64(0))
	assert.Equal(t, metrics["dropped_messages"], float64(0))
	assert.Equal(t, metrics["state_store_ms"], float64(6.155008))
	assert.Equal(t, metrics["p9999"], float64(7.74204928))
}

func TestParseMetricsSlave(t *testing.T) {
	str := `{
  "slave/cpus_percent": 0.71875,
  "slave/cpus_total": 8,
  "slave/cpus_used": 5.75,
  "slave/disk_percent": 0,
  "slave/disk_total": 317417,
  "slave/disk_used": 0,
  "slave/executors_registering": 0,
  "slave/executors_running": 10,
  "slave/executors_terminated": 55176,
  "slave/executors_terminating": 0,
  "slave/frameworks_active": 1,
  "slave/invalid_framework_messages": 0,
  "slave/invalid_status_updates": 82,
  "slave/mem_percent": 0.215765207817231,
  "slave/mem_total": 29064,
  "slave/mem_used": 6271,
  "slave/recovery_errors": 0,
  "slave/registered": 1,
  "slave/tasks_failed": 2146,
  "slave/tasks_finished": 0,
  "slave/tasks_killed": 52947,
  "slave/tasks_lost": 0,
  "slave/tasks_running": 10,
  "slave/tasks_staging": 0,
  "slave/tasks_starting": 0,
  "slave/uptime_secs": 32198455.5071821,
  "slave/valid_framework_messages": 0,
  "slave/valid_status_updates": 110183,
  "system/cpus_total": 8,
  "system/load_15min": 2.99,
  "system/load_1min": 9.92,
  "system/load_5min": 4.23,
  "system/mem_free_bytes": 14588207104,
  "system/mem_total_bytes": 31550550016
}`

	var p MesosPlugin
	data := make(map[string]float64)
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		t.Fatal(err)
	}

	metrics := p.parseMetrics(data)
	assert.Equal(t, metrics["cpus_percent"], float64(0.71875))
	assert.Equal(t, metrics["cpus_total"], float64(8))
	assert.Equal(t, metrics["cpus_used"], float64(5.75))
	assert.Equal(t, metrics["executors_registering"], float64(0))
	assert.Equal(t, metrics["executors_running"], float64(10))
	assert.Equal(t, metrics["executors_terminated"], float64(55176))
	assert.Equal(t, metrics["executors_terminating"], float64(0))
}
