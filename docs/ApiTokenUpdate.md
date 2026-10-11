# ApiTokenUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The server-to-server authentication security token name. | [optional] 
**Description** | Pointer to **string** | The server-to-server authentication security token description. | [optional] 
**Expire** | Pointer to **bool** | Whether the API token expires at some stage. | [optional] 
**ExpirationDate** | Pointer to **string** | Expiration date of the token (YYYY-MM-DD format). | [optional] 

## Methods

### NewApiTokenUpdate

`func NewApiTokenUpdate() *ApiTokenUpdate`

NewApiTokenUpdate instantiates a new ApiTokenUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenUpdateWithDefaults

`func NewApiTokenUpdateWithDefaults() *ApiTokenUpdate`

NewApiTokenUpdateWithDefaults instantiates a new ApiTokenUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiTokenUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiTokenUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiTokenUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiTokenUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiTokenUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiTokenUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiTokenUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiTokenUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpire

`func (o *ApiTokenUpdate) GetExpire() bool`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *ApiTokenUpdate) GetExpireOk() (*bool, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *ApiTokenUpdate) SetExpire(v bool)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *ApiTokenUpdate) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ApiTokenUpdate) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ApiTokenUpdate) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ApiTokenUpdate) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ApiTokenUpdate) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


