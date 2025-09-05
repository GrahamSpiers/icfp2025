package aedificium

import (
	"math/rand"
	"slices"
)

type Builder struct {
	Size      int
	UsedDoors [][]int // List of used doors for each room.
	Connected []int   // List of connected rooms.
}

func NewBuilder(size int) *Builder {
	b := &Builder{
		Size:      size,
		UsedDoors: make([][]int, size),
		Connected: []int{},
	}
	for room := range size {
		b.UsedDoors[room] = []int{}
	}
	return b
}

func (b *Builder) UnusedDoor(room int) int {
	//fmt.Printf("%d ", room)
	//fmt.Printf("%+v\n", b.UsedDoors[room])
	if len(b.UsedDoors[room]) == 0 {
		return rand.Intn(6)
	}
	for {
		door := rand.Intn(6)
		if !slices.Contains(b.UsedDoors[room], door) {
			b.UsedDoors[room] = append(b.UsedDoors[room], door)
			return door
		}
	}
}

func (b *Builder) IsDoorUsed(room, door int) bool {
	return slices.Contains(b.UsedDoors[room], door)
}

// Pick a room from the connected set that has at least one unused door.
func (b *Builder) PickConnected() int {
	n := len(b.Connected)
	offset := rand.Intn(n)
	for i := range b.Size {
		j := (i + offset) % n
		room := b.Connected[j]
		if len(b.UsedDoors[room]) < 6 {
			return room
		}
	}
	//fmt.Printf("no room in %v", b.Connected)
	return -1
}

func (b *Builder) AddConnected(room int) {
	if !slices.Contains(b.Connected, room) {
		b.Connected = append(b.Connected, room)
	}
}

func (b *Builder) UseDoor(room, door int) {
	if !b.IsDoorUsed(room, door) {
		b.UsedDoors[room] = append(b.UsedDoors[room], door)
	}
}

func (b *Builder) RandomUnusedNode() Node {
	for {
		room := rand.Intn(b.Size)
		if len(b.UsedDoors[room]) == 6 {
			continue
		}
		// Just pick the first...
		for door := range 6 {
			if !b.IsDoorUsed(room, door) {
				return Node{Room: room, Door: door}
			}
		}
	}
}

// Build creates a library of the given size with all doors
// leading to random rooms.  It assigns a random path to all rooms.  Note
// that there may be more that one path throught the complete library.
// It returns the Library and a path that visits all rooms.
func (b *Builder) Build(size int) LibMap {
	//fmt.Printf("size %d\n", size)
	lib := MakeEmptyLibMap(size)
	if size == 0 {
		return lib
	}
	for i := range size {
		lib.Labels[i] = i
		for d := range 6 {
			lib.Connections[d][i] = Node{Room: i, Door: d}
		}
	}
	// At this point we have a legal map with all rooms isolated.
	// We need to add the minimal connections to connect all rooms.
	// Create a starting room.
	lib.StartingRoom = rand.Intn(size)
	b.Connected = []int{lib.StartingRoom}
	for to := range size {
		//fmt.Printf("to %d from %v\n", to, b.Connected)
		// Find a connectable room.
		from := b.PickConnected()
		//fmt.Printf("from %d doors %v\n", from, b.UsedDoors[from])
		fromDoor := b.UnusedDoor(from)
		toDoor := b.UnusedDoor(to)
		lib.Connect(Node{Room: from, Door: fromDoor}, Node{Room: to, Door: toDoor})
		b.AddConnected(to)
		b.UseDoor(from, fromDoor)
		b.UseDoor(to, toDoor)
		connection := Edge{Node{from, fromDoor}, Node{to, toDoor}}
		lib.Minimal = append(lib.Minimal, connection)
	}
	// Randomize connections.
	for from := range size {
		for fromDoor := range 6 {
			if slices.Contains(b.UsedDoors[from], fromDoor) {
				continue
			}
			// This door is unused.
			// Pick a random room/door.  It may be the same room or same room
			// and door.
			to := b.RandomUnusedNode()
			lib.Connect(Node{Room: from, Door: fromDoor}, to)
			b.UseDoor(from, fromDoor)
			b.UseDoor(to.Room, to.Door)
			if len(b.UsedDoors[from]) == 6 {
				break
			}
		}
	}
	return lib
}

func MakeLibMap(size int) LibMap {
	builder := NewBuilder(size)
	return builder.Build(size)
}
