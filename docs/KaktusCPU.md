# KaktusCPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | **string** | The Kaktus computing node CPU architecture. | 
**Model** | **string** | The Kaktus computing node CPU model. | 
**Vendor** | **string** | The Kaktus computing node CPU vendor. | 
**Sockets** | **int32** | The Kaktus computing node CPU number of sockets. | 
**Cores** | **int32** | The Kaktus computing node CPU number of cores. | 
**Threads** | **int32** | The Kaktus computing node CPU number of threads. | 

## Methods

### NewKaktusCPU

`func NewKaktusCPU(arch string, model string, vendor string, sockets int32, cores int32, threads int32, ) *KaktusCPU`

NewKaktusCPU instantiates a new KaktusCPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKaktusCPUWithDefaults

`func NewKaktusCPUWithDefaults() *KaktusCPU`

NewKaktusCPUWithDefaults instantiates a new KaktusCPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *KaktusCPU) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *KaktusCPU) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *KaktusCPU) SetArch(v string)`

SetArch sets Arch field to given value.


### GetModel

`func (o *KaktusCPU) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *KaktusCPU) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *KaktusCPU) SetModel(v string)`

SetModel sets Model field to given value.


### GetVendor

`func (o *KaktusCPU) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *KaktusCPU) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *KaktusCPU) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetSockets

`func (o *KaktusCPU) GetSockets() int32`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *KaktusCPU) GetSocketsOk() (*int32, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *KaktusCPU) SetSockets(v int32)`

SetSockets sets Sockets field to given value.


### GetCores

`func (o *KaktusCPU) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *KaktusCPU) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *KaktusCPU) SetCores(v int32)`

SetCores sets Cores field to given value.


### GetThreads

`func (o *KaktusCPU) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *KaktusCPU) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *KaktusCPU) SetThreads(v int32)`

SetThreads sets Threads field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


