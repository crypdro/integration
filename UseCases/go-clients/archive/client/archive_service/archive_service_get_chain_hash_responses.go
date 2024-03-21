// Code generated by go-swagger; DO NOT EDIT.

package archive_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/qubic/integration/UseCases/go-clients/archive/models"
)

// ArchiveServiceGetChainHashReader is a Reader for the ArchiveServiceGetChainHash structure.
type ArchiveServiceGetChainHashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArchiveServiceGetChainHashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArchiveServiceGetChainHashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewArchiveServiceGetChainHashDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewArchiveServiceGetChainHashOK creates a ArchiveServiceGetChainHashOK with default headers values
func NewArchiveServiceGetChainHashOK() *ArchiveServiceGetChainHashOK {
	return &ArchiveServiceGetChainHashOK{}
}

/*
ArchiveServiceGetChainHashOK describes a response with status code 200, with default header values.

A successful response.
*/
type ArchiveServiceGetChainHashOK struct {
	Payload *models.PbGetChainHashResponse
}

// IsSuccess returns true when this archive service get chain hash o k response has a 2xx status code
func (o *ArchiveServiceGetChainHashOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this archive service get chain hash o k response has a 3xx status code
func (o *ArchiveServiceGetChainHashOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive service get chain hash o k response has a 4xx status code
func (o *ArchiveServiceGetChainHashOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this archive service get chain hash o k response has a 5xx status code
func (o *ArchiveServiceGetChainHashOK) IsServerError() bool {
	return false
}

// IsCode returns true when this archive service get chain hash o k response a status code equal to that given
func (o *ArchiveServiceGetChainHashOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the archive service get chain hash o k response
func (o *ArchiveServiceGetChainHashOK) Code() int {
	return 200
}

func (o *ArchiveServiceGetChainHashOK) Error() string {
	return fmt.Sprintf("[GET /ticks/{tickNumber}/chain-hash][%d] archiveServiceGetChainHashOK  %+v", 200, o.Payload)
}

func (o *ArchiveServiceGetChainHashOK) String() string {
	return fmt.Sprintf("[GET /ticks/{tickNumber}/chain-hash][%d] archiveServiceGetChainHashOK  %+v", 200, o.Payload)
}

func (o *ArchiveServiceGetChainHashOK) GetPayload() *models.PbGetChainHashResponse {
	return o.Payload
}

func (o *ArchiveServiceGetChainHashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PbGetChainHashResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveServiceGetChainHashDefault creates a ArchiveServiceGetChainHashDefault with default headers values
func NewArchiveServiceGetChainHashDefault(code int) *ArchiveServiceGetChainHashDefault {
	return &ArchiveServiceGetChainHashDefault{
		_statusCode: code,
	}
}

/*
ArchiveServiceGetChainHashDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ArchiveServiceGetChainHashDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this archive service get chain hash default response has a 2xx status code
func (o *ArchiveServiceGetChainHashDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this archive service get chain hash default response has a 3xx status code
func (o *ArchiveServiceGetChainHashDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this archive service get chain hash default response has a 4xx status code
func (o *ArchiveServiceGetChainHashDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this archive service get chain hash default response has a 5xx status code
func (o *ArchiveServiceGetChainHashDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this archive service get chain hash default response a status code equal to that given
func (o *ArchiveServiceGetChainHashDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the archive service get chain hash default response
func (o *ArchiveServiceGetChainHashDefault) Code() int {
	return o._statusCode
}

func (o *ArchiveServiceGetChainHashDefault) Error() string {
	return fmt.Sprintf("[GET /ticks/{tickNumber}/chain-hash][%d] ArchiveService_GetChainHash default  %+v", o._statusCode, o.Payload)
}

func (o *ArchiveServiceGetChainHashDefault) String() string {
	return fmt.Sprintf("[GET /ticks/{tickNumber}/chain-hash][%d] ArchiveService_GetChainHash default  %+v", o._statusCode, o.Payload)
}

func (o *ArchiveServiceGetChainHashDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ArchiveServiceGetChainHashDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
