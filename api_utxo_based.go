/*
 * CryptoAPIs
 *
 * Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.
 *
 * API version: 2.0.0
 * Contact: developers@cryptoapis.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// UTXOBasedApiService UTXOBasedApi service
type UTXOBasedApiService service

type ApiGetHDWalletXPubYPubZPubDetailsRequest struct {
	ctx _context.Context
	ApiService *UTXOBasedApiService
	blockchain string
	extendedPublicKey string
	network string
	context *string
	derivation *string
}

func (r ApiGetHDWalletXPubYPubZPubDetailsRequest) Context(context string) ApiGetHDWalletXPubYPubZPubDetailsRequest {
	r.context = &context
	return r
}
func (r ApiGetHDWalletXPubYPubZPubDetailsRequest) Derivation(derivation string) ApiGetHDWalletXPubYPubZPubDetailsRequest {
	r.derivation = &derivation
	return r
}

func (r ApiGetHDWalletXPubYPubZPubDetailsRequest) Execute() (GetHDWalletXPubYPubZPubDetailsR, *_nethttp.Response, error) {
	return r.ApiService.GetHDWalletXPubYPubZPubDetailsExecute(r)
}

/*
 * GetHDWalletXPubYPubZPubDetails Get HD Wallet (xPub, yPub, zPub) Details
 * HD wallet details is useful endpoint to get the most important data about HD wallet without the need to do a lot of calculations, once the HD Wallet is synced using Sync endpoint we keep it up to date and we calculate these details in advance.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param blockchain Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
 * @param extendedPublicKey Defines the account extended publicly known key which is used to derive all child public keys.
 * @param network Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
 * @return ApiGetHDWalletXPubYPubZPubDetailsRequest
 */
func (a *UTXOBasedApiService) GetHDWalletXPubYPubZPubDetails(ctx _context.Context, blockchain string, extendedPublicKey string, network string) ApiGetHDWalletXPubYPubZPubDetailsRequest {
	return ApiGetHDWalletXPubYPubZPubDetailsRequest{
		ApiService: a,
		ctx: ctx,
		blockchain: blockchain,
		extendedPublicKey: extendedPublicKey,
		network: network,
	}
}

/*
 * Execute executes the request
 * @return GetHDWalletXPubYPubZPubDetailsR
 */
