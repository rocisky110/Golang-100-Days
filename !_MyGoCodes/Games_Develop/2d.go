package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	playerImage *ebiten.Image
	playerX     = 0.0
	playerY     = 0.0
)

func init() {
	var err error
	playerImage, _, err = ebitenutil.NewImageFromFile("path/to/player.png", ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
}

func update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		playerX += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		playerX -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		playerY += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		playerY -= 5
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(playerX, playerY)
	screen.DrawImage(playerImage, op)
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "My Game"); err != nil {
		panic(err)
	}
}
