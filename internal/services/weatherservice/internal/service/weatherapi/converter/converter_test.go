package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	c := NewConverter()

	tests := []struct {
		input    float64
		expected float64
	}{
		{0, 32},
		{10, 50},
		{-10, 14},
	}

	for _, test := range tests {
		actual := c.CelsiusToFahrenheit(test.input)
		assert.Equal(t, test.expected, actual)
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	c := NewConverter()

	tests := []struct {
		input    float64
		expected float64
	}{
		{0, 273.15},
		{10, 283.15},
		{-10, 263.15},
	}

	for _, test := range tests {
		actual := c.CelsiusToKelvin(test.input)
		assert.Equal(t, test.expected, actual)
	}
}
