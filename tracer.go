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

func computeColor(hitAngle float64, ray Ray, sphere Sphere, light Light, ambientLight Light) Color {
  if hitAngle == math.Inf(1) {
    return Color{0, 0, 0}
  } else {
    vectorCoefficient := ray.multiply(hitAngle)
    shadedCoefficient := light.Position.subtract(sphere.Position).norm().multiplyFold(vectorCoefficient.subtract(sphere.Position).norm())
    finalColor := sphere.Color.multiply(light.Color.multiplyFloat(math.Max(shadedCoefficient, 0)).add(ambientLight.Color))
    return finalColor
  }
}

func getObjectColor(scene Scene, obj Sphere, ray Ray) Color {
  hitAngle := obj.intersectRay(ray)
  return computeColor(hitAngle, ray, obj, scene.Lights[0], scene.AmbientLight)
}

func getClosestSphere(ray Ray, spheres []Sphere) Sphere {
  min := Sphere{}
  for _, element := range spheres {
    intersectedRay := element.intersectRay(ray)
    if intersectedRay < min.intersectRay(ray) {
      min = element
    }
  }
  return min
}

func trace(scene Scene, width int, height int) *image.RGBA  {
  image := image.NewRGBA(image.Rect(0, 0, width, height)) 
  for x := 0; x <= width; x++ {
    for y := 0; y <= height; y++ {
      ray := Ray{ Vector{float64(x), float64(y), -1000}, Vector{0, 0, 1}.norm() }
      closestSphere := getClosestSphere(ray, scene.Objects)
      finalColor := getObjectColor(scene, closestSphere, ray)
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

