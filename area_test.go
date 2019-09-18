package area

import (
	"testing"
	"time"
)

func TestArea(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name   string
		start  time.Time
		end    time.Time
		points []Point
		want   int64
	}{
		{
			name:  "two points inside timewindow",
			start: now,
			end:   now.Add(10 * time.Second),
			points: []Point{
				{Timestamp: now, Value: 100},
				{Timestamp: now.Add(10 * time.Second), Value: 100},
			},
			want: int64(1000),
		},
		{
			name:  "there points inside timewindow",
			start: now,
			end:   now.Add(10 * time.Second),
			points: []Point{
				{Timestamp: now, Value: 100},
				{Timestamp: now.Add(5 * time.Second), Value: 200},
				{Timestamp: now.Add(10 * time.Second), Value: 200},
			},
			want: int64(1500),
		},
		{
			name:  "four points inside timewindow",
			start: now,
			end:   now.Add(10 * time.Second),
			points: []Point{
				{Timestamp: now, Value: 100},
				{Timestamp: now.Add(2 * time.Second), Value: 200},  // 100 * 2sec
				{Timestamp: now.Add(4 * time.Second), Value: 50},   // 200 * 2 sec
				{Timestamp: now.Add(10 * time.Second), Value: 200}, // 50 * 6 sec
			},
			want: int64(900),
		},
		{
			name:  "two points one before one inside timewindow",
			start: now,
			end:   now.Add(10 * time.Second),
			points: []Point{
				{Timestamp: now.Add(-2 * time.Second), Value: 100},
				{Timestamp: now.Add(2 * time.Second), Value: 200}, // 100 * 2sec
			},
			want: int64(200),
		},
		{
			name:  "two points one before one after timewindow",
			start: now,
			end:   now.Add(10 * time.Second),
			points: []Point{
				{Timestamp: now.Add(-2 * time.Second), Value: 100},
				{Timestamp: now.Add(12 * time.Second), Value: 200}, // 100 * 10sec
			},
			want: int64(1000),
		},
		{
			name:  "three points one before one inside and one after timewindow",
			start: now,
			end:   now.Add(10 * time.Second),
			points: []Point{
				{Timestamp: now.Add(-2 * time.Second), Value: 100},
				{Timestamp: now.Add(5 * time.Second), Value: 200},  // 100 * 5sec
				{Timestamp: now.Add(12 * time.Second), Value: 200}, // 200 * 5sec
			},
			want: int64(1500),
		},
		{
			name:  "two points outside of timewindow",
			start: now,
			end:   now.Add(10 * time.Second),
			points: []Point{
				{Timestamp: now.Add(-20 * time.Second), Value: 100},
				{Timestamp: now.Add(-10 * time.Second), Value: 200},
			},
			want: int64(0),
		},
		{
			name:  "one points inside the timewindow",
			start: now,
			end:   now.Add(10 * time.Second),
			points: []Point{
				{Timestamp: now.Add(5 * time.Second), Value: 100}, // 5 * 100
			},
			want: int64(500),
		},
		{
			name:  "one points before the timewindow",
			start: now,
			end:   now.Add(10 * time.Second),
			points: []Point{
				{Timestamp: now.Add(-5 * time.Second), Value: 100}, // 10 * 100
			},
			want: int64(1000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Area(tt.start, tt.end, tt.points); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}
