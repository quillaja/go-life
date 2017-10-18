package main

import (
	"image"
	_ "image/png"
	"math"
	"os"

	"golang.org/x/image/colornames"

	g "local/life/game"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// demos the stuff
func main() {

	// var pat = flag.String("p", "blinker", "name of an initial pattern.")
	// var iter = flag.Int("i", 100, "number of iterations to run")
	// var pause = flag.Duration("w", 500, "duration to wait between each iteration")
	// flag.Parse()

	// patterns := map[string]g.Board{
	// 	// Blinker
	// 	"blinker": g.Board{
	// 		g.Point{0, 0}: true,
	// 		g.Point{0, 1}: true,
	// 		g.Point{0, 2}: true,
	// 	},

	// 	// R-Pentomino
	// 	"rpentomino": g.Board{
	// 		g.Point{1, 0}: true,
	// 		g.Point{2, 0}: true,
	// 		g.Point{0, 1}: true,
	// 		g.Point{1, 1}: true,
	// 		g.Point{1, 2}: true,
	// 	},

	// 	// Acorn
	// 	"acorn": g.Board{
	// 		g.Point{1, 0}: true,
	// 		g.Point{3, 1}: true,
	// 		g.Point{0, 2}: true,
	// 		g.Point{1, 2}: true,
	// 		g.Point{4, 2}: true,
	// 		g.Point{5, 2}: true,
	// 		g.Point{6, 2}: true,
	// 	},
	// }

	pixelgl.Run(loop)
}

func loadPicture(filename string) (pixel.Picture, error) {
	fr, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	img, _, err := image.Decode(fr)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

func loop() {

	acorn := g.Board{
		g.Point{1, 0}: true,
		g.Point{3, 1}: true,
		g.Point{0, 2}: true,
		g.Point{1, 2}: true,
		g.Point{4, 2}: true,
		g.Point{5, 2}: true,
		g.Point{6, 2}: true,
	}

	cfg := pixelgl.WindowConfig{
		Title:  "Life...finds a way.",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	redPic, _ := loadPicture("red.png")
	red := pixel.NewSprite(redPic, redPic.Bounds())
	batch := pixel.NewBatch(&pixel.TrianglesData{}, redPic)

	camZoom := 1.0
	camZoomSpeed := 1.2

	for !win.Closed() {

		win.SetMatrix(pixel.IM.Scaled(win.Bounds().Center(), camZoom))
		camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

		// draw
		batch.Clear()
		for p := range acorn {
			red.Draw(batch, pixel.IM.
				Moved(win.Bounds().Center()).
				Moved(pToV(p)))
		}

		win.Clear(colornames.White)
		batch.Draw(win)
		win.Update()

		// update state
		acorn = g.Advance(acorn)
	}
}

func pToV(p g.Point) pixel.Vec {
	return pixel.V(float64(p.X), float64(p.Y))
}
