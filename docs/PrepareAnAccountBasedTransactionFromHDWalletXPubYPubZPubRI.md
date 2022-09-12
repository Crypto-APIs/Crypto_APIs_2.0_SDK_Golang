# PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Representation of the amount of the transaction | 
**DataHex** | **string** | Representation of the data in hex value | 
**DerivationIndex** | **int64** | Representation of the derivation index of the xpub address | 
**Fee** | [**PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRIFee**](PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRIFee.md) |  | 
**Nonce** | **string** | Represents the sequential running number for an address, starting from 0 for the first transaction. E.g., if the nonce of a transaction is 10, it would be the 11th transaction sent from the sender&#39;s address. | 
**Recipient** | **string** | Represents a recipient addresses. In account-based protocols like Ethereum there is only one address in this list. | 
**Sender** | **string** | Represents a sender address. In account-based protocols like Ethereum there is only one address in this list. | 
**SigHash** | **string** | Representation of the hash that should be signed. | 
**TransactionType** | **string** | Representation of the transaction type | 

## Methods

### NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI

`func NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI(amount string, dataHex string, derivationIndex int64, fee PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRIFee, nonce string, recipient string, sender string, sigHash string, transactionType string, ) *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI`

NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI instantiates a new PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRIWithDefaults

`func NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRIWithDefaults() *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI`

NewPrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRIWithDefaults instantiates a new PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDataHex

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetDataHex() string`

GetDataHex returns the DataHex field if non-nil, zero value otherwise.

### GetDataHexOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetDataHexOk() (*string, bool)`

GetDataHexOk returns a tuple with the DataHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataHex

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) SetDataHex(v string)`

SetDataHex sets DataHex field to given value.


### GetDerivationIndex

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetDerivationIndex() int64`

GetDerivationIndex returns the DerivationIndex field if non-nil, zero value otherwise.

### GetDerivationIndexOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetDerivationIndexOk() (*int64, bool)`

GetDerivationIndexOk returns a tuple with the DerivationIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivationIndex

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) SetDerivationIndex(v int64)`

SetDerivationIndex sets DerivationIndex field to given value.


### GetFee

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetFee() PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetFeeOk() (*PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) SetFee(v PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRIFee)`

SetFee sets Fee field to given value.


### GetNonce

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetRecipient

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetSender

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) SetSender(v string)`

SetSender sets Sender field to given value.


### GetSigHash

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetSigHash() string`

GetSigHash returns the SigHash field if non-nil, zero value otherwise.

### GetSigHashOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetSigHashOk() (*string, bool)`

GetSigHashOk returns a tuple with the SigHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigHash

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) SetSigHash(v string)`

SetSigHash sets SigHash field to given value.


### GetTransactionType

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *PrepareAnAccountBasedTransactionFromHDWalletXPubYPubZPubRI) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


