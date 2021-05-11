# GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptPubKey**](GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout

`func NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout(isSpent bool, scriptPubKey GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptPubKey, value string, ) *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout`

NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout instantiates a new GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVoutWithDefaults

`func NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVoutWithDefaults() *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout`

NewGetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVoutWithDefaults instantiates a new GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout) GetScriptPubKey() GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout) GetScriptPubKeyOk() (*GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout) SetScriptPubKey(v GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetTransactionDetailsByTransactionIDResponseItemBlockchainSpecificLitecoinVout) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


