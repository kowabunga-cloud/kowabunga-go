# Kompute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Kompute ID (auto-generated). | [optional] 
**Name** | **string** | The Kompute name. | 
**Description** | Pointer to **string** | The Kompute description. | [optional] 
**Memory** | **int64** | The Kompute memory size (in bytes). | 
**Vcpus** | **int64** | The Kompute number of vCPUs. | 
**Disk** | **int64** | The Kompute OS disk size (in bytes). | 
**DataDisk** | Pointer to **int64** | The Kompute extra data disk size (in bytes). If unspecified, no extra data disk will be assigned. | [optional] [default to 0]
**Ip** | Pointer to **string** | The Kompute assigned private IPv4 address (read-only). | [optional] 

## Methods

### NewKompute

`func NewKompute(name string, memory int64, vcpus int64, disk int64, ) *Kompute`

NewKompute instantiates a new Kompute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKomputeWithDefaults

`func NewKomputeWithDefaults() *Kompute`

NewKomputeWithDefaults instantiates a new Kompute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Kompute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Kompute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Kompute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Kompute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Kompute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Kompute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Kompute) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Kompute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Kompute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Kompute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Kompute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMemory

`func (o *Kompute) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Kompute) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Kompute) SetMemory(v int64)`

SetMemory sets Memory field to given value.


### GetVcpus

`func (o *Kompute) GetVcpus() int64`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *Kompute) GetVcpusOk() (*int64, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *Kompute) SetVcpus(v int64)`

SetVcpus sets Vcpus field to given value.


### GetDisk

`func (o *Kompute) GetDisk() int64`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *Kompute) GetDiskOk() (*int64, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *Kompute) SetDisk(v int64)`

SetDisk sets Disk field to given value.


### GetDataDisk

`func (o *Kompute) GetDataDisk() int64`

GetDataDisk returns the DataDisk field if non-nil, zero value otherwise.

### GetDataDiskOk

`func (o *Kompute) GetDataDiskOk() (*int64, bool)`

GetDataDiskOk returns a tuple with the DataDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDisk

`func (o *Kompute) SetDataDisk(v int64)`

SetDataDisk sets DataDisk field to given value.

### HasDataDisk

`func (o *Kompute) HasDataDisk() bool`

HasDataDisk returns a boolean if a field has been set.

### GetIp

`func (o *Kompute) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Kompute) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Kompute) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Kompute) HasIp() bool`

HasIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


