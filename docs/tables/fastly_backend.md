# Table: fastly_backend

A backend (also sometimes called an origin server) is a server identified by IP address or hostname, from which Fastly will fetch your content. There can be multiple backends attached to a service, but each backend is specific to one service.

## Examples

### List all backends for all versions of the service

```sql
select
  *
from
  fastly_backend
```

### List all backends for the active service version

```sql
select
  name,
  address,
  port
from
  fastly_backend
where
  service_version in (select number from fastly_service_version where active)
```

### List backends for a given service version

```sql
select
  name,
  address,
  port
from
  fastly_backend
where
  service_version = 2
```
