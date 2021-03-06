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

// RevisionsApiService RevisionsApi service
type RevisionsApiService service

type apiEnumerateRevisionsRequest struct {
	ctx _context.Context
	apiService *RevisionsApiService
	cid string
	elementType *int32
	limit *int32
	offset *int32
	latestOnly *bool
	after *int64
}


func (r apiEnumerateRevisionsRequest) ElementType(elementType int32) apiEnumerateRevisionsRequest {
	r.elementType = &elementType
	return r
}

func (r apiEnumerateRevisionsRequest) Limit(limit int32) apiEnumerateRevisionsRequest {
	r.limit = &limit
	return r
}

func (r apiEnumerateRevisionsRequest) Offset(offset int32) apiEnumerateRevisionsRequest {
	r.offset = &offset
	return r
}

func (r apiEnumerateRevisionsRequest) LatestOnly(latestOnly bool) apiEnumerateRevisionsRequest {
	r.latestOnly = &latestOnly
	return r
}

func (r apiEnumerateRevisionsRequest) After(after int64) apiEnumerateRevisionsRequest {
	r.after = &after
	return r
}

/*
EnumerateRevisions Method for EnumerateRevisions
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cid
@return apiEnumerateRevisionsRequest
*/
func (a *RevisionsApiService) EnumerateRevisions(ctx _context.Context, cid string) apiEnumerateRevisionsRequest {
	return apiEnumerateRevisionsRequest{
		apiService: a,
		ctx: ctx,
		cid: cid,
	}
}

/*
Execute executes the request
 @return BTListResponseBTRevisionInfo
*/
func (r apiEnumerateRevisionsRequest) Execute() (BTListResponseBTRevisionInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BTListResponseBTRevisionInfo
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "RevisionsApiService.EnumerateRevisions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/revisions/companies/{cid}"
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", _neturl.QueryEscape(parameterToString(r.cid, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
					
	if r.elementType != nil {
		localVarQueryParams.Add("elementType", parameterToString(*r.elementType, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.latestOnly != nil {
		localVarQueryParams.Add("latestOnly", parameterToString(*r.latestOnly, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
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
			var v BTListResponseBTRevisionInfo
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
type apiGetLatestInDocumentOrCompanyRequest struct {
	ctx _context.Context
	apiService *RevisionsApiService
	cd string
	cdid string
	pnum string
	et *string
}


func (r apiGetLatestInDocumentOrCompanyRequest) Et(et string) apiGetLatestInDocumentOrCompanyRequest {
	r.et = &et
	return r
}

/*
GetLatestInDocumentOrCompany Method for GetLatestInDocumentOrCompany
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cd
 * @param cdid
 * @param pnum
@return apiGetLatestInDocumentOrCompanyRequest
*/
func (a *RevisionsApiService) GetLatestInDocumentOrCompany(ctx _context.Context, cd string, cdid string, pnum string) apiGetLatestInDocumentOrCompanyRequest {
	return apiGetLatestInDocumentOrCompanyRequest{
		apiService: a,
		ctx: ctx,
		cd: cd,
		cdid: cdid,
		pnum: pnum,
	}
}

/*
Execute executes the request
 @return BTRevisionInfo
*/
func (r apiGetLatestInDocumentOrCompanyRequest) Execute() (BTRevisionInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BTRevisionInfo
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "RevisionsApiService.GetLatestInDocumentOrCompany")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/revisions/{cd}/{cdid}/p/{pnum}/latest"
	localVarPath = strings.Replace(localVarPath, "{"+"cd"+"}", _neturl.QueryEscape(parameterToString(r.cd, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cdid"+"}", _neturl.QueryEscape(parameterToString(r.cdid, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pnum"+"}", _neturl.QueryEscape(parameterToString(r.pnum, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	
	
	if r.et != nil {
		localVarQueryParams.Add("et", parameterToString(*r.et, ""))
	}
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
			var v BTRevisionInfo
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
type apiGetRevisionHistoryInCompanyRequest struct {
	ctx _context.Context
	apiService *RevisionsApiService
	cid string
	pnum string
	elementType *string
	fillApprovers *bool
	fillExportPermission *bool
}


func (r apiGetRevisionHistoryInCompanyRequest) ElementType(elementType string) apiGetRevisionHistoryInCompanyRequest {
	r.elementType = &elementType
	return r
}

func (r apiGetRevisionHistoryInCompanyRequest) FillApprovers(fillApprovers bool) apiGetRevisionHistoryInCompanyRequest {
	r.fillApprovers = &fillApprovers
	return r
}

func (r apiGetRevisionHistoryInCompanyRequest) FillExportPermission(fillExportPermission bool) apiGetRevisionHistoryInCompanyRequest {
	r.fillExportPermission = &fillExportPermission
	return r
}

/*
GetRevisionHistoryInCompany Method for GetRevisionHistoryInCompany
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cid
 * @param pnum
@return apiGetRevisionHistoryInCompanyRequest
*/
func (a *RevisionsApiService) GetRevisionHistoryInCompany(ctx _context.Context, cid string, pnum string) apiGetRevisionHistoryInCompanyRequest {
	return apiGetRevisionHistoryInCompanyRequest{
		apiService: a,
		ctx: ctx,
		cid: cid,
		pnum: pnum,
	}
}

/*
Execute executes the request
 @return BTListResponseBTRevisionInfo
*/
func (r apiGetRevisionHistoryInCompanyRequest) Execute() (BTListResponseBTRevisionInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BTListResponseBTRevisionInfo
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "RevisionsApiService.GetRevisionHistoryInCompany")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/revisions/companies/{cid}/partnumber/{pnum}"
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", _neturl.QueryEscape(parameterToString(r.cid, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pnum"+"}", _neturl.QueryEscape(parameterToString(r.pnum, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
			
	if r.elementType != nil {
		localVarQueryParams.Add("elementType", parameterToString(*r.elementType, ""))
	}
	if r.fillApprovers != nil {
		localVarQueryParams.Add("fillApprovers", parameterToString(*r.fillApprovers, ""))
	}
	if r.fillExportPermission != nil {
		localVarQueryParams.Add("fillExportPermission", parameterToString(*r.fillExportPermission, ""))
	}
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
			var v BTListResponseBTRevisionInfo
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
