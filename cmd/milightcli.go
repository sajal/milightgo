package main

import (
	//"log"
	"github.com/sajal/milightgo"
	"os"
	//"time"
)

func main() {
	c := milight.NewController(os.Args[1])
	c.SetBrightness(1, 5)
	/*
		log.Println(i)
		c.ZoneOff(i)
		time.Sleep(time.Second * 3)
		c.ZoneOn(i)
		time.Sleep(time.Second * 3)
	*/
}
