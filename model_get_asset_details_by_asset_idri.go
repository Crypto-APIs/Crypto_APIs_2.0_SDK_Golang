/*
CryptoAPIs

Crypto APIs is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.

API version: 2021-03-20
Contact: developers@cryptoapis.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// GetAssetDetailsByAssetIDRI struct for GetAssetDetailsByAssetIDRI
type GetAssetDetailsByAssetIDRI struct {
	// Defines the unique ID of the specific asset.
	AssetId string `json:"assetId"`
	AssetLogo GetAssetDetailsByAssetIDRIAssetLogo `json:"assetLogo"`
	// Specifies the name of the asset in question.
	AssetName string `json:"assetName"`
	// Specifies the asset's original symbol as introduced by its founders.
	AssetOriginalSymbol string `json:"assetOriginalSymbol"`
	// Specifies the asset's unique symbol in the Crypto APIs listings.
	AssetSymbol string `json:"assetSymbol"`
	// Defines the type of the supported asset. This could be either \"crypto\" or \"fiat\".
	AssetType string `json:"assetType"`
	LatestRate GetAssetDetailsByAssetIDRILatestRate `json:"latestRate"`
	// Represents the asset`s unique slug string in Crypto APIs listings.
	Slug *string `json:"slug,omitempty"`
	SpecificData GetAssetDetailsByAssetIDRIS `json:"specificData"`
}

// NewGetAssetDetailsByAssetIDRI instantiates a new GetAssetDetailsByAssetIDRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAssetDetailsByAssetIDRI(assetId string, assetLogo GetAssetDetailsByAssetIDRIAssetLogo, assetName string, assetOriginalSymbol string, assetSymbol string, assetType string, latestRate GetAssetDetailsByAssetIDRILatestRate, specificData GetAssetDetailsByAssetIDRIS) *GetAssetDetailsByAssetIDRI {
	this := GetAssetDetailsByAssetIDRI{}
	this.AssetId = assetId
	this.AssetLogo = assetLogo
	this.AssetName = assetName
	this.AssetOriginalSymbol = assetOriginalSymbol
	this.AssetSymbol = assetSymbol
	this.AssetType = assetType
	this.LatestRate = latestRate
	this.SpecificData = specificData
	return &this
}

// NewGetAssetDetailsByAssetIDRIWithDefaults instantiates a new GetAssetDetailsByAssetIDRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAssetDetailsByAssetIDRIWithDefaults() *GetAssetDetailsByAssetIDRI {
	this := GetAssetDetailsByAssetIDRI{}
	return &this
}

// GetAssetId returns the AssetId field value
func (o *GetAssetDetailsByAssetIDRI) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *GetAssetDetailsByAssetIDRI) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *GetAssetDetailsByAssetIDRI) SetAssetId(v string) {
	o.AssetId = v
}

// GetAssetLogo returns the AssetLogo field value
func (o *GetAssetDetailsByAssetIDRI) GetAssetLogo() GetAssetDetailsByAssetIDRIAssetLogo {
	if o == nil {
		var ret GetAssetDetailsByAssetIDRIAssetLogo
		return ret
	}

	return o.AssetLogo
}

// GetAssetLogoOk returns a tuple with the AssetLogo field value
// and a boolean to check if the value has been set.
func (o *GetAssetDetailsByAssetIDRI) GetAssetLogoOk() (*GetAssetDetailsByAssetIDRIAssetLogo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetLogo, true
}

// SetAssetLogo sets field value
func (o *GetAssetDetailsByAssetIDRI) SetAssetLogo(v GetAssetDetailsByAssetIDRIAssetLogo) {
	o.AssetLogo = v
}

// GetAssetName returns the AssetName field value
func (o *GetAssetDetailsByAssetIDRI) GetAssetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetName
}

// GetAssetNameOk returns a tuple with the AssetName field value
// and a boolean to check if the value has been set.
func (o *GetAssetDetailsByAssetIDRI) GetAssetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetName, true
}

// SetAssetName sets field value
func (o *GetAssetDetailsByAssetIDRI) SetAssetName(v string) {
	o.AssetName = v
}

// GetAssetOriginalSymbol returns the AssetOriginalSymbol field value
func (o *GetAssetDetailsByAssetIDRI) GetAssetOriginalSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetOriginalSymbol
}

// GetAssetOriginalSymbolOk returns a tuple with the AssetOriginalSymbol field value
// and a boolean to check if the value has been set.
func (o *GetAssetDetailsByAssetIDRI) GetAssetOriginalSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetOriginalSymbol, true
}

// SetAssetOriginalSymbol sets field value
func (o *GetAssetDetailsByAssetIDRI) SetAssetOriginalSymbol(v string) {
	o.AssetOriginalSymbol = v
}

// GetAssetSymbol returns the AssetSymbol field value
func (o *GetAssetDetailsByAssetIDRI) GetAssetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetSymbol
}

// GetAssetSymbolOk returns a tuple with the AssetSymbol field value
// and a boolean to check if the value has been set.
func (o *GetAssetDetailsByAssetIDRI) GetAssetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetSymbol, true
}

// SetAssetSymbol sets field value
func (o *GetAssetDetailsByAssetIDRI) SetAssetSymbol(v string) {
	o.AssetSymbol = v
}

// GetAssetType returns the AssetType field value
func (o *GetAssetDetailsByAssetIDRI) GetAssetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value
// and a boolean to check if the value has been set.
func (o *GetAssetDetailsByAssetIDRI) GetAssetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetType, true
}

// SetAssetType sets field value
func (o *GetAssetDetailsByAssetIDRI) SetAssetType(v string) {
	o.AssetType = v
}

// GetLatestRate returns the LatestRate field value
func (o *GetAssetDetailsByAssetIDRI) GetLatestRate() GetAssetDetailsByAssetIDRILatestRate {
	if o == nil {
		var ret GetAssetDetailsByAssetIDRILatestRate
		return ret
	}

	return o.LatestRate
}

// GetLatestRateOk returns a tuple with the LatestRate field value
// and a boolean to check if the value has been set.
func (o *GetAssetDetailsByAssetIDRI) GetLatestRateOk() (*GetAssetDetailsByAssetIDRILatestRate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LatestRate, true
}

// SetLatestRate sets field value
func (o *GetAssetDetailsByAssetIDRI) SetLatestRate(v GetAssetDetailsByAssetIDRILatestRate) {
	o.LatestRate = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *GetAssetDetailsByAssetIDRI) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDetailsByAssetIDRI) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *GetAssetDetailsByAssetIDRI) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *GetAssetDetailsByAssetIDRI) SetSlug(v string) {
	o.Slug = &v
}

// GetSpecificData returns the SpecificData field value
func (o *GetAssetDetailsByAssetIDRI) GetSpecificData() GetAssetDetailsByAssetIDRIS {
	if o == nil {
		var ret GetAssetDetailsByAssetIDRIS
		return ret
	}

	return o.SpecificData
}

// GetSpecificDataOk returns a tuple with the SpecificData field value
// and a boolean to check if the value has been set.
func (o *GetAssetDetailsByAssetIDRI) GetSpecificDataOk() (*GetAssetDetailsByAssetIDRIS, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecificData, true
}

// SetSpecificData sets field value
func (o *GetAssetDetailsByAssetIDRI) SetSpecificData(v GetAssetDetailsByAssetIDRIS) {
	o.SpecificData = v
}

func (o GetAssetDetailsByAssetIDRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assetId"] = o.AssetId
	}
	if true {
		toSerialize["assetLogo"] = o.AssetLogo
	}
	if true {
		toSerialize["assetName"] = o.AssetName
	}
	if true {
		toSerialize["assetOriginalSymbol"] = o.AssetOriginalSymbol
	}
	if true {
		toSerialize["assetSymbol"] = o.AssetSymbol
	}
	if true {
		toSerialize["assetType"] = o.AssetType
	}
	if true {
		toSerialize["latestRate"] = o.LatestRate
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["specificData"] = o.SpecificData
	}
	return json.Marshal(toSerialize)
}

type NullableGetAssetDetailsByAssetIDRI struct {
	value *GetAssetDetailsByAssetIDRI
	isSet bool
}

func (v NullableGetAssetDetailsByAssetIDRI) Get() *GetAssetDetailsByAssetIDRI {
	return v.value
}

func (v *NullableGetAssetDetailsByAssetIDRI) Set(val *GetAssetDetailsByAssetIDRI) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssetDetailsByAssetIDRI) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssetDetailsByAssetIDRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssetDetailsByAssetIDRI(val *GetAssetDetailsByAssetIDRI) *NullableGetAssetDetailsByAssetIDRI {
	return &NullableGetAssetDetailsByAssetIDRI{value: val, isSet: true}
}

func (v NullableGetAssetDetailsByAssetIDRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssetDetailsByAssetIDRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


