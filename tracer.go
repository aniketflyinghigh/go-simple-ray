package main

import (
  "math"
  "fmt"
  "encoding/json" 
  "io/ioutil"
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
    finalColor := sphere.Color.multiply(light.Color.multiplyFloat(math.Max(shadedCoefficient, 0))).add(ambientLight.Color)
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
    fmt.Println(element)
  }
  return min
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

