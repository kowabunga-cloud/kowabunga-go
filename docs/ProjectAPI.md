# \ProjectAPI

All URIs are relative to *https://raw.githubusercontent.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProject**](ProjectAPI.md#CreateProject) | **Post** /project | 
[**CreateProjectDnsRecord**](ProjectAPI.md#CreateProjectDnsRecord) | **Post** /project/{projectId}/record | 
[**CreateProjectRegionKawaii**](ProjectAPI.md#CreateProjectRegionKawaii) | **Post** /project/{projectId}/region/{regionId}/kawaii | 
[**CreateProjectRegionKonvey**](ProjectAPI.md#CreateProjectRegionKonvey) | **Post** /project/{projectId}/region/{regionId}/konvey | 
[**CreateProjectRegionKylo**](ProjectAPI.md#CreateProjectRegionKylo) | **Post** /project/{projectId}/region/{regionId}/kylo | 
[**CreateProjectRegionVolume**](ProjectAPI.md#CreateProjectRegionVolume) | **Post** /project/{projectId}/region/{regionId}/volume | 
[**CreateProjectZoneInstance**](ProjectAPI.md#CreateProjectZoneInstance) | **Post** /project/{projectId}/zone/{zoneId}/instance | 
[**CreateProjectZoneKompute**](ProjectAPI.md#CreateProjectZoneKompute) | **Post** /project/{projectId}/zone/{zoneId}/kompute | 
[**CreateProjectZoneKonvey**](ProjectAPI.md#CreateProjectZoneKonvey) | **Post** /project/{projectId}/zone/{zoneId}/konvey | 
[**DeleteProject**](ProjectAPI.md#DeleteProject) | **Delete** /project/{projectId} | 
[**ListProjectDnsRecords**](ProjectAPI.md#ListProjectDnsRecords) | **Get** /project/{projectId}/records | 
[**ListProjectRegionKawaiis**](ProjectAPI.md#ListProjectRegionKawaiis) | **Get** /project/{projectId}/region/{regionId}/kawaiis | 
[**ListProjectRegionKonveys**](ProjectAPI.md#ListProjectRegionKonveys) | **Get** /project/{projectId}/region/{regionId}/konveys | 
[**ListProjectRegionKylos**](ProjectAPI.md#ListProjectRegionKylos) | **Get** /project/{projectId}/region/{regionId}/kylo | 
[**ListProjectRegionVolumes**](ProjectAPI.md#ListProjectRegionVolumes) | **Get** /project/{projectId}/region/{regionId}/volumes | 
[**ListProjectZoneInstances**](ProjectAPI.md#ListProjectZoneInstances) | **Get** /project/{projectId}/zone/{zoneId}/instances | 
[**ListProjectZoneKomputes**](ProjectAPI.md#ListProjectZoneKomputes) | **Get** /project/{projectId}/zone/{zoneId}/komputes | 
[**ListProjectZoneKonveys**](ProjectAPI.md#ListProjectZoneKonveys) | **Get** /project/{projectId}/zone/{zoneId}/konveys | 
[**ListProjects**](ProjectAPI.md#ListProjects) | **Get** /project | 
[**ReadProject**](ProjectAPI.md#ReadProject) | **Get** /project/{projectId} | 
[**ReadProjectCost**](ProjectAPI.md#ReadProjectCost) | **Get** /project/{projectId}/cost | 
[**ReadProjectUsage**](ProjectAPI.md#ReadProjectUsage) | **Get** /project/{projectId}/usage | 
[**UpdateProject**](ProjectAPI.md#UpdateProject) | **Put** /project/{projectId} | 



## CreateProject

> Project CreateProject(ctx).Project(project).SubnetSize(subnetSize).Execute()





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
	project := *openapiclient.NewProject("Name_example", []string{"Teams_example"}, []string{"Regions_example"}) // Project | Project payload.
	subnetSize := int32(56) // int32 | The minimum VPC subnet size to be affected to the project. WARNING, this cannot be changed later. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.CreateProject(context.Background()).Project(project).SubnetSize(subnetSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.CreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProject`: Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**Project**](Project.md) | Project payload. | 
 **subnetSize** | **int32** | The minimum VPC subnet size to be affected to the project. WARNING, this cannot be changed later. | 

### Return type

[**Project**](Project.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectDnsRecord

> DnsRecord CreateProjectDnsRecord(ctx, projectId).DnsRecord(dnsRecord).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	dnsRecord := *openapiclient.NewDnsRecord("Name_example", []string{"Addresses_example"}) // DnsRecord | DnsRecord payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.CreateProjectDnsRecord(context.Background(), projectId).DnsRecord(dnsRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.CreateProjectDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectDnsRecord`: DnsRecord
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.CreateProjectDnsRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsRecord** | [**DnsRecord**](DnsRecord.md) | DnsRecord payload. | 

### Return type

[**DnsRecord**](DnsRecord.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectRegionKawaii

> Kawaii CreateProjectRegionKawaii(ctx, projectId, regionId).Kawaii(kawaii).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	regionId := "regionId_example" // string | The ID of the region.
	kawaii := *openapiclient.NewKawaii() // Kawaii | Kawaii payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.CreateProjectRegionKawaii(context.Background(), projectId, regionId).Kawaii(kawaii).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.CreateProjectRegionKawaii``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectRegionKawaii`: Kawaii
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.CreateProjectRegionKawaii`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 
**regionId** | **string** | The ID of the region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRegionKawaiiRequest struct via the builder pattern


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


## CreateProjectRegionKonvey

> Konvey CreateProjectRegionKonvey(ctx, projectId, regionId).Konvey(konvey).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	regionId := "regionId_example" // string | The ID of the region.
	konvey := *openapiclient.NewKonvey([]openapiclient.KonveyEndpoint{*openapiclient.NewKonveyEndpoint("Name_example", int64(123), "Protocol_example", *openapiclient.NewKonveyBackends([]string{"Hosts_example"}, int64(123)))}) // Konvey | Konvey payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.CreateProjectRegionKonvey(context.Background(), projectId, regionId).Konvey(konvey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.CreateProjectRegionKonvey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectRegionKonvey`: Konvey
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.CreateProjectRegionKonvey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 
**regionId** | **string** | The ID of the region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRegionKonveyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **konvey** | [**Konvey**](Konvey.md) | Konvey payload. | 

### Return type

[**Konvey**](Konvey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectRegionKylo

> Kylo CreateProjectRegionKylo(ctx, projectId, regionId).Kylo(kylo).NfsId(nfsId).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	regionId := "regionId_example" // string | The ID of the region.
	kylo := *openapiclient.NewKylo("Name_example") // Kylo | Kylo payload.
	nfsId := "nfsId_example" // string | NFS storage ID (optional, region's default if unspecified). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.CreateProjectRegionKylo(context.Background(), projectId, regionId).Kylo(kylo).NfsId(nfsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.CreateProjectRegionKylo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectRegionKylo`: Kylo
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.CreateProjectRegionKylo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 
**regionId** | **string** | The ID of the region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRegionKyloRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kylo** | [**Kylo**](Kylo.md) | Kylo payload. | 
 **nfsId** | **string** | NFS storage ID (optional, region&#39;s default if unspecified). | 

### Return type

[**Kylo**](Kylo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectRegionVolume

> Volume CreateProjectRegionVolume(ctx, projectId, regionId).Volume(volume).PoolId(poolId).TemplateId(templateId).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	regionId := "regionId_example" // string | The ID of the region.
	volume := *openapiclient.NewVolume("Name_example", "Type_example", int64(123)) // Volume | Volume payload.
	poolId := "poolId_example" // string | Storage pool ID (optional, region's default if unspecified). (optional)
	templateId := "templateId_example" // string | Template to clone the storage volume from (optional, region's default if unspecified). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.CreateProjectRegionVolume(context.Background(), projectId, regionId).Volume(volume).PoolId(poolId).TemplateId(templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.CreateProjectRegionVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectRegionVolume`: Volume
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.CreateProjectRegionVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 
**regionId** | **string** | The ID of the region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRegionVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **volume** | [**Volume**](Volume.md) | Volume payload. | 
 **poolId** | **string** | Storage pool ID (optional, region&#39;s default if unspecified). | 
 **templateId** | **string** | Template to clone the storage volume from (optional, region&#39;s default if unspecified). | 

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


## CreateProjectZoneInstance

> Instance CreateProjectZoneInstance(ctx, projectId, zoneId).Instance(instance).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	zoneId := "zoneId_example" // string | The ID of the availability zone.
	instance := *openapiclient.NewInstance("Name_example", int64(123), int64(123)) // Instance | Instance payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.CreateProjectZoneInstance(context.Background(), projectId, zoneId).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.CreateProjectZoneInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectZoneInstance`: Instance
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.CreateProjectZoneInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 
**zoneId** | **string** | The ID of the availability zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectZoneInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instance** | [**Instance**](Instance.md) | Instance payload. | 

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


## CreateProjectZoneKompute

> Kompute CreateProjectZoneKompute(ctx, projectId, zoneId).Kompute(kompute).PoolId(poolId).TemplateId(templateId).Public(public).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	zoneId := "zoneId_example" // string | The ID of the availability zone.
	kompute := *openapiclient.NewKompute("Name_example", int64(123), int64(123), int64(123)) // Kompute | Kompute payload.
	poolId := "poolId_example" // string | Storage pool ID (optional, region's default if unspecified). (optional)
	templateId := "templateId_example" // string | Template to clone the storage volume from (optional, region's default if unspecified). (optional)
	public := true // bool | Should Kompute be exposed over public Internet ? (a public IPv4 address will then be auto-assigned, default to false). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.CreateProjectZoneKompute(context.Background(), projectId, zoneId).Kompute(kompute).PoolId(poolId).TemplateId(templateId).Public(public).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.CreateProjectZoneKompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectZoneKompute`: Kompute
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.CreateProjectZoneKompute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 
**zoneId** | **string** | The ID of the availability zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectZoneKomputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kompute** | [**Kompute**](Kompute.md) | Kompute payload. | 
 **poolId** | **string** | Storage pool ID (optional, region&#39;s default if unspecified). | 
 **templateId** | **string** | Template to clone the storage volume from (optional, region&#39;s default if unspecified). | 
 **public** | **bool** | Should Kompute be exposed over public Internet ? (a public IPv4 address will then be auto-assigned, default to false). | 

### Return type

[**Kompute**](Kompute.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectZoneKonvey

> Konvey CreateProjectZoneKonvey(ctx, projectId, zoneId).Konvey(konvey).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	zoneId := "zoneId_example" // string | The ID of the availability zone.
	konvey := *openapiclient.NewKonvey([]openapiclient.KonveyEndpoint{*openapiclient.NewKonveyEndpoint("Name_example", int64(123), "Protocol_example", *openapiclient.NewKonveyBackends([]string{"Hosts_example"}, int64(123)))}) // Konvey | Konvey payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.CreateProjectZoneKonvey(context.Background(), projectId, zoneId).Konvey(konvey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.CreateProjectZoneKonvey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectZoneKonvey`: Konvey
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.CreateProjectZoneKonvey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 
**zoneId** | **string** | The ID of the availability zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectZoneKonveyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **konvey** | [**Konvey**](Konvey.md) | Konvey payload. | 

### Return type

[**Konvey**](Konvey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProject

> DeleteProject(ctx, projectId).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.DeleteProject(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


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


## ListProjectDnsRecords

> []string ListProjectDnsRecords(ctx, projectId).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ListProjectDnsRecords(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ListProjectDnsRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjectDnsRecords`: []string
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ListProjectDnsRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectDnsRecordsRequest struct via the builder pattern


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


## ListProjectRegionKawaiis

> []string ListProjectRegionKawaiis(ctx, projectId, regionId).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	regionId := "regionId_example" // string | The ID of the region.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ListProjectRegionKawaiis(context.Background(), projectId, regionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ListProjectRegionKawaiis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjectRegionKawaiis`: []string
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ListProjectRegionKawaiis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 
**regionId** | **string** | The ID of the region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectRegionKawaiisRequest struct via the builder pattern


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


## ListProjectRegionKonveys

> []string ListProjectRegionKonveys(ctx, projectId, regionId).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	regionId := "regionId_example" // string | The ID of the region.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ListProjectRegionKonveys(context.Background(), projectId, regionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ListProjectRegionKonveys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjectRegionKonveys`: []string
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ListProjectRegionKonveys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 
**regionId** | **string** | The ID of the region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectRegionKonveysRequest struct via the builder pattern


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


## ListProjectRegionKylos

> []string ListProjectRegionKylos(ctx, projectId, regionId).NfsId(nfsId).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	regionId := "regionId_example" // string | The ID of the region.
	nfsId := "nfsId_example" // string | NFS storage ID (optional, region's default if unspecified). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ListProjectRegionKylos(context.Background(), projectId, regionId).NfsId(nfsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ListProjectRegionKylos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjectRegionKylos`: []string
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ListProjectRegionKylos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 
**regionId** | **string** | The ID of the region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectRegionKylosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nfsId** | **string** | NFS storage ID (optional, region&#39;s default if unspecified). | 

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


## ListProjectRegionVolumes

> []string ListProjectRegionVolumes(ctx, projectId, regionId).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	regionId := "regionId_example" // string | The ID of the region.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ListProjectRegionVolumes(context.Background(), projectId, regionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ListProjectRegionVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjectRegionVolumes`: []string
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ListProjectRegionVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 
**regionId** | **string** | The ID of the region. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectRegionVolumesRequest struct via the builder pattern


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


## ListProjectZoneInstances

> []string ListProjectZoneInstances(ctx, projectId, zoneId).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	zoneId := "zoneId_example" // string | The ID of the availability zone.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ListProjectZoneInstances(context.Background(), projectId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ListProjectZoneInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjectZoneInstances`: []string
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ListProjectZoneInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 
**zoneId** | **string** | The ID of the availability zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectZoneInstancesRequest struct via the builder pattern


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


## ListProjectZoneKomputes

> []string ListProjectZoneKomputes(ctx, projectId, zoneId).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	zoneId := "zoneId_example" // string | The ID of the availability zone.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ListProjectZoneKomputes(context.Background(), projectId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ListProjectZoneKomputes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjectZoneKomputes`: []string
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ListProjectZoneKomputes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 
**zoneId** | **string** | The ID of the availability zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectZoneKomputesRequest struct via the builder pattern


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


## ListProjectZoneKonveys

> []string ListProjectZoneKonveys(ctx, projectId, zoneId).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	zoneId := "zoneId_example" // string | The ID of the availability zone.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ListProjectZoneKonveys(context.Background(), projectId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ListProjectZoneKonveys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjectZoneKonveys`: []string
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ListProjectZoneKonveys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 
**zoneId** | **string** | The ID of the availability zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectZoneKonveysRequest struct via the builder pattern


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


## ListProjects

> []string ListProjects(ctx).SubnetSize(subnetSize).Execute()





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
	subnetSize := int32(56) // int32 | The minimum VPC subnet size to be affected to the project. WARNING, this cannot be changed later. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ListProjects(context.Background()).SubnetSize(subnetSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ListProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProjects`: []string
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ListProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subnetSize** | **int32** | The minimum VPC subnet size to be affected to the project. WARNING, this cannot be changed later. | 

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


## ReadProject

> Project ReadProject(ctx, projectId).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ReadProject(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ReadProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadProject`: Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ReadProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadProjectCost

> Cost ReadProjectCost(ctx, projectId).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ReadProjectCost(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ReadProjectCost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadProjectCost`: Cost
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ReadProjectCost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadProjectCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Cost**](Cost.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadProjectUsage

> ProjectResources ReadProjectUsage(ctx, projectId).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ReadProjectUsage(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ReadProjectUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadProjectUsage`: ProjectResources
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ReadProjectUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadProjectUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectResources**](ProjectResources.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProject

> Project UpdateProject(ctx, projectId).Project(project).Execute()





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
	projectId := "projectId_example" // string | The ID of the project.
	project := *openapiclient.NewProject("Name_example", []string{"Teams_example"}, []string{"Regions_example"}) // Project | Project payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.UpdateProject(context.Background(), projectId).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.UpdateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProject`: Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**Project**](Project.md) | Project payload. | 

### Return type

[**Project**](Project.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

