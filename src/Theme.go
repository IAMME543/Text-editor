package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

func InputBorderColor() color.Color {
	return color.White
}

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.White
	case theme.ColorNameInputBackground:
		return color.Transparent
	case theme.ColorNamePrimary:
		return color.Black
	case theme.ColorNameForeground:
		return color.Black
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameInputBorder:
		return 1
	case theme.SizeNameText:
		return 16
	default:
		return theme.DefaultTheme().Size(name)
	}

}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}
