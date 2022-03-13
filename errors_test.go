package serial_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/albenik/go-serial/v2"
)

func TestPortError(t *testing.T) {
	testErr := errors.New("test error")
	portErr := serial.NewPortError(serial.OsError, testErr)
	err := error(portErr)
	assert.ErrorIs(t, err, testErr)
	portErr = new(serial.PortError)
	assert.ErrorAs(t, err, &portErr)
	assert.Equal(t, serial.OsError, portErr.Code())
}
