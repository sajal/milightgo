[![GoDoc](https://godoc.org/github.com/sajal/milightgo?status.svg)](https://godoc.org/github.com/sajal/milightgo)
# milightgo
Control Mi Light using cli. Currently only supports White light because thats what I got.

I tested it with :-

1. [9W Wifi 2.4G Dimmable Mi Light](http://www.aliexpress.com/item/E27-SMD-5730-6W-9W-Wifi-2-4G-Dimmable-Mi-Light-LED-Bulb-AC-85-265V/32318726954.html)
2. [Wifi LED Remote Controller](http://www.aliexpress.com/item/2015-New-Dimmable-Mi-Light-Wireless-Wifi-LED-Remote-Controller-for-2-4G-Smart-Light-RGBW/32318146114.html)

Total project cost: $92.43

Cli reference

IP address is required, either the controlers specific ip or the broadcast ip for your network.

Install

	go install github.com/sajal/milightgo/cli/milightcli

Examples

Turn on all lights

	milightcli --ip="192.168.1.255:8899" on

Turn on specific zone

	milightcli --ip="192.168.1.255:8899" --zone=2 on

Turn off all lights

	milightcli --ip="192.168.1.255:8899" off

Turn off specific zone

	milightcli --ip="192.168.1.255:8899" --zone=2 off

Reduce brightness of all lights

	milightcli --ip="192.168.1.255:8899" dim

Reduce brightness of specific zone

	milightcli --ip="192.168.1.255:8899" --zone=2 dim

Increase brightness of all lights

	milightcli --ip="192.168.1.255:8899" bright

Increase brightness of specific zone

	milightcli --ip="192.168.1.255:8899" --zone=2 bright

Set brightness of all lights (levels 1 - 10)

	milightcli --ip="192.168.1.255:8899" setbrightness --level=4

Set brightness of specific zone (levels 1 - 10)

	milightcli --ip="192.168.1.255:8899" --zone=2 setbrightness --level=4



