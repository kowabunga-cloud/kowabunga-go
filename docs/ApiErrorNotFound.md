# ApiErrorNotFound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** |  | 
**Error** | **string** |  | 
**Reason** | **string** |  | 

## Methods

### NewApiErrorNotFound

`func NewApiErrorNotFound(status int32, error_ string, reason string, ) *ApiErrorNotFound`

NewApiErrorNotFound instantiates a new ApiErrorNotFound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorNotFoundWithDefaults

`func NewApiErrorNotFoundWithDefaults() *ApiErrorNotFound`

NewApiErrorNotFoundWithDefaults instantiates a new ApiErrorNotFound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApiErrorNotFound) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiErrorNotFound) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiErrorNotFound) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetError

`func (o *ApiErrorNotFound) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiErrorNotFound) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiErrorNotFound) SetError(v string)`

SetError sets Error field to given value.


### GetReason

`func (o *ApiErrorNotFound) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ApiErrorNotFound) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ApiErrorNotFound) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


