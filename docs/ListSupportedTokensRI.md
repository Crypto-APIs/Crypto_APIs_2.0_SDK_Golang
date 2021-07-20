# ListSupportedTokensRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decimals** | **int32** | Defines the token&#39;s decimal number or all of its points after the zero. | 
**Identifier** | **string** | Represents a unique identifier for the specific blockchain and network, e.g. smart contract address, property ID, etc. | 
**Name** | **string** | Defines the token name. | 
**Symbol** | **string** | Defines the token&#39;s unique symbol in the Crypto APIs listings. | 
**Type** | **string** | Represents the token&#39;s type representation, e.g. ERC-20, Omni, etc. | 

## Methods

### NewListSupportedTokensRI

`func NewListSupportedTokensRI(decimals int32, identifier string, name string, symbol string, type_ string, ) *ListSupportedTokensRI`

NewListSupportedTokensRI instantiates a new ListSupportedTokensRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSupportedTokensRIWithDefaults

`func NewListSupportedTokensRIWithDefaults() *ListSupportedTokensRI`

NewListSupportedTokensRIWithDefaults instantiates a new ListSupportedTokensRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecimals

`func (o *ListSupportedTokensRI) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *ListSupportedTokensRI) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *ListSupportedTokensRI) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetIdentifier

`func (o *ListSupportedTokensRI) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ListSupportedTokensRI) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ListSupportedTokensRI) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *ListSupportedTokensRI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSupportedTokensRI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSupportedTokensRI) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *ListSupportedTokensRI) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ListSupportedTokensRI) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ListSupportedTokensRI) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetType

`func (o *ListSupportedTokensRI) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListSupportedTokensRI) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListSupportedTokensRI) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


