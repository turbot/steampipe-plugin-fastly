# Table: fastly_condition

Conditions are used to control whether logic defined in configured VCL objects is applied for a particular client request. A condition contains a VCL conditional expression that evaluates to either true or false and is used to determine whether the condition is met. The type of the condition determines where it is executed and the VCL variables that can be evaluated as part of the conditional logic.

## Examples

### Basic info

```sql
select
  name,
  service_id,
  service_version,
  type,
  priority,
  created_at
from
  fastly_condition;
```

### List conditions that are not deleted

```sql
select
  name,
  service_id,
  service_version,
  type,
  priority,
  created_at
from
  fastly_condition
where
  deleted_at is null;
```

### List conditions that are of 'CACHE' type

```sql
select
  name,
  service_id,
  service_version,
  type,
  priority,
  created_at
from
  fastly_condition
where
  type = 'CACHE';
```

### List conditions that are of high priority

```sql
select
  name,
  service_id,
  service_version,
  type,
  priority,
  created_at
from
  fastly_condition
where
  priority = 1;
```
