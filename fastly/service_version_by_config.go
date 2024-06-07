package fastly

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fastly/go-fastly/v8/fastly"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Usage:
// This function can be used as a parent where Service ID and Service Version is required for making an API call.
// 1. Check if Service details have been configured in the connection config or by environment variables:
//   - If yes, stream list items as configured in the SPC file.
//   - If Service ID is specified but Service Version is not, make an API call to get the active version.
//   - If the active version is not found, make an API call to get the latest version.
//
// 2. Check if Service details have NOT been configured in the connection config or by environment variables:
//   - Make an API call to list all available service versions for all services and stream list those.
func listServiceVersionsByConfig(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	// Check if Service details has been configured in connection config
	if conn.ServiceID != "" {
		res, err := configServiceVersionHydrate(ctx, d, h)
		if err != nil {
			return nil, err
		}
		items := res.([]*fastly.Version)
		for _, item := range items {
			d.StreamListItem(ctx, item)
		}
		return nil, nil
	}

	return listServicesVersions(ctx, d, h)

}

// We are not directly streaming the list of items in the "fastly_acl" table
// because this is a special case involving parent hydrate chaining.
// Since Steampipe does not yet support parent hydrate chaining, we return the result as required
// and use it accordingly.
func configServiceVersionHydrate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	serviceID, serviceVersion := "", ""

	sc, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("configServiceVersionHydrate", "connection_error", err)
		return nil, err
	}
	serviceID = sc.ServiceID
	serviceVersion = sc.ServiceVersion

	// This check is required while invoking this function from from the table "fastly_acl"
	if serviceID == "" {
		return nil, nil
	}

	// set active version if version is not provided in config
	if serviceID != "" && serviceVersion == "" {
		version, err := getActiveVersion(sc.Client, serviceID, d)
		if err != nil {
			plugin.Logger(ctx).Error("configServiceVersionHydrate.getActiveVersion", "api_error", err)
			return nil, err
		}

		serviceVersion = fmt.Sprint(*version)

		// set the latest version if there is no active version available for the service
		if serviceVersion == "" {
			latestVersion, err := getLatestVersion(sc.Client, serviceID, d)
			if err != nil {
				plugin.Logger(ctx).Error("configServiceVersionHydrate.getLatestVersion", "api_error", err)
				return nil, err
			}
			serviceVersion = fmt.Sprint(latestVersion.Number)
		}

	}

	num, err := strconv.Atoi(serviceVersion)
	if err != nil {
		plugin.Logger(ctx).Error("configServiceVersionHydrate", "error_in_parsing_service_version", err)
		return nil, err
	}

	res := &fastly.Version{
		ServiceID: serviceID,
		Number:    num,
	}

	var versionList []*fastly.Version

	versionList = append(versionList, res)

	return versionList, nil
}
