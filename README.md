<p align="center">
  <a href="https://www.kowabunga.cloud/?utm_source=github&utm_medium=logo" target="_blank">
    <picture>
      <source srcset="https://rawcdn.githack.com/kowabunga-cloud/infographics/master/art/kowabunga-square-600x600-2.png" media="(prefers-color-scheme: dark)" />
      <source srcset="https://rawcdn.githack.com/kowabunga-cloud/infographics/master/art/kowabunga-square-600x600-2.png" media="(prefers-color-scheme: light), (prefers-color-scheme: no-preference)" />
      <img src="https://rawcdn.githack.com/kowabunga-cloud/infographics/master/art/kowabunga-square-600x600-2.png" alt="Kowabunga" width="200">
    </picture>
  </a>
</p>

# Official Kowabunga SDK for Go

## Installation

`kowabunga-go` can be installed like any other Go library through `go get`:

```console
$ go get github.com/kowabunga-cloud/kowabunga-go@latest
```

Check out the [list of released versions](https://github.com/kowabunga-cloud/kowabunga-go/releases).

## Configuration

To use `kowabunga-go`, you’ll need to import the `kowabunga-go` package:

```go
import kowabunga "github.com/kowabunga-cloud/kowabunga-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Usage

Creating an API client can be done by:

```go

var client *kowabunga.APIClient

u, _ := url.Parse(uri)

cfg := kowabunga.NewConfiguration()
cfg.Host = u.Host
cfg.Scheme = u.Scheme
cfg.Debug = true
cfg.AddDefaultHeader("X-API-Key", token)

client = kowabunga.NewAPIClient(cfg), nil
```

where **uri** is *https://your_kowabunga_kahuna_server* and **token** is the associated API key.

## Documentation for API Endpoints

All URIs are relative to *https://your_kowabunga_kahuna_server/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdapterAPI* | [**DeleteAdapter**](docs/AdapterAPI.md#deleteadapter) | **Delete** /adapter/{adapterId} | 
*AdapterAPI* | [**ListAdapters**](docs/AdapterAPI.md#listadapters) | **Get** /adapter | 
*AdapterAPI* | [**ReadAdapter**](docs/AdapterAPI.md#readadapter) | **Get** /adapter/{adapterId} | 
*AdapterAPI* | [**UpdateAdapter**](docs/AdapterAPI.md#updateadapter) | **Put** /adapter/{adapterId} | 
*AgentAPI* | [**CreateAgent**](docs/AgentAPI.md#createagent) | **Post** /agent | 
*AgentAPI* | [**DeleteAgent**](docs/AgentAPI.md#deleteagent) | **Delete** /agent/{agentId} | 
*AgentAPI* | [**ListAgents**](docs/AgentAPI.md#listagents) | **Get** /agent | 
*AgentAPI* | [**ReadAgent**](docs/AgentAPI.md#readagent) | **Get** /agent/{agentId} | 
*AgentAPI* | [**SetAgentApiToken**](docs/AgentAPI.md#setagentapitoken) | **Patch** /agent/{agentId}/token | 
*AgentAPI* | [**UpdateAgent**](docs/AgentAPI.md#updateagent) | **Put** /agent/{agentId} | 
*InstanceAPI* | [**DeleteInstance**](docs/InstanceAPI.md#deleteinstance) | **Delete** /instance/{instanceId} | 
*InstanceAPI* | [**ListInstances**](docs/InstanceAPI.md#listinstances) | **Get** /instance | 
*InstanceAPI* | [**ReadInstance**](docs/InstanceAPI.md#readinstance) | **Get** /instance/{instanceId} | 
*InstanceAPI* | [**ReadInstanceRemoteConnection**](docs/InstanceAPI.md#readinstanceremoteconnection) | **Get** /instance/{instanceId}/connect | 
*InstanceAPI* | [**ReadInstanceState**](docs/InstanceAPI.md#readinstancestate) | **Get** /instance/{instanceId}/state | 
*InstanceAPI* | [**RebootInstance**](docs/InstanceAPI.md#rebootinstance) | **Patch** /instance/{instanceId}/reboot | 
*InstanceAPI* | [**ResetInstance**](docs/InstanceAPI.md#resetinstance) | **Patch** /instance/{instanceId}/reset | 
*InstanceAPI* | [**ResumeInstance**](docs/InstanceAPI.md#resumeinstance) | **Patch** /instance/{instanceId}/resume | 
*InstanceAPI* | [**ShutdownInstance**](docs/InstanceAPI.md#shutdowninstance) | **Patch** /instance/{instanceId}/shutdown | 
*InstanceAPI* | [**StartInstance**](docs/InstanceAPI.md#startinstance) | **Patch** /instance/{instanceId}/start | 
*InstanceAPI* | [**StopInstance**](docs/InstanceAPI.md#stopinstance) | **Patch** /instance/{instanceId}/stop | 
*InstanceAPI* | [**SuspendInstance**](docs/InstanceAPI.md#suspendinstance) | **Patch** /instance/{instanceId}/suspend | 
*InstanceAPI* | [**UpdateInstance**](docs/InstanceAPI.md#updateinstance) | **Put** /instance/{instanceId} | 
*KaktusAPI* | [**DeleteKaktus**](docs/KaktusAPI.md#deletekaktus) | **Delete** /kaktus/{kaktusId} | 
*KaktusAPI* | [**ListKaktusInstances**](docs/KaktusAPI.md#listkaktusinstances) | **Get** /kaktus/{kaktusId}/instances | 
*KaktusAPI* | [**ListKaktuss**](docs/KaktusAPI.md#listkaktuss) | **Get** /kaktus | 
*KaktusAPI* | [**ReadKaktus**](docs/KaktusAPI.md#readkaktus) | **Get** /kaktus/{kaktusId} | 
*KaktusAPI* | [**ReadKaktusCaps**](docs/KaktusAPI.md#readkaktuscaps) | **Get** /kaktus/{kaktusId}/caps | 
*KaktusAPI* | [**UpdateKaktus**](docs/KaktusAPI.md#updatekaktus) | **Put** /kaktus/{kaktusId} | 
*KawaiiAPI* | [**CreateKawaiiIpSec**](docs/KawaiiAPI.md#createkawaiiipsec) | **Post** /kawaii/{kawaiiId}/ipsec | 
*KawaiiAPI* | [**DeleteKawaii**](docs/KawaiiAPI.md#deletekawaii) | **Delete** /kawaii/{kawaiiId} | 
*KawaiiAPI* | [**DeleteKawaiiIpSec**](docs/KawaiiAPI.md#deletekawaiiipsec) | **Delete** /kawaii/{kawaiiId}/ipsec/{KawaiiIpSecId} | 
*KawaiiAPI* | [**ListKawaiiIpSecs**](docs/KawaiiAPI.md#listkawaiiipsecs) | **Get** /kawaii/{kawaiiId}/ipsec | 
*KawaiiAPI* | [**ListKawaiis**](docs/KawaiiAPI.md#listkawaiis) | **Get** /kawaii | 
*KawaiiAPI* | [**ReadKawaii**](docs/KawaiiAPI.md#readkawaii) | **Get** /kawaii/{kawaiiId} | 
*KawaiiAPI* | [**ReadKawaiiIpSec**](docs/KawaiiAPI.md#readkawaiiipsec) | **Get** /kawaii/{kawaiiId}/ipsec/{KawaiiIpSecId} | 
*KawaiiAPI* | [**UpdateKawaii**](docs/KawaiiAPI.md#updatekawaii) | **Put** /kawaii/{kawaiiId} | 
*KawaiiAPI* | [**UpdateKawaiiIpSec**](docs/KawaiiAPI.md#updatekawaiiipsec) | **Put** /kawaii/{kawaiiId}/ipsec/{KawaiiIpSecId} | 
*KiwiAPI* | [**DeleteKiwi**](docs/KiwiAPI.md#deletekiwi) | **Delete** /kiwi/{kiwiId} | 
*KiwiAPI* | [**ListKiwis**](docs/KiwiAPI.md#listkiwis) | **Get** /kiwi | 
*KiwiAPI* | [**ReadKiwi**](docs/KiwiAPI.md#readkiwi) | **Get** /kiwi/{kiwiId} | 
*KiwiAPI* | [**UpdateKiwi**](docs/KiwiAPI.md#updatekiwi) | **Put** /kiwi/{kiwiId} | 
*KomputeAPI* | [**DeleteKompute**](docs/KomputeAPI.md#deletekompute) | **Delete** /kompute/{komputeId} | 
*KomputeAPI* | [**ListKomputes**](docs/KomputeAPI.md#listkomputes) | **Get** /kompute | 
*KomputeAPI* | [**ReadKompute**](docs/KomputeAPI.md#readkompute) | **Get** /kompute/{komputeId} | 
*KomputeAPI* | [**ReadKomputeState**](docs/KomputeAPI.md#readkomputestate) | **Get** /kompute/{komputeId}/state | 
*KomputeAPI* | [**RebootKompute**](docs/KomputeAPI.md#rebootkompute) | **Patch** /kompute/{komputeId}/reboot | 
*KomputeAPI* | [**ResetKompute**](docs/KomputeAPI.md#resetkompute) | **Patch** /kompute/{komputeId}/reset | 
*KomputeAPI* | [**ResumeKompute**](docs/KomputeAPI.md#resumekompute) | **Patch** /kompute/{komputeId}/resume | 
*KomputeAPI* | [**ShutdownKompute**](docs/KomputeAPI.md#shutdownkompute) | **Patch** /kompute/{komputeId}/shutdown | 
*KomputeAPI* | [**StartKompute**](docs/KomputeAPI.md#startkompute) | **Patch** /kompute/{komputeId}/start | 
*KomputeAPI* | [**StopKompute**](docs/KomputeAPI.md#stopkompute) | **Patch** /kompute/{komputeId}/stop | 
*KomputeAPI* | [**SuspendKompute**](docs/KomputeAPI.md#suspendkompute) | **Patch** /kompute/{komputeId}/suspend | 
*KomputeAPI* | [**UpdateKompute**](docs/KomputeAPI.md#updatekompute) | **Put** /kompute/{komputeId} | 
*KonveyAPI* | [**DeleteKonvey**](docs/KonveyAPI.md#deletekonvey) | **Delete** /konvey/{konveyId} | 
*KonveyAPI* | [**ListKonveys**](docs/KonveyAPI.md#listkonveys) | **Get** /konvey | 
*KonveyAPI* | [**ReadKonvey**](docs/KonveyAPI.md#readkonvey) | **Get** /konvey/{konveyId} | 
*KonveyAPI* | [**UpdateKonvey**](docs/KonveyAPI.md#updatekonvey) | **Put** /konvey/{konveyId} | 
*KyloAPI* | [**DeleteKylo**](docs/KyloAPI.md#deletekylo) | **Delete** /kylo/{kyloId} | 
*KyloAPI* | [**ListKylos**](docs/KyloAPI.md#listkylos) | **Get** /kylo | 
*KyloAPI* | [**ReadKylo**](docs/KyloAPI.md#readkylo) | **Get** /kylo/{kyloId} | 
*KyloAPI* | [**UpdateKylo**](docs/KyloAPI.md#updatekylo) | **Put** /kylo/{kyloId} | 
*NfsAPI* | [**DeleteStorageNFS**](docs/NfsAPI.md#deletestoragenfs) | **Delete** /nfs/{nfsId} | 
*NfsAPI* | [**ListStorageNFSKylos**](docs/NfsAPI.md#liststoragenfskylos) | **Get** /nfs/{nfsId}/kylo | 
*NfsAPI* | [**ListStorageNFSs**](docs/NfsAPI.md#liststoragenfss) | **Get** /nfs | 
*NfsAPI* | [**ReadStorageNFS**](docs/NfsAPI.md#readstoragenfs) | **Get** /nfs/{nfsId} | 
*NfsAPI* | [**UpdateStorageNFS**](docs/NfsAPI.md#updatestoragenfs) | **Put** /nfs/{nfsId} | 
*PoolAPI* | [**CreateTemplate**](docs/PoolAPI.md#createtemplate) | **Post** /pool/{poolId}/template | 
*PoolAPI* | [**DeleteStoragePool**](docs/PoolAPI.md#deletestoragepool) | **Delete** /pool/{poolId} | 
*PoolAPI* | [**ListStoragePoolTemplates**](docs/PoolAPI.md#liststoragepooltemplates) | **Get** /pool/{poolId}/templates | 
*PoolAPI* | [**ListStoragePoolVolumes**](docs/PoolAPI.md#liststoragepoolvolumes) | **Get** /pool/{poolId}/volumes | 
*PoolAPI* | [**ListStoragePools**](docs/PoolAPI.md#liststoragepools) | **Get** /pool | 
*PoolAPI* | [**ReadStoragePool**](docs/PoolAPI.md#readstoragepool) | **Get** /pool/{poolId} | 
*PoolAPI* | [**SetStoragePoolDefaultTemplate**](docs/PoolAPI.md#setstoragepooldefaulttemplate) | **Patch** /pool/{poolId}/template/{templateId}/default | 
*PoolAPI* | [**UpdateStoragePool**](docs/PoolAPI.md#updatestoragepool) | **Put** /pool/{poolId} | 
*ProjectAPI* | [**CreateProject**](docs/ProjectAPI.md#createproject) | **Post** /project | 
*ProjectAPI* | [**CreateProjectDnsRecord**](docs/ProjectAPI.md#createprojectdnsrecord) | **Post** /project/{projectId}/record | 
*ProjectAPI* | [**CreateProjectRegionKawaii**](docs/ProjectAPI.md#createprojectregionkawaii) | **Post** /project/{projectId}/region/{regionId}/kawaii | 
*ProjectAPI* | [**CreateProjectRegionKonvey**](docs/ProjectAPI.md#createprojectregionkonvey) | **Post** /project/{projectId}/region/{regionId}/konvey | 
*ProjectAPI* | [**CreateProjectRegionKylo**](docs/ProjectAPI.md#createprojectregionkylo) | **Post** /project/{projectId}/region/{regionId}/kylo | 
*ProjectAPI* | [**CreateProjectRegionVolume**](docs/ProjectAPI.md#createprojectregionvolume) | **Post** /project/{projectId}/region/{regionId}/volume | 
*ProjectAPI* | [**CreateProjectZoneInstance**](docs/ProjectAPI.md#createprojectzoneinstance) | **Post** /project/{projectId}/zone/{zoneId}/instance | 
*ProjectAPI* | [**CreateProjectZoneKompute**](docs/ProjectAPI.md#createprojectzonekompute) | **Post** /project/{projectId}/zone/{zoneId}/kompute | 
*ProjectAPI* | [**CreateProjectZoneKonvey**](docs/ProjectAPI.md#createprojectzonekonvey) | **Post** /project/{projectId}/zone/{zoneId}/konvey | 
*ProjectAPI* | [**DeleteProject**](docs/ProjectAPI.md#deleteproject) | **Delete** /project/{projectId} | 
*ProjectAPI* | [**ListProjectDnsRecords**](docs/ProjectAPI.md#listprojectdnsrecords) | **Get** /project/{projectId}/records | 
*ProjectAPI* | [**ListProjectRegionKawaiis**](docs/ProjectAPI.md#listprojectregionkawaiis) | **Get** /project/{projectId}/region/{regionId}/kawaiis | 
*ProjectAPI* | [**ListProjectRegionKonveys**](docs/ProjectAPI.md#listprojectregionkonveys) | **Get** /project/{projectId}/region/{regionId}/konveys | 
*ProjectAPI* | [**ListProjectRegionKylos**](docs/ProjectAPI.md#listprojectregionkylos) | **Get** /project/{projectId}/region/{regionId}/kylo | 
*ProjectAPI* | [**ListProjectRegionVolumes**](docs/ProjectAPI.md#listprojectregionvolumes) | **Get** /project/{projectId}/region/{regionId}/volumes | 
*ProjectAPI* | [**ListProjectZoneInstances**](docs/ProjectAPI.md#listprojectzoneinstances) | **Get** /project/{projectId}/zone/{zoneId}/instances | 
*ProjectAPI* | [**ListProjectZoneKomputes**](docs/ProjectAPI.md#listprojectzonekomputes) | **Get** /project/{projectId}/zone/{zoneId}/komputes | 
*ProjectAPI* | [**ListProjectZoneKonveys**](docs/ProjectAPI.md#listprojectzonekonveys) | **Get** /project/{projectId}/zone/{zoneId}/konveys | 
*ProjectAPI* | [**ListProjects**](docs/ProjectAPI.md#listprojects) | **Get** /project | 
*ProjectAPI* | [**ReadProject**](docs/ProjectAPI.md#readproject) | **Get** /project/{projectId} | 
*ProjectAPI* | [**ReadProjectCost**](docs/ProjectAPI.md#readprojectcost) | **Get** /project/{projectId}/cost | 
*ProjectAPI* | [**ReadProjectUsage**](docs/ProjectAPI.md#readprojectusage) | **Get** /project/{projectId}/usage | 
*ProjectAPI* | [**UpdateProject**](docs/ProjectAPI.md#updateproject) | **Put** /project/{projectId} | 
*RecordAPI* | [**DeleteDnsRecord**](docs/RecordAPI.md#deletednsrecord) | **Delete** /record/{recordId} | 
*RecordAPI* | [**ReadDnsRecord**](docs/RecordAPI.md#readdnsrecord) | **Get** /record/{recordId} | 
*RecordAPI* | [**UpdateDnsRecord**](docs/RecordAPI.md#updatednsrecord) | **Put** /record/{recordId} | 
*RegionAPI* | [**CreateKiwi**](docs/RegionAPI.md#createkiwi) | **Post** /region/{regionId}/kiwi | 
*RegionAPI* | [**CreateRegion**](docs/RegionAPI.md#createregion) | **Post** /region | 
*RegionAPI* | [**CreateStorageNFS**](docs/RegionAPI.md#createstoragenfs) | **Post** /region/{regionId}/nfs | 
*RegionAPI* | [**CreateStoragePool**](docs/RegionAPI.md#createstoragepool) | **Post** /region/{regionId}/pool | 
*RegionAPI* | [**CreateVNet**](docs/RegionAPI.md#createvnet) | **Post** /region/{regionId}/vnet | 
*RegionAPI* | [**CreateZone**](docs/RegionAPI.md#createzone) | **Post** /region/{regionId}/zone | 
*RegionAPI* | [**DeleteRegion**](docs/RegionAPI.md#deleteregion) | **Delete** /region/{regionId} | 
*RegionAPI* | [**ListRegionKiwis**](docs/RegionAPI.md#listregionkiwis) | **Get** /region/{regionId}/kiwis | 
*RegionAPI* | [**ListRegionStorageNFSs**](docs/RegionAPI.md#listregionstoragenfss) | **Get** /region/{regionId}/nfs | 
*RegionAPI* | [**ListRegionStoragePools**](docs/RegionAPI.md#listregionstoragepools) | **Get** /region/{regionId}/pools | 
*RegionAPI* | [**ListRegionVNets**](docs/RegionAPI.md#listregionvnets) | **Get** /region/{regionId}/vnets | 
*RegionAPI* | [**ListRegionZones**](docs/RegionAPI.md#listregionzones) | **Get** /region/{regionId}/zones | 
*RegionAPI* | [**ListRegions**](docs/RegionAPI.md#listregions) | **Get** /region | 
*RegionAPI* | [**ReadRegion**](docs/RegionAPI.md#readregion) | **Get** /region/{regionId} | 
*RegionAPI* | [**SetRegionDefaultStorageNFS**](docs/RegionAPI.md#setregiondefaultstoragenfs) | **Patch** /region/{regionId}/nfs/{nfsId}/default | 
*RegionAPI* | [**SetRegionDefaultStoragePool**](docs/RegionAPI.md#setregiondefaultstoragepool) | **Patch** /region/{regionId}/pool/{poolId}/default | 
*RegionAPI* | [**UpdateRegion**](docs/RegionAPI.md#updateregion) | **Put** /region/{regionId} | 
*SubnetAPI* | [**CreateAdapter**](docs/SubnetAPI.md#createadapter) | **Post** /subnet/{subnetId}/adapter | 
*SubnetAPI* | [**DeleteSubnet**](docs/SubnetAPI.md#deletesubnet) | **Delete** /subnet/{subnetId} | 
*SubnetAPI* | [**ListSubnetAdapters**](docs/SubnetAPI.md#listsubnetadapters) | **Get** /subnet/{subnetId}/adapters | 
*SubnetAPI* | [**ListSubnets**](docs/SubnetAPI.md#listsubnets) | **Get** /subnet | 
*SubnetAPI* | [**ReadSubnet**](docs/SubnetAPI.md#readsubnet) | **Get** /subnet/{subnetId} | 
*SubnetAPI* | [**UpdateSubnet**](docs/SubnetAPI.md#updatesubnet) | **Put** /subnet/{subnetId} | 
*TeamAPI* | [**CreateTeam**](docs/TeamAPI.md#createteam) | **Post** /team | 
*TeamAPI* | [**DeleteTeam**](docs/TeamAPI.md#deleteteam) | **Delete** /team/{teamId} | 
*TeamAPI* | [**ListTeams**](docs/TeamAPI.md#listteams) | **Get** /team | 
*TeamAPI* | [**ReadTeam**](docs/TeamAPI.md#readteam) | **Get** /team/{teamId} | 
*TeamAPI* | [**UpdateTeam**](docs/TeamAPI.md#updateteam) | **Put** /team/{teamId} | 
*TemplateAPI* | [**DeleteTemplate**](docs/TemplateAPI.md#deletetemplate) | **Delete** /template/{templateId} | 
*TemplateAPI* | [**ListTemplates**](docs/TemplateAPI.md#listtemplates) | **Get** /template | 
*TemplateAPI* | [**ReadTemplate**](docs/TemplateAPI.md#readtemplate) | **Get** /template/{templateId} | 
*TemplateAPI* | [**UpdateTemplate**](docs/TemplateAPI.md#updatetemplate) | **Put** /template/{templateId} | 
*TokenAPI* | [**DeleteApiToken**](docs/TokenAPI.md#deleteapitoken) | **Delete** /token/{tokenId} | 
*TokenAPI* | [**ListApiTokens**](docs/TokenAPI.md#listapitokens) | **Get** /token | 
*TokenAPI* | [**ReadApiToken**](docs/TokenAPI.md#readapitoken) | **Get** /token/{tokenId} | 
*TokenAPI* | [**UpdateApiToken**](docs/TokenAPI.md#updateapitoken) | **Put** /token/{tokenId} | 
*UserAPI* | [**CreateUser**](docs/UserAPI.md#createuser) | **Post** /user | 
*UserAPI* | [**DeleteUser**](docs/UserAPI.md#deleteuser) | **Delete** /user/{userId} | 
*UserAPI* | [**ListUsers**](docs/UserAPI.md#listusers) | **Get** /user | 
*UserAPI* | [**Login**](docs/UserAPI.md#login) | **Post** /login | 
*UserAPI* | [**Logout**](docs/UserAPI.md#logout) | **Post** /logout | 
*UserAPI* | [**ReadUser**](docs/UserAPI.md#readuser) | **Get** /user/{userId} | 
*UserAPI* | [**ResetPassword**](docs/UserAPI.md#resetpassword) | **Put** /resetPassword | 
*UserAPI* | [**ResetUserPassword**](docs/UserAPI.md#resetuserpassword) | **Patch** /user/{userId}/resetPassword | 
*UserAPI* | [**SetUserApiToken**](docs/UserAPI.md#setuserapitoken) | **Patch** /user/{userId}/token | 
*UserAPI* | [**SetUserPassword**](docs/UserAPI.md#setuserpassword) | **Put** /user/{userId}/password | 
*UserAPI* | [**UpdateUser**](docs/UserAPI.md#updateuser) | **Put** /user/{userId} | 
*VnetAPI* | [**CreateSubnet**](docs/VnetAPI.md#createsubnet) | **Post** /vnet/{vnetId}/subnet | 
*VnetAPI* | [**DeleteVNet**](docs/VnetAPI.md#deletevnet) | **Delete** /vnet/{vnetId} | 
*VnetAPI* | [**ListVNetSubnets**](docs/VnetAPI.md#listvnetsubnets) | **Get** /vnet/{vnetId}/subnets | 
*VnetAPI* | [**ListVNets**](docs/VnetAPI.md#listvnets) | **Get** /vnet | 
*VnetAPI* | [**ReadVNet**](docs/VnetAPI.md#readvnet) | **Get** /vnet/{vnetId} | 
*VnetAPI* | [**SetVNetDefaultSubnet**](docs/VnetAPI.md#setvnetdefaultsubnet) | **Patch** /vnet/{vnetId}/subnet/{subnetId}/default | 
*VnetAPI* | [**UpdateVNet**](docs/VnetAPI.md#updatevnet) | **Put** /vnet/{vnetId} | 
*VolumeAPI* | [**DeleteVolume**](docs/VolumeAPI.md#deletevolume) | **Delete** /volume/{volumeId} | 
*VolumeAPI* | [**ListVolumes**](docs/VolumeAPI.md#listvolumes) | **Get** /volume | 
*VolumeAPI* | [**ReadVolume**](docs/VolumeAPI.md#readvolume) | **Get** /volume/{volumeId} | 
*VolumeAPI* | [**UpdateVolume**](docs/VolumeAPI.md#updatevolume) | **Put** /volume/{volumeId} | 
*ZoneAPI* | [**CreateKaktus**](docs/ZoneAPI.md#createkaktus) | **Post** /zone/{zoneId}/kaktus | 
*ZoneAPI* | [**DeleteZone**](docs/ZoneAPI.md#deletezone) | **Delete** /zone/{zoneId} | 
*ZoneAPI* | [**ListZoneKaktuses**](docs/ZoneAPI.md#listzonekaktuses) | **Get** /zone/{zoneId}/kaktuses | 
*ZoneAPI* | [**ListZones**](docs/ZoneAPI.md#listzones) | **Get** /zone | 
*ZoneAPI* | [**ReadZone**](docs/ZoneAPI.md#readzone) | **Get** /zone/{zoneId} | 
*ZoneAPI* | [**UpdateZone**](docs/ZoneAPI.md#updatezone) | **Put** /zone/{zoneId} | 

## Documentation For Models

 - [Adapter](docs/Adapter.md)
 - [Agent](docs/Agent.md)
 - [ApiErrorBadRequest](docs/ApiErrorBadRequest.md)
 - [ApiErrorConflict](docs/ApiErrorConflict.md)
 - [ApiErrorForbidden](docs/ApiErrorForbidden.md)
 - [ApiErrorInsufficientResource](docs/ApiErrorInsufficientResource.md)
 - [ApiErrorNotFound](docs/ApiErrorNotFound.md)
 - [ApiErrorUnauthorized](docs/ApiErrorUnauthorized.md)
 - [ApiErrorUnprocessableEntity](docs/ApiErrorUnprocessableEntity.md)
 - [ApiToken](docs/ApiToken.md)
 - [Cost](docs/Cost.md)
 - [DnsRecord](docs/DnsRecord.md)
 - [Instance](docs/Instance.md)
 - [InstanceRemoteAccess](docs/InstanceRemoteAccess.md)
 - [InstanceState](docs/InstanceState.md)
 - [IpRange](docs/IpRange.md)
 - [Kaktus](docs/Kaktus.md)
 - [KaktusCPU](docs/KaktusCPU.md)
 - [KaktusCaps](docs/KaktusCaps.md)
 - [Kawaii](docs/Kawaii.md)
 - [KawaiiDNatRule](docs/KawaiiDNatRule.md)
 - [KawaiiFirewall](docs/KawaiiFirewall.md)
 - [KawaiiFirewallEgressRule](docs/KawaiiFirewallEgressRule.md)
 - [KawaiiFirewallIngressRule](docs/KawaiiFirewallIngressRule.md)
 - [KawaiiIpSec](docs/KawaiiIpSec.md)
 - [KawaiiNetIp](docs/KawaiiNetIp.md)
 - [KawaiiNetIpZone](docs/KawaiiNetIpZone.md)
 - [KawaiiVpcForwardRule](docs/KawaiiVpcForwardRule.md)
 - [KawaiiVpcNetIpZone](docs/KawaiiVpcNetIpZone.md)
 - [KawaiiVpcPeering](docs/KawaiiVpcPeering.md)
 - [Kiwi](docs/Kiwi.md)
 - [Kompute](docs/Kompute.md)
 - [Konvey](docs/Konvey.md)
 - [KonveyBackends](docs/KonveyBackends.md)
 - [KonveyEndpoint](docs/KonveyEndpoint.md)
 - [Kylo](docs/Kylo.md)
 - [Metadata](docs/Metadata.md)
 - [Password](docs/Password.md)
 - [Project](docs/Project.md)
 - [ProjectResources](docs/ProjectResources.md)
 - [Region](docs/Region.md)
 - [RegionSubnet](docs/RegionSubnet.md)
 - [StorageNFS](docs/StorageNFS.md)
 - [StoragePool](docs/StoragePool.md)
 - [Subnet](docs/Subnet.md)
 - [Team](docs/Team.md)
 - [Template](docs/Template.md)
 - [User](docs/User.md)
 - [UserCredentials](docs/UserCredentials.md)
 - [UserEmail](docs/UserEmail.md)
 - [VNet](docs/VNet.md)
 - [Volume](docs/Volume.md)
 - [Zone](docs/Zone.md)

## License

Licensed under [Apache License, Version 2.0](https://opensource.org/license/apache-2-0), see [`LICENSE`](LICENSE).
