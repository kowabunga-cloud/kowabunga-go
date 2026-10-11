# DnsRecordUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The DNS record name. | [optional] 
**Description** | Pointer to **string** | The DNS record description. | [optional] 
**Addresses** | Pointer to **[]string** | A list of IPv4 addresses to be associated to the record. | [optional] 

## Methods

### NewDnsRecordUpdate

`func NewDnsRecordUpdate() *DnsRecordUpdate`

NewDnsRecordUpdate instantiates a new DnsRecordUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordUpdateWithDefaults

`func NewDnsRecordUpdateWithDefaults() *DnsRecordUpdate`

NewDnsRecordUpdateWithDefaults instantiates a new DnsRecordUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DnsRecordUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsRecordUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsRecordUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnsRecordUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DnsRecordUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DnsRecordUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DnsRecordUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DnsRecordUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAddresses

`func (o *DnsRecordUpdate) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *DnsRecordUpdate) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *DnsRecordUpdate) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *DnsRecordUpdate) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


