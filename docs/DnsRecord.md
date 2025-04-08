# DnsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The DNS record ID (auto-generated). | [optional] 
**Name** | **string** | The DNS record name. | 
**Description** | Pointer to **string** | The DNS record description. | [optional] 
**Domain** | Pointer to **string** | The DNS record associated domain (inherited from associated project). | [optional] 
**Addresses** | **[]string** | A list of IPv4 addresses to be associated to the record. | 

## Methods

### NewDnsRecord

`func NewDnsRecord(name string, addresses []string, ) *DnsRecord`

NewDnsRecord instantiates a new DnsRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordWithDefaults

`func NewDnsRecordWithDefaults() *DnsRecord`

NewDnsRecordWithDefaults instantiates a new DnsRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DnsRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DnsRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DnsRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsRecord) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DnsRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DnsRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DnsRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DnsRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *DnsRecord) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DnsRecord) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DnsRecord) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DnsRecord) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetAddresses

`func (o *DnsRecord) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *DnsRecord) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *DnsRecord) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


