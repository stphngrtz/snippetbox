package main

import (
	"testing"
	"time"

	"github.com/stphngrtz/snippetbox/internal/assert"
)

func TestHumanData(t *testing.T) {
	tm := time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC)
	hd := humanDate(tm)
	want := "17 Mar 2024 at 10:15"
	if hd != want {
		t.Errorf("got %q; want %q", hd, want)
	}
}

func TestHumanDataTableDriven(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2024 at 10:15",
		}, {
			name: "Emtpy",
			tm:   time.Time{},
			want: "",
		}, {
			name: "CET",
			tm:   time.Date(2024, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2024 at 09:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)
			assert.Equal(t, hd, tt.want)
		})
	}
}
