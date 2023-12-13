![image](https://hub.steampipe.io/images/plugins/turbot/fastly-social-graphic.png)

# Fastly Plugin for Steampipe

Use SQL to query services, acls, domains and more from Fastly.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/fastly)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/fastly/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-fastly/issues)

## Quick start

### Install

Download and install the latest Fastly plugin:

```bash
steampipe plugin install fastly
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/fastly#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/fastly#configuration).

Configure your account details in `~/.steampipe/config/fastly.spc`:

You may specify the API Key and Service ID to authenticate:

```hcl
connection "fastly" {
  plugin     = "fastly"

  # Authentication information
  api_key    = "cj9nU-sMOgUmo7FxcZ48tJsofuiVUhai"
  service_id = "2ctACCWV5PmZGadiS7Ft5T"
}
```

Or through environment variables:

```sh
export FASTLY_API_KEY=cj9nU-sMOgUmo7FxcZ48tJsofuiVUhai
export FASTLY_SERVICE_ID=2ctACCWV5PmZGadiS7Ft5T
```

Run steampipe:

```shell
steampipe query
```

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

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs/steampipe_sqlite/overview) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/overview) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-fastly.git
cd steampipe-plugin-fastly
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/fastly.spc
```

Try it!

```
steampipe query
> .inspect fastly
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Open Source & Contributing

This repository is published under the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) (source code) and [CC BY-NC-ND](https://creativecommons.org/licenses/by-nc-nd/2.0/) (docs) licenses. Please see our [code of conduct](https://github.com/turbot/.github/blob/main/CODE_OF_CONDUCT.md). We look forward to collaborating with you!

[Steampipe](https://steampipe.io) is a product produced from this open source software, exclusively by [Turbot HQ, Inc](https://turbot.com). It is distributed under our commercial terms. Others are allowed to make their own distribution of the software, but cannot use any of the Turbot trademarks, cloud services, etc. You can learn more in our [Open Source FAQ](https://turbot.com/open-source).

## Get Involved

**[Join #steampipe on Slack →](https://turbot.com/community/join)**

Want to help but don't know where to start? Pick up one of the `help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Fastly Plugin](https://github.com/turbot/steampipe-plugin-fastly/labels/help%20wanted)
