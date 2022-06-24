# GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cmu** | **string** | Represents the 𝑢-coordinate of the note commitment for the output note. | 
**Cv** | **string** | Defines a value commitment to the value of the input note. | 
**EncCipherText** | **string** | Represents a ciphertext component for the encrypted output note. | 
**EphemeralKey** | **string** | Represents an encoding of an ephemeral Jubjub public key. | 
**OutCipherText** | **string** | Represents a ciphertext component that allows the holder of the outgoing cipher key to recover the diversified transmission key pkd and ephemeral private key esk, hence the entire note plaintext. | 
**Proof** | **string** | Represents the proof. | 

## Methods

### NewGetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner

`func NewGetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner(cmu string, cv string, encCipherText string, ephemeralKey string, outCipherText string, proof string, ) *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner`

NewGetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner instantiates a new GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInnerWithDefaults

`func NewGetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInnerWithDefaults() *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner`

NewGetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInnerWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCmu

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) GetCmu() string`

GetCmu returns the Cmu field if non-nil, zero value otherwise.

### GetCmuOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) GetCmuOk() (*string, bool)`

GetCmuOk returns a tuple with the Cmu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmu

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) SetCmu(v string)`

SetCmu sets Cmu field to given value.


### GetCv

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) GetCv() string`

GetCv returns the Cv field if non-nil, zero value otherwise.

### GetCvOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) GetCvOk() (*string, bool)`

GetCvOk returns a tuple with the Cv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCv

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) SetCv(v string)`

SetCv sets Cv field to given value.


### GetEncCipherText

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) GetEncCipherText() string`

GetEncCipherText returns the EncCipherText field if non-nil, zero value otherwise.

### GetEncCipherTextOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) GetEncCipherTextOk() (*string, bool)`

GetEncCipherTextOk returns a tuple with the EncCipherText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncCipherText

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) SetEncCipherText(v string)`

SetEncCipherText sets EncCipherText field to given value.


### GetEphemeralKey

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) GetEphemeralKey() string`

GetEphemeralKey returns the EphemeralKey field if non-nil, zero value otherwise.

### GetEphemeralKeyOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) GetEphemeralKeyOk() (*string, bool)`

GetEphemeralKeyOk returns a tuple with the EphemeralKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralKey

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) SetEphemeralKey(v string)`

SetEphemeralKey sets EphemeralKey field to given value.


### GetOutCipherText

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) GetOutCipherText() string`

GetOutCipherText returns the OutCipherText field if non-nil, zero value otherwise.

### GetOutCipherTextOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) GetOutCipherTextOk() (*string, bool)`

GetOutCipherTextOk returns a tuple with the OutCipherText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutCipherText

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) SetOutCipherText(v string)`

SetOutCipherText sets OutCipherText field to given value.


### GetProof

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) GetProof() string`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) GetProofOk() (*string, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *GetTransactionDetailsByTransactionIDRIBSZVShieldedOutputInner) SetProof(v string)`

SetProof sets Proof field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


