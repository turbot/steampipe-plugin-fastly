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

### List dictionaries that are not deleted

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

### List private dictionaries

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
