# GetAssetDetailsByAssetSymbolRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** | Defines the unique ID of the specific asset. | 
**AssetLogo** | [**GetAssetDetailsByAssetIDRIAssetLogo**](GetAssetDetailsByAssetIDRIAssetLogo.md) |  | 
**AssetName** | **string** | Specifies the name of the asset in question. | 
**AssetOriginalSymbol** | **string** | Specifies the asset&#39;s original symbol as introduced by its founders. | 
**AssetSymbol** | **string** | Specifies the asset&#39;s unique symbol in the Crypto APIs listings. | 
**AssetType** | **string** | Defines the type of the supported asset. This could be either \&quot;crypto\&quot; or \&quot;fiat\&quot;. | 
**LatestRate** | [**GetAssetDetailsByAssetIDRILatestRate**](GetAssetDetailsByAssetIDRILatestRate.md) |  | 
**Slug** | Pointer to **string** | Represents the asset&#x60;s unique slug string in Crypto APIs listings. | [optional] 
**SpecificData** | [**GetAssetDetailsByAssetSymbolRIS**](GetAssetDetailsByAssetSymbolRIS.md) |  | 

## Methods

### NewGetAssetDetailsByAssetSymbolRI

`func NewGetAssetDetailsByAssetSymbolRI(assetId string, assetLogo GetAssetDetailsByAssetIDRIAssetLogo, assetName string, assetOriginalSymbol string, assetSymbol string, assetType string, latestRate GetAssetDetailsByAssetIDRILatestRate, specificData GetAssetDetailsByAssetSymbolRIS, ) *GetAssetDetailsByAssetSymbolRI`

NewGetAssetDetailsByAssetSymbolRI instantiates a new GetAssetDetailsByAssetSymbolRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetDetailsByAssetSymbolRIWithDefaults

`func NewGetAssetDetailsByAssetSymbolRIWithDefaults() *GetAssetDetailsByAssetSymbolRI`

NewGetAssetDetailsByAssetSymbolRIWithDefaults instantiates a new GetAssetDetailsByAssetSymbolRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *GetAssetDetailsByAssetSymbolRI) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *GetAssetDetailsByAssetSymbolRI) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *GetAssetDetailsByAssetSymbolRI) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetAssetLogo

`func (o *GetAssetDetailsByAssetSymbolRI) GetAssetLogo() GetAssetDetailsByAssetIDRIAssetLogo`

GetAssetLogo returns the AssetLogo field if non-nil, zero value otherwise.

### GetAssetLogoOk

`func (o *GetAssetDetailsByAssetSymbolRI) GetAssetLogoOk() (*GetAssetDetailsByAssetIDRIAssetLogo, bool)`

GetAssetLogoOk returns a tuple with the AssetLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetLogo

`func (o *GetAssetDetailsByAssetSymbolRI) SetAssetLogo(v GetAssetDetailsByAssetIDRIAssetLogo)`

SetAssetLogo sets AssetLogo field to given value.


### GetAssetName

`func (o *GetAssetDetailsByAssetSymbolRI) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *GetAssetDetailsByAssetSymbolRI) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *GetAssetDetailsByAssetSymbolRI) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.


### GetAssetOriginalSymbol

`func (o *GetAssetDetailsByAssetSymbolRI) GetAssetOriginalSymbol() string`

GetAssetOriginalSymbol returns the AssetOriginalSymbol field if non-nil, zero value otherwise.

### GetAssetOriginalSymbolOk

`func (o *GetAssetDetailsByAssetSymbolRI) GetAssetOriginalSymbolOk() (*string, bool)`

GetAssetOriginalSymbolOk returns a tuple with the AssetOriginalSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetOriginalSymbol

`func (o *GetAssetDetailsByAssetSymbolRI) SetAssetOriginalSymbol(v string)`

SetAssetOriginalSymbol sets AssetOriginalSymbol field to given value.


### GetAssetSymbol

`func (o *GetAssetDetailsByAssetSymbolRI) GetAssetSymbol() string`

GetAssetSymbol returns the AssetSymbol field if non-nil, zero value otherwise.

### GetAssetSymbolOk

`func (o *GetAssetDetailsByAssetSymbolRI) GetAssetSymbolOk() (*string, bool)`

GetAssetSymbolOk returns a tuple with the AssetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSymbol

`func (o *GetAssetDetailsByAssetSymbolRI) SetAssetSymbol(v string)`

SetAssetSymbol sets AssetSymbol field to given value.


### GetAssetType

`func (o *GetAssetDetailsByAssetSymbolRI) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *GetAssetDetailsByAssetSymbolRI) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *GetAssetDetailsByAssetSymbolRI) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetLatestRate

`func (o *GetAssetDetailsByAssetSymbolRI) GetLatestRate() GetAssetDetailsByAssetIDRILatestRate`

GetLatestRate returns the LatestRate field if non-nil, zero value otherwise.

### GetLatestRateOk

`func (o *GetAssetDetailsByAssetSymbolRI) GetLatestRateOk() (*GetAssetDetailsByAssetIDRILatestRate, bool)`

GetLatestRateOk returns a tuple with the LatestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRate

`func (o *GetAssetDetailsByAssetSymbolRI) SetLatestRate(v GetAssetDetailsByAssetIDRILatestRate)`

SetLatestRate sets LatestRate field to given value.


### GetSlug

`func (o *GetAssetDetailsByAssetSymbolRI) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *GetAssetDetailsByAssetSymbolRI) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *GetAssetDetailsByAssetSymbolRI) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *GetAssetDetailsByAssetSymbolRI) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSpecificData

`func (o *GetAssetDetailsByAssetSymbolRI) GetSpecificData() GetAssetDetailsByAssetSymbolRIS`

GetSpecificData returns the SpecificData field if non-nil, zero value otherwise.

### GetSpecificDataOk

`func (o *GetAssetDetailsByAssetSymbolRI) GetSpecificDataOk() (*GetAssetDetailsByAssetSymbolRIS, bool)`

GetSpecificDataOk returns a tuple with the SpecificData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificData

`func (o *GetAssetDetailsByAssetSymbolRI) SetSpecificData(v GetAssetDetailsByAssetSymbolRIS)`

SetSpecificData sets SpecificData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


