package main

import "time"

func sleep(s int) {
    select {
          case <-time.After(time.Duration(s) * time.Second):
            return
    }
}
