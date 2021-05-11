# ListAssetsDetailsResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** | Defines the unique ID of the specific asset. | 
**AssetLogo** | [**ListAssetsDetailsResponseItemAssetLogo**](ListAssetsDetailsResponseItemAssetLogo.md) |  | 
**AssetName** | **string** | Specifies the name of the asset in question. | 
**AssetOriginalSymbol** | **string** | Specifies the asset&#39;s original symbol as introduced by its founders. | 
**AssetSymbol** | **string** | Specifies the asset&#39;s unique symbol in the Crypto APIs listings. | 
**AssetType** | **string** | Defines the type of the supported asset. This could be either \&quot;crypto\&quot; or \&quot;fiat\&quot;. | 
**LatestRate** | [**ListAssetsDetailsResponseItemLatestRate**](ListAssetsDetailsResponseItemLatestRate.md) |  | 
**SpecificData** | [**ListAssetsDetailsResponseItemSpecificData**](ListAssetsDetailsResponseItemSpecificData.md) |  | 

## Methods

### NewListAssetsDetailsResponseItem

`func NewListAssetsDetailsResponseItem(assetId string, assetLogo ListAssetsDetailsResponseItemAssetLogo, assetName string, assetOriginalSymbol string, assetSymbol string, assetType string, latestRate ListAssetsDetailsResponseItemLatestRate, specificData ListAssetsDetailsResponseItemSpecificData, ) *ListAssetsDetailsResponseItem`

NewListAssetsDetailsResponseItem instantiates a new ListAssetsDetailsResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAssetsDetailsResponseItemWithDefaults

`func NewListAssetsDetailsResponseItemWithDefaults() *ListAssetsDetailsResponseItem`

NewListAssetsDetailsResponseItemWithDefaults instantiates a new ListAssetsDetailsResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *ListAssetsDetailsResponseItem) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ListAssetsDetailsResponseItem) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ListAssetsDetailsResponseItem) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetAssetLogo

`func (o *ListAssetsDetailsResponseItem) GetAssetLogo() ListAssetsDetailsResponseItemAssetLogo`

GetAssetLogo returns the AssetLogo field if non-nil, zero value otherwise.

### GetAssetLogoOk

`func (o *ListAssetsDetailsResponseItem) GetAssetLogoOk() (*ListAssetsDetailsResponseItemAssetLogo, bool)`

GetAssetLogoOk returns a tuple with the AssetLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetLogo

`func (o *ListAssetsDetailsResponseItem) SetAssetLogo(v ListAssetsDetailsResponseItemAssetLogo)`

SetAssetLogo sets AssetLogo field to given value.


### GetAssetName

`func (o *ListAssetsDetailsResponseItem) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *ListAssetsDetailsResponseItem) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *ListAssetsDetailsResponseItem) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.


### GetAssetOriginalSymbol

`func (o *ListAssetsDetailsResponseItem) GetAssetOriginalSymbol() string`

GetAssetOriginalSymbol returns the AssetOriginalSymbol field if non-nil, zero value otherwise.

### GetAssetOriginalSymbolOk

`func (o *ListAssetsDetailsResponseItem) GetAssetOriginalSymbolOk() (*string, bool)`

GetAssetOriginalSymbolOk returns a tuple with the AssetOriginalSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetOriginalSymbol

`func (o *ListAssetsDetailsResponseItem) SetAssetOriginalSymbol(v string)`

SetAssetOriginalSymbol sets AssetOriginalSymbol field to given value.


### GetAssetSymbol

`func (o *ListAssetsDetailsResponseItem) GetAssetSymbol() string`

GetAssetSymbol returns the AssetSymbol field if non-nil, zero value otherwise.

### GetAssetSymbolOk

`func (o *ListAssetsDetailsResponseItem) GetAssetSymbolOk() (*string, bool)`

GetAssetSymbolOk returns a tuple with the AssetSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSymbol

`func (o *ListAssetsDetailsResponseItem) SetAssetSymbol(v string)`

SetAssetSymbol sets AssetSymbol field to given value.


### GetAssetType

`func (o *ListAssetsDetailsResponseItem) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *ListAssetsDetailsResponseItem) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *ListAssetsDetailsResponseItem) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetLatestRate

`func (o *ListAssetsDetailsResponseItem) GetLatestRate() ListAssetsDetailsResponseItemLatestRate`

GetLatestRate returns the LatestRate field if non-nil, zero value otherwise.

### GetLatestRateOk

`func (o *ListAssetsDetailsResponseItem) GetLatestRateOk() (*ListAssetsDetailsResponseItemLatestRate, bool)`

GetLatestRateOk returns a tuple with the LatestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRate

`func (o *ListAssetsDetailsResponseItem) SetLatestRate(v ListAssetsDetailsResponseItemLatestRate)`

SetLatestRate sets LatestRate field to given value.


### GetSpecificData

`func (o *ListAssetsDetailsResponseItem) GetSpecificData() ListAssetsDetailsResponseItemSpecificData`

GetSpecificData returns the SpecificData field if non-nil, zero value otherwise.

### GetSpecificDataOk

`func (o *ListAssetsDetailsResponseItem) GetSpecificDataOk() (*ListAssetsDetailsResponseItemSpecificData, bool)`

GetSpecificDataOk returns a tuple with the SpecificData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificData

`func (o *ListAssetsDetailsResponseItem) SetSpecificData(v ListAssetsDetailsResponseItemSpecificData)`

SetSpecificData sets SpecificData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


