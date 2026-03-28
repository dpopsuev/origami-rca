package rca

// Node names used as step identifiers throughout the RCA circuit.
const (
	nodeRecall      = "recall"
	nodeTriage      = "triage"
	nodeResolve     = "resolve"
	nodeInvestigate = "investigate"
	nodeCorrelate   = "correlate"
	nodeReview      = "review"
	nodeReport      = "report"
)

// String constants used across multiple files.
const (
	valueUnknown = "unknown"
	valueNoSteps = "(no steps)"

	valueInvestigationPending = "investigation pending (no error message available)"

	categoryProduct = "product"
	defectPB001     = "pb001"

	valueRCA = "rca"

	metricM17 = "M17"
	metricM18 = "M18"
	metricM19 = "M19"
	metricM20 = "M20"

	categoryInfra = "infra"
	categoryTest  = "test"

	statusTriaged      = "triaged"
	statusInvestigated = "investigated"
	statusReviewed     = "reviewed"

	tagValueTrue = "true"

	backendLLM  = "llm"
	backendStub = "stub"
)
