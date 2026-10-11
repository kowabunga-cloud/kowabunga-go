# AgentUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Kowabunga remote agent name. | [optional] 
**Description** | Pointer to **string** | The Kowabunga remote agent description. | [optional] 

## Methods

### NewAgentUpdate

`func NewAgentUpdate() *AgentUpdate`

NewAgentUpdate instantiates a new AgentUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentUpdateWithDefaults

`func NewAgentUpdateWithDefaults() *AgentUpdate`

NewAgentUpdateWithDefaults instantiates a new AgentUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AgentUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AgentUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


