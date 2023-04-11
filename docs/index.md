---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/fastly.svg"
brand_color: "#eb423b"
display_name: "Fastly"
short_name: "fastly"
description: "Steampipe plugin to query services, data centers, tokens and more from Fastly."
og_description: "Query Fastly with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/fastly-social-graphic.png"
---

# Fastly + Steampipe

[Fastly](https://fastly.com) provides an edge cloud platform, including content delivery network (CDN), image optimization, video and streaming, cloud security, and load balancing services.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List service versions in your Fastly account:

```sql
select
  name,
  service_type,
  active_version
from
  fastly_service
```

```
+------------+--------------+----------------+
| name       | service_type | active_version |
+------------+--------------+----------------+
| Infinity   | vcl          | 2              |
| Legacy Web | vcl          | 1              |
+------------+--------------+----------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/fastly/tables)**

## Get started

### Install

Download and install the latest Fastly plugin:

```bash
steampipe plugin install fastly
```

### Credentials

No credentials are required.

### Configuration

Installing the latest fastly plugin will create a config file (`~/.steampipe/config/fastly.spc`) with a single connection named `fastly`:

```hcl
connection "fastly" {
  plugin = "fastly"
  api_key = "cj9nU-sMOgUmo7FxcZ48tJsofuiVUhai"
}
```

- `api_key` - API key from Fastly.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-fastly
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
