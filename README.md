# Provisioning Scalable Observability Workspace

<p align="center">
  <a href="https://github.com/qclaogui/codelab-monitoring/actions/workflows/ci.yml">
    <img src="https://github.com/qclaogui/codelab-monitoring/actions/workflows/ci.yml/badge.svg">
  </a>
  <a href="https://github.com/qclaogui/codelab-monitoring/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/qclaogui/codelab-monitoring.svg" alt="License">
  </a>
  <a href="https://github.com/qclaogui/codelab-monitoring/tags">
    <img src="https://img.shields.io/github/last-commit/qclaogui/codelab-monitoring" alt="GitHub Last Commit">
  </a>
</p>

## Grafana LGTMP Stack (Loki Grafana Tempo Mimir Pyroscope)

Base on Flow mode of Grafana Agent.

[Grafana Agent Configuration Generator](https://github.com/grafana/agent-configurator) a tool allows for easy configuration of Grafana Agents Flow system

## Dependencies

This project uses [bingo](https://github.com/bwplotka/bingo) (located in [.bingo/](.bingo/)), a tool to automate the versioning of Go packages.

Run `make install-build-deps` to install dependencies tools.

## Docker Compose

These instructions will get you through the deploying samples with Docker Compose.

***Prerequisites:*** Make sure that you have Docker and Docker Compose installed

> NOTE:
> `include` is available in Docker Compose version 2.20 and later, and Docker Desktop version 4.22 and later.

### [Monolithic mode (单体模式)](./docker-compose/monolithic-mode)

- [Logs [Loki]](./docker-compose/monolithic-mode/logs)
- [Traces [Tempo]](./docker-compose/monolithic-mode/traces)
- [Metrics [Mimir]](./docker-compose/monolithic-mode/metrics)
- [Profiles [Pyroscope]](./docker-compose/monolithic-mode/profiles)
- [All In One [Loki + Tempo + Mimir + Pyroscope]](./docker-compose/monolithic-mode/all-in-one)

### [Read-Write mode (读写模式)](./docker-compose/read-write-mode)

- [Logs [Loki (Read + Write + Backend)]](./docker-compose/read-write-mode/logs)
- Traces
- [Metrics [Mimir (Read + Write + Backend)]](./docker-compose/read-write-mode/metrics)
- Profiles

### [Microservices mode (微服务模式)](./docker-compose/microservices-mode)

- [Logs [Loki (Query-Frontend + Querier + Ruler + Distributor + Ingester)]](./docker-compose/microservices-mode/logs)
- [Traces [Tempo (Query-Frontend + Querier + Distributor + Ingester + Compactor)]](./docker-compose/microservices-mode/traces)
- [Metrics [Mimir (Query-Frontend + Querier + Store-Gateway + Distributor + Ingester + Compactor)]](./docker-compose/microservices-mode/metrics)
- [Profiles [Pyroscope (Query-Frontend + Query-Scheduler + Querier + Store-Gateway + Distributor + Ingester + Compactor)]](./docker-compose/microservices-mode/profiles)

### Quick Start(docker-compose)

```shell
git clone https://github.com/qclaogui/codelab-monitoring.git && cd "$(basename "$_" .git)"

make up-monolithic-mode-metrics
```

That's it.

Once all containers are up and running you can search for metrics in Grafana.

Navigate to [http://localhost:3000/explore](http://localhost:3000/explore) and select the search tab.

### Clean up(docker-compose)

```shell
make down-monolithic-mode-metrics
```

## Kubernetes

[k3d](https://github.com/k3d-io/k3d) makes it very easy to create single- and multi-node [k3s](https://github.com/rancher/k3s) clusters in docker, e.g. for local development on Kubernetes.

### [Monolithic mode (单体模式)](./kubernetes/monolithic-mode)

- [Logs [Loki]](./kubernetes/monolithic-mode/logs)
- [Traces [Tempo]](./kubernetes/monolithic-mode/traces)
- [Metrics [Mimir]](./kubernetes/monolithic-mode/metrics)
- [Profiles [Pyroscope]](./kubernetes/monolithic-mode/profiles)
- [All In One [Loki + Tempo + Mimir + Pyroscope]](./kubernetes/monolithic-mode/all-in-one)

### [Read-Write mode (读写模式)](./kubernetes/read-write-mode)

- [Logs [Loki (Read + Write + Backend)]](./kubernetes/read-write-mode/logs)
- Traces
- [Metrics [Mimir (Read + Write + Backend)]](./kubernetes/read-write-mode/metrics)
- Profiles

### [Microservices mode (微服务模式)](./kubernetes/microservices-mode)

- [Logs [Loki (Query-Frontend + Querier + Ruler + Distributor + Ingester)]](./kubernetes/microservices-mode/logs)
- [Traces [Tempo (Query-Frontend + Querier + Distributor + Ingester + Compactor)]](./kubernetes/microservices-mode/traces)
- [Metrics [Mimir (Query-Frontend + Querier + Store-Gateway + Distributor + Ingester + Compactor)]](./kubernetes/microservices-mode/metrics)
- [Profiles [Pyroscope (Query-Frontend + Query-Scheduler + Querier + Store-Gateway + Distributor + Ingester + Compactor)]](./kubernetes/microservices-mode/profiles)

### Quick Start(kubernetes)

Install dependencies tools

```shell
git clone https://github.com/qclaogui/codelab-monitoring.git && cd "$(basename "$_" .git)"

make install-build-deps
```

Create a cluster and mapping the ingress port 80 to localhost:8080

```shell
make cluster
```

Deploy manifests

```shell
make deploy-monolithic-mode-logs
```

Once all containers are up and running you can search for logs in Grafana.

Navigate to [http://localhost:8080/explore](http://localhost:8080/explore) and select the search tab.

### Clean up(kubernetes)

```shell
make delete-monolithic-mode-logs
```

help

```shell
❯ make help

Usage:
  make <target>

Dependencies
  install-build-deps                        Install dependencies tools

Lint & fmt
  copyright                                 Add Copyright header to .go files.
  fmt                                       Uses Grafana Agent to fmt the river config

Docker compose
  up-monolithic-mode-metrics                Run monolithic-mode metrics
  up-monolithic-mode-logs                   Run monolithic-mode logs
  up-monolithic-mode-traces                 Run monolithic-mode traces
  up-monolithic-mode-profiles               Run monolithic-mode profiles
  up-monolithic-mode-all-in-one             Run monolithic-mode all-in-one
  up-read-write-mode-metrics                Run read-write-mode metrics
  up-read-write-mode-logs                   Run read-write-mode logs
  up-microservices-mode-metrics             Run microservices-mode metrics
  up-microservices-mode-logs                Run microservices-mode logs
  up-microservices-mode-traces              Run microservices-mode traces
  up-microservices-mode-profiles            Run microservices-mode profiles

Kubernetes
  cluster                                   Create k3s cluster
  clean                                     Clean cluster
  manifests                                 Generates k8s manifests
  deploy-kube-prometheus-stack              Deploy kube-prometheus-stack manifests
  deploy-monolithic-mode-metrics            Deploy monolithic-mode metrics
  deploy-monolithic-mode-logs               Deploy monolithic-mode logs
  deploy-monolithic-mode-profiles           Deploy monolithic-mode profiles
  deploy-monolithic-mode-traces             Deploy monolithic-mode traces
  deploy-monolithic-mode-all-in-one         Deploy monolithic-mode all-in-one
  deploy-read-write-mode-metrics            Deploy read-write-mode metrics
  deploy-read-write-mode-logs               Deploy read-write-mode logs
  deploy-microservices-mode-logs            Deploy microservices-mode logs
  deploy-microservices-mode-metrics         Deploy microservices-mode metrics
  deploy-microservices-mode-profiles        Deploy microservices-mode profiles
  deploy-microservices-mode-traces          Deploy microservices-mode traces

Grafana Agent Integrations
  deploy-memcached                          Deploy integration memcached manifests
  deploy-mysql                              Deploy integration mysql manifests
  deploy-redis                              Deploy integration redis manifests

Build
  generate                                  generate embed deps
  build                                     Build binary for current OS and place it at ./bin/lgtmp_$(GOOS)_$(GOARCH)
  build-all                                 Build binaries for Linux and Mac and place them in dist/

Release
  prepare-release-candidate                 Create release candidate
  prepare-release                           Create release
  print-version                             Prints the upcoming release number

General
  console-token                             Prints the minio-operator console jwt token
  help                                      Display this help. Thanks to https://www.thapaliya.com/en/writings/well-documented-makefiles/

```
