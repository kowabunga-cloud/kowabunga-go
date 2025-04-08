# ApiErrorConflict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** |  | 
**Error** | **string** |  | 
**Reason** | **string** |  | 

## Methods

### NewApiErrorConflict

`func NewApiErrorConflict(status int32, error_ string, reason string, ) *ApiErrorConflict`

NewApiErrorConflict instantiates a new ApiErrorConflict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorConflictWithDefaults

`func NewApiErrorConflictWithDefaults() *ApiErrorConflict`

NewApiErrorConflictWithDefaults instantiates a new ApiErrorConflict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApiErrorConflict) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiErrorConflict) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiErrorConflict) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetError

`func (o *ApiErrorConflict) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiErrorConflict) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiErrorConflict) SetError(v string)`

SetError sets Error field to given value.


### GetReason

`func (o *ApiErrorConflict) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ApiErrorConflict) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ApiErrorConflict) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


