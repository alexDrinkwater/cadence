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

// HistoryService_RecordDecisionTaskStarted_Args represents the arguments for the HistoryService.RecordDecisionTaskStarted function.
//
// The arguments for RecordDecisionTaskStarted are sent and received over the wire as this struct.
type HistoryService_RecordDecisionTaskStarted_Args struct {
	AddRequest *RecordDecisionTaskStartedRequest `json:"addRequest,omitempty"`
}

// ToWire translates a HistoryService_RecordDecisionTaskStarted_Args struct into a Thrift-level intermediate
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
func (v *HistoryService_RecordDecisionTaskStarted_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.AddRequest != nil {
		w, err = v.AddRequest.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _RecordDecisionTaskStartedRequest_Read(w wire.Value) (*RecordDecisionTaskStartedRequest, error) {
	var v RecordDecisionTaskStartedRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a HistoryService_RecordDecisionTaskStarted_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a HistoryService_RecordDecisionTaskStarted_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v HistoryService_RecordDecisionTaskStarted_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *HistoryService_RecordDecisionTaskStarted_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.AddRequest, err = _RecordDecisionTaskStartedRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a HistoryService_RecordDecisionTaskStarted_Args
// struct.
func (v *HistoryService_RecordDecisionTaskStarted_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.AddRequest != nil {
		fields[i] = fmt.Sprintf("AddRequest: %v", v.AddRequest)
		i++
	}

	return fmt.Sprintf("HistoryService_RecordDecisionTaskStarted_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this HistoryService_RecordDecisionTaskStarted_Args match the
// provided HistoryService_RecordDecisionTaskStarted_Args.
//
// This function performs a deep comparison.
func (v *HistoryService_RecordDecisionTaskStarted_Args) Equals(rhs *HistoryService_RecordDecisionTaskStarted_Args) bool {
	if !((v.AddRequest == nil && rhs.AddRequest == nil) || (v.AddRequest != nil && rhs.AddRequest != nil && v.AddRequest.Equals(rhs.AddRequest))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "RecordDecisionTaskStarted" for this struct.
func (v *HistoryService_RecordDecisionTaskStarted_Args) MethodName() string {
	return "RecordDecisionTaskStarted"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *HistoryService_RecordDecisionTaskStarted_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// HistoryService_RecordDecisionTaskStarted_Helper provides functions that aid in handling the
// parameters and return values of the HistoryService.RecordDecisionTaskStarted
// function.
var HistoryService_RecordDecisionTaskStarted_Helper = struct {
	// Args accepts the parameters of RecordDecisionTaskStarted in-order and returns
	// the arguments struct for the function.
	Args func(
		addRequest *RecordDecisionTaskStartedRequest,
	) *HistoryService_RecordDecisionTaskStarted_Args

	// IsException returns true if the given error can be thrown
	// by RecordDecisionTaskStarted.
	//
	// An error can be thrown by RecordDecisionTaskStarted only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for RecordDecisionTaskStarted
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// RecordDecisionTaskStarted into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by RecordDecisionTaskStarted
	//
	//   value, err := RecordDecisionTaskStarted(args)
	//   result, err := HistoryService_RecordDecisionTaskStarted_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from RecordDecisionTaskStarted: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*RecordDecisionTaskStartedResponse, error) (*HistoryService_RecordDecisionTaskStarted_Result, error)

	// UnwrapResponse takes the result struct for RecordDecisionTaskStarted
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if RecordDecisionTaskStarted threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := HistoryService_RecordDecisionTaskStarted_Helper.UnwrapResponse(result)
	UnwrapResponse func(*HistoryService_RecordDecisionTaskStarted_Result) (*RecordDecisionTaskStartedResponse, error)
}{}

func init() {
	HistoryService_RecordDecisionTaskStarted_Helper.Args = func(
		addRequest *RecordDecisionTaskStartedRequest,
	) *HistoryService_RecordDecisionTaskStarted_Args {
		return &HistoryService_RecordDecisionTaskStarted_Args{
			AddRequest: addRequest,
		}
	}

	HistoryService_RecordDecisionTaskStarted_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.BadRequestError:
			return true
		case *shared.InternalServiceError:
			return true
		case *EventAlreadyStartedError:
			return true
		case *shared.EntityNotExistsError:
			return true
		case *ShardOwnershipLostError:
			return true
		default:
			return false
		}
	}

	HistoryService_RecordDecisionTaskStarted_Helper.WrapResponse = func(success *RecordDecisionTaskStartedResponse, err error) (*HistoryService_RecordDecisionTaskStarted_Result, error) {
		if err == nil {
			return &HistoryService_RecordDecisionTaskStarted_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_RecordDecisionTaskStarted_Result.BadRequestError")
			}
			return &HistoryService_RecordDecisionTaskStarted_Result{BadRequestError: e}, nil
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_RecordDecisionTaskStarted_Result.InternalServiceError")
			}
			return &HistoryService_RecordDecisionTaskStarted_Result{InternalServiceError: e}, nil
		case *EventAlreadyStartedError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_RecordDecisionTaskStarted_Result.EventAlreadyStartedError")
			}
			return &HistoryService_RecordDecisionTaskStarted_Result{EventAlreadyStartedError: e}, nil
		case *shared.EntityNotExistsError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_RecordDecisionTaskStarted_Result.EntityNotExistError")
			}
			return &HistoryService_RecordDecisionTaskStarted_Result{EntityNotExistError: e}, nil
		case *ShardOwnershipLostError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_RecordDecisionTaskStarted_Result.ShardOwnershipLostError")
			}
			return &HistoryService_RecordDecisionTaskStarted_Result{ShardOwnershipLostError: e}, nil
		}

		return nil, err
	}
	HistoryService_RecordDecisionTaskStarted_Helper.UnwrapResponse = func(result *HistoryService_RecordDecisionTaskStarted_Result) (success *RecordDecisionTaskStartedResponse, err error) {
		if result.BadRequestError != nil {
			err = result.BadRequestError
			return
		}
		if result.InternalServiceError != nil {
			err = result.InternalServiceError
			return
		}
		if result.EventAlreadyStartedError != nil {
			err = result.EventAlreadyStartedError
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

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// HistoryService_RecordDecisionTaskStarted_Result represents the result of a HistoryService.RecordDecisionTaskStarted function call.
//
// The result of a RecordDecisionTaskStarted execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type HistoryService_RecordDecisionTaskStarted_Result struct {
	// Value returned by RecordDecisionTaskStarted after a successful execution.
	Success                  *RecordDecisionTaskStartedResponse `json:"success,omitempty"`
	BadRequestError          *shared.BadRequestError            `json:"badRequestError,omitempty"`
	InternalServiceError     *shared.InternalServiceError       `json:"internalServiceError,omitempty"`
	EventAlreadyStartedError *EventAlreadyStartedError          `json:"eventAlreadyStartedError,omitempty"`
	EntityNotExistError      *shared.EntityNotExistsError       `json:"entityNotExistError,omitempty"`
	ShardOwnershipLostError  *ShardOwnershipLostError           `json:"shardOwnershipLostError,omitempty"`
}

// ToWire translates a HistoryService_RecordDecisionTaskStarted_Result struct into a Thrift-level intermediate
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
func (v *HistoryService_RecordDecisionTaskStarted_Result) ToWire() (wire.Value, error) {
	var (
		fields [6]wire.Field
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
	if v.EventAlreadyStartedError != nil {
		w, err = v.EventAlreadyStartedError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.EntityNotExistError != nil {
		w, err = v.EntityNotExistError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	if v.ShardOwnershipLostError != nil {
		w, err = v.ShardOwnershipLostError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 5, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("HistoryService_RecordDecisionTaskStarted_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _RecordDecisionTaskStartedResponse_Read(w wire.Value) (*RecordDecisionTaskStartedResponse, error) {
	var v RecordDecisionTaskStartedResponse
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a HistoryService_RecordDecisionTaskStarted_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a HistoryService_RecordDecisionTaskStarted_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v HistoryService_RecordDecisionTaskStarted_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *HistoryService_RecordDecisionTaskStarted_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _RecordDecisionTaskStartedResponse_Read(field.Value)
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
				v.EventAlreadyStartedError, err = _EventAlreadyStartedError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 4:
			if field.Value.Type() == wire.TStruct {
				v.EntityNotExistError, err = _EntityNotExistsError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 5:
			if field.Value.Type() == wire.TStruct {
				v.ShardOwnershipLostError, err = _ShardOwnershipLostError_Read(field.Value)
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
	if v.EventAlreadyStartedError != nil {
		count++
	}
	if v.EntityNotExistError != nil {
		count++
	}
	if v.ShardOwnershipLostError != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("HistoryService_RecordDecisionTaskStarted_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a HistoryService_RecordDecisionTaskStarted_Result
// struct.
func (v *HistoryService_RecordDecisionTaskStarted_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [6]string
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
	if v.EventAlreadyStartedError != nil {
		fields[i] = fmt.Sprintf("EventAlreadyStartedError: %v", v.EventAlreadyStartedError)
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

	return fmt.Sprintf("HistoryService_RecordDecisionTaskStarted_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this HistoryService_RecordDecisionTaskStarted_Result match the
// provided HistoryService_RecordDecisionTaskStarted_Result.
//
// This function performs a deep comparison.
func (v *HistoryService_RecordDecisionTaskStarted_Result) Equals(rhs *HistoryService_RecordDecisionTaskStarted_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.BadRequestError == nil && rhs.BadRequestError == nil) || (v.BadRequestError != nil && rhs.BadRequestError != nil && v.BadRequestError.Equals(rhs.BadRequestError))) {
		return false
	}
	if !((v.InternalServiceError == nil && rhs.InternalServiceError == nil) || (v.InternalServiceError != nil && rhs.InternalServiceError != nil && v.InternalServiceError.Equals(rhs.InternalServiceError))) {
		return false
	}
	if !((v.EventAlreadyStartedError == nil && rhs.EventAlreadyStartedError == nil) || (v.EventAlreadyStartedError != nil && rhs.EventAlreadyStartedError != nil && v.EventAlreadyStartedError.Equals(rhs.EventAlreadyStartedError))) {
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
// This will always be "RecordDecisionTaskStarted" for this struct.
func (v *HistoryService_RecordDecisionTaskStarted_Result) MethodName() string {
	return "RecordDecisionTaskStarted"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *HistoryService_RecordDecisionTaskStarted_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
