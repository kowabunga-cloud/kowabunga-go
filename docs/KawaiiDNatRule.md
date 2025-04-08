# KawaiiDNatRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | **string** | Target private IP address to forward public traffic to. | 
**Protocol** | Pointer to **string** | The transport layer protocol to forward public traffic to. | [optional] [default to "tcp"]
**Ports** | **string** | The port (or list of ports) to forward public traffic from. Ranges are accepted. Format is a-b,c-d (e.g. 443; 22,80,443; 80,443,3000-3005). | 

## Methods

### NewKawaiiDNatRule

`func NewKawaiiDNatRule(destination string, ports string, ) *KawaiiDNatRule`

NewKawaiiDNatRule instantiates a new KawaiiDNatRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKawaiiDNatRuleWithDefaults

`func NewKawaiiDNatRuleWithDefaults() *KawaiiDNatRule`

NewKawaiiDNatRuleWithDefaults instantiates a new KawaiiDNatRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *KawaiiDNatRule) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *KawaiiDNatRule) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *KawaiiDNatRule) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetProtocol

`func (o *KawaiiDNatRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *KawaiiDNatRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *KawaiiDNatRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *KawaiiDNatRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPorts

`func (o *KawaiiDNatRule) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *KawaiiDNatRule) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *KawaiiDNatRule) SetPorts(v string)`

SetPorts sets Ports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


