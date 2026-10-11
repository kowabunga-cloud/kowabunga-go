# SubnetUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The network subnet name. | [optional] 
**Description** | Pointer to **string** | The network subnet description. | [optional] 
**Gateway** | Pointer to **string** | The network subnet router/gateway IP address (e.g. 192.168.0.254). | [optional] 
**Dns** | Pointer to **string** | The network subnet DNS server IP address (gateway value if unspecified). | [optional] 
**ExtraRoutes** | Pointer to **[]string** | The list of extra routes to be accessed through designated gateway (format is 10.0.0.0/8). | [optional] 
**Reserved** | Pointer to [**[]IpRange**](IpRange.md) | The network subnet reserved IPv4 ranges (i.e. no IP address can be assigned from there). | [optional] 
**GwPool** | Pointer to [**[]IpRange**](IpRange.md) | The network subnet IPv4 ranges reserved for per-zone local network gateways (range size must be at least equal to region number of zones). | [optional] 
**Application** | Pointer to **string** | Optional application service type. | [optional] 

## Methods

### NewSubnetUpdate

`func NewSubnetUpdate() *SubnetUpdate`

NewSubnetUpdate instantiates a new SubnetUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetUpdateWithDefaults

`func NewSubnetUpdateWithDefaults() *SubnetUpdate`

NewSubnetUpdateWithDefaults instantiates a new SubnetUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SubnetUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubnetUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubnetUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubnetUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SubnetUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubnetUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubnetUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubnetUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGateway

`func (o *SubnetUpdate) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *SubnetUpdate) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *SubnetUpdate) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *SubnetUpdate) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetDns

`func (o *SubnetUpdate) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *SubnetUpdate) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *SubnetUpdate) SetDns(v string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *SubnetUpdate) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetExtraRoutes

`func (o *SubnetUpdate) GetExtraRoutes() []string`

GetExtraRoutes returns the ExtraRoutes field if non-nil, zero value otherwise.

### GetExtraRoutesOk

`func (o *SubnetUpdate) GetExtraRoutesOk() (*[]string, bool)`

GetExtraRoutesOk returns a tuple with the ExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes

`func (o *SubnetUpdate) SetExtraRoutes(v []string)`

SetExtraRoutes sets ExtraRoutes field to given value.

### HasExtraRoutes

`func (o *SubnetUpdate) HasExtraRoutes() bool`

HasExtraRoutes returns a boolean if a field has been set.

### GetReserved

`func (o *SubnetUpdate) GetReserved() []IpRange`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *SubnetUpdate) GetReservedOk() (*[]IpRange, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *SubnetUpdate) SetReserved(v []IpRange)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *SubnetUpdate) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetGwPool

`func (o *SubnetUpdate) GetGwPool() []IpRange`

GetGwPool returns the GwPool field if non-nil, zero value otherwise.

### GetGwPoolOk

`func (o *SubnetUpdate) GetGwPoolOk() (*[]IpRange, bool)`

GetGwPoolOk returns a tuple with the GwPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwPool

`func (o *SubnetUpdate) SetGwPool(v []IpRange)`

SetGwPool sets GwPool field to given value.

### HasGwPool

`func (o *SubnetUpdate) HasGwPool() bool`

HasGwPool returns a boolean if a field has been set.

### GetApplication

`func (o *SubnetUpdate) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *SubnetUpdate) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *SubnetUpdate) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *SubnetUpdate) HasApplication() bool`

HasApplication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


