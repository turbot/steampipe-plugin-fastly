---
title: "Steampipe Table: fastly_condition - Query Fastly Conditions using SQL"
description: "Allows users to query Fastly Conditions, specifically the condition name, service id, type, and statement, providing insights into Fastly CDN configuration."
---

# Table: fastly_condition - Query Fastly Conditions using SQL

Fastly Conditions are logical expressions that specify when a particular behavior should be triggered in Fastly's Content Delivery Network (CDN) service. Conditions are used to control how and when certain actions are executed within Fastly. They are associated with specific Fastly services and can be used in various configurations like caching, routing, and response manipulation.

## Table Usage Guide

The `fastly_condition` table provides insights into the conditions configured within Fastly Content Delivery Network (CDN). As a site reliability engineer, explore condition-specific details through this table, including its type, statement, and associated service. Utilize it to uncover information about conditions, such as those used for caching policies, routing rules, or response manipulations.

## Examples

### Basic info
Explore the fundamental details of specific conditions within your Fastly service to understand their priority and creation time. This could be particularly useful for auditing and managing your service configurations effectively.

```sql
select
  name,
  service_id,
  service_version,
  type,
  priority,
  created_at
from
  fastly_condition;
```

### List conditions that are not deleted
Explore conditions within your Fastly service configuration that are still active and have not been deleted. This allows you to review and manage your current settings and priorities effectively.

```sql
select
  name,
  service_id,
  service_version,
  type,
  priority,
  created_at
from
  fastly_condition
where
  deleted_at is null;
```

### List conditions that are of 'CACHE' type
Explore which conditions in your Fastly configuration are set to 'CACHE' type. This can help in managing cache behavior and troubleshooting caching issues.

```sql
select
  name,
  service_id,
  service_version,
  type,
  priority,
  created_at
from
  fastly_condition
where
  type = 'CACHE';
```

### List conditions that are of high priority
Explore which conditions within the Fastly service are marked as high priority. This can be useful in identifying areas that require immediate attention or action.

```sql
select
  name,
  service_id,
  service_version,
  type,
  priority,
  created_at
from
  fastly_condition
where
  priority = 1;
```