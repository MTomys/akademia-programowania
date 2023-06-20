package ships_test

import (
	"ships"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		a        ships.Point
		b        ships.Point
		expected ships.Point
	}{
		{
			name:     "Zero points added together give zero a point",
			a:        ships.Point{X: 0, Y: 0},
			b:        ships.Point{X: 0, Y: 0},
			expected: ships.Point{X: 0, Y: 0},
		},
		{
			name:     "Positive point reduced gives a negative point",
			a:        ships.Point{X: 2, Y: 2},
			b:        ships.Point{X: -5, Y: -5},
			expected: ships.Point{X: -3, Y: -3},
		},
		{
			name:     "Positive point added to positive point gives a positive point",
			a:        ships.Point{X: 2, Y: 2},
			b:        ships.Point{X: 5, Y: 5},
			expected: ships.Point{X: 7, Y: 7},
		},
	}

	for _, testCase := range testCases {
		t.Run(t.Name(), func(t *testing.T) {
			pointA := ships.Point{testCase.a.X, testCase.a.Y}
			pointB := ships.Point{testCase.b.X, testCase.b.Y}
			result := pointA.Add(pointB)
			if result != testCase.expected {
				t.Errorf("Expected point %v, but got %v", testCase.expected, result)
			}
		})
	}

}

func TestMoveTo(t *testing.T) {
	testsName := "Ship's points should be displaced by the difference between the destination point and ship's first coordinate"
	testCases := []struct {
		receiver ships.Ship
		moveTo   ships.Point
		expected ships.Ship
	}{
		{
			receiver: []ships.Point{{X: 2, Y: 2}, {X: 4, Y: 4}},
			moveTo:   ships.Point{X: 0, Y: 0},
			expected: []ships.Point{{X: 0, Y: 0}, {X: 2, Y: 2}},
		},
		{
			receiver: []ships.Point{{X: 0, Y: 0}, {X: 0, Y: 0}},
			moveTo:   ships.Point{X: 0, Y: 0},
			expected: []ships.Point{{X: 0, Y: 0}, {X: 0, Y: 0}},
		},
		{
			receiver: []ships.Point{{X: 0, Y: 0}, {X: 0, Y: 0}},
			moveTo:   ships.Point{X: -5, Y: -5},
			expected: []ships.Point{{X: -5, Y: -5}, {X: -5, Y: -5}},
		},
		{
			receiver: []ships.Point{{X: -5, Y: -5}, {X: -6, Y: -6}},
			moveTo:   ships.Point{X: -3, Y: -3},
			expected: []ships.Point{{X: -3, Y: -3}, {X: -4, Y: -4}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testsName, func(t *testing.T) {
			result := testCase.receiver.MoveTo(testCase.moveTo)
			if !compareShips(testCase.expected, result) {
				t.Errorf("Expected ship: %v, but got %v", testCase.expected, result)
			}
		})
	}
}

func compareShips(left, right ships.Ship) bool {
	if len(left) != len(right) {
		return false
	}

	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}
