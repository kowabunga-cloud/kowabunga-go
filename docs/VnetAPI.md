# \VnetAPI

All URIs are relative to *https://your_kowabunga_kahuna_server/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVNetSubnet**](VnetAPI.md#CreateVNetSubnet) | **Post** /vnet/{vnetId}/subnet | 
[**DeleteVNet**](VnetAPI.md#DeleteVNet) | **Delete** /vnet/{vnetId} | 
[**ListVNetSubnets**](VnetAPI.md#ListVNetSubnets) | **Get** /vnet/{vnetId}/subnets | 
[**ListVNets**](VnetAPI.md#ListVNets) | **Get** /vnet | 
[**ReadVNet**](VnetAPI.md#ReadVNet) | **Get** /vnet/{vnetId} | 
[**SetVNetDefaultSubnet**](VnetAPI.md#SetVNetDefaultSubnet) | **Patch** /vnet/{vnetId}/subnet/{subnetId}/default | 
[**UpdateVNet**](VnetAPI.md#UpdateVNet) | **Put** /vnet/{vnetId} | 



## CreateVNetSubnet

> Subnet CreateVNetSubnet(ctx, vnetId).Subnet(subnet).Execute()





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
	vnetId := "vnetId_example" // string | The ID of the virtual network.
	subnet := *kowabunga.NewSubnet("Name_example", "Cidr_example", "Gateway_example") // Subnet | Subnet payload.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.VnetAPI.CreateVNetSubnet(context.Background(), vnetId).Subnet(subnet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VnetAPI.CreateVNetSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVNetSubnet`: Subnet
	fmt.Fprintf(os.Stdout, "Response from `VnetAPI.CreateVNetSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vnetId** | **string** | The ID of the virtual network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVNetSubnetRequest struct via the builder pattern


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
	kowabunga "github.com/kowabunga-cloud/kowabunga-go"
)

func main() {
	vnetId := "vnetId_example" // string | The ID of the virtual network.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
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
	kowabunga "github.com/kowabunga-cloud/kowabunga-go"
)

func main() {
	vnetId := "vnetId_example" // string | The ID of the virtual network.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
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
	kowabunga "github.com/kowabunga-cloud/kowabunga-go"
)

func main() {

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
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
	kowabunga "github.com/kowabunga-cloud/kowabunga-go"
)

func main() {
	vnetId := "vnetId_example" // string | The ID of the virtual network.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
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
	kowabunga "github.com/kowabunga-cloud/kowabunga-go"
)

func main() {
	vnetId := "vnetId_example" // string | The ID of the virtual network.
	subnetId := "subnetId_example" // string | The ID of the network subnet.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
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

> VNet UpdateVNet(ctx, vnetId).VNetUpdate(vNetUpdate).Execute()





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
	vnetId := "vnetId_example" // string | The ID of the virtual network.
	vNetUpdate := *kowabunga.NewVNetUpdate() // VNetUpdate | VNetUpdate payload.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.VnetAPI.UpdateVNet(context.Background(), vnetId).VNetUpdate(vNetUpdate).Execute()
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

 **vNetUpdate** | [**VNetUpdate**](VNetUpdate.md) | VNetUpdate payload. | 

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

