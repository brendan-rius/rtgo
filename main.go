package main

import (
	"math"
	"github.com/hajimehoshi/ebiten"
	"fmt"
	"github.com/ungerik/go3d/vec3"
)

type Scene struct {
	Elements []Primitive
}

type Ray struct {
	Origin    *vec3.T
	Direction *vec3.T
}

type Buffer struct {
	Pixels []uint8
}

func (this *Buffer) Set(x int, y int, r uint8, g uint8, b uint8) {
	index := 4 * (y * Width + x)
	this.Pixels[index] = r
	this.Pixels[index + 1] = g
	this.Pixels[index + 2] = b
	this.Pixels[index + 3] = 255
}

func NewBuffer(width, height uint) Buffer {
	return Buffer{make([]uint8, width * height * 4)}
}

const (
	Width = 640
	Height = 480
	HalfHeight = Height / 2
	HalfWidth = Width / 2
	DegToRad = math.Pi / 180
	FieldOfView = 60 * DegToRad
)

var focalDistance = HalfWidth / math.Tan(FieldOfView / 2)
var camPos = vec3.T{0, 0, -500}
var scene = Scene{[]Primitive{Sphere{30}}}
var buffer = NewBuffer(Width, Height)

var upVector = vec3.T{0, 1, 0}
var downVector = vec3.T{0, -1, 0}
var leftVector = vec3.T{-1, 0, 0}
var rightVector = vec3.T{1, 0, 0}
var forwardVector = vec3.T{0, 0, 1}
var backVector = vec3.T{0, 0, -1}

func render() {
	for y := 0; y < Height; y++ {
		go (func(y int) {
			for x := 0; x < Width; x++ {
				pointOnScreen := vec3.T{float32(x - HalfWidth), float32(-y + HalfHeight), float32(focalDistance)}
				direction := vec3.Sub(&pointOnScreen, &camPos)
				ray := Ray{&camPos, &direction}
				for _, p := range scene.Elements {
					if p.Intersect(ray) {
						buffer.Set(x, y, 255, 0, 0)
					} else {
						buffer.Set(x, y, 0, 0, 255)
					}
				}
			}
		})(y)
	}
}

func update(screen *ebiten.Image) error {
	fmt.Println(ebiten.CurrentFPS())
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		camPos.Add(&upVector)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		camPos.Add(&downVector)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		camPos.Add(&leftVector)
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		camPos.Add(&rightVector)
	}
	if ebiten.IsKeyPressed(ebiten.KeyN) {
		camPos.Add(&forwardVector)
	}
	if ebiten.IsKeyPressed(ebiten.KeyM) {
		camPos.Add(&backVector)
	}

	render()

	screen.ReplacePixels(buffer.Pixels[:])
	return nil
}

func main() {
	ebiten.Run(update, Width, Height, 1, "rtgo")
}
