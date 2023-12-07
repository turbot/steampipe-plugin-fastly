---
title: "Steampipe Table: fastly_dictionary - Query Fastly Dictionaries using SQL"
description: "Allows users to query Fastly Dictionaries, which are key-value stores associated with a particular Fastly service."
---

# Table: fastly_dictionary - Query Fastly Dictionaries using SQL

Fastly Dictionaries are key-value stores that are associated with a particular Fastly service. They are designed to help you store and manage data that can be referenced from VCL, Fastly's own caching configuration language. These dictionaries can be used to customize how cached content is served, making them a crucial part of Fastly's edge cloud platform.

## Table Usage Guide

The `fastly_dictionary` table provides insights into Fastly Dictionaries, which are key-value stores associated with a particular Fastly service. As a DevOps engineer, you can use this table to manage and manipulate data that can be referenced from VCL, Fastly's own caching configuration language. This table can be particularly useful when you need to customize how cached content is served, enhancing your control over Fastly's edge cloud platform.

## Examples

### Basic info
Explore which Fastly dictionaries have been created with write-only permissions. This can help you understand how your data is being secured and managed within your Fastly service.

```sql+postgres
select
  id,
  name,
  service_id,
  service_version,
  write_only,
  created_at
from
  fastly_dictionary;
```

```sql+sqlite
select
  id,
  name,
  service_id,
  service_version,
  write_only,
  created_at
from
  fastly_dictionary;
```

### List dictionaries created in the last 30 days
Explore which dictionaries have been created recently to keep track of changes and updates. This is useful for maintaining an up-to-date understanding of the system and identifying any unexpected modifications.

```sql+postgres
select
  id,
  name,
  service_id,
  service_version,
  write_only,
  created_at
from
  fastly_dictionary
where
  created_at >= now() - interval '30 days';
```

```sql+sqlite
select
  id,
  name,
  service_id,
  service_version,
  write_only,
  created_at
from
  fastly_dictionary
where
  created_at >= datetime('now', '-30 days');
```

### List dictionaries that have not been deleted
Explore which dictionaries within your Fastly service are still active and have not been deleted. This can be useful for managing your data resources and ensuring they are up-to-date.

```sql+postgres
select
  id,
  name,
  service_id,
  service_version,
  write_only,
  created_at
from
  fastly_dictionary
where
  deleted_at is null;
```

```sql+sqlite
select
  id,
  name,
  service_id,
  service_version,
  write_only,
  created_at
from
  fastly_dictionary
where
  deleted_at is null;
```

### List write-only dictionaries
Explore which dictionaries in your Fastly service are set to write-only. This can help identify areas where data is being stored but not read, potentially highlighting inefficiencies or security concerns.

```sql+postgres
select
  id,
  name,
  service_id,
  service_version,
  write_only,
  created_at
from
  fastly_dictionary
where
  write_only;
```

```sql+sqlite
select
  id,
  name,
  service_id,
  service_version,
  write_only,
  created_at
from
  fastly_dictionary
where
  write_only = 1;
```