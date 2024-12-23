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
	"os"
)

// Linger please
var (
	_ _context.Context
)

// FileuploadApiService FileuploadApi service
type FileuploadApiService service

type ApiPcssuperfile2Request struct {
	ctx _context.Context
	ApiService *FileuploadApiService
	accessToken *string
	partseq *string
	path *string
	uploadid *string
	type_ *string
	file **os.File
}

func (r ApiPcssuperfile2Request) AccessToken(accessToken string) ApiPcssuperfile2Request {
	r.accessToken = &accessToken
	return r
}
func (r ApiPcssuperfile2Request) Partseq(partseq string) ApiPcssuperfile2Request {
	r.partseq = &partseq
	return r
}
func (r ApiPcssuperfile2Request) Path(path string) ApiPcssuperfile2Request {
	r.path = &path
	return r
}
func (r ApiPcssuperfile2Request) Uploadid(uploadid string) ApiPcssuperfile2Request {
	r.uploadid = &uploadid
	return r
}
func (r ApiPcssuperfile2Request) Type_(type_ string) ApiPcssuperfile2Request {
	r.type_ = &type_
	return r
}
// 要进行传送的本地文件分片
func (r ApiPcssuperfile2Request) File(file *os.File) ApiPcssuperfile2Request {
	r.file = &file
	return r
}

func (r ApiPcssuperfile2Request) Execute() (string, *_nethttp.Response, error) {
	return r.ApiService.Pcssuperfile2Execute(r)
}

