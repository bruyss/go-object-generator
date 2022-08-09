package utils

import (
	"testing"
)

func TestNextDigital(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"1", "M0.0"},
		{"2", "M0.1"},
		{"3", "M0.2"},
		{"4", "M0.3"},
		{"5", "M0.4"},
		{"6", "M0.5"},
		{"7", "M0.6"},
		{"8", "M0.7"},
		{"9", "M1.0"},
		{"10", "M1.1"},
		{"11", "M1.2"},
		{"12", "M1.3"},
		{"13", "M1.4"},
		{"14", "M1.5"},
		{"15", "M1.6"},
		{"16", "M1.7"},
		{"17", "M2.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextDigital(); got != tt.want {
				t.Errorf("NextDigital() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextAnalog(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"1", "MW100"},
		{"2", "MW102"},
		{"3", "MW104"},
		{"4", "MW106"},
		{"5", "MW108"},
		{"6", "MW110"},
		{"7", "MW112"},
		{"8", "MW114"},
		{"9", "MW116"},
		{"10", "MW118"},
		{"11", "MW120"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextAnalog(); got != tt.want {
				t.Errorf("NextAnalog() = %v, want %v", got, tt.want)
			}
		})
	}
}
