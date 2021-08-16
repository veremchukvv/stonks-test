// Code generated by go-swagger; DO NOT EDIT.

package stocks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetStocksNotFoundCode is the HTTP code returned for type GetStocksNotFound
const GetStocksNotFoundCode int = 404

/*GetStocksNotFound Stocks not found

swagger:response getStocksNotFound
*/
type GetStocksNotFound struct {
}

// NewGetStocksNotFound creates GetStocksNotFound with default headers values
func NewGetStocksNotFound() *GetStocksNotFound {

	return &GetStocksNotFound{}
}

// WriteResponse to the client
func (o *GetStocksNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetStocksMethodNotAllowedCode is the HTTP code returned for type GetStocksMethodNotAllowed
const GetStocksMethodNotAllowedCode int = 405

/*GetStocksMethodNotAllowed Invalid input

swagger:response getStocksMethodNotAllowed
*/
type GetStocksMethodNotAllowed struct {
}

// NewGetStocksMethodNotAllowed creates GetStocksMethodNotAllowed with default headers values
func NewGetStocksMethodNotAllowed() *GetStocksMethodNotAllowed {

	return &GetStocksMethodNotAllowed{}
}

// WriteResponse to the client
func (o *GetStocksMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}
