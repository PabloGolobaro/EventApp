package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type EmailService service

type Email struct {
	From    string `json:"from"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	To      string `json:"to"`
	To_name string `json:"to_name"`
	Html    string `json:"html"`
	Text    string `json:"text"`
}

func (a *EmailService) SendEmail(ctx context.Context, authorization string, data *Email) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/smtp/send"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(authorization, "")

	localVarPostBody = &data

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		if err == nil {
			return localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, fmt.Errorf("Error ocured")
	}

	return localVarHttpResponse, nil
}
