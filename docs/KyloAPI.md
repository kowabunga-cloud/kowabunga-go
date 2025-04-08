# \KyloAPI

All URIs are relative to *https://raw.githubusercontent.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKylo**](KyloAPI.md#DeleteKylo) | **Delete** /kylo/{kyloId} | 
[**ListKylos**](KyloAPI.md#ListKylos) | **Get** /kylo | 
[**ReadKylo**](KyloAPI.md#ReadKylo) | **Get** /kylo/{kyloId} | 
[**UpdateKylo**](KyloAPI.md#UpdateKylo) | **Put** /kylo/{kyloId} | 



## DeleteKylo

> DeleteKylo(ctx, kyloId).Execute()





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
	kyloId := "kyloId_example" // string | The ID of the Kylo.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KyloAPI.DeleteKylo(context.Background(), kyloId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KyloAPI.DeleteKylo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kyloId** | **string** | The ID of the Kylo. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKyloRequest struct via the builder pattern


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


## ListKylos

> []string ListKylos(ctx).Execute()





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
	resp, r, err := apiClient.KyloAPI.ListKylos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KyloAPI.ListKylos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKylos`: []string
	fmt.Fprintf(os.Stdout, "Response from `KyloAPI.ListKylos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKylosRequest struct via the builder pattern


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


## ReadKylo

> Kylo ReadKylo(ctx, kyloId).Execute()





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
	kyloId := "kyloId_example" // string | The ID of the Kylo.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KyloAPI.ReadKylo(context.Background(), kyloId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KyloAPI.ReadKylo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadKylo`: Kylo
	fmt.Fprintf(os.Stdout, "Response from `KyloAPI.ReadKylo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kyloId** | **string** | The ID of the Kylo. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadKyloRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Kylo**](Kylo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKylo

> Kylo UpdateKylo(ctx, kyloId).Kylo(kylo).Execute()





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
	kyloId := "kyloId_example" // string | The ID of the Kylo.
	kylo := *openapiclient.NewKylo("Name_example") // Kylo | Kylo payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KyloAPI.UpdateKylo(context.Background(), kyloId).Kylo(kylo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KyloAPI.UpdateKylo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKylo`: Kylo
	fmt.Fprintf(os.Stdout, "Response from `KyloAPI.UpdateKylo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kyloId** | **string** | The ID of the Kylo. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKyloRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kylo** | [**Kylo**](Kylo.md) | Kylo payload. | 

### Return type

[**Kylo**](Kylo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

