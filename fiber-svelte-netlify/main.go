// Package main is a placeholder, not a recipe.
//
// Commit d85c8d2 (Nov 2024) renamed six recipe directories at once. GitHub's
// dependency graph never dropped the manifests under the old paths, so they
// stayed frozen on the fiber v2 they pinned then. Every new fiber v2 advisory
// still raises alerts against paths that no longer exist, and the security
// update job that would close them aborts with dependency_file_not_found.
//
// Re-creating the manifest makes the graph re-scan the path and find a module
// with no dependencies, which clears the stale fiber v2 entry. This directory
// is deleted again afterwards, that time as a delete the graph can observe.
// The real recipe lives in svelte-netlify/.
package main

func main() {}
