# GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Asm** | **string** | Represents the assembly of the script public key of the address. | 
**Hex** | **string** | Represents the hex of the script public key of the address. | 
**ReqSigs** | Pointer to **int32** | Represents the required signatures. | [optional] 
**Type** | **string** | Represents the script type. | 

## Methods

### NewGetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey

`func NewGetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey(addresses []string, asm string, hex string, type_ string, ) *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey`

NewGetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey instantiates a new GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKeyWithDefaults

`func NewGetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKeyWithDefaults() *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey`

NewGetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKeyWithDefaults instantiates a new GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetAsm

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) GetAsm() string`

GetAsm returns the Asm field if non-nil, zero value otherwise.

### GetAsmOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) GetAsmOk() (*string, bool)`

GetAsmOk returns a tuple with the Asm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsm

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) SetAsm(v string)`

SetAsm sets Asm field to given value.


### GetHex

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) GetHex() string`

GetHex returns the Hex field if non-nil, zero value otherwise.

### GetHexOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) GetHexOk() (*string, bool)`

GetHexOk returns a tuple with the Hex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHex

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) SetHex(v string)`

SetHex sets Hex field to given value.


### GetReqSigs

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) GetReqSigs() int32`

GetReqSigs returns the ReqSigs field if non-nil, zero value otherwise.

### GetReqSigsOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) GetReqSigsOk() (*int32, bool)`

GetReqSigsOk returns a tuple with the ReqSigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqSigs

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) SetReqSigs(v int32)`

SetReqSigs sets ReqSigs field to given value.

### HasReqSigs

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) HasReqSigs() bool`

HasReqSigs returns a boolean if a field has been set.

### GetType

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTransactionDetailsByTransactionIDRIBSBVoutInnerScriptPubKey) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


