# GetAssetDetailsByAssetIDRI

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
**SpecificData** | [**GetAssetDetailsByAssetIDRIS**](GetAssetDetailsByAssetIDRIS.md) |  | 

## Methods

### NewGetAssetDetailsByAssetIDRI

`func NewGetAssetDetailsByAssetIDRI(assetId string, assetLogo GetAssetDetailsByAssetIDRIAssetLogo, assetName string, assetOriginalSymbol string, assetSymbol string, assetType string, latestRate GetAssetDetailsByAssetIDRILatestRate, specificData GetAssetDetailsByAssetIDRIS, ) *GetAssetDetailsByAssetIDRI`

NewGetAssetDetailsByAssetIDRI instantiates a new GetAssetDetailsByAssetIDRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetDetailsByAssetIDRIWithDefaults

`func NewGetAssetDetailsByAssetIDRIWithDefaults() *GetAssetDetailsByAssetIDRI`

NewGetAssetDetailsByAssetIDRIWithDefaults instantiates a new GetAssetDetailsByAssetIDRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *GetAssetDetailsByAssetIDRI) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *GetAssetDetailsByAssetIDRI) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *GetAssetDetailsByAssetIDRI) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetAssetLogo

`func (o *GetAssetDetailsByAssetIDRI) GetAssetLogo() GetAssetDetailsByAssetIDRIAssetLogo`

GetAssetLogo returns the AssetLogo field if non-nil, zero value otherwise.

### GetAssetLogoOk

`func (o *GetAssetDetailsByAssetIDRI) GetAssetLogoOk() (*GetAssetDetailsByAssetIDRIAssetLogo, bool)`

GetAssetLogoOk returns a tuple with the AssetLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetLogo

`func (o *GetAssetDetailsByAssetIDRI) SetAssetLogo(v GetAssetDetailsByAssetIDRIAssetLogo)`

SetAssetLogo sets AssetLogo field to given value.


### GetAssetName

`func (o *GetAssetDetailsByAssetIDRI) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *GetAssetDetailsByAssetIDRI) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *GetAssetDetailsByAssetIDRI) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.


### GetAssetOriginalSymbol

`func (o *GetAssetDetailsByAssetIDRI) GetAssetOriginalSymbol() string`

GetAssetOriginalSymbol returns the AssetOriginalSymbol field if non-nil, zero value otherwise.

### GetAssetOriginalSymbolOk

`func (o *GetAssetDetailsByAssetIDRI) GetAssetOriginalSymbolOk() (*string, bool)`

GetAssetOriginalSymbolOk returns a tuple with the AssetOriginalSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetOriginalSymbol

`func (o *GetAssetDetailsByAssetIDRI) SetAssetOriginalSymbol(v string)`

SetAssetOriginalSymbol sets AssetOriginalSymbol field to given value.


### GetAssetSymbol

`func (o *GetAssetDetailsByAssetIDRI) GetAssetSymbol() string`

GetAssetSymbol returns the AssetSymbol field if non-nil, zero value otherwise.

### GetAssetSymbolOk

`func (o *GetAssetDetailsByAssetIDRI) GetAssetSymbolOk() (*string, bool)`

GetAssetSymbolOk returns a tuple with the AssetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSymbol

`func (o *GetAssetDetailsByAssetIDRI) SetAssetSymbol(v string)`

SetAssetSymbol sets AssetSymbol field to given value.


### GetAssetType

`func (o *GetAssetDetailsByAssetIDRI) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *GetAssetDetailsByAssetIDRI) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *GetAssetDetailsByAssetIDRI) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetLatestRate

`func (o *GetAssetDetailsByAssetIDRI) GetLatestRate() GetAssetDetailsByAssetIDRILatestRate`

GetLatestRate returns the LatestRate field if non-nil, zero value otherwise.

### GetLatestRateOk

`func (o *GetAssetDetailsByAssetIDRI) GetLatestRateOk() (*GetAssetDetailsByAssetIDRILatestRate, bool)`

GetLatestRateOk returns a tuple with the LatestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRate

`func (o *GetAssetDetailsByAssetIDRI) SetLatestRate(v GetAssetDetailsByAssetIDRILatestRate)`

SetLatestRate sets LatestRate field to given value.


### GetSlug

`func (o *GetAssetDetailsByAssetIDRI) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *GetAssetDetailsByAssetIDRI) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *GetAssetDetailsByAssetIDRI) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *GetAssetDetailsByAssetIDRI) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSpecificData

`func (o *GetAssetDetailsByAssetIDRI) GetSpecificData() GetAssetDetailsByAssetIDRIS`

GetSpecificData returns the SpecificData field if non-nil, zero value otherwise.

### GetSpecificDataOk

`func (o *GetAssetDetailsByAssetIDRI) GetSpecificDataOk() (*GetAssetDetailsByAssetIDRIS, bool)`

GetSpecificDataOk returns a tuple with the SpecificData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificData

`func (o *GetAssetDetailsByAssetIDRI) SetSpecificData(v GetAssetDetailsByAssetIDRIS)`

SetSpecificData sets SpecificData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


