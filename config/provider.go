/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	// "github.com/vdveldet/provider-github/config/null"
	"github.com/vdveldet/provider-github/config/branch"
	"github.com/vdveldet/provider-github/config/repository"
)

const (
	resourcePrefix = "github"
	modulePath     = "github.com/vdveldet/provider-github"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		repository.Configure,
		branch.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
