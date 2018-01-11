// Code generated by thriftrw v1.8.0. DO NOT EDIT.
// @generated

package history

import (
	"errors"
	"fmt"
	"github.com/uber/cadence/.gen/go/shared"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// HistoryService_RespondActivityTaskCompleted_Args represents the arguments for the HistoryService.RespondActivityTaskCompleted function.
//
// The arguments for RespondActivityTaskCompleted are sent and received over the wire as this struct.
type HistoryService_RespondActivityTaskCompleted_Args struct {
	CompleteRequest *RespondActivityTaskCompletedRequest `json:"completeRequest,omitempty"`
}

// ToWire translates a HistoryService_RespondActivityTaskCompleted_Args struct into a Thrift-level intermediate
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
func (v *HistoryService_RespondActivityTaskCompleted_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.CompleteRequest != nil {
		w, err = v.CompleteRequest.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _RespondActivityTaskCompletedRequest_1_Read(w wire.Value) (*RespondActivityTaskCompletedRequest, error) {
	var v RespondActivityTaskCompletedRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a HistoryService_RespondActivityTaskCompleted_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a HistoryService_RespondActivityTaskCompleted_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v HistoryService_RespondActivityTaskCompleted_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *HistoryService_RespondActivityTaskCompleted_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.CompleteRequest, err = _RespondActivityTaskCompletedRequest_1_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a HistoryService_RespondActivityTaskCompleted_Args
// struct.
func (v *HistoryService_RespondActivityTaskCompleted_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.CompleteRequest != nil {
		fields[i] = fmt.Sprintf("CompleteRequest: %v", v.CompleteRequest)
		i++
	}

	return fmt.Sprintf("HistoryService_RespondActivityTaskCompleted_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this HistoryService_RespondActivityTaskCompleted_Args match the
// provided HistoryService_RespondActivityTaskCompleted_Args.
//
// This function performs a deep comparison.
func (v *HistoryService_RespondActivityTaskCompleted_Args) Equals(rhs *HistoryService_RespondActivityTaskCompleted_Args) bool {
	if !((v.CompleteRequest == nil && rhs.CompleteRequest == nil) || (v.CompleteRequest != nil && rhs.CompleteRequest != nil && v.CompleteRequest.Equals(rhs.CompleteRequest))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "RespondActivityTaskCompleted" for this struct.
func (v *HistoryService_RespondActivityTaskCompleted_Args) MethodName() string {
	return "RespondActivityTaskCompleted"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *HistoryService_RespondActivityTaskCompleted_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// HistoryService_RespondActivityTaskCompleted_Helper provides functions that aid in handling the
// parameters and return values of the HistoryService.RespondActivityTaskCompleted
// function.
var HistoryService_RespondActivityTaskCompleted_Helper = struct {
	// Args accepts the parameters of RespondActivityTaskCompleted in-order and returns
	// the arguments struct for the function.
	Args func(
		completeRequest *RespondActivityTaskCompletedRequest,
	) *HistoryService_RespondActivityTaskCompleted_Args

	// IsException returns true if the given error can be thrown
	// by RespondActivityTaskCompleted.
	//
	// An error can be thrown by RespondActivityTaskCompleted only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for RespondActivityTaskCompleted
	// given the error returned by it. The provided error may
	// be nil if RespondActivityTaskCompleted did not fail.
	//
	// This allows mapping errors returned by RespondActivityTaskCompleted into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// RespondActivityTaskCompleted
	//
	//   err := RespondActivityTaskCompleted(args)
	//   result, err := HistoryService_RespondActivityTaskCompleted_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from RespondActivityTaskCompleted: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*HistoryService_RespondActivityTaskCompleted_Result, error)

	// UnwrapResponse takes the result struct for RespondActivityTaskCompleted
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if RespondActivityTaskCompleted threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := HistoryService_RespondActivityTaskCompleted_Helper.UnwrapResponse(result)
	UnwrapResponse func(*HistoryService_RespondActivityTaskCompleted_Result) error
}{}

func init() {
	HistoryService_RespondActivityTaskCompleted_Helper.Args = func(
		completeRequest *RespondActivityTaskCompletedRequest,
	) *HistoryService_RespondActivityTaskCompleted_Args {
		return &HistoryService_RespondActivityTaskCompleted_Args{
			CompleteRequest: completeRequest,
		}
	}

	HistoryService_RespondActivityTaskCompleted_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.BadRequestError:
			return true
		case *shared.InternalServiceError:
			return true
		case *shared.EntityNotExistsError:
			return true
		case *ShardOwnershipLostError:
			return true
		default:
			return false
		}
	}

	HistoryService_RespondActivityTaskCompleted_Helper.WrapResponse = func(err error) (*HistoryService_RespondActivityTaskCompleted_Result, error) {
		if err == nil {
			return &HistoryService_RespondActivityTaskCompleted_Result{}, nil
		}

		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_RespondActivityTaskCompleted_Result.BadRequestError")
			}
			return &HistoryService_RespondActivityTaskCompleted_Result{BadRequestError: e}, nil
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_RespondActivityTaskCompleted_Result.InternalServiceError")
			}
			return &HistoryService_RespondActivityTaskCompleted_Result{InternalServiceError: e}, nil
		case *shared.EntityNotExistsError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_RespondActivityTaskCompleted_Result.EntityNotExistError")
			}
			return &HistoryService_RespondActivityTaskCompleted_Result{EntityNotExistError: e}, nil
		case *ShardOwnershipLostError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_RespondActivityTaskCompleted_Result.ShardOwnershipLostError")
			}
			return &HistoryService_RespondActivityTaskCompleted_Result{ShardOwnershipLostError: e}, nil
		}

		return nil, err
	}
	HistoryService_RespondActivityTaskCompleted_Helper.UnwrapResponse = func(result *HistoryService_RespondActivityTaskCompleted_Result) (err error) {
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
		if result.ShardOwnershipLostError != nil {
			err = result.ShardOwnershipLostError
			return
		}
		return
	}

}

// HistoryService_RespondActivityTaskCompleted_Result represents the result of a HistoryService.RespondActivityTaskCompleted function call.
//
// The result of a RespondActivityTaskCompleted execution is sent and received over the wire as this struct.
type HistoryService_RespondActivityTaskCompleted_Result struct {
	BadRequestError         *shared.BadRequestError      `json:"badRequestError,omitempty"`
	InternalServiceError    *shared.InternalServiceError `json:"internalServiceError,omitempty"`
	EntityNotExistError     *shared.EntityNotExistsError `json:"entityNotExistError,omitempty"`
	ShardOwnershipLostError *ShardOwnershipLostError     `json:"shardOwnershipLostError,omitempty"`
}

// ToWire translates a HistoryService_RespondActivityTaskCompleted_Result struct into a Thrift-level intermediate
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
func (v *HistoryService_RespondActivityTaskCompleted_Result) ToWire() (wire.Value, error) {
	var (
		fields [4]wire.Field
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
	if v.ShardOwnershipLostError != nil {
		w, err = v.ShardOwnershipLostError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}

	if i > 1 {
		return wire.Value{}, fmt.Errorf("HistoryService_RespondActivityTaskCompleted_Result should have at most one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a HistoryService_RespondActivityTaskCompleted_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a HistoryService_RespondActivityTaskCompleted_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v HistoryService_RespondActivityTaskCompleted_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *HistoryService_RespondActivityTaskCompleted_Result) FromWire(w wire.Value) error {
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
		case 4:
			if field.Value.Type() == wire.TStruct {
				v.ShardOwnershipLostError, err = _ShardOwnershipLostError_Read(field.Value)
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
	if v.ShardOwnershipLostError != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("HistoryService_RespondActivityTaskCompleted_Result should have at most one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a HistoryService_RespondActivityTaskCompleted_Result
// struct.
func (v *HistoryService_RespondActivityTaskCompleted_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [4]string
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
	if v.ShardOwnershipLostError != nil {
		fields[i] = fmt.Sprintf("ShardOwnershipLostError: %v", v.ShardOwnershipLostError)
		i++
	}

	return fmt.Sprintf("HistoryService_RespondActivityTaskCompleted_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this HistoryService_RespondActivityTaskCompleted_Result match the
// provided HistoryService_RespondActivityTaskCompleted_Result.
//
// This function performs a deep comparison.
func (v *HistoryService_RespondActivityTaskCompleted_Result) Equals(rhs *HistoryService_RespondActivityTaskCompleted_Result) bool {
	if !((v.BadRequestError == nil && rhs.BadRequestError == nil) || (v.BadRequestError != nil && rhs.BadRequestError != nil && v.BadRequestError.Equals(rhs.BadRequestError))) {
		return false
	}
	if !((v.InternalServiceError == nil && rhs.InternalServiceError == nil) || (v.InternalServiceError != nil && rhs.InternalServiceError != nil && v.InternalServiceError.Equals(rhs.InternalServiceError))) {
		return false
	}
	if !((v.EntityNotExistError == nil && rhs.EntityNotExistError == nil) || (v.EntityNotExistError != nil && rhs.EntityNotExistError != nil && v.EntityNotExistError.Equals(rhs.EntityNotExistError))) {
		return false
	}
	if !((v.ShardOwnershipLostError == nil && rhs.ShardOwnershipLostError == nil) || (v.ShardOwnershipLostError != nil && rhs.ShardOwnershipLostError != nil && v.ShardOwnershipLostError.Equals(rhs.ShardOwnershipLostError))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "RespondActivityTaskCompleted" for this struct.
func (v *HistoryService_RespondActivityTaskCompleted_Result) MethodName() string {
	return "RespondActivityTaskCompleted"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *HistoryService_RespondActivityTaskCompleted_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
