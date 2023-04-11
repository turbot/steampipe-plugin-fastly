# Table: fastly_data_center

Data centers in the Fastly network.

## Examples

### List all data centers

```sql
select
  *
from
  fastly_data_center
order by
  name
```

### Get data center by code

```sql
select
  *
from
  fastly_data_center
where
  code = 'BNE'
```

### Data center counts by location group

```sql
select
  location_group,
  count(*)
from
  fastly_data_center
group by
  location_group
order by
  count desc
```
