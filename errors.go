package serial

// PortErrorCode is a code to easily identify the type of error
type PortErrorCode int

const (
	// UnknownCode the error code is unknown
	UnknownCode PortErrorCode = iota
	// PortNotFound the requested port doesn't exist
	PortNotFound
	// PortClosed the port is not opened or was closed from go code or it's os handle became invalid
	PortClosed
	// PortBusy the serial port is already in used by another process
	PortBusy
	// InvalidSerialPort the requested port is not a serial port
	InvalidSerialPort
	// PermissionDenied the user doesn't have enough priviledges
	PermissionDenied
	// InvalidSpeed the requested speed is not valid or not supported
	InvalidSpeed
	// InvalidDataBits the number of data bits is not valid or not supported
	InvalidDataBits
	// InvalidParity the selected parity is not valid or not supported
	InvalidParity
	// InvalidStopBits the selected number of stop bits is not valid or not supported
	InvalidStopBits
	// InvalidTimeoutValue Invalid timeout value passed
	InvalidTimeoutValue
	// ErrorEnumeratingPorts an error occurred while listing serial port
	ErrorEnumeratingPorts
	// FunctionNotImplemented the requested function is not implemented
	FunctionNotImplemented
	// OsError Operating system function error
	OsError
	// WriteFailed Port write failed
	WriteFailed
	// ReadFailed Port read failed
	ReadFailed
)

// PortError is a platform independent error type for serial ports
type PortError struct {
	code    PortErrorCode
	wrapped error
}

func NewPortError(code PortErrorCode, err error) *PortError {
	return &PortError{code: code, wrapped: err}
}

func newOSError(err error) *PortError {
	return &PortError{code: OsError, wrapped: err}
}

// EncodedErrorString returns a string explaining the error code
func (e *PortError) EncodedErrorString() string {
	switch e.code {
	case PortNotFound:
		return "port not found"
	case PortClosed:
		return "port is closed"
	case PortBusy:
		return "port busy"
	case InvalidSerialPort:
		return "invalid port"
	case PermissionDenied:
		return "permission denied"
	case InvalidSpeed:
		return "port speed invalid or not supported"
	case InvalidDataBits:
		return "port data bits invalid or not supported"
	case InvalidParity:
		return "port parity invalid or not supported"
	case InvalidStopBits:
		return "port stop bits invalid or not supported"
	case InvalidTimeoutValue:
		return "timeout value invalid or not supported"
	case ErrorEnumeratingPorts:
		return "could not enumerate serial ports"
	case FunctionNotImplemented:
		return "function not implemented"
	case OsError:
		return "operating system error"
	case WriteFailed:
		return "write failed"
	default:
		return "other error"
	}
}

// Error returns the complete error code with details on the cause of the error
func (e *PortError) Error() string {
	if e.wrapped != nil {
		return e.EncodedErrorString() + ": " + e.wrapped.Error()
	}
	return e.EncodedErrorString()
}

// Code returns an identifier for the kind of error occurred
func (e *PortError) Code() PortErrorCode {
	return e.code
}

// Cause returns the cause for the error
// Deprecated: Use go113 error compatible `Unwrap()`
func (e *PortError) Cause() error {
	return e.wrapped
}

// Unwrap returns a wrapped error
func (e *PortError) Unwrap() error {
	return e.wrapped
}
