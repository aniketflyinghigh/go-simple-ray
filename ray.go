package main

type Ray struct {
  Position Vector
  Direction Vector
}

func(this Ray) multiply(v float64) Vector {
  return this.Direction.multiplyFloat(v).add(this.Position)
}