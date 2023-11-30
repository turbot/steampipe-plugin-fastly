---
title: "Steampipe Table: fastly_acl - Query Fastly Access Control Lists using SQL"
description: "Allows users to query Fastly Access Control Lists, providing details on the access control rules applied to Fastly services."
---

# Table: fastly_acl - Query Fastly Access Control Lists using SQL

Fastly Access Control Lists (ACLs) are a security feature that allows you to manage client access to your Fastly services. ACLs enable you to create rules that allow or deny requests from specific IP addresses, subnets, or geographical regions. They are essential for maintaining the security and integrity of your Fastly services.

## Table Usage Guide

The `fastly_acl` table provides insights into Access Control Lists within Fastly. As a security or DevOps professional, explore ACL-specific details through this table, including associated services, rules, and IP addresses. Utilize it to uncover information about the access control mechanisms, such as the allowed or denied IP addresses and the corresponding services, ensuring the robust security of your Fastly services.

## Examples

### Basic info
Explore which Access Control Lists (ACLs) have been created or updated in Fastly, a crucial step in managing network access and ensuring optimal security measures are in place.

```sql
select
  id,
  name,
  service_id,
  service_version,
  created_at,
  updated_at
from
  fastly_acl;
```

### List ACLs created in the last 30 days
Explore ACLs that have been established in the past month. This can help you understand recent changes and maintain up-to-date security configurations.

```sql
select
  id,
  name,
  service_id,
  service_version,
  created_at,
  updated_at
from
  fastly_acl
where
  created_at >= now() - interval '30 days';
```

### List ACLs that are not deleted
Discover the segments that have active Access Control Lists (ACLs) in Fastly. This can help in maintaining security by ensuring only authorized users have access to specific services.

```sql
select
  id,
  name,
  service_id,
  service_version,
  created_at,
  updated_at
from
  fastly_acl
where
  deleted_at is null;
```

### List ACLs where the service version is inactive
Explore which Access Control Lists (ACLs) are associated with inactive versions of services. This can be useful in identifying potential security risks or redundant ACLs that need to be updated or removed.

```sql
select
  id,
  name,
  a.service_id,
  service_version,
  a.created_at
from
  fastly_acl as a,
  fastly_service_version as v
where
  a.service_id = v.service_id
  and a.service_version = v.number
  and not v.active;
```