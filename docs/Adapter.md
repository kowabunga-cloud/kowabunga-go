# Adapter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The network adapter ID (auto-generated). | [optional] 
**Name** | **string** | The network adapter name. | 
**Description** | Pointer to **string** | The network adapter description. | [optional] 
**Mac** | Pointer to **string** | The network adapter hardware address (e.g. 00:11:22:33:44:55). Auto-generated if unspecified. | [optional] 
**Addresses** | Pointer to **[]string** | The network adapter list of associated IPv4 addresses. | [optional] 
**Reserved** | Pointer to **bool** | The network adapter is a reserved adapter (e.g. router), where the same hardware address can be reused over several subnets. | [optional] [default to false]

## Methods

### NewAdapter

`func NewAdapter(name string, ) *Adapter`

NewAdapter instantiates a new Adapter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterWithDefaults

`func NewAdapterWithDefaults() *Adapter`

NewAdapterWithDefaults instantiates a new Adapter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Adapter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Adapter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Adapter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Adapter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Adapter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Adapter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Adapter) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Adapter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Adapter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Adapter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Adapter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMac

`func (o *Adapter) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *Adapter) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *Adapter) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *Adapter) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetAddresses

`func (o *Adapter) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Adapter) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Adapter) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *Adapter) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetReserved

`func (o *Adapter) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *Adapter) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *Adapter) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *Adapter) HasReserved() bool`

HasReserved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


