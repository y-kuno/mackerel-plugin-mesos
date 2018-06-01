# mackerel-plugin-mesos [![Build Status](https://travis-ci.org/y-kuno/mackerel-plugin-mesos.svg?branch=master)](https://travis-ci.org/y-kuno/mackerel-plugin-mesos)

Mesos plugin for mackerel.io agent. This repository releases an artifact to Github Releases, which satisfy the format for mkr plugin installer.

## Install

```shell
mkr plugin install y-kuno/mackerel-plugin-mesos
```

## Synopsis

```shell
mackerel-plugin-mesos [--host=<host>] [--port=<port>] [--node=<node>] [--metric-key-prefix=<prefix>]
```

### Master Node

```
[plugin.metrics.mesos]
command = "/path/to/mackerel-plugin-mesos --node=master"
```

### Slave Node

```
[plugin.metrics.mesos]
command = "/path/to/mackerel-plugin-mesos --node=salve"
```

## Documents

* [Monitoring](http://mesos.apache.org/documentation/latest/monitoring/)
