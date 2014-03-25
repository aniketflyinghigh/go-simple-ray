package main

import (
  "math"
  "fmt"
  "encoding/json" 
  "io/ioutil"
  "image"
  "image/color"
)

type Scene struct {
  Eye Vector
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

func (this Intersection) Position() Vector {
  return this.Ray.Position.add(this.Ray.Direction.scale(this.Distance))
}

func (this Intersection) Normal() Vector {
  return this.Position().subtract(this.Sphere.Position).unitVector()
}

func shade(intersection Intersection, scene Scene) Color {
  lambert := 0.0

  light := scene.AmbientLight
  // Calculate the lambertian reflectance, which is essentially  a 'diffuse' lighting system.
  // direct light is bright, and from there, less direct light is gradually, less light.
  contribution := light.Position.subtract(intersection.Position()).unitVector().dot(intersection.Normal())

  // sometimes this formula can return negatives, so we check:
  // we only want positive values for lighting.
  if contribution > 0 {
    lambert += contribution
  }

  // lambert should never 'blow out' the lighting of an object,
  // even if the ray bounces between a lot of things and hits lights
  lambert = math.Min(1, lambert)
  
  lambertColor := intersection.Sphere.Color.scale(lambert)
  ambientColor := intersection.Sphere.Color.scale(0.1)

  // clamp the colors so it doesn't go create artifacts
  return lambertColor.add(ambientColor).clamp()
}

func traceRay(ray Ray, scene Scene) Color {
  intersection := closestIntersection(ray, scene)
  if intersection.Distance == math.Inf(1) {
    return Color{0, 0, 0};
  } else {
    return shade(intersection, scene);
  }
}

func closestIntersection(ray Ray, scene Scene) Intersection {
  min := Intersection{ Sphere{}, ray, math.Inf(1)}
  for _, element := range scene.Objects {
    intersection := element.intersectRay(ray)
    if intersection.Distance < min.Sphere.intersectRay(ray).Distance {
      min = intersection
    }
  }
  return min
}

func render(scene Scene, width int, height int) *image.RGBA  {
  image := image.NewRGBA(image.Rect(0, 0, width, height)) 
  for x := 0; x <= width; x++ {
    for y := 0; y <= height; y++ {
      ray := Ray{ Vector{float64(x), float64(y), -1000}, scene.Eye }
      finalColor := traceRay(ray, scene)
      color := color.RGBA{uint8(finalColor.Red * 255), uint8(finalColor.Green * 255), uint8(finalColor.Blue * 255), 255}
      image.Set(x, y, color) 
    }
  }
  return image
}

func parseScene(filename string) Scene {
  file, e := ioutil.ReadFile("./scene.json")
  if e != nil {
      fmt.Printf("File error: %v\n", e)
  }
  res := &Scene{}
  
  err := json.Unmarshal(file, &res)
  if err != nil {
    fmt.Println(err)
  }
  return *res
}

