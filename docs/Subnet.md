# Subnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The network subnet ID (auto-generated). | [optional] 
**Name** | **string** | The network subnet name. | 
**Description** | Pointer to **string** | The network subnet description. | [optional] 
**Cidr** | **string** | The network subnet CIDR (e.g. 192.168.0.0/24). | 
**Gateway** | **string** | The network subnet router/gateway IP address (e.g. 192.168.0.254). | 
**Dns** | Pointer to **string** | The network subnet DNS server IP address (gateway value if unspecified). | [optional] 
**ExtraRoutes** | Pointer to **[]string** | The list of extra routes to be access through designated gateway (format is 10.0.0.0/8). | [optional] 
**Reserved** | Pointer to [**[]IpRange**](IpRange.md) | The network subnet reserved IPv4 ranges (i.e. no IP address can be assigned from there). | [optional] 
**GwPool** | Pointer to [**[]IpRange**](IpRange.md) | The network subnet IPv4 ranges reserved for per-zone local network gateways (range size must be at least equal to region number of zones). | [optional] 
**Application** | Pointer to **string** | Optional application service type. | [optional] [default to "user"]

## Methods

### NewSubnet

`func NewSubnet(name string, cidr string, gateway string, ) *Subnet`

NewSubnet instantiates a new Subnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetWithDefaults

`func NewSubnetWithDefaults() *Subnet`

NewSubnetWithDefaults instantiates a new Subnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subnet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subnet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subnet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Subnet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Subnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subnet) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Subnet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Subnet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Subnet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Subnet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCidr

`func (o *Subnet) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *Subnet) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *Subnet) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetGateway

`func (o *Subnet) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Subnet) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Subnet) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetDns

`func (o *Subnet) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *Subnet) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *Subnet) SetDns(v string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *Subnet) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetExtraRoutes

`func (o *Subnet) GetExtraRoutes() []string`

GetExtraRoutes returns the ExtraRoutes field if non-nil, zero value otherwise.

### GetExtraRoutesOk

`func (o *Subnet) GetExtraRoutesOk() (*[]string, bool)`

GetExtraRoutesOk returns a tuple with the ExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRoutes

`func (o *Subnet) SetExtraRoutes(v []string)`

SetExtraRoutes sets ExtraRoutes field to given value.

### HasExtraRoutes

`func (o *Subnet) HasExtraRoutes() bool`

HasExtraRoutes returns a boolean if a field has been set.

### GetReserved

`func (o *Subnet) GetReserved() []IpRange`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *Subnet) GetReservedOk() (*[]IpRange, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *Subnet) SetReserved(v []IpRange)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *Subnet) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetGwPool

`func (o *Subnet) GetGwPool() []IpRange`

GetGwPool returns the GwPool field if non-nil, zero value otherwise.

### GetGwPoolOk

`func (o *Subnet) GetGwPoolOk() (*[]IpRange, bool)`

GetGwPoolOk returns a tuple with the GwPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwPool

`func (o *Subnet) SetGwPool(v []IpRange)`

SetGwPool sets GwPool field to given value.

### HasGwPool

`func (o *Subnet) HasGwPool() bool`

HasGwPool returns a boolean if a field has been set.

### GetApplication

`func (o *Subnet) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *Subnet) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *Subnet) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *Subnet) HasApplication() bool`

HasApplication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


