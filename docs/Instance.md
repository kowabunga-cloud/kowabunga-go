# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The virtual machine instance ID (auto-generated). | [optional] 
**Name** | **string** | The virtual machine instance name. | 
**Description** | Pointer to **string** | The virtual machine instance description. | [optional] 
**Memory** | **int64** | The virtual machine instance memory size (in bytes). | 
**Vcpus** | **int64** | The virtual machine instance number of vCPUs. | 
**Adapters** | Pointer to **[]string** | a list of existing network adapters to be connected to the instance. | [optional] 
**Volumes** | Pointer to **[]string** | volumes list of existing storage volumes (i.e. disks) to be connected to the instance. | [optional] 

## Methods

### NewInstance

`func NewInstance(name string, memory int64, vcpus int64, ) *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Instance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Instance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Instance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Instance) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Instance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Instance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Instance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Instance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMemory

`func (o *Instance) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Instance) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Instance) SetMemory(v int64)`

SetMemory sets Memory field to given value.


### GetVcpus

`func (o *Instance) GetVcpus() int64`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *Instance) GetVcpusOk() (*int64, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *Instance) SetVcpus(v int64)`

SetVcpus sets Vcpus field to given value.


### GetAdapters

`func (o *Instance) GetAdapters() []string`

GetAdapters returns the Adapters field if non-nil, zero value otherwise.

### GetAdaptersOk

`func (o *Instance) GetAdaptersOk() (*[]string, bool)`

GetAdaptersOk returns a tuple with the Adapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapters

`func (o *Instance) SetAdapters(v []string)`

SetAdapters sets Adapters field to given value.

### HasAdapters

`func (o *Instance) HasAdapters() bool`

HasAdapters returns a boolean if a field has been set.

### GetVolumes

`func (o *Instance) GetVolumes() []string`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *Instance) GetVolumesOk() (*[]string, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *Instance) SetVolumes(v []string)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *Instance) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


