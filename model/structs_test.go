package model

import (
	"testing"

	"github.com/docker/docker/api/types/swarm"
)

func TestToNormalCPU(t *testing.T) {
	tests := []struct {
		input int64
		want  float64
	}{
		{500000000, 0.5},
		{1000000000, 1},
		{1500000000, 1.5},
		{2000000000, 2},
	}
	for _, test := range tests {
		r := resources{NanoCPUs: test.input}
		actual := r.ToNormalCPU()
		if actual != test.want {
			t.Errorf("Got %f, expected: %f", actual, test.want)
		}
	}
}

func TestManagerReachability(t *testing.T) {
	tests := []struct {
		input swarm.Reachability
		want  swarm.Reachability
	}{
		{swarm.ReachabilityReachable, swarm.ReachabilityReachable},
		{swarm.ReachabilityUnknown, swarm.ReachabilityUnknown},
		{swarm.ReachabilityUnreachable, swarm.ReachabilityUnreachable},
	}
	for _, test := range tests {
		nm := NodeMetrics{ManagerInfo: managerInfo{Reachability: string(test.input)}}
		actual := nm.ManagerReachability()
		if actual != string(test.want) {
			t.Errorf("Got %s, expected: %s", actual, string(test.want))
		}
	}
	// check if below is able to work
	// nm := NodeMetrics{ManagerInfo: managerInfo{}}
	// actual := nm.ManagerReachability()
	// if actual != string(swarm.ReachabilityUnreachable) {
	// 	t.Errorf("Got %s, expected: %s", actual, string(swarm.ReachabilityUnreachable))
	// }
}

func TestGetSanitizedID(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"abc", "abc"},
		{"a:b;c", "a_b_c"},
		{"a-b~c", "a_b_c"},
		{"a,b.c", "a_b_c"},
		{"a b=c", "a_b_c"},
		{"a+b/c", "a_b_c"},
		{"a<b>c", "a_b_c"},
	}
	for _, test := range tests {
		sm := SwarmMetrics{ID: test.input}
		actual := sm.GetSanitizedID()
		if actual != test.want {
			t.Errorf("Got %s, expected: %s", actual, test.want)
		}
	}
}
