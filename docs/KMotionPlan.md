# KMotionPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The kMotion plan ID (auto-generated) (read-only). | [optional] [readonly] 
**Instance** | **string** | The ID of the virtual machine instance this kMotion plan applies to. | 
**Source** | **string** | The ID of the Kaktus computing node currently hosting the instance. | 
**Target** | **string** | The ID of the Kaktus computing node elected to host the instance once the kMotion completes. | 
**Live** | **bool** | Whether the plan is for a live kMotion, without any service interruption, or the instance will be shutdown and restarted on its new host. | 

## Methods

### NewKMotionPlan

`func NewKMotionPlan(instance string, source string, target string, live bool, ) *KMotionPlan`

NewKMotionPlan instantiates a new KMotionPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKMotionPlanWithDefaults

`func NewKMotionPlanWithDefaults() *KMotionPlan`

NewKMotionPlanWithDefaults instantiates a new KMotionPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KMotionPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KMotionPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KMotionPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KMotionPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstance

`func (o *KMotionPlan) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *KMotionPlan) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *KMotionPlan) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetSource

`func (o *KMotionPlan) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *KMotionPlan) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *KMotionPlan) SetSource(v string)`

SetSource sets Source field to given value.


### GetTarget

`func (o *KMotionPlan) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *KMotionPlan) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *KMotionPlan) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetLive

`func (o *KMotionPlan) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *KMotionPlan) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *KMotionPlan) SetLive(v bool)`

SetLive sets Live field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


