# Table: fastly_acl

An access control list or "ACL" specifies individual IP addresses or subnet ranges and can be accessed and used from Fastly VCL.

Note: A `service_id` and `service_version` must be provided in all queries to this table.

## Examples

### List all ACLs for a service version

```sql
select
  *
from
  fastly_acl
where
  service_id = '1crAGGWV3PnZEibiZ9FsJT'
  and service_version = 2
```

### List all ACLs for all service versions

```sql
select
  *
from
  fastly_service_version as v,
  fastly_acl as acl
where
  acl.service_id = v.service_id
  and acl.service_version = v.number
```

### List all ACLs and their entries for all service versions

```sql
select
  *
from
  fastly_service_version as v,
  fastly_acl as acl,
  fastly_acl_entry as e
where
  acl.service_id = v.service_id
  and acl.service_version = v.number
  and e.service_id = v.service_id
  and e.acl_id = acl.id
```
