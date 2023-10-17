package ansible

import (
	"errors"
	"os"
	"testing"
)

func TestNewAnsibleCfgFile(t *testing.T) {
	expected := &AnsibleCfg{Config: map[string]map[string]string{
		"defaults": {
			"callback_plugins":        "plugins/callback",
			"display_skipped_hosts":   "False",
			"fact_caching":            "jsonfile",
			"fact_caching_connection": "/tmp/ansible",
			"forks":                   "50",
			"gathering":               "smart",
			"host_key_checking":       "False",
			"inventory":               "inventory/hosts,hosts,../inventory/hosts",
			"log_path":                "/tmp/ansible/play.log",
			"remote_user":             "root",
			"retry_files_enabled":     "False",
			"roles_path":              "roles:upstream/roles",
			"strategy":                "linear",
			"timeout":                 "86400",
		},
		"ssh_connection": {
			"control_path": "/tmp/ansible/ssh-%%h-%%p-%%r",
			"pipelining":   "True",
		},
	}}

	actual, err := NewAnsibleCfgFile("testdata/ansible.cfg")
	if err != nil {
		t.Error("file parsing failed", err)
	}
	if actual == nil { //nolint:staticcheck // yes, that's the point
		t.Error("parsed ansible.cfg is nil")
	}
	t.Log("actual:")
	for k, v := range actual.Config { //nolint:staticcheck // that's intended
		t.Log(k, "=", v)
	}
	if len(expected.Config) != len(actual.Config) { //nolint:staticcheck // that's intended
		t.Error("length is not expected. Expected:", len(expected.Config), "actual:", len(actual.Config))
	}
	for sc := range expected.Config {
		if len(expected.Config[sc]) != len(actual.Config[sc]) {
			t.Error(sc, "section length is not expected. Expected:", len(expected.Config[sc]), "actual:", len(actual.Config[sc]))
		}
		for k, v := range expected.Config[sc] {
			if v != actual.Config[sc][k] {
				t.Error(sc+"."+k, "is invalid. Expected: '"+sc+"."+k+"="+v+"', actual: '"+sc+"."+k+"="+v+"'")
			}
		}
	}
}

func TestNewAnsibleCfgFile_Empty(t *testing.T) {
	actual, err := NewAnsibleCfgFile("")
	if err != nil {
		t.Error("file parsing failed", err)
	}
	if actual != nil {
		t.Error("parsed ansible.cfg is nil")
	}
}

func TestNewAnsibleCfgFile_NoExists(t *testing.T) {
	_, err := NewAnsibleCfgFile("testdata/not-exists")

	if err == nil {
		t.Error("error is nil")
	}
	if !errors.Is(err, os.ErrNotExist) {
		t.Error("error is not os.ErrNotExists, it is", err)
	}
}
