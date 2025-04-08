# \NfsAPI

All URIs are relative to *https://raw.githubusercontent.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteStorageNFS**](NfsAPI.md#DeleteStorageNFS) | **Delete** /nfs/{nfsId} | 
[**ListStorageNFSKylos**](NfsAPI.md#ListStorageNFSKylos) | **Get** /nfs/{nfsId}/kylo | 
[**ListStorageNFSs**](NfsAPI.md#ListStorageNFSs) | **Get** /nfs | 
[**ReadStorageNFS**](NfsAPI.md#ReadStorageNFS) | **Get** /nfs/{nfsId} | 
[**UpdateStorageNFS**](NfsAPI.md#UpdateStorageNFS) | **Put** /nfs/{nfsId} | 



## DeleteStorageNFS

> DeleteStorageNFS(ctx, nfsId).Execute()





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
	nfsId := "nfsId_example" // string | The ID of the NFS storage.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NfsAPI.DeleteStorageNFS(context.Background(), nfsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsAPI.DeleteStorageNFS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfsId** | **string** | The ID of the NFS storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorageNFSRequest struct via the builder pattern


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


## ListStorageNFSKylos

> []string ListStorageNFSKylos(ctx, nfsId).Execute()





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
	nfsId := "nfsId_example" // string | The ID of the NFS storage.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsAPI.ListStorageNFSKylos(context.Background(), nfsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsAPI.ListStorageNFSKylos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStorageNFSKylos`: []string
	fmt.Fprintf(os.Stdout, "Response from `NfsAPI.ListStorageNFSKylos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfsId** | **string** | The ID of the NFS storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStorageNFSKylosRequest struct via the builder pattern


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


## ListStorageNFSs

> []string ListStorageNFSs(ctx).Execute()





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
	resp, r, err := apiClient.NfsAPI.ListStorageNFSs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsAPI.ListStorageNFSs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStorageNFSs`: []string
	fmt.Fprintf(os.Stdout, "Response from `NfsAPI.ListStorageNFSs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListStorageNFSsRequest struct via the builder pattern


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


## ReadStorageNFS

> StorageNFS ReadStorageNFS(ctx, nfsId).Execute()





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
	nfsId := "nfsId_example" // string | The ID of the NFS storage.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsAPI.ReadStorageNFS(context.Background(), nfsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsAPI.ReadStorageNFS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadStorageNFS`: StorageNFS
	fmt.Fprintf(os.Stdout, "Response from `NfsAPI.ReadStorageNFS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfsId** | **string** | The ID of the NFS storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadStorageNFSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StorageNFS**](StorageNFS.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStorageNFS

> StorageNFS UpdateStorageNFS(ctx, nfsId).StorageNFS(storageNFS).Execute()





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
	nfsId := "nfsId_example" // string | The ID of the NFS storage.
	storageNFS := *openapiclient.NewStorageNFS("Name_example", "Endpoint_example") // StorageNFS | StorageNFS payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NfsAPI.UpdateStorageNFS(context.Background(), nfsId).StorageNFS(storageNFS).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NfsAPI.UpdateStorageNFS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStorageNFS`: StorageNFS
	fmt.Fprintf(os.Stdout, "Response from `NfsAPI.UpdateStorageNFS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfsId** | **string** | The ID of the NFS storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorageNFSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageNFS** | [**StorageNFS**](StorageNFS.md) | StorageNFS payload. | 

### Return type

[**StorageNFS**](StorageNFS.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

