package fastly

import (
	"context"

	"github.com/fastly/go-fastly/v3/fastly"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableFastlyBackend(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "fastly_backend",
		Description: "Backends for the service version.",
		List: &plugin.ListConfig{
			KeyColumns:    plugin.OptionalColumns([]string{"service_id", "service_version"}),
			ParentHydrate: hydrateServiceVersion,
			//ParentHydrate: listServiceVersion,
			Hydrate: listBackend,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getBackend,
			KeyColumns: plugin.AllColumns([]string{"service_version", "name"}),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the backend."},
			{Name: "address", Type: proto.ColumnType_STRING, Description: "A hostname, IPv4, or IPv6 address for the backend. This is the preferred way to specify the location of your backend."},
			{Name: "port", Type: proto.ColumnType_INT, Description: "Port on which the backend server is listening for connections from Fastly. Setting port to 80 or 443 will also set use_ssl automatically (to false and true respectively), unless explicitly overridden by setting use_ssl in the same request."},
			{Name: "use_ssl", Type: proto.ColumnType_BOOL, Description: "Whether or not to require SSL for connections to this backend."},
			// Other columns
			{Name: "auto_loadbalance", Type: proto.ColumnType_BOOL, Description: "Whether or not this backend should be automatically load balanced. If true, all backends with this setting that don't have a request_condition will be selected based on their weight."},
			{Name: "between_bytes_timeout", Type: proto.ColumnType_INT, Description: "Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, the response received so far will be considered complete and the fetch will end."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "A freeform descriptive note."},
			{Name: "connect_timeout", Type: proto.ColumnType_INT, Description: "Maximum duration in milliseconds to wait for a connection to this backend to be established. If exceeded, the connection is aborted and a synthethic 503 response will be presented instead."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the backend was created."},
			{Name: "deleted_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the backend was deleted."},
			{Name: "first_byte_timeout", Type: proto.ColumnType_INT, Description: "Maximum duration in milliseconds to wait for the server response to begin after a TCP connection is established and the request has been sent. If exceeded, the connection is aborted and a synthethic 503 response will be presented instead."},
			{Name: "healthcheck", Type: proto.ColumnType_STRING, Transform: transform.FromField("HealthCheck"), Description: "The name of the healthcheck to use with this backend."},
			{Name: "hostname", Type: proto.ColumnType_STRING, Description: "The hostname of the backend. May be used as an alternative to address to set the backend location."},
			{Name: "max_conn", Type: proto.ColumnType_INT, Description: "Maximum number of concurrent connections this backend will accept."},
			{Name: "max_tls_version", Type: proto.ColumnType_STRING, Description: "Maximum allowed TLS version on connections to this backend."},
			{Name: "min_tls_version", Type: proto.ColumnType_STRING, Description: "Minimum allowed TLS version on connections to this backend."},
			{Name: "override_host", Type: proto.ColumnType_STRING, Description: "If set, will replace the client-supplied HTTP Host header on connections to this backend. Applied after VCL has been processed, so this setting will take precedence over changing bereq.http.Host in VCL."},
			{Name: "request_condition", Type: proto.ColumnType_STRING, Description: "Name of a Condition, which if satisfied, will select this backend during a request. If set, will override any auto_loadbalance setting. By default, the first backend added to a service is selected for all requests."},
			{Name: "service_id", Type: proto.ColumnType_STRING, Description: "Alphanumeric string identifying the service."},
			{Name: "service_version", Type: proto.ColumnType_INT, Description: "Integer identifying a service version."},
			{Name: "shield", Type: proto.ColumnType_STRING, Description: "Data center POP code of the data center to use as a shield."},
			{Name: "ssl_ca_cert", Type: proto.ColumnType_STRING, Description: "CA certificate attached to origin."},
			{Name: "ssl_cert_hostname", Type: proto.ColumnType_STRING, Description: "Overrides ssl_hostname, but only for cert verification. Does not affect SNI at all."},
			{Name: "ssl_check_cert", Type: proto.ColumnType_BOOL, Description: "Be strict on checking SSL certs."},
			{Name: "ssl_ciphers", Type: proto.ColumnType_STRING, Description: "List of OpenSSL ciphers to support for connections to this origin. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic 503 error response will be generated."},
			{Name: "ssl_client_cert", Type: proto.ColumnType_STRING, Description: "Client certificate attached to origin."},
			{Name: "ssl_client_key", Type: proto.ColumnType_STRING, Description: "Client key attached to origin."},
			{Name: "ssl_hostname", Type: proto.ColumnType_STRING, Description: "Use ssl_cert_hostname and ssl_sni_hostname to configure certificate validation."},
			{Name: "ssl_sni_hostname", Type: proto.ColumnType_STRING, Description: "Overrides ssl_hostname, but only for SNI in the handshake. Does not affect cert validation at all."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp (UTC) of when the backend was updated."},
			{Name: "weight", Type: proto.ColumnType_INT, Description: "Weight used to load balance this backend against others. May be any positive integer. If auto_loadbalance is true, the chance of this backend being selected is equal to its own weight over the sum of all weights for backends that have auto_loadbalance set to true."},
		},
	}
}

func listBackend(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, _, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_backend.listBackend", "connection_error", err)
		return nil, err
	}
	version := h.Item.(*fastly.Version)
	items, err := conn.ListBackends(&fastly.ListBackendsInput{ServiceID: version.ServiceID, ServiceVersion: version.Number})
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_backend.listBackend", "query_error", err)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getBackend(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, serviceID, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_backend.getBackend", "connection_error", err)
		return nil, err
	}
	serviceVersion := int(d.KeyColumnQuals["service_version"].GetInt64Value())
	name := d.KeyColumnQuals["name"].GetStringValue()
	input := fastly.GetBackendInput{ServiceID: serviceID, ServiceVersion: serviceVersion, Name: name}
	result, err := conn.GetBackend(&input)
	if err != nil {
		plugin.Logger(ctx).Error("fastly_service_backend.getBackend", "query_error", err, "input", input)
		return nil, err
	}
	return result, nil
}
