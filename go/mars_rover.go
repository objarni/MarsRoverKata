package marsrover

type Coordinates struct {
	x int
	y int
}

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

func (d Direction) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
}

type Command int

const (
	M Command = iota
	L
	R
)

func (c Command) String() string {
	return [...]string{"M", "L", "R"}[c]
}

type Obstacle struct {
	position Coordinates
}

type Plateau struct {
	maxX      int
	maxY      int
	obstacles []Obstacle
}

type Status int

const (
	OK Status = iota
	NOK
)

func (s Status) String() string {
	return [...]string{"OK", "NOK"}[s]
}

type MarsRover struct {
	plateau  Plateau
	heading  Direction
	position Coordinates
	status   Status
}

func (r *MarsRover) acceptCommands(commands ...Command) {
	for _, command := range commands {
		switch command {
		case L:
			r.heading = W
		case R:
			r.heading = E
		case M:
			r.position = Coordinates{x: 1, y: r.position.y + 1}
		}
	}
}
