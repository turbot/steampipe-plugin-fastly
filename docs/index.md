---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/fastly.svg"
brand_color: "#FF282D"
display_name: "Fastly"
short_name: "fastly"
description: "Steampipe plugin to query services, acls, domains and more from Fastly."
og_description: "Query Fastly with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/fastly-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Fastly + Steampipe

[Fastly](https://fastly.com) provides an edge cloud platform, including content delivery network (CDN), image optimization, video and streaming, cloud security, and load balancing services.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List your Fastly domains:

```sql
select
  name,
  service_id,
  service_version,
  comment,
  created_at,
  updated_at
from
  fastly_service_domain;
```

```
+----------------+------------------------+-----------------+---------+---------------------------+---------------------------+
| name           | service_id             | service_version | comment | created_at                | updated_at                |
+----------------+------------------------+-----------------+---------+---------------------------+---------------------------+
| testnumbe3.com | FxvqNlNxUKWRWrioX313Q6 | 4               | <null>  | 2023-04-18T21:58:07+05:30 | 2023-04-19T15:42:19+05:30 |
+----------------+------------------------+-----------------+---------+---------------------------+---------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/fastly/tables)**

## Quick start

### Install

Download and install the latest Fastly plugin:

```sh
steampipe plugin install fastly
```

### Credentials

| Item        | Description                                                                                                                                                                                         |
| ----------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | Fastly requires an [API Key](https://docs.fastly.com/en/guides/using-api-tokens) and Service ID for all requests.                                                                                   |
| Permissions | The permission scope of API Keys is set by the Admin at the creation time of the API tokens.                                                                                                        |
| Radius      | Each connection represents a single Fastly Installation.                                                                                                                                            |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/fastly.spc`)<br />2. Credentials specified in environment variables, e.g., `FASTLY_API_KEY` and `FASTLY_SERVICE_ID`. |

### Configuration

Installing the latest fastly plugin will create a config file (`~/.steampipe/config/fastly.spc`) with a single connection named `fastly`:

Configure your account details in `~/.steampipe/config/fastly.spc`:

```hcl
connection "fastly" {
  plugin = "fastly"

  # api_key - The fastly API token. Required.
  # Get your API token from Fastly https://docs.fastly.com/en/guides/using-api-tokens
  # Can also be set with the FASTLY_API_KEY environment variable.
  # api_key = "cj9nU-sMOgUmo7FxcZ48tJsofuiVUhai"

  # service_id - Each connection represents a single service in Fastly. Required.
  # Can also be set with the FASTLY_SERVICE_ID environment variable.
  # service_id = "2ctACCWV5PmZGadiS7Ft5T"

  # base_url - The fastly base URL. By default plugin will use https://api.fastly.com. Optional.
  # Can also be set with the FASTLY_API_URL environment variable.
  # base_url = "https://api.fastly.com"

  # service_version - The fastly service version. By default, the plugin will use the active version of the provided service; if no active version is available, then the plugin will use the latest version. Optional.
  # Can also be set with the FASTLY_SERVICE_VERSION environment variable.
  # service_version = "1"
}
```

Alternatively, you can also use the standard Fastly environment variables to obtain credentials **only if the `api_key`,`service_id`, `base_url` and `service_version` are not specified** in the connection:

```sh
export FASTLY_API_KEY=cj9nU-sMOgUmo7FxcZ48tJsofuiVUhai
export FASTLY_SERVICE_ID=2ctACCWV5PmZGadiS7Ft5T
export FASTLY_API_URL=https://api.fastly.com
export FASTLY_SERVICE_VERSION=1
```


