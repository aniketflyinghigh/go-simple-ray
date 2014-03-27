package main

import (
  "math"
)

type Vector struct {
  X, Y, Z float64
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

func (this Vector) scale(that float64) Vector {
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

func(this Vector) unitVector() Vector {
  return this.scale(1 / this.length())
}

func(this Vector) reflectThrough(normal Vector) Vector {
  d := normal.scale(this.dot(normal))
  return d.scale(2).subtract(this)
}
