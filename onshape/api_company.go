/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://cad.onshape.com/appstore/dev-portal): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// CompanyApiService CompanyApi service
type CompanyApiService service

type ApiAddUserToCompanyRequest struct {
	ctx                 context.Context
	ApiService          *CompanyApiService
	cid                 string
	bTCompanyUserParams *BTCompanyUserParams
}

func (r ApiAddUserToCompanyRequest) BTCompanyUserParams(bTCompanyUserParams BTCompanyUserParams) ApiAddUserToCompanyRequest {
	r.bTCompanyUserParams = &bTCompanyUserParams
	return r
}

func (r ApiAddUserToCompanyRequest) Execute() (*BTCompanyUserInfo, *http.Response, error) {
	return r.ApiService.AddUserToCompanyExecute(r)
}

/*
AddUserToCompany Add a user to a company.

Returns the company user info.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cid
	@return ApiAddUserToCompanyRequest
*/
func (a *CompanyApiService) AddUserToCompany(ctx context.Context, cid string) ApiAddUserToCompanyRequest {
	return ApiAddUserToCompanyRequest{
		ApiService: a,
		ctx:        ctx,
		cid:        cid,
	}
}

// Execute executes the request
//
//	@return BTCompanyUserInfo
func (a *CompanyApiService) AddUserToCompanyExecute(r ApiAddUserToCompanyRequest) (*BTCompanyUserInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTCompanyUserInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanyApiService.AddUserToCompany")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/companies/{cid}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTCompanyUserParams == nil {
		return localVarReturnValue, nil, reportError("bTCompanyUserParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bTCompanyUserParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTCompanyUserInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFindCompanyRequest struct {
	ctx        context.Context
	ApiService *CompanyApiService
	uid        *string
	activeOnly *bool
	includeAll *bool
}

func (r ApiFindCompanyRequest) Uid(uid string) ApiFindCompanyRequest {
	r.uid = &uid
	return r
}

func (r ApiFindCompanyRequest) ActiveOnly(activeOnly bool) ApiFindCompanyRequest {
	r.activeOnly = &activeOnly
	return r
}

func (r ApiFindCompanyRequest) IncludeAll(includeAll bool) ApiFindCompanyRequest {
	r.includeAll = &includeAll
	return r
}

func (r ApiFindCompanyRequest) Execute() (*BTListResponseBTCompanyInfo, *http.Response, error) {
	return r.ApiService.FindCompanyExecute(r)
}

/*
FindCompany Get all companies to which the specified user belongs.

If no user is specified, will return all companies associated with the current user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFindCompanyRequest
*/
func (a *CompanyApiService) FindCompany(ctx context.Context) ApiFindCompanyRequest {
	return ApiFindCompanyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BTListResponseBTCompanyInfo
func (a *CompanyApiService) FindCompanyExecute(r ApiFindCompanyRequest) (*BTListResponseBTCompanyInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTListResponseBTCompanyInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanyApiService.FindCompany")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/companies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.uid != nil {
		localVarQueryParams.Add("uid", parameterToString(*r.uid, ""))
	}
	if r.activeOnly != nil {
		localVarQueryParams.Add("activeOnly", parameterToString(*r.activeOnly, ""))
	}
	if r.includeAll != nil {
		localVarQueryParams.Add("includeAll", parameterToString(*r.includeAll, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

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

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTListResponseBTCompanyInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCompanyRequest struct {
	ctx        context.Context
	ApiService *CompanyApiService
	cid        string
}

func (r ApiGetCompanyRequest) Execute() (*BTCompanyInfo, *http.Response, error) {
	return r.ApiService.GetCompanyExecute(r)
}

/*
GetCompany Get company information by company ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cid
	@return ApiGetCompanyRequest
*/
func (a *CompanyApiService) GetCompany(ctx context.Context, cid string) ApiGetCompanyRequest {
	return ApiGetCompanyRequest{
		ApiService: a,
		ctx:        ctx,
		cid:        cid,
	}
}

// Execute executes the request
//
//	@return BTCompanyInfo
func (a *CompanyApiService) GetCompanyExecute(r ApiGetCompanyRequest) (*BTCompanyInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTCompanyInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanyApiService.GetCompany")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/companies/{cid}"
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

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

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTCompanyInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetDocumentsByNameRequest struct {
	ctx        context.Context
	ApiService *CompanyApiService
	cid        string
	name       *string
}

func (r ApiGetDocumentsByNameRequest) Name(name string) ApiGetDocumentsByNameRequest {
	r.name = &name
	return r
}

func (r ApiGetDocumentsByNameRequest) Execute() ([]BTDocumentSummaryInfo, *http.Response, error) {
	return r.ApiService.GetDocumentsByNameExecute(r)
}

/*
GetDocumentsByName Get document by exact document name.

This API can only be called by company admins. Use the `name` field for the exact document name string.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cid
	@return ApiGetDocumentsByNameRequest
*/
func (a *CompanyApiService) GetDocumentsByName(ctx context.Context, cid string) ApiGetDocumentsByNameRequest {
	return ApiGetDocumentsByNameRequest{
		ApiService: a,
		ctx:        ctx,
		cid:        cid,
	}
}

// Execute executes the request
//
//	@return []BTDocumentSummaryInfo
func (a *CompanyApiService) GetDocumentsByNameExecute(r ApiGetDocumentsByNameRequest) ([]BTDocumentSummaryInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []BTDocumentSummaryInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanyApiService.GetDocumentsByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/companies/{cid}/documentsbyname"
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}

	localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

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

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v []BTDocumentSummaryInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRemoveUserFromCompanyRequest struct {
	ctx                context.Context
	ApiService         *CompanyApiService
	cid                string
	uid                string
	removeFromTeams    *bool
	removeDirectShares *bool
}

func (r ApiRemoveUserFromCompanyRequest) RemoveFromTeams(removeFromTeams bool) ApiRemoveUserFromCompanyRequest {
	r.removeFromTeams = &removeFromTeams
	return r
}

func (r ApiRemoveUserFromCompanyRequest) RemoveDirectShares(removeDirectShares bool) ApiRemoveUserFromCompanyRequest {
	r.removeDirectShares = &removeDirectShares
	return r
}

func (r ApiRemoveUserFromCompanyRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.RemoveUserFromCompanyExecute(r)
}

/*
RemoveUserFromCompany Remove a user from a company, company teams, and all the direct shares.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cid
	@param uid
	@return ApiRemoveUserFromCompanyRequest
*/
func (a *CompanyApiService) RemoveUserFromCompany(ctx context.Context, cid string, uid string) ApiRemoveUserFromCompanyRequest {
	return ApiRemoveUserFromCompanyRequest{
		ApiService: a,
		ctx:        ctx,
		cid:        cid,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *CompanyApiService) RemoveUserFromCompanyExecute(r ApiRemoveUserFromCompanyRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanyApiService.RemoveUserFromCompany")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/companies/{cid}/users/{uid}"
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterToString(r.uid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.removeFromTeams != nil {
		localVarQueryParams.Add("removeFromTeams", parameterToString(*r.removeFromTeams, ""))
	}
	if r.removeDirectShares != nil {
		localVarQueryParams.Add("removeDirectShares", parameterToString(*r.removeDirectShares, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

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

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v map[string]interface{}
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateCompanyUserRequest struct {
	ctx                 context.Context
	ApiService          *CompanyApiService
	cid                 string
	uid                 string
	bTCompanyUserParams *BTCompanyUserParams
}

func (r ApiUpdateCompanyUserRequest) BTCompanyUserParams(bTCompanyUserParams BTCompanyUserParams) ApiUpdateCompanyUserRequest {
	r.bTCompanyUserParams = &bTCompanyUserParams
	return r
}

func (r ApiUpdateCompanyUserRequest) Execute() (*BTCompanyUserInfo, *http.Response, error) {
	return r.ApiService.UpdateCompanyUserExecute(r)
}

/*
UpdateCompanyUser Update the company's information for a user.

Returns updated company user info. Global permissions can only be updated by the company admin.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cid
	@param uid
	@return ApiUpdateCompanyUserRequest
*/
func (a *CompanyApiService) UpdateCompanyUser(ctx context.Context, cid string, uid string) ApiUpdateCompanyUserRequest {
	return ApiUpdateCompanyUserRequest{
		ApiService: a,
		ctx:        ctx,
		cid:        cid,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return BTCompanyUserInfo
func (a *CompanyApiService) UpdateCompanyUserExecute(r ApiUpdateCompanyUserRequest) (*BTCompanyUserInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTCompanyUserInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanyApiService.UpdateCompanyUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/companies/{cid}/users/{uid}"
	localVarPath = strings.Replace(localVarPath, "{"+"cid"+"}", url.PathEscape(parameterToString(r.cid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterToString(r.uid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTCompanyUserParams == nil {
		return localVarReturnValue, nil, reportError("bTCompanyUserParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bTCompanyUserParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTCompanyUserInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
