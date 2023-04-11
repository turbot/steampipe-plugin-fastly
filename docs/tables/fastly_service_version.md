# Table: fastly_service_version

A Version represents a specific instance of the configuration for a service. A Version can be cloned, locked, activated, or deactivated.

## Examples

### List all versions of all services

```sql
select
  *
from
  fastly_service_version
```

### List all active versions

```sql
select
  *
from
  fastly_service_version
where
  active
```
