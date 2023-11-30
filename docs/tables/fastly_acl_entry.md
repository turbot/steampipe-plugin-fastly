---
title: "Steampipe Table: fastly_acl_entry - Query Fastly Access Control List Entries using SQL"
description: "Allows users to query Fastly Access Control List Entries, providing detailed information about each entry within an Access Control List (ACL)."
---

# Table: fastly_acl_entry - Query Fastly Access Control List Entries using SQL

Fastly Access Control List Entries are individual rules within an Access Control List (ACL) in Fastly, a cloud computing services provider. These entries determine what traffic is allowed or denied based on the IP address or the subnet. Fastly Access Control List Entries offer granular control over the traffic to your services, enhancing security by blocking or allowing specific IP addresses or subnets.

## Table Usage Guide

The `fastly_acl_entry` table provides insights into individual rules within an Access Control List (ACL) in Fastly. As a security engineer, you can explore details about each ACL entry through this table, including IP addresses, subnet details, and the actions associated with them. Use this table to analyze and manage traffic to your services by blocking or allowing specific IP addresses or subnets.

## Examples

### Basic info
Explore which access control list (ACL) entries have been negated to understand potential vulnerabilities in your network security. This information can be crucial in identifying areas that require immediate attention or improvement.

```sql
select
  id,
  acl_id,
  ip,
  negated,
  service_id,
  created_at
from
  fastly_acl_entry;
```

### List entries created in the last 30 days
Discover the most recent entries to understand your system's activity over the past month. This allows you to stay updated on changes and identify any unusual patterns or anomalies.

```sql
select
  id,
  acl_id,
  ip,
  negated,
  service_id,
  created_at
from
  fastly_acl_entry
where
  created_at >= now() - interval '30 days';
```

### List entries that are not deleted
Uncover the details of active access control list (ACL) entries in your Fastly configuration to maintain the security and access management of your network resources. This query is useful in monitoring the overall health of your ACLs by identifying entries that are currently in effect.

```sql
select
  id,
  acl_id,
  ip,
  negated,
  service_id,
  created_at
from
  fastly_acl_entry
where
  deleted_at is null;
```

### List entries that are negated
Discover the segments that have been negated to understand the impact on your Fastly Access Control List (ACL). This can help pinpoint specific areas requiring attention or modification to enhance your security measures.

```sql
select
  id,
  acl_id,
  ip,
  negated,
  service_id,
  created_at
from
  fastly_acl_entry
where
  negated;
```

### List entries of a particular ACL
Analyze the settings to understand the specific entries within a particular Access Control List (ACL), allowing you to assess the configuration for better security management.

```sql
select
  e.id,
  acl_id,
  ip,
  negated,
  e.service_id,
  e.created_at
from
  fastly_acl_entry as e,
  fastly_acl as a
where
  e.acl_id = a.id
  and name = 'acl_entry';
```