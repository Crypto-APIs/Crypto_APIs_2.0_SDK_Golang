# ListTokensByAddressResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmedBalance** | **string** | Defines the token balance that has been confirmed. | 
**ContractAddress** | **string** | Represents the contract address of the token, which controls its logic. It is not the address that holds the tokens. | 
**Name** | **string** | Defines the token&#39;s name as a string. | 
**Symbol** | **string** | Defines the token symbol by which the token contract is known. It is usually 3-4 characters in length. | 
**Type** | **string** | Defines the specific token type. | 

## Methods

### NewListTokensByAddressResponseItem

`func NewListTokensByAddressResponseItem(confirmedBalance string, contractAddress string, name string, symbol string, type_ string, ) *ListTokensByAddressResponseItem`

NewListTokensByAddressResponseItem instantiates a new ListTokensByAddressResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTokensByAddressResponseItemWithDefaults

`func NewListTokensByAddressResponseItemWithDefaults() *ListTokensByAddressResponseItem`

NewListTokensByAddressResponseItemWithDefaults instantiates a new ListTokensByAddressResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmedBalance

`func (o *ListTokensByAddressResponseItem) GetConfirmedBalance() string`

GetConfirmedBalance returns the ConfirmedBalance field if non-nil, zero value otherwise.

### GetConfirmedBalanceOk

`func (o *ListTokensByAddressResponseItem) GetConfirmedBalanceOk() (*string, bool)`

GetConfirmedBalanceOk returns a tuple with the ConfirmedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedBalance

`func (o *ListTokensByAddressResponseItem) SetConfirmedBalance(v string)`

SetConfirmedBalance sets ConfirmedBalance field to given value.


### GetContractAddress

`func (o *ListTokensByAddressResponseItem) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ListTokensByAddressResponseItem) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ListTokensByAddressResponseItem) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetName

`func (o *ListTokensByAddressResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListTokensByAddressResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListTokensByAddressResponseItem) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *ListTokensByAddressResponseItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ListTokensByAddressResponseItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ListTokensByAddressResponseItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetType

`func (o *ListTokensByAddressResponseItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListTokensByAddressResponseItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListTokensByAddressResponseItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


