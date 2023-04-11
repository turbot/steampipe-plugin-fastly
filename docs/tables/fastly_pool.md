# Table: fastly_pool

A pool is responsible for balancing requests among a group of servers.

Note: A `service_id` and `service_version` must be provided in all queries to this table.

## Examples

### List all pools for a service version

```sql
select
  *
from
  fastly_pool
where
  service_id = '1crAGGWV3PnZEibiZ9FsJT'
  and service_version = 2
```

### All pools for active service versions

```sql
select
  c.*
from
  fastly_service as s,
  fastly_pool as c
where
  s.id = c.service_id
  and s.active_version = c.service_version
```
