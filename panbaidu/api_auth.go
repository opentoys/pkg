/*
xpan

xpanapi

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package panbaidu

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// AuthApiService AuthApi service
type AuthApiService service

type ApiOauthTokenCode2tokenRequest struct {
	ctx _context.Context
	ApiService *AuthApiService
	code *string
	clientId *string
	clientSecret *string
	redirectUri *string
}

func (r ApiOauthTokenCode2tokenRequest) Code(code string) ApiOauthTokenCode2tokenRequest {
	r.code = &code
	return r
}
func (r ApiOauthTokenCode2tokenRequest) ClientId(clientId string) ApiOauthTokenCode2tokenRequest {
	r.clientId = &clientId
	return r
}
func (r ApiOauthTokenCode2tokenRequest) ClientSecret(clientSecret string) ApiOauthTokenCode2tokenRequest {
	r.clientSecret = &clientSecret
	return r
}
func (r ApiOauthTokenCode2tokenRequest) RedirectUri(redirectUri string) ApiOauthTokenCode2tokenRequest {
	r.redirectUri = &redirectUri
	return r
}

func (r ApiOauthTokenCode2tokenRequest) Execute() (OauthTokenAuthorizationCodeResponse, *_nethttp.Response, error) {
	return r.ApiService.OauthTokenCode2tokenExecute(r)
}

/*
OauthTokenCode2token Method for OauthTokenCode2token

get accesstoken by authorization code

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOauthTokenCode2tokenRequest
*/
func (a *AuthApiService) OauthTokenCode2token(ctx _context.Context) ApiOauthTokenCode2tokenRequest {
	return ApiOauthTokenCode2tokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OauthTokenAuthorizationCodeResponse
func (a *AuthApiService) OauthTokenCode2tokenExecute(r ApiOauthTokenCode2tokenRequest) (OauthTokenAuthorizationCodeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  OauthTokenAuthorizationCodeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.OauthTokenCode2token")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/2.0/token?grant_type=authorization_code&openapi=xpansdk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.code == nil {
		return localVarReturnValue, nil, reportError("code is required and must be specified")
	}
	if r.clientId == nil {
		return localVarReturnValue, nil, reportError("clientId is required and must be specified")
	}
	if r.clientSecret == nil {
		return localVarReturnValue, nil, reportError("clientSecret is required and must be specified")
	}
	if r.redirectUri == nil {
		return localVarReturnValue, nil, reportError("redirectUri is required and must be specified")
	}

	localVarQueryParams.Add("code", parameterToString(*r.code, ""))
	localVarQueryParams.Add("client_id", parameterToString(*r.clientId, ""))
	localVarQueryParams.Add("client_secret", parameterToString(*r.clientSecret, ""))
	localVarQueryParams.Add("redirect_uri", parameterToString(*r.redirectUri, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=UTF-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type ApiOauthTokenDeviceCodeRequest struct {
	ctx _context.Context
	ApiService *AuthApiService
	clientId *string
	scope *string
}

func (r ApiOauthTokenDeviceCodeRequest) ClientId(clientId string) ApiOauthTokenDeviceCodeRequest {
	r.clientId = &clientId
	return r
}
func (r ApiOauthTokenDeviceCodeRequest) Scope(scope string) ApiOauthTokenDeviceCodeRequest {
	r.scope = &scope
	return r
}

func (r ApiOauthTokenDeviceCodeRequest) Execute() (OauthTokenDeviceCodeResponse, *_nethttp.Response, error) {
	return r.ApiService.OauthTokenDeviceCodeExecute(r)
}

/*
OauthTokenDeviceCode Method for OauthTokenDeviceCode

get device code and user code

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOauthTokenDeviceCodeRequest
*/
func (a *AuthApiService) OauthTokenDeviceCode(ctx _context.Context) ApiOauthTokenDeviceCodeRequest {
	return ApiOauthTokenDeviceCodeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OauthTokenDeviceCodeResponse
func (a *AuthApiService) OauthTokenDeviceCodeExecute(r ApiOauthTokenDeviceCodeRequest) (OauthTokenDeviceCodeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  OauthTokenDeviceCodeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.OauthTokenDeviceCode")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/2.0/device/code?response_type=device_code&openapi=xpansdk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.clientId == nil {
		return localVarReturnValue, nil, reportError("clientId is required and must be specified")
	}
	if r.scope == nil {
		return localVarReturnValue, nil, reportError("scope is required and must be specified")
	}

	localVarQueryParams.Add("client_id", parameterToString(*r.clientId, ""))
	localVarQueryParams.Add("scope", parameterToString(*r.scope, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=UTF-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type ApiOauthTokenDeviceTokenRequest struct {
	ctx _context.Context
	ApiService *AuthApiService
	code *string
	clientId *string
	clientSecret *string
}

func (r ApiOauthTokenDeviceTokenRequest) Code(code string) ApiOauthTokenDeviceTokenRequest {
	r.code = &code
	return r
}
func (r ApiOauthTokenDeviceTokenRequest) ClientId(clientId string) ApiOauthTokenDeviceTokenRequest {
	r.clientId = &clientId
	return r
}
func (r ApiOauthTokenDeviceTokenRequest) ClientSecret(clientSecret string) ApiOauthTokenDeviceTokenRequest {
	r.clientSecret = &clientSecret
	return r
}

func (r ApiOauthTokenDeviceTokenRequest) Execute() (OauthTokenDeviceTokenResponse, *_nethttp.Response, error) {
	return r.ApiService.OauthTokenDeviceTokenExecute(r)
}

/*
OauthTokenDeviceToken Method for OauthTokenDeviceToken

get access_token

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOauthTokenDeviceTokenRequest
*/
func (a *AuthApiService) OauthTokenDeviceToken(ctx _context.Context) ApiOauthTokenDeviceTokenRequest {
	return ApiOauthTokenDeviceTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OauthTokenDeviceTokenResponse
func (a *AuthApiService) OauthTokenDeviceTokenExecute(r ApiOauthTokenDeviceTokenRequest) (OauthTokenDeviceTokenResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  OauthTokenDeviceTokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.OauthTokenDeviceToken")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/2.0/token?grant_type=device_token&openapi=xpansdk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.code == nil {
		return localVarReturnValue, nil, reportError("code is required and must be specified")
	}
	if r.clientId == nil {
		return localVarReturnValue, nil, reportError("clientId is required and must be specified")
	}
	if r.clientSecret == nil {
		return localVarReturnValue, nil, reportError("clientSecret is required and must be specified")
	}

	localVarQueryParams.Add("code", parameterToString(*r.code, ""))
	localVarQueryParams.Add("client_id", parameterToString(*r.clientId, ""))
	localVarQueryParams.Add("client_secret", parameterToString(*r.clientSecret, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=UTF-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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

type ApiOauthTokenRefreshTokenRequest struct {
	ctx _context.Context
	ApiService *AuthApiService
	refreshToken *string
	clientId *string
	clientSecret *string
}

func (r ApiOauthTokenRefreshTokenRequest) RefreshToken(refreshToken string) ApiOauthTokenRefreshTokenRequest {
	r.refreshToken = &refreshToken
	return r
}
func (r ApiOauthTokenRefreshTokenRequest) ClientId(clientId string) ApiOauthTokenRefreshTokenRequest {
	r.clientId = &clientId
	return r
}
func (r ApiOauthTokenRefreshTokenRequest) ClientSecret(clientSecret string) ApiOauthTokenRefreshTokenRequest {
	r.clientSecret = &clientSecret
	return r
}

func (r ApiOauthTokenRefreshTokenRequest) Execute() (OauthTokenRefreshTokenResponse, *_nethttp.Response, error) {
	return r.ApiService.OauthTokenRefreshTokenExecute(r)
}

/*
OauthTokenRefreshToken Method for OauthTokenRefreshToken

authorization code

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOauthTokenRefreshTokenRequest
*/
func (a *AuthApiService) OauthTokenRefreshToken(ctx _context.Context) ApiOauthTokenRefreshTokenRequest {
	return ApiOauthTokenRefreshTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OauthTokenRefreshTokenResponse
func (a *AuthApiService) OauthTokenRefreshTokenExecute(r ApiOauthTokenRefreshTokenRequest) (OauthTokenRefreshTokenResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  OauthTokenRefreshTokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.OauthTokenRefreshToken")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/2.0/token?grant_type=refresh_token&openapi=xpansdk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.refreshToken == nil {
		return localVarReturnValue, nil, reportError("refreshToken is required and must be specified")
	}
	if r.clientId == nil {
		return localVarReturnValue, nil, reportError("clientId is required and must be specified")
	}
	if r.clientSecret == nil {
		return localVarReturnValue, nil, reportError("clientSecret is required and must be specified")
	}

	localVarQueryParams.Add("refresh_token", parameterToString(*r.refreshToken, ""))
	localVarQueryParams.Add("client_id", parameterToString(*r.clientId, ""))
	localVarQueryParams.Add("client_secret", parameterToString(*r.clientSecret, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=UTF-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
