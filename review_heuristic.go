package rca

import (
	"context"

	"github.com/dpopsuev/origami/engine"
)

type reviewHeuristic struct{}

func (t *reviewHeuristic) Name() string        { return "review-heuristic" }
func (t *reviewHeuristic) Deterministic() bool { return true }

func (t *reviewHeuristic) Transform(_ context.Context, _ *engine.TransformerContext) (any, error) {
	return map[string]any{"decision": "approve"}, nil
}
