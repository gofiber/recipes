package main

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// complexityLimit works like extension.FixedComplexityLimit, but computes each named fragment once.
// gqlgen walks a fragment again at every spread, which makes nested spreads exponential.
type complexityLimit struct {
	limit int
	es    graphql.ExecutableSchema
}

var _ interface {
	graphql.HandlerExtension
	graphql.OperationContextMutator
} = (*complexityLimit)(nil)

func (*complexityLimit) ExtensionName() string { return "ComplexityLimit" }

func (c *complexityLimit) Validate(es graphql.ExecutableSchema) error {
	c.es = es
	return nil
}

func (c *complexityLimit) MutateOperationContext(ctx context.Context, opCtx *graphql.OperationContext) *gqlerror.Error {
	w := complexityWalker{limit: c.limit, es: c.es, vars: opCtx.Variables, fragments: map[string]int{}}
	if n := w.selectionSet(ctx, opCtx.Operation.SelectionSet); n > c.limit {
		err := gqlerror.Errorf("operation has complexity %d, which exceeds the limit of %d", n, c.limit)
		errcode.Set(err, "COMPLEXITY_LIMIT_EXCEEDED")
		return err
	}
	return nil
}

type complexityWalker struct {
	limit     int
	es        graphql.ExecutableSchema
	vars      map[string]any
	fragments map[string]int // cost of each named fragment
}

func (w *complexityWalker) selectionSet(ctx context.Context, set ast.SelectionSet) int {
	total := 0
	for _, sel := range set {
		var cost int
		switch s := sel.(type) {
		case *ast.Field:
			cost = w.field(ctx, s)
		case *ast.FragmentSpread:
			var ok bool
			if cost, ok = w.fragments[s.Name]; !ok {
				cost = w.selectionSet(ctx, s.Definition.SelectionSet)
				w.fragments[s.Name] = cost
			}
		case *ast.InlineFragment:
			cost = w.selectionSet(ctx, s.SelectionSet)
		}
		// Returning as soon as the limit is passed keeps the sums far from overflow.
		if total += cost; total > w.limit {
			return total
		}
	}
	return total
}

// field returns the custom cost from the schema config, or 1 plus the cost of the selection, as gqlgen does.
// Interfaces are not handled. The schema has none.
func (w *complexityWalker) field(ctx context.Context, f *ast.Field) int {
	child := 0
	if len(f.SelectionSet) > 0 {
		child = w.selectionSet(ctx, f.SelectionSet)
	}
	// A child already over the limit decides the result.
	if child > w.limit {
		return child
	}
	if cost, ok := w.es.Complexity(ctx, f.ObjectDefinition.Name, f.Name, child, f.ArgumentMap(w.vars)); ok && cost >= 1 {
		return cost
	}
	return 1 + child
}
