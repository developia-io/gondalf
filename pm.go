package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh"
)

type PackageManager int

const (
	Npm PackageManager = iota
	Yarn
	Pnpm
)

func (pm PackageManager) GetBinName() string {
	switch pm {
	case Npm:
		return "npm"
	case Yarn:
		return "yarn"
	case Pnpm:
		return "pnpm"
	default:
		return ""
	}
}

func (pm PackageManager) GetInstallParam() string {
	switch pm {
	case Npm:
		return "install"
	case Yarn:
		return ""
	case Pnpm:
		return "install"
	default:
		return ""
	}
}

func (pm PackageManager) Install() error {
	cmd := exec.Command(pm.GetBinName(), pm.GetInstallParam())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func PackageManagerOptions(pms []PackageManager) []huh.Option[PackageManager] {
	options := make([]huh.Option[PackageManager], len(pms))
	for i, pm := range pms {
		options[i] = huh.NewOption(pm.GetBinName(), pm)
	}
	return options
}

func GetAvailablePackageManagers() []PackageManager {
	var availablePackageManagers []PackageManager

	for _, manager := range []PackageManager{Npm, Yarn, Pnpm} {
		cmd := exec.Command(manager.GetBinName(), "--version")
		output, err := cmd.Output()
		if err == nil && cmd.ProcessState.Success() {
			availablePackageManagers = append(availablePackageManagers, manager)
			fmt.Println(string(output))
		}
	}

	return availablePackageManagers
}
