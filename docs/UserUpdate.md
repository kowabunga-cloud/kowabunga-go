# UserUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Kowabunga user name. | [optional] 
**Description** | Pointer to **string** | The Kowabunga user description. | [optional] 
**Email** | Pointer to **string** | User email address. | [optional] 
**Role** | Pointer to **string** | The Kowabunga user role. | [optional] 
**Notifications** | Pointer to **bool** | Whether or not to receive email notifications on events. | [optional] 

## Methods

### NewUserUpdate

`func NewUserUpdate() *UserUpdate`

NewUserUpdate instantiates a new UserUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateWithDefaults

`func NewUserUpdateWithDefaults() *UserUpdate`

NewUserUpdateWithDefaults instantiates a new UserUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UserUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *UserUpdate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserUpdate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserUpdate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserUpdate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRole

`func (o *UserUpdate) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserUpdate) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserUpdate) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserUpdate) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetNotifications

`func (o *UserUpdate) GetNotifications() bool`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *UserUpdate) GetNotificationsOk() (*bool, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *UserUpdate) SetNotifications(v bool)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *UserUpdate) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


