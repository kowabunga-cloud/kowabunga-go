# ApiErrorUnprocessableEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** |  | 
**Error** | **string** |  | 
**Reason** | **string** |  | 

## Methods

### NewApiErrorUnprocessableEntity

`func NewApiErrorUnprocessableEntity(status int32, error_ string, reason string, ) *ApiErrorUnprocessableEntity`

NewApiErrorUnprocessableEntity instantiates a new ApiErrorUnprocessableEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorUnprocessableEntityWithDefaults

`func NewApiErrorUnprocessableEntityWithDefaults() *ApiErrorUnprocessableEntity`

NewApiErrorUnprocessableEntityWithDefaults instantiates a new ApiErrorUnprocessableEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApiErrorUnprocessableEntity) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiErrorUnprocessableEntity) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiErrorUnprocessableEntity) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetError

`func (o *ApiErrorUnprocessableEntity) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiErrorUnprocessableEntity) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiErrorUnprocessableEntity) SetError(v string)`

SetError sets Error field to given value.


### GetReason

`func (o *ApiErrorUnprocessableEntity) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ApiErrorUnprocessableEntity) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ApiErrorUnprocessableEntity) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


