# KMotionPlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | How the destination Kaktus computing node must be elected. Use \&quot;manual\&quot; to explicitly target a given node (see kaktus), \&quot;zone\&quot; to automatically elect the least busy node within the instance own availability zone, \&quot;region\&quot; to automatically elect the least busy node anywhere within the instance region, or \&quot;auto\&quot; to let Kowabunga fully decide the best possible destination across the region, regardless of the number of availability zones it contains. | [optional] [default to "auto"]
**Kaktus** | Pointer to **string** | The explicit target Kaktus computing node ID to kMotion the instance to. Only used (and required) when destination is set to manual. | [optional] 
**Live** | Pointer to **bool** | Whether the kMotion must be performed live, without any service interruption, or the instance can instead be shutdown and restarted on its new host. | [optional] [default to true]

## Methods

### NewKMotionPlanRequest

`func NewKMotionPlanRequest() *KMotionPlanRequest`

NewKMotionPlanRequest instantiates a new KMotionPlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKMotionPlanRequestWithDefaults

`func NewKMotionPlanRequestWithDefaults() *KMotionPlanRequest`

NewKMotionPlanRequestWithDefaults instantiates a new KMotionPlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *KMotionPlanRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *KMotionPlanRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *KMotionPlanRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *KMotionPlanRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetKaktus

`func (o *KMotionPlanRequest) GetKaktus() string`

GetKaktus returns the Kaktus field if non-nil, zero value otherwise.

### GetKaktusOk

`func (o *KMotionPlanRequest) GetKaktusOk() (*string, bool)`

GetKaktusOk returns a tuple with the Kaktus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKaktus

`func (o *KMotionPlanRequest) SetKaktus(v string)`

SetKaktus sets Kaktus field to given value.

### HasKaktus

`func (o *KMotionPlanRequest) HasKaktus() bool`

HasKaktus returns a boolean if a field has been set.

### GetLive

`func (o *KMotionPlanRequest) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *KMotionPlanRequest) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *KMotionPlanRequest) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *KMotionPlanRequest) HasLive() bool`

HasLive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


