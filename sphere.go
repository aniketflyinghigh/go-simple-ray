package main

import(
  "math"
)

type Sphere struct {
  Position Vector
  Radius int
  Color Color
  Diffuse Color
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