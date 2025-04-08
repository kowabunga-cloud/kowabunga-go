# \KaktusAPI

All URIs are relative to *https://your_kowabunga_kahuna_server/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKaktus**](KaktusAPI.md#DeleteKaktus) | **Delete** /kaktus/{kaktusId} | 
[**ListKaktusInstances**](KaktusAPI.md#ListKaktusInstances) | **Get** /kaktus/{kaktusId}/instances | 
[**ListKaktuss**](KaktusAPI.md#ListKaktuss) | **Get** /kaktus | 
[**ReadKaktus**](KaktusAPI.md#ReadKaktus) | **Get** /kaktus/{kaktusId} | 
[**ReadKaktusCaps**](KaktusAPI.md#ReadKaktusCaps) | **Get** /kaktus/{kaktusId}/caps | 
[**UpdateKaktus**](KaktusAPI.md#UpdateKaktus) | **Put** /kaktus/{kaktusId} | 



## DeleteKaktus

> DeleteKaktus(ctx, kaktusId).Execute()





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
	kaktusId := "kaktusId_example" // string | The ID of the Kaktus computing node.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.KaktusAPI.DeleteKaktus(context.Background(), kaktusId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaktusAPI.DeleteKaktus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kaktusId** | **string** | The ID of the Kaktus computing node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKaktusRequest struct via the builder pattern


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


## ListKaktusInstances

> []string ListKaktusInstances(ctx, kaktusId).Execute()





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
	kaktusId := "kaktusId_example" // string | The ID of the Kaktus computing node.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.KaktusAPI.ListKaktusInstances(context.Background(), kaktusId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaktusAPI.ListKaktusInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKaktusInstances`: []string
	fmt.Fprintf(os.Stdout, "Response from `KaktusAPI.ListKaktusInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kaktusId** | **string** | The ID of the Kaktus computing node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKaktusInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListKaktuss

> []string ListKaktuss(ctx).Execute()





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
	resp, r, err := apiClient.KaktusAPI.ListKaktuss(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaktusAPI.ListKaktuss``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKaktuss`: []string
	fmt.Fprintf(os.Stdout, "Response from `KaktusAPI.ListKaktuss`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKaktussRequest struct via the builder pattern


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


## ReadKaktus

> Kaktus ReadKaktus(ctx, kaktusId).Execute()





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
	kaktusId := "kaktusId_example" // string | The ID of the Kaktus computing node.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.KaktusAPI.ReadKaktus(context.Background(), kaktusId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaktusAPI.ReadKaktus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadKaktus`: Kaktus
	fmt.Fprintf(os.Stdout, "Response from `KaktusAPI.ReadKaktus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kaktusId** | **string** | The ID of the Kaktus computing node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadKaktusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Kaktus**](Kaktus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadKaktusCaps

> KaktusCaps ReadKaktusCaps(ctx, kaktusId).Execute()





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
	kaktusId := "kaktusId_example" // string | The ID of the Kaktus computing node.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.KaktusAPI.ReadKaktusCaps(context.Background(), kaktusId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaktusAPI.ReadKaktusCaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadKaktusCaps`: KaktusCaps
	fmt.Fprintf(os.Stdout, "Response from `KaktusAPI.ReadKaktusCaps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kaktusId** | **string** | The ID of the Kaktus computing node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadKaktusCapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KaktusCaps**](KaktusCaps.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKaktus

> Kaktus UpdateKaktus(ctx, kaktusId).Kaktus(kaktus).Execute()





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
	kaktusId := "kaktusId_example" // string | The ID of the Kaktus computing node.
	kaktus := *kowabunga.NewKaktus("Name_example", []string{"Agents_example"}) // Kaktus | Kaktus payload.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.KaktusAPI.UpdateKaktus(context.Background(), kaktusId).Kaktus(kaktus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KaktusAPI.UpdateKaktus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKaktus`: Kaktus
	fmt.Fprintf(os.Stdout, "Response from `KaktusAPI.UpdateKaktus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kaktusId** | **string** | The ID of the Kaktus computing node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKaktusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kaktus** | [**Kaktus**](Kaktus.md) | Kaktus payload. | 

### Return type

[**Kaktus**](Kaktus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

