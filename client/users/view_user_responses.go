// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/gursew/go-usrmgr/models"
)

// ViewUserReader is a Reader for the ViewUser structure.
type ViewUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ViewUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewViewUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewViewUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewViewUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewViewUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewViewUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewViewUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewViewUserOK creates a ViewUserOK with default headers values
func NewViewUserOK() *ViewUserOK {
	return &ViewUserOK{}
}

/* ViewUserOK describes a response with status code 200, with default header values.

successfully view user
*/
type ViewUserOK struct {
	Payload *models.UserView
}

// IsSuccess returns true when this view user o k response has a 2xx status code
func (o *ViewUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this view user o k response has a 3xx status code
func (o *ViewUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this view user o k response has a 4xx status code
func (o *ViewUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this view user o k response has a 5xx status code
func (o *ViewUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this view user o k response a status code equal to that given
func (o *ViewUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *ViewUserOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] viewUserOK  %+v", 200, o.Payload)
}

func (o *ViewUserOK) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] viewUserOK  %+v", 200, o.Payload)
}

func (o *ViewUserOK) GetPayload() *models.UserView {
	return o.Payload
}

func (o *ViewUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewViewUserBadRequest creates a ViewUserBadRequest with default headers values
func NewViewUserBadRequest() *ViewUserBadRequest {
	return &ViewUserBadRequest{}
}

/* ViewUserBadRequest describes a response with status code 400, with default header values.

error
*/
type ViewUserBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this view user bad request response has a 2xx status code
func (o *ViewUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this view user bad request response has a 3xx status code
func (o *ViewUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this view user bad request response has a 4xx status code
func (o *ViewUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this view user bad request response has a 5xx status code
func (o *ViewUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this view user bad request response a status code equal to that given
func (o *ViewUserBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ViewUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] viewUserBadRequest  %+v", 400, o.Payload)
}

func (o *ViewUserBadRequest) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] viewUserBadRequest  %+v", 400, o.Payload)
}

func (o *ViewUserBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ViewUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewViewUserUnauthorized creates a ViewUserUnauthorized with default headers values
func NewViewUserUnauthorized() *ViewUserUnauthorized {
	return &ViewUserUnauthorized{}
}

/* ViewUserUnauthorized describes a response with status code 401, with default header values.

error
*/
type ViewUserUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this view user unauthorized response has a 2xx status code
func (o *ViewUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this view user unauthorized response has a 3xx status code
func (o *ViewUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this view user unauthorized response has a 4xx status code
func (o *ViewUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this view user unauthorized response has a 5xx status code
func (o *ViewUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this view user unauthorized response a status code equal to that given
func (o *ViewUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ViewUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] viewUserUnauthorized  %+v", 401, o.Payload)
}

func (o *ViewUserUnauthorized) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] viewUserUnauthorized  %+v", 401, o.Payload)
}

func (o *ViewUserUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ViewUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewViewUserForbidden creates a ViewUserForbidden with default headers values
func NewViewUserForbidden() *ViewUserForbidden {
	return &ViewUserForbidden{}
}

/* ViewUserForbidden describes a response with status code 403, with default header values.

error
*/
type ViewUserForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this view user forbidden response has a 2xx status code
func (o *ViewUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this view user forbidden response has a 3xx status code
func (o *ViewUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this view user forbidden response has a 4xx status code
func (o *ViewUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this view user forbidden response has a 5xx status code
func (o *ViewUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this view user forbidden response a status code equal to that given
func (o *ViewUserForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ViewUserForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] viewUserForbidden  %+v", 403, o.Payload)
}

func (o *ViewUserForbidden) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] viewUserForbidden  %+v", 403, o.Payload)
}

func (o *ViewUserForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ViewUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewViewUserNotFound creates a ViewUserNotFound with default headers values
func NewViewUserNotFound() *ViewUserNotFound {
	return &ViewUserNotFound{}
}

/* ViewUserNotFound describes a response with status code 404, with default header values.

error
*/
type ViewUserNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this view user not found response has a 2xx status code
func (o *ViewUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this view user not found response has a 3xx status code
func (o *ViewUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this view user not found response has a 4xx status code
func (o *ViewUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this view user not found response has a 5xx status code
func (o *ViewUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this view user not found response a status code equal to that given
func (o *ViewUserNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ViewUserNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] viewUserNotFound  %+v", 404, o.Payload)
}

func (o *ViewUserNotFound) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] viewUserNotFound  %+v", 404, o.Payload)
}

func (o *ViewUserNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ViewUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewViewUserDefault creates a ViewUserDefault with default headers values
func NewViewUserDefault(code int) *ViewUserDefault {
	return &ViewUserDefault{
		_statusCode: code,
	}
}

/* ViewUserDefault describes a response with status code -1, with default header values.

error
*/
type ViewUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the view user default response
func (o *ViewUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this view user default response has a 2xx status code
func (o *ViewUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this view user default response has a 3xx status code
func (o *ViewUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this view user default response has a 4xx status code
func (o *ViewUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this view user default response has a 5xx status code
func (o *ViewUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this view user default response a status code equal to that given
func (o *ViewUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ViewUserDefault) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] viewUser default  %+v", o._statusCode, o.Payload)
}

func (o *ViewUserDefault) String() string {
	return fmt.Sprintf("[GET /users/{id}][%d] viewUser default  %+v", o._statusCode, o.Payload)
}

func (o *ViewUserDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ViewUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
