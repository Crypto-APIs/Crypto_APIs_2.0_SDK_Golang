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

// DecodeRawTransactionHexRISE Ethereum
type DecodeRawTransactionHexRISE struct {
	// Defines the approximate fee value. When isConfirmed is True - Defines the amount of the transaction fee When isConfirmed is False - For ETH-based blockchains this attribute represents the max fee value.
	ApproximateFee *string `json:"approximateFee,omitempty"`
	// Defines the approximate minimum fee that is required for the transaction.
	ApproximateMinimumRequiredFee *string `json:"approximateMinimumRequiredFee,omitempty"`
	// Represents the amount of gas used by this specific transaction alone.
	GasLimit string `json:"gasLimit"`
	// Represents the amount of gas paid for the data in the transaction.
	GasPaidForData *string `json:"gasPaidForData,omitempty"`
	// Represents the price offered to the miner to purchase this amount of gas.
	GasPrice *string `json:"gasPrice,omitempty"`
	// Represents additional information that is required for the transaction.
	InputData *string `json:"inputData,omitempty"`
	// Defines the maximum amount that customer is willing to pay per unit of gas to get his transaction included in a block.
	MaxFeePerGas *string `json:"maxFeePerGas,omitempty"`
	// Represents determined by the user value that is paid directly to miners.
	MaxFeePriorityPerGas *string `json:"maxFeePriorityPerGas,omitempty"`
	// Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender's address.
	Nonce int32 `json:"nonce"`
	// Represents output of an ECDSA signature.
	R *string `json:"r,omitempty"`
	// The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one recipient.
	Recipient string `json:"recipient"`
	// Represents output of an ECDSA signature.
	S *string `json:"s,omitempty"`
	// Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender.
	Sender string `json:"sender"`
	// Specifies the transaction type as one from three options: if response returns a `\"0\"` it means the raw transaction includes legacy transaction data, if it is `\"1\"` - includes access lists for EIP2930, and if it is `\"2\"` - EIP1559 data.
	Type int32 `json:"type"`
	// Defines the the recovery id.
	V *string `json:"v,omitempty"`
	// Represents the sent/received amount.
	Value *string `json:"value,omitempty"`
}

// NewDecodeRawTransactionHexRISE instantiates a new DecodeRawTransactionHexRISE object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecodeRawTransactionHexRISE(gasLimit string, nonce int32, recipient string, sender string, type_ int32) *DecodeRawTransactionHexRISE {
	this := DecodeRawTransactionHexRISE{}
	this.GasLimit = gasLimit
	this.Nonce = nonce
	this.Recipient = recipient
	this.Sender = sender
	this.Type = type_
	return &this
}

// NewDecodeRawTransactionHexRISEWithDefaults instantiates a new DecodeRawTransactionHexRISE object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodeRawTransactionHexRISEWithDefaults() *DecodeRawTransactionHexRISE {
	this := DecodeRawTransactionHexRISE{}
	return &this
}

// GetApproximateFee returns the ApproximateFee field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISE) GetApproximateFee() string {
	if o == nil || o.ApproximateFee == nil {
		var ret string
		return ret
	}
	return *o.ApproximateFee
}

// GetApproximateFeeOk returns a tuple with the ApproximateFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetApproximateFeeOk() (*string, bool) {
	if o == nil || o.ApproximateFee == nil {
		return nil, false
	}
	return o.ApproximateFee, true
}

// HasApproximateFee returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISE) HasApproximateFee() bool {
	if o != nil && o.ApproximateFee != nil {
		return true
	}

	return false
}

// SetApproximateFee gets a reference to the given string and assigns it to the ApproximateFee field.
func (o *DecodeRawTransactionHexRISE) SetApproximateFee(v string) {
	o.ApproximateFee = &v
}

// GetApproximateMinimumRequiredFee returns the ApproximateMinimumRequiredFee field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISE) GetApproximateMinimumRequiredFee() string {
	if o == nil || o.ApproximateMinimumRequiredFee == nil {
		var ret string
		return ret
	}
	return *o.ApproximateMinimumRequiredFee
}

// GetApproximateMinimumRequiredFeeOk returns a tuple with the ApproximateMinimumRequiredFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetApproximateMinimumRequiredFeeOk() (*string, bool) {
	if o == nil || o.ApproximateMinimumRequiredFee == nil {
		return nil, false
	}
	return o.ApproximateMinimumRequiredFee, true
}

