---
title: "Steampipe Table: fastly_service_version - Query Fastly Service Versions using SQL"
description: "Allows users to query Fastly Service Versions, providing detailed insights into the configuration of each version of a Fastly service."
---

# Table: fastly_service_version - Query Fastly Service Versions using SQL

Fastly Service Versions are an integral part of Fastly's edge cloud platform. Each version of a service represents a specific configuration. Users can create new versions, clone existing versions, activate a particular version, and much more. 

## Table Usage Guide

The `fastly_service_version` table provides deep visibility into the versions of Fastly services. If you are a DevOps engineer or an IT administrator, this table can be extremely useful for tracking changes, understanding the configuration of each version, and managing your Fastly services more efficiently. It provides details about each version, including its status, settings, and associated services.

## Examples

### Basic info
Explore active services within Fastly by identifying the services that are currently active, when they were created, and whether they are locked. This query is useful for understanding the status of your services and their history.

```sql+postgres
select
  service_id,
  number,
  active,
  created_at,
  locked
from
  fastly_service_version;
```

```sql+sqlite
select
  service_id,
  number,
  active,
  created_at,
  locked
from
  fastly_service_version;
```

### List versions created in the last 30 days
Explore which service versions have been created in the last 30 days. This is useful for tracking recent changes and updates, determining the active status, and identifying if any versions have been locked.

```sql+postgres
select
  service_id,
  number,
  active,
  created_at,
  locked
from
  fastly_service_version
where
  created_at >= now() - interval '30 days';
```

```sql+sqlite
select
  service_id,
  number,
  active,
  created_at,
  locked
from
  fastly_service_version
where
  created_at >= datetime('now', '-30 days');
```

### List all inactive versions
Explore which service versions are inactive in Fastly, allowing you to understand the history and management of your service versions. This can be particularly useful when reviewing changes or troubleshooting issues.

```sql+postgres
select
  service_id,
  number,
  active,
  created_at,
  locked
from
  fastly_service_version
where
  not active;
```

```sql+sqlite
select
  service_id,
  number,
  active,
  created_at,
  locked
from
  fastly_service_version
where
  active = 0;
```

### List all locked versions
Explore which service versions are currently locked. This could be useful in understanding the status of your services and identifying any potential issues or bottlenecks.

```sql+postgres
select
  service_id,
  number,
  active,
  created_at,
  locked
from
  fastly_service_version
where
  locked;
```

```sql+sqlite
select
  service_id,
  number,
  active,
  created_at,
  locked
from
  fastly_service_version
where
  locked = 1;
```

### List versions that are not deleted
Analyze the settings of your Fastly service to identify all versions that are still active and not deleted, allowing you to manage and streamline your services effectively. This can be particularly useful in maintaining a clean and efficient service environment.

```sql+postgres
select
  service_id,
  number,
  active,
  created_at,
  locked
from
  fastly_service_version
where
  deleted_at is null;
```

```sql+sqlite
select
  service_id,
  number,
  active,
  created_at,
  locked
from
  fastly_service_version
where
  deleted_at is null;
```