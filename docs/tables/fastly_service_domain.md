# Table: fastly_service_domain

Domains are used to route requests to your service. You associate your domain names with your origin when provisioning a Fastly service so you can properly route requests to your website and ensure that others cannot serve requests to that domain.

## Examples

### Basic info

```sql
select
  name,
  service_id,
  service_version,
  comment,
  created_at,
  updated_at
from
  fastly_service_domain;
```

### List domains that are not deleted

```sql
select
  name,
  service_id,
  service_version,
  created_at
from
  fastly_service_domain
where
  deleted_at is null;
```

### List domains of a particular service

```sql
select
  d.name,
  service_id,
  service_version,
  d.created_at
from
  fastly_service_domain as d,
  fastly_service as s
where
  d.service_id = s.id
  and s.name = 'service-check';
```
