package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/png"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/image/colornames"

	g "github.com/quillaja/go-life/game"

	p "github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type camera struct {
	Position p.Vec
	Speed    float64
	Zoom     float64
	ZSpeed   float64
}

var (
	initPattern string
	iterWait    time.Duration
	console     bool
	winSizeStr  string

	winBounds p.Rect
	resPath   string
)

// demos the stuff
func main() {

	// command line options processing
	flag.Usage = func() {
		msg :=
			`Plays Conway's Game of Life. It can produce a nice graphical display 
(requires OpenGL 3.3+), or console output (ANSI compatible console recommended).
In graphical mode, use the arrow keys to scroll the view area, mouse wheel to 
zoom, and space to pause/unpause. When paused, use the left mouse button to 
turn a cell on or off, and the right mouse button to perform one iteration 
of the game.`
		fmt.Fprintln(os.Stderr, msg)
		flag.PrintDefaults()
	}

	flag.StringVar(&initPattern, "p", "blank",
		"Name of an initial pattern. Choices: "+
			strings.Join(g.PatternNames(), ", "))
	flag.DurationVar(&iterWait, "w", 100*time.Millisecond,
		"Duration to wait between each iteration.")
	flag.StringVar(&winSizeStr, "s", "800x600",
		"Window height and width--must have the 'x'.")
	flag.BoolVar(&console, "c", false,
		"Set to display output to console.")
	flag.Parse()

	if _, ok := g.Patterns[initPattern]; !ok {
		fmt.Fprintf(os.Stderr, "-p flag error: \"%s\" is an invalid pattern name.\n", initPattern)
		fmt.Fprintf(os.Stderr, "  Choices are: %v\n", g.PatternNames())
		os.Exit(1)
	}

	// process window dimensions option
	var err error
	pErr := func() {
		fmt.Fprintln(os.Stderr, "-s flag error: Must be in format \"<height>x<width>\".")
		os.Exit(1)
	}
	hwSplit := strings.Split(winSizeStr, "x")
	if len(hwSplit) != 2 {
		pErr()
	}
	winBounds.Max.X, err = strconv.ParseFloat(hwSplit[0], 64)
	if err != nil {
		pErr()
	}
	winBounds.Max.Y, err = strconv.ParseFloat(hwSplit[1], 64)
	if err != nil {
		pErr()
	}

	// construct path to image resources
	// goPath := strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator))
	// if len(goPath) == 0 {
	// 	goPath = append(goPath, build.Default.GOPATH)
	// } else if goPath[0] == "" {
	// 	goPath[0] = build.Default.GOPATH
	// }
	// resPath = filepath.Join(goPath[0], "src", "github.com", "quillaja", "go-life")

	// begin
	if console {
		g.Animate(g.Patterns[initPattern], 100, iterWait)
	} else {
		pixelgl.Run(loop)
	}
}

// the actual game loop. A variety of initialization followed by the draw loop
func loop() {

	cfg := pixelgl.WindowConfig{
		Title:  "Life...finds a way.",
		Bounds: winBounds,
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// sprite creation
	redPic := p.MakePictureData(p.R(0, 0, 1, 1))
	redPic.Pix[0] = colornames.Red
	red := p.NewSprite(redPic, redPic.Bounds())

	batch := p.NewBatch(&p.TrianglesData{}, redPic)

	// game state
	board := g.Patterns[initPattern]
	paused := false
	iterations := 0

	// update the game state (board) every iterWait duration
	// independently of the graphical draw loop
	go func() {
		wait := time.NewTicker(iterWait)
		for !win.Closed() {
			select {
			case <-wait.C:
				if !paused {
					board = g.Advance(board)
					iterations++
				}
			default:
			}
		}
		wait.Stop()
	}()

	// various state for drawing
	cam := camera{Position: p.ZV, Speed: 250.0, Zoom: 1.0, ZSpeed: 1.1}
	frames := 0
	second := time.Tick(time.Second)
	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		camMatrix := p.IM.
			Scaled(cam.Position, cam.Zoom).
			Moved(win.Bounds().Center().Sub(cam.Position))
		win.SetMatrix(camMatrix)

		// update user controlled things
		if win.Pressed(pixelgl.KeyLeft) {
			cam.Position.X -= cam.Speed * dt
		}
		if win.Pressed(pixelgl.KeyRight) {
			cam.Position.X += cam.Speed * dt
		}
		if win.Pressed(pixelgl.KeyDown) {
			cam.Position.Y -= cam.Speed * dt
		}
		if win.Pressed(pixelgl.KeyUp) {
			cam.Position.Y += cam.Speed * dt
		}
		if win.JustPressed(pixelgl.KeySpace) {
			paused = !paused
		}
		// win.SetTitle(fmt.Sprintf("Mouse (%.2f, %.2f)", mouse.X, mouse.Y))
		if win.JustPressed(pixelgl.MouseButtonLeft) && paused {
			// toggle a point's existence.
			// use Round to change mouse floats to ints. Simple
			// truncation will often place dot in wrong spot since
			// Pixel uses the sprite's center as it's position.
			mouse := camMatrix.Unproject(win.MousePosition())
			point := g.Point{round(mouse.X), round(mouse.Y)}
			if board[point] {
				delete(board, point)
			} else {
				board[point] = true
			}
		}
		if win.JustPressed(pixelgl.MouseButtonRight) && paused {
			// allow user to increment the board state 1 iteration
			board = g.Advance(board)
			iterations++
		}
		cam.Zoom *= math.Pow(cam.ZSpeed, win.MouseScroll().Y)

		// render game state
		batch.Clear()
		for point := range board {
			red.Draw(batch, p.IM.Moved(pToV(point)))
		}

		// draw
		if paused {
			win.Clear(colornames.Lightgray)
		} else {
			win.Clear(colornames.White)
		}
		batch.Draw(win)
		win.Update()

		// various metrics in titlebar
		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf(
				"%s | FPS: %d | Paused: %v | %d cells %d iterations",
				cfg.Title, frames, paused, len(board), iterations))
			frames = 0
		default:
		}
	}
}

// convenience function for loading an image file
func loadPicture(filename string) (p.Picture, error) {
	fr, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	img, _, err := image.Decode(fr)
	if err != nil {
		return nil, err
	}

	return p.PictureDataFromImage(img), nil
}

// convenience function to convert game.Point to pixel.Vec
func pToV(point g.Point) p.Vec {
	return p.V(float64(point.X), float64(point.Y))
}

// utility function to round floats to ints, since golang is so
// omniscient to realize that we don't need this crap in the std libary
func round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}
