package configwrapper_test

import (
	"context"
	"testing"
	"time"

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
