package api

import (
	"testing"

	"github.com/hashicorp/consul/testutil/retry"
)

func TestCoordinate_Datacenters(t *testing.T) {
	t.Parallel()
	c, s := makeClient(t)
	defer s.Stop()

	coordinate := c.Coordinate()
	for r := retry.OneSec(); r.NextOr(t.FailNow); {
		datacenters, err := coordinate.Datacenters()
		if err != nil {
			t.Log(err)
			continue
		}
		if len(datacenters) == 0 {
			t.Logf("Bad: %v", datacenters)
			continue
		}
		break
	}
}

func TestCoordinate_Nodes(t *testing.T) {
	t.Parallel()
	c, s := makeClient(t)
	defer s.Stop()

	coordinate := c.Coordinate()
	for r := retry.OneSec(); r.NextOr(t.FailNow); {
		_, _, err := coordinate.Nodes(nil)
		if err != nil {
			t.Log(err)
			continue
		}
		break

		// There's not a good way to populate coordinates without
		// waiting for them to calculate and update, so the best
		// we can do is call the endpoint and make sure we don't
		// get an error.
	}
}
