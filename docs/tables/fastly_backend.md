# Table: fastly_backend

A backend (also sometimes called an origin server) is a server identified by IP address or hostname, from which Fastly will fetch your content. There can be multiple backends attached to a service, but each backend is specific to one service.

## Examples

### Basic info

```sql
select
  name,
  address,
  port,
  use_ssl,
  auto_loadbalance,
  created_at,
  service_id,
  service_version
from
  fastly_backend;
```

### List all backends that are not deleted

```sql
select
  name,
  address,
  port,
  use_ssl,
  auto_loadbalance,
  created_at,
  service_id,
  service_version
from
  fastly_backend
where
  deleted_at is null;
```

### List all backends that are using SSL

```sql
select
  name,
  address,
  port,
  use_ssl,
  auto_loadbalance,
  created_at,
  service_id,
  service_version
from
  fastly_backend
where
  use_ssl;
```

### List all backends where auto load balance is enabled

```sql
select
  name,
  address,
  port,
  use_ssl,
  auto_loadbalance,
  created_at,
  service_id,
  service_version
from
  fastly_backend
where
  auto_loadbalance;
```

### List backends for a particular service

```sql
select
  b.name,
  address,
  port,
  use_ssl,
  auto_loadbalance,
  b.created_at,
  service_id,
  service_version
from
  fastly_backend as b,
  fastly_service as s
where
  b.service_id = s.id
  and s.name = 'check-service';
```
