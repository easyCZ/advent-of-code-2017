package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLayerMovementSizeZero(t *testing.T) {
	layer := &Layer{
		depth:     0,
		direction: upDirection,
		position:  0,
		size:      0,
	}
	layer.Move()
	assert.Equal(t, 0, layer.position)
}

func TestLayerMovement(t *testing.T) {
	layer := &Layer{
		depth:     0,
		direction: upDirection,
		position:  0,
		size:      3,
	}
	layer.Move()
	assert.Equal(t, 1, layer.position)
	assert.Equal(t, upDirection, layer.direction)

	layer.Move()
	assert.Equal(t, 2, layer.position)
	assert.Equal(t, downDirection, layer.direction)

	layer.Move()
	assert.Equal(t, 1, layer.position)
	assert.Equal(t, downDirection, layer.direction)

	layer.Move()
	assert.Equal(t, 0, layer.position)
	assert.Equal(t, upDirection, layer.direction)
}

func TestTraverseFirewall(t *testing.T) {
	firewall := &Firewall{
		layers: []*Layer{
			NewLayer(0, 3),
			NewLayer(1, 2),
			nil,
			nil,
			NewLayer(4, 4),
			nil,
			NewLayer(6, 4),
		},
	}

	severity := TraverseFirewall(firewall, 0)
	assert.Equal(t, 24, severity)
}

func TestTraverseFirewallWithDelayOf10(t *testing.T) {
	firewall := &Firewall{
		layers: []*Layer{
			NewLayer(0, 3),
			NewLayer(1, 2),
			nil,
			nil,
			NewLayer(4, 4),
			nil,
			NewLayer(6, 4),
		},
	}

	severity := TraverseFirewall(firewall, 10)
	assert.Equal(t, 0, severity)
}

func TestTraverseFirewallWithoutGettingCought(t *testing.T) {
	firewall := &Firewall{
		layers: []*Layer{
			NewLayer(0, 3),
			NewLayer(1, 2),
			nil,
			nil,
			NewLayer(4, 4),
			nil,
			NewLayer(6, 4),
		},
	}

	delay := TraverseFirewallWithoutGettingCought(firewall)
	assert.Equal(t, 10, delay)
}
