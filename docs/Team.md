# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Kowabunga users team ID (auto-generated). | [optional] 
**Name** | **string** | The Kowabunga users team name. | 
**Description** | Pointer to **string** | The Kowabunga users team description. | [optional] 
**Users** | **[]string** | List of user IDs that are part of the team. | 

## Methods

### NewTeam

`func NewTeam(name string, users []string, ) *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Team) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Team) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Team) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Team) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Team) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Team) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Team) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Team) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Team) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Team) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Team) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUsers

`func (o *Team) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Team) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Team) SetUsers(v []string)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


