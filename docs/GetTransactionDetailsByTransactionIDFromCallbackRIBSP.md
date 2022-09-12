# GetTransactionDetailsByTransactionIDFromCallbackRIBSP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Representation of the amount value. | 
**Contract** | **string** | Represents the specific transaction contract. | 
**Gas** | **string** | Represents the price offered to the miner to purchase this amount of gas. | 
**GasPrice** | **string** | Represents the price offered to the miner to purchase this amount of gas. | 
**GasUsed** | **string** | Represents the exact unit of gas that was used for the transaction. | 
**Input** | **string** | Represents additional information that is required for the transaction. | 
**Nonce** | **int32** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**Recipients** | **string** | Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Senders** | **string** | Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**TransactionStatus** | **string** | Represents the status of this transaction. | 
**Txid** | **string** | Represents the unique identifier of a transaction, i.e. it could be transactionId in UTXO-based protocols like Bitcoin, and transaction hash in Ethereum blockchain. | 

## Methods

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSP

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSP(amount string, contract string, gas string, gasPrice string, gasUsed string, input string, nonce int32, recipients string, senders string, transactionStatus string, txid string, ) *GetTransactionDetailsByTransactionIDFromCallbackRIBSP`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSP instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSPWithDefaults

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSPWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSP`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSPWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetContract

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetContract(v string)`

SetContract sets Contract field to given value.


### GetGas

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetGas() string`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetGasOk() (*string, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetGas(v string)`

SetGas sets Gas field to given value.


### GetGasPrice

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetGasUsed

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetInput

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetInput(v string)`

SetInput sets Input field to given value.


### GetNonce

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetRecipients

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetRecipients() string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetRecipientsOk() (*string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetRecipients(v string)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetSenders() string`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetSendersOk() (*string, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetSenders(v string)`

SetSenders sets Senders field to given value.


### GetTransactionStatus

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.


### GetTxid

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSP) SetTxid(v string)`

SetTxid sets Txid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


