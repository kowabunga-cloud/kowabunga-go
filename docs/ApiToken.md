# ApiToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The server-to-server authentication security token ID (auto-generated). | [optional] 
**Name** | **string** | The server-to-server authentication security token name. | 
**Description** | Pointer to **string** | The server-to-server authentication security token description. | [optional] 
**Expire** | **bool** | Does the API token expires at some stage ?. | [default to false]
**ExpirationDate** | Pointer to **string** | Expiration date of the token (YYYY-MM-DD format). | [optional] 

## Methods

### NewApiToken

`func NewApiToken(name string, expire bool, ) *ApiToken`

NewApiToken instantiates a new ApiToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenWithDefaults

`func NewApiTokenWithDefaults() *ApiToken`

NewApiTokenWithDefaults instantiates a new ApiToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApiToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiToken) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ApiToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpire

`func (o *ApiToken) GetExpire() bool`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *ApiToken) GetExpireOk() (*bool, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *ApiToken) SetExpire(v bool)`

SetExpire sets Expire field to given value.


### GetExpirationDate

`func (o *ApiToken) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ApiToken) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ApiToken) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ApiToken) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


