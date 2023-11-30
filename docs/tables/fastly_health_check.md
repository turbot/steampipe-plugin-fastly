---
title: "Steampipe Table: fastly_health_check - Query Fastly Health Checks using SQL"
description: "Allows users to query Fastly Health Checks, providing insights into the health status of Fastly services and potential issues."
---

# Table: fastly_health_check - Query Fastly Health Checks using SQL

A Fastly Health Check is a service within Fastly that monitors the status of your services and servers. It provides a way to determine if your services are functioning correctly and if there are any issues that need to be addressed. Fastly Health Checks help you stay informed about the health and performance of your services and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `fastly_health_check` table provides insights into the health status of Fastly services. As an operations engineer, explore service-specific details through this table, including health check status, response timeout, and associated metadata. Utilize it to uncover information about the health and performance of your services, helping to identify potential issues and take appropriate actions.

## Examples

### Basic info
Explore the configuration of health checks in Fastly to understand the frequency of checks and their creation time. This is useful to assess the performance and reliability of your services.

```sql
select
  name,
  service_id,
  service_version,
  method,
  host,
  path,
  check_interval,
  created_at
from
  fastly_health_check;
```

### List health checks created in the last 7 days
Explore recent health checks by identifying those created within the last week. This can be useful in maintaining system health and identifying any recent changes or issues.

```sql
select
  name,
  service_id,
  service_version,
  method,
  host,
  path,
  check_interval,
  created_at
from
  fastly_health_check
where
  created_at >= now() - interval '7 days';
```

### List health checks that are not deleted
Uncover the details of active health checks within your Fastly services. This can help you manage and monitor the performance and health of your services, ensuring they are functioning optimally.

```sql
select
  name,
  service_id,
  service_version,
  method,
  host,
  path,
  check_interval,
  created_at
from
  fastly_health_check
where
  deleted_at is null;
```

### Show health check details for a particular host
Explore the health check details for a specific host to gain insights into its service ID, version, method, and other key information. This is useful for assessing the health and performance of a particular host within the Fastly service.

```sql
select
  name,
  service_id,
  service_version,
  method,
  host,
  path,
  check_interval,
  created_at
from
  fastly_health_check
where
  host = 'health.com';
```

### List health checks where threshold is less than 3
Determine the areas in which health checks have a threshold less than 3 to review the configuration for potential vulnerabilities or issues. This is useful for maintaining optimal system health and performance.

```sql
select
  name,
  service_id,
  service_version,
  method,
  host,
  path,
  check_interval,
  created_at
from
  fastly_health_check
where
  threshold < 3;
```

### List health checks where the service version is inactive
Determine the areas in which health checks are being conducted on inactive service versions, allowing for the identification of potential issues or areas of improvement in your Fastly services. This can be particularly useful in maintaining efficient service operations and ensuring all versions are properly monitored.

```sql
select
  name,
  c.service_id,
  service_version,
  method,
  host,
  path,
  check_interval,
  c.created_at
from
  fastly_health_check as c,
  fastly_service_version as v
where
  c.service_id = v.service_id
  and c.service_version = v.number
  and not v.active;
```