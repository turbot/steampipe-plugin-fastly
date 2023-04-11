![image](https://hub.steampipe.io/images/plugins/turbot/fastly-social-graphic.png)

# Fastly Plugin for Steampipe

Use SQL to query instances, domains and more from Fastly.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/fastly)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/fastly/tables)
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-fastly/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install fastly
```

Run a query:

```sql
select
  name,
  service_type,
  active_version
from
  fastly_service
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