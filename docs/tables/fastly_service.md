# Table: fastly_service

A Service represents the configuration for a website, app, API, or anything else to be served through Fastly. A Service can have many Versions, through which Backends, Domains, and more can be configured.

## Examples

### List all services the user has access to

```sql
select
  *
from
  fastly_service
```

### List services that are not deleted

```sql
select
  *
from
  fastly_service
where
  deleted is null
```

### Group services by type

```sql
select
  service_type,
  count(*)
from
  fastly_service
group by
  service_type
order by
  count desc
```
