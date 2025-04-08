# KonveyEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The Konvey (Kowabunga Network Load-Balancer) endpoint name. | 
**Port** | **int64** | The port to be exposed. | 
**Protocol** | **string** | The transport layer protocol to be exposed. | [default to "tcp"]
**Backends** | [**KonveyBackends**](KonveyBackends.md) |  | 

## Methods

### NewKonveyEndpoint

`func NewKonveyEndpoint(name string, port int64, protocol string, backends KonveyBackends, ) *KonveyEndpoint`

NewKonveyEndpoint instantiates a new KonveyEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKonveyEndpointWithDefaults

`func NewKonveyEndpointWithDefaults() *KonveyEndpoint`

NewKonveyEndpointWithDefaults instantiates a new KonveyEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KonveyEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KonveyEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KonveyEndpoint) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *KonveyEndpoint) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KonveyEndpoint) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KonveyEndpoint) SetPort(v int64)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *KonveyEndpoint) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *KonveyEndpoint) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *KonveyEndpoint) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetBackends

`func (o *KonveyEndpoint) GetBackends() KonveyBackends`

GetBackends returns the Backends field if non-nil, zero value otherwise.

### GetBackendsOk

`func (o *KonveyEndpoint) GetBackendsOk() (*KonveyBackends, bool)`

GetBackendsOk returns a tuple with the Backends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackends

`func (o *KonveyEndpoint) SetBackends(v KonveyBackends)`

SetBackends sets Backends field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


