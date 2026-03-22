package rca

import (
	"github.com/dpopsuev/origami/circuit"
)

// DoneNodeName is the terminal pseudo-node name used in circuit definitions.
const DoneNodeName = "DONE"

// WrapNodeArtifact wraps an artifact as a circuit.Artifact using the node name as type.
func WrapNodeArtifact(nodeName string, artifact any) circuit.Artifact {
	if artifact == nil {
		return nil
	}
	return &bridgeArtifact{
		raw:      artifact,
		typeName: nodeName,
	}
}

type bridgeArtifact struct {
	raw      any
	typeName string
}

func (a *bridgeArtifact) Type() string       { return a.typeName }
func (a *bridgeArtifact) Confidence() float64 { return 0 }
func (a *bridgeArtifact) Raw() any            { return a.raw }


