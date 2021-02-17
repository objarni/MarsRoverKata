package marsrover

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)







func TestCanRotateLeft(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	startingPosition := Coordinates{1,2}
	marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}

	marsRover.turnLeft()
	if marsRover.heading != W {
		t.Fail()
	}
}

func TestConveyCanRotateLeft(t *testing.T) {
	Convey("Rovers can turn", t, func() {
		plateau := Plateau{maxX: 5, maxY: 5}
		startingPosition := Coordinates{1,2}
		marsRover := MarsRover{plateau: plateau, heading: N, position: startingPosition}
		Convey("Left", func() {
			marsRover.turnLeft()
			So(marsRover.heading, ShouldEqual, W)
		})
	})
}

func TestConvey(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}
