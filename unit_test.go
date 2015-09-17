package main
import "testing" 
import "time"


type TestPair struct{
    values uint
    res uint
}

type TimeDetails struct {
        secs int
        expect time.Duration
}

var Tests = []TestPair{
    { 10, 55 },
    { 5, 5 },
}

var TestsTime = []TimeDetails {
        {2, 2*time.Second}, {5, 5*time.Second}, 
}

func TestFib(t *testing.T){
    for _, pair := range Tests {
        v := fib(pair.values)
        if v != pair.res {
            t.Error(
                 "For", pair.values,
                "expected", pair.res,
                "got", v,
            )
        }
    }
}

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

func TestSleep(t *testing.T) {
    for _, timer := range TestsTime {
        startTime := time.Now()
        sleep(timer.secs)
        stopTime := time.Now()
        var timeDelay time.Duration = stopTime.Sub(startTime)   
        if timeDelay < timer.expect{
            t.Errorf("Test case Failed as the expected output did not meet the actual output")
        }
    }
}
