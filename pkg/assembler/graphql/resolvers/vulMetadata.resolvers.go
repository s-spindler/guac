package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// IngestVulnerabilityMetadata is the resolver for the ingestVulnerabilityMetadata field.
func (r *mutationResolver) IngestVulnerabilityMetadata(ctx context.Context, vulnerability model.VulnerabilityInputSpec, score model.VulnerabilityScoreInputSpec, vulnerabilityMetadata model.VulnerabilityMetadataInputSpec) (string, error) {
	panic(fmt.Errorf("not implemented: IngestVulnerabilityMetadata - ingestVulnerabilityMetadata"))
}

// IngestVulnerabilityMetadatas is the resolver for the ingestVulnerabilityMetadatas field.
func (r *mutationResolver) IngestVulnerabilityMetadatas(ctx context.Context, vulnerabilities []*model.VulnerabilityInputSpec, scores []*model.VulnerabilityScoreInputSpec, vulnerabilityMetadatas []*model.VulnerabilityMetadataInputSpec) ([]string, error) {
	panic(fmt.Errorf("not implemented: IngestVulnerabilityMetadatas - ingestVulnerabilityMetadatas"))
}

// VulnerabilityMetadata is the resolver for the vulnerabilityMetadata field.
func (r *queryResolver) VulnerabilityMetadata(ctx context.Context, vulnerabilityMetadataSpec model.VulnerabilityMetadataSpec) ([]*model.VulnerabilityMetadata, error) {
	panic(fmt.Errorf("not implemented: VulnerabilityMetadata - vulnerabilityMetadata"))
}
