/*
 * Onshape REST API
 *
 * The Onshape REST API consumed by all clients.
 *
 * API version: 1.113
 * Contact: api-support@onshape.zendesk.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package onshape

import (
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

// FeatureStudiosApiService FeatureStudiosApi service
type FeatureStudiosApiService service

type apiCreateFeatureStudioRequest struct {
	ctx _context.Context
	apiService *FeatureStudiosApiService
	did string
	wid string
	bTModelElementParams *BTModelElementParams
}


func (r apiCreateFeatureStudioRequest) BTModelElementParams(bTModelElementParams BTModelElementParams) apiCreateFeatureStudioRequest {
	r.bTModelElementParams = &bTModelElementParams
	return r
}

/*
CreateFeatureStudio Create Feature Studio
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param did
 * @param wid
@return apiCreateFeatureStudioRequest
*/
func (a *FeatureStudiosApiService) CreateFeatureStudio(ctx _context.Context, did string, wid string) apiCreateFeatureStudioRequest {
	return apiCreateFeatureStudioRequest{
		apiService: a,
		ctx: ctx,
		did: did,
		wid: wid,
	}
}

/*
Execute executes the request
 @return BTDocumentElementInfo
*/
func (r apiCreateFeatureStudioRequest) Execute() (BTDocumentElementInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BTDocumentElementInfo
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "FeatureStudiosApiService.CreateFeatureStudio")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/featurestudios/d/{did}/w/{wid}"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", _neturl.QueryEscape(parameterToString(r.did, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wid"+"}", _neturl.QueryEscape(parameterToString(r.wid, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	
	if r.bTModelElementParams == nil {
		return localVarReturnValue, nil, reportError("bTModelElementParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bTModelElementParams
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v BTDocumentElementInfo
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiGetFeatureStudioContentsRequest struct {
	ctx _context.Context
	apiService *FeatureStudiosApiService
	did string
	wvm string
	wvmid string
	eid string
}


/*
GetFeatureStudioContents Get Feature Studio Contents.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param did
 * @param wvm
 * @param wvmid
 * @param eid
@return apiGetFeatureStudioContentsRequest
*/
func (a *FeatureStudiosApiService) GetFeatureStudioContents(ctx _context.Context, did string, wvm string, wvmid string, eid string) apiGetFeatureStudioContentsRequest {
	return apiGetFeatureStudioContentsRequest{
		apiService: a,
		ctx: ctx,
		did: did,
		wvm: wvm,
		wvmid: wvmid,
		eid: eid,
	}
}

/*
Execute executes the request
 @return BTFeatureStudioContents2239
*/
func (r apiGetFeatureStudioContentsRequest) Execute() (BTFeatureStudioContents2239, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BTFeatureStudioContents2239
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "FeatureStudiosApiService.GetFeatureStudioContents")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/featurestudios/d/{did}/{wvm}/{wvmid}/e/{eid}"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", _neturl.QueryEscape(parameterToString(r.did, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvm"+"}", _neturl.QueryEscape(parameterToString(r.wvm, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvmid"+"}", _neturl.QueryEscape(parameterToString(r.wvmid, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", _neturl.QueryEscape(parameterToString(r.eid, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	
	

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v BTFeatureStudioContents2239
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiGetFeatureStudioSpecsRequest struct {
	ctx _context.Context
	apiService *FeatureStudiosApiService
	did string
	wvm string
	wvmid string
	eid string
}


/*
GetFeatureStudioSpecs Get Feature Studio Specs
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param did
 * @param wvm
 * @param wvmid
 * @param eid
@return apiGetFeatureStudioSpecsRequest
*/
func (a *FeatureStudiosApiService) GetFeatureStudioSpecs(ctx _context.Context, did string, wvm string, wvmid string, eid string) apiGetFeatureStudioSpecsRequest {
	return apiGetFeatureStudioSpecsRequest{
		apiService: a,
		ctx: ctx,
		did: did,
		wvm: wvm,
		wvmid: wvmid,
		eid: eid,
	}
}

/*
Execute executes the request
 @return BTFeatureSpecsResponse664
*/
func (r apiGetFeatureStudioSpecsRequest) Execute() (BTFeatureSpecsResponse664, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BTFeatureSpecsResponse664
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "FeatureStudiosApiService.GetFeatureStudioSpecs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/featurestudios/d/{did}/{wvm}/{wvmid}/e/{eid}/featurespecs"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", _neturl.QueryEscape(parameterToString(r.did, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvm"+"}", _neturl.QueryEscape(parameterToString(r.wvm, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvmid"+"}", _neturl.QueryEscape(parameterToString(r.wvmid, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", _neturl.QueryEscape(parameterToString(r.eid, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	
	

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v BTFeatureSpecsResponse664
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiUpdateFeatureStudioContentsRequest struct {
	ctx _context.Context
	apiService *FeatureStudiosApiService
	did string
	wvm string
	wvmid string
	eid string
	body *string
}


func (r apiUpdateFeatureStudioContentsRequest) Body(body string) apiUpdateFeatureStudioContentsRequest {
	r.body = &body
	return r
}

/*
UpdateFeatureStudioContents Update Feature Studio contents
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param did
 * @param wvm
 * @param wvmid
 * @param eid
@return apiUpdateFeatureStudioContentsRequest
*/
func (a *FeatureStudiosApiService) UpdateFeatureStudioContents(ctx _context.Context, did string, wvm string, wvmid string, eid string) apiUpdateFeatureStudioContentsRequest {
	return apiUpdateFeatureStudioContentsRequest{
		apiService: a,
		ctx: ctx,
		did: did,
		wvm: wvm,
		wvmid: wvmid,
		eid: eid,
	}
}

/*
Execute executes the request
 @return BTFeatureStudioContents2239
*/
func (r apiUpdateFeatureStudioContentsRequest) Execute() (BTFeatureStudioContents2239, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BTFeatureStudioContents2239
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "FeatureStudiosApiService.UpdateFeatureStudioContents")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/featurestudios/d/{did}/{wvm}/{wvmid}/e/{eid}"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", _neturl.QueryEscape(parameterToString(r.did, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvm"+"}", _neturl.QueryEscape(parameterToString(r.wvm, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvmid"+"}", _neturl.QueryEscape(parameterToString(r.wvmid, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", _neturl.QueryEscape(parameterToString(r.eid, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	
	
	
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v BTFeatureStudioContents2239
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
