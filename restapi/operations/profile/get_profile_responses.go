// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetProfileOKCode is the HTTP code returned for type GetProfileOK
const GetProfileOKCode int = 200

/*GetProfileOK successful operation

swagger:response getProfileOK
*/
type GetProfileOK struct {
}

// NewGetProfileOK creates GetProfileOK with default headers values
func NewGetProfileOK() *GetProfileOK {

	return &GetProfileOK{}
}

// WriteResponse to the client
func (o *GetProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
