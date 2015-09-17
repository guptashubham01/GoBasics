package main
import "testing" 

func TestRectanglePerimeter (t *testing.T) {
    r := Rectangle{x1: 1, y1: 1, x2:10, y2:10}
    var x Shapes
    x = r 
    perimeterRectangle :=  x.perimeter()
    if perimeterRectangle != 36 {
        t.Error("expected 36 and got ", perimeterRectangle)
    }
}

func TestCirclePerimeter (t *testing.T) {
    c := Circle{x:1, y:1, r: 5}
    var x Shapes
    x = c 
    perimeterCircle := x.perimeter()
    if perimeterCircle != 31.41592653589793 {
        t.Error("expected 31.41592653589793 and got ", perimeterCircle)
    }
}