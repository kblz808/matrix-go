package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const spacing = 220
const screen_width = 1000 + spacing
const screen_height = 600
const size = 10

var fps = int32(24)

const number_of_cells = (screen_width - spacing) / size

var cells []*Cell

var (
	gridSpacing     = 16
	panelRec        = rl.Rectangle{X: 1000, Y: 0, Width: spacing, Height: screen_height}
	panelContentRec = rl.Rectangle{X: 1000, Y: 0, Width: spacing, Height: screen_height}
	panelScroll     = rl.Vector2{X: 100, Y: 100}
	color4          = rl.NewColor(0, 51, 51, 255)
)

func main() {
	cells = InitCells()

	rl.SetTraceLogLevel(rl.TraceLogLevel(6))
	rl.InitWindow(screen_width, screen_height, "matrix-go")
	defer rl.CloseWindow()
	rl.SetTargetFPS(fps)

	font := rl.LoadFont("./assets/font.ttf")

	shader := rl.LoadShader("", "./assets/shader.fs")

	rect := rl.Rectangle{
		X:      0.0,
		Y:      0.0,
		Width:  float32(screen_width - spacing),
		Height: float32(screen_height),
	}
	opacity := uint8(45)
	color := rl.NewColor(0, 0, 0, opacity)

	target := rl.LoadRenderTexture(screen_width-spacing, screen_height)

	for !rl.WindowShouldClose() {
		rl.SetTargetFPS(fps)
		color.A = opacity

		rl.BeginDrawing()

		view := gui.ScrollPanel(panelRec, "text", panelContentRec, &panelScroll)
		rl.BeginScissorMode(int32(view.X), int32(view.Y), int32(view.Width), int32(view.Height))
		gui.Grid(rl.Rectangle{
			X:      float32(panelRec.X + panelScroll.X),
			Y:      float32(panelRec.Y + panelScroll.Y),
			Width:  float32(panelContentRec.Width),
			Height: float32(panelContentRec.Height),
		}, "", float32(gridSpacing), 1)

		color4 = gui.ColorPicker(rl.NewRectangle(panelRec.X+float32(gridSpacing), panelRec.Y+float32(gridSpacing*3)+panelScroll.Y, float32(gridSpacing*9), float32(gridSpacing*9)), "color", color4)
		fps = int32(gui.Slider(rl.NewRectangle(panelRec.X+float32(gridSpacing), panelRec.Y+float32(gridSpacing*14)+panelScroll.Y, 150, 20), "", "fps", float32(fps), 18, 64))
		opacity = uint8(gui.Slider(rl.NewRectangle(panelRec.X+float32(gridSpacing), panelRec.Y+float32(gridSpacing*16)+panelScroll.Y, 150, 20), "", "alpha", float32(opacity), 10, 100))
		rl.EndScissorMode()

		rl.BeginTextureMode(target)
		for _, cell := range cells {
			rl.DrawTextEx(font, string(cell.char), rl.Vector2{X: float32(cell.xPos), Y: float32(cell.yPos)}, float32(size), 0, rl.LightGray)
			rl.DrawTextEx(font, string(cell.prevChar), rl.Vector2{X: float32(cell.xPos), Y: float32(cell.yPos - int(cell.speed))}, float32(size), 0, color4)

			cell.step()
			cell.check()
		}

		rl.DrawRectangleRounded(rect, 0, 0, color)
		rl.EndTextureMode()

		rl.BeginShaderMode(shader)
		rl.DrawTextureRec(target.Texture, rl.NewRectangle(0, 0, float32(target.Texture.Width), float32(-target.Texture.Height)), rl.NewVector2(0, 0), rl.White)
		rl.EndShaderMode()

		rl.EndDrawing()
	}
}
