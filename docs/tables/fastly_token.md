# Table: fastly_token

API tokens are unique authentication identifiers that you can create for the users and applications authorized to interact with your account and services.

## Examples

### List all tokens

```sql
select
  *
from
  fastly_token
```

### Tokens expiring in the next 30 days

```sql
select
  name,
  created_at,
  expires_at
from
  fastly_token
where
  expires_at < current_timestamp + interval '30 days'
```

### Tokens with no expiration

```sql
select
  name,
  created_at,
  expires_at
from
  fastly_token
where
  expires_at is null
```

### Oldest tokens

```sql
select
  name,
  created_at,
  date_part('day', now() - created_at) as age_in_days
from
  fastly_token
order by
  created_at
```

### Tokens not used recently

```sql
select
  name,
  last_used_at,
  date_part('day', now() - last_used_at) as last_used_age_in_days
from
  fastly_token
order by
  last_used_at
```

### Tokens that have never been used

```sql
select
  name,
  last_used_at
from
  fastly_token
where
  last_used_at is null
```

### Tokens with access to a given service

```sql
select
  name,
  scopes,
  services
from
  fastly_token
where
  jsonb_array_length(services) = 0
  or services ? '1crAFFWV5PmZEzbiZ9FsJT'
```

### Tokens used from an IP outside expected range

```sql
select
  name,
  last_used_at,
  ip
from
  fastly_token
where
  not (ip << '123.0.0.0/8')
```

### Tokens with full access

```sql
select
  name,
  scopes,
  services
from
  fastly_token
where
  scopes ? 'global'
```
