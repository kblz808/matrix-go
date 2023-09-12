package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const screen_width = 1200
const screen_height = 600
const size = 10

var fps = int32(24)

const number_of_cells = (screen_width - 200) / size

var cells []*Cell

func main() {
	cells = InitCells()

	rl.SetTraceLogLevel(rl.TraceLogLevel(6))
	rl.InitWindow(screen_width, screen_height, "matrix-go")
	defer rl.CloseWindow()
	rl.SetTargetFPS(fps)

	font := rl.LoadFont("./assets/font.ttf")

	color4 := rl.NewColor(0, 51, 51, 255)

	rect := rl.Rectangle{
		X:      0.0,
		Y:      0.0,
		Width:  float32(screen_width - 200),
		Height: float32(screen_height),
	}
	opacity := uint8(45)
	color := rl.NewColor(10, 10, 10, opacity)

	rect2 := rl.Rectangle{
		X:      1000.0,
		Y:      0.0,
		Width:  200.0,
		Height: float32(screen_height),
	}
	color2 := rl.NewColor(25, 25, 25, 255)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		for _, cell := range cells {
			rl.DrawTextEx(font, string(cell.char), rl.Vector2{X: float32(cell.xPos), Y: float32(cell.yPos)}, float32(size), 0, rl.LightGray)
			rl.DrawTextEx(font, string(cell.prevChar), rl.Vector2{X: float32(cell.xPos), Y: float32(cell.yPos - int(cell.speed))}, float32(size), 0, color4)

			cell.step()
			cell.check()
		}

		rl.DrawRectangleRounded(rect, 0, 0, color)
		rl.DrawRectangleRounded(rect2, 0, 0, color2)

		fps = int32(gui.Slider(rl.NewRectangle(1015, 570, 150, 20), "", "fps", float32(fps), 18, 64))
		color4 = gui.ColorPicker(rl.NewRectangle(1015, 400, 150, 150), "color", color4)
		opacity = uint8(gui.Slider(rl.NewRectangle(1015, 360, 150, 20), "", "alpha", float32(opacity), 10, 100))
		color = rl.NewColor(0, 0, 0, opacity)

		rl.SetTargetFPS(fps)

		rl.EndDrawing()
	}
}
