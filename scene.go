package main

import (
  "math"
)

type Color struct {
  Red, Green, Blue float64
}

type Camera struct {
  Position Vector
  Direction Vector
}

type Scene struct {
  Camera Camera
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

type Sphere struct {
  Position Vector
  Radius int
  Color Color
  Diffuse Color
}

type Plane struct {
  Position Vector
  Length int
  Color Color
}

type Ray struct {
  Position Vector 
  Direction Vector
}

func(this Ray) multiply(v float64) Vector {
  return this.Direction.scale(v).add(this.Position)
}

func (this Intersection) Position() Vector {
  return this.Ray.Position.add(this.Ray.Direction.scale(this.Distance))
}

func (this Intersection) Normal() Vector {
  return this.Position().subtract(this.Sphere.Position).unitVector()
}

func (this Color) multiply(that Color) Color {
  return Color{ this.Red * that.Red, this.Green * that.Green, this.Blue * that.Blue }
}

func (this Color) add(that Color) Color {
  return Color{ this.Red + that.Red, this.Green + that.Green, this.Blue + that.Blue }
}

func (this Color) scale(that float64) Color {
  return Color{ this.Red * that, this.Green * that, this.Blue * that }
}

func (this Color) clamp() Color {
  return Color{ math.Min(this.Red, 1.0), math.Min(this.Green, 1.0), math.Min(this.Blue, 1.0)}
}

func (this Plane) intersectRay(ray Ray) Intersection {
  return Intersection{Sphere{}, ray, math.Inf(1)}
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
