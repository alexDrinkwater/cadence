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

// WorkflowService_RespondActivityTaskCanceledByID_Args represents the arguments for the WorkflowService.RespondActivityTaskCanceledByID function.
//
// The arguments for RespondActivityTaskCanceledByID are sent and received over the wire as this struct.
type WorkflowService_RespondActivityTaskCanceledByID_Args struct {
	CanceledRequest *shared.RespondActivityTaskCanceledByIDRequest `json:"canceledRequest,omitempty"`
}

// ToWire translates a WorkflowService_RespondActivityTaskCanceledByID_Args struct into a Thrift-level intermediate
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
func (v *WorkflowService_RespondActivityTaskCanceledByID_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.CanceledRequest != nil {
		w, err = v.CanceledRequest.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _RespondActivityTaskCanceledByIDRequest_Read(w wire.Value) (*shared.RespondActivityTaskCanceledByIDRequest, error) {
	var v shared.RespondActivityTaskCanceledByIDRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a WorkflowService_RespondActivityTaskCanceledByID_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_RespondActivityTaskCanceledByID_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_RespondActivityTaskCanceledByID_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_RespondActivityTaskCanceledByID_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.CanceledRequest, err = _RespondActivityTaskCanceledByIDRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_RespondActivityTaskCanceledByID_Args
// struct.
func (v *WorkflowService_RespondActivityTaskCanceledByID_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.CanceledRequest != nil {
		fields[i] = fmt.Sprintf("CanceledRequest: %v", v.CanceledRequest)
		i++
	}

	return fmt.Sprintf("WorkflowService_RespondActivityTaskCanceledByID_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_RespondActivityTaskCanceledByID_Args match the
// provided WorkflowService_RespondActivityTaskCanceledByID_Args.
//
// This function performs a deep comparison.
func (v *WorkflowService_RespondActivityTaskCanceledByID_Args) Equals(rhs *WorkflowService_RespondActivityTaskCanceledByID_Args) bool {
	if !((v.CanceledRequest == nil && rhs.CanceledRequest == nil) || (v.CanceledRequest != nil && rhs.CanceledRequest != nil && v.CanceledRequest.Equals(rhs.CanceledRequest))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "RespondActivityTaskCanceledByID" for this struct.
func (v *WorkflowService_RespondActivityTaskCanceledByID_Args) MethodName() string {
	return "RespondActivityTaskCanceledByID"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *WorkflowService_RespondActivityTaskCanceledByID_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// WorkflowService_RespondActivityTaskCanceledByID_Helper provides functions that aid in handling the
// parameters and return values of the WorkflowService.RespondActivityTaskCanceledByID
// function.
var WorkflowService_RespondActivityTaskCanceledByID_Helper = struct {
	// Args accepts the parameters of RespondActivityTaskCanceledByID in-order and returns
	// the arguments struct for the function.
	Args func(
		canceledRequest *shared.RespondActivityTaskCanceledByIDRequest,
	) *WorkflowService_RespondActivityTaskCanceledByID_Args

	// IsException returns true if the given error can be thrown
	// by RespondActivityTaskCanceledByID.
	//
	// An error can be thrown by RespondActivityTaskCanceledByID only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for RespondActivityTaskCanceledByID
	// given the error returned by it. The provided error may
	// be nil if RespondActivityTaskCanceledByID did not fail.
	//
	// This allows mapping errors returned by RespondActivityTaskCanceledByID into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// RespondActivityTaskCanceledByID
	//
	//   err := RespondActivityTaskCanceledByID(args)
	//   result, err := WorkflowService_RespondActivityTaskCanceledByID_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from RespondActivityTaskCanceledByID: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*WorkflowService_RespondActivityTaskCanceledByID_Result, error)

	// UnwrapResponse takes the result struct for RespondActivityTaskCanceledByID
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if RespondActivityTaskCanceledByID threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := WorkflowService_RespondActivityTaskCanceledByID_Helper.UnwrapResponse(result)
	UnwrapResponse func(*WorkflowService_RespondActivityTaskCanceledByID_Result) error
}{}

func init() {
	WorkflowService_RespondActivityTaskCanceledByID_Helper.Args = func(
		canceledRequest *shared.RespondActivityTaskCanceledByIDRequest,
	) *WorkflowService_RespondActivityTaskCanceledByID_Args {
		return &WorkflowService_RespondActivityTaskCanceledByID_Args{
			CanceledRequest: canceledRequest,
		}
	}

	WorkflowService_RespondActivityTaskCanceledByID_Helper.IsException = func(err error) bool {
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

	WorkflowService_RespondActivityTaskCanceledByID_Helper.WrapResponse = func(err error) (*WorkflowService_RespondActivityTaskCanceledByID_Result, error) {
		if err == nil {
			return &WorkflowService_RespondActivityTaskCanceledByID_Result{}, nil
		}

		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_RespondActivityTaskCanceledByID_Result.BadRequestError")
			}
			return &WorkflowService_RespondActivityTaskCanceledByID_Result{BadRequestError: e}, nil
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_RespondActivityTaskCanceledByID_Result.InternalServiceError")
			}
			return &WorkflowService_RespondActivityTaskCanceledByID_Result{InternalServiceError: e}, nil
		case *shared.EntityNotExistsError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_RespondActivityTaskCanceledByID_Result.EntityNotExistError")
			}
			return &WorkflowService_RespondActivityTaskCanceledByID_Result{EntityNotExistError: e}, nil
		}

		return nil, err
	}
	WorkflowService_RespondActivityTaskCanceledByID_Helper.UnwrapResponse = func(result *WorkflowService_RespondActivityTaskCanceledByID_Result) (err error) {
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
		return
	}

}

// WorkflowService_RespondActivityTaskCanceledByID_Result represents the result of a WorkflowService.RespondActivityTaskCanceledByID function call.
//
// The result of a RespondActivityTaskCanceledByID execution is sent and received over the wire as this struct.
type WorkflowService_RespondActivityTaskCanceledByID_Result struct {
	BadRequestError      *shared.BadRequestError      `json:"badRequestError,omitempty"`
	InternalServiceError *shared.InternalServiceError `json:"internalServiceError,omitempty"`
	EntityNotExistError  *shared.EntityNotExistsError `json:"entityNotExistError,omitempty"`
}

// ToWire translates a WorkflowService_RespondActivityTaskCanceledByID_Result struct into a Thrift-level intermediate
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
func (v *WorkflowService_RespondActivityTaskCanceledByID_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

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

	if i > 1 {
		return wire.Value{}, fmt.Errorf("WorkflowService_RespondActivityTaskCanceledByID_Result should have at most one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a WorkflowService_RespondActivityTaskCanceledByID_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_RespondActivityTaskCanceledByID_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_RespondActivityTaskCanceledByID_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_RespondActivityTaskCanceledByID_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
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
	if v.BadRequestError != nil {
		count++
	}
	if v.InternalServiceError != nil {
		count++
	}
	if v.EntityNotExistError != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("WorkflowService_RespondActivityTaskCanceledByID_Result should have at most one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_RespondActivityTaskCanceledByID_Result
// struct.
func (v *WorkflowService_RespondActivityTaskCanceledByID_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
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

	return fmt.Sprintf("WorkflowService_RespondActivityTaskCanceledByID_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_RespondActivityTaskCanceledByID_Result match the
// provided WorkflowService_RespondActivityTaskCanceledByID_Result.
//
// This function performs a deep comparison.
func (v *WorkflowService_RespondActivityTaskCanceledByID_Result) Equals(rhs *WorkflowService_RespondActivityTaskCanceledByID_Result) bool {
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
// This will always be "RespondActivityTaskCanceledByID" for this struct.
func (v *WorkflowService_RespondActivityTaskCanceledByID_Result) MethodName() string {
	return "RespondActivityTaskCanceledByID"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *WorkflowService_RespondActivityTaskCanceledByID_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
