// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"livepeer/models/components"
	"net/http"
)

type GetSigningKeyRequest struct {
	// ID of the signing key
	KeyID string `pathParam:"style=simple,explode=false,name=keyId"`
}

func (o *GetSigningKeyRequest) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

type GetSigningKeyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	SigningKey *components.SigningKey
}

func (o *GetSigningKeyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSigningKeyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSigningKeyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSigningKeyResponse) GetSigningKey() *components.SigningKey {
	if o == nil {
		return nil
	}
	return o.SigningKey
}
