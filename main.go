package main

import rl "github.com/gen2brain/raylib-go/raylib"

var red = []uint32{255,0,0,255}

type sphere struct { 
	radius float32,
	color []uint8,
	center []float32
}

/*
O = (0, 0, 0)
for x = -Cw/2 to Cw/2 {
    for y = -Ch/2 to Ch/2 {
        D = CanvasToViewport(x, y)
        color = TraceRay(O, D, 1, inf)
        canvas.PutPixel(x, y, color)
    }
}
*/
var sphere1 = sphere { 
	radius = 1 
	center = 
}
func main() {
	rl.InitWindow(1920, 1080, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	origin := {0, 0, 0}
	
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		for x := -960; x < 960; x++{
			for y := -540 ; y < 540; y++{
				D := CanvasToViewPort(x, y)
				color := TraceRay(origin, D, 1, 9999999)
			}
		}
		//draw_triangle(0,0,0,50,50,50)

		rl.EndDrawing()
	}
}
func CanvasToViewPort(x int, y int){
	return {x,y,2}
}

func TraceRay(O []int, D []int, t_min int, t_max int){
	closest_t = 9999999
	closest_sphere = nil 

}












//x2-x1,y2-y1
/*
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
*/