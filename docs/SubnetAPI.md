# \SubnetAPI

All URIs are relative to *https://your_kowabunga_kahuna_server/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubnetAdapter**](SubnetAPI.md#CreateSubnetAdapter) | **Post** /subnet/{subnetId}/adapter | 
[**DeleteSubnet**](SubnetAPI.md#DeleteSubnet) | **Delete** /subnet/{subnetId} | 
[**ListSubnetAdapters**](SubnetAPI.md#ListSubnetAdapters) | **Get** /subnet/{subnetId}/adapters | 
[**ListSubnets**](SubnetAPI.md#ListSubnets) | **Get** /subnet | 
[**ReadSubnet**](SubnetAPI.md#ReadSubnet) | **Get** /subnet/{subnetId} | 
[**UpdateSubnet**](SubnetAPI.md#UpdateSubnet) | **Put** /subnet/{subnetId} | 



## CreateSubnetAdapter

> Adapter CreateSubnetAdapter(ctx, subnetId).Adapter(adapter).AssignIP(assignIP).Execute()





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
	subnetId := "subnetId_example" // string | The ID of the network subnet.
	adapter := *kowabunga.NewAdapter("Name_example") // Adapter | Adapter payload.
	assignIP := true // bool | Whether Kowabunga should pick and assign an IP address to this adapter. (optional) (default to false)

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.CreateSubnetAdapter(context.Background(), subnetId).Adapter(adapter).AssignIP(assignIP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.CreateSubnetAdapter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubnetAdapter`: Adapter
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.CreateSubnetAdapter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | The ID of the network subnet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubnetAdapterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adapter** | [**Adapter**](Adapter.md) | Adapter payload. | 
 **assignIP** | **bool** | Whether Kowabunga should pick and assign an IP address to this adapter. | [default to false]

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


## DeleteSubnet

> DeleteSubnet(ctx, subnetId).Execute()





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
	subnetId := "subnetId_example" // string | The ID of the network subnet.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.SubnetAPI.DeleteSubnet(context.Background(), subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.DeleteSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | The ID of the network subnet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubnetRequest struct via the builder pattern


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


## ListSubnetAdapters

> []string ListSubnetAdapters(ctx, subnetId).Execute()





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
	subnetId := "subnetId_example" // string | The ID of the network subnet.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.ListSubnetAdapters(context.Background(), subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.ListSubnetAdapters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubnetAdapters`: []string
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.ListSubnetAdapters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | The ID of the network subnet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubnetAdaptersRequest struct via the builder pattern


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


## ListSubnets

> []string ListSubnets(ctx).Execute()





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
	resp, r, err := apiClient.SubnetAPI.ListSubnets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.ListSubnets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubnets`: []string
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.ListSubnets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSubnetsRequest struct via the builder pattern


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


## ReadSubnet

> Subnet ReadSubnet(ctx, subnetId).Execute()





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
	subnetId := "subnetId_example" // string | The ID of the network subnet.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.ReadSubnet(context.Background(), subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.ReadSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadSubnet`: Subnet
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.ReadSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | The ID of the network subnet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subnet**](Subnet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubnet

> Subnet UpdateSubnet(ctx, subnetId).SubnetUpdate(subnetUpdate).Execute()





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
	subnetId := "subnetId_example" // string | The ID of the network subnet.
	subnetUpdate := *kowabunga.NewSubnetUpdate() // SubnetUpdate | SubnetUpdate payload.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.UpdateSubnet(context.Background(), subnetId).SubnetUpdate(subnetUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.UpdateSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubnet`: Subnet
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.UpdateSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | The ID of the network subnet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subnetUpdate** | [**SubnetUpdate**](SubnetUpdate.md) | SubnetUpdate payload. | 

### Return type

[**Subnet**](Subnet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

