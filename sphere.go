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

// http://www.macwright.org/literate-raytracer/
func(this Sphere) intersectRay(ray Ray) Intersection {
  eyeToCenter := this.Position.subtract(ray.Position)
  rayToEdgeOfSphere := eyeToCenter.dot(ray.Direction)
  rayToCenterOfSphere := eyeToCenter.dot(eyeToCenter)
  discriminant := float64(this.Radius * this.Radius) - rayToCenterOfSphere + (rayToEdgeOfSphere * rayToEdgeOfSphere)
  if discriminant < 0 {
    return Intersection{this, ray, math.Inf(1)}
  } else {
    return Intersection{this, ray, rayToEdgeOfSphere - math.Sqrt(discriminant)}
  }
}
   