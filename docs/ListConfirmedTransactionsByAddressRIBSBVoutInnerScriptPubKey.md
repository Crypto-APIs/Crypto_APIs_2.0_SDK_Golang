# ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** |  | 
**Asm** | **string** | Represents the assembly of the script public key of the address. | 
**Hex** | **string** | Represents the hex of the script public key of the address. | 
**ReqSigs** | Pointer to **int32** | Represents the required signatures. | [optional] 
**Type** | **string** | Represents the script type. | 

## Methods

### NewListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey

`func NewListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey(addresses []string, asm string, hex string, type_ string, ) *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey`

NewListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey instantiates a new ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKeyWithDefaults

`func NewListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKeyWithDefaults() *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey`

NewListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKeyWithDefaults instantiates a new ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetAsm

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) GetAsm() string`

GetAsm returns the Asm field if non-nil, zero value otherwise.

### GetAsmOk

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) GetAsmOk() (*string, bool)`

GetAsmOk returns a tuple with the Asm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsm

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) SetAsm(v string)`

SetAsm sets Asm field to given value.


### GetHex

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) GetHex() string`

GetHex returns the Hex field if non-nil, zero value otherwise.

### GetHexOk

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) GetHexOk() (*string, bool)`

GetHexOk returns a tuple with the Hex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHex

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) SetHex(v string)`

SetHex sets Hex field to given value.


### GetReqSigs

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) GetReqSigs() int32`

GetReqSigs returns the ReqSigs field if non-nil, zero value otherwise.

### GetReqSigsOk

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) GetReqSigsOk() (*int32, bool)`

GetReqSigsOk returns a tuple with the ReqSigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqSigs

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) SetReqSigs(v int32)`

SetReqSigs sets ReqSigs field to given value.

### HasReqSigs

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) HasReqSigs() bool`

HasReqSigs returns a boolean if a field has been set.

### GetType

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListConfirmedTransactionsByAddressRIBSBVoutInnerScriptPubKey) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


