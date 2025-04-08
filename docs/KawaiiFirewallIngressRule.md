# KawaiiFirewallIngressRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | The source IP or CIDR to accept public traffic from. | [optional] [default to "0.0.0.0/0"]
**Protocol** | Pointer to **string** | The transport layer protocol to accept public traffic from. | [optional] [default to "tcp"]
**Ports** | **string** | The port (or list of ports) to accept public traffic from. Ranges are accepted. Format is a-b,c-d (e.g. 443; 22,80,443; 80,443,3000-3005). | 

## Methods

### NewKawaiiFirewallIngressRule

`func NewKawaiiFirewallIngressRule(ports string, ) *KawaiiFirewallIngressRule`

NewKawaiiFirewallIngressRule instantiates a new KawaiiFirewallIngressRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKawaiiFirewallIngressRuleWithDefaults

`func NewKawaiiFirewallIngressRuleWithDefaults() *KawaiiFirewallIngressRule`

NewKawaiiFirewallIngressRuleWithDefaults instantiates a new KawaiiFirewallIngressRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *KawaiiFirewallIngressRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *KawaiiFirewallIngressRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *KawaiiFirewallIngressRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *KawaiiFirewallIngressRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetProtocol

`func (o *KawaiiFirewallIngressRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *KawaiiFirewallIngressRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *KawaiiFirewallIngressRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *KawaiiFirewallIngressRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPorts

`func (o *KawaiiFirewallIngressRule) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *KawaiiFirewallIngressRule) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *KawaiiFirewallIngressRule) SetPorts(v string)`

SetPorts sets Ports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


