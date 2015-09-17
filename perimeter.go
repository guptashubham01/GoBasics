package main

import "math"

type Shapes interface {
    perimeter() float64
    area() float64
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}

type Circle struct {
    x, y, r float64
}


func (r Rectangle) perimeter() float64 {
    length := distance(r.x1, r.y1, r.x1, r.y2)
    width := distance(r.x1, r.y1, r.x2, r.y1)
    return 2 * (length + width)
}

func (r Rectangle) area() float64 {
    length := distance(r.x1, r.y1, r.x1, r.y2)
    width := distance(r.x1, r.y1, r.x2, r.y1)
    return length * width
}

func distance(x1, y1, x2, y2 float64) float64 {
    q := x2 - x1
    w := y2 - y1
    return math.Sqrt(q*q + w*w)
}

func (c Circle) area() float64 {
    return math.Pi * c.r*c.r
}

func (c Circle) perimeter() float64 {
    return 2 * math.Pi * c.r
}
