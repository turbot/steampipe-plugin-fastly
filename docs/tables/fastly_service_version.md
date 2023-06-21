# Table: fastly_service_version

A Version represents a specific instance of the configuration for a service. A Version can be cloned, locked, activated, or deactivated.

## Examples

### Basic info

```sql
select
  service_id,
  number,
  active,
  created_at,
  locked
from
  fastly_service_version;
```

### List versions created in the last 30 days

```sql
select
  service_id,
  number,
  active,
  created_at,
  locked
from
  fastly_service_version
where
  created_at >= now() - interval '30 days';
```

### List all inactive versions

```sql
select
  service_id,
  number,
  active,
  created_at,
  locked
from
  fastly_service_version
where
  not active;
```

### List all locked versions

```sql
select
  service_id,
  number,
  active,
  created_at,
  locked
from
  fastly_service_version
where
  locked;
```

### List versions that are not deleted

```sql
select
  service_id,
  number,
  active,
  created_at,
  locked
from
  fastly_service_version
where
  deleted_at is null;
```
