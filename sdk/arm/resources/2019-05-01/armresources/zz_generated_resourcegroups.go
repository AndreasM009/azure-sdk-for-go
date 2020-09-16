// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// ResourceGroupsOperations contains the methods for the ResourceGroups group.
type ResourceGroupsOperations interface {
	// CheckExistence - Checks whether a resource group exists.
	CheckExistence(ctx context.Context, resourceGroupName string) (*http.Response, error)
	// CreateOrUpdate - Creates or updates a resource group.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, parameters ResourceGroup) (*ResourceGroupResponse, error)
	// BeginDelete - When you delete a resource group, all of its resources are also deleted. Deleting a resource group deletes all of its template deployments and currently stored operations.
	BeginDelete(ctx context.Context, resourceGroupName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// ExportTemplate - Captures the specified resource group as a template.
	ExportTemplate(ctx context.Context, resourceGroupName string, parameters ExportTemplateRequest) (*ResourceGroupExportResultResponse, error)
	// Get - Gets a resource group.
	Get(ctx context.Context, resourceGroupName string) (*ResourceGroupResponse, error)
	// List - Gets all the resource groups for a subscription.
	List(resourceGroupsListOptions *ResourceGroupsListOptions) ResourceGroupListResultPager
	// Update - Resource groups can be updated through a simple PATCH operation to a group address. The format of the request is the same as that for creating a resource group. If a field is unspecified, the current value is retained.
	Update(ctx context.Context, resourceGroupName string, parameters ResourceGroupPatchable) (*ResourceGroupResponse, error)
}

// ResourceGroupsClient implements the ResourceGroupsOperations interface.
// Don't use this type directly, use NewResourceGroupsClient() instead.
type ResourceGroupsClient struct {
	*Client
	subscriptionID string
}

// NewResourceGroupsClient creates a new instance of ResourceGroupsClient with the specified values.
func NewResourceGroupsClient(c *Client, subscriptionID string) ResourceGroupsOperations {
	return &ResourceGroupsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *ResourceGroupsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CheckExistence - Checks whether a resource group exists.
func (client *ResourceGroupsClient) CheckExistence(ctx context.Context, resourceGroupName string) (*http.Response, error) {
	req, err := client.CheckExistenceCreateRequest(ctx, resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusNoContent, http.StatusNotFound) {
		return nil, client.CheckExistenceHandleError(resp)
	}
	return resp.Response, nil
}

// CheckExistenceCreateRequest creates the CheckExistence request.
func (client *ResourceGroupsClient) CheckExistenceCreateRequest(ctx context.Context, resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-05-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// CheckExistenceHandleError handles the CheckExistence error response.
func (client *ResourceGroupsClient) CheckExistenceHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// CreateOrUpdate - Creates or updates a resource group.
func (client *ResourceGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, parameters ResourceGroup) (*ResourceGroupResponse, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	result, err := client.CreateOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ResourceGroupsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, parameters ResourceGroup) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-05-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ResourceGroupsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*ResourceGroupResponse, error) {
	result := ResourceGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ResourceGroup)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ResourceGroupsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// Delete - When you delete a resource group, all of its resources are also deleted. Deleting a resource group deletes all of its template deployments and currently stored operations.
func (client *ResourceGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string) (*HTTPPollerResponse, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.DeleteHandleError(resp)
	}
	result, err := client.DeleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("ResourceGroupsClient.Delete", "", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *ResourceGroupsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ResourceGroupsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *ResourceGroupsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-05-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// DeleteHandleResponse handles the Delete response.
func (client *ResourceGroupsClient) DeleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// DeleteHandleError handles the Delete error response.
func (client *ResourceGroupsClient) DeleteHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// ExportTemplate - Captures the specified resource group as a template.
func (client *ResourceGroupsClient) ExportTemplate(ctx context.Context, resourceGroupName string, parameters ExportTemplateRequest) (*ResourceGroupExportResultResponse, error) {
	req, err := client.ExportTemplateCreateRequest(ctx, resourceGroupName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ExportTemplateHandleError(resp)
	}
	result, err := client.ExportTemplateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExportTemplateCreateRequest creates the ExportTemplate request.
func (client *ResourceGroupsClient) ExportTemplateCreateRequest(ctx context.Context, resourceGroupName string, parameters ExportTemplateRequest) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/exportTemplate"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-05-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// ExportTemplateHandleResponse handles the ExportTemplate response.
func (client *ResourceGroupsClient) ExportTemplateHandleResponse(resp *azcore.Response) (*ResourceGroupExportResultResponse, error) {
	result := ResourceGroupExportResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ResourceGroupExportResult)
}

// ExportTemplateHandleError handles the ExportTemplate error response.
func (client *ResourceGroupsClient) ExportTemplateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// Get - Gets a resource group.
func (client *ResourceGroupsClient) Get(ctx context.Context, resourceGroupName string) (*ResourceGroupResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *ResourceGroupsClient) GetCreateRequest(ctx context.Context, resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-05-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *ResourceGroupsClient) GetHandleResponse(resp *azcore.Response) (*ResourceGroupResponse, error) {
	result := ResourceGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ResourceGroup)
}

// GetHandleError handles the Get error response.
func (client *ResourceGroupsClient) GetHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// List - Gets all the resource groups for a subscription.
func (client *ResourceGroupsClient) List(resourceGroupsListOptions *ResourceGroupsListOptions) ResourceGroupListResultPager {
	return &resourceGroupListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupsListOptions)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ResourceGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ResourceGroupListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *ResourceGroupsClient) ListCreateRequest(ctx context.Context, resourceGroupsListOptions *ResourceGroupsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	if resourceGroupsListOptions != nil && resourceGroupsListOptions.Filter != nil {
		query.Set("$filter", *resourceGroupsListOptions.Filter)
	}
	if resourceGroupsListOptions != nil && resourceGroupsListOptions.Top != nil {
		query.Set("$top", strconv.FormatInt(int64(*resourceGroupsListOptions.Top), 10))
	}
	query.Set("api-version", "2019-05-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *ResourceGroupsClient) ListHandleResponse(resp *azcore.Response) (*ResourceGroupListResultResponse, error) {
	result := ResourceGroupListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ResourceGroupListResult)
}

// ListHandleError handles the List error response.
func (client *ResourceGroupsClient) ListHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// Update - Resource groups can be updated through a simple PATCH operation to a group address. The format of the request is the same as that for creating a resource group. If a field is unspecified, the current value is retained.
func (client *ResourceGroupsClient) Update(ctx context.Context, resourceGroupName string, parameters ResourceGroupPatchable) (*ResourceGroupResponse, error) {
	req, err := client.UpdateCreateRequest(ctx, resourceGroupName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateHandleError(resp)
	}
	result, err := client.UpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateCreateRequest creates the Update request.
func (client *ResourceGroupsClient) UpdateCreateRequest(ctx context.Context, resourceGroupName string, parameters ResourceGroupPatchable) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-05-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// UpdateHandleResponse handles the Update response.
func (client *ResourceGroupsClient) UpdateHandleResponse(resp *azcore.Response) (*ResourceGroupResponse, error) {
	result := ResourceGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ResourceGroup)
}

// UpdateHandleError handles the Update error response.
func (client *ResourceGroupsClient) UpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}