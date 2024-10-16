package ansible

import "testing"

var testAnsibleInventory = &Inventory{Hosts: map[string]*Host{
	"customer.host": {
		Name:        "customer.host",
		Group:       "customers",
		Host:        "321.321.321.321",
		Port:        22,
		User:        "root",
		PrivateKeys: []string{"/from/group/vars", "/from/host/vars"},
	},
	"setup.host": {
		Name:        "setup.host",
		Group:       "setup",
		Host:        "123.123.123.123",
		Port:        123,
		User:        "setup-user",
		BecomePass:  "a(*sEtuP1pass_\"5wOrD",
		PrivateKeys: []string{"/from/group/vars", "/from/host/vars"},
		OrderedAt:   "2012-01-01_15:04:05",
	},
	"todo.host": {
		Name:        "todo.host",
		Group:       "setup",
		Host:        "TODO",
		Port:        22,
		User:        "root",
		PrivateKeys: []string{"/from/group/vars"},
	},
}}

func BenchmarkParseInventory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseInventory("testdata/ansible.cfg", "testdata/hosts", "")
	}
}

func TestParseInventory(t *testing.T) {
	expected := testAnsibleInventory

	actual := ParseInventory("testdata/ansible.cfg", "testdata/hosts", "")
	a := actual.Hosts
	for name, host := range expected.Hosts {
		equal(t, host.Name, a[name].Name)
		equal(t, host.Group, a[name].Group)
		equal(t, host.Host, a[name].Host)
		equal(t, host.User, a[name].User)
		equal(t, host.Port, a[name].Port)
		equal(t, host.SSHPass, a[name].SSHPass)
		equal(t, host.BecomePass, a[name].BecomePass)
		for i, key := range host.PrivateKeys {
			equal(t, key, a[name].PrivateKeys[i])
		}
	}

}
