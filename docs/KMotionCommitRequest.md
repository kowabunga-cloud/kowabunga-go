# KMotionCommitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | **string** | The ID of the previously computed kMotion plan (see kMotion plan webservice) to execute. | 

## Methods

### NewKMotionCommitRequest

`func NewKMotionCommitRequest(plan string, ) *KMotionCommitRequest`

NewKMotionCommitRequest instantiates a new KMotionCommitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKMotionCommitRequestWithDefaults

`func NewKMotionCommitRequestWithDefaults() *KMotionCommitRequest`

NewKMotionCommitRequestWithDefaults instantiates a new KMotionCommitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *KMotionCommitRequest) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *KMotionCommitRequest) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *KMotionCommitRequest) SetPlan(v string)`

SetPlan sets Plan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


