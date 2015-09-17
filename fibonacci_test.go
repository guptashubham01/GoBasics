package main
import "testing" 

type TestPair struct{
    values uint
    res uint
}

var Tests = []TestPair{
    { 10, 55 },
    { 5, 5 },
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
