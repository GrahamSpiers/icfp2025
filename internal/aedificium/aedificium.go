package aedificium

// A Node is a specific room and door.
type Node struct {
	Room int
	Door int
}

func (n Node) Key(size int) int {
	return size*n.Door + n.Room
}

type Edge struct {
	From Node
	To   Node
}

func (e Edge) Key(size int) int {
	return size*size*e.From.Key(size) + e.To.Key(size)
}

// A Library is a collection of rooms connected by doors.
// Each room has a label only part of which is visible.
// A full map contains the information for all rooms.  A minimal map contains
// just enough connections to navigate the library.
type LibMap struct {
	Labels       []int     // Room labels indexed by room index [room].
	Connections  [6][]Node // Connections index by door (0-5), and room.
	StartingRoom int       // What room did we build the library from?
	Minimal      []Edge    // Minimal connections to navigate the library.
}

func (lm *LibMap) Size() int {
	return len(lm.Labels)
}

func (lm *LibMap) Label(room int) int {
	return lm.Labels[room]
}

func (lm *LibMap) VisibleLabel(room int) int {
	return lm.Labels[room] % 4
}

func (lm *LibMap) Connect(from, to Node) {
	lm.Connections[from.Door][from.Room] = to
	lm.Connections[to.Door][to.Room] = from
}

func (lm *LibMap) ConnectedRoom(from Node) Node {
	return lm.Connections[from.Door][from.Room]
}

// Returns the rooms the doors connect to.
func (lm *LibMap) Doors(room int) [6]int {
	var doors [6]int
	for d := range 6 {
		doors[d] = lm.Connections[d][room].Room
	}
	return doors
}

func (lm *LibMap) Id(room int) int {
	var id = lm.Labels[room] % 4
	for d := range 6 {
		id = (id << 2) | (lm.Connections[d][room].Room % 4)
	}
	return id
}

func MakeEmptyLibMap(size int) LibMap {
	return LibMap{
		Labels: make([]int, size),
		Connections: [6][]Node{
			make([]Node, size),
			make([]Node, size),
			make([]Node, size),
			make([]Node, size),
			make([]Node, size),
			make([]Node, size),
		},
		Minimal: []Edge{},
	}
}
