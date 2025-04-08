# Konvey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Konvey (Kowabunga Network Load-Balancer) ID (auto-generated). | [optional] 
**Name** | Pointer to **string** | The Konvey (Kowabunga Network Load-Balancer) name. | [optional] 
**Description** | Pointer to **string** | The Konvey (Kowabunga Network Load-Balancer) description. | [optional] 
**Vip** | Pointer to **string** | The Konvey (Kowabunga Network Load-Balancer) assigned private virtual IP address (read-only). | [optional] 
**Failover** | Pointer to **bool** | Whether Konvey (Kowabunga Network Load-Balancer) must be deployed in a highly-available replicated state to support service failover. | [optional] [default to true]
**Endpoints** | [**[]KonveyEndpoint**](KonveyEndpoint.md) | The Konvey (Kowabunga Network Load-Balancer) list of load-balanced endpoints. | 

## Methods

### NewKonvey

`func NewKonvey(endpoints []KonveyEndpoint, ) *Konvey`

NewKonvey instantiates a new Konvey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKonveyWithDefaults

`func NewKonveyWithDefaults() *Konvey`

NewKonveyWithDefaults instantiates a new Konvey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Konvey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Konvey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Konvey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Konvey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Konvey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Konvey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Konvey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Konvey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Konvey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Konvey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Konvey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Konvey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVip

`func (o *Konvey) GetVip() string`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *Konvey) GetVipOk() (*string, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *Konvey) SetVip(v string)`

SetVip sets Vip field to given value.

### HasVip

`func (o *Konvey) HasVip() bool`

HasVip returns a boolean if a field has been set.

### GetFailover

`func (o *Konvey) GetFailover() bool`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *Konvey) GetFailoverOk() (*bool, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *Konvey) SetFailover(v bool)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *Konvey) HasFailover() bool`

HasFailover returns a boolean if a field has been set.

### GetEndpoints

`func (o *Konvey) GetEndpoints() []KonveyEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *Konvey) GetEndpointsOk() (*[]KonveyEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *Konvey) SetEndpoints(v []KonveyEndpoint)`

SetEndpoints sets Endpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


