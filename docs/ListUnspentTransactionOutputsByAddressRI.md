# ListUnspentTransactionOutputsByAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Represents the index position of the transaction in the block. | 
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain | 
**MinedInBlockHash** | **string** | Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm. | 
**MinedInBlockHeight** | **int32** | Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block. | 
**Recipients** | [**[]GetTransactionDetailsByTransactionIDRIRecipients**](GetTransactionDetailsByTransactionIDRIRecipients.md) | Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Senders** | [**[]ListUnspentTransactionOutputsByAddressRISenders**](ListUnspentTransactionOutputsByAddressRISenders.md) | Object Array representation of transaction senders | 
**Size** | **int32** | Represents the total size of this transaction | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionHash** | **string** | Represents the same as &#x60;transactionId&#x60; for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols &#x60;hash&#x60; is different from &#x60;transactionId&#x60; for SegWit transactions. | 
**TransactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be &#x60;transactionId&#x60; in UTXO-based protocols like Bitcoin, and transaction &#x60;hash&#x60; in Ethereum blockchain. | 
**Version** | **int32** | Represents the transaction version number. | 
**Vin** | [**[]ListUnspentTransactionOutputsByAddressRIVin**](ListUnspentTransactionOutputsByAddressRIVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListConfirmedTransactionsByAddressRIBSBVout**](ListConfirmedTransactionsByAddressRIBSBVout.md) | Represents the transaction outputs. | 
**Fee** | [**ListUnspentTransactionOutputsByAddressRIFee**](ListUnspentTransactionOutputsByAddressRIFee.md) |  | 
**BlockchainSpecific** | [**ListUnspentTransactionOutputsByAddressRIBlockchainSpecific**](ListUnspentTransactionOutputsByAddressRIBlockchainSpecific.md) |  | 

## Methods

### NewListUnspentTransactionOutputsByAddressRI

`func NewListUnspentTransactionOutputsByAddressRI(index int32, locktime int32, minedInBlockHash string, minedInBlockHeight int32, recipients []GetTransactionDetailsByTransactionIDRIRecipients, senders []ListUnspentTransactionOutputsByAddressRISenders, size int32, timestamp int32, transactionHash string, transactionId string, version int32, vin []ListUnspentTransactionOutputsByAddressRIVin, vout []ListConfirmedTransactionsByAddressRIBSBVout, fee ListUnspentTransactionOutputsByAddressRIFee, blockchainSpecific ListUnspentTransactionOutputsByAddressRIBlockchainSpecific, ) *ListUnspentTransactionOutputsByAddressRI`

NewListUnspentTransactionOutputsByAddressRI instantiates a new ListUnspentTransactionOutputsByAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnspentTransactionOutputsByAddressRIWithDefaults

`func NewListUnspentTransactionOutputsByAddressRIWithDefaults() *ListUnspentTransactionOutputsByAddressRI`

NewListUnspentTransactionOutputsByAddressRIWithDefaults instantiates a new ListUnspentTransactionOutputsByAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *ListUnspentTransactionOutputsByAddressRI) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ListUnspentTransactionOutputsByAddressRI) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetLocktime

`func (o *ListUnspentTransactionOutputsByAddressRI) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListUnspentTransactionOutputsByAddressRI) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetMinedInBlockHash

`func (o *ListUnspentTransactionOutputsByAddressRI) GetMinedInBlockHash() string`

GetMinedInBlockHash returns the MinedInBlockHash field if non-nil, zero value otherwise.

### GetMinedInBlockHashOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetMinedInBlockHashOk() (*string, bool)`

GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHash

`func (o *ListUnspentTransactionOutputsByAddressRI) SetMinedInBlockHash(v string)`

SetMinedInBlockHash sets MinedInBlockHash field to given value.


### GetMinedInBlockHeight

`func (o *ListUnspentTransactionOutputsByAddressRI) GetMinedInBlockHeight() int32`

GetMinedInBlockHeight returns the MinedInBlockHeight field if non-nil, zero value otherwise.

### GetMinedInBlockHeightOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetMinedInBlockHeightOk() (*int32, bool)`

GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinedInBlockHeight

`func (o *ListUnspentTransactionOutputsByAddressRI) SetMinedInBlockHeight(v int32)`

SetMinedInBlockHeight sets MinedInBlockHeight field to given value.


### GetRecipients

`func (o *ListUnspentTransactionOutputsByAddressRI) GetRecipients() []GetTransactionDetailsByTransactionIDRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetRecipientsOk() (*[]GetTransactionDetailsByTransactionIDRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListUnspentTransactionOutputsByAddressRI) SetRecipients(v []GetTransactionDetailsByTransactionIDRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *ListUnspentTransactionOutputsByAddressRI) GetSenders() []ListUnspentTransactionOutputsByAddressRISenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetSendersOk() (*[]ListUnspentTransactionOutputsByAddressRISenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *ListUnspentTransactionOutputsByAddressRI) SetSenders(v []ListUnspentTransactionOutputsByAddressRISenders)`

SetSenders sets Senders field to given value.


### GetSize

`func (o *ListUnspentTransactionOutputsByAddressRI) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListUnspentTransactionOutputsByAddressRI) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTimestamp

`func (o *ListUnspentTransactionOutputsByAddressRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListUnspentTransactionOutputsByAddressRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionHash

`func (o *ListUnspentTransactionOutputsByAddressRI) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ListUnspentTransactionOutputsByAddressRI) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetTransactionId

`func (o *ListUnspentTransactionOutputsByAddressRI) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ListUnspentTransactionOutputsByAddressRI) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetVersion

`func (o *ListUnspentTransactionOutputsByAddressRI) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListUnspentTransactionOutputsByAddressRI) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListUnspentTransactionOutputsByAddressRI) GetVin() []ListUnspentTransactionOutputsByAddressRIVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetVinOk() (*[]ListUnspentTransactionOutputsByAddressRIVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListUnspentTransactionOutputsByAddressRI) SetVin(v []ListUnspentTransactionOutputsByAddressRIVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListUnspentTransactionOutputsByAddressRI) GetVout() []ListConfirmedTransactionsByAddressRIBSBVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetVoutOk() (*[]ListConfirmedTransactionsByAddressRIBSBVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListUnspentTransactionOutputsByAddressRI) SetVout(v []ListConfirmedTransactionsByAddressRIBSBVout)`

SetVout sets Vout field to given value.


### GetFee

`func (o *ListUnspentTransactionOutputsByAddressRI) GetFee() ListUnspentTransactionOutputsByAddressRIFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetFeeOk() (*ListUnspentTransactionOutputsByAddressRIFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ListUnspentTransactionOutputsByAddressRI) SetFee(v ListUnspentTransactionOutputsByAddressRIFee)`

SetFee sets Fee field to given value.


### GetBlockchainSpecific

`func (o *ListUnspentTransactionOutputsByAddressRI) GetBlockchainSpecific() ListUnspentTransactionOutputsByAddressRIBlockchainSpecific`

GetBlockchainSpecific returns the BlockchainSpecific field if non-nil, zero value otherwise.

### GetBlockchainSpecificOk

`func (o *ListUnspentTransactionOutputsByAddressRI) GetBlockchainSpecificOk() (*ListUnspentTransactionOutputsByAddressRIBlockchainSpecific, bool)`

GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainSpecific

`func (o *ListUnspentTransactionOutputsByAddressRI) SetBlockchainSpecific(v ListUnspentTransactionOutputsByAddressRIBlockchainSpecific)`

SetBlockchainSpecific sets BlockchainSpecific field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


