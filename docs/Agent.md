# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Kowabunga remote agent ID (auto-generated). | [optional] 
**Name** | **string** | The Kowabunga remote agent name. | 
**Description** | Pointer to **string** | The Kowabunga remote agent description. | [optional] 
**Type** | **string** | The Kowabunga agent type. | 

## Methods

### NewAgent

`func NewAgent(name string, type_ string, ) *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Agent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Agent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Agent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Agent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Agent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Agent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Agent) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Agent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Agent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Agent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Agent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *Agent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Agent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Agent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


