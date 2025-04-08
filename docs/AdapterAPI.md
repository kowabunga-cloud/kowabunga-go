# \AdapterAPI

All URIs are relative to *https://your_kowabunga_kahuna_server/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAdapter**](AdapterAPI.md#DeleteAdapter) | **Delete** /adapter/{adapterId} | 
[**ListAdapters**](AdapterAPI.md#ListAdapters) | **Get** /adapter | 
[**ReadAdapter**](AdapterAPI.md#ReadAdapter) | **Get** /adapter/{adapterId} | 
[**UpdateAdapter**](AdapterAPI.md#UpdateAdapter) | **Put** /adapter/{adapterId} | 



## DeleteAdapter

> DeleteAdapter(ctx, adapterId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	kowabunga "github.com/kowabunga-cloud/kowabunga-go"
)

func main() {
	adapterId := "adapterId_example" // string | The ID of the network adapter.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.AdapterAPI.DeleteAdapter(context.Background(), adapterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdapterAPI.DeleteAdapter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adapterId** | **string** | The ID of the network adapter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdapterRequest struct via the builder pattern


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


## ListAdapters

> []string ListAdapters(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	kowabunga "github.com/kowabunga-cloud/kowabunga-go"
)

func main() {

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.AdapterAPI.ListAdapters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdapterAPI.ListAdapters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAdapters`: []string
	fmt.Fprintf(os.Stdout, "Response from `AdapterAPI.ListAdapters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAdaptersRequest struct via the builder pattern


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


## ReadAdapter

> Adapter ReadAdapter(ctx, adapterId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	kowabunga "github.com/kowabunga-cloud/kowabunga-go"
)

func main() {
	adapterId := "adapterId_example" // string | The ID of the network adapter.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.AdapterAPI.ReadAdapter(context.Background(), adapterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdapterAPI.ReadAdapter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadAdapter`: Adapter
	fmt.Fprintf(os.Stdout, "Response from `AdapterAPI.ReadAdapter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adapterId** | **string** | The ID of the network adapter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAdapterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Adapter**](Adapter.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdapter

> Adapter UpdateAdapter(ctx, adapterId).Adapter(adapter).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	kowabunga "github.com/kowabunga-cloud/kowabunga-go"
)

func main() {
	adapterId := "adapterId_example" // string | The ID of the network adapter.
	adapter := *kowabunga.NewAdapter("Name_example") // Adapter | Adapter payload.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.AdapterAPI.UpdateAdapter(context.Background(), adapterId).Adapter(adapter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdapterAPI.UpdateAdapter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAdapter`: Adapter
	fmt.Fprintf(os.Stdout, "Response from `AdapterAPI.UpdateAdapter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adapterId** | **string** | The ID of the network adapter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdapterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adapter** | [**Adapter**](Adapter.md) | Adapter payload. | 

### Return type

[**Adapter**](Adapter.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

