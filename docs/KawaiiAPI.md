# \KawaiiAPI

All URIs are relative to *https://raw.githubusercontent.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKawaiiIpSec**](KawaiiAPI.md#CreateKawaiiIpSec) | **Post** /kawaii/{kawaiiId}/ipsec | 
[**DeleteKawaii**](KawaiiAPI.md#DeleteKawaii) | **Delete** /kawaii/{kawaiiId} | 
[**DeleteKawaiiIpSec**](KawaiiAPI.md#DeleteKawaiiIpSec) | **Delete** /kawaii/{kawaiiId}/ipsec/{KawaiiIpSecId} | 
[**ListKawaiiIpSecs**](KawaiiAPI.md#ListKawaiiIpSecs) | **Get** /kawaii/{kawaiiId}/ipsec | 
[**ListKawaiis**](KawaiiAPI.md#ListKawaiis) | **Get** /kawaii | 
[**ReadKawaii**](KawaiiAPI.md#ReadKawaii) | **Get** /kawaii/{kawaiiId} | 
[**ReadKawaiiIpSec**](KawaiiAPI.md#ReadKawaiiIpSec) | **Get** /kawaii/{kawaiiId}/ipsec/{KawaiiIpSecId} | 
[**UpdateKawaii**](KawaiiAPI.md#UpdateKawaii) | **Put** /kawaii/{kawaiiId} | 
[**UpdateKawaiiIpSec**](KawaiiAPI.md#UpdateKawaiiIpSec) | **Put** /kawaii/{kawaiiId}/ipsec/{KawaiiIpSecId} | 



## CreateKawaiiIpSec

> KawaiiIpSec CreateKawaiiIpSec(ctx, kawaiiId).KawaiiIpSec(kawaiiIpSec).Execute()





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
	kawaiiId := "kawaiiId_example" // string | The ID of the Kawaii.
	kawaiiIpSec := *openapiclient.NewKawaiiIpSec("Name_example", "RemoteIp_example", "RemoteSubnet_example", "PreSharedKey_example", int64(123), "Phase1IntegrityAlgorithm_example", "Phase1EncryptionAlgorithm_example", int64(123), "Phase2IntegrityAlgorithm_example", "Phase2EncryptionAlgorithm_example") // KawaiiIpSec | KawaiiIpSec payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KawaiiAPI.CreateKawaiiIpSec(context.Background(), kawaiiId).KawaiiIpSec(kawaiiIpSec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KawaiiAPI.CreateKawaiiIpSec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKawaiiIpSec`: KawaiiIpSec
	fmt.Fprintf(os.Stdout, "Response from `KawaiiAPI.CreateKawaiiIpSec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kawaiiId** | **string** | The ID of the Kawaii. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKawaiiIpSecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kawaiiIpSec** | [**KawaiiIpSec**](KawaiiIpSec.md) | KawaiiIpSec payload. | 

### Return type

[**KawaiiIpSec**](KawaiiIpSec.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKawaii

> DeleteKawaii(ctx, kawaiiId).Execute()





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
	kawaiiId := "kawaiiId_example" // string | The ID of the Kawaii.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KawaiiAPI.DeleteKawaii(context.Background(), kawaiiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KawaiiAPI.DeleteKawaii``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kawaiiId** | **string** | The ID of the Kawaii. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKawaiiRequest struct via the builder pattern


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


## DeleteKawaiiIpSec

> DeleteKawaiiIpSec(ctx, kawaiiId, kawaiiIpSecId).Execute()





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
	kawaiiId := "kawaiiId_example" // string | The ID of the Kawaii.
	kawaiiIpSecId := "kawaiiIpSecId_example" // string | The ID of the Kawaii IPsec connection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KawaiiAPI.DeleteKawaiiIpSec(context.Background(), kawaiiId, kawaiiIpSecId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KawaiiAPI.DeleteKawaiiIpSec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kawaiiId** | **string** | The ID of the Kawaii. | 
**kawaiiIpSecId** | **string** | The ID of the Kawaii IPsec connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKawaiiIpSecRequest struct via the builder pattern


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


## ListKawaiiIpSecs

> []string ListKawaiiIpSecs(ctx, kawaiiId).Execute()





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
	kawaiiId := "kawaiiId_example" // string | The ID of the Kawaii.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KawaiiAPI.ListKawaiiIpSecs(context.Background(), kawaiiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KawaiiAPI.ListKawaiiIpSecs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKawaiiIpSecs`: []string
	fmt.Fprintf(os.Stdout, "Response from `KawaiiAPI.ListKawaiiIpSecs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kawaiiId** | **string** | The ID of the Kawaii. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKawaiiIpSecsRequest struct via the builder pattern


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


## ListKawaiis

> []string ListKawaiis(ctx).Execute()





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
	resp, r, err := apiClient.KawaiiAPI.ListKawaiis(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KawaiiAPI.ListKawaiis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKawaiis`: []string
	fmt.Fprintf(os.Stdout, "Response from `KawaiiAPI.ListKawaiis`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKawaiisRequest struct via the builder pattern


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


## ReadKawaii

> Kawaii ReadKawaii(ctx, kawaiiId).Execute()





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
	kawaiiId := "kawaiiId_example" // string | The ID of the Kawaii.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KawaiiAPI.ReadKawaii(context.Background(), kawaiiId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KawaiiAPI.ReadKawaii``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadKawaii`: Kawaii
	fmt.Fprintf(os.Stdout, "Response from `KawaiiAPI.ReadKawaii`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kawaiiId** | **string** | The ID of the Kawaii. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadKawaiiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Kawaii**](Kawaii.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadKawaiiIpSec

> KawaiiIpSec ReadKawaiiIpSec(ctx, kawaiiId, kawaiiIpSecId).Execute()





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
	kawaiiId := "kawaiiId_example" // string | The ID of the Kawaii.
	kawaiiIpSecId := "kawaiiIpSecId_example" // string | The ID of the Kawaii IPsec connection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KawaiiAPI.ReadKawaiiIpSec(context.Background(), kawaiiId, kawaiiIpSecId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KawaiiAPI.ReadKawaiiIpSec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadKawaiiIpSec`: KawaiiIpSec
	fmt.Fprintf(os.Stdout, "Response from `KawaiiAPI.ReadKawaiiIpSec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kawaiiId** | **string** | The ID of the Kawaii. | 
**kawaiiIpSecId** | **string** | The ID of the Kawaii IPsec connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadKawaiiIpSecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KawaiiIpSec**](KawaiiIpSec.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKawaii

> Kawaii UpdateKawaii(ctx, kawaiiId).Kawaii(kawaii).Execute()





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
	kawaiiId := "kawaiiId_example" // string | The ID of the Kawaii.
	kawaii := *openapiclient.NewKawaii() // Kawaii | Kawaii payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KawaiiAPI.UpdateKawaii(context.Background(), kawaiiId).Kawaii(kawaii).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KawaiiAPI.UpdateKawaii``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKawaii`: Kawaii
	fmt.Fprintf(os.Stdout, "Response from `KawaiiAPI.UpdateKawaii`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kawaiiId** | **string** | The ID of the Kawaii. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKawaiiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kawaii** | [**Kawaii**](Kawaii.md) | Kawaii payload. | 

### Return type

[**Kawaii**](Kawaii.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKawaiiIpSec

> KawaiiIpSec UpdateKawaiiIpSec(ctx, kawaiiId, kawaiiIpSecId).KawaiiIpSec(kawaiiIpSec).Execute()





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
	kawaiiId := "kawaiiId_example" // string | The ID of the Kawaii.
	kawaiiIpSecId := "kawaiiIpSecId_example" // string | The ID of the Kawaii IPsec connection.
	kawaiiIpSec := *openapiclient.NewKawaiiIpSec("Name_example", "RemoteIp_example", "RemoteSubnet_example", "PreSharedKey_example", int64(123), "Phase1IntegrityAlgorithm_example", "Phase1EncryptionAlgorithm_example", int64(123), "Phase2IntegrityAlgorithm_example", "Phase2EncryptionAlgorithm_example") // KawaiiIpSec | KawaiiIpSec payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KawaiiAPI.UpdateKawaiiIpSec(context.Background(), kawaiiId, kawaiiIpSecId).KawaiiIpSec(kawaiiIpSec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KawaiiAPI.UpdateKawaiiIpSec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKawaiiIpSec`: KawaiiIpSec
	fmt.Fprintf(os.Stdout, "Response from `KawaiiAPI.UpdateKawaiiIpSec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kawaiiId** | **string** | The ID of the Kawaii. | 
**kawaiiIpSecId** | **string** | The ID of the Kawaii IPsec connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKawaiiIpSecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kawaiiIpSec** | [**KawaiiIpSec**](KawaiiIpSec.md) | KawaiiIpSec payload. | 

### Return type

[**KawaiiIpSec**](KawaiiIpSec.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

