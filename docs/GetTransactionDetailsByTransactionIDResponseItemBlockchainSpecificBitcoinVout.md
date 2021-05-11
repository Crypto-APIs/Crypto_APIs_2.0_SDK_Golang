# GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptPubKey**](GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout

`func NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptPubKey, value string, ) *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout`

NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout instantiates a new GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVoutWithDefaults

`func NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVoutWithDefaults() *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout`

NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVoutWithDefaults instantiates a new GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout) GetScriptPubKey() GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout) SetScriptPubKey(v GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificBitcoinVout) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


