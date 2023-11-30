---
title: "Steampipe Table: fastly_token - Query Fastly API Tokens using SQL"
description: "Allows users to query Fastly API Tokens, providing detailed information about each token, such as its ID, user ID, service ID, and access level."
---

# Table: fastly_token - Query Fastly API Tokens using SQL

Fastly API Tokens are a resource within the Fastly API service. They are used to authenticate and authorize requests made to the Fastly API. Each token is associated with specific user and service IDs, and has a defined access level that determines what actions can be performed using the token.

## Table Usage Guide

The `fastly_token` table provides insights into API tokens within Fastly's API service. As a DevOps engineer, explore token-specific details through this table, including the associated user and service IDs, and the access level of each token. Utilize it to manage and monitor the use of API tokens, ensuring that each token has the appropriate access level for its intended use.

## Examples

### Basic info
Discover the segments that have been created, their expiration date, their last used date and associated user details within Fastly, to better manage and monitor access. This could be particularly useful for enhancing security and ensuring optimal utilization of resources.

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
Explore which tokens have been created in the past 30 days. This can be useful for auditing purposes, allowing you to keep track of newly generated tokens and their associated user activity.

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
Discover the tokens that are due to expire in the next 30 days. This can be useful for proactive management and renewal of these tokens to prevent any service disruptions.

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
Identify all Fastly tokens that have been set to never expire. This can be useful for managing security risks and ensuring appropriate access control.

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
Discover the segments that contain unused tokens, which can be instrumental in identifying potential security risks or optimizing resource allocation. This provides a way to assess your system's efficiency and security by pinpointing unused tokens.

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
Discover the tokens that have access to a specific service. This is useful for managing access control and ensuring only the appropriate tokens have access to certain services.

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
Discover the instances where tokens have been used from an IP address outside of an expected range. This is beneficial in identifying potential security breaches or unauthorized access.

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
Explore which tokens have full access across your network. This can be used to monitor and manage security by identifying potentially risky permissions.

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