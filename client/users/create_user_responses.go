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

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateUserCreated creates a CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {
	return &CreateUserCreated{}
}

/* CreateUserCreated describes a response with status code 201, with default header values.

user created
*/
type CreateUserCreated struct {

	/* id of newly created user
	 */
	XUserID string
}

// IsSuccess returns true when this create user created response has a 2xx status code
func (o *CreateUserCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user created response has a 3xx status code
func (o *CreateUserCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user created response has a 4xx status code
func (o *CreateUserCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user created response has a 5xx status code
func (o *CreateUserCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create user created response a status code equal to that given
func (o *CreateUserCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateUserCreated) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserCreated ", 201)
}

func (o *CreateUserCreated) String() string {
	return fmt.Sprintf("[POST /users][%d] createUserCreated ", 201)
}

func (o *CreateUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-User-ID
	hdrXUserID := response.GetHeader("X-User-ID")

	if hdrXUserID != "" {
		o.XUserID = hdrXUserID
	}

	return nil
}

// NewCreateUserBadRequest creates a CreateUserBadRequest with default headers values
func NewCreateUserBadRequest() *CreateUserBadRequest {
	return &CreateUserBadRequest{}
}

/* CreateUserBadRequest describes a response with status code 400, with default header values.

error
*/
type CreateUserBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create user bad request response has a 2xx status code
func (o *CreateUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user bad request response has a 3xx status code
func (o *CreateUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user bad request response has a 4xx status code
func (o *CreateUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user bad request response has a 5xx status code
func (o *CreateUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create user bad request response a status code equal to that given
func (o *CreateUserBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserBadRequest) String() string {
	return fmt.Sprintf("[POST /users][%d] createUserBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserUnauthorized creates a CreateUserUnauthorized with default headers values
func NewCreateUserUnauthorized() *CreateUserUnauthorized {
	return &CreateUserUnauthorized{}
}

/* CreateUserUnauthorized describes a response with status code 401, with default header values.

error
*/
type CreateUserUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this create user unauthorized response has a 2xx status code
func (o *CreateUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user unauthorized response has a 3xx status code
func (o *CreateUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user unauthorized response has a 4xx status code
func (o *CreateUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user unauthorized response has a 5xx status code
func (o *CreateUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create user unauthorized response a status code equal to that given
func (o *CreateUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateUserUnauthorized) String() string {
	return fmt.Sprintf("[POST /users][%d] createUserUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateUserUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserForbidden creates a CreateUserForbidden with default headers values
func NewCreateUserForbidden() *CreateUserForbidden {
	return &CreateUserForbidden{}
}

/* CreateUserForbidden describes a response with status code 403, with default header values.

error
*/
type CreateUserForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this create user forbidden response has a 2xx status code
func (o *CreateUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user forbidden response has a 3xx status code
func (o *CreateUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user forbidden response has a 4xx status code
func (o *CreateUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user forbidden response has a 5xx status code
func (o *CreateUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create user forbidden response a status code equal to that given
func (o *CreateUserForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateUserForbidden) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserForbidden  %+v", 403, o.Payload)
}

func (o *CreateUserForbidden) String() string {
	return fmt.Sprintf("[POST /users][%d] createUserForbidden  %+v", 403, o.Payload)
}

func (o *CreateUserForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserDefault creates a CreateUserDefault with default headers values
func NewCreateUserDefault(code int) *CreateUserDefault {
	return &CreateUserDefault{
		_statusCode: code,
	}
}

/* CreateUserDefault describes a response with status code -1, with default header values.

error
*/
type CreateUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create user default response
func (o *CreateUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create user default response has a 2xx status code
func (o *CreateUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create user default response has a 3xx status code
func (o *CreateUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create user default response has a 4xx status code
func (o *CreateUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create user default response has a 5xx status code
func (o *CreateUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create user default response a status code equal to that given
func (o *CreateUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateUserDefault) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUser default  %+v", o._statusCode, o.Payload)
}

func (o *CreateUserDefault) String() string {
	return fmt.Sprintf("[POST /users][%d] createUser default  %+v", o._statusCode, o.Payload)
}

func (o *CreateUserDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
