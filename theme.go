package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type CustomTheme struct{}

func (t *CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{0, 0, 0, 255} // Set the background color to black
	case theme.ColorNameForeground:
		return color.RGBA{80, 80, 80, 255} // Set the default text color to dark grey
	case theme.ColorNamePrimary:
		return color.RGBA{163, 35, 24, 255} // Set the primary color to dark red
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (t *CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t *CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

var appTheme = &CustomTheme{}
