package main

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/png"
	"log"
	"math/rand"
)

//go:embed assets/*
var EmbeddedAssets embed.FS

type enemy Sprite

const (
	GameWidth   = 700
	GameHeight  = 700
	PlayerSpeed = 2
)

type Sprite struct {
	pict *ebiten.Image
	xloc int
	yloc int
	dX   int
	dY   int
}

type Game struct {
	player     Sprite
	score      int
	drawOps    ebiten.DrawImageOptions
	enemySlice []enemy
}

func (g *Game) Update() error {
	if doesCollide {
		
	}

	processPlayerInput(g)
	return nil

}

func (g Game) Draw(screen *ebiten.Image) {
	g.drawOps.GeoM.Reset()
	g.drawOps.GeoM.Translate(float64(g.player.xloc), float64(g.player.yloc))

	screen.DrawImage(g.player.pict, &g.drawOps)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GameWidth, GameHeight
}

func main() {
	ebiten.SetWindowSize(GameWidth, GameHeight)
	ebiten.SetWindowTitle("Minimal Game")
	simpleGame := Game{score: 0}
	simpleGame.player = Sprite{
		pict: loadPNGImageFromEmbedded("tanjiro-kamado-demon-slayer-color-by-number_icon_android.png"),
		xloc: 200,
		yloc: 300,
		dX:   0,
		dY:   0,
	}

	//making a slice of 10 enemies
	//looping through enemy sprites
	enemySlice := make([]enemy, 10)
	i := 0
	max := 700
	min := 0
	random := rand.Intn(max-min) + min

	for range enemySlice {
		current := enemy{pict: loadPNGImageFromEmbedded("PngItem_5308340.png"), xloc: random, yloc: random}
		enemySlice[i] = current
		i++
	}

	if err := ebiten.RunGame(&simpleGame); err != nil {
		log.Fatal("Oh no! something terrible happened and the game crashed", err)
	}

}

func loadPNGImageFromEmbedded(name string) *ebiten.Image {
	pictNames, err := EmbeddedAssets.ReadDir("assets")
	if err != nil {
		log.Fatal("failed to read embedded dir ", pictNames, " ", err)
	}
	embeddedFile, err := EmbeddedAssets.Open("assets/" + name)
	if err != nil {
		log.Fatal("failed to load embedded image ", embeddedFile, err)
	}
	rawImage, err := png.Decode(embeddedFile)
	if err != nil {
		log.Fatal("failed to load embedded image ", name, err)
	}
	gameImage := ebiten.NewImageFromImage(rawImage)
	return gameImage

}

func processPlayerInput(theGame *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		theGame.player.dY = -PlayerSpeed
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		theGame.player.dY = PlayerSpeed
	} else if inpututil.IsKeyJustReleased(ebiten.KeyUp) || inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		theGame.player.dY = 0
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		theGame.player.dX = -PlayerSpeed
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		theGame.player.dX = PlayerSpeed
	} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) || inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) {
		theGame.player.dX = 0

	}

	theGame.player.yloc += theGame.player.dY
	if theGame.player.yloc <= 0 {
		theGame.player.dY = 0
		theGame.player.yloc = 0
	} else if theGame.player.yloc+theGame.player.pict.Bounds().Size().Y > GameHeight {
		theGame.player.dY = 0
		theGame.player.yloc = GameHeight - theGame.player.pict.Bounds().Size().Y
	}
}

func playerScore(g Game) {
	//scoreboard, err := Game{}

}

func init() {

}
