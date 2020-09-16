// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// SSHPublicKeysOperations contains the methods for the SSHPublicKeys group.
type SSHPublicKeysOperations interface {
	// Create - Creates a new SSH public key resource.
	Create(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyResource) (*SSHPublicKeyResourceResponse, error)
	// Delete - Delete an SSH public key.
	Delete(ctx context.Context, resourceGroupName string, sshPublicKeyName string) (*http.Response, error)
	// GenerateKeyPair - Generates and returns a public/private key pair and populates the SSH public key resource with the public key. The length of the key will be 3072 bits. This operation can only be performed once per SSH public key resource.
	GenerateKeyPair(ctx context.Context, resourceGroupName string, sshPublicKeyName string) (*SSHPublicKeyGenerateKeyPairResultResponse, error)
	// Get - Retrieves information about an SSH public key.
	Get(ctx context.Context, resourceGroupName string, sshPublicKeyName string) (*SSHPublicKeyResourceResponse, error)
	// ListByResourceGroup - Lists all of the SSH public keys in the specified resource group. Use the nextLink property in the response to get the next page of SSH public keys.
	ListByResourceGroup(resourceGroupName string) SSHPublicKeysGroupListResultPager
	// ListBySubscription - Lists all of the SSH public keys in the subscription. Use the nextLink property in the response to get the next page of SSH public keys.
	ListBySubscription() SSHPublicKeysGroupListResultPager
	// Update - Updates a new SSH public key resource.
	Update(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyUpdateResource) (*SSHPublicKeyResourceResponse, error)
}

// SSHPublicKeysClient implements the SSHPublicKeysOperations interface.
// Don't use this type directly, use NewSSHPublicKeysClient() instead.
type SSHPublicKeysClient struct {
	*Client
	subscriptionID string
}

// NewSSHPublicKeysClient creates a new instance of SSHPublicKeysClient with the specified values.
func NewSSHPublicKeysClient(c *Client, subscriptionID string) SSHPublicKeysOperations {
	return &SSHPublicKeysClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *SSHPublicKeysClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Create - Creates a new SSH public key resource.
func (client *SSHPublicKeysClient) Create(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyResource) (*SSHPublicKeyResourceResponse, error) {
	req, err := client.CreateCreateRequest(ctx, resourceGroupName, sshPublicKeyName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateHandleError(resp)
	}
	result, err := client.CreateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateCreateRequest creates the Create request.
func (client *SSHPublicKeysClient) CreateCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyResource) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateHandleResponse handles the Create response.
func (client *SSHPublicKeysClient) CreateHandleResponse(resp *azcore.Response) (*SSHPublicKeyResourceResponse, error) {
	result := SSHPublicKeyResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SSHPublicKeyResource)
}

// CreateHandleError handles the Create error response.
func (client *SSHPublicKeysClient) CreateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// Delete - Delete an SSH public key.
func (client *SSHPublicKeysClient) Delete(ctx context.Context, resourceGroupName string, sshPublicKeyName string) (*http.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, sshPublicKeyName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp.Response, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *SSHPublicKeysClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *SSHPublicKeysClient) DeleteHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// GenerateKeyPair - Generates and returns a public/private key pair and populates the SSH public key resource with the public key. The length of the key will be 3072 bits. This operation can only be performed once per SSH public key resource.
func (client *SSHPublicKeysClient) GenerateKeyPair(ctx context.Context, resourceGroupName string, sshPublicKeyName string) (*SSHPublicKeyGenerateKeyPairResultResponse, error) {
	req, err := client.GenerateKeyPairCreateRequest(ctx, resourceGroupName, sshPublicKeyName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GenerateKeyPairHandleError(resp)
	}
	result, err := client.GenerateKeyPairHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GenerateKeyPairCreateRequest creates the GenerateKeyPair request.
func (client *SSHPublicKeysClient) GenerateKeyPairCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}/generateKeyPair"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GenerateKeyPairHandleResponse handles the GenerateKeyPair response.
func (client *SSHPublicKeysClient) GenerateKeyPairHandleResponse(resp *azcore.Response) (*SSHPublicKeyGenerateKeyPairResultResponse, error) {
	result := SSHPublicKeyGenerateKeyPairResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SSHPublicKeyGenerateKeyPairResult)
}

// GenerateKeyPairHandleError handles the GenerateKeyPair error response.
func (client *SSHPublicKeysClient) GenerateKeyPairHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// Get - Retrieves information about an SSH public key.
func (client *SSHPublicKeysClient) Get(ctx context.Context, resourceGroupName string, sshPublicKeyName string) (*SSHPublicKeyResourceResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, sshPublicKeyName)
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
func (client *SSHPublicKeysClient) GetCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *SSHPublicKeysClient) GetHandleResponse(resp *azcore.Response) (*SSHPublicKeyResourceResponse, error) {
	result := SSHPublicKeyResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SSHPublicKeyResource)
}

// GetHandleError handles the Get error response.
func (client *SSHPublicKeysClient) GetHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// ListByResourceGroup - Lists all of the SSH public keys in the specified resource group. Use the nextLink property in the response to get the next page of SSH public keys.
func (client *SSHPublicKeysClient) ListByResourceGroup(resourceGroupName string) SSHPublicKeysGroupListResultPager {
	return &sshPublicKeysGroupListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, resourceGroupName)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *SSHPublicKeysGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SSHPublicKeysGroupListResult.NextLink)
		},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SSHPublicKeysClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SSHPublicKeysClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*SSHPublicKeysGroupListResultResponse, error) {
	result := SSHPublicKeysGroupListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SSHPublicKeysGroupListResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *SSHPublicKeysClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// ListBySubscription - Lists all of the SSH public keys in the subscription. Use the nextLink property in the response to get the next page of SSH public keys.
func (client *SSHPublicKeysClient) ListBySubscription() SSHPublicKeysGroupListResultPager {
	return &sshPublicKeysGroupListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListBySubscriptionCreateRequest(ctx)
		},
		responder: client.ListBySubscriptionHandleResponse,
		errorer:   client.ListBySubscriptionHandleError,
		advancer: func(ctx context.Context, resp *SSHPublicKeysGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SSHPublicKeysGroupListResult.NextLink)
		},
	}
}

// ListBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SSHPublicKeysClient) ListBySubscriptionCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/sshPublicKeys"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SSHPublicKeysClient) ListBySubscriptionHandleResponse(resp *azcore.Response) (*SSHPublicKeysGroupListResultResponse, error) {
	result := SSHPublicKeysGroupListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SSHPublicKeysGroupListResult)
}

// ListBySubscriptionHandleError handles the ListBySubscription error response.
func (client *SSHPublicKeysClient) ListBySubscriptionHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// Update - Updates a new SSH public key resource.
func (client *SSHPublicKeysClient) Update(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyUpdateResource) (*SSHPublicKeyResourceResponse, error) {
	req, err := client.UpdateCreateRequest(ctx, resourceGroupName, sshPublicKeyName, parameters)
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
func (client *SSHPublicKeysClient) UpdateCreateRequest(ctx context.Context, resourceGroupName string, sshPublicKeyName string, parameters SSHPublicKeyUpdateResource) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{sshPublicKeyName}", url.PathEscape(sshPublicKeyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// UpdateHandleResponse handles the Update response.
func (client *SSHPublicKeysClient) UpdateHandleResponse(resp *azcore.Response) (*SSHPublicKeyResourceResponse, error) {
	result := SSHPublicKeyResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SSHPublicKeyResource)
}

// UpdateHandleError handles the Update error response.
func (client *SSHPublicKeysClient) UpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}