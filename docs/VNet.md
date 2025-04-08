# VNet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The virtual network ID (auto-generated). | [optional] 
**Name** | **string** | The virtual network name. | 
**Description** | Pointer to **string** | The virtual network description. | [optional] 
**Vlan** | Pointer to **int64** | The VLAN identifier (0 if unspecified). | [optional] 
**Interface** | **string** | The libvirt&#39;s bridge network interface (brX). | 
**Private** | Pointer to **bool** | Is the virtual network adapter connected to private (LAN) or public (WAN) physical network ?. | [optional] [default to true]

## Methods

### NewVNet

`func NewVNet(name string, interface_ string, ) *VNet`

NewVNet instantiates a new VNet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVNetWithDefaults

`func NewVNetWithDefaults() *VNet`

NewVNetWithDefaults instantiates a new VNet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VNet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VNet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VNet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VNet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VNet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VNet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VNet) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VNet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VNet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VNet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VNet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVlan

`func (o *VNet) GetVlan() int64`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *VNet) GetVlanOk() (*int64, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *VNet) SetVlan(v int64)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *VNet) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetInterface

`func (o *VNet) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *VNet) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *VNet) SetInterface(v string)`

SetInterface sets Interface field to given value.


### GetPrivate

`func (o *VNet) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *VNet) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *VNet) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *VNet) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