// HasApproximateMinimumRequiredFee returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISE) HasApproximateMinimumRequiredFee() bool {
	if o != nil && o.ApproximateMinimumRequiredFee != nil {
		return true
	}

	return false
}

// SetApproximateMinimumRequiredFee gets a reference to the given string and assigns it to the ApproximateMinimumRequiredFee field.
func (o *DecodeRawTransactionHexRISE) SetApproximateMinimumRequiredFee(v string) {
	o.ApproximateMinimumRequiredFee = &v
}

// GetGasLimit returns the GasLimit field value
func (o *DecodeRawTransactionHexRISE) GetGasLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetGasLimitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *DecodeRawTransactionHexRISE) SetGasLimit(v string) {
	o.GasLimit = v
}

// GetGasPaidForData returns the GasPaidForData field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISE) GetGasPaidForData() string {
	if o == nil || o.GasPaidForData == nil {
		var ret string
		return ret
	}
	return *o.GasPaidForData
}

// GetGasPaidForDataOk returns a tuple with the GasPaidForData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetGasPaidForDataOk() (*string, bool) {
	if o == nil || o.GasPaidForData == nil {
		return nil, false
	}
	return o.GasPaidForData, true
}

// HasGasPaidForData returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISE) HasGasPaidForData() bool {
	if o != nil && o.GasPaidForData != nil {
		return true
	}

	return false
}

// SetGasPaidForData gets a reference to the given string and assigns it to the GasPaidForData field.
func (o *DecodeRawTransactionHexRISE) SetGasPaidForData(v string) {
	o.GasPaidForData = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISE) GetGasPrice() string {
	if o == nil || o.GasPrice == nil {
		var ret string
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetGasPriceOk() (*string, bool) {
	if o == nil || o.GasPrice == nil {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISE) HasGasPrice() bool {
	if o != nil && o.GasPrice != nil {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given string and assigns it to the GasPrice field.
func (o *DecodeRawTransactionHexRISE) SetGasPrice(v string) {
	o.GasPrice = &v
}

// GetInputData returns the InputData field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISE) GetInputData() string {
	if o == nil || o.InputData == nil {
		var ret string
		return ret
	}
	return *o.InputData
}

// GetInputDataOk returns a tuple with the InputData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetInputDataOk() (*string, bool) {
	if o == nil || o.InputData == nil {
		return nil, false
	}
	return o.InputData, true
}

// HasInputData returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISE) HasInputData() bool {
	if o != nil && o.InputData != nil {
		return true
	}

	return false
}

// SetInputData gets a reference to the given string and assigns it to the InputData field.
func (o *DecodeRawTransactionHexRISE) SetInputData(v string) {
	o.InputData = &v
}

// GetMaxFeePerGas returns the MaxFeePerGas field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISE) GetMaxFeePerGas() string {
	if o == nil || o.MaxFeePerGas == nil {
		var ret string
		return ret
	}
	return *o.MaxFeePerGas
}

// GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetMaxFeePerGasOk() (*string, bool) {
	if o == nil || o.MaxFeePerGas == nil {
		return nil, false
	}
	return o.MaxFeePerGas, true
}

// HasMaxFeePerGas returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISE) HasMaxFeePerGas() bool {
	if o != nil && o.MaxFeePerGas != nil {
		return true
	}

	return false
}

// SetMaxFeePerGas gets a reference to the given string and assigns it to the MaxFeePerGas field.
func (o *DecodeRawTransactionHexRISE) SetMaxFeePerGas(v string) {
	o.MaxFeePerGas = &v
}

// GetMaxFeePriorityPerGas returns the MaxFeePriorityPerGas field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISE) GetMaxFeePriorityPerGas() string {
	if o == nil || o.MaxFeePriorityPerGas == nil {
		var ret string
		return ret
	}
	return *o.MaxFeePriorityPerGas
}

// GetMaxFeePriorityPerGasOk returns a tuple with the MaxFeePriorityPerGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetMaxFeePriorityPerGasOk() (*string, bool) {
	if o == nil || o.MaxFeePriorityPerGas == nil {
		return nil, false
	}
	return o.MaxFeePriorityPerGas, true
}

// HasMaxFeePriorityPerGas returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISE) HasMaxFeePriorityPerGas() bool {
	if o != nil && o.MaxFeePriorityPerGas != nil {
		return true
	}

	return false
}

