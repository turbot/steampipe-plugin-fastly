---
title: "Steampipe Table: fastly_backend - Query Fastly Backends using SQL"
description: "Allows users to query Fastly Backends, specifically providing detailed information about the configuration and status of backends in a Fastly service."
---

# Table: fastly_backend - Query Fastly Backends using SQL

A Fastly Backend represents an origin server to which Fastly will connect to fetch content. Backends are defined in the configuration for a Fastly service and are associated with a specific service and version. They provide the critical link between Fastly and the origin server, defining how Fastly connects to the origin to fetch the content it will cache and deliver.

## Table Usage Guide

The `fastly_backend` table provides insights into the configuration and status of backends within Fastly's edge cloud platform. As an infrastructure engineer, explore backend-specific details through this table, including the service it is associated with, its version, and its connection details. Utilize it to uncover information about backends, such as the connection timeouts, max connections, and the healthcheck parameters.

## Examples

### Basic info
Explore the configuration of your backend servers on Fastly to understand their security settings, load balancing status, and creation dates. This can help you assess your current setup and identify areas for potential improvement or troubleshooting.

```sql+postgres
select
  name,
  address,
  port,
  use_ssl,
  auto_loadbalance,
  created_at,
  service_id,
  service_version
from
  fastly_backend;
```

```sql+sqlite
select
  name,
  address,
  port,
  use_ssl,
  auto_loadbalance,
  created_at,
  service_id,
  service_version
from
  fastly_backend;
```

### List all backends that are not deleted
Discover the segments that contain all active backends in your Fastly service. This information is useful for maintaining an overview of your current operational infrastructure and understanding the configuration of each backend.

```sql+postgres
select
  name,
  address,
  port,
  use_ssl,
  auto_loadbalance,
  created_at,
  service_id,
  service_version
from
  fastly_backend
where
  deleted_at is null;
```

```sql+sqlite
select
  name,
  address,
  port,
  use_ssl,
  auto_loadbalance,
  created_at,
  service_id,
  service_version
from
  fastly_backend
where
  deleted_at is null;
```

### List all backends that are using SSL
Discover the segments that are utilizing SSL for secure communication, which can be crucial for maintaining data privacy and enhancing security measures. This can be particularly useful in identifying and managing secure backends in your network infrastructure.

```sql+postgres
select
  name,
  address,
  port,
  use_ssl,
  auto_loadbalance,
  created_at,
  service_id,
  service_version
from
  fastly_backend
where
  use_ssl;
```

```sql+sqlite
select
  name,
  address,
  port,
  use_ssl,
  auto_loadbalance,
  created_at,
  service_id,
  service_version
from
  fastly_backend
where
  use_ssl = 1;
```

### List all backends where auto load balance is enabled
Discover the segments that are auto load balanced in your backends. This is useful for identifying areas where load distribution is automated, allowing for efficient resource management.

```sql+postgres
select
  name,
  address,
  port,
  use_ssl,
  auto_loadbalance,
  created_at,
  service_id,
  service_version
from
  fastly_backend
where
  auto_loadbalance;
```

```sql+sqlite
select
  name,
  address,
  port,
  use_ssl,
  auto_loadbalance,
  created_at,
  service_id,
  service_version
from
  fastly_backend
where
  auto_loadbalance = 1;
```

### List backends for a particular service
Uncover the details of specific backends associated with a particular service in Fastly. This is useful for understanding the configuration and settings of the backends, including whether they use SSL, auto load balance, their address, and port, providing insights into the service's operation and performance.

```sql+postgres
select
  b.name,
  address,
  port,
  use_ssl,
  auto_loadbalance,
  b.created_at,
  service_id,
  service_version
from
  fastly_backend as b,
  fastly_service as s
where
  b.service_id = s.id
  and s.name = 'check-service';
```

```sql+sqlite
select
  b.name,
  address,
  port,
  use_ssl,
  auto_loadbalance,
  b.created_at,
  service_id,
  service_version
from
  fastly_backend as b,
  fastly_service as s
where
  b.service_id = s.id
  and s.name = 'check-service';
```