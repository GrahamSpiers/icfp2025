package aedificium

import (
	"fmt"
	"strings"
	"testing"
)

func TestMakeLibMap(t *testing.T) {
	for size := 1; size <= 4; size++ {
		lib := MakeLibMap(size)
		if lib.Size() != size {
			t.Errorf("len(Rooms) = %d, want %d", lib.Size(), size)
		}
		if len(lib.Minimal) != size {
			t.Errorf("len(Path) = %d, want %d", len(lib.Minimal), size)
		}
		if lib.StartingRoom < 0 || lib.StartingRoom >= lib.Size() {
			t.Errorf("lib.StartingRoom = %d size %d", lib.StartingRoom, lib.Size())
		}
		for room := range lib.Size() {
			if lib.Label(room) != room {
				t.Errorf("lib.Label(%d) = %d, want %d", room, lib.Label(room), room)
			}
			if lib.VisibleLabel(room) != room%4 {
				t.Errorf("lib.VisibleLabel(%d) = %d, want %d", room, lib.VisibleLabel(room), room%4)
			}
			for door := range 6 {
				to := lib.Connections[door][room]
				if to.Room < 0 || to.Room >= size {
					t.Errorf("lib.Connections[%d][%d] = %d, want 0 <= x < %d", door, room, to.Room, size)
				}
			}
		}
		connected := make(map[int]bool)
		for i, connection := range lib.Minimal {
			if connection.From.Room < 0 || connection.From.Room > lib.Size() {
				t.Errorf("bad room in minimal connection[0] %d %v", i, connection)
			}
			if connection.To.Room < 0 || connection.To.Room > lib.Size() {
				t.Errorf("bad room in minimal connection[1] %d %v", i, connection)
			}
			if connection.From.Door < 0 || connection.From.Door >= 6 {
				t.Errorf("bad door in minimal connection[0] %d %v", i, connection)
			}
			if connection.To.Door < 0 || connection.To.Door >= 6 {
				t.Errorf("bad door in minimal connection[1] %d %v", i, connection)
			}
			connected[connection.From.Room] = true
			connected[connection.To.Room] = true
		}
		if len(connected) != lib.Size() {
			t.Errorf("only %d rooms connected out of %d", len(connected), lib.Size())
		}
		for room := range size {
			for door := range 6 {
				node1 := lib.ConnectedRoom(Node{room, door})
				node0 := lib.ConnectedRoom(node1)
				if room != node0.Room || door != node0.Door {
					t.Errorf("size %d bad connection %d,%d -> %v -> %v", size, room, door, node1, node0)
				}
			}
		}
		if size > 1 {
			plan := strings.Repeat("1", size)
			result := lib.Explore(plan)
			if len(result) != size+1 {
				t.Errorf("size %d explore %s failed got %v", size, plan, result)
			}
		}
	}
}

func TestSolver(t *testing.T) {
	for size := 1; size <= 4; size++ {
		lib := MakeLibMap(size)
		server := NewXServer(lib)
		testName := fmt.Sprintf("test%d", size)
		solver := NewSolver(server, testName, size)
		solved, queryCount, err := solver.Solve()
		if err != nil {
			t.Fatalf("%s solver.Solve() returned error: %v", testName, err)
		}
		if queryCount <= 0 {
			t.Errorf("%s bad query count %d", testName, queryCount)
		}
		if !solved {
			t.Fatalf("%s solver.Solve() failed", testName)
		}
	}
}
