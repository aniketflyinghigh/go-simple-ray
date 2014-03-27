package main

import (
  "math"
  "fmt"
  "encoding/json" 
  "io/ioutil"
  "image"
  "image/color"
)

func shade(intersection Intersection, scene Scene, depth int) Color {
  lambert := 0.0
  reflect := Color{0, 0, 0}

  for _, light := range scene.Lights {

    if !isLightVisible(intersection, scene, light){
      continue
    }
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
  }

  reflectedThrough := intersection.Ray.Direction.reflectThrough(intersection.Normal())
  reflectedRay := Ray{intersection.Normal(), reflectedThrough}

  reflectedColor := traceRay(reflectedRay, scene, depth + 1)
  blank := Color {0, 0, 0}
  if reflectedColor != blank  {
    reflect = reflect.add(reflectedColor.scale(intersection.Sphere.Specular))
  }

  lambertColor := intersection.Sphere.Color.scale(lambert)

  // combine the sphere color and the ambient light contribution
  ambientColor := intersection.Sphere.Color.multiply(scene.AmbientLight.Color)
 
  // clamp the colors so it doesn't go create artifacts 
  return lambertColor.add(ambientColor).add(reflect).clamp()
}

func isLightVisible(intersection Intersection, scene Scene, light Light) bool {
  ray := Ray{intersection.Position(), intersection.Position().subtract(light.Position).unitVector()}
  distObject := closestIntersection(ray, scene)
  return distObject.Distance > -0.001
}

func traceRay(ray Ray, scene Scene, depth int) Color {
  if depth > 1 {
    return Color{0, 0, 0}
  }
  intersection := closestIntersection(ray, scene)
  if intersection.Distance == math.Inf(1) {
    return Color{0, 0, 0};
  } else {
    return shade(intersection, scene, depth);
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
      ray := Ray{ Vector{float64(x), float64(y), -1000}, scene.Camera.Direction }
      finalColor := traceRay(ray, scene, 0)
      color := color.RGBA{uint8(finalColor.Red * 255), uint8(finalColor.Green * 255), uint8(finalColor.Blue * 255), 255}
      image.Set(x, y, color) 
    }
  } 
  return image
}

func parseScene(filename string) Scene {
  file, e := ioutil.ReadFile(filename)
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