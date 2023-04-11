# Table: fastly_condition

Conditions are used to control whether logic defined in configured VCL objects is applied for a particular client request. A condition contains a VCL conditional expression that evaluates to either true or false and is used to determine whether the condition is met. The type of the condition determines where it is executed and the VCL variables that can be evaluated as part of the conditional logic.

## Examples

### List all conditions for a service version

```sql
select
  *
from
  fastly_condition
where
  service_id = '1crAGGWV3PnZEibiZ9FsJT'
  and service_version = 2
```

### All conditions for active service versions

```sql
select
  c.*
from
  fastly_service as s,
  fastly_condition as c
where
  s.id = c.service_id
  and s.active_version = c.service_version
```

### All conditions for all service versions

```sql
select
  c.*
from
  fastly_service as s,
  fastly_service_version as v,
  fastly_condition as c
where
  s.id = v.service_id
  and s.id = c.service_id
  and v.number = c.service_version
```

```sql
select
  c.*
from
  fastly_service as s,
  fastly_service_version as v,
  fastly_condition as c
where
  s.id = v.service_id
  and v.service_id = c.service_id
  and v.number = c.service_version
```

Unfortunately, this approach through the JSONB object does not work (yet, as of Jul 2021):

```sql
select
  c.*
from
  fastly_service as s,
  jsonb_array_elements(s.versions) as v,
  fastly_condition as c
where
  s.id = c.service_id
  and (v->'Number')::int = c.service_version
```
