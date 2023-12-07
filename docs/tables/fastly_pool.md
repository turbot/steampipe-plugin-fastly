---
title: "Steampipe Table: fastly_pool - Query Fastly Pools using SQL"
description: "Allows users to query Fastly Pools, providing details about each pool's configuration, health, and associated backend servers."
---

# Table: fastly_pool - Query Fastly Pools using SQL

Fastly Pools are a key component in Fastly's Content Delivery Network (CDN) services. They allow users to group backend servers together, which Fastly then uses to balance load and route traffic. Pools are also used to manage failover scenarios, where traffic can be rerouted to backup servers if the primary servers become unavailable.

## Table Usage Guide

The `fastly_pool` table provides insights into Fastly Pools within the Fastly CDN services. As a DevOps engineer or site reliability engineer, explore pool-specific details through this table, including configuration settings, health status, and associated backend servers. Utilize it to understand the load balancing and failover strategies of your Fastly CDN, and to ensure optimal performance and reliability of your web services.

## Examples

### Basic info
Gain insights into the health and configuration of your Fastly service pools, including their creation date, version, and connection timeout settings. This can be useful in monitoring and maintaining optimal service performance.

```sql+postgres
select
  id,
  name,
  connect_timeout,
  created_at,
  healthcheck,
  service_id,
  service_version
from
  fastly_pool;
```

```sql+sqlite
select
  id,
  name,
  connect_timeout,
  created_at,
  healthcheck,
  service_id,
  service_version
from
  fastly_pool;
```

### List pools that are not deleted
Explore which service pools are still active and not marked for deletion. This can assist in managing resources and maintaining an efficient service structure.

```sql+postgres
select
  id,
  name,
  connect_timeout,
  created_at,
  healthcheck,
  service_id,
  service_version
from
  fastly_pool
where
  deleted_at is null;
```

```sql+sqlite
select
  id,
  name,
  connect_timeout,
  created_at,
  healthcheck,
  service_id,
  service_version
from
  fastly_pool
where
  deleted_at is null;
```

### List random pools
Analyze the settings to understand the configuration and health status of various service pools, particularly those of the 'random' type. This can help in managing resources effectively, identifying potential issues, and optimizing service performance.

```sql+postgres
select
  id,
  name,
  connect_timeout,
  created_at,
  healthcheck,
  service_id,
  service_version
from
  fastly_pool
where
  pool_type = 'random';
```

```sql+sqlite
select
  id,
  name,
  connect_timeout,
  created_at,
  healthcheck,
  service_id,
  service_version
from
  fastly_pool
where
  pool_type = 'random';
```

### List pools of the active versions
Analyze the settings to understand the active versions of your service by identifying their associated pools. This information can help optimize your service's performance and resource allocation.

```sql+postgres
select
  id,
  name,
  connect_timeout,
  p.created_at,
  healthcheck,
  p.service_id,
  service_version
from
  fastly_pool as p,
  fastly_service_version as v
where
  p.service_id = v.service_id
  and p.service_version = v.number
  and v.active;
```

```sql+sqlite
select
  p.id,
  p.name,
  p.connect_timeout,
  p.created_at,
  p.healthcheck,
  p.service_id,
  p.service_version
from
  fastly_pool as p,
  fastly_service_version as v
where
  p.service_id = v.service_id
  and p.service_version = v.number
  and v.active = 1;
```