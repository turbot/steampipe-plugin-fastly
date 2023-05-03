package fastly

import (
	"context"
	"path"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// shouldIgnoreErrors:: function which returns an ErrorPredicate for Fastly API calls
func shouldIgnoreErrors(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		for _, pattern := range notFoundErrors {
			// handle not found error
			if ok, _ := path.Match(err.Error(), pattern); ok {
				return true
			}
		}
		return false
	}
}

func shouldRetryError(retryErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		for _, pattern := range retryErrors {
			// handle retry error
			if strings.Contains(err.Error(), pattern) {
				return true
			}
		}
		return false
	}
}
