# TeamUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Kowabunga users team name. | [optional] 
**Description** | Pointer to **string** | The Kowabunga users team description. | [optional] 
**Users** | Pointer to **[]string** | List of user IDs that are part of the team. | [optional] 

## Methods

### NewTeamUpdate

`func NewTeamUpdate() *TeamUpdate`

NewTeamUpdate instantiates a new TeamUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamUpdateWithDefaults

`func NewTeamUpdateWithDefaults() *TeamUpdate`

NewTeamUpdateWithDefaults instantiates a new TeamUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TeamUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TeamUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUsers

`func (o *TeamUpdate) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *TeamUpdate) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *TeamUpdate) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *TeamUpdate) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


