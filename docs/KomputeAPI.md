# \KomputeAPI

All URIs are relative to *https://raw.githubusercontent.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKompute**](KomputeAPI.md#DeleteKompute) | **Delete** /kompute/{komputeId} | 
[**ListKomputes**](KomputeAPI.md#ListKomputes) | **Get** /kompute | 
[**ReadKompute**](KomputeAPI.md#ReadKompute) | **Get** /kompute/{komputeId} | 
[**ReadKomputeState**](KomputeAPI.md#ReadKomputeState) | **Get** /kompute/{komputeId}/state | 
[**RebootKompute**](KomputeAPI.md#RebootKompute) | **Patch** /kompute/{komputeId}/reboot | 
[**ResetKompute**](KomputeAPI.md#ResetKompute) | **Patch** /kompute/{komputeId}/reset | 
[**ResumeKompute**](KomputeAPI.md#ResumeKompute) | **Patch** /kompute/{komputeId}/resume | 
[**ShutdownKompute**](KomputeAPI.md#ShutdownKompute) | **Patch** /kompute/{komputeId}/shutdown | 
[**StartKompute**](KomputeAPI.md#StartKompute) | **Patch** /kompute/{komputeId}/start | 
[**StopKompute**](KomputeAPI.md#StopKompute) | **Patch** /kompute/{komputeId}/stop | 
[**SuspendKompute**](KomputeAPI.md#SuspendKompute) | **Patch** /kompute/{komputeId}/suspend | 
[**UpdateKompute**](KomputeAPI.md#UpdateKompute) | **Put** /kompute/{komputeId} | 



## DeleteKompute

> DeleteKompute(ctx, komputeId).Execute()





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
	komputeId := "komputeId_example" // string | The ID of the Kompute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KomputeAPI.DeleteKompute(context.Background(), komputeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KomputeAPI.DeleteKompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**komputeId** | **string** | The ID of the Kompute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKomputeRequest struct via the builder pattern


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


## ListKomputes

> []string ListKomputes(ctx).Execute()





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
	resp, r, err := apiClient.KomputeAPI.ListKomputes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KomputeAPI.ListKomputes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKomputes`: []string
	fmt.Fprintf(os.Stdout, "Response from `KomputeAPI.ListKomputes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKomputesRequest struct via the builder pattern


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


## ReadKompute

> Kompute ReadKompute(ctx, komputeId).Execute()





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
	komputeId := "komputeId_example" // string | The ID of the Kompute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KomputeAPI.ReadKompute(context.Background(), komputeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KomputeAPI.ReadKompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadKompute`: Kompute
	fmt.Fprintf(os.Stdout, "Response from `KomputeAPI.ReadKompute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**komputeId** | **string** | The ID of the Kompute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadKomputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Kompute**](Kompute.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadKomputeState

> InstanceState ReadKomputeState(ctx, komputeId).Execute()





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
	komputeId := "komputeId_example" // string | The ID of the Kompute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KomputeAPI.ReadKomputeState(context.Background(), komputeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KomputeAPI.ReadKomputeState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadKomputeState`: InstanceState
	fmt.Fprintf(os.Stdout, "Response from `KomputeAPI.ReadKomputeState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**komputeId** | **string** | The ID of the Kompute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadKomputeStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InstanceState**](InstanceState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootKompute

> RebootKompute(ctx, komputeId).Execute()





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
	komputeId := "komputeId_example" // string | The ID of the Kompute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KomputeAPI.RebootKompute(context.Background(), komputeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KomputeAPI.RebootKompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**komputeId** | **string** | The ID of the Kompute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootKomputeRequest struct via the builder pattern


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


## ResetKompute

> ResetKompute(ctx, komputeId).Execute()





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
	komputeId := "komputeId_example" // string | The ID of the Kompute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KomputeAPI.ResetKompute(context.Background(), komputeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KomputeAPI.ResetKompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**komputeId** | **string** | The ID of the Kompute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetKomputeRequest struct via the builder pattern


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


## ResumeKompute

> ResumeKompute(ctx, komputeId).Execute()





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
	komputeId := "komputeId_example" // string | The ID of the Kompute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KomputeAPI.ResumeKompute(context.Background(), komputeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KomputeAPI.ResumeKompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**komputeId** | **string** | The ID of the Kompute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeKomputeRequest struct via the builder pattern


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


## ShutdownKompute

> ShutdownKompute(ctx, komputeId).Execute()





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
	komputeId := "komputeId_example" // string | The ID of the Kompute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KomputeAPI.ShutdownKompute(context.Background(), komputeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KomputeAPI.ShutdownKompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**komputeId** | **string** | The ID of the Kompute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownKomputeRequest struct via the builder pattern


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


## StartKompute

> StartKompute(ctx, komputeId).Execute()





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
	komputeId := "komputeId_example" // string | The ID of the Kompute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KomputeAPI.StartKompute(context.Background(), komputeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KomputeAPI.StartKompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**komputeId** | **string** | The ID of the Kompute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartKomputeRequest struct via the builder pattern


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


## StopKompute

> StopKompute(ctx, komputeId).Execute()





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
	komputeId := "komputeId_example" // string | The ID of the Kompute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KomputeAPI.StopKompute(context.Background(), komputeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KomputeAPI.StopKompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**komputeId** | **string** | The ID of the Kompute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopKomputeRequest struct via the builder pattern


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


## SuspendKompute

> SuspendKompute(ctx, komputeId).Execute()





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
	komputeId := "komputeId_example" // string | The ID of the Kompute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KomputeAPI.SuspendKompute(context.Background(), komputeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KomputeAPI.SuspendKompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**komputeId** | **string** | The ID of the Kompute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendKomputeRequest struct via the builder pattern


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


## UpdateKompute

> Kompute UpdateKompute(ctx, komputeId).Kompute(kompute).Execute()





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
	komputeId := "komputeId_example" // string | The ID of the Kompute.
	kompute := *openapiclient.NewKompute("Name_example", int64(123), int64(123), int64(123)) // Kompute | Kompute payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KomputeAPI.UpdateKompute(context.Background(), komputeId).Kompute(kompute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KomputeAPI.UpdateKompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKompute`: Kompute
	fmt.Fprintf(os.Stdout, "Response from `KomputeAPI.UpdateKompute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**komputeId** | **string** | The ID of the Kompute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKomputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kompute** | [**Kompute**](Kompute.md) | Kompute payload. | 

### Return type

[**Kompute**](Kompute.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

