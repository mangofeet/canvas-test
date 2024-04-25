package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

func main() {

	files := []string{
		"HB",
		"JINTEKI",
		"WEYLAND",
	}

	var data []*canvas.Canvas

	width, height := 0.0, 0.0
	x := 0.0

	for _, name := range files {

		file, err := os.Open(fmt.Sprintf("%s.svg", name))
		if err != nil {
			panic(err)
		}

		c, err := canvas.ParseSVG(file)
		if err != nil {
			panic(err)
		}

		// c.Transform(canvas.Identity.Scale(0.1, 0.1))
		c.Transform(canvas.Identity.Translate(width, 0))

		width += c.W
		height = math.Max(height, c.H)

		data = append(data, c)

	}

	cnv := canvas.New(width, height)

	ctx := canvas.NewContext(cnv)

	for _, c := range data {
		c.RenderTo(ctx)
		x += c.W
	}

	filename := "out.png"

	log.Printf("rendering output to %s", filename)
	if err := renderers.Write(filename, cnv, canvas.DPMM(1)); err != nil {
		panic(err)
	}
	log.Println("done")

}
