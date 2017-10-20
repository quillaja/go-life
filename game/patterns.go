package game

// Patterns contains a number of pre-set patterns in `Board`s.
var Patterns = map[string]Board{

	// Blank
	"blank": Board{},

	// Blinker
	"blinker": Board{
		Point{0, 0}: true,
		Point{0, 1}: true,
		Point{0, 2}: true,
	},

	// Toad
	"toad": Board{
		Point{1, 0}: true,
		Point{2, 0}: true,
		Point{3, 0}: true,
		Point{0, 1}: true,
		Point{1, 1}: true,
		Point{2, 1}: true,
	},

	// Beacon
	"beacon": Board{
		Point{0, 0}: true,
		Point{0, 1}: true,
		Point{1, 0}: true,
		Point{3, 3}: true,
		Point{2, 3}: true,
		Point{3, 2}: true,
	},

	// Diehard
	"diehard": Board{
		Point{0, 1}: true,
		Point{1, 1}: true,
		Point{1, 2}: true,
		Point{6, 0}: true,
		Point{5, 2}: true,
		Point{6, 2}: true,
		Point{7, 2}: true,
	},

	// R-Pentomino
	"rpentomino": Board{
		Point{1, 0}: true,
		Point{2, 0}: true,
		Point{0, 1}: true,
		Point{1, 1}: true,
		Point{1, 2}: true,
	},

	// Acorn
	"acorn": Board{
		Point{1, 0}: true,
		Point{3, 1}: true,
		Point{0, 2}: true,
		Point{1, 2}: true,
		Point{4, 2}: true,
		Point{5, 2}: true,
		Point{6, 2}: true,
	},

	// 5 x 5
	"five_by_five": Board{
		Point{0, 0}: true,
		Point{1, 0}: true,
		Point{2, 0}: true,
		Point{4, 0}: true,
		Point{0, 1}: true,
		Point{3, 2}: true,
		Point{4, 2}: true,
		Point{1, 3}: true,
		Point{2, 3}: true,
		Point{4, 3}: true,
		Point{0, 4}: true,
		Point{2, 4}: true,
		Point{4, 4}: true,
	},

	// 10-cell
	"ten_cell": Board{
		Point{6, 0}: true,
		Point{4, 1}: true,
		Point{6, 1}: true,
		Point{7, 1}: true,
		Point{4, 2}: true,
		Point{6, 2}: true,
		Point{4, 3}: true,
		Point{2, 4}: true,
		Point{0, 5}: true,
		Point{2, 5}: true,
	},

	// Line
	"line": Board{
		Point{0, 0}:  true,
		Point{1, 0}:  true,
		Point{2, 0}:  true,
		Point{3, 0}:  true,
		Point{4, 0}:  true,
		Point{5, 0}:  true,
		Point{6, 0}:  true,
		Point{7, 0}:  true,
		Point{9, 0}:  true,
		Point{10, 0}: true,
		Point{11, 0}: true,
		Point{12, 0}: true,
		Point{13, 0}: true,
		Point{17, 0}: true,
		Point{18, 0}: true,
		Point{19, 0}: true,
		Point{26, 0}: true,
		Point{27, 0}: true,
		Point{28, 0}: true,
		Point{29, 0}: true,
		Point{30, 0}: true,
		Point{31, 0}: true,
		Point{32, 0}: true,
		Point{34, 0}: true,
		Point{35, 0}: true,
		Point{36, 0}: true,
		Point{37, 0}: true,
		Point{38, 0}: true,
	},
}

// PatternNames returns a slice containing the names of all the predefined patterns.
func PatternNames() []string {
	names := []string{}
	for k := range Patterns {
		names = append(names, k)
	}
	return names
}
