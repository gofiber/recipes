// Package main is a placeholder, not a recipe. The real example lives in
// envoy-extauthz/.
//
// Renaming fiber-envoy-extauthz/ to envoy-extauthz/ in d85c8d2 (Nov 2024) left
// GitHub's dependency graph holding manifests for the old path, frozen on the
// fiber v2 they pinned back then. Every new fiber v2 advisory still raises
// alerts against them, and the security update job that would close those
// alerts aborts with dependency_file_not_found because the path is gone.
//
// Re-creating the manifest makes the graph re-scan the path and find a module
// with no dependencies, which clears the stale fiber v2 entry. This directory
// is deleted again afterwards, that time as a delete the graph can observe.
package main

func main() {}
