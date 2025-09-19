package generated

import (
	"context"

	"job-board/backend/graph/model"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

//go:embed "schema.graphqls"
var sourcesBytes []byte

var sources = &ast.Source{
	Name:    "schema.graphqls",
	Input:   string(sourcesBytes),
	BuiltIn: false,
}

type Config struct {
	Resolvers ResolverRoot
}

type ResolverRoot interface {
	Query() QueryResolver
	Mutation() MutationResolver
}

type QueryResolver interface {
	Jobs(ctx context.Context) ([]*model.Job, error)
	Job(ctx context.Context, id string) (*model.Job, error)
	Videos(ctx context.Context) ([]*model.Video, error)
	Video(ctx context.Context, id string) (*model.Video, error)
}

type MutationResolver interface {
	CreateJob(ctx context.Context, input model.JobInput) (*model.Job, error)
	UpdateJob(ctx context.Context, id string, input model.JobInput) (*model.Job, error)
	DeleteJob(ctx context.Context, id string) (bool, error)
	CreateVideo(ctx context.Context, input model.VideoInput) (*model.Video, error)
}

func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{cfg}
}

type executableSchema struct {
	cfg Config
}

func (e *executableSchema) Schema() *ast.Schema {
	// This is a simplified implementation
	// In a real app, you'd parse the schema from sources
	return &ast.Schema{}
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	// This is a simplified implementation
	// In a real app, you'd implement the full GraphQL execution
	return func(ctx context.Context) *graphql.Response {
		return &graphql.Response{
			Data: []byte(`{"data":{}}`),
		}
	}
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, args map[string]interface{}) (int, bool) {
	// Return default complexity for all fields
	return childComplexity, true
}
