# KonveyBackends

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | **[]string** | The Konvey (Kowabunga Network Load-Balancer) endpoint list of load-balanced backend hosts. | 
**Port** | **int64** | The Konvey (Kowabunga Network Load-Balancer) endpoint backend service port. | 

## Methods

### NewKonveyBackends

`func NewKonveyBackends(hosts []string, port int64, ) *KonveyBackends`

NewKonveyBackends instantiates a new KonveyBackends object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKonveyBackendsWithDefaults

`func NewKonveyBackendsWithDefaults() *KonveyBackends`

NewKonveyBackendsWithDefaults instantiates a new KonveyBackends object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *KonveyBackends) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *KonveyBackends) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *KonveyBackends) SetHosts(v []string)`

SetHosts sets Hosts field to given value.


### GetPort

`func (o *KonveyBackends) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KonveyBackends) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KonveyBackends) SetPort(v int64)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


