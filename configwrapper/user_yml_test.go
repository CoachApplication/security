package configwrapper

import (
	"context"
	"github.com/CoachApplication/config/provider/buffered"
	"github.com/CoachApplication/config/provider/yaml"
	"testing"
	"time"
)

var cuTest = ConfigUser{
	FromId:   "test",
	FromName: "Test User",
}
var cuTestBytes = []byte(`
Id: test
Name: Test User
`)

func TestConfigUser_YmlGet(t *testing.T) {
	dur, _ := time.ParseDuration("2s")

	c := yaml.NewConfig("key", "scope", buffered.NewSingle("key", "scope", cuTestBytes))

	var cu ConfigUser
	res := c.Get(&cu)

	ctx, _ := context.WithTimeout(context.Background(), dur)
	select {
	case <-res.Finished():

		if cu.Id() != "test" {
			t.Error("ConfigUser provided the wrong id: ", cu)
		}
		if cu.Name() != "Test User" {
			t.Error("ConfigUser provided the wrong name: ", cu)
		}

	case <-ctx.Done():
		t.Error("ConfigUser marshalling timed out: ", ctx.Err().Error())
	}

}
func TestConfigUser_YmlSet(t *testing.T) {
	dur, _ := time.ParseDuration("2s")

	con := buffered.NewSingle("key", "scope", []byte{})
	c := yaml.NewConfig("key", "scope", con)

	res := c.Set(cuTest)

	ctx, _ := context.WithTimeout(context.Background(), dur)
	select {
	case <-res.Finished():

		if !res.Success() {
			t.Error("ConfigUser set failed: ", res.Errors())
		} else {

			var cu ConfigUser
			res = c.Get(&cu)

			select {
			case <-res.Finished():

				if cu.Id() != "test" {
					t.Error("ConfigUser provided the wrong id: ", cu)
				}
				if cu.Name() != "Test User" {
					t.Error("ConfigUser provided the wrong name: ", cu)
				}

			case <-ctx.Done():
				t.Error("ConfigUser marshalling timed out: ", ctx.Err().Error())
			}

		}

	case <-ctx.Done():
		t.Error("ConfigUser marshalling timed out: ", ctx.Err().Error())
	}

}