func (a *UTXOBasedApiService) GetHDWalletXPubYPubZPubDetailsExecute(r ApiGetHDWalletXPubYPubZPubDetailsRequest) (GetHDWalletXPubYPubZPubDetailsR, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetHDWalletXPubYPubZPubDetailsR
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UTXOBasedApiService.GetHDWalletXPubYPubZPubDetails")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/blockchain-data/{blockchain}/{network}/hd/{extendedPublicKey}/details"
	localVarPath = strings.Replace(localVarPath, "{"+"blockchain"+"}", _neturl.PathEscape(parameterToString(r.blockchain, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extendedPublicKey"+"}", _neturl.PathEscape(parameterToString(r.extendedPublicKey, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"network"+"}", _neturl.PathEscape(parameterToString(r.network, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.context != nil {
		localVarQueryParams.Add("context", parameterToString(*r.context, ""))
	}
	if r.derivation != nil {
		localVarQueryParams.Add("derivation", parameterToString(*r.derivation, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v XpubNotSynced
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InvalidApiKey
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v InsufficientCredits
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v FeatureMainnetsNotAllowedForPlan
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v InvalidData
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v UnsupportedMediaType
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v XpubSyncInProgress
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RequestLimitReached
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v UnexpectedServerError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListHDWalletXPubYPubZPubTransactionsRequest struct {
	ctx _context.Context
	ApiService *UTXOBasedApiService
	blockchain string
	extendedPublicKey string
	network string
	context *string
	derivation *string
	limit *int32
	offset *int32
}

func (r ApiListHDWalletXPubYPubZPubTransactionsRequest) Context(context string) ApiListHDWalletXPubYPubZPubTransactionsRequest {
	r.context = &context
	return r
}
func (r ApiListHDWalletXPubYPubZPubTransactionsRequest) Derivation(derivation string) ApiListHDWalletXPubYPubZPubTransactionsRequest {
	r.derivation = &derivation
	return r
}
func (r ApiListHDWalletXPubYPubZPubTransactionsRequest) Limit(limit int32) ApiListHDWalletXPubYPubZPubTransactionsRequest {
	r.limit = &limit
	return r
}
func (r ApiListHDWalletXPubYPubZPubTransactionsRequest) Offset(offset int32) ApiListHDWalletXPubYPubZPubTransactionsRequest {
	r.offset = &offset
	return r
}

func (r ApiListHDWalletXPubYPubZPubTransactionsRequest) Execute() (ListHDWalletXPubYPubZPubTransactionsR, *_nethttp.Response, error) {
	return r.ApiService.ListHDWalletXPubYPubZPubTransactionsExecute(r)
}

/*
 * ListHDWalletXPubYPubZPubTransactions List HD Wallet (xPub, yPub, zPub) Transactions
 * This endpoint will list HD Wallet transactions.

{note}Please note that listing data from the same type will apply pagination on the results.{/note}
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param blockchain Represents the specific blockchain.
 * @param extendedPublicKey Defines the master public key (xPub) of the account.
 * @param network Represents the specific network.
 * @return ApiListHDWalletXPubYPubZPubTransactionsRequest
 */
func (a *UTXOBasedApiService) ListHDWalletXPubYPubZPubTransactions(ctx _context.Context, blockchain string, extendedPublicKey string, network string) ApiListHDWalletXPubYPubZPubTransactionsRequest {
	return ApiListHDWalletXPubYPubZPubTransactionsRequest{
		ApiService: a,
		ctx: ctx,
		blockchain: blockchain,
		extendedPublicKey: extendedPublicKey,
		network: network,
	}
}

/*
 * Execute executes the request
 * @return ListHDWalletXPubYPubZPubTransactionsR
 */
func (a *UTXOBasedApiService) ListHDWalletXPubYPubZPubTransactionsExecute(r ApiListHDWalletXPubYPubZPubTransactionsRequest) (ListHDWalletXPubYPubZPubTransactionsR, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListHDWalletXPubYPubZPubTransactionsR
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UTXOBasedApiService.ListHDWalletXPubYPubZPubTransactions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/blockchain-data/{blockchain}/{network}/hd/{extendedPublicKey}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"blockchain"+"}", _neturl.PathEscape(parameterToString(r.blockchain, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extendedPublicKey"+"}", _neturl.PathEscape(parameterToString(r.extendedPublicKey, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"network"+"}", _neturl.PathEscape(parameterToString(r.network, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.context != nil {
		localVarQueryParams.Add("context", parameterToString(*r.context, ""))
	}
	if r.derivation != nil {
		localVarQueryParams.Add("derivation", parameterToString(*r.derivation, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v XpubNotSynced
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InvalidApiKey
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v InsufficientCredits
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v FeatureMainnetsNotAllowedForPlan
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v InvalidData
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v UnsupportedMediaType
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v XpubSyncInProgress
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RequestLimitReached
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v UnexpectedServerError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSyncHDWalletXPubYPubZPubRequest struct {
	ctx _context.Context
	ApiService *UTXOBasedApiService
	blockchain string
	network string
	context *string
	syncHDWalletXPubYPubZPubRB *SyncHDWalletXPubYPubZPubRB
}

func (r ApiSyncHDWalletXPubYPubZPubRequest) Context(context string) ApiSyncHDWalletXPubYPubZPubRequest {
	r.context = &context
	return r
}
func (r ApiSyncHDWalletXPubYPubZPubRequest) SyncHDWalletXPubYPubZPubRB(syncHDWalletXPubYPubZPubRB SyncHDWalletXPubYPubZPubRB) ApiSyncHDWalletXPubYPubZPubRequest {
	r.syncHDWalletXPubYPubZPubRB = &syncHDWalletXPubYPubZPubRB
	return r
}

func (r ApiSyncHDWalletXPubYPubZPubRequest) Execute() (SyncHDWalletXPubYPubZPubR, *_nethttp.Response, error) {
	return r.ApiService.SyncHDWalletXPubYPubZPubExecute(r)
}

/*
 * SyncHDWalletXPubYPubZPub Sync HD Wallet (xPub, yPub, zPub)
 * HD wallets usually have a lot of addresses and transactions, getting the data on demand is a heavy operation. That's why we have created this feature, to be able to get HD wallet details or transactions this HD wallet must be synced first. In addition to the initial sync we keep updating the synced HD wallets all the time.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param blockchain Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
 * @param network Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
 * @return ApiSyncHDWalletXPubYPubZPubRequest
 */
func (a *UTXOBasedApiService) SyncHDWalletXPubYPubZPub(ctx _context.Context, blockchain string, network string) ApiSyncHDWalletXPubYPubZPubRequest {
	return ApiSyncHDWalletXPubYPubZPubRequest{
		ApiService: a,
		ctx: ctx,
		blockchain: blockchain,
		network: network,
	}
}

/*
 * Execute executes the request
 * @return SyncHDWalletXPubYPubZPubR
 */
func (a *UTXOBasedApiService) SyncHDWalletXPubYPubZPubExecute(r ApiSyncHDWalletXPubYPubZPubRequest) (SyncHDWalletXPubYPubZPubR, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SyncHDWalletXPubYPubZPubR
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UTXOBasedApiService.SyncHDWalletXPubYPubZPub")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/blockchain-data/{blockchain}/{network}/hd/sync"
	localVarPath = strings.Replace(localVarPath, "{"+"blockchain"+"}", _neturl.PathEscape(parameterToString(r.blockchain, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"network"+"}", _neturl.PathEscape(parameterToString(r.network, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.context != nil {
		localVarQueryParams.Add("context", parameterToString(*r.context, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.syncHDWalletXPubYPubZPubRB
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InvalidXpub
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InvalidApiKey
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v InsufficientCredits
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v FeatureMainnetsNotAllowedForPlan
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v AlreadyExists
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v UnsupportedMediaType
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v XpubSyncInProgress
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RequestLimitReached
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v UnexpectedServerError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
