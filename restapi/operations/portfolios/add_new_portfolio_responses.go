// Code generated by go-swagger; DO NOT EDIT.

package portfolios

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddNewPortfolioMethodNotAllowedCode is the HTTP code returned for type AddNewPortfolioMethodNotAllowed
const AddNewPortfolioMethodNotAllowedCode int = 405

/*AddNewPortfolioMethodNotAllowed Invalid input

swagger:response addNewPortfolioMethodNotAllowed
*/
type AddNewPortfolioMethodNotAllowed struct {
}

// NewAddNewPortfolioMethodNotAllowed creates AddNewPortfolioMethodNotAllowed with default headers values
func NewAddNewPortfolioMethodNotAllowed() *AddNewPortfolioMethodNotAllowed {

	return &AddNewPortfolioMethodNotAllowed{}
}

// WriteResponse to the client
func (o *AddNewPortfolioMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}
