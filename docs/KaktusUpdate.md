# KaktusUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Kaktus computing node name. | [optional] 
**Description** | Pointer to **string** | The Kaktus computing node description. | [optional] 
**CpuCost** | Pointer to [**Cost**](Cost.md) | Cost associated to the Kaktus node&#39;s CPU resources. | [optional] 
**MemoryCost** | Pointer to [**Cost**](Cost.md) | Cost associated to the Kaktus node&#39;s memory resources. | [optional] 
**OvercommitCpuRatio** | Pointer to **int64** | The Kaktus node CPU resource over-commit ratio. Overcommitting CPU resources for VMs means allocating more virtual CPUs (vCPUs) to the virtual machines (VMs) than the physical cores available on the node. This can help optimize the utilization of the node CPU and increase the density of VMs per node. | [optional] 
**OvercommitMemoryRatio** | Pointer to **int64** | The Kaktus node memory resource over-commit ratio. Memory overcommitment is a concept in computing that covers the assignment of more memory to virtual computing devices (or processes) than the physical machine they are hosted, or running on, actually has. | [optional] 
**Agents** | Pointer to **[]string** | a list of existing remote agents managing the Kaktus node. | [optional] 

## Methods

### NewKaktusUpdate

`func NewKaktusUpdate() *KaktusUpdate`

NewKaktusUpdate instantiates a new KaktusUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKaktusUpdateWithDefaults

`func NewKaktusUpdateWithDefaults() *KaktusUpdate`

NewKaktusUpdateWithDefaults instantiates a new KaktusUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KaktusUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KaktusUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KaktusUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KaktusUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *KaktusUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KaktusUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KaktusUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KaktusUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCpuCost

`func (o *KaktusUpdate) GetCpuCost() Cost`

GetCpuCost returns the CpuCost field if non-nil, zero value otherwise.

### GetCpuCostOk

`func (o *KaktusUpdate) GetCpuCostOk() (*Cost, bool)`

GetCpuCostOk returns a tuple with the CpuCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCost

`func (o *KaktusUpdate) SetCpuCost(v Cost)`

SetCpuCost sets CpuCost field to given value.

### HasCpuCost

`func (o *KaktusUpdate) HasCpuCost() bool`

HasCpuCost returns a boolean if a field has been set.

### GetMemoryCost

`func (o *KaktusUpdate) GetMemoryCost() Cost`

GetMemoryCost returns the MemoryCost field if non-nil, zero value otherwise.

### GetMemoryCostOk

`func (o *KaktusUpdate) GetMemoryCostOk() (*Cost, bool)`

GetMemoryCostOk returns a tuple with the MemoryCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCost

`func (o *KaktusUpdate) SetMemoryCost(v Cost)`

SetMemoryCost sets MemoryCost field to given value.

### HasMemoryCost

`func (o *KaktusUpdate) HasMemoryCost() bool`

HasMemoryCost returns a boolean if a field has been set.

### GetOvercommitCpuRatio

`func (o *KaktusUpdate) GetOvercommitCpuRatio() int64`

GetOvercommitCpuRatio returns the OvercommitCpuRatio field if non-nil, zero value otherwise.

### GetOvercommitCpuRatioOk

`func (o *KaktusUpdate) GetOvercommitCpuRatioOk() (*int64, bool)`

GetOvercommitCpuRatioOk returns a tuple with the OvercommitCpuRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvercommitCpuRatio

`func (o *KaktusUpdate) SetOvercommitCpuRatio(v int64)`

SetOvercommitCpuRatio sets OvercommitCpuRatio field to given value.

### HasOvercommitCpuRatio

`func (o *KaktusUpdate) HasOvercommitCpuRatio() bool`

HasOvercommitCpuRatio returns a boolean if a field has been set.

### GetOvercommitMemoryRatio

`func (o *KaktusUpdate) GetOvercommitMemoryRatio() int64`

GetOvercommitMemoryRatio returns the OvercommitMemoryRatio field if non-nil, zero value otherwise.

### GetOvercommitMemoryRatioOk

`func (o *KaktusUpdate) GetOvercommitMemoryRatioOk() (*int64, bool)`

GetOvercommitMemoryRatioOk returns a tuple with the OvercommitMemoryRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvercommitMemoryRatio

`func (o *KaktusUpdate) SetOvercommitMemoryRatio(v int64)`

SetOvercommitMemoryRatio sets OvercommitMemoryRatio field to given value.

### HasOvercommitMemoryRatio

`func (o *KaktusUpdate) HasOvercommitMemoryRatio() bool`

HasOvercommitMemoryRatio returns a boolean if a field has been set.

### GetAgents

`func (o *KaktusUpdate) GetAgents() []string`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *KaktusUpdate) GetAgentsOk() (*[]string, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *KaktusUpdate) SetAgents(v []string)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *KaktusUpdate) HasAgents() bool`

HasAgents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


