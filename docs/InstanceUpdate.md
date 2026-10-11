# InstanceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The virtual machine instance name. | [optional] 
**Description** | Pointer to **string** | The virtual machine instance description. | [optional] 
**Memory** | Pointer to **int64** | The virtual machine instance memory size (in bytes). | [optional] 
**Vcpus** | Pointer to **int32** | The virtual machine instance number of vCPUs. | [optional] 
**Adapters** | Pointer to **[]string** | A list of existing network adapters to be connected to the instance. | [optional] 
**Volumes** | Pointer to **[]string** | A list of existing storage volumes (i.e. disks) to be connected to the instance. | [optional] 
**KmotionEnabled** | Pointer to **bool** | Whether kMotion (live/cold host migration) is enabled for the virtual machine instance. When disabled, the instance remains bound to its original hosting node. | [optional] 

## Methods

### NewInstanceUpdate

`func NewInstanceUpdate() *InstanceUpdate`

NewInstanceUpdate instantiates a new InstanceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceUpdateWithDefaults

`func NewInstanceUpdateWithDefaults() *InstanceUpdate`

NewInstanceUpdateWithDefaults instantiates a new InstanceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstanceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMemory

`func (o *InstanceUpdate) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *InstanceUpdate) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *InstanceUpdate) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *InstanceUpdate) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetVcpus

`func (o *InstanceUpdate) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *InstanceUpdate) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *InstanceUpdate) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *InstanceUpdate) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetAdapters

`func (o *InstanceUpdate) GetAdapters() []string`

GetAdapters returns the Adapters field if non-nil, zero value otherwise.

### GetAdaptersOk

`func (o *InstanceUpdate) GetAdaptersOk() (*[]string, bool)`

GetAdaptersOk returns a tuple with the Adapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapters

`func (o *InstanceUpdate) SetAdapters(v []string)`

SetAdapters sets Adapters field to given value.

### HasAdapters

`func (o *InstanceUpdate) HasAdapters() bool`

HasAdapters returns a boolean if a field has been set.

### GetVolumes

`func (o *InstanceUpdate) GetVolumes() []string`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *InstanceUpdate) GetVolumesOk() (*[]string, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *InstanceUpdate) SetVolumes(v []string)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *InstanceUpdate) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetKmotionEnabled

`func (o *InstanceUpdate) GetKmotionEnabled() bool`

GetKmotionEnabled returns the KmotionEnabled field if non-nil, zero value otherwise.

### GetKmotionEnabledOk

`func (o *InstanceUpdate) GetKmotionEnabledOk() (*bool, bool)`

GetKmotionEnabledOk returns a tuple with the KmotionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmotionEnabled

`func (o *InstanceUpdate) SetKmotionEnabled(v bool)`

SetKmotionEnabled sets KmotionEnabled field to given value.

### HasKmotionEnabled

`func (o *InstanceUpdate) HasKmotionEnabled() bool`

HasKmotionEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


