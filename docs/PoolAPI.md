# \PoolAPI

All URIs are relative to *https://your_kowabunga_kahuna_server/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTemplate**](PoolAPI.md#CreateTemplate) | **Post** /pool/{poolId}/template | 
[**DeleteStoragePool**](PoolAPI.md#DeleteStoragePool) | **Delete** /pool/{poolId} | 
[**ListStoragePoolTemplates**](PoolAPI.md#ListStoragePoolTemplates) | **Get** /pool/{poolId}/templates | 
[**ListStoragePoolVolumes**](PoolAPI.md#ListStoragePoolVolumes) | **Get** /pool/{poolId}/volumes | 
[**ListStoragePools**](PoolAPI.md#ListStoragePools) | **Get** /pool | 
[**ReadStoragePool**](PoolAPI.md#ReadStoragePool) | **Get** /pool/{poolId} | 
[**SetStoragePoolDefaultTemplate**](PoolAPI.md#SetStoragePoolDefaultTemplate) | **Patch** /pool/{poolId}/template/{templateId}/default | 
[**UpdateStoragePool**](PoolAPI.md#UpdateStoragePool) | **Put** /pool/{poolId} | 



## CreateTemplate

> Template CreateTemplate(ctx, poolId).Template(template).Execute()





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
	poolId := "poolId_example" // string | The ID of the storage pool.
	template := *kowabunga.NewTemplate("Name_example", "Source_example") // Template | Template payload.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolAPI.CreateTemplate(context.Background(), poolId).Template(template).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.CreateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `PoolAPI.CreateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | The ID of the storage pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **template** | [**Template**](Template.md) | Template payload. | 

### Return type

[**Template**](Template.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStoragePool

> DeleteStoragePool(ctx, poolId).Execute()





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
	poolId := "poolId_example" // string | The ID of the storage pool.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.PoolAPI.DeleteStoragePool(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.DeleteStoragePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | The ID of the storage pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoragePoolRequest struct via the builder pattern


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


## ListStoragePoolTemplates

> []string ListStoragePoolTemplates(ctx, poolId).Execute()





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
	poolId := "poolId_example" // string | The ID of the storage pool.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolAPI.ListStoragePoolTemplates(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.ListStoragePoolTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStoragePoolTemplates`: []string
	fmt.Fprintf(os.Stdout, "Response from `PoolAPI.ListStoragePoolTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | The ID of the storage pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStoragePoolTemplatesRequest struct via the builder pattern


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


## ListStoragePoolVolumes

> []string ListStoragePoolVolumes(ctx, poolId).Execute()





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
	poolId := "poolId_example" // string | The ID of the storage pool.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolAPI.ListStoragePoolVolumes(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.ListStoragePoolVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStoragePoolVolumes`: []string
	fmt.Fprintf(os.Stdout, "Response from `PoolAPI.ListStoragePoolVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | The ID of the storage pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStoragePoolVolumesRequest struct via the builder pattern


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


## ListStoragePools

> []string ListStoragePools(ctx).Execute()





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
	resp, r, err := apiClient.PoolAPI.ListStoragePools(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.ListStoragePools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStoragePools`: []string
	fmt.Fprintf(os.Stdout, "Response from `PoolAPI.ListStoragePools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListStoragePoolsRequest struct via the builder pattern


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


## ReadStoragePool

> StoragePool ReadStoragePool(ctx, poolId).Execute()





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
	poolId := "poolId_example" // string | The ID of the storage pool.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolAPI.ReadStoragePool(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.ReadStoragePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadStoragePool`: StoragePool
	fmt.Fprintf(os.Stdout, "Response from `PoolAPI.ReadStoragePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | The ID of the storage pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadStoragePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StoragePool**](StoragePool.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetStoragePoolDefaultTemplate

> SetStoragePoolDefaultTemplate(ctx, poolId, templateId).Execute()





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
	poolId := "poolId_example" // string | The ID of the storage pool.
	templateId := "templateId_example" // string | The ID of the image template.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.PoolAPI.SetStoragePoolDefaultTemplate(context.Background(), poolId, templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.SetStoragePoolDefaultTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | The ID of the storage pool. | 
**templateId** | **string** | The ID of the image template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetStoragePoolDefaultTemplateRequest struct via the builder pattern


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


## UpdateStoragePool

> StoragePool UpdateStoragePool(ctx, poolId).StoragePool(storagePool).Execute()





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
	poolId := "poolId_example" // string | The ID of the storage pool.
	storagePool := *kowabunga.NewStoragePool("Name_example", "Pool_example", []string{"Agents_example"}) // StoragePool | StoragePool payload.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolAPI.UpdateStoragePool(context.Background(), poolId).StoragePool(storagePool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.UpdateStoragePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStoragePool`: StoragePool
	fmt.Fprintf(os.Stdout, "Response from `PoolAPI.UpdateStoragePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | The ID of the storage pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStoragePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storagePool** | [**StoragePool**](StoragePool.md) | StoragePool payload. | 

### Return type

[**StoragePool**](StoragePool.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