// SetMaxFeePriorityPerGas gets a reference to the given string and assigns it to the MaxFeePriorityPerGas field.
func (o *DecodeRawTransactionHexRISE) SetMaxFeePriorityPerGas(v string) {
	o.MaxFeePriorityPerGas = &v
}

// GetNonce returns the Nonce field value
func (o *DecodeRawTransactionHexRISE) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *DecodeRawTransactionHexRISE) SetNonce(v int32) {
	o.Nonce = v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISE) GetR() string {
	if o == nil || o.R == nil {
		var ret string
		return ret
	}
	return *o.R
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetROk() (*string, bool) {
	if o == nil || o.R == nil {
		return nil, false
	}
	return o.R, true
}

// HasR returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISE) HasR() bool {
	if o != nil && o.R != nil {
		return true
	}

	return false
}

// SetR gets a reference to the given string and assigns it to the R field.
func (o *DecodeRawTransactionHexRISE) SetR(v string) {
	o.R = &v
}

// GetRecipient returns the Recipient field value
func (o *DecodeRawTransactionHexRISE) GetRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *DecodeRawTransactionHexRISE) SetRecipient(v string) {
	o.Recipient = v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISE) GetS() string {
	if o == nil || o.S == nil {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetSOk() (*string, bool) {
	if o == nil || o.S == nil {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISE) HasS() bool {
	if o != nil && o.S != nil {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *DecodeRawTransactionHexRISE) SetS(v string) {
	o.S = &v
}

// GetSender returns the Sender field value
func (o *DecodeRawTransactionHexRISE) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *DecodeRawTransactionHexRISE) SetSender(v string) {
	o.Sender = v
}

// GetType returns the Type field value
func (o *DecodeRawTransactionHexRISE) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DecodeRawTransactionHexRISE) SetType(v int32) {
	o.Type = v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISE) GetV() string {
	if o == nil || o.V == nil {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetVOk() (*string, bool) {
	if o == nil || o.V == nil {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISE) HasV() bool {
	if o != nil && o.V != nil {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *DecodeRawTransactionHexRISE) SetV(v string) {
	o.V = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DecodeRawTransactionHexRISE) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodeRawTransactionHexRISE) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DecodeRawTransactionHexRISE) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DecodeRawTransactionHexRISE) SetValue(v string) {
	o.Value = &v
}

func (o DecodeRawTransactionHexRISE) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApproximateFee != nil {
		toSerialize["approximateFee"] = o.ApproximateFee
	}
	if o.ApproximateMinimumRequiredFee != nil {
		toSerialize["approximateMinimumRequiredFee"] = o.ApproximateMinimumRequiredFee
	}
	if true {
		toSerialize["gasLimit"] = o.GasLimit
	}
	if o.GasPaidForData != nil {
		toSerialize["gasPaidForData"] = o.GasPaidForData
	}
	if o.GasPrice != nil {
		toSerialize["gasPrice"] = o.GasPrice
	}
	if o.InputData != nil {
		toSerialize["inputData"] = o.InputData
	}
	if o.MaxFeePerGas != nil {
		toSerialize["maxFeePerGas"] = o.MaxFeePerGas
	}
	if o.MaxFeePriorityPerGas != nil {
		toSerialize["maxFeePriorityPerGas"] = o.MaxFeePriorityPerGas
	}
	if true {
		toSerialize["nonce"] = o.Nonce
	}
	if o.R != nil {
		toSerialize["r"] = o.R
	}
	if true {
		toSerialize["recipient"] = o.Recipient
	}
	if o.S != nil {
		toSerialize["s"] = o.S
	}
	if true {
		toSerialize["sender"] = o.Sender
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.V != nil {
		toSerialize["v"] = o.V
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableDecodeRawTransactionHexRISE struct {
	value *DecodeRawTransactionHexRISE
	isSet bool
}

func (v NullableDecodeRawTransactionHexRISE) Get() *DecodeRawTransactionHexRISE {
	return v.value
}

func (v *NullableDecodeRawTransactionHexRISE) Set(val *DecodeRawTransactionHexRISE) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodeRawTransactionHexRISE) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodeRawTransactionHexRISE) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodeRawTransactionHexRISE(val *DecodeRawTransactionHexRISE) *NullableDecodeRawTransactionHexRISE {
	return &NullableDecodeRawTransactionHexRISE{value: val, isSet: true}
}

func (v NullableDecodeRawTransactionHexRISE) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodeRawTransactionHexRISE) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


