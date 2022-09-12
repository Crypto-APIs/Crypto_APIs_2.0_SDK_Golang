# ListConfirmedTokensTransfersByAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractAddress** | **string** | Represents the contract address of the token, which controls its logic. It is not the address that holds the tokens. | 
**MinedInBlockHeight** | **int64** | Defines the block height in which this transaction was confirmed/mined. | 
**RecipientAddress** | **string** | Defines the address to which the recipient receives the transferred tokens. | 
**SenderAddress** | **string** | Defines the address from which the sender transfers tokens. | 
**TokenDecimals** | **int32** | Defines the decimals of the token, i.e. the number of digits that come after the decimal coma of the token. | 
**TokenId** | Pointer to **string** | Represents the unique token identifier. | [optional] 
**TokenName** | **string** | Defines the token&#39;s name as a string. | 
**TokenSymbol** | **string** | Defines the token symbol by which the token contract is known. It is usually 3-4 characters in length. | 
**TokenType** | **string** | Defines the specific token type. | 
**TokensAmount** | Pointer to **string** | Defines the token amount of the transfer. | [optional] 
**TransactionHash** | **string** | Represents the hash of the transaction, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**TransactionTimestamp** | **int32** | Defines the specific time/date when the transaction was created in Unix Timestamp. | 
**TransactionFee** | [**ListTokensTransfersByTransactionHashRITransactionFee**](ListTokensTransfersByTransactionHashRITransactionFee.md) |  | 

## Methods

### NewListConfirmedTokensTransfersByAddressRI

`func NewListConfirmedTokensTransfersByAddressRI(contractAddress string, minedInBlockHeight int64, recipientAddress string, senderAddress string, tokenDecimals int32, tokenName string, tokenSymbol string, tokenType string, transactionHash string, transactionTimestamp int32, transactionFee ListTokensTransfersByTransactionHashRITransactionFee, ) *ListConfirmedTokensTransfersByAddressRI`

NewListConfirmedTokensTransfersByAddressRI instantiates a new ListConfirmedTokensTransfersByAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTokensTransfersByAddressRIWithDefaults

`func NewListConfirmedTokensTransfersByAddressRIWithDefaults() *ListConfirmedTokensTransfersByAddressRI`

NewListConfirmedTokensTransfersByAddressRIWithDefaults instantiates a new ListConfirmedTokensTransfersByAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractAddress

`func (o *ListConfirmedTokensTransfersByAddressRI) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ListConfirmedTokensTransfersByAddressRI) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ListConfirmedTokensTransfersByAddressRI) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetMinedInBlockHeight

`func (o *ListConfirmedTokensTransfersByAddressRI) GetMinedInBlockHeight() int64`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *ListConfirmedTokensTransfersByAddressRI) GetMinedInBlockHeightOk() (*int64, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *ListConfirmedTokensTransfersByAddressRI) SetMinedInBlockHeight(v int64)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetRecipientAddress

`func (o *ListConfirmedTokensTransfersByAddressRI) GetRecipientAddress() string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *ListConfirmedTokensTransfersByAddressRI) GetRecipientAddressOk() (*string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *ListConfirmedTokensTransfersByAddressRI) SetRecipientAddress(v string)`

SetRecipientAddress sets RecipientAddress field to given value.


### GetSenderAddress

`func (o *ListConfirmedTokensTransfersByAddressRI) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *ListConfirmedTokensTransfersByAddressRI) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *ListConfirmedTokensTransfersByAddressRI) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetTokenDecimals

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenDecimals() int32`

GetTokenDecimals returns the TokenDecimals field if non-nil, zero value otherwise.

### GetTokenDecimalsOk

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenDecimalsOk() (*int32, bool)`

GetTokenDecimalsOk returns a tuple with the TokenDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDecimals

`func (o *ListConfirmedTokensTransfersByAddressRI) SetTokenDecimals(v int32)`

SetTokenDecimals sets TokenDecimals field to given value.


### GetTokenId

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ListConfirmedTokensTransfersByAddressRI) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ListConfirmedTokensTransfersByAddressRI) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTokenName

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *ListConfirmedTokensTransfersByAddressRI) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetTokenSymbol

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenSymbol() string`

GetTokenSymbol returns the TokenSymbol field if non-nil, zero value otherwise.

### GetTokenSymbolOk

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenSymbolOk() (*string, bool)`

GetTokenSymbolOk returns a tuple with the TokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSymbol

`func (o *ListConfirmedTokensTransfersByAddressRI) SetTokenSymbol(v string)`

SetTokenSymbol sets TokenSymbol field to given value.


### GetTokenType

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *ListConfirmedTokensTransfersByAddressRI) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetTokensAmount

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTokensAmount() string`

GetTokensAmount returns the TokensAmount field if non-nil, zero value otherwise.

### GetTokensAmountOk

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTokensAmountOk() (*string, bool)`

GetTokensAmountOk returns a tuple with the TokensAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensAmount

`func (o *ListConfirmedTokensTransfersByAddressRI) SetTokensAmount(v string)`

SetTokensAmount sets TokensAmount field to given value.

### HasTokensAmount

`func (o *ListConfirmedTokensTransfersByAddressRI) HasTokensAmount() bool`

HasTokensAmount returns a boolean if a field has been set.

### GetTransactionHash

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ListConfirmedTokensTransfersByAddressRI) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetTransactionTimestamp

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTransactionTimestamp() int32`

GetTransactionTimestamp returns the TransactionTimestamp field if non-nil, zero value otherwise.

### GetTransactionTimestampOk

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTransactionTimestampOk() (*int32, bool)`

GetTransactionTimestampOk returns a tuple with the TransactionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTimestamp

`func (o *ListConfirmedTokensTransfersByAddressRI) SetTransactionTimestamp(v int32)`

SetTransactionTimestamp sets TransactionTimestamp field to given value.


### GetTransactionFee

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTransactionFee() ListTokensTransfersByTransactionHashRITransactionFee`

GetTransactionFee returns the TransactionFee field if non-nil, zero value otherwise.

### GetTransactionFeeOk

`func (o *ListConfirmedTokensTransfersByAddressRI) GetTransactionFeeOk() (*ListTokensTransfersByTransactionHashRITransactionFee, bool)`

GetTransactionFeeOk returns a tuple with the TransactionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionFee

`func (o *ListConfirmedTokensTransfersByAddressRI) SetTransactionFee(v ListTokensTransfersByTransactionHashRITransactionFee)`

SetTransactionFee sets TransactionFee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


