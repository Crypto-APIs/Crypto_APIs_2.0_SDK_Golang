# ListHDWalletXPubYPubZPubUTXOsRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents the public address, which is a compressed and shortened form of a public key. | 
**AddressPath** | **string** | Defines a data which tells a Hierarchical Deterministic wallet how to derive a specific key within a tree of keys. | 
**Amount** | **string** | Represents the UTXO amount value. | 
**Derivation** | **string** | The way how the HD walled derives, for example when the type is ACCOUNT, it derives change and receive addresses while when the type is BIP32 it derives directly. | 
**Index** | **int32** | Represents the output index. It refers to the UTXO sequence in the transaction outputs (vout). | 
**IsAvailable** | **bool** | Represents if the UTXO has been used from another unconfirmed transaction. If it is - the value will be \&quot;false\&quot;. | 
**IsConfirmed** | **bool** | Represents the state of the transaction whether it is confirmed or not confirmed. | 
**ReferenceId** | **string** | Represents the reference id of the record. It may be used for the startingAfter pagination attribute. | 
**TransactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be transactionId in UTXO-based protocols like Bitcoin, and transaction hash in Ethereum blockchain. | 

## Methods

### NewListHDWalletXPubYPubZPubUTXOsRI

`func NewListHDWalletXPubYPubZPubUTXOsRI(address string, addressPath string, amount string, derivation string, index int32, isAvailable bool, isConfirmed bool, referenceId string, transactionId string, ) *ListHDWalletXPubYPubZPubUTXOsRI`

NewListHDWalletXPubYPubZPubUTXOsRI instantiates a new ListHDWalletXPubYPubZPubUTXOsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHDWalletXPubYPubZPubUTXOsRIWithDefaults

`func NewListHDWalletXPubYPubZPubUTXOsRIWithDefaults() *ListHDWalletXPubYPubZPubUTXOsRI`

NewListHDWalletXPubYPubZPubUTXOsRIWithDefaults instantiates a new ListHDWalletXPubYPubZPubUTXOsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddressPath

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetAddressPath() string`

GetAddressPath returns the AddressPath field if non-nil, zero value otherwise.

### GetAddressPathOk

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetAddressPathOk() (*string, bool)`

GetAddressPathOk returns a tuple with the AddressPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPath

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) SetAddressPath(v string)`

SetAddressPath sets AddressPath field to given value.


### GetAmount

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDerivation

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetDerivation() string`

GetDerivation returns the Derivation field if non-nil, zero value otherwise.

### GetDerivationOk

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetDerivationOk() (*string, bool)`

GetDerivationOk returns a tuple with the Derivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivation

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) SetDerivation(v string)`

SetDerivation sets Derivation field to given value.


### GetIndex

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetIsAvailable

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetIsAvailable() bool`

GetIsAvailable returns the IsAvailable field if non-nil, zero value otherwise.

### GetIsAvailableOk

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetIsAvailableOk() (*bool, bool)`

GetIsAvailableOk returns a tuple with the IsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailable

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) SetIsAvailable(v bool)`

SetIsAvailable sets IsAvailable field to given value.


### GetIsConfirmed

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetIsConfirmed() bool`

GetIsConfirmed returns the IsConfirmed field if non-nil, zero value otherwise.

### GetIsConfirmedOk

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetIsConfirmedOk() (*bool, bool)`

GetIsConfirmedOk returns a tuple with the IsConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfirmed

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) SetIsConfirmed(v bool)`

SetIsConfirmed sets IsConfirmed field to given value.


### GetReferenceId

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetTransactionId

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ListHDWalletXPubYPubZPubUTXOsRI) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


