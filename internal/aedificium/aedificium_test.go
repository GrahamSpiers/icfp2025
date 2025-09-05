package aedificium

import "testing"

func TestMakeComplexLibrary(t *testing.T) {
	for size := 1; size <= 10; size++ {
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
				toRoom := lib.Connections[door][room]
				if toRoom < 0 || toRoom >= size {
					t.Errorf("lib.Connections[%d][%d] = %d, want 0 <= x < %d", door, room, toRoom, size)
				}
			}
		}
		connected := make(map[int]bool)
		for i, connection := range lib.Minimal {
			if connection[0][0] < 0 || connection[0][0] > lib.Size() {
				t.Errorf("bad room in minimal connection[0] %d %v", i, connection[0])
			}
			if connection[1][0] < 0 || connection[1][0] > lib.Size() {
				t.Errorf("bad room in minimal connection[1] %d %v", i, connection[1])
			}
			if connection[0][1] < 0 || connection[0][1] >= 6 {
				t.Errorf("bad door in minimal connection[0] %d %v", i, connection[0])
			}
			if connection[1][1] < 0 || connection[1][1] >= 6 {
				t.Errorf("bad door in minimal connection[1] %d %v", i, connection[0])
			}
			connected[connection[0][0]] = true
			connected[connection[1][0]] = true
		}
		if len(connected) != lib.Size() {
			t.Errorf("only %d rooms connected out of %d", len(connected), lib.Size())
		}
	}
}
