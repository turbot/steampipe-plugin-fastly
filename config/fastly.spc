connection "fastly" {
  plugin = "fastly"

  # The fastly API token is required for requests. Required.
  # Get your API token from Fastly https://docs.fastly.com/en/guides/using-api-tokens
  # Can also be set with the FASTLY_API_KEY environment variable.
  # api_key = "cj9nU-sMOgUmo7FxcZ48tJsofuiVUhai"

  # The service ID in Fastly. Optional.
  # If the service version is configured, the service ID must also be configured.
  # Can also be set with the FASTLY_SERVICE_ID environment variable.
  # service_id = "2ctACCWV5PmZGadiS7Ft5T"

  # The fastly base URL is the API server hostname. Optional.
  # It is required if using a private instance of the API and otherwise defaults to the public Fastly production service By default plugin will use https://api.fastly.com.
  # Can also be set with the FASTLY_API_URL environment variable.
  # base_url = "https://api.fastly.com"

  # The fastly service version. By default, the plugin will use the active version of the provided service; if no active version is available, then the plugin will use the latest version. Optional.
  # Can also be set with the FASTLY_SERVICE_VERSION environment variable.
  # service_version = "1"
}
