# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The project ID (auto-generated). | [optional] 
**Name** | **string** | The project name. | 
**Description** | Pointer to **string** | The project description. | [optional] 
**Domain** | Pointer to **string** | Internal domain name (e.g. myproject.acme.com). | [optional] 
**RootPassword** | Pointer to **string** | Default root password, set at cloud-init instance bootstrap phase. Will be randomly auto-generated at each instance creation if unspecified. | [optional] 
**BootstrapUser** | Pointer to **string** | Default service user name, created at cloud-init instance bootstrap phase. Will use Kowabunga&#39;s default configuration one if unspecified. | [optional] 
**BootstrapPubkey** | Pointer to **string** | Default public SSH key, to be associated to bootstrap user. Will use Kowabunga&#39;s default configuration one if unspecified. | [optional] 
**Tags** | Pointer to **[]string** | A list of tags to be associated to the project. | [optional] 
**Metadatas** | Pointer to [**[]Metadata**](Metadata.md) | A list of metadata to be associated to the project. | [optional] 
**Quotas** | Pointer to [**ProjectResources**](ProjectResources.md) |  | [optional] 
**PrivateSubnets** | Pointer to [**[]RegionSubnet**](RegionSubnet.md) | The assigned project VPC private subnets IDs (read-only). | [optional] 
**ReservedVrrpIds** | Pointer to **[]int32** | The list of VRRP IDs used by -as-a-service resources within the project virtual network (read-only). Should your application use VRRP for service redundancy, you should use different IDs to prevent issues.. | [optional] 
**Teams** | **[]string** | A list of user teams allowed to administrate the project (i.e. capable of managing internal resources). | 
**Regions** | **[]string** | A list of Kowabunga regions the project is managing resources from. | 

## Methods

### NewProject

`func NewProject(name string, teams []string, regions []string, ) *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Project) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Project) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Project) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Project) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Project) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Project) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *Project) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Project) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Project) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Project) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetRootPassword

`func (o *Project) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *Project) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *Project) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *Project) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetBootstrapUser

`func (o *Project) GetBootstrapUser() string`

GetBootstrapUser returns the BootstrapUser field if non-nil, zero value otherwise.

### GetBootstrapUserOk

`func (o *Project) GetBootstrapUserOk() (*string, bool)`

GetBootstrapUserOk returns a tuple with the BootstrapUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapUser

`func (o *Project) SetBootstrapUser(v string)`

SetBootstrapUser sets BootstrapUser field to given value.

### HasBootstrapUser

`func (o *Project) HasBootstrapUser() bool`

HasBootstrapUser returns a boolean if a field has been set.

### GetBootstrapPubkey

`func (o *Project) GetBootstrapPubkey() string`

GetBootstrapPubkey returns the BootstrapPubkey field if non-nil, zero value otherwise.

### GetBootstrapPubkeyOk

`func (o *Project) GetBootstrapPubkeyOk() (*string, bool)`

GetBootstrapPubkeyOk returns a tuple with the BootstrapPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapPubkey

`func (o *Project) SetBootstrapPubkey(v string)`

SetBootstrapPubkey sets BootstrapPubkey field to given value.

### HasBootstrapPubkey

`func (o *Project) HasBootstrapPubkey() bool`

HasBootstrapPubkey returns a boolean if a field has been set.

### GetTags

`func (o *Project) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Project) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Project) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Project) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadatas

`func (o *Project) GetMetadatas() []Metadata`

GetMetadatas returns the Metadatas field if non-nil, zero value otherwise.

### GetMetadatasOk

`func (o *Project) GetMetadatasOk() (*[]Metadata, bool)`

GetMetadatasOk returns a tuple with the Metadatas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadatas

`func (o *Project) SetMetadatas(v []Metadata)`

SetMetadatas sets Metadatas field to given value.

### HasMetadatas

`func (o *Project) HasMetadatas() bool`

HasMetadatas returns a boolean if a field has been set.

### GetQuotas

`func (o *Project) GetQuotas() ProjectResources`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *Project) GetQuotasOk() (*ProjectResources, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *Project) SetQuotas(v ProjectResources)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *Project) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *Project) GetPrivateSubnets() []RegionSubnet`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *Project) GetPrivateSubnetsOk() (*[]RegionSubnet, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *Project) SetPrivateSubnets(v []RegionSubnet)`

SetPrivateSubnets sets PrivateSubnets field to given value.

### HasPrivateSubnets

`func (o *Project) HasPrivateSubnets() bool`

HasPrivateSubnets returns a boolean if a field has been set.

### GetReservedVrrpIds

`func (o *Project) GetReservedVrrpIds() []int32`

GetReservedVrrpIds returns the ReservedVrrpIds field if non-nil, zero value otherwise.

### GetReservedVrrpIdsOk

`func (o *Project) GetReservedVrrpIdsOk() (*[]int32, bool)`

GetReservedVrrpIdsOk returns a tuple with the ReservedVrrpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedVrrpIds

`func (o *Project) SetReservedVrrpIds(v []int32)`

SetReservedVrrpIds sets ReservedVrrpIds field to given value.

### HasReservedVrrpIds

`func (o *Project) HasReservedVrrpIds() bool`

HasReservedVrrpIds returns a boolean if a field has been set.

### GetTeams

`func (o *Project) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *Project) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *Project) SetTeams(v []string)`

SetTeams sets Teams field to given value.


### GetRegions

`func (o *Project) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Project) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Project) SetRegions(v []string)`

SetRegions sets Regions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


