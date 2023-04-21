# Table: fastly_acl

An access control list or "ACL" specifies individual IP addresses or subnet ranges and can be accessed and used from Fastly VCL.

## Examples

### Basic info

```sql
select
  id,
  name,
  service_id,
  service_version,
  created_at,
  updated_at
from
  fastly_acl;
```

### List ACLs that are deleted

```sql
select
  id,
  name,
  service_id,
  service_version,
  created_at,
  updated_at
from
  fastly_acl
where
  deleted_at is null;
```

### List ACLs where the service version is inactive

```sql
select
  id,
  name,
  a.service_id,
  service_version,
  a.created_at
from
  fastly_acl as a,
  fastly_service_version as v
where
  a.service_id = v.service_id
  and a.service_version = v.number
  and not v.active;
```
