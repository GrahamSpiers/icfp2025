package aedificium

import "math/rand"

// A Room has a label and 6 doors.
type Room struct {
	// Each room has a label 0-3.
	Label int
	// Each room has 6 doors which each lead to another room or the same room.
	Doors [6]int
}

// What is the visible label of this room?
func (r Room) Bits() int {
	return r.Label % 4
}

// A Library is a collection of connected rooms.
type Library []Room

// A Path is a sequence of door index choices.
type Path []int

// A Result is a sequence of room labels.
type Result []int

// MakeSimpleLibrary creates a library of the given size with all doors
// but one leading to room 0 the other door leads to the next room in the
// sequence, wrapping around to room 0 from the last room.
func MakeSimpleLibrary(size int) Library {
	rooms := make([]Room, size)
	for i := range rooms {
		rooms[i].Label = i
		for j := range rooms[i].Doors {
			rooms[i].Doors[j] = 0
		}
		rooms[i].Doors[0] = (i + 1) % size
	}
	return rooms
}

// MakeComplexLibrary creates a library of the given size with all doors
// leading to random rooms.  It assigns a random path to all rooms.  Note
// that there may be more that one path throught the complete library.
func MakeComplexLibrary(size int) Library {
	rooms := make([]Room, size)
	for i := range rooms {
		rooms[i].Label = i
		for j := range rooms[i].Doors {
			rooms[i].Doors[j] = rand.Intn(size)
		}
		rooms[i].Doors[rand.Intn(6)] = (i + 1) % size
	}
	return rooms
}
