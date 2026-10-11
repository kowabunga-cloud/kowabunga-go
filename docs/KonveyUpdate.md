# KonveyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Konvey (Kowabunga Network Load-Balancer) name, prefixed with konvey- (project name if unspecified). | [optional] 
**Description** | Pointer to **string** | The Konvey (Kowabunga Network Load-Balancer) description. | [optional] 
**Endpoints** | Pointer to [**[]KonveyEndpoint**](KonveyEndpoint.md) | The Konvey (Kowabunga Network Load-Balancer) list of load-balanced endpoints. | [optional] 

## Methods

### NewKonveyUpdate

`func NewKonveyUpdate() *KonveyUpdate`

NewKonveyUpdate instantiates a new KonveyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKonveyUpdateWithDefaults

`func NewKonveyUpdateWithDefaults() *KonveyUpdate`

NewKonveyUpdateWithDefaults instantiates a new KonveyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KonveyUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KonveyUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KonveyUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KonveyUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *KonveyUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KonveyUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KonveyUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KonveyUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpoints

`func (o *KonveyUpdate) GetEndpoints() []KonveyEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *KonveyUpdate) GetEndpointsOk() (*[]KonveyEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *KonveyUpdate) SetEndpoints(v []KonveyEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *KonveyUpdate) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


