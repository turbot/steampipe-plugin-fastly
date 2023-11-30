---
title: "Steampipe Table: fastly_service - Query OCI Fastly Services using SQL"
description: "Allows users to query Fastly Services, specifically providing detailed information about the services like ID, name, customer ID, etc. It also includes details about the version of the services."
---

# Table: fastly_service - Query OCI Fastly Services using SQL

Fastly service is a flexible, secure, and scalable platform designed to deliver fast and reliable online experiences. It provides edge cloud platform, web application firewall, bot management, and other services to businesses. It enables developers to build, secure, and deliver digital experiences at the edge of the internet.

## Table Usage Guide

The `fastly_service` table provides insights into Fastly services within Oracle Cloud Infrastructure (OCI). As a DevOps engineer, you can explore service-specific details through this table, including service ID, name, customer ID, and version details. Utilize it to uncover information about services, such as their active versions, configurations, and associated customer details.

## Examples

### Basic info
Explore the basic information of your Fastly services to understand the active versions, their creation dates, and types. This can help in managing and tracking the progress and status of your services.

```sql
select
  id,
  name,
  active_version,
  comment,
  created_at,
  type
from
  fastly_service;
```

### List services created in the last 30 days
Explore which services have been activated in the recent past. This could be useful to track new additions and assess the growth and expansion of your application or platform.

```sql
select
  id,
  name,
  active_version,
  comment,
  created_at,
  type
from
  fastly_service
where
  created_at >= now() - interval '30 days';
```

### List services that have not been updated in the last 90 days
Discover the segments that have been inactive for the last 90 days. This could be useful in identifying potential areas of neglect or inactivity within your services.

```sql
select
  id,
  name,
  active_version,
  comment,
  updated_at,
  type
from
  fastly_service
where
  updated_at < now() - interval '90 days';
```

### List services that are not deleted
Identify active services in your system by checking which ones have not been marked as deleted. This can help maintain an accurate overview of your current operational services.

```sql
select
  id,
  name,
  active_version,
  comment,
  created_at,
  type
from
  fastly_service
where
  deleted_at is null;
```

### List wasm type services
Explore which services are of the 'wasm' type within your Fastly configuration. This query is useful for identifying these specific services and understanding their distribution within your setup.

```sql
select
  id,
  name,
  active_version,
  comment,
  created_at,
  type
from
  fastly_service
where
  type = 'wasm';
```