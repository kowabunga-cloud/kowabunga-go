# AdapterUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The network adapter name. | [optional] 
**Description** | Pointer to **string** | The network adapter description. | [optional] 
**Mac** | Pointer to **string** | The network adapter hardware address (e.g. 00:11:22:33:44:55). Auto-generated if unspecified. | [optional] 
**Addresses** | Pointer to **[]string** | The network adapter list of associated IPv4 addresses. | [optional] 
**Reserved** | Pointer to **bool** | The network adapter is a reserved adapter (e.g. router), where the same hardware address can be reused over several subnets. | [optional] 

## Methods

### NewAdapterUpdate

`func NewAdapterUpdate() *AdapterUpdate`

NewAdapterUpdate instantiates a new AdapterUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterUpdateWithDefaults

`func NewAdapterUpdateWithDefaults() *AdapterUpdate`

NewAdapterUpdateWithDefaults instantiates a new AdapterUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdapterUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdapterUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdapterUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdapterUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AdapterUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdapterUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdapterUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdapterUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMac

`func (o *AdapterUpdate) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *AdapterUpdate) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *AdapterUpdate) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *AdapterUpdate) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetAddresses

`func (o *AdapterUpdate) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *AdapterUpdate) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *AdapterUpdate) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *AdapterUpdate) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetReserved

`func (o *AdapterUpdate) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *AdapterUpdate) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *AdapterUpdate) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *AdapterUpdate) HasReserved() bool`

HasReserved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


