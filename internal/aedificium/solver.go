package aedificium

import "fmt"

const Id = "grahamsspiers@gmail.com ffrbCqDWn7ARkZ9pR26Frg"

type Server interface {
	Select(problemName string) (string, error)
	Explore(plans []string) ([][]int, int, error)
	Guess() (bool, error)
}

type XServer struct {
	Lib        LibMap
	Server     Server
	QueryCount int
}

func NewXServer(lib LibMap) *XServer {
	xs := &XServer{Lib: lib, QueryCount: 0}
	return xs
}

// Choose a problem to work on.
func (xs *XServer) Select(problemName string) (string, error) {
	fmt.Printf("select %s\n", problemName)
	return problemName, nil
}

// Follow the given plans.
func (xs *XServer) Explore(plans []string) ([][]int, int, error) {
	fmt.Printf("explore %d %v\n", xs.QueryCount, plans)
	results := make([][]int, len(plans))
	for i, plan := range plans {
		results[i] = xs.Lib.Explore(plan)
		//fmt.Printf("  %s %v\n", plan, results[i])
	}
	xs.QueryCount += 1 + len(plans)
	return results, xs.QueryCount, nil
}

func (xs *XServer) Guess() (bool, error) {
	fmt.Printf("guess\n")
	return true, nil
}

type RoomInfo struct {
	Label  int // 0-3
	ToHere string
	Doors  [6]int // Visible labels for each door.
}

func (info RoomInfo) Equals(other RoomInfo) bool {
	return info.Label == other.Label && info.Doors == other.Doors
}

type Solver struct {
	ProblemName string
	Size        int
	Server      Server
	Info        []RoomInfo
	Plans       []string
}

func NewSolver(server Server, problemName string, size int) *Solver {
	return &Solver{
		ProblemName: problemName,
		Size:        size,
		Server:      server,
		Info:        []RoomInfo{},
		// Add a plan for "solving the entry."
		Plans: []string{"0", "1", "2", "3", "4", "5"},
	}
}

func (solver *Solver) Solve() (bool, int, error) {
	// Select
	solver.Server.Select(solver.ProblemName)
	var queryCount = 0
	for {
		// plan
		if len(solver.Plans) == 0 {
			return false, queryCount, fmt.Errorf("%s %d no plans", solver.ProblemName, solver.Size)
		}
		//fmt.Printf("plans %v\n", solver.Plans)
		// explore
		var results [][]int
		var err error
		results, queryCount, err = solver.Server.Explore(solver.Plans)
		if err != nil {
			return false, queryCount, fmt.Errorf("%s %d explore got %v", solver.ProblemName, solver.Size, err)
		}
		// add to map
		solver.LearnAndPlan(results)
		//fmt.Printf("%d rooms\n", len(solver.Info))
		// map complete?
		if len(solver.Info) == solver.Size {
			//	guess
			correct, err := solver.Server.Guess()
			if err != nil {
				return false, queryCount, fmt.Errorf("%s %d guess got %v", solver.ProblemName, solver.Size, err)
			}
			return correct, queryCount, nil
		}
		if len(solver.Info) > solver.Size {
			return false, queryCount, fmt.Errorf("%s %d bad number of rooms %d", solver.ProblemName, solver.Size, len(solver.Info))
		}
	}
}

func (solver *Solver) LearnAndPlan(results [][]int) {
	nextPlans := []string{}
	//fmt.Printf("%d plans\n", len(solver.Plans))
	for i := range len(solver.Plans) / 6 {
		//first := i * 6
		plans := solver.Plans[i*6 : i*6+6]
		result := results[i*6 : i*6+6]
		roomInfo := RoomInfo{
			Label:  result[0][len(result[0])-2],
			ToHere: plans[0][0 : len(plans[0])-1],
		}
		for door := range 6 {
			//j := first + door
			//fmt.Printf("  plan %s result %v\n", solver.Plans[j], result[door])
			roomInfo.Doors[door] = result[door][len(result[door])-1]
		}
		//fmt.Printf("  room %+v\n", roomInfo)
		var shouldAdd = true
		for _, oldInfo := range solver.Info {
			if oldInfo.Equals(roomInfo) {
				shouldAdd = false
				break
			}
		}
		if shouldAdd {
			//fmt.Printf("  added\n")
			solver.Info = append(solver.Info, roomInfo)
			for _, plan := range plans {
				for _, door := range []string{"0", "1", "2", "3", "4", "5"} {
					nextPlans = append(nextPlans, plan+door)
				}
			}
		}
	}
	solver.Plans = nextPlans
}
