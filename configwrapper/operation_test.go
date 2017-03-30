package configwrapper_test

import (
	"context"
	"testing"
	"time"

	"github.com/CoachApplication/base"
	"github.com/CoachApplication/config"
	"github.com/CoachApplication/config/provider"
	"github.com/CoachApplication/config/provider/buffered"
	"github.com/CoachApplication/config/provider/yaml"
	"github.com/CoachApplication/security"
	"github.com/CoachApplication/security/configwrapper"
	"reflect"
)

var cuTest = configwrapper.ConfigUser{
	FromId:   "test",
	FromName: "Test User",
}
var cuTestBytes = []byte(`
Id: test
Name: Test User
`)

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

func TestUserOperation_Exec(t *testing.T) {
	dur, _ := time.ParseDuration("2s")
	ctx, _ := context.WithTimeout(context.Background(), dur)
	b := MakeOperationBase(t, ctx)

	gOp := (&configwrapper.UserOperation{OperationBase: *b}).Operation()

	res := gOp.Exec(gOp.Properties())

	select {
	case <-res.Finished():

		if uPr, err := res.Properties().Get(security.PROPERTY_ID_USER); err != nil {
			t.Error("UserOperation did not provide a user Property: ", err.Error())
		} else if !uPr.Validate() {
			t.Error("UserOperation User Property says that it is invalid.")
		} else if u, good := uPr.Get().(security.User); !good {
			t.Error("UserOperation User property returned invalid data: ", reflect.TypeOf(interface{}(uPr.Get())))
		} else {

			if u.Name() != cuTest.Name() {
				t.Error("UserOperation User Property has a bad Name: ", u.Name())
			}
			if u.Id() != cuTest.Id() {
				t.Error("UserOperation User Property has a bad Id: ", u.Id())
			}

		}

	case <-ctx.Done():
		t.Error("UserOperation exec timed out: ", ctx.Err().Error())
	}
}
