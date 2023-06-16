# Table: fastly_service

A Service represents the configuration for a website, app, API, or anything else to be served through Fastly. A Service can have many Versions, through which Backends, Domains, and more can be configured.

## Examples

### Basic info

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

### List services created in last the 30 days

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
