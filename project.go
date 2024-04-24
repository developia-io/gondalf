package main

import (
	"os"

	"github.com/go-git/go-git/v5"
)

func CloneRepository(url string, dest string) error {
	_, err := git.PlainClone(dest, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return err
	}

	return nil
}

func InstallDependencies(pm PackageManager, dest string) error {
	err := os.Chdir(dest)
	if err != nil {
		return err
	}

	return pm.Install()
}

func CreateProject(recipe Recipe) error {
	err := CloneRepository(recipe.Url, recipe.ProjectName)
	if err != nil {
		return err
	}

	if recipe.Install {
		err = InstallDependencies(recipe.PackageManager, recipe.ProjectName)
		if err != nil {
			return err
		}
	}

	return nil
}
