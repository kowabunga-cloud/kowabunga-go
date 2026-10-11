# ApiErrorInternalServerError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** |  | 
**Error** | **string** |  | 
**Reason** | **string** |  | 

## Methods

### NewApiErrorInternalServerError

`func NewApiErrorInternalServerError(status int32, error_ string, reason string, ) *ApiErrorInternalServerError`

NewApiErrorInternalServerError instantiates a new ApiErrorInternalServerError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorInternalServerErrorWithDefaults

`func NewApiErrorInternalServerErrorWithDefaults() *ApiErrorInternalServerError`

NewApiErrorInternalServerErrorWithDefaults instantiates a new ApiErrorInternalServerError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApiErrorInternalServerError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiErrorInternalServerError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiErrorInternalServerError) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetError

`func (o *ApiErrorInternalServerError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiErrorInternalServerError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiErrorInternalServerError) SetError(v string)`

SetError sets Error field to given value.


### GetReason

`func (o *ApiErrorInternalServerError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ApiErrorInternalServerError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ApiErrorInternalServerError) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


