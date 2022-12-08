// Code generated by github.com/beyondan/gqlgen, DO NOT EDIT.

package fileupload

import (
	"context"

	"github.com/beyondan/gqlgen/example/fileupload/model"
	"github.com/beyondan/gqlgen/graphql"
)

type Stub struct {
	MutationResolver struct {
		SingleUpload              func(ctx context.Context, file graphql.Upload) (*model.File, error)
		SingleUploadWithPayload   func(ctx context.Context, req model.UploadFile) (*model.File, error)
		MultipleUpload            func(ctx context.Context, files []*graphql.Upload) ([]*model.File, error)
		MultipleUploadWithPayload func(ctx context.Context, req []*model.UploadFile) ([]*model.File, error)
	}
	QueryResolver struct {
		Empty func(ctx context.Context) (string, error)
	}
}

func (r *Stub) Mutation() MutationResolver {
	return &stubMutation{r}
}
func (r *Stub) Query() QueryResolver {
	return &stubQuery{r}
}

type stubMutation struct{ *Stub }

func (r *stubMutation) SingleUpload(ctx context.Context, file graphql.Upload) (*model.File, error) {
	return r.MutationResolver.SingleUpload(ctx, file)
}
func (r *stubMutation) SingleUploadWithPayload(ctx context.Context, req model.UploadFile) (*model.File, error) {
	return r.MutationResolver.SingleUploadWithPayload(ctx, req)
}
func (r *stubMutation) MultipleUpload(ctx context.Context, files []*graphql.Upload) ([]*model.File, error) {
	return r.MutationResolver.MultipleUpload(ctx, files)
}
func (r *stubMutation) MultipleUploadWithPayload(ctx context.Context, req []*model.UploadFile) ([]*model.File, error) {
	return r.MutationResolver.MultipleUploadWithPayload(ctx, req)
}

type stubQuery struct{ *Stub }

func (r *stubQuery) Empty(ctx context.Context) (string, error) {
	return r.QueryResolver.Empty(ctx)
}
