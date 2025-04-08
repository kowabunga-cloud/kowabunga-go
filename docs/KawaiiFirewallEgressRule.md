# KawaiiFirewallEgressRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | The destination IP or CIDR to accept/drop public traffic to. | [optional] [default to "0.0.0.0/0"]
**Protocol** | Pointer to **string** | The transport layer protocol to accept/drop public traffic to. | [optional] [default to "tcp"]
**Ports** | **string** | The port (or list of ports) to accept/drop public traffic from. Ranges are accepted. Format is a-b,c-d (e.g. 443; 22,80,443; 80,443,3000-3005). | 

## Methods

### NewKawaiiFirewallEgressRule

`func NewKawaiiFirewallEgressRule(ports string, ) *KawaiiFirewallEgressRule`

NewKawaiiFirewallEgressRule instantiates a new KawaiiFirewallEgressRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKawaiiFirewallEgressRuleWithDefaults

`func NewKawaiiFirewallEgressRuleWithDefaults() *KawaiiFirewallEgressRule`

NewKawaiiFirewallEgressRuleWithDefaults instantiates a new KawaiiFirewallEgressRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *KawaiiFirewallEgressRule) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *KawaiiFirewallEgressRule) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *KawaiiFirewallEgressRule) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *KawaiiFirewallEgressRule) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetProtocol

`func (o *KawaiiFirewallEgressRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *KawaiiFirewallEgressRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *KawaiiFirewallEgressRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *KawaiiFirewallEgressRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPorts

`func (o *KawaiiFirewallEgressRule) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *KawaiiFirewallEgressRule) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *KawaiiFirewallEgressRule) SetPorts(v string)`

SetPorts sets Ports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


