## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#31](https://github.com/turbot/steampipe-plugin-fastly/pull/31))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#31](https://github.com/turbot/steampipe-plugin-fastly/pull/31))

## v0.2.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#24](https://github.com/turbot/steampipe-plugin-fastly/pull/24))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#24](https://github.com/turbot/steampipe-plugin-fastly/pull/24))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-fastly/blob/main/docs/LICENSE). ([#24](https://github.com/turbot/steampipe-plugin-fastly/pull/24))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#23](https://github.com/turbot/steampipe-plugin-fastly/pull/23))

## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#11](https://github.com/turbot/steampipe-plugin-fastly/pull/11))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#7](https://github.com/turbot/steampipe-plugin-fastly/pull/7))
- Recompiled plugin with Go version `1.21`. ([#7](https://github.com/turbot/steampipe-plugin-fastly/pull/7))

## v0.0.1 [2023-06-23]

_What's new?_

- New tables added
  - [fastly_acl](https://hub.steampipe.io/plugins/turbot/fastly/tables/fastly_acl)
  - [fastly_acl_entry](https://hub.steampipe.io/plugins/turbot/fastly/tables/fastly_acl_entry)
  - [fastly_backend](https://hub.steampipe.io/plugins/turbot/fastly/tables/fastly_backend)
  - [fastly_condition](https://hub.steampipe.io/plugins/turbot/fastly/tables/fastly_condition)
  - [fastly_data_center](https://hub.steampipe.io/plugins/turbot/fastly/tables/fastly_data_center)
  - [fastly_dictionary](https://hub.steampipe.io/plugins/turbot/fastly/tables/fastly_dictionary)
  - [fastly_health_check](https://hub.steampipe.io/plugins/turbot/fastly/tables/fastly_health_check)
  - [fastly_ip_range](https://hub.steampipe.io/plugins/turbot/fastly/tables/fastly_ip_range)
  - [fastly_pool](https://hub.steampipe.io/plugins/turbot/fastly/tables/fastly_pool)
  - [fastly_service](https://hub.steampipe.io/plugins/turbot/fastly/tables/fastly_service)
  - [fastly_service_domain](https://hub.steampipe.io/plugins/turbot/fastly/tables/fastly_service_domain)
  - [fastly_service_version](https://hub.steampipe.io/plugins/turbot/fastly/tables/fastly_service_version)
  - [fastly_token](https://hub.steampipe.io/plugins/turbot/fastly/tables/fastly_token)
