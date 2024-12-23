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

// FilemanagerApiService FilemanagerApi service
type FilemanagerApiService service

type ApiFilemanagercopyRequest struct {
	ctx _context.Context
	ApiService *FilemanagerApiService
	accessToken *string
	async *int32
	filelist *string
	ondup *string
}

func (r ApiFilemanagercopyRequest) AccessToken(accessToken string) ApiFilemanagercopyRequest {
	r.accessToken = &accessToken
	return r
}
// async
func (r ApiFilemanagercopyRequest) Async(async int32) ApiFilemanagercopyRequest {
	r.async = &async
	return r
}
// filelist
func (r ApiFilemanagercopyRequest) Filelist(filelist string) ApiFilemanagercopyRequest {
	r.filelist = &filelist
	return r
}
// ondup
func (r ApiFilemanagercopyRequest) Ondup(ondup string) ApiFilemanagercopyRequest {
	r.ondup = &ondup
	return r
}

func (r ApiFilemanagercopyRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FilemanagercopyExecute(r)
}

/*
Filemanagercopy Method for Filemanagercopy

filemanager copy

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFilemanagercopyRequest
*/
func (a *FilemanagerApiService) Filemanagercopy(ctx _context.Context) ApiFilemanagercopyRequest {
	return ApiFilemanagercopyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *FilemanagerApiService) FilemanagercopyExecute(r ApiFilemanagercopyRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilemanagerApiService.Filemanagercopy")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rest/2.0/xpan/file?method=filemanager&opera=copy&openapi=xpansdk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.accessToken == nil {
		return nil, reportError("accessToken is required and must be specified")
	}
	if r.async == nil {
		return nil, reportError("async is required and must be specified")
	}
	if r.filelist == nil {
		return nil, reportError("filelist is required and must be specified")
	}

	localVarQueryParams.Add("access_token", parameterToString(*r.accessToken, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("async", parameterToString(*r.async, ""))
	localVarFormParams.Add("filelist", parameterToString(*r.filelist, ""))
	if r.ondup != nil {
		localVarFormParams.Add("ondup", parameterToString(*r.ondup, ""))
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFilemanagerdeleteRequest struct {
	ctx _context.Context
	ApiService *FilemanagerApiService
	accessToken *string
	async *int32
	filelist *string
	ondup *string
}

func (r ApiFilemanagerdeleteRequest) AccessToken(accessToken string) ApiFilemanagerdeleteRequest {
	r.accessToken = &accessToken
	return r
}
// async
func (r ApiFilemanagerdeleteRequest) Async(async int32) ApiFilemanagerdeleteRequest {
	r.async = &async
	return r
}
// filelist
func (r ApiFilemanagerdeleteRequest) Filelist(filelist string) ApiFilemanagerdeleteRequest {
	r.filelist = &filelist
	return r
}
// ondup
func (r ApiFilemanagerdeleteRequest) Ondup(ondup string) ApiFilemanagerdeleteRequest {
	r.ondup = &ondup
	return r
}

func (r ApiFilemanagerdeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FilemanagerdeleteExecute(r)
}

/*
Filemanagerdelete Method for Filemanagerdelete

filemanager delete

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFilemanagerdeleteRequest
*/
func (a *FilemanagerApiService) Filemanagerdelete(ctx _context.Context) ApiFilemanagerdeleteRequest {
	return ApiFilemanagerdeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *FilemanagerApiService) FilemanagerdeleteExecute(r ApiFilemanagerdeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilemanagerApiService.Filemanagerdelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rest/2.0/xpan/file?method=filemanager&opera=delete&openapi=xpansdk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.accessToken == nil {
		return nil, reportError("accessToken is required and must be specified")
	}
	if r.async == nil {
		return nil, reportError("async is required and must be specified")
	}
	if r.filelist == nil {
		return nil, reportError("filelist is required and must be specified")
	}

	localVarQueryParams.Add("access_token", parameterToString(*r.accessToken, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("async", parameterToString(*r.async, ""))
	localVarFormParams.Add("filelist", parameterToString(*r.filelist, ""))
	if r.ondup != nil {
		localVarFormParams.Add("ondup", parameterToString(*r.ondup, ""))
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFilemanagermoveRequest struct {
	ctx _context.Context
	ApiService *FilemanagerApiService
	accessToken *string
	async *int32
	filelist *string
	ondup *string
}

func (r ApiFilemanagermoveRequest) AccessToken(accessToken string) ApiFilemanagermoveRequest {
	r.accessToken = &accessToken
	return r
}
// async
func (r ApiFilemanagermoveRequest) Async(async int32) ApiFilemanagermoveRequest {
	r.async = &async
	return r
}
// filelist
func (r ApiFilemanagermoveRequest) Filelist(filelist string) ApiFilemanagermoveRequest {
	r.filelist = &filelist
	return r
}
// ondup
func (r ApiFilemanagermoveRequest) Ondup(ondup string) ApiFilemanagermoveRequest {
	r.ondup = &ondup
	return r
}

func (r ApiFilemanagermoveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FilemanagermoveExecute(r)
}

/*
Filemanagermove Method for Filemanagermove

filemanager move

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFilemanagermoveRequest
*/
func (a *FilemanagerApiService) Filemanagermove(ctx _context.Context) ApiFilemanagermoveRequest {
	return ApiFilemanagermoveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *FilemanagerApiService) FilemanagermoveExecute(r ApiFilemanagermoveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilemanagerApiService.Filemanagermove")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rest/2.0/xpan/file?method=filemanager&opera=move&openapi=xpansdk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.accessToken == nil {
		return nil, reportError("accessToken is required and must be specified")
	}
	if r.async == nil {
		return nil, reportError("async is required and must be specified")
	}
	if r.filelist == nil {
		return nil, reportError("filelist is required and must be specified")
	}

	localVarQueryParams.Add("access_token", parameterToString(*r.accessToken, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("async", parameterToString(*r.async, ""))
	localVarFormParams.Add("filelist", parameterToString(*r.filelist, ""))
	if r.ondup != nil {
		localVarFormParams.Add("ondup", parameterToString(*r.ondup, ""))
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFilemanagerrenameRequest struct {
	ctx _context.Context
	ApiService *FilemanagerApiService
	accessToken *string
	async *int32
	filelist *string
	ondup *string
}

func (r ApiFilemanagerrenameRequest) AccessToken(accessToken string) ApiFilemanagerrenameRequest {
	r.accessToken = &accessToken
	return r
}
// async
func (r ApiFilemanagerrenameRequest) Async(async int32) ApiFilemanagerrenameRequest {
	r.async = &async
	return r
}
// filelist
func (r ApiFilemanagerrenameRequest) Filelist(filelist string) ApiFilemanagerrenameRequest {
	r.filelist = &filelist
	return r
}
// ondup
func (r ApiFilemanagerrenameRequest) Ondup(ondup string) ApiFilemanagerrenameRequest {
	r.ondup = &ondup
	return r
}

func (r ApiFilemanagerrenameRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.FilemanagerrenameExecute(r)
}

/*
Filemanagerrename Method for Filemanagerrename

filemanager rename

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFilemanagerrenameRequest
*/
func (a *FilemanagerApiService) Filemanagerrename(ctx _context.Context) ApiFilemanagerrenameRequest {
	return ApiFilemanagerrenameRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *FilemanagerApiService) FilemanagerrenameExecute(r ApiFilemanagerrenameRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilemanagerApiService.Filemanagerrename")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rest/2.0/xpan/file?method=filemanager&opera=rename&openapi=xpansdk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.accessToken == nil {
		return nil, reportError("accessToken is required and must be specified")
	}
	if r.async == nil {
		return nil, reportError("async is required and must be specified")
	}
	if r.filelist == nil {
		return nil, reportError("filelist is required and must be specified")
	}

	localVarQueryParams.Add("access_token", parameterToString(*r.accessToken, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("async", parameterToString(*r.async, ""))
	localVarFormParams.Add("filelist", parameterToString(*r.filelist, ""))
	if r.ondup != nil {
		localVarFormParams.Add("ondup", parameterToString(*r.ondup, ""))
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
