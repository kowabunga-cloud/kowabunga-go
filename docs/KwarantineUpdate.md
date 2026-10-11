# KwarantineUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Kwarantine name. | [optional] 
**Description** | Pointer to **string** | The Kwarantine description. | [optional] 
**Policy** | Pointer to **string** | The anti-affinity enforcement policy. Use \&quot;host\&quot; to guarantee no two group members are ever scheduled on the same Kaktus computing node, or \&quot;zone\&quot; to guarantee no two group members are ever scheduled within the same availability zone. | [optional] 

## Methods

### NewKwarantineUpdate

`func NewKwarantineUpdate() *KwarantineUpdate`

NewKwarantineUpdate instantiates a new KwarantineUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKwarantineUpdateWithDefaults

`func NewKwarantineUpdateWithDefaults() *KwarantineUpdate`

NewKwarantineUpdateWithDefaults instantiates a new KwarantineUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KwarantineUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KwarantineUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KwarantineUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KwarantineUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *KwarantineUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KwarantineUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KwarantineUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KwarantineUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicy

`func (o *KwarantineUpdate) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *KwarantineUpdate) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *KwarantineUpdate) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *KwarantineUpdate) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


