# ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSpent** | **bool** | Defines whether the output is spent or not. | 
**ScriptPubKey** | [**ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinScriptPubKey**](ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinScriptPubKey.md) |  | 
**Value** | **string** | Represents the sent/received amount. | 

## Methods

### NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout

`func NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout(isSpent bool, scriptPubKey ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinScriptPubKey, value string, ) *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout`

NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVoutWithDefaults

`func NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVoutWithDefaults() *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout`

NewListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVoutWithDefaults instantiates a new ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSpent

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) GetIsSpent() bool`

GetIsSpent returns the IsSpent field if non-nil, zero value otherwise.

### GetIsSpentOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) GetIsSpentOk() (*bool, bool)`

GetIsSpentOk returns a tuple with the IsSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpent

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) SetIsSpent(v bool)`

SetIsSpent sets IsSpent field to given value.


### GetScriptPubKey

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) GetScriptPubKey() ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinScriptPubKey`

GetScriptPubKey returns the ScriptPubKey field if non-nil, zero value otherwise.

### GetScriptPubKeyOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) GetScriptPubKeyOk() (*ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinScriptPubKey, bool)`

GetScriptPubKeyOk returns a tuple with the ScriptPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPubKey

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) SetScriptPubKey(v ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinScriptPubKey)`

SetScriptPubKey sets ScriptPubKey field to given value.


### GetValue

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListTransactionsByBlockHashResponseItemBlockchainSpecificLitecoinVout) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


