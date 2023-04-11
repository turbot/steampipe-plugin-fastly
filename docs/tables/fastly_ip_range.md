# Table: fastly_ip_range

IP ranges in the Fastly network.

## Examples

### List all IP ranges

```sql
select
  *
from
  fastly_ip_range
```

### IP ranges inside a range

```sql
select
  *
from
  fastly_ip_range
where
  cidr << '199.0.0.0/8'
```

### IPv6 ranges

```sql
select
  *
from
  fastly_ip_range
where
  version = 6
```
