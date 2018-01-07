package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
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

	severity := TraverseFirewall(*firewall, 0)
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

	severity := TraverseFirewall(*firewall, 10)
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

	delay := TraverseFirewallWithoutGettingCought(*firewall)
	assert.Equal(t, 10, delay)
}

func TestClone(t *testing.T) {
	firewall := &Firewall{
		time: 10,
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

	clone := firewall.Clone()

	firewall.time = 20
	assert.Equal(t, 10, clone.time)
	assert.Equal(t, 20, firewall.time)

	firewall.layers = []*Layer{
		NewLayer(0, 10),
		NewLayer(1, 20),
	}
	assert.Equal(t, 7, len(clone.layers))
	assert.Equal(t, 2, len(firewall.layers))

	firewall.layers[0].depth = 15
	assert.Equal(t, 15, firewall.layers[0].depth)
	assert.Equal(t, 0, clone.layers[0].depth)

	fmt.Println(firewall)
	fmt.Println(clone)
}
