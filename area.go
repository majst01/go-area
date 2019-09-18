package area

import (
	"sort"
	"time"
)

// Point represents a value to a given time
type Point struct {
	Timestamp time.Time
	Value     int64
}

// square is a internal representation of two points connected
// where the timestamp of the second point is taken as end time.
// only linear connected points can be expressed in that way
type square struct {
	start time.Time
	end   time.Time
	value int64
}

// Area calculates to integral between a list of points which are connected linear between start and end time
func Area(start, end time.Time, points []Point) int64 {
	// Points must be sorted ascending by timestamp
	sort.Slice(points, func(i, j int) bool {
		return points[i].Timestamp.Before(points[j].Timestamp)
	})

	if len(points) == 1 {
		virtualPoint := Point{
			Timestamp: end,
			Value:     points[0].Value,
		}
		points = append(points, virtualPoint)
	}
	squares := squaresOf(points)
	var area int64
	for _, s := range squares {
		if s.end.Before(start) {
			continue
		}
		if s.start.After(end) {
			continue
		}
		area += areaOf(start, end, s)
	}

	return area
}

func squaresOf(points []Point) []square {
	squares := []square{}
	for i := range points {
		if i == len(points)-1 {
			break
		}
		s := square{
			start: points[i].Timestamp,
			end:   points[i+1].Timestamp,
			value: points[i].Value,
		}
		squares = append(squares, s)
	}
	return squares
}

// areaOf calculates the area of a square between start and end
func areaOf(start, end time.Time, square square) int64 {
	realstart := square.start
	realend := square.end
	if square.start.Before(start) {
		realstart = start
	}
	if square.end.After(end) {
		realend = end
	}

	area := realend.Sub(realstart).Milliseconds() * square.value
	return area / 1000
}
