package services

import (
	"reflect"
	"testing"
	"time"
)

func Test_identifyLimit(t *testing.T) {
	type args struct {
		interval string
		now      time.Time
	}

	now := time.Now()
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"day", args{interval: "day", now: now}, now.AddDate(0, 0, -1)},
		{"week", args{interval: "week", now: now}, now.AddDate(0, 0, -7)},
		{"month", args{interval: "month", now: now}, now.AddDate(0, -1, 0)},
		{"year", args{interval: "year", now: now}, now.AddDate(-1, 0, 0)},
		{"empty string", args{interval: "", now: now}, now.AddDate(-1, 0, 0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := identifyLimit(tt.args.interval, tt.args.now); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("identifyLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isInterval(t *testing.T) {
	type args struct {
		stack  [4]string
		needle string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"interval present in array", args{intervals, "week"}, true},
		{"interval absent in array", args{intervals, "we"}, false},
		{"interval is empty string", args{intervals, ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterval(tt.args.stack, tt.args.needle); got != tt.want {
				t.Errorf("isInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
