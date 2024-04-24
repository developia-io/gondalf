package main

import (
	"testing"
)

func TestGetBinName(t *testing.T) {
	tests := []struct {
		pm       PackageManager
		expected string
	}{
		{Npm, "npm"},
		{Yarn, "yarn"},
		{Pnpm, "pnpm"},
	}

	for _, test := range tests {
		result := test.pm.GetBinName()
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}

func TestGetInstallParam(t *testing.T) {
	tests := []struct {
		pm       PackageManager
		expected string
	}{
		{Npm, "install"},
		{Yarn, ""},
		{Pnpm, "install"},
	}

	for _, test := range tests {
		result := test.pm.GetInstallParam()
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}

func TestGetAvailablePackageManagers(t *testing.T) {
	pms := GetAvailablePackageManagers()
	for _, pm := range pms {
		t.Logf("Found package manager: %s", pm.GetBinName())
	}
}
