package main

type Ray struct {
  Position Vector
  Direction Vector
}

func(this Ray) multiply(v float64) Vector {
  return this.Position.multiply(this.Direction).multiplyFloat(v)
}
