package resource

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/hub/api/gen/resource"
	"github.com/tektoncd/hub/api/pkg/testutils"
)

func TestVersionsByID(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.VersionsByIDPayload{ID: 6}
	res, err := resourceSvc.VersionsByID(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, res.Location, "/v1/resource/6/versions")
}

func TestByCatalogKindNameVersion(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByCatalogKindNameVersionPayload{Catalog: "catalog-official", Kind: "task", Name: "tkn", Version: "0.1"}
	res, err := resourceSvc.ByCatalogKindNameVersion(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, res.Location, "/v1/resource/catalog-official/task/tkn/0.1")
}
