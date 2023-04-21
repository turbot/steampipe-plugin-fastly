# Table: fastly_acl_entry

An ACL entry holds an individual IP address or subnet range and is a member of an ACL. ACL entries are versionless, which means they can be created, modified, or deleted without activating a new version of your service.

## Examples

### Basic info

```sql
select
  id,
  acl_id,
  ip,
  negated,
  service_id,
  created_at
from
  fastly_acl_entry;
```

### List entries which are negated

```sql
select
  id,
  acl_id,
  ip,
  negated,
  service_id,
  created_at
from
  fastly_acl_entry
where
  negated;
```

### List entries of a particular ACL

```sql
select
  e.id,
  acl_id,
  ip,
  negated,
  e.service_id,
  e.created_at
from
  fastly_acl_entry as e,
  fastly_acl as a
where
  e.acl_id = a.id
  and name = 'acl_entry';
```
