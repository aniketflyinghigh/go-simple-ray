package main

import (
  "os"
  "image/png"
)

func main() {
  width, height := 720, 480
  scene := parseScene("/.scene.json")
  image := render(scene, width, height)
  toimg, _ := os.Create("out.jpg")
  png.Encode(toimg, image)
}