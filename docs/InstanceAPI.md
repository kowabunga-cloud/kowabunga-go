# \InstanceAPI

All URIs are relative to *https://your_kowabunga_kahuna_server/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommitInstanceKMotion**](InstanceAPI.md#CommitInstanceKMotion) | **Patch** /instance/{instanceId}/kmotion/commit | 
[**DeleteInstance**](InstanceAPI.md#DeleteInstance) | **Delete** /instance/{instanceId} | 
[**ListInstances**](InstanceAPI.md#ListInstances) | **Get** /instance | 
[**PlanInstanceKMotion**](InstanceAPI.md#PlanInstanceKMotion) | **Post** /instance/{instanceId}/kmotion/plan | 
[**ReadInstance**](InstanceAPI.md#ReadInstance) | **Get** /instance/{instanceId} | 
[**ReadInstanceRemoteConnection**](InstanceAPI.md#ReadInstanceRemoteConnection) | **Get** /instance/{instanceId}/connect | 
[**ReadInstanceState**](InstanceAPI.md#ReadInstanceState) | **Get** /instance/{instanceId}/state | 
[**RebootInstance**](InstanceAPI.md#RebootInstance) | **Patch** /instance/{instanceId}/reboot | 
[**ResetInstance**](InstanceAPI.md#ResetInstance) | **Patch** /instance/{instanceId}/reset | 
[**ResumeInstance**](InstanceAPI.md#ResumeInstance) | **Patch** /instance/{instanceId}/resume | 
[**ShutdownInstance**](InstanceAPI.md#ShutdownInstance) | **Patch** /instance/{instanceId}/shutdown | 
[**StartInstance**](InstanceAPI.md#StartInstance) | **Patch** /instance/{instanceId}/start | 
[**StopInstance**](InstanceAPI.md#StopInstance) | **Patch** /instance/{instanceId}/stop | 
[**SuspendInstance**](InstanceAPI.md#SuspendInstance) | **Patch** /instance/{instanceId}/suspend | 
[**UpdateInstance**](InstanceAPI.md#UpdateInstance) | **Put** /instance/{instanceId} | 



## CommitInstanceKMotion

> CommitInstanceKMotion(ctx, instanceId).KMotionCommitRequest(kMotionCommitRequest).Execute()





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
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.
	kMotionCommitRequest := *kowabunga.NewKMotionCommitRequest("Plan_example") // KMotionCommitRequest | KMotionCommitRequest payload.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.InstanceAPI.CommitInstanceKMotion(context.Background(), instanceId).KMotionCommitRequest(kMotionCommitRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.CommitInstanceKMotion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommitInstanceKMotionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kMotionCommitRequest** | [**KMotionCommitRequest**](KMotionCommitRequest.md) | KMotionCommitRequest payload. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstance

> DeleteInstance(ctx, instanceId).Execute()





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
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.InstanceAPI.DeleteInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.DeleteInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


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


## ListInstances

> []string ListInstances(ctx).Execute()





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
	resp, r, err := apiClient.InstanceAPI.ListInstances(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ListInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstances`: []string
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.ListInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListInstancesRequest struct via the builder pattern


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


## PlanInstanceKMotion

> KMotionPlan PlanInstanceKMotion(ctx, instanceId).KMotionPlanRequest(kMotionPlanRequest).Execute()





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
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.
	kMotionPlanRequest := *kowabunga.NewKMotionPlanRequest() // KMotionPlanRequest | KMotionPlanRequest payload.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.PlanInstanceKMotion(context.Background(), instanceId).KMotionPlanRequest(kMotionPlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.PlanInstanceKMotion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanInstanceKMotion`: KMotionPlan
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.PlanInstanceKMotion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlanInstanceKMotionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kMotionPlanRequest** | [**KMotionPlanRequest**](KMotionPlanRequest.md) | KMotionPlanRequest payload. | 

### Return type

[**KMotionPlan**](KMotionPlan.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadInstance

> Instance ReadInstance(ctx, instanceId).Execute()





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
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.ReadInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ReadInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadInstance`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.ReadInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Instance**](Instance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadInstanceRemoteConnection

> InstanceRemoteAccess ReadInstanceRemoteConnection(ctx, instanceId).Execute()





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
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.ReadInstanceRemoteConnection(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ReadInstanceRemoteConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadInstanceRemoteConnection`: InstanceRemoteAccess
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.ReadInstanceRemoteConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadInstanceRemoteConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InstanceRemoteAccess**](InstanceRemoteAccess.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadInstanceState

> InstanceState ReadInstanceState(ctx, instanceId).Execute()





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
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.ReadInstanceState(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ReadInstanceState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadInstanceState`: InstanceState
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.ReadInstanceState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadInstanceStateRequest struct via the builder pattern


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


## RebootInstance

> RebootInstance(ctx, instanceId).Execute()





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
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.InstanceAPI.RebootInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.RebootInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootInstanceRequest struct via the builder pattern


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


## ResetInstance

> ResetInstance(ctx, instanceId).Execute()





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
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.InstanceAPI.ResetInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ResetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetInstanceRequest struct via the builder pattern


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


## ResumeInstance

> ResumeInstance(ctx, instanceId).Execute()





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
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.InstanceAPI.ResumeInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ResumeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeInstanceRequest struct via the builder pattern


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


## ShutdownInstance

> ShutdownInstance(ctx, instanceId).Execute()





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
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.InstanceAPI.ShutdownInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ShutdownInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownInstanceRequest struct via the builder pattern


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


## StartInstance

> StartInstance(ctx, instanceId).Execute()





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
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.InstanceAPI.StartInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.StartInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartInstanceRequest struct via the builder pattern


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


## StopInstance

> StopInstance(ctx, instanceId).Execute()





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
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.InstanceAPI.StopInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.StopInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopInstanceRequest struct via the builder pattern


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


## SuspendInstance

> SuspendInstance(ctx, instanceId).Execute()





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
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	r, err := apiClient.InstanceAPI.SuspendInstance(context.Background(), instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.SuspendInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendInstanceRequest struct via the builder pattern


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


## UpdateInstance

> Instance UpdateInstance(ctx, instanceId).InstanceUpdate(instanceUpdate).Execute()





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
	instanceId := "instanceId_example" // string | The ID of the virtual machine instance.
	instanceUpdate := *kowabunga.NewInstanceUpdate() // InstanceUpdate | InstanceUpdate payload.

	configuration := kowabunga.NewConfiguration()
	apiClient := kowabunga.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.UpdateInstance(context.Background(), instanceId).InstanceUpdate(instanceUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.UpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInstance`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.UpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | The ID of the virtual machine instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceUpdate** | [**InstanceUpdate**](InstanceUpdate.md) | InstanceUpdate payload. | 

### Return type

[**Instance**](Instance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

