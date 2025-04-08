# \KonveyAPI

All URIs are relative to *https://raw.githubusercontent.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKonvey**](KonveyAPI.md#DeleteKonvey) | **Delete** /konvey/{konveyId} | 
[**ListKonveys**](KonveyAPI.md#ListKonveys) | **Get** /konvey | 
[**ReadKonvey**](KonveyAPI.md#ReadKonvey) | **Get** /konvey/{konveyId} | 
[**UpdateKonvey**](KonveyAPI.md#UpdateKonvey) | **Put** /konvey/{konveyId} | 



## DeleteKonvey

> DeleteKonvey(ctx, konveyId).Execute()





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
	konveyId := "konveyId_example" // string | The ID of the Konvey (Kowabunga Network Load-Balancer).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KonveyAPI.DeleteKonvey(context.Background(), konveyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KonveyAPI.DeleteKonvey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**konveyId** | **string** | The ID of the Konvey (Kowabunga Network Load-Balancer). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKonveyRequest struct via the builder pattern


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


## ListKonveys

> []string ListKonveys(ctx).Execute()





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
	resp, r, err := apiClient.KonveyAPI.ListKonveys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KonveyAPI.ListKonveys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKonveys`: []string
	fmt.Fprintf(os.Stdout, "Response from `KonveyAPI.ListKonveys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKonveysRequest struct via the builder pattern


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


## ReadKonvey

> Konvey ReadKonvey(ctx, konveyId).Execute()





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
	konveyId := "konveyId_example" // string | The ID of the Konvey (Kowabunga Network Load-Balancer).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KonveyAPI.ReadKonvey(context.Background(), konveyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KonveyAPI.ReadKonvey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadKonvey`: Konvey
	fmt.Fprintf(os.Stdout, "Response from `KonveyAPI.ReadKonvey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**konveyId** | **string** | The ID of the Konvey (Kowabunga Network Load-Balancer). | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadKonveyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Konvey**](Konvey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKonvey

> Konvey UpdateKonvey(ctx, konveyId).Konvey(konvey).Execute()





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
	konveyId := "konveyId_example" // string | The ID of the Konvey (Kowabunga Network Load-Balancer).
	konvey := *openapiclient.NewKonvey([]openapiclient.KonveyEndpoint{*openapiclient.NewKonveyEndpoint("Name_example", int64(123), "Protocol_example", *openapiclient.NewKonveyBackends([]string{"Hosts_example"}, int64(123)))}) // Konvey | Konvey payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KonveyAPI.UpdateKonvey(context.Background(), konveyId).Konvey(konvey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KonveyAPI.UpdateKonvey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKonvey`: Konvey
	fmt.Fprintf(os.Stdout, "Response from `KonveyAPI.UpdateKonvey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**konveyId** | **string** | The ID of the Konvey (Kowabunga Network Load-Balancer). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKonveyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **konvey** | [**Konvey**](Konvey.md) | Konvey payload. | 

### Return type

[**Konvey**](Konvey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

