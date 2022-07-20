/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.150.5616-564f6a8676f1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

// MetadataCategoryApiService MetadataCategoryApi service
type MetadataCategoryApiService service

type ApiGetCategoryPropertiesRequest struct {
	ctx                                      context.Context
	ApiService                               *MetadataCategoryApiService
	ownerId                                  *string
	ownerType                                *int32
	documentId                               *string
	categoryIds                              *[]string
	objectType                               *int32
	strict                                   *bool
	includeObjectTypeDefaults                *bool
	includeComputedProperties                *bool
	includePartPropertiesTableOnlyProperties *bool
	onlyActive                               *bool
	onlyObjectTypeDefaults                   *bool
}

func (r ApiGetCategoryPropertiesRequest) OwnerId(ownerId string) ApiGetCategoryPropertiesRequest {
	r.ownerId = &ownerId
	return r
}

func (r ApiGetCategoryPropertiesRequest) OwnerType(ownerType int32) ApiGetCategoryPropertiesRequest {
	r.ownerType = &ownerType
	return r
}

func (r ApiGetCategoryPropertiesRequest) DocumentId(documentId string) ApiGetCategoryPropertiesRequest {
	r.documentId = &documentId
	return r
}

func (r ApiGetCategoryPropertiesRequest) CategoryIds(categoryIds []string) ApiGetCategoryPropertiesRequest {
	r.categoryIds = &categoryIds
	return r
}

func (r ApiGetCategoryPropertiesRequest) ObjectType(objectType int32) ApiGetCategoryPropertiesRequest {
	r.objectType = &objectType
	return r
}

func (r ApiGetCategoryPropertiesRequest) Strict(strict bool) ApiGetCategoryPropertiesRequest {
	r.strict = &strict
	return r
}

func (r ApiGetCategoryPropertiesRequest) IncludeObjectTypeDefaults(includeObjectTypeDefaults bool) ApiGetCategoryPropertiesRequest {
	r.includeObjectTypeDefaults = &includeObjectTypeDefaults
	return r
}

func (r ApiGetCategoryPropertiesRequest) IncludeComputedProperties(includeComputedProperties bool) ApiGetCategoryPropertiesRequest {
	r.includeComputedProperties = &includeComputedProperties
	return r
}

func (r ApiGetCategoryPropertiesRequest) IncludePartPropertiesTableOnlyProperties(includePartPropertiesTableOnlyProperties bool) ApiGetCategoryPropertiesRequest {
	r.includePartPropertiesTableOnlyProperties = &includePartPropertiesTableOnlyProperties
	return r
}

func (r ApiGetCategoryPropertiesRequest) OnlyActive(onlyActive bool) ApiGetCategoryPropertiesRequest {
	r.onlyActive = &onlyActive
	return r
}

func (r ApiGetCategoryPropertiesRequest) OnlyObjectTypeDefaults(onlyObjectTypeDefaults bool) ApiGetCategoryPropertiesRequest {
	r.onlyObjectTypeDefaults = &onlyObjectTypeDefaults
	return r
}

func (r ApiGetCategoryPropertiesRequest) Execute() (*BTListResponseBTCategoryPropertyInfo, *http.Response, error) {
	return r.ApiService.GetCategoryPropertiesExecute(r)
}

/*
GetCategoryProperties Retrieve category properties for metadata.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCategoryPropertiesRequest
*/
func (a *MetadataCategoryApiService) GetCategoryProperties(ctx context.Context) ApiGetCategoryPropertiesRequest {
	return ApiGetCategoryPropertiesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BTListResponseBTCategoryPropertyInfo
func (a *MetadataCategoryApiService) GetCategoryPropertiesExecute(r ApiGetCategoryPropertiesRequest) (*BTListResponseBTCategoryPropertyInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTListResponseBTCategoryPropertyInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataCategoryApiService.GetCategoryProperties")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metadatacategory/categoryproperties"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ownerId != nil {
		localVarQueryParams.Add("ownerId", parameterToString(*r.ownerId, ""))
	}
	if r.ownerType != nil {
		localVarQueryParams.Add("ownerType", parameterToString(*r.ownerType, ""))
	}
	if r.documentId != nil {
		localVarQueryParams.Add("documentId", parameterToString(*r.documentId, ""))
	}
	if r.categoryIds != nil {
		t := *r.categoryIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("categoryIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("categoryIds", parameterToString(t, "multi"))
		}
	}
	if r.objectType != nil {
		localVarQueryParams.Add("objectType", parameterToString(*r.objectType, ""))
	}
	if r.strict != nil {
		localVarQueryParams.Add("strict", parameterToString(*r.strict, ""))
	}
	if r.includeObjectTypeDefaults != nil {
		localVarQueryParams.Add("includeObjectTypeDefaults", parameterToString(*r.includeObjectTypeDefaults, ""))
	}
	if r.includeComputedProperties != nil {
		localVarQueryParams.Add("includeComputedProperties", parameterToString(*r.includeComputedProperties, ""))
	}
	if r.includePartPropertiesTableOnlyProperties != nil {
		localVarQueryParams.Add("includePartPropertiesTableOnlyProperties", parameterToString(*r.includePartPropertiesTableOnlyProperties, ""))
	}
	if r.onlyActive != nil {
		localVarQueryParams.Add("onlyActive", parameterToString(*r.onlyActive, ""))
	}
	if r.onlyObjectTypeDefaults != nil {
		localVarQueryParams.Add("onlyObjectTypeDefaults", parameterToString(*r.onlyObjectTypeDefaults, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTListResponseBTCategoryPropertyInfo
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}