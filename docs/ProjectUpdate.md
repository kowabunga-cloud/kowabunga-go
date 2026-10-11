# ProjectUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The project name. | [optional] 
**Description** | Pointer to **string** | The project description. | [optional] 
**RootPassword** | Pointer to **string** | Default root password, set at cloud-init instance bootstrap phase. Will be randomly auto-generated at each instance creation if unspecified. | [optional] 
**BootstrapUser** | Pointer to **string** | Default service user name, created at cloud-init instance bootstrap phase. Will use Kowabunga&#39;s default configuration one if unspecified. | [optional] 
**BootstrapPubkey** | Pointer to **string** | Default public SSH key, to be associated to bootstrap user. Will use Kowabunga&#39;s default configuration one if unspecified. | [optional] 
**Tags** | Pointer to **[]string** | A list of tags to be associated to the project. | [optional] 
**Metadatas** | Pointer to [**[]Metadata**](Metadata.md) | A list of metadata to be associated to the project. | [optional] 
**Quotas** | Pointer to [**ProjectResources**](ProjectResources.md) | The global project resource quotas (0 for unlimited). | [optional] 
**Teams** | Pointer to **[]string** | A list of user teams allowed to administrate the project (i.e. capable of managing internal resources). | [optional] 
**Regions** | Pointer to **[]string** | A list of Kowabunga regions the project is managing resources from. | [optional] 

## Methods

### NewProjectUpdate

`func NewProjectUpdate() *ProjectUpdate`

NewProjectUpdate instantiates a new ProjectUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectUpdateWithDefaults

`func NewProjectUpdateWithDefaults() *ProjectUpdate`

NewProjectUpdateWithDefaults instantiates a new ProjectUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRootPassword

`func (o *ProjectUpdate) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *ProjectUpdate) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *ProjectUpdate) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *ProjectUpdate) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetBootstrapUser

`func (o *ProjectUpdate) GetBootstrapUser() string`

GetBootstrapUser returns the BootstrapUser field if non-nil, zero value otherwise.

### GetBootstrapUserOk

`func (o *ProjectUpdate) GetBootstrapUserOk() (*string, bool)`

GetBootstrapUserOk returns a tuple with the BootstrapUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapUser

`func (o *ProjectUpdate) SetBootstrapUser(v string)`

SetBootstrapUser sets BootstrapUser field to given value.

### HasBootstrapUser

`func (o *ProjectUpdate) HasBootstrapUser() bool`

HasBootstrapUser returns a boolean if a field has been set.

### GetBootstrapPubkey

`func (o *ProjectUpdate) GetBootstrapPubkey() string`

GetBootstrapPubkey returns the BootstrapPubkey field if non-nil, zero value otherwise.

### GetBootstrapPubkeyOk

`func (o *ProjectUpdate) GetBootstrapPubkeyOk() (*string, bool)`

GetBootstrapPubkeyOk returns a tuple with the BootstrapPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapPubkey

`func (o *ProjectUpdate) SetBootstrapPubkey(v string)`

SetBootstrapPubkey sets BootstrapPubkey field to given value.

### HasBootstrapPubkey

`func (o *ProjectUpdate) HasBootstrapPubkey() bool`

HasBootstrapPubkey returns a boolean if a field has been set.

### GetTags

`func (o *ProjectUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProjectUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadatas

`func (o *ProjectUpdate) GetMetadatas() []Metadata`

GetMetadatas returns the Metadatas field if non-nil, zero value otherwise.

### GetMetadatasOk

`func (o *ProjectUpdate) GetMetadatasOk() (*[]Metadata, bool)`

GetMetadatasOk returns a tuple with the Metadatas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadatas

`func (o *ProjectUpdate) SetMetadatas(v []Metadata)`

SetMetadatas sets Metadatas field to given value.

### HasMetadatas

`func (o *ProjectUpdate) HasMetadatas() bool`

HasMetadatas returns a boolean if a field has been set.

### GetQuotas

`func (o *ProjectUpdate) GetQuotas() ProjectResources`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *ProjectUpdate) GetQuotasOk() (*ProjectResources, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *ProjectUpdate) SetQuotas(v ProjectResources)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *ProjectUpdate) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.

### GetTeams

`func (o *ProjectUpdate) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ProjectUpdate) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ProjectUpdate) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *ProjectUpdate) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetRegions

`func (o *ProjectUpdate) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ProjectUpdate) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ProjectUpdate) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *ProjectUpdate) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


