# Kwarantine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Kwarantine ID (auto-generated) (read-only). | [optional] [readonly] 
**Name** | **string** | The Kwarantine name. | 
**Description** | Pointer to **string** | The Kwarantine description. | [optional] 
**Policy** | Pointer to **string** | The anti-affinity enforcement policy. Use \&quot;host\&quot; to guarantee no two group members are ever scheduled on the same Kaktus computing node, or \&quot;zone\&quot; to guarantee no two group members are ever scheduled within the same availability zone. | [optional] [default to "host"]
**Instances** | Pointer to **[]string** | A list of existing virtual machine instance IDs that are part of the Kwarantine group. Members are guaranteed to never be co-scheduled according to the group policy, regardless of Kaktus host scheduling score. Managed through the dedicated add/remove operations (read-only). | [optional] [readonly] 
**Komputes** | Pointer to **[]string** | A list of existing Kompute IDs that are part of the Kwarantine group. Members are guaranteed to never be co-scheduled according to the group policy, regardless of Kaktus host scheduling score. Managed through the dedicated add/remove operations (read-only). | [optional] [readonly] 

## Methods

### NewKwarantine

`func NewKwarantine(name string, ) *Kwarantine`

NewKwarantine instantiates a new Kwarantine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKwarantineWithDefaults

`func NewKwarantineWithDefaults() *Kwarantine`

NewKwarantineWithDefaults instantiates a new Kwarantine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Kwarantine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Kwarantine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Kwarantine) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Kwarantine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Kwarantine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Kwarantine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Kwarantine) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Kwarantine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Kwarantine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Kwarantine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Kwarantine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicy

`func (o *Kwarantine) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *Kwarantine) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *Kwarantine) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *Kwarantine) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetInstances

`func (o *Kwarantine) GetInstances() []string`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *Kwarantine) GetInstancesOk() (*[]string, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *Kwarantine) SetInstances(v []string)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *Kwarantine) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetKomputes

`func (o *Kwarantine) GetKomputes() []string`

GetKomputes returns the Komputes field if non-nil, zero value otherwise.

### GetKomputesOk

`func (o *Kwarantine) GetKomputesOk() (*[]string, bool)`

GetKomputesOk returns a tuple with the Komputes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKomputes

`func (o *Kwarantine) SetKomputes(v []string)`

SetKomputes sets Komputes field to given value.

### HasKomputes

`func (o *Kwarantine) HasKomputes() bool`

HasKomputes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


