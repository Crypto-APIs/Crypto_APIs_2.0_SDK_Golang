# EstimateGasLimitRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **string** | Represents an optional field to add additonal data. | [optional] 
**Amount** | **string** | Represents transactions&#39; amount. | 
**Recipient** | **string** | The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one recipient. | 
**Sender** | **string** | Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender. | 

## Methods

### NewEstimateGasLimitRBDataItem

`func NewEstimateGasLimitRBDataItem(amount string, recipient string, sender string, ) *EstimateGasLimitRBDataItem`

NewEstimateGasLimitRBDataItem instantiates a new EstimateGasLimitRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateGasLimitRBDataItemWithDefaults

`func NewEstimateGasLimitRBDataItemWithDefaults() *EstimateGasLimitRBDataItem`

NewEstimateGasLimitRBDataItemWithDefaults instantiates a new EstimateGasLimitRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *EstimateGasLimitRBDataItem) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *EstimateGasLimitRBDataItem) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *EstimateGasLimitRBDataItem) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *EstimateGasLimitRBDataItem) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAmount

`func (o *EstimateGasLimitRBDataItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EstimateGasLimitRBDataItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EstimateGasLimitRBDataItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetRecipient

`func (o *EstimateGasLimitRBDataItem) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *EstimateGasLimitRBDataItem) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *EstimateGasLimitRBDataItem) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetSender

`func (o *EstimateGasLimitRBDataItem) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *EstimateGasLimitRBDataItem) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *EstimateGasLimitRBDataItem) SetSender(v string)`

SetSender sets Sender field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


