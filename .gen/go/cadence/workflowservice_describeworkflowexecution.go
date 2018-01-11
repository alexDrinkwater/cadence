// Code generated by thriftrw v1.8.0. DO NOT EDIT.
// @generated

package cadence

import (
	"errors"
	"fmt"
	"github.com/uber/cadence/.gen/go/shared"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// WorkflowService_DescribeWorkflowExecution_Args represents the arguments for the WorkflowService.DescribeWorkflowExecution function.
//
// The arguments for DescribeWorkflowExecution are sent and received over the wire as this struct.
type WorkflowService_DescribeWorkflowExecution_Args struct {
	DescribeRequest *shared.DescribeWorkflowExecutionRequest `json:"describeRequest,omitempty"`
}

// ToWire translates a WorkflowService_DescribeWorkflowExecution_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_DescribeWorkflowExecution_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.DescribeRequest != nil {
		w, err = v.DescribeRequest.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _DescribeWorkflowExecutionRequest_Read(w wire.Value) (*shared.DescribeWorkflowExecutionRequest, error) {
	var v shared.DescribeWorkflowExecutionRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a WorkflowService_DescribeWorkflowExecution_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_DescribeWorkflowExecution_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_DescribeWorkflowExecution_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_DescribeWorkflowExecution_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DescribeRequest, err = _DescribeWorkflowExecutionRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_DescribeWorkflowExecution_Args
// struct.
func (v *WorkflowService_DescribeWorkflowExecution_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.DescribeRequest != nil {
		fields[i] = fmt.Sprintf("DescribeRequest: %v", v.DescribeRequest)
		i++
	}

	return fmt.Sprintf("WorkflowService_DescribeWorkflowExecution_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_DescribeWorkflowExecution_Args match the
// provided WorkflowService_DescribeWorkflowExecution_Args.
//
// This function performs a deep comparison.
func (v *WorkflowService_DescribeWorkflowExecution_Args) Equals(rhs *WorkflowService_DescribeWorkflowExecution_Args) bool {
	if !((v.DescribeRequest == nil && rhs.DescribeRequest == nil) || (v.DescribeRequest != nil && rhs.DescribeRequest != nil && v.DescribeRequest.Equals(rhs.DescribeRequest))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "DescribeWorkflowExecution" for this struct.
func (v *WorkflowService_DescribeWorkflowExecution_Args) MethodName() string {
	return "DescribeWorkflowExecution"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *WorkflowService_DescribeWorkflowExecution_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// WorkflowService_DescribeWorkflowExecution_Helper provides functions that aid in handling the
// parameters and return values of the WorkflowService.DescribeWorkflowExecution
// function.
var WorkflowService_DescribeWorkflowExecution_Helper = struct {
	// Args accepts the parameters of DescribeWorkflowExecution in-order and returns
	// the arguments struct for the function.
	Args func(
		describeRequest *shared.DescribeWorkflowExecutionRequest,
	) *WorkflowService_DescribeWorkflowExecution_Args

	// IsException returns true if the given error can be thrown
	// by DescribeWorkflowExecution.
	//
	// An error can be thrown by DescribeWorkflowExecution only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for DescribeWorkflowExecution
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// DescribeWorkflowExecution into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by DescribeWorkflowExecution
	//
	//   value, err := DescribeWorkflowExecution(args)
	//   result, err := WorkflowService_DescribeWorkflowExecution_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from DescribeWorkflowExecution: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*shared.DescribeWorkflowExecutionResponse, error) (*WorkflowService_DescribeWorkflowExecution_Result, error)

	// UnwrapResponse takes the result struct for DescribeWorkflowExecution
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if DescribeWorkflowExecution threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := WorkflowService_DescribeWorkflowExecution_Helper.UnwrapResponse(result)
	UnwrapResponse func(*WorkflowService_DescribeWorkflowExecution_Result) (*shared.DescribeWorkflowExecutionResponse, error)
}{}

func init() {
	WorkflowService_DescribeWorkflowExecution_Helper.Args = func(
		describeRequest *shared.DescribeWorkflowExecutionRequest,
	) *WorkflowService_DescribeWorkflowExecution_Args {
		return &WorkflowService_DescribeWorkflowExecution_Args{
			DescribeRequest: describeRequest,
		}
	}

	WorkflowService_DescribeWorkflowExecution_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.BadRequestError:
			return true
		case *shared.InternalServiceError:
			return true
		case *shared.EntityNotExistsError:
			return true
		default:
			return false
		}
	}

	WorkflowService_DescribeWorkflowExecution_Helper.WrapResponse = func(success *shared.DescribeWorkflowExecutionResponse, err error) (*WorkflowService_DescribeWorkflowExecution_Result, error) {
		if err == nil {
			return &WorkflowService_DescribeWorkflowExecution_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_DescribeWorkflowExecution_Result.BadRequestError")
			}
			return &WorkflowService_DescribeWorkflowExecution_Result{BadRequestError: e}, nil
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_DescribeWorkflowExecution_Result.InternalServiceError")
			}
			return &WorkflowService_DescribeWorkflowExecution_Result{InternalServiceError: e}, nil
		case *shared.EntityNotExistsError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_DescribeWorkflowExecution_Result.EntityNotExistError")
			}
			return &WorkflowService_DescribeWorkflowExecution_Result{EntityNotExistError: e}, nil
		}

		return nil, err
	}
	WorkflowService_DescribeWorkflowExecution_Helper.UnwrapResponse = func(result *WorkflowService_DescribeWorkflowExecution_Result) (success *shared.DescribeWorkflowExecutionResponse, err error) {
		if result.BadRequestError != nil {
			err = result.BadRequestError
			return
		}
		if result.InternalServiceError != nil {
			err = result.InternalServiceError
			return
		}
		if result.EntityNotExistError != nil {
			err = result.EntityNotExistError
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// WorkflowService_DescribeWorkflowExecution_Result represents the result of a WorkflowService.DescribeWorkflowExecution function call.
//
// The result of a DescribeWorkflowExecution execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type WorkflowService_DescribeWorkflowExecution_Result struct {
	// Value returned by DescribeWorkflowExecution after a successful execution.
	Success              *shared.DescribeWorkflowExecutionResponse `json:"success,omitempty"`
	BadRequestError      *shared.BadRequestError                   `json:"badRequestError,omitempty"`
	InternalServiceError *shared.InternalServiceError              `json:"internalServiceError,omitempty"`
	EntityNotExistError  *shared.EntityNotExistsError              `json:"entityNotExistError,omitempty"`
}

// ToWire translates a WorkflowService_DescribeWorkflowExecution_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_DescribeWorkflowExecution_Result) ToWire() (wire.Value, error) {
	var (
		fields [4]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.BadRequestError != nil {
		w, err = v.BadRequestError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.InternalServiceError != nil {
		w, err = v.InternalServiceError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.EntityNotExistError != nil {
		w, err = v.EntityNotExistError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("WorkflowService_DescribeWorkflowExecution_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _DescribeWorkflowExecutionResponse_Read(w wire.Value) (*shared.DescribeWorkflowExecutionResponse, error) {
	var v shared.DescribeWorkflowExecutionResponse
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a WorkflowService_DescribeWorkflowExecution_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_DescribeWorkflowExecution_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_DescribeWorkflowExecution_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_DescribeWorkflowExecution_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _DescribeWorkflowExecutionResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BadRequestError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.InternalServiceError, err = _InternalServiceError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TStruct {
				v.EntityNotExistError, err = _EntityNotExistsError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.BadRequestError != nil {
		count++
	}
	if v.InternalServiceError != nil {
		count++
	}
	if v.EntityNotExistError != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("WorkflowService_DescribeWorkflowExecution_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_DescribeWorkflowExecution_Result
// struct.
func (v *WorkflowService_DescribeWorkflowExecution_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [4]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.BadRequestError != nil {
		fields[i] = fmt.Sprintf("BadRequestError: %v", v.BadRequestError)
		i++
	}
	if v.InternalServiceError != nil {
		fields[i] = fmt.Sprintf("InternalServiceError: %v", v.InternalServiceError)
		i++
	}
	if v.EntityNotExistError != nil {
		fields[i] = fmt.Sprintf("EntityNotExistError: %v", v.EntityNotExistError)
		i++
	}

	return fmt.Sprintf("WorkflowService_DescribeWorkflowExecution_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_DescribeWorkflowExecution_Result match the
// provided WorkflowService_DescribeWorkflowExecution_Result.
//
// This function performs a deep comparison.
func (v *WorkflowService_DescribeWorkflowExecution_Result) Equals(rhs *WorkflowService_DescribeWorkflowExecution_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.BadRequestError == nil && rhs.BadRequestError == nil) || (v.BadRequestError != nil && rhs.BadRequestError != nil && v.BadRequestError.Equals(rhs.BadRequestError))) {
		return false
	}
	if !((v.InternalServiceError == nil && rhs.InternalServiceError == nil) || (v.InternalServiceError != nil && rhs.InternalServiceError != nil && v.InternalServiceError.Equals(rhs.InternalServiceError))) {
		return false
	}
	if !((v.EntityNotExistError == nil && rhs.EntityNotExistError == nil) || (v.EntityNotExistError != nil && rhs.EntityNotExistError != nil && v.EntityNotExistError.Equals(rhs.EntityNotExistError))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "DescribeWorkflowExecution" for this struct.
func (v *WorkflowService_DescribeWorkflowExecution_Result) MethodName() string {
	return "DescribeWorkflowExecution"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *WorkflowService_DescribeWorkflowExecution_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
