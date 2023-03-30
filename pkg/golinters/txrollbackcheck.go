package golinters

import (
	"github.com/herpiko/txrollbackcheck/pkg/analyzer"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
)

func NewTxRollbackCheck() *goanalysis.Linter {
	return goanalysis.NewLinter(
		"txrollbackcheck",
		"Checks that sql.Tx are having a defer Rollback().",
		[]*analysis.Analyzer{
			analyzer.NewAnalyzer(),
		},
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
