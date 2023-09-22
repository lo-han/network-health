package ipv4

import (
	entity "network-health/core/entity/device"
	"testing"
)

func Test_IPv4Address(t *testing.T) {
	testCases := []struct {
		name    string
		address string
		err     error
	}{
		{
			name:    "Succesfull IPv4 set",
			address: "192.169.0.0",
			err:     nil,
		},
		{
			name:    "Failed to set invalid IPv4 with non number characters",
			address: "192.169.error.0",
			err:     entity.HealthErrorInvalidAddress("IPv4"),
		},
		{
			name:    "Failed to set invalid too big IPv4",
			address: "192.169.error.0.0",
			err:     entity.HealthErrorInvalidAddress("IPv4"),
		},
		{
			name:    "Failed to set invalid too short IPv4",
			address: "192.169.0",
			err:     entity.HealthErrorInvalidAddress("IPv4"),
		},
		{
			name:    "Failed to set invalid IPv4 numeral value (bigger than 254)",
			address: "192.169.255.0",
			err:     entity.HealthErrorInvalidAddress("IPv4"),
		},
		{
			name:    "Failed to set invalid IPv4 numeral value (smaller than 0)",
			address: "192.169.-1.0",
			err:     entity.HealthErrorInvalidAddress("IPv4"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := NewIPv4Address(testCase.address)

			if err == nil {
				if err != testCase.err {
					t.Error("Test_IPv4Address() failed on success test case")
				}
			}

			if err != nil {
				if err.Error() != testCase.err.Error() {
					t.Errorf("Test_IPv4Address().Err = %s, expected %s", err.Error(), testCase.err.Error())
				}
			}

		})
	}
}
