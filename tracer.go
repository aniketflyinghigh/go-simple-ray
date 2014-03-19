package main

import (
  "math"
  // "image/color"
  "fmt"
  "encoding/json"
  // "io"
  "io/ioutil"
)


type Vector struct {
  X, Y, Z float64
}

type Scene struct {
  Eye Vector
  Objects [] Sphere
  Lights  [] Light
  AmbientLight Light
}

type Sphere struct {
  Position Vector
  Radius int
  Color Color
  Diffuse Color
}

type Ray struct {
  Position Vector
  Direction Vector
}

type Light struct {
  Position Vector
  Color Color
}

type Color struct {
  Red, Green, Blue float64
}

func (this Vector) add(that Vector) Vector {
    return Vector{this.X + that.X, this.Y + that.Y, this.Z + that.Z}
}

func (this Vector) subtract(that Vector) Vector {
    return Vector{this.X - that.X, this.Y - that.Y, this.Z - that.Z}
}

func (this Vector) multiply(that Vector) Vector {
    return Vector{this.X * that.X, this.Y * that.Y, this.Z * that.Z}
}

func (this Vector) divide(that Vector) Vector {
    return Vector{this.X / that.X, this.Y / that.Y, this.Z / that.Z}
}

func (this Vector) multiplyFloat(that float64) Vector {
    return Vector{this.X * that, this.Y * that, this.Z * that}
}

func (this Vector) divideFloat(that float64) Vector {
    return Vector{this.X / that, this.Y / that, this.Z / that}
}

func (this Vector) multiplyFold(that Vector) float64 {
    t := this.multiply(that)
    return t.X + t.Y + t.Z
}

func (this Vector) dot(that Vector) float64 {
    return this.X * that.X + this.Y * that.Y + this.Z * that.Z
}

func (this Vector) lengthSquared() float64 {
    return this.multiplyFold(this);
}

func(this Vector) length() float64 {
  return math.Sqrt(this.lengthSquared())
}
func(this Vector) norm() Vector {
  return this.divideFloat(this.length())
}

func(this Ray) multiply(v float64) Vector {
  return this.Position.multiply(this.Direction).multiplyFloat(v)
}

func (this Color) multiply(that Color) Color {
  return Color{ this.Red * that.Red, this.Green * that.Green, this.Blue * that.Blue }
}

func (this Color) add(that Color) Color {
  return Color{ this.Red + that.Red, this.Green + that.Green, this.Blue + that.Blue }
}

func (this Color) multiplyFloat(that float64) Color {
  return Color{ this.Red * that, this.Green * that, this.Blue * that }
}

func(this Sphere) intersectRay(ray Ray) float64 {
  v := this.Position.subtract(ray.Position)
  a := v.multiplyFold(ray.Direction)
  b := v.lengthSquared() - float64(this.Radius * this.Radius)
  c := a * a -b
  if c >= 0 {
    return a - math.Sqrt(c)
  } else {
    return math.Inf(1)
  }
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

func main() {
  res := parseScene("/.scene.json")
  fmt.Println(res)
}

