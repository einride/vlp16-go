//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"

	// mage:import
	"go.einride.tech/mage-tools/targets/mgconvco"

	// mage:import
	"go.einride.tech/mage-tools/targets/mggo"

	// mage:import
	"go.einride.tech/mage-tools/targets/mggolangcilint"

	// mage:import
	"go.einride.tech/mage-tools/targets/mgmarkdownfmt"

	// mage:import
	"go.einride.tech/mage-tools/targets/mggitverifynodiff"
)

func All() {
	mg.Deps(
		mg.F(mgconvco.ConvcoCheck, "origin/master..HEAD"),
		mggolangcilint.GolangciLint,
		mggo.GoTest,
		mgmarkdownfmt.FormatMarkdown,
	)
	mg.SerialDeps(
		mggo.GoModTidy,
		mggitverifynodiff.GitVerifyNoDiff,
	)
}
