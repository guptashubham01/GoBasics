package main
import "testing" 
import "time"

type TimeDetails struct {
        secs int
        expect time.Duration
}

var TestsTime = []TimeDetails {
        {2, 2*time.Second}, {5, 5*time.Second}, 
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
