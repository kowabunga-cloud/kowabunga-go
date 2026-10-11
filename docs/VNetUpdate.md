# VNetUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The virtual network name. | [optional] 
**Description** | Pointer to **string** | The virtual network description. | [optional] 
**Vlan** | Pointer to **int32** | The VLAN identifier (0 if unspecified). | [optional] 
**Interface** | Pointer to **string** | The libvirt&#39;s bridge network interface (brX). | [optional] 

## Methods

### NewVNetUpdate

`func NewVNetUpdate() *VNetUpdate`

NewVNetUpdate instantiates a new VNetUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVNetUpdateWithDefaults

`func NewVNetUpdateWithDefaults() *VNetUpdate`

NewVNetUpdateWithDefaults instantiates a new VNetUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VNetUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VNetUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VNetUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VNetUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VNetUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VNetUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VNetUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VNetUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVlan

`func (o *VNetUpdate) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *VNetUpdate) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *VNetUpdate) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *VNetUpdate) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetInterface

`func (o *VNetUpdate) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *VNetUpdate) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *VNetUpdate) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *VNetUpdate) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


