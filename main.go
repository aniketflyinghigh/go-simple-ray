package main

import (
  "os"
  "image/png"
)

func main() {
  width, height := 640, 480
  scene := parseScene("/.scene.json")
  image := trace(scene, width, height)
  toimg, _ := os.Create("out.jpg")
  png.Encode(toimg, image)
}