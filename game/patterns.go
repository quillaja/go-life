package game

// Patterns contains a number of pre-set patterns in `Board`s.
var Patterns = map[string]Board{
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
}
