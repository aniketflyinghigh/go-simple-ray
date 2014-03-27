package main

import (
  "os"
  "image/png"
  "fmt"
)

func main() {
  width, height := 1920, 1200
  scene := parseScene("scene_hd.json")
  fmt.Println(scene.Objects[0].Radius)
  image := render(scene, width, height)
  toimg, _ := os.Create("out.jpg")
  png.Encode(toimg, image)
}