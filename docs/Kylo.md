# Kylo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Kylo ID (auto-generated). | [optional] 
**Name** | **string** | The Kylo name. | 
**Description** | Pointer to **string** | The Kylo description. | [optional] 
**Access** | Pointer to **string** | The Kylo volume access type. | [optional] [default to "RW"]
**Protocols** | Pointer to **[]int32** | The Kylo NFS protocol versions to be supported. | [optional] [default to [3, 4]]
**Endpoint** | Pointer to **string** | The Kylo endpoint FQDN (read-only). | [optional] 
**Size** | Pointer to **int64** | The Kylo volume bytes used (read-only). | [optional] 

## Methods

### NewKylo

`func NewKylo(name string, ) *Kylo`

NewKylo instantiates a new Kylo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKyloWithDefaults

`func NewKyloWithDefaults() *Kylo`

NewKyloWithDefaults instantiates a new Kylo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Kylo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Kylo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Kylo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Kylo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Kylo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Kylo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Kylo) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Kylo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Kylo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Kylo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Kylo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccess

`func (o *Kylo) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *Kylo) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *Kylo) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *Kylo) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetProtocols

`func (o *Kylo) GetProtocols() []int32`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *Kylo) GetProtocolsOk() (*[]int32, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *Kylo) SetProtocols(v []int32)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *Kylo) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetEndpoint

`func (o *Kylo) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *Kylo) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *Kylo) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *Kylo) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetSize

`func (o *Kylo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Kylo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Kylo) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *Kylo) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


