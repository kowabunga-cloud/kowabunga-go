# InstanceState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | The virtual machine instance state. | 
**Reason** | **string** | The virtual machine instance reason of the state. | 

## Methods

### NewInstanceState

`func NewInstanceState(state string, reason string, ) *InstanceState`

NewInstanceState instantiates a new InstanceState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceStateWithDefaults

`func NewInstanceStateWithDefaults() *InstanceState`

NewInstanceStateWithDefaults instantiates a new InstanceState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *InstanceState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InstanceState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InstanceState) SetState(v string)`

SetState sets State field to given value.


### GetReason

`func (o *InstanceState) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InstanceState) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InstanceState) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


