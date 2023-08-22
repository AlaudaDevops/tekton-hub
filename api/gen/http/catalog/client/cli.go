// Code generated by goa v3.12.4, DO NOT EDIT.
//
// catalog HTTP client CLI support package
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	catalog "github.com/tektoncd/hub/api/gen/catalog"
)

// BuildRefreshPayload builds the payload for the catalog Refresh endpoint from
// CLI flags.
func BuildRefreshPayload(catalogRefreshCatalogName string, catalogRefreshToken string) (*catalog.RefreshPayload, error) {
	var catalogName string
	{
		catalogName = catalogRefreshCatalogName
	}
	var token string
	{
		token = catalogRefreshToken
	}
	v := &catalog.RefreshPayload{}
	v.CatalogName = catalogName
	v.Token = token

	return v, nil
}

// BuildRefreshAllPayload builds the payload for the catalog RefreshAll
// endpoint from CLI flags.
func BuildRefreshAllPayload(catalogRefreshAllToken string) (*catalog.RefreshAllPayload, error) {
	var token string
	{
		token = catalogRefreshAllToken
	}
	v := &catalog.RefreshAllPayload{}
	v.Token = token

	return v, nil
}

// BuildCatalogErrorPayload builds the payload for the catalog CatalogError
// endpoint from CLI flags.
func BuildCatalogErrorPayload(catalogCatalogErrorCatalogName string, catalogCatalogErrorToken string) (*catalog.CatalogErrorPayload, error) {
	var catalogName string
	{
		catalogName = catalogCatalogErrorCatalogName
	}
	var token string
	{
		token = catalogCatalogErrorToken
	}
	v := &catalog.CatalogErrorPayload{}
	v.CatalogName = catalogName
	v.Token = token

	return v, nil
}
