# Read-Write mode (读写模式) - Metrics

The read-write mode provides an alternative to monolithic and microservices modes.

In read-write mode, components are grouped into three services to ease the operational overhead whilst still allowing scale to be tuned separately on the read and write paths. The services group the components as follows:

- read
  - query-frontend
  - querier
- backend
  - store-gateway
  - compactor
  - (optional) ruler
  - (optional) alertmanager
  - (optional) query-scheduler
  - (optional) overrides-exporter
- write
  - distributor
  - ingester

Similar to the other modes, each Grafana Mimir process is invoked with its `-target` parameter set to the specific service: `-target=read`, `-target=write`, or `-target=backend`.

## Diagram

The below diagram describes the various components of this deployment, and how data flows between them.

```mermaid
flowchart LR
    A  -->|writes| GW  -->|writes| D   -->|writes| I  -->|writes| M
    G -.->|reads | GW -.->|reads | QF -.->|reads | Q -.->|reads | SG -.->|reads| M

    subgraph Minio
        M{"Object Storage"}
    end
    subgraph Agent["Grafana Alloy"]
        A("alloy")
    end
    subgraph Grafana
        G("grafana")
    end
    subgraph Gateway["Load Balancer"]
        GW{"Nginx"}
    end

    subgraph MimirWrite["mimir -target=write"]
        D("distributor")
        I("ingester")
    end

    subgraph MimirBackend["mimir -target=backend"]
        SG("store-gateway")
        C("compactor")
        
        C  -->|writes| M
        C -.->|reads | M
    end

    subgraph MimirRead["mimir -target=read"]
        Q("Querier") -.->|reads| I
        QF("query-frontend")
    end
```

## Quick Start

```shell
git clone https://github.com/qclaogui/codelab-monitoring.git && cd "$(basename "$_" .git)"

make up-read-write-mode-metrics
```

That's it.

Once all containers are up and running you can search for logs in Grafana.

Navigate to [http://localhost:3000/explore](http://localhost:3000/explore) and select the search tab.

## Clean up

```shell
make down-read-write-mode-metrics
```
