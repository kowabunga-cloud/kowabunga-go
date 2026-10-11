# KiwiUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Kiwi (Kowabunga Inner Wan Interface) provides edge-network services. name. | [optional] 
**Description** | Pointer to **string** | The Kiwi (Kowabunga Inner Wan Interface) provides edge-network services. description. | [optional] 
**Agents** | Pointer to **[]string** | a list of existing remote agents managing the network gateway. | [optional] 

## Methods

### NewKiwiUpdate

`func NewKiwiUpdate() *KiwiUpdate`

NewKiwiUpdate instantiates a new KiwiUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKiwiUpdateWithDefaults

`func NewKiwiUpdateWithDefaults() *KiwiUpdate`

NewKiwiUpdateWithDefaults instantiates a new KiwiUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KiwiUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KiwiUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KiwiUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KiwiUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *KiwiUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KiwiUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KiwiUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KiwiUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAgents

`func (o *KiwiUpdate) GetAgents() []string`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *KiwiUpdate) GetAgentsOk() (*[]string, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *KiwiUpdate) SetAgents(v []string)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *KiwiUpdate) HasAgents() bool`

HasAgents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


