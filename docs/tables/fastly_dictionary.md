# Table: fastly_dictionary

A Dictionary is a VCL data table that stores key-value pairs accessible to VCL during request processing.

## Examples

### Basic info

```sql
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

```sql
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

### List dictionaries that have not been deleted

```sql
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

```sql
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
