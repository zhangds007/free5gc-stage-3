/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsmf_PDUSession

import (
	"context"
	"fmt"
	"free5gc/lib/openapi/common"
	"free5gc/lib/openapi/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type IndividualPDUSessionHSMFApiService service

/*
IndividualPDUSessionHSMFApiService Release
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pduSessionRef PDU session reference
 * @param optional nil or *ReleasePduSessionParamOpts - Optional Parameters:
 * @param "ReleaseData" (optional.Interface of ReleaseData) -  representation of the data to be sent to H-SMF when releasing the PDU session
*/

type ReleasePduSessionParamOpts struct {
	ReleaseData optional.Interface
}

func (a *IndividualPDUSessionHSMFApiService) ReleasePduSession(ctx context.Context, pduSessionRef string, localVarOptionals *ReleasePduSessionParamOpts) (*http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/pdu-sessions/{pduSessionRef}/release"
	localVarPath = strings.Replace(localVarPath, "{"+"pduSessionRef"+"}", fmt.Sprintf("%v", pduSessionRef), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHttpContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHttpContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHttpHeaderAccept := common.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	// body params
	if localVarOptionals != nil && localVarOptionals.ReleaseData.IsSet() {
		localVarOptionalReleaseData, localVarOptionalReleaseDataok := localVarOptionals.ReleaseData.Value().(models.ReleaseData)
		if !localVarOptionalReleaseDataok {
			return nil, common.ReportError("releaseData should be ReleaseData")
		}
		localVarPostBody = &localVarOptionalReleaseData
	}

	r, err := common.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := common.CallAPI(a.client.cfg, r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	apiError := common.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHttpResponse.Status,
	}
	switch localVarHttpResponse.StatusCode {
	case 204:
		return localVarHttpResponse, nil
	case 400:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 403:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 404:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 411:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 413:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 415:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 429:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 500:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 503:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	default:
		return localVarHttpResponse, common.ReportError("%d is not a valid status code in ReleasePduSession", localVarHttpResponse.StatusCode)
	}
}

/*
IndividualPDUSessionHSMFApiService Update (initiated by V-SMF)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pduSessionRef PDU session reference
 * @param hsmfUpdateData representation of the updates to apply to the PDU session
@return HsmfUpdatedData
*/

func (a *IndividualPDUSessionHSMFApiService) UpdatePduSession(ctx context.Context, pduSessionRef string, updatePduSessionRequest models.UpdatePduSessionRequest) (models.UpdatePduSessionResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  models.UpdatePduSessionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/pdu-sessions/{pduSessionRef}/modify"
	localVarPath = strings.Replace(localVarPath, "{"+"pduSessionRef"+"}", fmt.Sprintf("%v", pduSessionRef), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the request Content-Type header
	if updatePduSessionRequest.BinaryDataN1SmInfoFromUe != nil || updatePduSessionRequest.BinaryDataUnknownN1SmInfo != nil {
		localVarHeaderParams["Content-Type"] = "multipart/related"
		localVarPostBody = &updatePduSessionRequest
	} else {
		localVarHeaderParams["Content-Type"] = "application/json"
		localVarPostBody = updatePduSessionRequest.JsonData
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "multipart/related", "application/problem+json"}

	// set Accept header
	localVarHttpHeaderAccept := common.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := common.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := common.CallAPI(a.client.cfg, r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	apiError := common.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHttpResponse.Status,
	}
	switch localVarHttpResponse.StatusCode {
	case 200:
		err = common.Decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
		}
		return localVarReturnValue, localVarHttpResponse, nil
	case 400:
		var v models.UpdatePduSessionErrorResponse
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 403:
		var v models.UpdatePduSessionErrorResponse
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 404:
		var v models.UpdatePduSessionErrorResponse
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 411:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 413:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 415:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 429:
		var v models.ProblemDetails
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 500:
		var v models.UpdatePduSessionErrorResponse
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 503:
		var v models.UpdatePduSessionErrorResponse
		err = common.Decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	default:
		return localVarReturnValue, localVarHttpResponse, common.ReportError("%d is not a valid status code in UpdatePduSession", localVarHttpResponse.StatusCode)
	}
}
