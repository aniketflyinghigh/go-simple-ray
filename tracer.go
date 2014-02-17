package main

import (
  "fmt"
  "os"
  "math"
)

type Vector struct {
  x, y, z float64
}

func (this Vector) add(that Vector) Vector {
    return Vector{this.x + that.x, this.y + that.y, this.z + that.z}
}

func (this Vector) subtract(that Vector) Vector {
    return Vector{this.x - that.x, this.y - that.y, this.z - that.z}
}

func (this Vector) multiply(that Vector) Vector {
    return Vector{this.x * that.x, this.y * that.y, this.z * that.z}
}

func (this Vector) divide(that Vector) Vector {
    return Vector{this.x / that.x, this.y / that.y, this.z / that.z}
}

func (this Vector) multiplyFloat(that float64) Vector {
    return Vector{this.x * that, this.y * that, this.z * that}
}

func (this Vector) multiplyFold(that Vector) float64 {
    t := this.multiply(that)
    return t.x + t.y + t.z
}

func (this Vector) dot(that Vector) float64 {
    return this.x * that.x + this.y * that.y + this.z * that.z
}

func (this Vector) lengthSquared() float64 {
    return this.multiplyFold(this);
}

func(this Vector) length() float64 {
  return math.Sqrt(this.lengthSquared())
}



func main() {
  sceneFile := os.Args[len(os.Args)-1]
  fmt.Printf("%s\n", sceneFile)


  c := Vector{1, 2, 3}

  d := Vector{4, 5, 6}

  fmt.Printf("%s\n", c.add(d))
}

