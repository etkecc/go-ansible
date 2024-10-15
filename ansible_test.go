package ansible

import "testing"

func BenchmarkParseInventory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseInventory("testdata/ansible.cfg", "testdata/hosts", "")
	}
}

func TestParseInventory(t *testing.T) {
	expected := testInventory

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
		equal(t, host.PrivateKey, a[name].PrivateKey)
	}

}
