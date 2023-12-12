---
title: "Steampipe Table: fastly_service_domain - Query Fastly Service Domains using SQL"
description: "Allows users to query Fastly Service Domains, specifically the domain names associated with Fastly services, providing insights into content delivery network configurations and potential anomalies."
---

# Table: fastly_service_domain - Query Fastly Service Domains using SQL

Fastly is a cloud computing service provider that offers an edge cloud platform, which is designed to help developers extend their core cloud infrastructure to the edge of the network, closer to users. Fastly's edge cloud platform enhances web and mobile delivery by accelerating dynamic assets and caching static assets. It also provides security services, video & streaming, and cloud networking services.

## Table Usage Guide

The `fastly_service_domain` table provides insights into the domain names associated with Fastly services. As a DevOps engineer, explore domain-specific details through this table, including service IDs, version numbers, and associated metadata. Utilize it to uncover information about domains, such as those associated with specific services, the versions of those services, and the effective management of content delivery networks.

## Examples

### Basic info
Explore the essential details of your Fastly service domains, such as the name, version, and timestamps for creation and updates. This information can help you manage and track changes to your services over time.

```sql+postgres
select
  name,
  service_id,
  service_version,
  comment,
  created_at,
  updated_at
from
  fastly_service_domain;
```

```sql+sqlite
select
  name,
  service_id,
  service_version,
  comment,
  created_at,
  updated_at
from
  fastly_service_domain;
```

### List domains created in the last 30 days
Explore the recent additions to your web service by identifying domains that have been added in the past month. This can be useful for tracking growth, monitoring new domains, and maintaining an up-to-date overview of your service landscape.

```sql+postgres
select
  name,
  service_id,
  service_version,
  comment,
  created_at,
  updated_at
from
  fastly_service_domain
where
  created_at >= now() - interval '30 days';
```

```sql+sqlite
select
  name,
  service_id,
  service_version,
  comment,
  created_at,
  updated_at
from
  fastly_service_domain
where
  created_at >= datetime('now', '-30 days');
```

### List domains that are not deleted
Discover the segments that are actively in use in your Fastly services by identifying domains that have not been deleted. This can help maintain an efficient and streamlined service by focusing resources on active domains.

```sql+postgres
select
  name,
  service_id,
  service_version,
  created_at
from
  fastly_service_domain
where
  deleted_at is null;
```

```sql+sqlite
select
  name,
  service_id,
  service_version,
  created_at
from
  fastly_service_domain
where
  deleted_at is null;
```

### List domains of a particular service
Gain insights into the different domains associated with a specific service, enabling you to monitor service performance and version details over time. This can be particularly useful in managing and troubleshooting service-related issues.

```sql+postgres
select
  d.name as domain_name,
  service_id,
  service_version,
  d.created_at
from
  fastly_service_domain as d,
  fastly_service as s
where
  d.service_id = s.id
  and s.name = 'service-check';
```

```sql+sqlite
select
  d.name as domain_name,
  service_id,
  service_version,
  d.created_at
from
  fastly_service_domain as d,
  fastly_service as s
where
  d.service_id = s.id
  and s.name = 'service-check';
```