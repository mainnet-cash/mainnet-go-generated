/*
 * Mainnet Cash
 *
 * A developer friendly bitcoin cash wallet api  This API is currently in *active* development, breaking changes may be made prior to official release of version 1.0.0. 
 *
 * API version: 2.3.0
 * Contact: hello@mainnet.cash
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	_bytes "bytes"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// WalletSignedApiService WalletSignedApi service
type WalletSignedApiService service

// SignedMessageSignOpts Optional parameters for the method 'SignedMessageSign'
type SignedMessageSignOpts struct {
    CreateSignedMessageRequest optional.Interface
}

/*
SignedMessageSign Returns the message signature
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *SignedMessageSignOpts - Optional Parameters:
 * @param "CreateSignedMessageRequest" (optional.Interface of CreateSignedMessageRequest) -  Sign a message 
@return SignedMessageResponse
*/
func (a *WalletSignedApiService) SignedMessageSign(ctx _context.Context, localVarOptionals *SignedMessageSignOpts) (SignedMessageResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SignedMessageResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/wallet/signed/sign"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if localVarOptionals != nil && localVarOptionals.CreateSignedMessageRequest.IsSet() {
		localVarOptionalCreateSignedMessageRequest, localVarOptionalCreateSignedMessageRequestok := localVarOptionals.CreateSignedMessageRequest.Value().(CreateSignedMessageRequest)
		if !localVarOptionalCreateSignedMessageRequestok {
			return localVarReturnValue, nil, reportError("createSignedMessageRequest should be CreateSignedMessageRequest")
		}
		localVarPostBody = &localVarOptionalCreateSignedMessageRequest
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

// SignedMessageVerifyOpts Optional parameters for the method 'SignedMessageVerify'
type SignedMessageVerifyOpts struct {
    VerifySignedMessageRequest optional.Interface
}

/*
SignedMessageVerify Returns a boolean indicating whether message was valid for a given address
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *SignedMessageVerifyOpts - Optional Parameters:
 * @param "VerifySignedMessageRequest" (optional.Interface of VerifySignedMessageRequest) -  Sign a message 
@return VerifySignedMessageResponse
*/
func (a *WalletSignedApiService) SignedMessageVerify(ctx _context.Context, localVarOptionals *SignedMessageVerifyOpts) (VerifySignedMessageResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  VerifySignedMessageResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/wallet/signed/verify"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if localVarOptionals != nil && localVarOptionals.VerifySignedMessageRequest.IsSet() {
		localVarOptionalVerifySignedMessageRequest, localVarOptionalVerifySignedMessageRequestok := localVarOptionals.VerifySignedMessageRequest.Value().(VerifySignedMessageRequest)
		if !localVarOptionalVerifySignedMessageRequestok {
			return localVarReturnValue, nil, reportError("verifySignedMessageRequest should be VerifySignedMessageRequest")
		}
		localVarPostBody = &localVarOptionalVerifySignedMessageRequest
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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
