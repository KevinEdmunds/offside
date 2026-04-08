package main

import (
	"offside/internal/fpl"
	"testing"
)

func TestPositionLabel(t *testing.T) {
	tests := []struct {
		elementType int
		want        string
	}{
		{1, "GK"},
		{2, "DEF"},
		{3, "MID"},
		{4, "FWD"},
		{5, "UNK"},
	}

	for _, tt := range tests {
		player := fpl.Player{ElementType: tt.elementType}
		got := player.PositionLabel()
		if got != tt.want {
			t.Errorf("positionLabel(%d) = %s, want %s",
				tt.elementType, got, tt.want)
		}
	}
}