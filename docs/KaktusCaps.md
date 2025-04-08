# KaktusCaps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | [**KaktusCPU**](KaktusCPU.md) |  | 
**Memory** | **int64** | The Kaktus computing node memory size (bytes). | 

## Methods

### NewKaktusCaps

`func NewKaktusCaps(cpu KaktusCPU, memory int64, ) *KaktusCaps`

NewKaktusCaps instantiates a new KaktusCaps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKaktusCapsWithDefaults

`func NewKaktusCapsWithDefaults() *KaktusCaps`

NewKaktusCapsWithDefaults instantiates a new KaktusCaps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *KaktusCaps) GetCpu() KaktusCPU`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *KaktusCaps) GetCpuOk() (*KaktusCPU, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *KaktusCaps) SetCpu(v KaktusCPU)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *KaktusCaps) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *KaktusCaps) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *KaktusCaps) SetMemory(v int64)`

SetMemory sets Memory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


