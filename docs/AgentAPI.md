# \AgentAPI

All URIs are relative to *https://raw.githubusercontent.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgent**](AgentAPI.md#CreateAgent) | **Post** /agent | 
[**DeleteAgent**](AgentAPI.md#DeleteAgent) | **Delete** /agent/{agentId} | 
[**ListAgents**](AgentAPI.md#ListAgents) | **Get** /agent | 
[**ReadAgent**](AgentAPI.md#ReadAgent) | **Get** /agent/{agentId} | 
[**SetAgentApiToken**](AgentAPI.md#SetAgentApiToken) | **Patch** /agent/{agentId}/token | 
[**UpdateAgent**](AgentAPI.md#UpdateAgent) | **Put** /agent/{agentId} | 



## CreateAgent

> Agent CreateAgent(ctx).Agent(agent).Execute()





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
	agent := *openapiclient.NewAgent("Name_example", "Type_example") // Agent | Agent payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.CreateAgent(context.Background()).Agent(agent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.CreateAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAgent`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.CreateAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | [**Agent**](Agent.md) | Agent payload. | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgent

> DeleteAgent(ctx, agentId).Execute()





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
	agentId := "agentId_example" // string | The ID of the Kowabunga remote agent.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentAPI.DeleteAgent(context.Background(), agentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.DeleteAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | The ID of the Kowabunga remote agent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentRequest struct via the builder pattern


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


## ListAgents

> []string ListAgents(ctx).Execute()





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
	resp, r, err := apiClient.AgentAPI.ListAgents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.ListAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgents`: []string
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.ListAgents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAgentsRequest struct via the builder pattern


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


## ReadAgent

> Agent ReadAgent(ctx, agentId).Execute()





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
	agentId := "agentId_example" // string | The ID of the Kowabunga remote agent.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.ReadAgent(context.Background(), agentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.ReadAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadAgent`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.ReadAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | The ID of the Kowabunga remote agent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAgentApiToken

> ApiToken SetAgentApiToken(ctx, agentId).Expire(expire).ExpirationDate(expirationDate).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	agentId := "agentId_example" // string | The ID of the Kowabunga remote agent.
	expire := true // bool | Whether or not the token should expire. (optional)
	expirationDate := time.Now() // string | Token's expiration date (YYYY-MM-DD format). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.SetAgentApiToken(context.Background(), agentId).Expire(expire).ExpirationDate(expirationDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.SetAgentApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAgentApiToken`: ApiToken
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.SetAgentApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | The ID of the Kowabunga remote agent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAgentApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expire** | **bool** | Whether or not the token should expire. | 
 **expirationDate** | **string** | Token&#39;s expiration date (YYYY-MM-DD format). | 

### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgent

> Agent UpdateAgent(ctx, agentId).Agent(agent).Execute()





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
	agentId := "agentId_example" // string | The ID of the Kowabunga remote agent.
	agent := *openapiclient.NewAgent("Name_example", "Type_example") // Agent | Agent payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.UpdateAgent(context.Background(), agentId).Agent(agent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.UpdateAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAgent`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.UpdateAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | The ID of the Kowabunga remote agent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agent** | [**Agent**](Agent.md) | Agent payload. | 

### Return type

[**Agent**](Agent.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

