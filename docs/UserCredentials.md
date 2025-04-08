# UserCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The user email address for login. | 
**Password** | **string** | The user password for login. | 
**Jwt** | Pointer to **string** | The resulting server-generated JWT token for Bearer Authentication (read-only). | [optional] 

## Methods

### NewUserCredentials

`func NewUserCredentials(email string, password string, ) *UserCredentials`

NewUserCredentials instantiates a new UserCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCredentialsWithDefaults

`func NewUserCredentialsWithDefaults() *UserCredentials`

NewUserCredentialsWithDefaults instantiates a new UserCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserCredentials) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCredentials) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCredentials) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *UserCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetJwt

`func (o *UserCredentials) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *UserCredentials) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *UserCredentials) SetJwt(v string)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *UserCredentials) HasJwt() bool`

HasJwt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


