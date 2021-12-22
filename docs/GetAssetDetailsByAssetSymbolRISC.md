# GetAssetDetailsByAssetSymbolRISC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var1HourPriceChangeInPercentage** | **string** | Represents the percentage of the asset&#39;s current price against the its price from 1 hour ago. | 
**Var1WeekPriceChangeInPercentage** | **string** | Represents the percentage of the asset&#39;s current price against the its price from 1 week ago. | 
**Var24HoursPriceChangeInPercentage** | **string** | Represents the percentage of the asset&#39;s current price against the its price from 24 hours ago. | 
**Var24HoursTradingVolume** | **string** | Represents the trading volume of the asset for the time frame of 24 hours. | 
**AssetType** | **string** | Defines the type of the supported asset. This could be either \&quot;crypto\&quot; or \&quot;fiat\&quot;. | 
**CirculatingSupply** | **string** | Represents the amount of the asset that is circulating on the market and in public hands. | 
**MarketCapInUSD** | **string** | Defines the total market value of the asset&#39;s circulating supply in USD. | 
**MaxSupply** | **string** | Represents the maximum amount of all coins of a specific asset that will ever exist in its lifetime. | 

## Methods

### NewGetAssetDetailsByAssetSymbolRISC

`func NewGetAssetDetailsByAssetSymbolRISC(var1HourPriceChangeInPercentage string, var1WeekPriceChangeInPercentage string, var24HoursPriceChangeInPercentage string, var24HoursTradingVolume string, assetType string, circulatingSupply string, marketCapInUSD string, maxSupply string, ) *GetAssetDetailsByAssetSymbolRISC`

NewGetAssetDetailsByAssetSymbolRISC instantiates a new GetAssetDetailsByAssetSymbolRISC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetDetailsByAssetSymbolRISCWithDefaults

`func NewGetAssetDetailsByAssetSymbolRISCWithDefaults() *GetAssetDetailsByAssetSymbolRISC`

NewGetAssetDetailsByAssetSymbolRISCWithDefaults instantiates a new GetAssetDetailsByAssetSymbolRISC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar1HourPriceChangeInPercentage

`func (o *GetAssetDetailsByAssetSymbolRISC) GetVar1HourPriceChangeInPercentage() string`

GetVar1HourPriceChangeInPercentage returns the Var1HourPriceChangeInPercentage field if non-nil, zero value otherwise.

### GetVar1HourPriceChangeInPercentageOk

`func (o *GetAssetDetailsByAssetSymbolRISC) GetVar1HourPriceChangeInPercentageOk() (*string, bool)`

GetVar1HourPriceChangeInPercentageOk returns a tuple with the Var1HourPriceChangeInPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1HourPriceChangeInPercentage

`func (o *GetAssetDetailsByAssetSymbolRISC) SetVar1HourPriceChangeInPercentage(v string)`

SetVar1HourPriceChangeInPercentage sets Var1HourPriceChangeInPercentage field to given value.


### GetVar1WeekPriceChangeInPercentage

`func (o *GetAssetDetailsByAssetSymbolRISC) GetVar1WeekPriceChangeInPercentage() string`

GetVar1WeekPriceChangeInPercentage returns the Var1WeekPriceChangeInPercentage field if non-nil, zero value otherwise.

### GetVar1WeekPriceChangeInPercentageOk

`func (o *GetAssetDetailsByAssetSymbolRISC) GetVar1WeekPriceChangeInPercentageOk() (*string, bool)`

GetVar1WeekPriceChangeInPercentageOk returns a tuple with the Var1WeekPriceChangeInPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1WeekPriceChangeInPercentage

`func (o *GetAssetDetailsByAssetSymbolRISC) SetVar1WeekPriceChangeInPercentage(v string)`

SetVar1WeekPriceChangeInPercentage sets Var1WeekPriceChangeInPercentage field to given value.


### GetVar24HoursPriceChangeInPercentage

`func (o *GetAssetDetailsByAssetSymbolRISC) GetVar24HoursPriceChangeInPercentage() string`

GetVar24HoursPriceChangeInPercentage returns the Var24HoursPriceChangeInPercentage field if non-nil, zero value otherwise.

### GetVar24HoursPriceChangeInPercentageOk

`func (o *GetAssetDetailsByAssetSymbolRISC) GetVar24HoursPriceChangeInPercentageOk() (*string, bool)`

GetVar24HoursPriceChangeInPercentageOk returns a tuple with the Var24HoursPriceChangeInPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar24HoursPriceChangeInPercentage

`func (o *GetAssetDetailsByAssetSymbolRISC) SetVar24HoursPriceChangeInPercentage(v string)`

SetVar24HoursPriceChangeInPercentage sets Var24HoursPriceChangeInPercentage field to given value.


### GetVar24HoursTradingVolume

`func (o *GetAssetDetailsByAssetSymbolRISC) GetVar24HoursTradingVolume() string`

GetVar24HoursTradingVolume returns the Var24HoursTradingVolume field if non-nil, zero value otherwise.

### GetVar24HoursTradingVolumeOk

`func (o *GetAssetDetailsByAssetSymbolRISC) GetVar24HoursTradingVolumeOk() (*string, bool)`

GetVar24HoursTradingVolumeOk returns a tuple with the Var24HoursTradingVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar24HoursTradingVolume

`func (o *GetAssetDetailsByAssetSymbolRISC) SetVar24HoursTradingVolume(v string)`

SetVar24HoursTradingVolume sets Var24HoursTradingVolume field to given value.


### GetAssetType

`func (o *GetAssetDetailsByAssetSymbolRISC) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *GetAssetDetailsByAssetSymbolRISC) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *GetAssetDetailsByAssetSymbolRISC) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetCirculatingSupply

`func (o *GetAssetDetailsByAssetSymbolRISC) GetCirculatingSupply() string`

GetCirculatingSupply returns the CirculatingSupply field if non-nil, zero value otherwise.

### GetCirculatingSupplyOk

`func (o *GetAssetDetailsByAssetSymbolRISC) GetCirculatingSupplyOk() (*string, bool)`

GetCirculatingSupplyOk returns a tuple with the CirculatingSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCirculatingSupply

`func (o *GetAssetDetailsByAssetSymbolRISC) SetCirculatingSupply(v string)`

SetCirculatingSupply sets CirculatingSupply field to given value.


### GetMarketCapInUSD

`func (o *GetAssetDetailsByAssetSymbolRISC) GetMarketCapInUSD() string`

GetMarketCapInUSD returns the MarketCapInUSD field if non-nil, zero value otherwise.

### GetMarketCapInUSDOk

`func (o *GetAssetDetailsByAssetSymbolRISC) GetMarketCapInUSDOk() (*string, bool)`

GetMarketCapInUSDOk returns a tuple with the MarketCapInUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapInUSD

`func (o *GetAssetDetailsByAssetSymbolRISC) SetMarketCapInUSD(v string)`

SetMarketCapInUSD sets MarketCapInUSD field to given value.


### GetMaxSupply

`func (o *GetAssetDetailsByAssetSymbolRISC) GetMaxSupply() string`

GetMaxSupply returns the MaxSupply field if non-nil, zero value otherwise.

### GetMaxSupplyOk

`func (o *GetAssetDetailsByAssetSymbolRISC) GetMaxSupplyOk() (*string, bool)`

GetMaxSupplyOk returns a tuple with the MaxSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSupply

`func (o *GetAssetDetailsByAssetSymbolRISC) SetMaxSupply(v string)`

SetMaxSupply sets MaxSupply field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


