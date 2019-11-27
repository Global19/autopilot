// Code generated by Autopilot. DO NOT EDIT.

package metrics

import (
	"time"

	"github.com/solo-io/autopilot/pkg/metrics"
)

type CanaryDeploymentMetrics interface {
	metrics.Client
	GetMyquery(ctx context.Context,  string) (*metrics.QueryResult, error)
	GetIstioSuccessRate(ctx context.Context, Namespace, Name, Interval string) (*metrics.QueryResult, error)
	GetIstioRequestDuration(ctx context.Context, Namespace, Name, Interval string) (*metrics.QueryResult, error)
	GetEnvoySuccessRate(ctx context.Context, Namespace, Name, Interval string) (*metrics.QueryResult, error)
	GetEnvoyRequestDuration(ctx context.Context, Namespace, Name, Interval string) (*metrics.QueryResult, error)
}

type metricsClient struct {
	metrics.Client
}

func NewMetricsClient(client metrics.Client) *metricsClient {
	return &metricsClient{Client: client}
}

func (c *metricsClient) GetMyquery(ctx context.Context,  string) (*metrics.QueryResult, error) {
	queryTemplate := `sum`
	queryParameters := map[string]string{
	}
	return c.Client.RunQuery(ctx, queryTemplate, queryParameters)
}

func (c *metricsClient) GetIstioSuccessRate(ctx context.Context, Namespace, Name, Interval string) (*metrics.QueryResult, error) {
	queryTemplate := `sum(
		rate(
			istio_requests_total{
				destination_workload_namespace="{{ .Namespace }}",
				destination_workload=~"{{ .Name }}",
				response_code!~"5.*"
			}[{{ .Interval }}]
		)
	) 
	/ 
	sum(
		rate(
			istio_requests_total{
				destination_workload_namespace="{{ .Namespace }}",
				destination_workload=~"{{ .Name }}"
			}[{{ .Interval }}]
		)
	) 
	* 100`
	queryParameters := map[string]string{
	"Namespace": Namespace,
	"Name": Name,
	"Interval": Interval,
	}
	return c.Client.RunQuery(ctx, queryTemplate, queryParameters)
}

func (c *metricsClient) GetIstioRequestDuration(ctx context.Context, Namespace, Name, Interval string) (*metrics.QueryResult, error) {
	queryTemplate := `histogram_quantile(
		0.99,
		sum(
			rate(
				istio_request_duration_seconds_bucket{
					destination_workload_namespace="{{ .Namespace }}",
					destination_workload=~"{{ .Name }}"
				}[{{ .Interval }}]
			)
		) by (le)
	)`
	queryParameters := map[string]string{
	"Namespace": Namespace,
	"Name": Name,
	"Interval": Interval,
	}
	return c.Client.RunQuery(ctx, queryTemplate, queryParameters)
}

func (c *metricsClient) GetEnvoySuccessRate(ctx context.Context, Namespace, Name, Interval string) (*metrics.QueryResult, error) {
	queryTemplate := `sum(
		rate(
			envoy_cluster_upstream_rq{
				kubernetes_namespace="{{ .Namespace }}",
				kubernetes_pod_name=~"{{ .Name }}-[0-9a-zA-Z]+(-[0-9a-zA-Z]+)",
				envoy_response_code!~"5.*"
			}[{{ .Interval }}]
		)
	) 
	/ 
	sum(
		rate(
			envoy_cluster_upstream_rq{
				kubernetes_namespace="{{ .Namespace }}",
				kubernetes_pod_name=~"{{ .Name }}-[0-9a-zA-Z]+(-[0-9a-zA-Z]+)"
			}[{{ .Interval }}]
		)
	) 
	* 100`
	queryParameters := map[string]string{
	"Namespace": Namespace,
	"Name": Name,
	"Interval": Interval,
	}
	return c.Client.RunQuery(ctx, queryTemplate, queryParameters)
}

func (c *metricsClient) GetEnvoyRequestDuration(ctx context.Context, Namespace, Name, Interval string) (*metrics.QueryResult, error) {
	queryTemplate := `histogram_quantile(
		0.99,
		sum(
			rate(
				envoy_cluster_upstream_rq_time_bucket{
					kubernetes_namespace="{{ .Namespace }}",
					kubernetes_pod_name=~"{{ .Name }}-[0-9a-zA-Z]+(-[0-9a-zA-Z]+)"
				}[{{ .Interval }}]
			)
		) by (le)
	)`
	queryParameters := map[string]string{
	"Namespace": Namespace,
	"Name": Name,
	"Interval": Interval,
	}
	return c.Client.RunQuery(ctx, queryTemplate, queryParameters)
}
