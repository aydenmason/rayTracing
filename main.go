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

	var sphere1 = new(sphere)
	sphere1.radius = 1
	sphere1.color = [3]uint8{255,0,0}
	sphere1.center = [3]float32{0,-1,3}

	var sphere2 = new(sphere)
	sphere2.radius = 1
	sphere2.color = [3]uint8{0,255,0}
	sphere2.center = [3]float32{-2,0,4}
	
	var sphere3 = new(sphere)
	sphere3.radius = 1
	sphere3.color = [3]uint8{0,0,255}
	sphere3.center = [3]float32{2,-0,4}


	var scene = []sphere{*sphere1,*sphere2,*sphere3}
	
	rl.InitWindow(1920, 1080, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	origin := []float32{0, 0, 0}
	
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		for x := -960; x < 960; x++{
			for y := -540 ; y < 540; y++{
				D := CanvasToViewPort(x, y)
				color := TraceRay(origin, D, 1, 9999999, scene)
			}
		}
		//draw_triangle(0,0,0,50,50,50)

		rl.EndDrawing()
	}
}
func CanvasToViewPort(x int, y int)[]float32{
	return []float32{float32(x),float32(y),2}
}


func  IntersectRaySphere(O []float32, D []float32, shape sphere){
	r := shape.radius

	a := (D[0]*D[0] + D[1]* D[1])
	b := 2* (CO[0]*D[0] + CO[1]*D[1])
	c := (CO[0]*CO[0] + CO[1]*CO[1]) - r*r

	disc := b * b - 4*a*c 
	if disc < 0 {
		return []float32{99999999,99999999}
	}
	t1 := (-1*b + math.sqrt(disc)) / (2*a)
	t2 := (-1*b - math.sqrt(disc)) / (2*a)
	return t1,t2

}

func vec_subtract(vec1 [3]float32, vec2 [3]float32)[3]float32{
	return [3]float32{(vec1[0]-vec2[0],vec1[1]-vec2[1],vec1[2]-vec2[2])}
}

func vec_dotproduct(vec1 [3]float32, vec2 [3]float32)float32{
	return (vec1[0]*vec2[0] + vec1[1]*vec2[1] + vec1[2]*vec2[2])
}

func TraceRay(O []float32, D []float32, t_min int, t_max int, scene []sphere)rl.color{
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