package game

import (
	"image"
	"image/color"
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	layers [][]int
}

func NewGame() *Game {
	g := &Game{
		assets.L.Layers,
	}
	return g
}

const (
	screenWidth  = 640
	screenHeight = 480
)

const (
	tileSize = 16
)

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 40, 40, 255})
	w := assets.Tilemap.Bounds().Dx()
	tileXCount := w / tileSize

	const xCount = screenWidth / tileSize
	for _, l := range g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))

			sx := (t % tileXCount) * tileSize
			sy := (t / tileXCount) * tileSize
			screen.DrawImage(assets.Tilemap.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
