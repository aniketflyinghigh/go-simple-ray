package main

type Color struct {
  Red, Green, Blue float64
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