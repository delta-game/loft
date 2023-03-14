package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_ "image/png" // for loading PNG images
)

// Define the dimensions of the banner and showcase images.
var (
	bannerSize   = fyne.NewSize(854, 100)
	showcaseSize = fyne.NewSize(854, 300)
)

// main function
func main() {
	// Create a new Fyne app.
	app := app.New()

	// Set the app's theme to the custom theme.
	app.Settings().SetTheme(&CustomTheme{})

	// Create a new window.
	window := app.NewWindow("Loft")

	///////////////////////////////////////////////////////////////////////////
	/// The Main Content.
	///////////////////////////////////////////////////////////////////////////

	// Create the banner image and resize it to fit the window width.
	bannerImage := canvas.NewImageFromFile("banner.png")
	bannerImage.FillMode = canvas.ImageFillContain
	bannerImage.SetMinSize(bannerSize)

	// Create a showcase image and set its height to 300 pixels.
	showcaseImage := canvas.NewImageFromFile("showcase.png")
	showcaseImage.FillMode = canvas.ImageFillContain
	showcaseImage.SetMinSize(showcaseSize)

	// Create the pages for the tab bar.
	homePage := container.NewVBox(showcaseImage)
	newsPage := container.NewVBox(widget.NewLabel("WIP: News And Updates will be mirrored here."))
	refsPage := container.NewVBox(widget.NewLabel("WIP: The Reference Documentation will mirrored here."))
	setPage := container.NewVBox(widget.NewLabel("WIP: This is where settings will be."))

	// Create the tab bar and add the pages.
	// TODO: We want to figure out how to actually
	// center this. Maybe both it and the bannerImage.
	tabBar := container.NewAppTabs(
		container.NewTabItem("Home", homePage),
		container.NewTabItem("News", newsPage),
		container.NewTabItem("Refs", refsPage),
		container.NewTabItem("Set", setPage),
	)

	// container.NewCenter(tabBar) ?

	// Combine the image container and tab bar in a vertical box layout.
	top := container.NewVBox(bannerImage, tabBar)

	///////////////////////////////////////////////////////////////////////////
	/// The Sidebar UI.
	///////////////////////////////////////////////////////////////////////////

	VersLabel := widget.NewLabelWithStyle("VERSION MANAGER", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: false, Monospace: false})

	// Create the label for the commit header.
	CommitLabel := widget.NewLabel("Commits")

	// Create the buttons for the commit list.
	CommitBtnDown := widget.NewButton("Down", func() {})
	CommitBtnInfo := widget.NewButton("Info", func() {})
	CommitBtnDel := widget.NewButton("Del", func() {})

	// Combine the buttons in a horizontal box layout.
	CommitButts := container.NewHBox(CommitBtnDown, CommitBtnDel, CommitBtnInfo)

	// For Commits, combine the label and buttons into a 'header'.
	CommitHead := container.NewHBox(CommitLabel, CommitButts)

	// Make a list of all available commits for gh: delta-game/delta.
	// TODO: Need to figure out how to actually extend the list.
	// I want it to be like around 1/4th of the sidebar.
	CommitList := widget.NewList(
		func() int { return 7 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(fmt.Sprintf("Commit Hash %d", i))
		},
	)

	// For Commits, combine the header with the list.
	Commits := container.NewVBox(CommitHead, CommitList)
	CommitsContainer := container.NewVScroll(Commits)
	CommitsContainer.SetMinSize(fyne.NewSize(0, 200))

	// Create the header for the instances list.
	InstsLabel := widget.NewLabel("Instances")

	// Create the button for the instances list.
	InstBtnCopy := widget.NewButton("Copy", func() {})
	InstBtnStar := widget.NewButton("Star", func() {})
	InstBtnName := widget.NewButton("Name", func() {})

	// Combine the button in a horizontal box layout.
	InstButts := container.NewHBox(InstBtnStar, InstBtnCopy, InstBtnName)

	// For Instances, combine the label and buttons into a 'header'.
	InstsHead := container.NewHBox(InstsLabel, InstButts)

	// Make a list of all local instances. We'll eventually traverse an install dir.
	// TODO: Also need to extend this list too. Plan for it to take a majority of the
	// free space in the sidebar.
	InstsList := widget.NewList(
		func() int { return 7 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(fmt.Sprintf("Downloaded Instance %d", i))
		},
	)
	// For Instances, combine the header with the list.
	Insts := container.NewVBox(InstsHead, InstsList)
	InstsContainer := container.NewVScroll(Insts)
	InstsContainer.SetMinSize(fyne.NewSize(0, 400))

	// Create a separator for the commits and the insts.
	separator := widget.NewSeparator()

	///////////////////////////////////////////////////////////////////////////
	/// Presentation And Layout.
	///////////////////////////////////////////////////////////////////////////

	// Create a sidebar with the logo and some sample content.
	sidebar := container.NewVBox(
		VersLabel,
		CommitsContainer,
		separator,
		InstsContainer,
		widget.NewButton("LAUNCH", func() {}),
	)

	// Use the border container to position the sidebar on the right side of the window.
	content := container.NewBorder(
		nil,     // top
		nil,     // bottom
		nil,     // left
		sidebar, // right
		top,     // center
	)

	// Set the content of the window to the layout.
	window.SetContent(content)

	// Remove the fixed size limit on the window.
	window.SetFixedSize(false)

	// Resize the window to a larger default size.
	window.Resize(fyne.NewSize(1024, 720))

	// Center the window on the screen.
	window.CenterOnScreen()

	// Refresh the window to avoid potential deadlock with delta/vers creation.
	window.Canvas().Refresh(content)

	// Show the window and run the app.
	window.ShowAndRun()
}
