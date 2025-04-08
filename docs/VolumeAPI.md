# \VolumeAPI

All URIs are relative to *https://raw.githubusercontent.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVolume**](VolumeAPI.md#DeleteVolume) | **Delete** /volume/{volumeId} | 
[**ListVolumes**](VolumeAPI.md#ListVolumes) | **Get** /volume | 
[**ReadVolume**](VolumeAPI.md#ReadVolume) | **Get** /volume/{volumeId} | 
[**UpdateVolume**](VolumeAPI.md#UpdateVolume) | **Put** /volume/{volumeId} | 



## DeleteVolume

> DeleteVolume(ctx, volumeId).Execute()





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
	volumeId := "volumeId_example" // string | The ID of the storage volume.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VolumeAPI.DeleteVolume(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.DeleteVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | The ID of the storage volume. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeRequest struct via the builder pattern


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


## ListVolumes

> []string ListVolumes(ctx).Execute()





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
	resp, r, err := apiClient.VolumeAPI.ListVolumes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.ListVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumes`: []string
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.ListVolumes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVolumesRequest struct via the builder pattern


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


## ReadVolume

> Volume ReadVolume(ctx, volumeId).Execute()





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
	volumeId := "volumeId_example" // string | The ID of the storage volume.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.ReadVolume(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.ReadVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadVolume`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.ReadVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | The ID of the storage volume. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Volume**](Volume.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVolume

> Volume UpdateVolume(ctx, volumeId).Volume(volume).Execute()





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
	volumeId := "volumeId_example" // string | The ID of the storage volume.
	volume := *openapiclient.NewVolume("Name_example", "Type_example", int64(123)) // Volume | Volume payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.UpdateVolume(context.Background(), volumeId).Volume(volume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.UpdateVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVolume`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.UpdateVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | The ID of the storage volume. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volume** | [**Volume**](Volume.md) | Volume payload. | 

### Return type

[**Volume**](Volume.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

