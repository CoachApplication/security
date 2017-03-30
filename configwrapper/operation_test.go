package configwrapper_test

import (
	"context"
	"testing"

	"github.com/CoachApplication/base"
	"github.com/CoachApplication/config"
	"github.com/CoachApplication/config/provider"
	"github.com/CoachApplication/config/provider/buffered"
	"github.com/CoachApplication/config/provider/yaml"
	"github.com/CoachApplication/security/configwrapper"
)

func MakeOperationBase(t *testing.T, ctx context.Context) *configwrapper.OperationBase {
	// Build a config provider based on a buffered backend, and a yaml config factory
	con := buffered.NewSingle("user", "default", cuTestBytes) // us a single key-scope buffered connector
	fac := yaml.NewFactory(con)                               // use a yaml factory
	usa := (&provider.AllBackendUsage{}).BackendUsage()       // for now lie and say that this backend is for all the things
	b := provider.NewCompositeBackend(con, usa, fac)          // turn all of those into a provider backend

	p := provider.NewBackendConfigProvider() // Create a provider
	p.Add(b.Backend())                       // add our backend to that provider

	ops := base.NewOperations()                                  // Create a new list of operations for the config wrapper
	ops.Add(provider.NewGetOperation(p.Provider()).Operation())  // add get operation
	ops.Add(provider.NewListOperation(p.Provider()).Operation()) // add list operation

	wr := config.NewStandardWrapper(ops.Operations(), ctx).Wrapper() // create a Config Wrapper from all of this.

	return configwrapper.NewOperationBase(wr) // Now create our base security operation from the wrapper
}
