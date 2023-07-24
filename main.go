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
		draw_triangle(0,0,0,50,50,50)

		rl.EndDrawing()
	}
}

//x2-x1,y2-y1
func calc_vector(x1 float32, y1 float32,x2 float32, y2 float32)(float32,float32){
	return  x2-x1,y2-y1
}
func draw_triangle(x1 float32, y1 float32,x2 float32, y2 float32,x3 float32, y3 float32){
	//calcualte ray from points, then pass the ray to be drawn
	//i1,j1 := calc_vector(x1,y1,x2,y2)
	i2,j2 := calc_vector(x2,y2,x3,y3)
	//i3,j3 := calc_vector(x3,y3,x1,y1)
	//draw_line(i1,j1)
	draw_line(i2,j2,x2,y2)
	//draw_line(i3,j3)

}

func draw_line(x float32, y float32, a float32, b float32){
	for i := a; i < x; i++{
		j := i
		rl.DrawPixel(int32(i),int32(j), rl.NewColor(255, 0, 0, 255))
	}

}