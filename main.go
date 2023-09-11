package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const screen_width = 1200
const screen_height = 600
const size = 10

var fps = int32(24)

const number_of_cells = screen_width / size

var cells []*Cell

func main() {
	cells = InitCells()

	rl.SetTraceLogLevel(rl.TraceLogLevel(6))
	rl.InitWindow(screen_width, screen_height, "matrix-go")
	defer rl.CloseWindow()
	rl.SetTargetFPS(fps)

	font := rl.LoadFont("./assets/font.ttf")

	color4 := rl.NewColor(0, 51, 51, 255)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		for _, cell := range cells {
			rl.DrawTextEx(font, string(cell.char), rl.Vector2{X: float32(cell.xPos), Y: float32(cell.yPos)}, float32(size), 0, rl.LightGray)
			rl.DrawTextEx(font, string(cell.prevChar), rl.Vector2{X: float32(cell.xPos), Y: float32(cell.yPos - int(cell.speed))}, float32(size), 0, color4)

			cell.step()
			cell.check()
		}

		rect := rl.Rectangle{
			X:      0.0,
			Y:      0.0,
			Width:  float32(screen_width),
			Height: float32(screen_height),
		}

		color := rl.NewColor(0, 0, 0, 40)
		rl.DrawRectangleRounded(rect, 0, 0, color)

		fps = int32(gui.Slider(rl.NewRectangle(1040, 570, 150, 20), "speed", "", float32(fps), 18, 64))

		color4 = gui.ColorPicker(rl.NewRectangle(1015, 400, 150, 150), "color", color4)

		rl.SetTargetFPS(fps)

		rl.EndDrawing()
	}
}
