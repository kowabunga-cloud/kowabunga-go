# Password

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | The plain-text password. | 

## Methods

### NewPassword

`func NewPassword(value string, ) *Password`

NewPassword instantiates a new Password object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordWithDefaults

`func NewPasswordWithDefaults() *Password`

NewPasswordWithDefaults instantiates a new Password object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Password) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Password) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Password) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


