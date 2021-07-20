# GetContractDetailsByAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenDecimals** | **string** | Defines the number of decimals that the token possesses. | 
**TokenName** | Pointer to **string** | Specifies the token&#39;s name. | [optional] 
**TokenSymbol** | Pointer to **string** | Defines the unique symbol of the token. | [optional] 
**TokenType** | **string** | Defines the type of the token. | 
**TotalSupply** | **string** | Defines the total number of tokens created that exist on the market minus the ones that have been burned. | 

## Methods

### NewGetContractDetailsByAddressRI

`func NewGetContractDetailsByAddressRI(tokenDecimals string, tokenType string, totalSupply string, ) *GetContractDetailsByAddressRI`

NewGetContractDetailsByAddressRI instantiates a new GetContractDetailsByAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContractDetailsByAddressRIWithDefaults

`func NewGetContractDetailsByAddressRIWithDefaults() *GetContractDetailsByAddressRI`

NewGetContractDetailsByAddressRIWithDefaults instantiates a new GetContractDetailsByAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenDecimals

`func (o *GetContractDetailsByAddressRI) GetTokenDecimals() string`

GetTokenDecimals returns the TokenDecimals field if non-nil, zero value otherwise.

### GetTokenDecimalsOk

`func (o *GetContractDetailsByAddressRI) GetTokenDecimalsOk() (*string, bool)`

GetTokenDecimalsOk returns a tuple with the TokenDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDecimals

`func (o *GetContractDetailsByAddressRI) SetTokenDecimals(v string)`

SetTokenDecimals sets TokenDecimals field to given value.


### GetTokenName

`func (o *GetContractDetailsByAddressRI) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *GetContractDetailsByAddressRI) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *GetContractDetailsByAddressRI) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *GetContractDetailsByAddressRI) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.

### GetTokenSymbol

`func (o *GetContractDetailsByAddressRI) GetTokenSymbol() string`

GetTokenSymbol returns the TokenSymbol field if non-nil, zero value otherwise.

### GetTokenSymbolOk

`func (o *GetContractDetailsByAddressRI) GetTokenSymbolOk() (*string, bool)`

GetTokenSymbolOk returns a tuple with the TokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSymbol

`func (o *GetContractDetailsByAddressRI) SetTokenSymbol(v string)`

SetTokenSymbol sets TokenSymbol field to given value.

### HasTokenSymbol

`func (o *GetContractDetailsByAddressRI) HasTokenSymbol() bool`

HasTokenSymbol returns a boolean if a field has been set.

### GetTokenType

`func (o *GetContractDetailsByAddressRI) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *GetContractDetailsByAddressRI) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *GetContractDetailsByAddressRI) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetTotalSupply

`func (o *GetContractDetailsByAddressRI) GetTotalSupply() string`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *GetContractDetailsByAddressRI) GetTotalSupplyOk() (*string, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *GetContractDetailsByAddressRI) SetTotalSupply(v string)`

SetTotalSupply sets TotalSupply field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


