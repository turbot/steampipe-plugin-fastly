# Table: fastly_dictionary

A Dictionary is a VCL data table that stores key-value pairs accessible to VCL during request processing.

Note: A `service_id` and `service_version` must be provided in all queries to this table.

## Examples

### List all dictionaries for a service version

```sql
select
  *
from
  fastly_dictionary
where
  service_id = '1crAGGWV3PnZEibiZ9FsJT'
  and service_version = 2
```

### List all dictionaries for all service versions

```sql
select
  *
from
  fastly_service_version as v,
  fastly_dictionary as d
where
  d.service_id = v.service_id
  and d.service_version = v.number
```

### List all key value pairs from all dictionaries

```sql
select
  v.service_id,
  v.number as service_version,
  kv.key,
  kv.value
from
  fastly_service_version as v,
  fastly_dictionary as d,
  jsonb_each(d.data) as kv
where
  d.service_id = v.service_id
  and d.service_version = v.number
```
