package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
	"fmt"
)
var red = []uint32{255,0,0,255}

type sphere struct { 
	radius float32
	color [3]uint8
	center [3]float32
}


func main() {

	var sphere1 = sphere{1,{255,0,0},{0,-1,3}}
	var sphere2 = sphere{1,{0,255,0},{-2,0,4}}
	var sphere3 = sphere{1,{0,0,255},{2,0,4}}

	var scene = [3]sphere{sphere1,sphere2,sphere3}
	
	rl.InitWindow(1920, 1080, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	origin := [3]int{0, 0, 0}
	
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
	return [3]float32{x,y,2}
}
func  IntersectRaySphere(O []int, D []int, shape sphere){
	r := shape.radius
	CO := O - shape.center
	//dot product
	// SUM(0i*Di+Oj*Dj)
	a := (D[0]*D[0] + D[1]* D[1])
	b := 2* (CO[0]*D[0] + CO[1]*D[1])
	c := (CO[0]*CO[0] + CO[1]*CO[1]) - r*r

	disc = b * b - 4*a*c 
	if disc < 0 {
		return 99999999,99999999
	}
	t1 = (-1*b + math.sqrt(disc)) / (2*a)
	t2 = (-1*b - math.sqrt(disc)) / (2*a)
	return t1,t2

}
func TraceRay(O []int, D []int, t_min int, t_max int){
	closest_t = 99999999
	closest_sphere = nil 
	for sphere := range scene{
		t1,t2 = IntersectRaySphere(O, D, sphere)
		if (t1 <= t_max && t1 >= t_min) && t1 < closest_t{
			closest_t = t1
			closest_sphere = sphere 
		}
		if (t2 <= t_max && t2 >= t_min) && t2 < closest_t{
			closest_t = t2 
			closest_sphere = sphere 
		}
	}

	if closest_sphere == nil{
		return rl.RayWhite
	}

	return closest_sphere.color
	
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