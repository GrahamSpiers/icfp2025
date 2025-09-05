package aedificium

import "testing"

func TestMakeSimpleLibrary(t *testing.T) {
	for size := 1; size <= 10; size++ {
		lib := MakeSimpleLibrary(size)
		if len(lib) != size {
			t.Errorf("len(Rooms) = %d, want %d", len(lib), size)
		}
		for i, room := range lib {
			if room.Label != i {
				t.Errorf("Rooms[%d].Label = %d, want %d", i, room.Label, i)
			}
			if room.Bits() != i%4 {
				t.Errorf("Rooms[%d].Bits() = %d, want %d", i, room.Bits(), i%4)
			}
			for j, to := range room.Doors {
				if j == 0 {
					want := (i + 1) % size
					if to != want {
						t.Errorf("Rooms[%d][%d] = %d, want %d", i, j, to, want)
					}
				} else {
					if to != 0 {
						t.Errorf("Rooms[%d][%d] = %d, want %d", i, j, to, 0)
					}
				}
			}
		}
	}
}

func TestMakeComplexLibrary(t *testing.T) {
	for size := 1; size <= 10; size++ {
		lib := MakeSimpleLibrary(size)
		if len(lib) != size {
			t.Errorf("len(Rooms) = %d, want %d", len(lib), size)
		}
		for i, room := range lib {
			if room.Label != i {
				t.Errorf("Rooms[%d].Label = %d, want %d", i, room.Label, i)
			}
			for j, to := range room.Doors {
				if to < 0 || to >= size {
					t.Errorf("Rooms[%d][%d] = %d, want 0 <= x < %d", i, j, to, size)
				}
			}
		}
	}
}
