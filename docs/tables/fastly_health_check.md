# Table: fastly_health_check

Health checks monitor the status of your hosts. Fastly performs health checks on your origin server based on the Check frequency setting you select in the Create a new health check page and the packaged offering you may have purchased. The Check frequency setting you select specifies approximately how many requests per minute Fastly POPs are checked to see if they pass. There is roughly one health check per Fastly POP per period. Any checks that pass will be reported as "healthy".

## Examples

### Basic info

```sql
select
  name,
  service_id,
  service_version,
  method,
  host,
  path,
  check_interval,
  created_at
from
  fastly_health_check;
```

### List health checks that are not deleted

```sql
select
  name,
  service_id,
  service_version,
  method,
  host,
  path,
  check_interval,
  created_at
from
  fastly_health_check
where
  deleted_at is null;
```

### Show health check details for a particular host

```sql
select
  name,
  service_id,
  service_version,
  method,
  host,
  path,
  check_interval,
  created_at
from
  fastly_health_check
where
  host = 'health.com';
```

### List health checks where threshold is less than 3

```sql
select
  name,
  service_id,
  service_version,
  method,
  host,
  path,
  check_interval,
  created_at
from
  fastly_health_check
where
  threshold < 3;
```

### List health checks where the service version is inactive

```sql
select
  name,
  c.service_id,
  service_version,
  method,
  host,
  path,
  check_interval,
  c.created_at
from
  fastly_health_check as c,
  fastly_service_version as v
where
  c.service_id = v.service_id
  and c.service_version = v.number
  and not v.active;
```
