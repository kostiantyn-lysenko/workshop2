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
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{"day", args{interval: "day", now: now}, now.AddDate(0, 0, -1), false},
		{"week", args{interval: "week", now: now}, now.AddDate(0, 0, -7), false},
		{"month", args{interval: "month", now: now}, now.AddDate(0, -1, 0), false},
		{"year", args{interval: "year", now: now}, now.AddDate(-1, 0, 0), false},
		{"empty string", args{interval: "", now: now}, now.AddDate(-1, 0, 0), true},
		{"value isn't contained in intervals", args{interval: "1e", now: now}, now.AddDate(-1, 0, 0), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := identifyLimit(tt.args.interval, tt.args.now)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("identifyLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("identifyLimit() got = %v, want %v", got, tt.want)
			}
		})
	}
}
