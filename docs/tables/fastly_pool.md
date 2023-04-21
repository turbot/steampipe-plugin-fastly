# Table: fastly_pool

A pool is responsible for balancing requests among a group of servers.

## Examples

### Basic info

```sql
select
  id,
  name,
  connect_timeout,
  created_at,
  healthcheck,
  service_id,
  service_version
from
  fastly_pool;
```

### List pools that are not deleted

```sql
select
  id,
  name,
  connect_timeout,
  created_at,
  healthcheck,
  service_id,
  service_version
from
  fastly_pool
where
  deleted_at is null;
```

### List random pools

```sql
select
  id,
  name,
  connect_timeout,
  created_at,
  healthcheck,
  service_id,
  service_version
from
  fastly_pool
where
  pool_type = 'random';
```

### List pools of the active versions

```sql
select
  id,
  name,
  connect_timeout,
  p.created_at,
  healthcheck,
  p.service_id,
  service_version
from
  fastly_pool as p,
  fastly_service_version as v
where
  p.service_id = v.service_id
  and p.service_version = v.number
  and v.active;
```