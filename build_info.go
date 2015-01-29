package main

// Build information. Populated by Makefile.
var (
	buildVersion  string
	buildRevision string
	buildBranch   string
	buildDate     string
	goVersion     string
)

// BuildInfo encapsulates compile-time metadata about Prometheus made available
// via go tool ld such that this can be reported on-demand.
var BuildInfo = map[string]string{
	"version":    buildVersion,
	"revision":   buildRevision,
	"branch":     buildBranch,
	"date":       buildDate,
	"go_version": goVersion,
}
