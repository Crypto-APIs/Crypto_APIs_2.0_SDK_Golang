# EstimateTokenGasLimitRBDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Represents transactions&#39; amount. | 
**Contract** | **string** | Defines the specific token identifier.  For Ethereum-based transactions it is the contract. | 
**ContractType** | **string** | Represents the ERC contract type. It can be ERC20 or ERC721 | 
**Recipient** | **string** | The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one recipient. | 
**Sender** | **string** | Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender. | 

## Methods

### NewEstimateTokenGasLimitRBDataItem

`func NewEstimateTokenGasLimitRBDataItem(amount string, contract string, contractType string, recipient string, sender string, ) *EstimateTokenGasLimitRBDataItem`

NewEstimateTokenGasLimitRBDataItem instantiates a new EstimateTokenGasLimitRBDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateTokenGasLimitRBDataItemWithDefaults

`func NewEstimateTokenGasLimitRBDataItemWithDefaults() *EstimateTokenGasLimitRBDataItem`

NewEstimateTokenGasLimitRBDataItemWithDefaults instantiates a new EstimateTokenGasLimitRBDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EstimateTokenGasLimitRBDataItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EstimateTokenGasLimitRBDataItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EstimateTokenGasLimitRBDataItem) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetContract

`func (o *EstimateTokenGasLimitRBDataItem) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *EstimateTokenGasLimitRBDataItem) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *EstimateTokenGasLimitRBDataItem) SetContract(v string)`

SetContract sets Contract field to given value.


### GetContractType

`func (o *EstimateTokenGasLimitRBDataItem) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *EstimateTokenGasLimitRBDataItem) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *EstimateTokenGasLimitRBDataItem) SetContractType(v string)`

SetContractType sets ContractType field to given value.


### GetRecipient

`func (o *EstimateTokenGasLimitRBDataItem) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *EstimateTokenGasLimitRBDataItem) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *EstimateTokenGasLimitRBDataItem) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetSender

`func (o *EstimateTokenGasLimitRBDataItem) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *EstimateTokenGasLimitRBDataItem) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *EstimateTokenGasLimitRBDataItem) SetSender(v string)`

SetSender sets Sender field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


