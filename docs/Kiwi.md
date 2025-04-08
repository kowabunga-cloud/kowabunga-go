# Kiwi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Kiwi (Kowabunga Inner Wan Interface) provides edge-network services. ID (auto-generated). | [optional] 
**Name** | **string** | The Kiwi (Kowabunga Inner Wan Interface) provides edge-network services. name. | 
**Description** | Pointer to **string** | The Kiwi (Kowabunga Inner Wan Interface) provides edge-network services. description. | [optional] 
**Agents** | Pointer to **[]string** | a list of existing remote agents managing the network gateway. | [optional] 

## Methods

### NewKiwi

`func NewKiwi(name string, ) *Kiwi`

NewKiwi instantiates a new Kiwi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKiwiWithDefaults

`func NewKiwiWithDefaults() *Kiwi`

NewKiwiWithDefaults instantiates a new Kiwi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Kiwi) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Kiwi) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Kiwi) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Kiwi) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Kiwi) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Kiwi) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Kiwi) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Kiwi) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Kiwi) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Kiwi) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Kiwi) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAgents

`func (o *Kiwi) GetAgents() []string`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *Kiwi) GetAgentsOk() (*[]string, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *Kiwi) SetAgents(v []string)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *Kiwi) HasAgents() bool`

HasAgents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


