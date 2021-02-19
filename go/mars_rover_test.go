package marsrover

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_singleMarsRover(t *testing.T) {

	Convey("Given a 5x5 plateau", t, func() {

		plateau := Plateau{maxX: 5, maxY: 5}
		startingPosition := Coordinates{x: 1, y: 2}
		marsRover := MarsRover{
			plateau:  plateau,
			heading:  N,
			position: startingPosition,
		}

		Convey("Rover can turn", func() {
			Convey("Left", func() {
				marsRover.acceptCommands(L)
				So(marsRover.heading, ShouldEqual, W)
			})
			Convey("Right", func() {
				marsRover.acceptCommands(R)
				So(marsRover.heading, ShouldEqual, E)
			})
		})

		Convey("Rover can move", func() {

			Convey("single step", func() {
				marsRover.acceptCommands(M)
				So(marsRover.position.y, ShouldEqual, 3)
			})

			Convey("two steps", func() {
				marsRover.acceptCommands(M, M)
				So(marsRover.position.y, ShouldEqual, 4)
			})

			Convey("too far", func() {
				marsRover.acceptCommands(M, M, M, M, M, M, M)
				So(marsRover.status, ShouldEqual, NOK)
			})
		})
	})
}
