// +build int

package main

import "testing"

func TestCluster(t *testing.T) {
	testCluster(t,
		func(c *client) {
			// c.DoLoosly("CLUSTER", "SLOTS")
			c.DoLoosly("CLUSTER", "KEYSLOT", "{test}")
			c.DoLoosly("CLUSTER", "NODES")
			c.Do("CLUSTER")
		},
	)
}
