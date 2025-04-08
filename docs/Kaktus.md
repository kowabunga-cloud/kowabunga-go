# Kaktus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Kaktus computing node ID (auto-generated). | [optional] 
**Name** | **string** | The Kaktus computing node name. | 
**Description** | Pointer to **string** | The Kaktus computing node description. | [optional] 
**CpuCost** | Pointer to [**Cost**](Cost.md) |  | [optional] 
**MemoryCost** | Pointer to [**Cost**](Cost.md) |  | [optional] 
**OvercommitCpuRatio** | Pointer to **int64** | The Kaktus node CPU resource over-commit ratio. Overcommitting CPU resources for VMs means allocating more virtual CPUs (vCPUs) to the virtual machines (VMs) than the physical cores available on the node. This can help optimize the utilization of the node CPU and increase the density of VMs per node. | [optional] [default to 3]
**OvercommitMemoryRatio** | Pointer to **int64** | The Kaktus node memory resource over-commit ratio. Memory overcommitment is a concept in computing that covers the assignment of more memory to virtual computing devices (or processes) than the physical machine they are hosted, or running on, actually has. | [optional] [default to 2]
**Agents** | **[]string** | a list of existing remote agents managing the Kaktus node. | 

## Methods

### NewKaktus

`func NewKaktus(name string, agents []string, ) *Kaktus`

NewKaktus instantiates a new Kaktus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKaktusWithDefaults

`func NewKaktusWithDefaults() *Kaktus`

NewKaktusWithDefaults instantiates a new Kaktus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Kaktus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Kaktus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Kaktus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Kaktus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Kaktus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Kaktus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Kaktus) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Kaktus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Kaktus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Kaktus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Kaktus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCpuCost

`func (o *Kaktus) GetCpuCost() Cost`

GetCpuCost returns the CpuCost field if non-nil, zero value otherwise.

### GetCpuCostOk

`func (o *Kaktus) GetCpuCostOk() (*Cost, bool)`

GetCpuCostOk returns a tuple with the CpuCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCost

`func (o *Kaktus) SetCpuCost(v Cost)`

SetCpuCost sets CpuCost field to given value.

### HasCpuCost

`func (o *Kaktus) HasCpuCost() bool`

HasCpuCost returns a boolean if a field has been set.

### GetMemoryCost

`func (o *Kaktus) GetMemoryCost() Cost`

GetMemoryCost returns the MemoryCost field if non-nil, zero value otherwise.

### GetMemoryCostOk

`func (o *Kaktus) GetMemoryCostOk() (*Cost, bool)`

GetMemoryCostOk returns a tuple with the MemoryCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCost

`func (o *Kaktus) SetMemoryCost(v Cost)`

SetMemoryCost sets MemoryCost field to given value.

### HasMemoryCost

`func (o *Kaktus) HasMemoryCost() bool`

HasMemoryCost returns a boolean if a field has been set.

### GetOvercommitCpuRatio

`func (o *Kaktus) GetOvercommitCpuRatio() int64`

GetOvercommitCpuRatio returns the OvercommitCpuRatio field if non-nil, zero value otherwise.

### GetOvercommitCpuRatioOk

`func (o *Kaktus) GetOvercommitCpuRatioOk() (*int64, bool)`

GetOvercommitCpuRatioOk returns a tuple with the OvercommitCpuRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvercommitCpuRatio

`func (o *Kaktus) SetOvercommitCpuRatio(v int64)`

SetOvercommitCpuRatio sets OvercommitCpuRatio field to given value.

### HasOvercommitCpuRatio

`func (o *Kaktus) HasOvercommitCpuRatio() bool`

HasOvercommitCpuRatio returns a boolean if a field has been set.

### GetOvercommitMemoryRatio

`func (o *Kaktus) GetOvercommitMemoryRatio() int64`

GetOvercommitMemoryRatio returns the OvercommitMemoryRatio field if non-nil, zero value otherwise.

### GetOvercommitMemoryRatioOk

`func (o *Kaktus) GetOvercommitMemoryRatioOk() (*int64, bool)`

GetOvercommitMemoryRatioOk returns a tuple with the OvercommitMemoryRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvercommitMemoryRatio

`func (o *Kaktus) SetOvercommitMemoryRatio(v int64)`

SetOvercommitMemoryRatio sets OvercommitMemoryRatio field to given value.

### HasOvercommitMemoryRatio

`func (o *Kaktus) HasOvercommitMemoryRatio() bool`

HasOvercommitMemoryRatio returns a boolean if a field has been set.

### GetAgents

`func (o *Kaktus) GetAgents() []string`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *Kaktus) GetAgentsOk() (*[]string, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *Kaktus) SetAgents(v []string)`

SetAgents sets Agents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


