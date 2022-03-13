package serial_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/albenik/go-serial/v2"
)

func TestPort_Nil_SetDTR(t *testing.T) {
	err := (*serial.Port)(nil).SetDTR(false)
	var portErr *serial.PortError
	require.ErrorAs(t, err, &portErr)
	assert.Equal(t, serial.PortClosed, portErr.Code())
	assert.Nil(t, portErr.Unwrap())
}

func TestPort_Nil_String(t *testing.T) {
	var p *serial.Port
	assert.Equal(t, "Error: <nil> port instance", p.String())
}

func TestSerial_Operations(t *testing.T) {
	ports, err := serial.GetPortsList()
	require.NoError(t, err)

	if len(ports) == 0 {
		t.Log("No serial ports found")
	}

	for _, p := range ports {
		t.Logf("Found port: %v", p)
	}
}
