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

func(this Sphere) intersectRay(ray Ray) Intersection {
  distance := this.Position.subtract(ray.Position)
  a := ray.Direction.multiplyFold(distance)
  b := distance.lengthSquared() - float64(this.Radius * this.Radius)
  c := a * a -b
  if c >= 0 {
    return Intersection{this, ray, a - math.Sqrt(c)}
  } else {
    return Intersection{this, ray, math.Inf(1)}
  }
}
