// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strings"
)

// MetricDefinitionsOperations contains the methods for the MetricDefinitions group.
type MetricDefinitionsOperations interface {
	// List - Lists the metric definitions for the resource.
	List(ctx context.Context, resourceUri string, metricDefinitionsListOptions *MetricDefinitionsListOptions) (*MetricDefinitionCollectionResponse, error)
}

// MetricDefinitionsClient implements the MetricDefinitionsOperations interface.
// Don't use this type directly, use NewMetricDefinitionsClient() instead.
type MetricDefinitionsClient struct {
	*Client
}

// NewMetricDefinitionsClient creates a new instance of MetricDefinitionsClient with the specified values.
func NewMetricDefinitionsClient(c *Client) MetricDefinitionsOperations {
	return &MetricDefinitionsClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *MetricDefinitionsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// List - Lists the metric definitions for the resource.
func (client *MetricDefinitionsClient) List(ctx context.Context, resourceUri string, metricDefinitionsListOptions *MetricDefinitionsListOptions) (*MetricDefinitionCollectionResponse, error) {
	req, err := client.ListCreateRequest(ctx, resourceUri, metricDefinitionsListOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result, err := client.ListHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListCreateRequest creates the List request.
func (client *MetricDefinitionsClient) ListCreateRequest(ctx context.Context, resourceUri string, metricDefinitionsListOptions *MetricDefinitionsListOptions) (*azcore.Request, error) {
	urlPath := "/{resourceUri}/providers/microsoft.insights/metricDefinitions"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceUri)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2018-01-01")
	if metricDefinitionsListOptions != nil && metricDefinitionsListOptions.Metricnamespace != nil {
		query.Set("metricnamespace", *metricDefinitionsListOptions.Metricnamespace)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *MetricDefinitionsClient) ListHandleResponse(resp *azcore.Response) (*MetricDefinitionCollectionResponse, error) {
	result := MetricDefinitionCollectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.MetricDefinitionCollection)
}

// ListHandleError handles the List error response.
func (client *MetricDefinitionsClient) ListHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}