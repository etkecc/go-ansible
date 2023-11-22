package ansible

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	yaml "gopkg.in/yaml.v3"
)

func TestNewHostVarsFile(t *testing.T) {
	expected := readVarsYml(t, "testdata/host_vars/setup.host/vars.yml")
	actual, err := NewHostVarsFile("testdata/host_vars/setup.host/vars.yml")
	if err != nil {
		t.Error("file parsing failed", err)
	}
	if actual == nil {
		t.Error("parsed vars.yml is nil")
	}
	t.Log("actual:")
	for k, v := range actual {
		t.Log(k, "=", v)
	}
	if len(expected) != len(actual) {
		t.Error("length is not expected. Expected:", len(expected), "actual:", len(actual))
	}
	for k, v := range expected {
		equalRecursive(t, v, actual[k])
	}
}

func TestNewHostVars_NoExists(t *testing.T) {
	_, err := NewHostVarsFile("testdata/not-exists")

	if err == nil {
		t.Error("error is nil")
	}
	if !errors.Is(err, os.ErrNotExist) {
		t.Error("error is not os.ErrNotExists, it is", err)
	}
}

func TestHostVarsHasTODOs(t *testing.T) {
	tests := map[bool]string{
		true:  "ToDo",
		false: "any value",
	}
	for expected, input := range tests {
		t.Run(input, func(t *testing.T) {
			hv := HostVars{"key": input}

			actual := hv.HasTODOs()

			if expected != actual {
				t.Error("value is invalid. Expected: '" + fmt.Sprintf("%t", expected) + "', actual: '" + fmt.Sprintf("%t", actual) + "'")
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := map[string]any{
		"valid": "valid",
		"":      false,
	}
	for expected, input := range tests {
		t.Run(expected, func(t *testing.T) {
			hv := HostVars{"key": input}

			actual := hv.String("key")

			if expected != actual {
				t.Error("value is invalid. Expected: '" + expected + "', actual: '" + actual + "'")
			}
		})
	}
}

func TestStringSlice(t *testing.T) {
	tests := map[string]any{
		"valid,valid": []any{"valid", "valid"},
		"":            "invalid",
	}
	for expected, input := range tests {
		t.Run(expected, func(t *testing.T) {
			hv := HostVars{"key": input}

			actual := hv.StringSlice("key")

			actualStr := strings.Join(actual, ",")
			if expected != actualStr {
				t.Error("value is invalid. Expected: '" + expected + "', actual: '" + actualStr + "'")
			}
		})
	}
}

func TestMaintenanceEnabled(t *testing.T) {
	tests := map[bool]string{
		true:  "any value",
		false: "false",
	}
	for expected, input := range tests {
		t.Run(input, func(t *testing.T) {
			hv := HostVars{"etke_service_maintenance_enabled": input}

			actual := hv.MaintenanceEnabled()

			if expected != actual {
				t.Error("value is invalid. Expected: '" + fmt.Sprintf("%t", expected) + "', actual: '" + fmt.Sprintf("%t", actual) + "'")
			}
		})
	}
}

func readVarsYml(t *testing.T, filepath string) map[string]any {
	t.Helper()
	input, err := os.ReadFile(filepath)
	if err != nil {
		t.Error(err)
	}

	var vars map[string]any
	if err := yaml.Unmarshal(input, &vars); err != nil {
		t.Error(err)
	}
	return vars
}

func equalRecursive(t *testing.T, expected, actual any) {
	t.Helper()

	mv, ok := expected.(map[string]any)
	if ok {
		av := actual.(map[string]any) //nolint:forcetypeassert // that's ok
		for k, v := range mv {
			equalRecursive(t, v, av[k])
		}
		return
	}

	sv, ok := expected.([]any)
	if ok {
		av := actual.([]any) //nolint:forcetypeassert // that's ok
		for i, v := range sv {
			equalRecursive(t, v, av[i])
		}
		return
	}

	if expected != actual {
		t.Error("value is invalid. Expected: '" + fmt.Sprintf("%v", expected) + "', actual: '" + fmt.Sprintf("%v", actual) + "'")
	}
}
