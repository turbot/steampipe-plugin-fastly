# Table: fastly_acl_entry

An ACL entry holds an individual IP address or subnet range and is a member of an ACL. ACL entries are versionless, which means they can be created, modified, or deleted without activating a new version of your service.

Note: A `acl_id` must be provided in all queries to this table.

## Examples

### List all ACL entries for a given ACL

```sql
select
  *
from
  fastly_acl_entry
where
  service_id = '1crADDWV5PmZEabiZ9FsJT'
  and acl_id = '6AjT0uIOCxWS6b4hB3FSWT'
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
