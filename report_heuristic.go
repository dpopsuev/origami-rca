package rca

import (
	"context"

	"github.com/dpopsuev/origami/engine"
)

type reportHeuristic struct{}

func (t *reportHeuristic) Name() string        { return "report-heuristic" }
func (t *reportHeuristic) Deterministic() bool { return true }

func (t *reportHeuristic) Transform(_ context.Context, tc *engine.TransformerContext) (any, error) {
	fp := failureFromContext(tc.WalkerState)
	caseLabel, _ := tc.WalkerState.Context[KeyCaseLabel].(string)
	return map[string]any{
		"case_id":   caseLabel,
		"test_name": fp.name,
		"summary":   "automated baseline analysis",
	}, nil
}
