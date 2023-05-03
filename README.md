![image](https://hub.steampipe.io/images/plugins/turbot/fastly-social-graphic.png)

# Fastly Plugin for Steampipe

Use SQL to query services, acls, domains and more from Fastly.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/fastly)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/fastly/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-fastly/issues)

## Quick start

Download and install the latest Fastly plugin:

```bash
steampipe plugin install fastly
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/fastly#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/fastly#configuration).

### Configuring Fastly Credentials

Configure your account details in `~/.steampipe/config/fastly.spc`:

You may specify the API Key and Service ID to authenticate:

- `api_key`: The fastly API Token.
- `service_id`: The fastly Service ID.

```hcl
connection "fastly" {
  plugin     = "fastly"
  api_key    = "cj9nU-sMOgUmo7FxcZ48tJsofuiVUhai"
  service_id = "2ctACCWV5PmZGadiS7Ft5T"
}
```

or you may specify the API Key, Service ID, Base URL and Service Version to authenticate:

- `api_key`: The fastly API Token.
- `service_id`: The fastly Service ID.
- `base_url`: The fastly base URL.
- `service_version`: The fastly Service version.

```hcl
connection "fastly" {
  plugin          = "fastly"
  api_key         = "cj9nU-sMOgUmo7FxcZ48tJsofuiVUhai"
  service_id      = "2ctACCWV5PmZGadiS7Ft5T"
  base_url        = "https://api.fastly.com"
  service_version = "1"
}
```

or through environment variables

```sh
export FASTLY_API_KEY="cj9nU-sMOgUmo7FxcZ48tJsofuiVUhai"
export FASTLY_SERVICE_ID="2ctACCWV5PmZGadiS7Ft5T"
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

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-fastly/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Fastly Plugin](https://github.com/turbot/steampipe-plugin-fastly/labels/help%20wanted)