/*
Pcssuperfile2 Method for Pcssuperfile2

分片上传，这里是实际的文件内容传送部分。一般多为大于4MB的文件，需将文件以4MB为单位切分，对切分后得到的n个分片一一调用该接口进行传送，以实现对原文件的传送（当然若不大于4MB，则直接该对文件进行传送即可）。

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPcssuperfile2Request
*/
func (a *FileuploadApiService) Pcssuperfile2(ctx _context.Context) ApiPcssuperfile2Request {
	return ApiPcssuperfile2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *FileuploadApiService) Pcssuperfile2Execute(r ApiPcssuperfile2Request) (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileuploadApiService.Pcssuperfile2")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rest/2.0/pcs/superfile2?method=upload&openapi=xpansdk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.accessToken == nil {
		return localVarReturnValue, nil, reportError("accessToken is required and must be specified")
	}
	if r.partseq == nil {
		return localVarReturnValue, nil, reportError("partseq is required and must be specified")
	}
	if r.path == nil {
		return localVarReturnValue, nil, reportError("path is required and must be specified")
	}
	if r.uploadid == nil {
		return localVarReturnValue, nil, reportError("uploadid is required and must be specified")
	}
	if r.type_ == nil {
		return localVarReturnValue, nil, reportError("type_ is required and must be specified")
	}

	localVarQueryParams.Add("access_token", parameterToString(*r.accessToken, ""))
	localVarQueryParams.Add("partseq", parameterToString(*r.partseq, ""))
	localVarQueryParams.Add("path", parameterToString(*r.path, ""))
	localVarQueryParams.Add("uploadid", parameterToString(*r.uploadid, ""))
	localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/html;charset=utf8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"

	var fileLocalVarFile *os.File
	if r.file != nil {
		fileLocalVarFile = *r.file
	}
	if fileLocalVarFile != nil {
		fbs, _ := _ioutil.ReadAll(fileLocalVarFile)
		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
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

type ApiXpanfilecreateRequest struct {
	ctx _context.Context
	ApiService *FileuploadApiService
	accessToken *string
	path *string
	isdir *int32
	size *int32
	uploadid *string
	blockList *string
	rtype *int32
}

func (r ApiXpanfilecreateRequest) AccessToken(accessToken string) ApiXpanfilecreateRequest {
	r.accessToken = &accessToken
	return r
}
// 与precreate的path值保持一致
func (r ApiXpanfilecreateRequest) Path(path string) ApiXpanfilecreateRequest {
	r.path = &path
	return r
}
// isdir
func (r ApiXpanfilecreateRequest) Isdir(isdir int32) ApiXpanfilecreateRequest {
	r.isdir = &isdir
	return r
}
// 与precreate的size值保持一致
func (r ApiXpanfilecreateRequest) Size(size int32) ApiXpanfilecreateRequest {
	r.size = &size
	return r
}
// precreate返回的uploadid
func (r ApiXpanfilecreateRequest) Uploadid(uploadid string) ApiXpanfilecreateRequest {
	r.uploadid = &uploadid
	return r
}
// 与precreate的block_list值保持一致
func (r ApiXpanfilecreateRequest) BlockList(blockList string) ApiXpanfilecreateRequest {
	r.blockList = &blockList
	return r
}
// rtype
func (r ApiXpanfilecreateRequest) Rtype(rtype int32) ApiXpanfilecreateRequest {
	r.rtype = &rtype
	return r
}

func (r ApiXpanfilecreateRequest) Execute() (Filecreateresponse, *_nethttp.Response, error) {
	return r.ApiService.XpanfilecreateExecute(r)
}

/*
Xpanfilecreate Method for Xpanfilecreate

将多个文件分片合并成一个文件，生成文件基本信息，完成文件的上传最后一步。

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiXpanfilecreateRequest
*/
func (a *FileuploadApiService) Xpanfilecreate(ctx _context.Context) ApiXpanfilecreateRequest {
	return ApiXpanfilecreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Filecreateresponse
func (a *FileuploadApiService) XpanfilecreateExecute(r ApiXpanfilecreateRequest) (Filecreateresponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  Filecreateresponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileuploadApiService.Xpanfilecreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rest/2.0/xpan/file?method=create&openapi=xpansdk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.accessToken == nil {
		return localVarReturnValue, nil, reportError("accessToken is required and must be specified")
	}
	if r.path == nil {
		return localVarReturnValue, nil, reportError("path is required and must be specified")
	}
	if r.isdir == nil {
		return localVarReturnValue, nil, reportError("isdir is required and must be specified")
	}
	if r.size == nil {
		return localVarReturnValue, nil, reportError("size is required and must be specified")
	}
	if r.uploadid == nil {
		return localVarReturnValue, nil, reportError("uploadid is required and must be specified")
	}
	if r.blockList == nil {
		return localVarReturnValue, nil, reportError("blockList is required and must be specified")
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
	localVarHTTPHeaderAccepts := []string{"application/json; charset=UTF-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("path", parameterToString(*r.path, ""))
	localVarFormParams.Add("isdir", parameterToString(*r.isdir, ""))
	localVarFormParams.Add("size", parameterToString(*r.size, ""))
	localVarFormParams.Add("uploadid", parameterToString(*r.uploadid, ""))
	localVarFormParams.Add("block_list", parameterToString(*r.blockList, ""))
	if r.rtype != nil {
		localVarFormParams.Add("rtype", parameterToString(*r.rtype, ""))
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

type ApiXpanfileprecreateRequest struct {
	ctx _context.Context
	ApiService *FileuploadApiService
	accessToken *string
	path *string
	isdir *int32
	size *int32
	autoinit *int32
	blockList *string
	rtype *int32
}

func (r ApiXpanfileprecreateRequest) AccessToken(accessToken string) ApiXpanfileprecreateRequest {
	r.accessToken = &accessToken
	return r
}
// 对于一般的第三方软件应用，路径以 \\\&quot;/apps/your-app-name/\\\&quot; 开头。对于小度等硬件应用，路径一般 \\\&quot;/来自：小度设备/\\\&quot; 开头。对于定制化配置的硬件应用，根据配置情况进行填写。
func (r ApiXpanfileprecreateRequest) Path(path string) ApiXpanfileprecreateRequest {
	r.path = &path
	return r
}
// isdir
func (r ApiXpanfileprecreateRequest) Isdir(isdir int32) ApiXpanfileprecreateRequest {
	r.isdir = &isdir
	return r
}
// size
func (r ApiXpanfileprecreateRequest) Size(size int32) ApiXpanfileprecreateRequest {
	r.size = &size
	return r
}
// autoinit
func (r ApiXpanfileprecreateRequest) Autoinit(autoinit int32) ApiXpanfileprecreateRequest {
	r.autoinit = &autoinit
	return r
}
// 由MD5字符串组成的list
func (r ApiXpanfileprecreateRequest) BlockList(blockList string) ApiXpanfileprecreateRequest {
	r.blockList = &blockList
	return r
}
// rtype
func (r ApiXpanfileprecreateRequest) Rtype(rtype int32) ApiXpanfileprecreateRequest {
	r.rtype = &rtype
	return r
}

func (r ApiXpanfileprecreateRequest) Execute() (Fileprecreateresponse, *_nethttp.Response, error) {
	return r.ApiService.XpanfileprecreateExecute(r)
}

/*
Xpanfileprecreate Method for Xpanfileprecreate

文件预上传，用于获取上传任务id，既uploadid

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiXpanfileprecreateRequest
*/
func (a *FileuploadApiService) Xpanfileprecreate(ctx _context.Context) ApiXpanfileprecreateRequest {
	return ApiXpanfileprecreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Fileprecreateresponse
func (a *FileuploadApiService) XpanfileprecreateExecute(r ApiXpanfileprecreateRequest) (Fileprecreateresponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  Fileprecreateresponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileuploadApiService.Xpanfileprecreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rest/2.0/xpan/file?method=precreate&openapi=xpansdk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.accessToken == nil {
		return localVarReturnValue, nil, reportError("accessToken is required and must be specified")
	}
	if r.path == nil {
		return localVarReturnValue, nil, reportError("path is required and must be specified")
	}
	if r.isdir == nil {
		return localVarReturnValue, nil, reportError("isdir is required and must be specified")
	}
	if r.size == nil {
		return localVarReturnValue, nil, reportError("size is required and must be specified")
	}
	if r.autoinit == nil {
		return localVarReturnValue, nil, reportError("autoinit is required and must be specified")
	}
	if r.blockList == nil {
		return localVarReturnValue, nil, reportError("blockList is required and must be specified")
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
	localVarHTTPHeaderAccepts := []string{"application/json; charset=UTF-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("path", parameterToString(*r.path, ""))
	localVarFormParams.Add("isdir", parameterToString(*r.isdir, ""))
	localVarFormParams.Add("size", parameterToString(*r.size, ""))
	localVarFormParams.Add("autoinit", parameterToString(*r.autoinit, ""))
	localVarFormParams.Add("block_list", parameterToString(*r.blockList, ""))
	if r.rtype != nil {
		localVarFormParams.Add("rtype", parameterToString(*r.rtype, ""))
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
