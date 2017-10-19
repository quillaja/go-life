package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/png"
	"math"
	"os"
	"time"

	"golang.org/x/image/colornames"

	g "local/life/game"

	p "github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Camera struct {
	Position p.Vec
	Speed    float64
	Zoom     float64
	ZSpeed   float64
}

var (
	initPattern string
	// pause       time.Duration
)

// demos the stuff
func main() {

	flag.StringVar(&initPattern, "p", "", "name of an initial pattern.")
	// flag.DurationVar(&pause, "w", 500, "duration to wait between each iteration")
	flag.Parse()

	if _, ok := g.Patterns[initPattern]; !ok {
		fmt.Printf("-p flag error: \"%s\" is an invalid pattern name.", initPattern)
		os.Exit(1)
	}

	pixelgl.Run(loop)
}

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

func loop() {

	cfg := pixelgl.WindowConfig{
		Title:  "Life...finds a way.",
		Bounds: p.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// sprite loading
	redPic, err := loadPicture("red.png")
	if err != nil {
		panic(err)
	}
	red := p.NewSprite(redPic, redPic.Bounds())
	batch := p.NewBatch(&p.TrianglesData{}, redPic)

	// game state
	board := g.Patterns[initPattern]
	paused := false

	// various state for drawing
	cam := Camera{Position: p.ZV, Speed: 500.0, Zoom: 1.0, ZSpeed: 1.1}
	frames := 0
	second := time.Tick(time.Second)
	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

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
		cam.Zoom *= math.Pow(cam.ZSpeed, win.MouseScroll().Y)

		win.SetMatrix(p.IM.
			Scaled(cam.Position, cam.Zoom).
			Moved(win.Bounds().Center().Sub(cam.Position)))

		// render game state
		batch.Clear()
		for point := range board {
			red.Draw(batch, p.IM.Moved(pToV(point)))
		}

		// draw
		if paused {
			win.Clear(colornames.Gray)
		} else {
		win.Clear(colornames.White)
		}
		batch.Draw(win)
		win.Update()

		// basic speed metric in titlebar
		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}

		// update state
		if !paused {
			board = g.Advance(board)
		}
	}
}

func pToV(point g.Point) p.Vec {
	return p.V(float64(point.X), float64(point.Y))
}
