connection "fastly" {
  plugin = "fastly"

  # api_key(required) - The fastly API Token.
  # Get your API token from Fastly https://docs.fastly.com/en/guides/using-api-tokens
  # Can also be set with the FASTLY_API_KEY environment variable.
  # api_key = "cj9nU-sMOgUmo7FxcZ48tJsofuiVUhai"

  # service_id(required) - Each connection represents a single service in Fastly.
  # Can also be set with the FASTLY_SERVICE_ID environment variable.
  # service_id = "2ctACCWV5PmZGadiS7Ft5T"

  # base_url(optional) - The fastly base URL. By default plugin will use https://api.fastly.com.
  # Can also be set with the FASTLY_API_URL environment variable.
  # base_url = "https://api.fastly.com"

  # service_version(optional) - The fastly service version. By default, the plugin will use the active version of the provided service; if no active version is available, then the plugin will use the latest version.
  # Can also be set with the FASTLY_SERVICE_VERSION environment variable.
  # service_version = "1"
}
