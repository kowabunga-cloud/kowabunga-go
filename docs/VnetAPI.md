# \VnetAPI

All URIs are relative to *https://raw.githubusercontent.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubnet**](VnetAPI.md#CreateSubnet) | **Post** /vnet/{vnetId}/subnet | 
[**DeleteVNet**](VnetAPI.md#DeleteVNet) | **Delete** /vnet/{vnetId} | 
[**ListVNetSubnets**](VnetAPI.md#ListVNetSubnets) | **Get** /vnet/{vnetId}/subnets | 
[**ListVNets**](VnetAPI.md#ListVNets) | **Get** /vnet | 
[**ReadVNet**](VnetAPI.md#ReadVNet) | **Get** /vnet/{vnetId} | 
[**SetVNetDefaultSubnet**](VnetAPI.md#SetVNetDefaultSubnet) | **Patch** /vnet/{vnetId}/subnet/{subnetId}/default | 
[**UpdateVNet**](VnetAPI.md#UpdateVNet) | **Put** /vnet/{vnetId} | 



## CreateSubnet

> Subnet CreateSubnet(ctx, vnetId).Subnet(subnet).Execute()





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
	vnetId := "vnetId_example" // string | The ID of the virtual network.
	subnet := *openapiclient.NewSubnet("Name_example", "Cidr_example", "Gateway_example") // Subnet | Subnet payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VnetAPI.CreateSubnet(context.Background(), vnetId).Subnet(subnet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VnetAPI.CreateSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubnet`: Subnet
	fmt.Fprintf(os.Stdout, "Response from `VnetAPI.CreateSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vnetId** | **string** | The ID of the virtual network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subnet** | [**Subnet**](Subnet.md) | Subnet payload. | 

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


## DeleteVNet

> DeleteVNet(ctx, vnetId).Execute()





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
	vnetId := "vnetId_example" // string | The ID of the virtual network.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VnetAPI.DeleteVNet(context.Background(), vnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VnetAPI.DeleteVNet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vnetId** | **string** | The ID of the virtual network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVNetRequest struct via the builder pattern


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


## ListVNetSubnets

> []string ListVNetSubnets(ctx, vnetId).Execute()





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
	vnetId := "vnetId_example" // string | The ID of the virtual network.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VnetAPI.ListVNetSubnets(context.Background(), vnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VnetAPI.ListVNetSubnets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVNetSubnets`: []string
	fmt.Fprintf(os.Stdout, "Response from `VnetAPI.ListVNetSubnets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vnetId** | **string** | The ID of the virtual network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVNetSubnetsRequest struct via the builder pattern


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


## ListVNets

> []string ListVNets(ctx).Execute()





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
	resp, r, err := apiClient.VnetAPI.ListVNets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VnetAPI.ListVNets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVNets`: []string
	fmt.Fprintf(os.Stdout, "Response from `VnetAPI.ListVNets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVNetsRequest struct via the builder pattern


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


## ReadVNet

> VNet ReadVNet(ctx, vnetId).Execute()





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
	vnetId := "vnetId_example" // string | The ID of the virtual network.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VnetAPI.ReadVNet(context.Background(), vnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VnetAPI.ReadVNet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadVNet`: VNet
	fmt.Fprintf(os.Stdout, "Response from `VnetAPI.ReadVNet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vnetId** | **string** | The ID of the virtual network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadVNetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VNet**](VNet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVNetDefaultSubnet

> SetVNetDefaultSubnet(ctx, vnetId, subnetId).Execute()





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
	vnetId := "vnetId_example" // string | The ID of the virtual network.
	subnetId := "subnetId_example" // string | The ID of the network subnet.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VnetAPI.SetVNetDefaultSubnet(context.Background(), vnetId, subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VnetAPI.SetVNetDefaultSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vnetId** | **string** | The ID of the virtual network. | 
**subnetId** | **string** | The ID of the network subnet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVNetDefaultSubnetRequest struct via the builder pattern


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


## UpdateVNet

> VNet UpdateVNet(ctx, vnetId).VNet(vNet).Execute()





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
	vnetId := "vnetId_example" // string | The ID of the virtual network.
	vNet := *openapiclient.NewVNet("Name_example", "Interface_example") // VNet | VNet payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VnetAPI.UpdateVNet(context.Background(), vnetId).VNet(vNet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VnetAPI.UpdateVNet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVNet`: VNet
	fmt.Fprintf(os.Stdout, "Response from `VnetAPI.UpdateVNet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vnetId** | **string** | The ID of the virtual network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVNetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vNet** | [**VNet**](VNet.md) | VNet payload. | 

### Return type

[**VNet**](VNet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

