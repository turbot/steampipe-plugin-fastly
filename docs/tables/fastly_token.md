# Table: fastly_token

API tokens are unique authentication identifiers that you can create for the users and applications authorized to interact with your account and services.

## Examples

### Basic info

```sql
select
  id,
  name,
  created_at,
  expires_at,
  ip,
  last_used_at,
  user_id
from
  fastly_token;
```

### List Tokens created in the last 30 days

```sql
select
  id,
  name,
  created_at,
  expires_at,
  ip,
  last_used_at,
  user_id
from
  fastly_token
where
  created_at >= now() - interval '30 days';
```

### List Tokens expiring in the next 30 days

```sql
select
  id,
  name,
  created_at,
  expires_at,
  ip,
  last_used_at,
  user_id
from
  fastly_token
where
  expires_at < current_timestamp + interval '30 days';
```

### List Tokens that will never expire

```sql
select
  id,
  name,
  created_at,
  expires_at,
  ip,
  last_used_at,
  user_id
from
  fastly_token
where
  expires_at is null;
```

### List Tokens that have never been used

```sql
select
  id,
  name,
  created_at,
  expires_at,
  ip,
  last_used_at,
  user_id
from
  fastly_token
where
  last_used_at is null;
```

### List Tokens with access to a given service

```sql
select
  id,
  name,
  scopes,
  services
from
  fastly_token
where
  jsonb_array_length(services) = 0
  or services ? '1crAFFWV5PmZEzbiZ9FsJT';
```

### List Tokens used from an IP outside the expected range

```sql
select
  id,
  name,
  last_used_at,
  ip,
  user_id
from
  fastly_token
where
  not (ip << '123.0.0.0/8');
```

### List Tokens with full access

```sql
select
  name,
  created_at,
  expires_at,
  ip,
  last_used_at,
  user_id
from
  fastly_token
where
  scopes ? 'global';
```
