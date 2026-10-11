# \KwarantineAPI

All URIs are relative to *https://your_kowabunga_kahuna_server/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKwarantineInstance**](KwarantineAPI.md#AddKwarantineInstance) | **Put** /kwarantine/{kwarantineId}/instance/{instanceId} | 
[**AddKwarantineKompute**](KwarantineAPI.md#AddKwarantineKompute) | **Put** /kwarantine/{kwarantineId}/kompute/{komputeId} | 
[**DeleteKwarantine**](KwarantineAPI.md#DeleteKwarantine) | **Delete** /kwarantine/{kwarantineId} | 
[**ListKwarantines**](KwarantineAPI.md#ListKwarantines) | **Get** /kwarantine | 
[**ReadKwarantine**](KwarantineAPI.md#ReadKwarantine) | **Get** /kwarantine/{kwarantineId} | 
[**RemoveKwarantineInstance**](KwarantineAPI.md#RemoveKwarantineInstance) | **Delete** /kwarantine/{kwarantineId}/instance/{instanceId} | 
[**RemoveKwarantineKompute**](KwarantineAPI.md#RemoveKwarantineKompute) | **Delete** /kwarantine/{kwarantineId}/kompute/{komputeId} | 
[**UpdateKwarantine**](KwarantineAPI.md#UpdateKwarantine) | **Put** /kwarantine/{kwarantineId} | 



## AddKwarantineInstance

> Kwarantine AddKwarantineInstance(ctx, kwarantineId, instanceId).Execute()





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
	kwarantineId := "kwarantineId_example" // string | The ID of the Kwarantine.
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.KwarantineAPI.AddKwarantineInstance(context.Background(), kwarantineId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KwarantineAPI.AddKwarantineInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddKwarantineInstance`: Kwarantine
	fmt.Fprintf(os.Stdout, "Response from `KwarantineAPI.AddKwarantineInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kwarantineId** | **string** | The ID of the Kwarantine. | 
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddKwarantineInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Kwarantine**](Kwarantine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddKwarantineKompute

> Kwarantine AddKwarantineKompute(ctx, kwarantineId, komputeId).Execute()





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
	kwarantineId := "kwarantineId_example" // string | The ID of the Kwarantine.
	komputeId := "komputeId_example" // string | The ID of the Kompute.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.KwarantineAPI.AddKwarantineKompute(context.Background(), kwarantineId, komputeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KwarantineAPI.AddKwarantineKompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddKwarantineKompute`: Kwarantine
	fmt.Fprintf(os.Stdout, "Response from `KwarantineAPI.AddKwarantineKompute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kwarantineId** | **string** | The ID of the Kwarantine. | 
**komputeId** | **string** | The ID of the Kompute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddKwarantineKomputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Kwarantine**](Kwarantine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKwarantine

> DeleteKwarantine(ctx, kwarantineId).Execute()





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
	kwarantineId := "kwarantineId_example" // string | The ID of the Kwarantine.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.KwarantineAPI.DeleteKwarantine(context.Background(), kwarantineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KwarantineAPI.DeleteKwarantine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kwarantineId** | **string** | The ID of the Kwarantine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKwarantineRequest struct via the builder pattern


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


## ListKwarantines

> []string ListKwarantines(ctx).Execute()





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
	resp, r, err := apiClient.KwarantineAPI.ListKwarantines(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KwarantineAPI.ListKwarantines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKwarantines`: []string
	fmt.Fprintf(os.Stdout, "Response from `KwarantineAPI.ListKwarantines`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKwarantinesRequest struct via the builder pattern


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


## ReadKwarantine

> Kwarantine ReadKwarantine(ctx, kwarantineId).Execute()





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
	kwarantineId := "kwarantineId_example" // string | The ID of the Kwarantine.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.KwarantineAPI.ReadKwarantine(context.Background(), kwarantineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KwarantineAPI.ReadKwarantine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadKwarantine`: Kwarantine
	fmt.Fprintf(os.Stdout, "Response from `KwarantineAPI.ReadKwarantine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kwarantineId** | **string** | The ID of the Kwarantine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadKwarantineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Kwarantine**](Kwarantine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveKwarantineInstance

> RemoveKwarantineInstance(ctx, kwarantineId, instanceId).Execute()





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
	kwarantineId := "kwarantineId_example" // string | The ID of the Kwarantine.
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.KwarantineAPI.RemoveKwarantineInstance(context.Background(), kwarantineId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KwarantineAPI.RemoveKwarantineInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kwarantineId** | **string** | The ID of the Kwarantine. | 
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveKwarantineInstanceRequest struct via the builder pattern


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


## RemoveKwarantineKompute

> RemoveKwarantineKompute(ctx, kwarantineId, komputeId).Execute()





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
	kwarantineId := "kwarantineId_example" // string | The ID of the Kwarantine.
	komputeId := "komputeId_example" // string | The ID of the Kompute.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.KwarantineAPI.RemoveKwarantineKompute(context.Background(), kwarantineId, komputeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KwarantineAPI.RemoveKwarantineKompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kwarantineId** | **string** | The ID of the Kwarantine. | 
**komputeId** | **string** | The ID of the Kompute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveKwarantineKomputeRequest struct via the builder pattern


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


## UpdateKwarantine

> Kwarantine UpdateKwarantine(ctx, kwarantineId).Kwarantine(kwarantine).Execute()





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
	kwarantineId := "kwarantineId_example" // string | The ID of the Kwarantine.
	kwarantine := *kowabunga.NewKwarantine("Name_example") // Kwarantine | Kwarantine payload.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.KwarantineAPI.UpdateKwarantine(context.Background(), kwarantineId).Kwarantine(kwarantine).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KwarantineAPI.UpdateKwarantine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKwarantine`: Kwarantine
	fmt.Fprintf(os.Stdout, "Response from `KwarantineAPI.UpdateKwarantine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kwarantineId** | **string** | The ID of the Kwarantine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKwarantineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kwarantine** | [**Kwarantine**](Kwarantine.md) | Kwarantine payload. | 

### Return type

[**Kwarantine**](Kwarantine.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

