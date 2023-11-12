/*
 * SPDX-License-Identifier: GPL-3.0
 * Vencord Installer, a cross platform gui/cli app for installing Vencord
 * Copyright (c) 2023 Vendicated and Vencord contributors
 */

package main

import (
	"fmt"
	"runtime"
)

var IsInstallerOutdated = false

func CheckSelfUpdate() {
	fmt.Println("Checking for Installer Updates...")

	res, err := GetGithubRelease(InstallerReleaseUrl)
	if err == nil {
		IsInstallerOutdated = res.TagName != InstallerTag
	}
}

func GetInstallerDownloadLink() string {
	switch runtime.GOOS {
	case "windows":
		return "https://github.com/Stormcord/Installer/releases/latest/download/StormcordInstaller.exe"
	case "darwin":
		return "https://github.com/Stormcord/Installer/releases/latest/download/StormcordInstaller.MacOS.zip"
	default:
		return ""
	}
}

func GetInstallerDownloadMarkdown() string {
	link := GetInstallerDownloadLink()
	if link == "" {
		return ""
	}
	return " [Download the latest Installer](" + link + ")"
}
