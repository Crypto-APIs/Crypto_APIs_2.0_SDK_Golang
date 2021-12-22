# GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anchor** | **string** | Defines a Merkle tree root of a note commitment tree which uniquely identifies a note commitment tree state given the assumed security properties of the Merkle tree’s  hash function. | 
**Cv** | **string** | Defines a value commitment to the value of the input note. | 
**Nullifier** | **string** | Represents a sequence of nullifiers of the input notes. | 
**Proof** | **string** | Represents the proof. | 
**Rk** | **string** | Represents the randomized validating key for spendAuthSig. | 
**SpendAuthSig** | **string** | Used to prove knowledge of the spending key authorizing spending of an input note. | 

## Methods

### NewGetTransactionDetailsByTransactionIDRIBSZVShieldedSpend

`func NewGetTransactionDetailsByTransactionIDRIBSZVShieldedSpend(anchor string, cv string, nullifier string, proof string, rk string, spendAuthSig string, ) *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend`

NewGetTransactionDetailsByTransactionIDRIBSZVShieldedSpend instantiates a new GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDRIBSZVShieldedSpendWithDefaults

`func NewGetTransactionDetailsByTransactionIDRIBSZVShieldedSpendWithDefaults() *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend`

NewGetTransactionDetailsByTransactionIDRIBSZVShieldedSpendWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnchor

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) GetAnchor() string`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) GetAnchorOk() (*string, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) SetAnchor(v string)`

SetAnchor sets Anchor field to given value.


### GetCv

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) GetCv() string`

GetCv returns the Cv field if non-nil, zero value otherwise.

### GetCvOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) GetCvOk() (*string, bool)`

GetCvOk returns a tuple with the Cv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCv

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) SetCv(v string)`

SetCv sets Cv field to given value.


### GetNullifier

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) GetNullifier() string`

GetNullifier returns the Nullifier field if non-nil, zero value otherwise.

### GetNullifierOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) GetNullifierOk() (*string, bool)`

GetNullifierOk returns a tuple with the Nullifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullifier

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) SetNullifier(v string)`

SetNullifier sets Nullifier field to given value.


### GetProof

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) GetProof() string`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) GetProofOk() (*string, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) SetProof(v string)`

SetProof sets Proof field to given value.


### GetRk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) GetRk() string`

GetRk returns the Rk field if non-nil, zero value otherwise.

### GetRkOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) GetRkOk() (*string, bool)`

GetRkOk returns a tuple with the Rk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) SetRk(v string)`

SetRk sets Rk field to given value.


### GetSpendAuthSig

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) GetSpendAuthSig() string`

GetSpendAuthSig returns the SpendAuthSig field if non-nil, zero value otherwise.

### GetSpendAuthSigOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) GetSpendAuthSigOk() (*string, bool)`

GetSpendAuthSigOk returns a tuple with the SpendAuthSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendAuthSig

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedSpend) SetSpendAuthSig(v string)`

SetSpendAuthSig sets SpendAuthSig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


