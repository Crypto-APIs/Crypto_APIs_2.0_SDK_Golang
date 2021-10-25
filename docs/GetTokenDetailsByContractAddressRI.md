# GetTokenDetailsByContractAddressRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenDecimals** | **string** | Defines the number of decimals that the token possesses. | 
**TokenName** | Pointer to **string** | Specifies the token&#39;s name. | [optional] 
**TokenSymbol** | Pointer to **string** | Defines the unique symbol of the token. | [optional] 
**TokenType** | **string** | Defines the type of the token. | 
**TotalSupply** | **string** | Defines the total number of tokens created that exist on the market minus the ones that have been burned. | 

## Methods

### NewGetTokenDetailsByContractAddressRI

`func NewGetTokenDetailsByContractAddressRI(tokenDecimals string, tokenType string, totalSupply string, ) *GetTokenDetailsByContractAddressRI`

NewGetTokenDetailsByContractAddressRI instantiates a new GetTokenDetailsByContractAddressRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTokenDetailsByContractAddressRIWithDefaults

`func NewGetTokenDetailsByContractAddressRIWithDefaults() *GetTokenDetailsByContractAddressRI`

NewGetTokenDetailsByContractAddressRIWithDefaults instantiates a new GetTokenDetailsByContractAddressRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenDecimals

`func (o *GetTokenDetailsByContractAddressRI) GetTokenDecimals() string`

GetTokenDecimals returns the TokenDecimals field if non-nil, zero value otherwise.

### GetTokenDecimalsOk

`func (o *GetTokenDetailsByContractAddressRI) GetTokenDecimalsOk() (*string, bool)`

GetTokenDecimalsOk returns a tuple with the TokenDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDecimals

`func (o *GetTokenDetailsByContractAddressRI) SetTokenDecimals(v string)`

SetTokenDecimals sets TokenDecimals field to given value.


### GetTokenName

`func (o *GetTokenDetailsByContractAddressRI) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *GetTokenDetailsByContractAddressRI) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *GetTokenDetailsByContractAddressRI) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *GetTokenDetailsByContractAddressRI) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.

### GetTokenSymbol

`func (o *GetTokenDetailsByContractAddressRI) GetTokenSymbol() string`

GetTokenSymbol returns the TokenSymbol field if non-nil, zero value otherwise.

### GetTokenSymbolOk

`func (o *GetTokenDetailsByContractAddressRI) GetTokenSymbolOk() (*string, bool)`

GetTokenSymbolOk returns a tuple with the TokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSymbol

`func (o *GetTokenDetailsByContractAddressRI) SetTokenSymbol(v string)`

SetTokenSymbol sets TokenSymbol field to given value.

### HasTokenSymbol

`func (o *GetTokenDetailsByContractAddressRI) HasTokenSymbol() bool`

HasTokenSymbol returns a boolean if a field has been set.

### GetTokenType

`func (o *GetTokenDetailsByContractAddressRI) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *GetTokenDetailsByContractAddressRI) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *GetTokenDetailsByContractAddressRI) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetTotalSupply

`func (o *GetTokenDetailsByContractAddressRI) GetTotalSupply() string`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *GetTokenDetailsByContractAddressRI) GetTotalSupplyOk() (*string, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *GetTokenDetailsByContractAddressRI) SetTotalSupply(v string)`

SetTotalSupply sets TotalSupply field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


