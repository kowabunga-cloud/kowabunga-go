# \TokenAPI

All URIs are relative to *https://raw.githubusercontent.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiToken**](TokenAPI.md#DeleteApiToken) | **Delete** /token/{tokenId} | 
[**ListApiTokens**](TokenAPI.md#ListApiTokens) | **Get** /token | 
[**ReadApiToken**](TokenAPI.md#ReadApiToken) | **Get** /token/{tokenId} | 
[**UpdateApiToken**](TokenAPI.md#UpdateApiToken) | **Put** /token/{tokenId} | 



## DeleteApiToken

> DeleteApiToken(ctx, tokenId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tokenId := "tokenId_example" // string | The ID of the server-to-server authentication security token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TokenAPI.DeleteApiToken(context.Background(), tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.DeleteApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | The ID of the server-to-server authentication security token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiTokens

> []string ListApiTokens(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.ListApiTokens(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.ListApiTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApiTokens`: []string
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.ListApiTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiTokensRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadApiToken

> ApiToken ReadApiToken(ctx, tokenId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tokenId := "tokenId_example" // string | The ID of the server-to-server authentication security token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.ReadApiToken(context.Background(), tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.ReadApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadApiToken`: ApiToken
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.ReadApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | The ID of the server-to-server authentication security token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiToken

> ApiToken UpdateApiToken(ctx, tokenId).ApiToken(apiToken).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tokenId := "tokenId_example" // string | The ID of the server-to-server authentication security token.
	apiToken := *openapiclient.NewApiToken("Name_example", false) // ApiToken | ApiToken payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenAPI.UpdateApiToken(context.Background(), tokenId).ApiToken(apiToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.UpdateApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiToken`: ApiToken
	fmt.Fprintf(os.Stdout, "Response from `TokenAPI.UpdateApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | The ID of the server-to-server authentication security token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiToken** | [**ApiToken**](ApiToken.md) | ApiToken payload. | 

### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

