# \TeamAPI

All URIs are relative to *https://raw.githubusercontent.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTeam**](TeamAPI.md#CreateTeam) | **Post** /team | 
[**DeleteTeam**](TeamAPI.md#DeleteTeam) | **Delete** /team/{teamId} | 
[**ListTeams**](TeamAPI.md#ListTeams) | **Get** /team | 
[**ReadTeam**](TeamAPI.md#ReadTeam) | **Get** /team/{teamId} | 
[**UpdateTeam**](TeamAPI.md#UpdateTeam) | **Put** /team/{teamId} | 



## CreateTeam

> Team CreateTeam(ctx).Team(team).Execute()





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
	team := *openapiclient.NewTeam("Name_example", []string{"Users_example"}) // Team | Team payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.CreateTeam(context.Background()).Team(team).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CreateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTeam`: Team
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.CreateTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **team** | [**Team**](Team.md) | Team payload. | 

### Return type

[**Team**](Team.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeam

> DeleteTeam(ctx, teamId).Execute()





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
	teamId := "teamId_example" // string | The ID of the Kowabunga users team.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamAPI.DeleteTeam(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.DeleteTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | The ID of the Kowabunga users team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamRequest struct via the builder pattern


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


## ListTeams

> []string ListTeams(ctx).Execute()





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
	resp, r, err := apiClient.TeamAPI.ListTeams(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.ListTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTeams`: []string
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.ListTeams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamsRequest struct via the builder pattern


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


## ReadTeam

> Team ReadTeam(ctx, teamId).Execute()





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
	teamId := "teamId_example" // string | The ID of the Kowabunga users team.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.ReadTeam(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.ReadTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadTeam`: Team
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.ReadTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | The ID of the Kowabunga users team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Team**](Team.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeam

> Team UpdateTeam(ctx, teamId).Team(team).Execute()





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
	teamId := "teamId_example" // string | The ID of the Kowabunga users team.
	team := *openapiclient.NewTeam("Name_example", []string{"Users_example"}) // Team | Team payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.UpdateTeam(context.Background(), teamId).Team(team).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.UpdateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeam`: Team
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.UpdateTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | The ID of the Kowabunga users team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **team** | [**Team**](Team.md) | Team payload. | 

### Return type

[**Team**](Team.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

