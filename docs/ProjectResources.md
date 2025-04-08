# ProjectResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vcpus** | Pointer to **int32** | The maximum total number of vCPUs allowed to be consumed by sum of all instances. | [optional] 
**Memory** | Pointer to **int64** | The maximum total memory (in bytes) allowed to be consumed by sum of all instances. | [optional] 
**Storage** | Pointer to **int64** | The maximum total disk capacity allowed to be consumed by sum of all instances. | [optional] 
**Instances** | Pointer to **int32** | The maximum number of instances allowed to be spawned. | [optional] 

## Methods

### NewProjectResources

`func NewProjectResources() *ProjectResources`

NewProjectResources instantiates a new ProjectResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectResourcesWithDefaults

`func NewProjectResourcesWithDefaults() *ProjectResources`

NewProjectResourcesWithDefaults instantiates a new ProjectResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcpus

`func (o *ProjectResources) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *ProjectResources) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *ProjectResources) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *ProjectResources) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetMemory

`func (o *ProjectResources) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ProjectResources) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ProjectResources) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ProjectResources) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetStorage

`func (o *ProjectResources) GetStorage() int64`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ProjectResources) GetStorageOk() (*int64, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ProjectResources) SetStorage(v int64)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ProjectResources) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetInstances

`func (o *ProjectResources) GetInstances() int32`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ProjectResources) GetInstancesOk() (*int32, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ProjectResources) SetInstances(v int32)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ProjectResources) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


