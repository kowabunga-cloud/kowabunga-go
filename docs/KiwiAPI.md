# \KiwiAPI

All URIs are relative to *https://raw.githubusercontent.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKiwi**](KiwiAPI.md#DeleteKiwi) | **Delete** /kiwi/{kiwiId} | 
[**ListKiwis**](KiwiAPI.md#ListKiwis) | **Get** /kiwi | 
[**ReadKiwi**](KiwiAPI.md#ReadKiwi) | **Get** /kiwi/{kiwiId} | 
[**UpdateKiwi**](KiwiAPI.md#UpdateKiwi) | **Put** /kiwi/{kiwiId} | 



## DeleteKiwi

> DeleteKiwi(ctx, kiwiId).Execute()





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
	kiwiId := "kiwiId_example" // string | The ID of the Kiwi (Kowabunga Inner Wan Interface) provides edge-network services..

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KiwiAPI.DeleteKiwi(context.Background(), kiwiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KiwiAPI.DeleteKiwi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kiwiId** | **string** | The ID of the Kiwi (Kowabunga Inner Wan Interface) provides edge-network services.. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKiwiRequest struct via the builder pattern


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


## ListKiwis

> []string ListKiwis(ctx).Execute()





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
	resp, r, err := apiClient.KiwiAPI.ListKiwis(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KiwiAPI.ListKiwis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKiwis`: []string
	fmt.Fprintf(os.Stdout, "Response from `KiwiAPI.ListKiwis`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKiwisRequest struct via the builder pattern


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


## ReadKiwi

> Kiwi ReadKiwi(ctx, kiwiId).Execute()





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
	kiwiId := "kiwiId_example" // string | The ID of the Kiwi (Kowabunga Inner Wan Interface) provides edge-network services..

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KiwiAPI.ReadKiwi(context.Background(), kiwiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KiwiAPI.ReadKiwi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadKiwi`: Kiwi
	fmt.Fprintf(os.Stdout, "Response from `KiwiAPI.ReadKiwi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kiwiId** | **string** | The ID of the Kiwi (Kowabunga Inner Wan Interface) provides edge-network services.. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadKiwiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Kiwi**](Kiwi.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKiwi

> Kiwi UpdateKiwi(ctx, kiwiId).Kiwi(kiwi).Execute()





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
	kiwiId := "kiwiId_example" // string | The ID of the Kiwi (Kowabunga Inner Wan Interface) provides edge-network services..
	kiwi := *openapiclient.NewKiwi("Name_example") // Kiwi | Kiwi payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KiwiAPI.UpdateKiwi(context.Background(), kiwiId).Kiwi(kiwi).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KiwiAPI.UpdateKiwi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKiwi`: Kiwi
	fmt.Fprintf(os.Stdout, "Response from `KiwiAPI.UpdateKiwi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kiwiId** | **string** | The ID of the Kiwi (Kowabunga Inner Wan Interface) provides edge-network services.. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKiwiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kiwi** | [**Kiwi**](Kiwi.md) | Kiwi payload. | 

### Return type

[**Kiwi**](Kiwi.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

