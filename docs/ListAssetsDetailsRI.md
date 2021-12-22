# ListAssetsDetailsRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** | Defines the unique ID of the specific asset. | 
**AssetLogo** | [**ListAssetsDetailsRIAssetLogo**](ListAssetsDetailsRIAssetLogo.md) |  | 
**AssetName** | **string** | Specifies the name of the asset in question. | 
**AssetOriginalSymbol** | **string** | Specifies the asset&#39;s original symbol as introduced by its founders. | 
**AssetSymbol** | **string** | Specifies the asset&#39;s unique symbol in the Crypto APIs listings. | 
**AssetType** | **string** | Defines the type of the supported asset. This could be either \&quot;crypto\&quot; or \&quot;fiat\&quot;. | 
**LatestRate** | [**ListAssetsDetailsRILatestRate**](ListAssetsDetailsRILatestRate.md) |  | 
**Slug** | Pointer to **string** | Represents the asset&#x60;s unique slug string in Crypto APIs listings. | [optional] 
**SpecificData** | [**ListAssetsDetailsRIS**](ListAssetsDetailsRIS.md) |  | 

## Methods

### NewListAssetsDetailsRI

`func NewListAssetsDetailsRI(assetId string, assetLogo ListAssetsDetailsRIAssetLogo, assetName string, assetOriginalSymbol string, assetSymbol string, assetType string, latestRate ListAssetsDetailsRILatestRate, specificData ListAssetsDetailsRIS, ) *ListAssetsDetailsRI`

NewListAssetsDetailsRI instantiates a new ListAssetsDetailsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAssetsDetailsRIWithDefaults

`func NewListAssetsDetailsRIWithDefaults() *ListAssetsDetailsRI`

NewListAssetsDetailsRIWithDefaults instantiates a new ListAssetsDetailsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *ListAssetsDetailsRI) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ListAssetsDetailsRI) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ListAssetsDetailsRI) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetAssetLogo

`func (o *ListAssetsDetailsRI) GetAssetLogo() ListAssetsDetailsRIAssetLogo`

GetAssetLogo returns the AssetLogo field if non-nil, zero value otherwise.

### GetAssetLogoOk

`func (o *ListAssetsDetailsRI) GetAssetLogoOk() (*ListAssetsDetailsRIAssetLogo, bool)`

GetAssetLogoOk returns a tuple with the AssetLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetLogo

`func (o *ListAssetsDetailsRI) SetAssetLogo(v ListAssetsDetailsRIAssetLogo)`

SetAssetLogo sets AssetLogo field to given value.


### GetAssetName

`func (o *ListAssetsDetailsRI) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *ListAssetsDetailsRI) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *ListAssetsDetailsRI) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.


### GetAssetOriginalSymbol

`func (o *ListAssetsDetailsRI) GetAssetOriginalSymbol() string`

GetAssetOriginalSymbol returns the AssetOriginalSymbol field if non-nil, zero value otherwise.

### GetAssetOriginalSymbolOk

`func (o *ListAssetsDetailsRI) GetAssetOriginalSymbolOk() (*string, bool)`

GetAssetOriginalSymbolOk returns a tuple with the AssetOriginalSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetOriginalSymbol

`func (o *ListAssetsDetailsRI) SetAssetOriginalSymbol(v string)`

SetAssetOriginalSymbol sets AssetOriginalSymbol field to given value.


### GetAssetSymbol

`func (o *ListAssetsDetailsRI) GetAssetSymbol() string`

GetAssetSymbol returns the AssetSymbol field if non-nil, zero value otherwise.

### GetAssetSymbolOk

`func (o *ListAssetsDetailsRI) GetAssetSymbolOk() (*string, bool)`

GetAssetSymbolOk returns a tuple with the AssetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSymbol

`func (o *ListAssetsDetailsRI) SetAssetSymbol(v string)`

SetAssetSymbol sets AssetSymbol field to given value.


### GetAssetType

`func (o *ListAssetsDetailsRI) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *ListAssetsDetailsRI) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *ListAssetsDetailsRI) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetLatestRate

`func (o *ListAssetsDetailsRI) GetLatestRate() ListAssetsDetailsRILatestRate`

GetLatestRate returns the LatestRate field if non-nil, zero value otherwise.

### GetLatestRateOk

`func (o *ListAssetsDetailsRI) GetLatestRateOk() (*ListAssetsDetailsRILatestRate, bool)`

GetLatestRateOk returns a tuple with the LatestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRate

`func (o *ListAssetsDetailsRI) SetLatestRate(v ListAssetsDetailsRILatestRate)`

SetLatestRate sets LatestRate field to given value.


### GetSlug

`func (o *ListAssetsDetailsRI) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ListAssetsDetailsRI) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ListAssetsDetailsRI) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ListAssetsDetailsRI) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSpecificData

`func (o *ListAssetsDetailsRI) GetSpecificData() ListAssetsDetailsRIS`

GetSpecificData returns the SpecificData field if non-nil, zero value otherwise.

### GetSpecificDataOk

`func (o *ListAssetsDetailsRI) GetSpecificDataOk() (*ListAssetsDetailsRIS, bool)`

GetSpecificDataOk returns a tuple with the SpecificData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificData

`func (o *ListAssetsDetailsRI) SetSpecificData(v ListAssetsDetailsRIS)`

SetSpecificData sets SpecificData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


