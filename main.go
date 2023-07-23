package main

import rl "github.com/gen2brain/raylib-go/raylib"

var red = []uint32{255,0,0,255}

func main() {
	rl.InitWindow(800, 800, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		draw_line()
		draw_triangle()

		rl.EndDrawing()
	}
}


func draw_triangle(){
	for i := 0; i < 350; i++{
		for j:=i; j <350; j++{
			rl.DrawPixel(int32(i),int32(j), rl.NewColor(0, 0, 0, 255))
		}
	}
}

func draw_line(){
	for i := 350; i < 450; i++{
		j := i
		rl.DrawPixel(int32(i),int32(j), rl.NewColor(255, 0, 0, 255))
	}
}