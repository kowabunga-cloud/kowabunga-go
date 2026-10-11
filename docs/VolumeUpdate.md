# VolumeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The storage volume name. | [optional] 
**Description** | Pointer to **string** | The storage volume description. | [optional] 
**Size** | Pointer to **int64** | The storage volume size (in bytes), can only be increased. | [optional] 

## Methods

### NewVolumeUpdate

`func NewVolumeUpdate() *VolumeUpdate`

NewVolumeUpdate instantiates a new VolumeUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeUpdateWithDefaults

`func NewVolumeUpdateWithDefaults() *VolumeUpdate`

NewVolumeUpdateWithDefaults instantiates a new VolumeUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VolumeUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VolumeUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VolumeUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VolumeUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VolumeUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSize

`func (o *VolumeUpdate) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumeUpdate) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumeUpdate) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *VolumeUpdate) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


