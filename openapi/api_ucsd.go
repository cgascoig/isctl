/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-04-17T15:33:06-07:00.
 *
 * API version: 1.0.9-1628
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

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

// UcsdApiService UcsdApi service
type UcsdApiService service

type apiDeleteUcsdBackupInfoRequest struct {
	ctx        _context.Context
	apiService *UcsdApiService
	moid       string
}

/*
DeleteUcsdBackupInfo Delete a 'ucsd.BackupInfo' resource.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param moid The unique Moid identifier of a resource instance.
@return apiDeleteUcsdBackupInfoRequest
*/
func (a *UcsdApiService) DeleteUcsdBackupInfo(ctx _context.Context, moid string) apiDeleteUcsdBackupInfoRequest {
	return apiDeleteUcsdBackupInfoRequest{
		apiService: a,
		ctx:        ctx,
		moid:       moid,
	}
}

/*
Execute executes the request

*/
func (r apiDeleteUcsdBackupInfoRequest) Execute() (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "UcsdApiService.DeleteUcsdBackupInfo")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ucsd/BackupInfos/{Moid}"
	localVarPath = strings.Replace(localVarPath, "{"+"Moid"+"}", _neturl.QueryEscape(parameterToString(r.moid, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		var v Error
		err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type apiGetUcsdBackupInfoByMoidRequest struct {
	ctx        _context.Context
	apiService *UcsdApiService
	moid       string
}

/*
GetUcsdBackupInfoByMoid Read a 'ucsd.BackupInfo' resource.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param moid The unique Moid identifier of a resource instance.
@return apiGetUcsdBackupInfoByMoidRequest
*/
func (a *UcsdApiService) GetUcsdBackupInfoByMoid(ctx _context.Context, moid string) apiGetUcsdBackupInfoByMoidRequest {
	return apiGetUcsdBackupInfoByMoidRequest{
		apiService: a,
		ctx:        ctx,
		moid:       moid,
	}
}

/*
Execute executes the request
 @return UcsdBackupInfo
*/
func (r apiGetUcsdBackupInfoByMoidRequest) Execute() (UcsdBackupInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UcsdBackupInfo
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "UcsdApiService.GetUcsdBackupInfoByMoid")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ucsd/BackupInfos/{Moid}"
	localVarPath = strings.Replace(localVarPath, "{"+"Moid"+"}", _neturl.QueryEscape(parameterToString(r.moid, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "text/csv", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Error
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

type apiGetUcsdBackupInfoListRequest struct {
	ctx         _context.Context
	apiService  *UcsdApiService
	filter      *string
	orderby     *string
	top         *int32
	skip        *int32
	select_     *string
	expand      *string
	apply       *string
	count       *bool
	inlinecount *string
	at          *string
	tags        *string
}

func (r apiGetUcsdBackupInfoListRequest) Filter(filter string) apiGetUcsdBackupInfoListRequest {
	r.filter = &filter
	return r
}

func (r apiGetUcsdBackupInfoListRequest) Orderby(orderby string) apiGetUcsdBackupInfoListRequest {
	r.orderby = &orderby
	return r
}

func (r apiGetUcsdBackupInfoListRequest) Top(top int32) apiGetUcsdBackupInfoListRequest {
	r.top = &top
	return r
}

func (r apiGetUcsdBackupInfoListRequest) Skip(skip int32) apiGetUcsdBackupInfoListRequest {
	r.skip = &skip
	return r
}

func (r apiGetUcsdBackupInfoListRequest) Select_(select_ string) apiGetUcsdBackupInfoListRequest {
	r.select_ = &select_
	return r
}

func (r apiGetUcsdBackupInfoListRequest) Expand(expand string) apiGetUcsdBackupInfoListRequest {
	r.expand = &expand
	return r
}

func (r apiGetUcsdBackupInfoListRequest) Apply(apply string) apiGetUcsdBackupInfoListRequest {
	r.apply = &apply
	return r
}

func (r apiGetUcsdBackupInfoListRequest) Count(count bool) apiGetUcsdBackupInfoListRequest {
	r.count = &count
	return r
}

func (r apiGetUcsdBackupInfoListRequest) Inlinecount(inlinecount string) apiGetUcsdBackupInfoListRequest {
	r.inlinecount = &inlinecount
	return r
}

func (r apiGetUcsdBackupInfoListRequest) At(at string) apiGetUcsdBackupInfoListRequest {
	r.at = &at
	return r
}

func (r apiGetUcsdBackupInfoListRequest) Tags(tags string) apiGetUcsdBackupInfoListRequest {
	r.tags = &tags
	return r
}

/*
GetUcsdBackupInfoList Read a 'ucsd.BackupInfo' resource.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiGetUcsdBackupInfoListRequest
*/
func (a *UcsdApiService) GetUcsdBackupInfoList(ctx _context.Context) apiGetUcsdBackupInfoListRequest {
	return apiGetUcsdBackupInfoListRequest{
		apiService: a,
		ctx:        ctx,
	}
}

/*
Execute executes the request
 @return UcsdBackupInfoResponse
*/
func (r apiGetUcsdBackupInfoListRequest) Execute() (UcsdBackupInfoResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UcsdBackupInfoResponse
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "UcsdApiService.GetUcsdBackupInfoList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ucsd/BackupInfos"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.filter != nil {
		localVarQueryParams.Add("$filter", parameterToString(*r.filter, ""))
	}
	if r.orderby != nil {
		localVarQueryParams.Add("$orderby", parameterToString(*r.orderby, ""))
	}
	if r.top != nil {
		localVarQueryParams.Add("$top", parameterToString(*r.top, ""))
	}
	if r.skip != nil {
		localVarQueryParams.Add("$skip", parameterToString(*r.skip, ""))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("$select", parameterToString(*r.select_, ""))
	}
	if r.expand != nil {
		localVarQueryParams.Add("$expand", parameterToString(*r.expand, ""))
	}
	if r.apply != nil {
		localVarQueryParams.Add("$apply", parameterToString(*r.apply, ""))
	}
	if r.count != nil {
		localVarQueryParams.Add("$count", parameterToString(*r.count, ""))
	}
	if r.inlinecount != nil {
		localVarQueryParams.Add("$inlinecount", parameterToString(*r.inlinecount, ""))
	}
	if r.at != nil {
		localVarQueryParams.Add("at", parameterToString(*r.at, ""))
	}
	if r.tags != nil {
		localVarQueryParams.Add("tags", parameterToString(*r.tags, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/csv", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Error
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
