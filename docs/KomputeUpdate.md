# KomputeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Kompute name. | [optional] 
**Description** | Pointer to **string** | The Kompute description. | [optional] 
**Memory** | Pointer to **int64** | The Kompute memory size (in bytes). | [optional] 
**Vcpus** | Pointer to **int32** | The Kompute number of vCPUs. | [optional] 
**Disk** | Pointer to **int64** | The Kompute OS disk size (in bytes). | [optional] 
**DataDisk** | Pointer to **int64** | The Kompute extra data disk size (in bytes). If unspecified, no extra data disk will be assigned. | [optional] 
**KmotionEnabled** | Pointer to **bool** | Whether kMotion (live/cold host migration) is enabled for the Kompute. When disabled, the instance remains bound to its original hosting node. | [optional] 

## Methods

### NewKomputeUpdate

`func NewKomputeUpdate() *KomputeUpdate`

NewKomputeUpdate instantiates a new KomputeUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKomputeUpdateWithDefaults

`func NewKomputeUpdateWithDefaults() *KomputeUpdate`

NewKomputeUpdateWithDefaults instantiates a new KomputeUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KomputeUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KomputeUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KomputeUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KomputeUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *KomputeUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KomputeUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KomputeUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KomputeUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMemory

`func (o *KomputeUpdate) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *KomputeUpdate) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *KomputeUpdate) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *KomputeUpdate) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetVcpus

`func (o *KomputeUpdate) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *KomputeUpdate) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *KomputeUpdate) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *KomputeUpdate) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetDisk

`func (o *KomputeUpdate) GetDisk() int64`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *KomputeUpdate) GetDiskOk() (*int64, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *KomputeUpdate) SetDisk(v int64)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *KomputeUpdate) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetDataDisk

`func (o *KomputeUpdate) GetDataDisk() int64`

GetDataDisk returns the DataDisk field if non-nil, zero value otherwise.

### GetDataDiskOk

`func (o *KomputeUpdate) GetDataDiskOk() (*int64, bool)`

GetDataDiskOk returns a tuple with the DataDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDisk

`func (o *KomputeUpdate) SetDataDisk(v int64)`

SetDataDisk sets DataDisk field to given value.

### HasDataDisk

`func (o *KomputeUpdate) HasDataDisk() bool`

HasDataDisk returns a boolean if a field has been set.

### GetKmotionEnabled

`func (o *KomputeUpdate) GetKmotionEnabled() bool`

GetKmotionEnabled returns the KmotionEnabled field if non-nil, zero value otherwise.

### GetKmotionEnabledOk

`func (o *KomputeUpdate) GetKmotionEnabledOk() (*bool, bool)`

GetKmotionEnabledOk returns a tuple with the KmotionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmotionEnabled

`func (o *KomputeUpdate) SetKmotionEnabled(v bool)`

SetKmotionEnabled sets KmotionEnabled field to given value.

### HasKmotionEnabled

`func (o *KomputeUpdate) HasKmotionEnabled() bool`

HasKmotionEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


