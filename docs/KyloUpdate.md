# KyloUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Kylo name. | [optional] 
**Description** | Pointer to **string** | The Kylo description. | [optional] 
**Access** | Pointer to **string** | The Kylo volume access type. | [optional] 
**Protocols** | Pointer to **[]int32** | The Kylo NFS protocol versions to be supported. | [optional] 

## Methods

### NewKyloUpdate

`func NewKyloUpdate() *KyloUpdate`

NewKyloUpdate instantiates a new KyloUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKyloUpdateWithDefaults

`func NewKyloUpdateWithDefaults() *KyloUpdate`

NewKyloUpdateWithDefaults instantiates a new KyloUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KyloUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KyloUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KyloUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KyloUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *KyloUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KyloUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KyloUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KyloUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccess

`func (o *KyloUpdate) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *KyloUpdate) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *KyloUpdate) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *KyloUpdate) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetProtocols

`func (o *KyloUpdate) GetProtocols() []int32`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *KyloUpdate) GetProtocolsOk() (*[]int32, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *KyloUpdate) SetProtocols(v []int32)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *KyloUpdate) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


