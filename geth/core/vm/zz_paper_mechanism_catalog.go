//go:build paper_evaluation

package vm

// Bind the bundled executor catalog to this release of the lowering logic.
func init() {
	if !paperWithoutCheckElimination {
		precompiledFlows.version = flowPreparationVersion()
	}
}
