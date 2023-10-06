//TODO 
//dont short hand type
//auto generate random spehere 
//dont hard code any numbers, go change 

//NEXT 
//light!

package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type sphere struct { 
	radius float32
	color rl.Color 
	center [3]float32
}

func main() {

	
	var sphere1 = new(sphere)
	sphere1.radius = .5
	sphere1.color = rl.Color{255,0,0,255} //
	sphere1.center = [3]float32{0,1,3}
	
	var sphere2 = new(sphere)
	sphere2.radius = .5
	sphere2.color = rl.Color{0,255,0,255} //
	sphere2.center = [3]float32{1,0,4}
	
	var sphere3 = new(sphere)
	sphere3.radius = .5
	sphere3.color = rl.Color{0,0,255,255} //
	sphere3.center = [3]float32{1,1,4}
	
	var scene = []sphere{ *sphere1,*sphere3, *sphere2}
	
	//creates the window
	rl.InitWindow(600, 600, "raylibtard")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	origin := [3]float32{0, 0, 0}
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		for x := -300; x < 300; x++{
			for y := -300 ; y < 300; y++{
				D := CanvasToViewPort(x, y)
				color := TraceRay(origin, D, 1, 99999999, scene)
				rl.DrawPixel(int32(x),int32(y), color)
			}
		}
		rl.EndDrawing()
	}
}
func CanvasToViewPort(x int, y int)[3]float32{
	return [3]float32{float32(x)*float32(1)/float32(600),float32(y)*float32(1)/float32(600),1}
}

func  IntersectRaySphere(O [3]float32, D [3]float32, shape sphere)(float32,float32){
	
	r := shape.radius
	CO := vec_subtract(O, shape.center)
	a := vec_dotproduct(D, D)
	b := 2 * vec_dotproduct(CO, D)
	c := vec_dotproduct(CO, CO) - (r * r)

	disc := (b * b) - (4 * a* c)
	if disc < 0 {
		return 99999999,99999999
	}
	t1 := (-1*b) + (float32(math.Sqrt(float64(disc))) / (2*a))
	t2 := (-1*b) - (float32(math.Sqrt(float64(disc))) / (2*a))
	
	return t1,t2

}

func vec_subtract(vec1 [3]float32, vec2 [3]float32)[3]float32{
	return [3]float32{vec1[0]-vec2[0],vec1[1]-vec2[1],vec1[2]-vec2[2]}
}

func vec_dotproduct(vec1 [3]float32, vec2 [3]float32)float32{
	return (vec1[0]*vec2[0] + vec1[1]*vec2[1] + vec1[2]*vec2[2])
}



func TraceRay(O [3]float32, D [3]float32, t_min float32, t_max float32, scene []sphere)rl.Color{
	var closest_t float32 = 999999999
	var closest_sphere = new(sphere)
	for i := 0; i < len(scene); i++{
		t1,t2 := IntersectRaySphere(O, D, scene[i])
		if t1 < t_max && t1 > t_min && t1 < closest_t{
			closest_t = t1
			closest_sphere = &scene[i]
		}
		if t2 < t_max && t2 > t_min && t2 < closest_t{
			closest_t = t2 
			closest_sphere = &scene[i] 
		}
	}	
	
	if closest_sphere == nil { 
		return rl.RayWhite
	}

	return closest_sphere.color	
}
