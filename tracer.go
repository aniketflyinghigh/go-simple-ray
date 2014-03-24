package main

import (
  "math"
  "fmt"
  "encoding/json" 
  "io/ioutil"
  "image"
  "image/color"
)

type Scene struct {
  Eye Vector
  Objects [] Sphere
  Lights  [] Light
  AmbientLight Light
}

type Light struct {
  Position Vector
  Color Color
}

type Intersection struct {
  Sphere Sphere
  Ray Ray
  Distance float64
}

func computeColor(hitAngle float64, ray Ray, sphere Sphere, lights []Light, ambientLight Light) Color {
  // if the ray continues on forever, return black
  if hitAngle == math.Inf(1) {
    return Color{0, 0, 0}
  } else {

    // vectorCoefficient := ray.multiply(hitAngle)


    // shadedCoefficient := lights[0].Position.subtract(sphere.Position).norm().multiplyFold(vectorCoefficient.subtract(sphere.Position).norm())


    // color := Color{0, 0, 0}.multip(shadedCoefficient)
    // // shadedCoefficient := sphere.Position.norm().multiplyFold(vectorCoefficient.subtract(sphere.Position).norm())
    // // lighting := Color{0, 0, 0}
    
    // // add ambient light
    // // lighting = lighting.multiply(lights[1].Color)
    
    // // add the scene lights
    // // for _, light := range lights {
    // //   lighting = lighting.add(light.Color)
    // // }

    // // add the lighting to the sphere colors 
    // color := sphere.Color.add(lights[0].Color.multiplyFloat(math.Max(shadedCoefficient, 0)))
    // // color := sphere.Color.add(lighting)
    // // color := sphere.Color.multiply(lighting.multiplyFloat(math.Max(shadedCoefficient, 0)))
    color := sphere.Color
    return color
    
  }
}

func shade(intersection Intersection, scene Scene, depth int) Color {
  return intersection.Sphere.Color
}

func traceRay(ray Ray, scene Scene) Color {
  intersection := closestIntersection(ray, scene)
  fmt.Println(intersection.Distance)
  if intersection.Distance == math.Inf(1) {
    return Color{0, 0, 0};
  } else {
    return shade(intersection, scene, 2);
  }
}

func closestIntersection(ray Ray, scene Scene) Intersection {
  min := Intersection{ Sphere{}, ray, math.Inf(1) }
  for _, element := range scene.Objects {
    intersection := element.intersectRay(ray)
    if intersection.Distance < min.Sphere.intersectRay(ray).Distance {
      min = intersection
    }
  }
  return min
}

func render(scene Scene, width int, height int) *image.RGBA  {
  image := image.NewRGBA(image.Rect(0, 0, width, height)) 
  for x := 0; x <= width; x++ {
    for y := 0; y <= height; y++ {
      ray := Ray{ Vector{float64(x), float64(y), -1000}, scene.Eye }
      finalColor := traceRay(ray, scene)
      color := color.RGBA{uint8(finalColor.Red * 255), uint8(finalColor.Green * 255), uint8(finalColor.Blue * 255), 255}
      image.Set(x, y, color) 
    }
  }
  return image
}

func parseScene(filename string) Scene {
  file, e := ioutil.ReadFile("./scene.json")
  if e != nil {
      fmt.Printf("File error: %v\n", e)
  }
  res := &Scene{}
  
  err := json.Unmarshal(file, &res)
  if err != nil {
    fmt.Println(err)
  }
  return *res
}

