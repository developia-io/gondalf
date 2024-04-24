package main

import "github.com/charmbracelet/huh"

func RunForm(recipe *Recipe) error {
	pms := GetAvailablePackageManagers()
	templateConf, err := LoadTemplates("")
	if err != nil {
		return err
	}
	templateOptions := templateConf.Options()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().Title("gondalf").Description("Create a new project"),
			huh.NewInput().Value(&recipe.ProjectName).Title("Project Name").Placeholder("Enter project name here"),
		),
		huh.NewGroup(
			huh.NewSelect[string]().Value(&recipe.Url).Title("Template").Description("Choose a template").Options(
				templateOptions...,
			),
		),
		huh.NewGroup(
			huh.NewConfirm().Title("Install").Description("Would you like to install the project dependencies?").Value(&recipe.Install),
		).WithHideFunc(func() bool {
			return len(pms) <= 0
		},
		),
		huh.NewGroup(
			huh.NewSelect[PackageManager]().Value(&recipe.PackageManager).Title("Package Manager").Description("Choose a package manager").Options(
				PackageManagerOptions(pms)...,
			),
		).WithHideFunc(func() bool {
			return !recipe.Install
		},
		),
	)

	return form.Run()
}
