# Table: fastly_data_center

Data centers in the Fastly network.

## Examples

### Basic info

```sql
select
  name,
  code,
  location_group,
  longitude,
  latitude,
  shield
from
  fastly_data_center;
```

### Show data center detail for a particular code

```sql
select
  name,
  code,
  location_group,
  longitude,
  latitude,
  shield
from
  fastly_data_center
where
  code = 'BNE';
```
### List data centers in Europe

```sql
select
  name,
  code,
  location_group,
  longitude,
  latitude,
  shield
from
  fastly_data_center
where
  location_group = 'Europe';
```

### Show data center count by location group

```sql
select
  location_group,
  count(*)
from
  fastly_data_center
group by
  location_group
order by
  count desc;
```
