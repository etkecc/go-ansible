package ansible

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

var testInventory = &Inventory{Hosts: map[string]*Host{
	"customer.host": {
		Name:       "customer.host",
		Group:      "customers",
		Host:       "321.321.321.321",
		Port:       22,
		User:       "root",
		PrivateKey: "/from/group/vars",
	},
	"setup.host": {
		Name:       "setup.host",
		Group:      "setup",
		Host:       "123.123.123.123",
		Port:       123,
		User:       "setup-user",
		BecomePass: "a(*sEtuP1pass_\"5wOrD",
		PrivateKey: "/from/group/vars",
	},
	"todo.host": {
		Name:       "todo.host",
		Group:      "setup",
		Host:       "TODO",
		Port:       22,
		User:       "root",
		PrivateKey: "/from/group/vars",
	},
}}

func TestNewHostsFile(t *testing.T) {
	expected := testInventory

	actual, err := NewHostsFile("testdata/hosts", nil)
	if err != nil {
		t.Error("file parsing failed", err)
	}
	if actual == nil { //nolint:staticcheck // yes, that's the point
		t.Error("parsed hosts file is nil")
		return
	}
	t.Log("actual:")
	for k, v := range actual.Hosts { //nolint:staticcheck // that's intended
		t.Log(k, "=", fmt.Sprintf("%+v", v))
	}
	if len(expected.Hosts) != len(actual.Hosts) { //nolint:staticcheck // that's intended
		t.Error("length is not expected. Expected:", len(expected.Hosts), "actual:", len(actual.Hosts))
	}
	a := actual.Hosts
	for name, host := range expected.Hosts {
		equal(t, host.Name, a[name].Name)
		equal(t, host.Group, a[name].Group)
		equal(t, host.Host, a[name].Host)
		equal(t, host.User, a[name].User)
		equal(t, host.Port, a[name].Port)
		equal(t, host.SSHPass, a[name].SSHPass)
		equal(t, host.BecomePass, a[name].BecomePass)
		equal(t, host.PrivateKey, a[name].PrivateKey)
	}
}

func TestNewHostsFile_NoExists(t *testing.T) {
	_, err := NewHostsFile("testdata/not-exists", nil)

	if err == nil {
		t.Error("error is nil")
	}
	if !errors.Is(err, os.ErrNotExist) {
		t.Error("error is not os.ErrNotExists, it is", err)
	}
}

func TestMatch(t *testing.T) {
	inv := testInventory
	expected := inv.Hosts["setup.host"]

	actual := inv.Match("setup.host")
	if actual == nil { //nolint:staticcheck // yes, that's the point
		t.Error("matched host is nil")
		return
	}
	if expected.Name != actual.Name { //nolint:staticcheck // that's intended
		t.Error("name is invalid. Expected: '" + expected.Name + "', actual: '" + actual.Name + "'")
	}
}

func TestMerge(t *testing.T) {
	expected := testInventory

	actual := &Inventory{}
	actual.Merge(testInventory)

	t.Log("actual:")
	for k, v := range actual.Hosts {
		t.Log(k, "=", fmt.Sprintf("%+v", v))
	}
	if len(expected.Hosts) != len(actual.Hosts) {
		t.Error("length is not expected. Expected:", len(expected.Hosts), "actual:", len(actual.Hosts))
	}
	a := actual.Hosts
	for name, host := range expected.Hosts {
		equal(t, host.Name, a[name].Name)
		equal(t, host.Group, a[name].Group)
		equal(t, host.Host, a[name].Host)
		equal(t, host.User, a[name].User)
		equal(t, host.Port, a[name].Port)
		equal(t, host.SSHPass, a[name].SSHPass)
		equal(t, host.BecomePass, a[name].BecomePass)
		equal(t, host.PrivateKey, a[name].PrivateKey)
	}
}

func TestHasTODOs(t *testing.T) {
	if !testInventory.Hosts["todo.host"].HasTODOs() {
		t.Error("host has no TODOs")
	}
}

func equal[T comparable](t *testing.T, expected, actual T) {
	t.Helper()

	if expected != actual {
		t.Error("value is not equal. Expected:", expected, ", actual:", actual)
	}
}
