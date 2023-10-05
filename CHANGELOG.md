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
