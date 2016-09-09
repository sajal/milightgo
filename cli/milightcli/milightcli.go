package main

import (
	"github.com/codegangsta/cli"
	"github.com/sajal/milightgo"
	"log"
	"os"
)

func check(ip string) {
	if ip == "" {
		log.Fatal("--ip must be provided")
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Mi-Light"
	app.Usage = "Control Mi Light"
	app.Action = func(c *cli.Context) {
		println("Hello friend!", c.Args()[0])
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "ip",
			Value: "",
			Usage: "IP address of controler",
		},
		cli.IntFlag{
			Name:  "zone",
			Value: 0,
			Usage: "Zone to act upon",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "on",
			Usage: "Turn on a LED",
			Action: func(c *cli.Context) {
				ip := c.GlobalString("ip")
				check(ip)
				zone := c.GlobalInt("zone")
				println("ip: ", ip, "Zone: ", zone)
				m := milight.NewController(ip)
				err := m.ZoneOn(zone)
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:  "off",
			Usage: "Turn off a LED",
			Action: func(c *cli.Context) {
				ip := c.GlobalString("ip")
				check(ip)
				zone := c.GlobalInt("zone")
				println("ip: ", ip, "Zone: ", zone)
				m := milight.NewController(ip)
				err := m.ZoneOff(zone)
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:  "bright",
			Usage: "Increase brightness",
			Action: func(c *cli.Context) {
				ip := c.GlobalString("ip")
				check(ip)
				zone := c.GlobalInt("zone")
				println("ip: ", ip, "Zone: ", zone)
				m := milight.NewController(ip)
				err := m.ZoneBright(zone)
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:  "dim",
			Usage: "Decrease brightness",
			Action: func(c *cli.Context) {
				ip := c.GlobalString("ip")
				check(ip)
				zone := c.GlobalInt("zone")
				println("ip: ", ip, "Zone: ", zone)
				m := milight.NewController(ip)
				err := m.ZoneDim(zone)
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:  "warm",
			Usage: "Increase warmth",
			Action: func(c *cli.Context) {
				ip := c.GlobalString("ip")
				check(ip)
				zone := c.GlobalInt("zone")
				println("ip: ", ip, "Zone: ", zone)
				m := milight.NewController(ip)
				err := m.ZoneWarm(zone)
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:  "cold",
			Usage: "Decrease warmth",
			Action: func(c *cli.Context) {
				ip := c.GlobalString("ip")
				check(ip)
				zone := c.GlobalInt("zone")
				println("ip: ", ip, "Zone: ", zone)
				m := milight.NewController(ip)
				err := m.ZoneCold(zone)
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:  "setbrightness",
			Usage: "Set brightness (1 - 10)",
			Action: func(c *cli.Context) {
				ip := c.GlobalString("ip")
				check(ip)
				zone := c.GlobalInt("zone")
				brightness := c.Int("level")
				println("ip: ", ip, "Zone: ", zone, "Brightness: ", brightness)
				m := milight.NewController(ip)
				err := m.SetBrightness(zone, brightness)
				if err != nil {
					log.Fatal(err)
				}
			},
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "level",
					Value: 0,
					Usage: "brightness level ( 1 - 10 )",
				},
			},
		},
		{
			Name:  "setwarmth",
			Usage: "Set setwarmth (1 - 10)",
			Action: func(c *cli.Context) {
				ip := c.GlobalString("ip")
				check(ip)
				zone := c.GlobalInt("zone")
				warmth := c.Int("level")
				println("ip: ", ip, "Zone: ", zone, "Warmth: ", warmth)
				m := milight.NewController(ip)
				err := m.SetWarmth(zone, warmth)
				if err != nil {
					log.Fatal(err)
				}
			},
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "level",
					Value: 0,
					Usage: "warmth level ( 1 - 10 )",
				},
			},
		},
	}
	app.Run(os.Args)
}
