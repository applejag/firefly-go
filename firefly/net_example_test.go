package firefly_test

import "github.com/firefly-zero/firefly-go/firefly"

var listOfPlayers []Player

type Player struct {
	Pos   firefly.Point
	Angle firefly.Angle
	Peer  firefly.Peer
}

func ExampleScanPeers() {
	scanner := firefly.ScanPeers()
	for scanner.Scan() {
		listOfPlayers = append(listOfPlayers, Player{
			Pos:   firefly.P(0, 0),
			Angle: firefly.Radians(0),
			Peer:  scanner.Peer(),
		})
	}
}
