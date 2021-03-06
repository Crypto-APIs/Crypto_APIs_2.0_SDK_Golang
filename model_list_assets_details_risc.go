/*
CryptoAPIs

Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.

API version: 2.0.0
Contact: developers@cryptoapis.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// ListAssetsDetailsRISC Crypto Type Data
type ListAssetsDetailsRISC struct {
	// Represents the percentage of the asset's current price against the its price from 1 hour ago.
	Var1HourPriceChangeInPercentage string `json:"1HourPriceChangeInPercentage"`
	// Represents the percentage of the asset's current price against the its price from 1 week ago.
	Var1WeekPriceChangeInPercentage string `json:"1WeekPriceChangeInPercentage"`
	// Represents the percentage of the asset's current price against the its price from 24 hours ago.
	Var24HoursPriceChangeInPercentage string `json:"24HoursPriceChangeInPercentage"`
	// Represents the trading volume of the asset for the time frame of 24 hours.
	Var24HoursTradingVolume string `json:"24HoursTradingVolume"`
	// Represent a subtype of the crypto assets. Could be COIN or TOKEN.
	AssetType string `json:"assetType"`
	// Represents the amount of the asset that is circulating on the market and in public hands.
	CirculatingSupply string `json:"circulatingSupply"`
	// Defines the total market value of the asset's circulating supply in USD.
	MarketCapInUSD string `json:"marketCapInUSD"`
	// Represents the maximum amount of all coins of a specific asset that will ever exist in its lifetime.
	MaxSupply string `json:"maxSupply"`
}

// NewListAssetsDetailsRISC instantiates a new ListAssetsDetailsRISC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAssetsDetailsRISC(var1HourPriceChangeInPercentage string, var1WeekPriceChangeInPercentage string, var24HoursPriceChangeInPercentage string, var24HoursTradingVolume string, assetType string, circulatingSupply string, marketCapInUSD string, maxSupply string) *ListAssetsDetailsRISC {
	this := ListAssetsDetailsRISC{}
	this.Var1HourPriceChangeInPercentage = var1HourPriceChangeInPercentage
	this.Var1WeekPriceChangeInPercentage = var1WeekPriceChangeInPercentage
	this.Var24HoursPriceChangeInPercentage = var24HoursPriceChangeInPercentage
	this.Var24HoursTradingVolume = var24HoursTradingVolume
	this.AssetType = assetType
	this.CirculatingSupply = circulatingSupply
	this.MarketCapInUSD = marketCapInUSD
	this.MaxSupply = maxSupply
	return &this
}

// NewListAssetsDetailsRISCWithDefaults instantiates a new ListAssetsDetailsRISC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAssetsDetailsRISCWithDefaults() *ListAssetsDetailsRISC {
	this := ListAssetsDetailsRISC{}
	return &this
}

// GetVar1HourPriceChangeInPercentage returns the Var1HourPriceChangeInPercentage field value
func (o *ListAssetsDetailsRISC) GetVar1HourPriceChangeInPercentage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Var1HourPriceChangeInPercentage
}

// GetVar1HourPriceChangeInPercentageOk returns a tuple with the Var1HourPriceChangeInPercentage field value
// and a boolean to check if the value has been set.
func (o *ListAssetsDetailsRISC) GetVar1HourPriceChangeInPercentageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var1HourPriceChangeInPercentage, true
}

// SetVar1HourPriceChangeInPercentage sets field value
func (o *ListAssetsDetailsRISC) SetVar1HourPriceChangeInPercentage(v string) {
	o.Var1HourPriceChangeInPercentage = v
}

// GetVar1WeekPriceChangeInPercentage returns the Var1WeekPriceChangeInPercentage field value
func (o *ListAssetsDetailsRISC) GetVar1WeekPriceChangeInPercentage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Var1WeekPriceChangeInPercentage
}

// GetVar1WeekPriceChangeInPercentageOk returns a tuple with the Var1WeekPriceChangeInPercentage field value
// and a boolean to check if the value has been set.
func (o *ListAssetsDetailsRISC) GetVar1WeekPriceChangeInPercentageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var1WeekPriceChangeInPercentage, true
}

// SetVar1WeekPriceChangeInPercentage sets field value
func (o *ListAssetsDetailsRISC) SetVar1WeekPriceChangeInPercentage(v string) {
	o.Var1WeekPriceChangeInPercentage = v
}

// GetVar24HoursPriceChangeInPercentage returns the Var24HoursPriceChangeInPercentage field value
func (o *ListAssetsDetailsRISC) GetVar24HoursPriceChangeInPercentage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Var24HoursPriceChangeInPercentage
}

// GetVar24HoursPriceChangeInPercentageOk returns a tuple with the Var24HoursPriceChangeInPercentage field value
// and a boolean to check if the value has been set.
func (o *ListAssetsDetailsRISC) GetVar24HoursPriceChangeInPercentageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var24HoursPriceChangeInPercentage, true
}

// SetVar24HoursPriceChangeInPercentage sets field value
func (o *ListAssetsDetailsRISC) SetVar24HoursPriceChangeInPercentage(v string) {
	o.Var24HoursPriceChangeInPercentage = v
}

// GetVar24HoursTradingVolume returns the Var24HoursTradingVolume field value
func (o *ListAssetsDetailsRISC) GetVar24HoursTradingVolume() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Var24HoursTradingVolume
}

// GetVar24HoursTradingVolumeOk returns a tuple with the Var24HoursTradingVolume field value
// and a boolean to check if the value has been set.
func (o *ListAssetsDetailsRISC) GetVar24HoursTradingVolumeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var24HoursTradingVolume, true
}

// SetVar24HoursTradingVolume sets field value
func (o *ListAssetsDetailsRISC) SetVar24HoursTradingVolume(v string) {
	o.Var24HoursTradingVolume = v
}

// GetAssetType returns the AssetType field value
func (o *ListAssetsDetailsRISC) GetAssetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value
// and a boolean to check if the value has been set.
func (o *ListAssetsDetailsRISC) GetAssetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetType, true
}

// SetAssetType sets field value
func (o *ListAssetsDetailsRISC) SetAssetType(v string) {
	o.AssetType = v
}

// GetCirculatingSupply returns the CirculatingSupply field value
func (o *ListAssetsDetailsRISC) GetCirculatingSupply() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CirculatingSupply
}

// GetCirculatingSupplyOk returns a tuple with the CirculatingSupply field value
// and a boolean to check if the value has been set.
func (o *ListAssetsDetailsRISC) GetCirculatingSupplyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CirculatingSupply, true
}

// SetCirculatingSupply sets field value
func (o *ListAssetsDetailsRISC) SetCirculatingSupply(v string) {
	o.CirculatingSupply = v
}

// GetMarketCapInUSD returns the MarketCapInUSD field value
func (o *ListAssetsDetailsRISC) GetMarketCapInUSD() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketCapInUSD
}

// GetMarketCapInUSDOk returns a tuple with the MarketCapInUSD field value
// and a boolean to check if the value has been set.
func (o *ListAssetsDetailsRISC) GetMarketCapInUSDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketCapInUSD, true
}

// SetMarketCapInUSD sets field value
func (o *ListAssetsDetailsRISC) SetMarketCapInUSD(v string) {
	o.MarketCapInUSD = v
}

// GetMaxSupply returns the MaxSupply field value
func (o *ListAssetsDetailsRISC) GetMaxSupply() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxSupply
}

// GetMaxSupplyOk returns a tuple with the MaxSupply field value
// and a boolean to check if the value has been set.
func (o *ListAssetsDetailsRISC) GetMaxSupplyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxSupply, true
}

// SetMaxSupply sets field value
func (o *ListAssetsDetailsRISC) SetMaxSupply(v string) {
	o.MaxSupply = v
}

func (o ListAssetsDetailsRISC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["1HourPriceChangeInPercentage"] = o.Var1HourPriceChangeInPercentage
	}
	if true {
		toSerialize["1WeekPriceChangeInPercentage"] = o.Var1WeekPriceChangeInPercentage
	}
	if true {
		toSerialize["24HoursPriceChangeInPercentage"] = o.Var24HoursPriceChangeInPercentage
	}
	if true {
		toSerialize["24HoursTradingVolume"] = o.Var24HoursTradingVolume
	}
	if true {
		toSerialize["assetType"] = o.AssetType
	}
	if true {
		toSerialize["circulatingSupply"] = o.CirculatingSupply
	}
	if true {
		toSerialize["marketCapInUSD"] = o.MarketCapInUSD
	}
	if true {
		toSerialize["maxSupply"] = o.MaxSupply
	}
	return json.Marshal(toSerialize)
}

type NullableListAssetsDetailsRISC struct {
	value *ListAssetsDetailsRISC
	isSet bool
}

func (v NullableListAssetsDetailsRISC) Get() *ListAssetsDetailsRISC {
	return v.value
}

func (v *NullableListAssetsDetailsRISC) Set(val *ListAssetsDetailsRISC) {
	v.value = val
	v.isSet = true
}

func (v NullableListAssetsDetailsRISC) IsSet() bool {
	return v.isSet
}

func (v *NullableListAssetsDetailsRISC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAssetsDetailsRISC(val *ListAssetsDetailsRISC) *NullableListAssetsDetailsRISC {
	return &NullableListAssetsDetailsRISC{value: val, isSet: true}
}

func (v NullableListAssetsDetailsRISC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAssetsDetailsRISC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


