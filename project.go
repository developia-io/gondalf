package main

import (
	"encoding/json"
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

func AddDependency(dir string, dependecy Dependency) error {
	packageJSON, err := os.ReadFile(dir + "/package.json")
	if err != nil {
		return err
	}

	var packageJSONData map[string]interface{}
	err = json.Unmarshal(packageJSON, &packageJSONData)
	if err != nil {
		return err
	}

	dependencies, ok := packageJSONData["dependencies"].(map[string]interface{})
	if !ok {
		dependencies = make(map[string]interface{})
	}

	dependencies[dependecy.Name] = dependecy.Version
	packageJSONData["dependencies"] = dependencies

	packageJSON, err = json.MarshalIndent(packageJSONData, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(dir+"/package.json", packageJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

func CreateProject(recipe Recipe) error {
	err := CloneRepository(recipe.Url, recipe.ProjectName)
	if err != nil {
		return err
	}

	if recipe.UseUILibrary {
		// TODO: Check if the library is already in the dependencies with the correct version
		// TODO: Get UI library name from config
		// TODO: Use the correct version of the library
		err = AddDependency(recipe.ProjectName, Dependency{Name: "@developia.io/developia-ui-lib", Version: "latest"})
		if err != nil {
			return err
		}
	}

	if recipe.InstallDependencies {
		err = InstallDependencies(recipe.PackageManager, recipe.ProjectName)
		if err != nil {
			return err
		}
	}

	return nil
}
