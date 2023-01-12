package geofox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSignature(t *testing.T) {
	cases := []struct {
		name              string
		inputMessage      []byte
		inputKey          []byte
		expectedSignature string
	}{
		{
			name:              "empty message - empty key",
			inputMessage:      []byte{},
			inputKey:          []byte{},
			expectedSignature: "+9sdGxiqbAgyS31ktx+3Y3BpDh0=",
		},
		{
			name:              "nil message - empty key",
			inputMessage:      nil,
			inputKey:          []byte{},
			expectedSignature: "+9sdGxiqbAgyS31ktx+3Y3BpDh0=",
		},
		{
			name:              "empty message - nil key",
			inputMessage:      []byte{},
			inputKey:          nil,
			expectedSignature: "+9sdGxiqbAgyS31ktx+3Y3BpDh0=",
		},
		{
			name:              "apple message - empty key",
			inputMessage:      []byte("apple"),
			inputKey:          []byte{},
			expectedSignature: "owHgjIirOQbvpNrkZ/PMomgjgno=",
		},
		{
			name:              "apple message - nil key",
			inputMessage:      []byte("apple"),
			inputKey:          nil,
			expectedSignature: "owHgjIirOQbvpNrkZ/PMomgjgno=",
		},
		{
			name:              "empty message - apple key",
			inputMessage:      []byte{},
			inputKey:          []byte("apple"),
			expectedSignature: "OlIRfD63z/DxE8x/5DT5e41ksJs=",
		},
		{
			name:              "nil message - apple key",
			inputMessage:      nil,
			inputKey:          []byte("apple"),
			expectedSignature: "OlIRfD63z/DxE8x/5DT5e41ksJs=",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			res := createSHA1HmacSignature(tt.inputMessage, tt.inputKey)
			assert.Equal(t, tt.expectedSignature, res)
		})
	}
}
