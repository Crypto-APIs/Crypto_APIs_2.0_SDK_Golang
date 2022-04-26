# ListTokensTransfersByTransactionHashRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractAddress** | **string** | Represents the contract address of the token, which controls its logic. It is not the address that holds the tokens. | 
**MinedInBlockHeight** | **int32** | Defines the block height in which this transaction was confirmed/mined. | 
**RecipientAddress** | **string** | Defines the address to which the recipient receives the transferred tokens. | 
**SenderAddress** | **string** | Defines the address from which the sender transfers tokens. | 
**TokenDecimals** | **int32** | Defines the decimals of the token, i.e. the number of digits that come after the decimal coma of the token. | 
**TokenName** | **string** | Defines the token&#39;s name as a string. | 
**TokenSymbol** | **string** | Defines the token symbol by which the token contract is known. It is usually 3-4 characters in length. | 
**TokenType** | **string** | Defines the specific token type. | 
**TokensAmount** | **string** | Defines the token amount of the transfer. | 
**TransactionHash** | **string** | Represents the hash of the transaction, which is its unique identifier. It represents a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**TransactionTimestamp** | **int32** | Defines the specific time/date when the transaction was created in Unix Timestamp. | 
**TransactionFee** | [**ListTokensTransfersByTransactionHashRITransactionFee**](ListTokensTransfersByTransactionHashRITransactionFee.md) |  | 

## Methods

### NewListTokensTransfersByTransactionHashRI

`func NewListTokensTransfersByTransactionHashRI(contractAddress string, minedInBlockHeight int32, recipientAddress string, senderAddress string, tokenDecimals int32, tokenName string, tokenSymbol string, tokenType string, tokensAmount string, transactionHash string, transactionTimestamp int32, transactionFee ListTokensTransfersByTransactionHashRITransactionFee, ) *ListTokensTransfersByTransactionHashRI`

NewListTokensTransfersByTransactionHashRI instantiates a new ListTokensTransfersByTransactionHashRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTokensTransfersByTransactionHashRIWithDefaults

`func NewListTokensTransfersByTransactionHashRIWithDefaults() *ListTokensTransfersByTransactionHashRI`

NewListTokensTransfersByTransactionHashRIWithDefaults instantiates a new ListTokensTransfersByTransactionHashRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractAddress

`func (o *ListTokensTransfersByTransactionHashRI) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ListTokensTransfersByTransactionHashRI) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ListTokensTransfersByTransactionHashRI) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetMinedInBlockHeight

`func (o *ListTokensTransfersByTransactionHashRI) GetMinedInBlockHeight() int32`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *ListTokensTransfersByTransactionHashRI) GetMinedInBlockHeightOk() (*int32, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *ListTokensTransfersByTransactionHashRI) SetMinedInBlockHeight(v int32)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetRecipientAddress

`func (o *ListTokensTransfersByTransactionHashRI) GetRecipientAddress() string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *ListTokensTransfersByTransactionHashRI) GetRecipientAddressOk() (*string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *ListTokensTransfersByTransactionHashRI) SetRecipientAddress(v string)`

SetRecipientAddress sets RecipientAddress field to given value.


### GetSenderAddress

`func (o *ListTokensTransfersByTransactionHashRI) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *ListTokensTransfersByTransactionHashRI) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *ListTokensTransfersByTransactionHashRI) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetTokenDecimals

`func (o *ListTokensTransfersByTransactionHashRI) GetTokenDecimals() int32`

GetTokenDecimals returns the TokenDecimals field if non-nil, zero value otherwise.

### GetTokenDecimalsOk

`func (o *ListTokensTransfersByTransactionHashRI) GetTokenDecimalsOk() (*int32, bool)`

GetTokenDecimalsOk returns a tuple with the TokenDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDecimals

`func (o *ListTokensTransfersByTransactionHashRI) SetTokenDecimals(v int32)`

SetTokenDecimals sets TokenDecimals field to given value.


### GetTokenName

`func (o *ListTokensTransfersByTransactionHashRI) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *ListTokensTransfersByTransactionHashRI) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *ListTokensTransfersByTransactionHashRI) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetTokenSymbol

`func (o *ListTokensTransfersByTransactionHashRI) GetTokenSymbol() string`

GetTokenSymbol returns the TokenSymbol field if non-nil, zero value otherwise.

### GetTokenSymbolOk

`func (o *ListTokensTransfersByTransactionHashRI) GetTokenSymbolOk() (*string, bool)`

GetTokenSymbolOk returns a tuple with the TokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSymbol

`func (o *ListTokensTransfersByTransactionHashRI) SetTokenSymbol(v string)`

SetTokenSymbol sets TokenSymbol field to given value.


### GetTokenType

`func (o *ListTokensTransfersByTransactionHashRI) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *ListTokensTransfersByTransactionHashRI) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *ListTokensTransfersByTransactionHashRI) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetTokensAmount

`func (o *ListTokensTransfersByTransactionHashRI) GetTokensAmount() string`

GetTokensAmount returns the TokensAmount field if non-nil, zero value otherwise.

### GetTokensAmountOk

`func (o *ListTokensTransfersByTransactionHashRI) GetTokensAmountOk() (*string, bool)`

GetTokensAmountOk returns a tuple with the TokensAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensAmount

`func (o *ListTokensTransfersByTransactionHashRI) SetTokensAmount(v string)`

SetTokensAmount sets TokensAmount field to given value.


### GetTransactionHash

`func (o *ListTokensTransfersByTransactionHashRI) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ListTokensTransfersByTransactionHashRI) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ListTokensTransfersByTransactionHashRI) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetTransactionTimestamp

`func (o *ListTokensTransfersByTransactionHashRI) GetTransactionTimestamp() int32`

GetTransactionTimestamp returns the TransactionTimestamp field if non-nil, zero value otherwise.

### GetTransactionTimestampOk

`func (o *ListTokensTransfersByTransactionHashRI) GetTransactionTimestampOk() (*int32, bool)`

GetTransactionTimestampOk returns a tuple with the TransactionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTimestamp

`func (o *ListTokensTransfersByTransactionHashRI) SetTransactionTimestamp(v int32)`

SetTransactionTimestamp sets TransactionTimestamp field to given value.


### GetTransactionFee

`func (o *ListTokensTransfersByTransactionHashRI) GetTransactionFee() ListTokensTransfersByTransactionHashRITransactionFee`

GetTransactionFee returns the TransactionFee field if non-nil, zero value otherwise.

### GetTransactionFeeOk

`func (o *ListTokensTransfersByTransactionHashRI) GetTransactionFeeOk() (*ListTokensTransfersByTransactionHashRITransactionFee, bool)`

GetTransactionFeeOk returns a tuple with the TransactionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionFee

`func (o *ListTokensTransfersByTransactionHashRI) SetTransactionFee(v ListTokensTransfersByTransactionHashRITransactionFee)`

SetTransactionFee sets TransactionFee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


