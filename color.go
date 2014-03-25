package main

import (
  "math"
)

type Color struct {
  Red, Green, Blue float64
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