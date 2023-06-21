# Table: fastly_ip_range

IP ranges in the Fastly network.

## Examples

### Basic info

```sql
select
  cidr,
  version
from
  fastly_ip_range;
```

### List IP ranges within a particular CIDR range

```sql
select
  cidr,
  version
from
  fastly_ip_range
where
  cidr << '199.0.0.0/8';
```

### Show IPv6 ranges

```sql
select
  cidr,
  version
from
  fastly_ip_range
where
  version = 6;
```
