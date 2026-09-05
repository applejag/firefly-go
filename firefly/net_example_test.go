package firefly_test

import "github.com/firefly-zero/firefly-go/firefly"

func ExampleScanPeers() {
	scanner := firefly.ScanPeers()
	for scanner.Scan() {
		peer := scanner.Peer()
		_ = peer
	}
}
