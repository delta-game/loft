package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	// for embedding static resources

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
	app.Settings().SetTheme(NewCustomTheme())

	// Create a new window.
	window := app.NewWindow("Loft")

	///////////////////////////////////////////////////////////////////////////
	/// The Main Content.
	///////////////////////////////////////////////////////////////////////////

	// Create the banner image and resize it to fit the window width.
	bannerImage := canvas.NewImageFromResource(resourceBannerPng)
	bannerImage.FillMode = canvas.ImageFillContain
	bannerImage.SetMinSize(bannerSize)

	// Create a showcase image and set its height to 300 pixels.
	showcaseImage := canvas.NewImageFromResource(resourceShowcasePng)
	showcaseImage.FillMode = canvas.ImageFillContain
	showcaseImage.SetMinSize(showcaseSize)

	// Create an entry widget with a wrapped text paragraph.
	loremEntry := widget.NewMultiLineEntry()
	loremEntry.Wrapping = fyne.TextWrapWord
	loremEntry.Text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus tempor velit vitae velit eleifend, at venenatis massa dictum. Pellentesque ut sapien magna. Nullam finibus, sem eu bibendum imperdiet, urna turpis placerat nulla, nec malesuada felis arcu ut enim. Integer malesuada massa vel ex efficitur gravida. Vivamus posuere mauris at ante rhoncus, vitae iaculis magna elementum. Nulla pulvinar mi vel urna porttitor, vel varius libero bibendum. Integer nec metus posuere, consectetur massa id, placerat ipsum. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Aliquam sed elit eget lectus pharetra venenatis. In vel tempor nulla, ut viverra mauris. Sed nec congue lacus, nec tempor arcu. Sed tincidunt tortor libero, vel ultrices orci aliquam at.\n\nSuspendisse potenti. Donec convallis, quam et pulvinar fermentum, quam libero accumsan nibh, nec elementum ex quam sit amet nulla. Sed fringilla imperdiet odio, vel ultricies magna tristique in. Sed euismod libero quis libero laoreet interdum. Fusce vestibulum bibendum sem, non ullamcorper risus molestie non. Nunc id enim sit amet lectus elementum luctus. Vivamus commodo turpis eget mauris semper, nec sagittis ipsum lacinia.\n\nProin sit amet turpis vitae sapien suscipit cursus eu vel est. Praesent venenatis, nibh vel dictum semper, turpis turpis elementum libero, vel consequat arcu neque ut massa. Ut ut ante ligula. Suspendisse eleifend mi ipsum, vel commodo mi vestibulum id. Aenean ut sapien et velit suscipit dapibus vitae in purus. Integer ut felis in elit ultricies tincidunt. Curabitur luctus bibendum tellus, nec consectetur tortor ullamcorper at. Aliquam erat volutpat. Quisque volutpat ex quis eros convallis ullamcorper. Sed euismod mi ut ex venenatis, quis feugiat justo faucibus. Sed bibendum eros ac massa fermentum interdum."

	// Add some checkboxes to fill in the settings page.
	checkbox1 := widget.NewCheck("Option 1", func(checked bool) {
	})

	checkbox2 := widget.NewCheck("Option 2", func(checked bool) {
	})

	checkbox3 := widget.NewCheck("Option 3", func(checked bool) {
	})

	checkboxes := container.NewVBox(
		checkbox1,
		checkbox2,
		checkbox3,
	)

	// Create the pages for the tab bar.
	homePage := container.NewVBox(showcaseImage, loremEntry)
	newsPage := container.NewVBox(widget.NewLabel("WIP: News And Updates will be mirrored here."), loremEntry)
	refsPage := container.NewVBox(widget.NewLabel("WIP: The Reference Documentation will mirrored here."), loremEntry)
	setPage := container.NewVBox(widget.NewLabel("WIP: This is where settings will be."), checkboxes)

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
	CommitBtnRefresh := widget.NewButton("Refresh", func() {})
	CommitBtnGrab := widget.NewButton("Grab", func() {})
	CommitBtnInfo := widget.NewButton("Info", func() {})

	// Create the dropdown box for the commit filter.
	CommitFilter := widget.NewSelect([]string{"All Vers", "Upstream", "Stable"}, func(selected string) {
		fmt.Println("Selected commit filter:", selected)
	})
	CommitFilter.SetSelected("All")

	// Combine the buttons in a horizontal box layout.
	CommitButts := container.NewHBox(CommitBtnGrab, CommitBtnInfo)

	// For Commits, combine the label and buttons into a 'header'.
	CommitHead := container.NewHBox(CommitBtnRefresh, CommitLabel, CommitButts, CommitFilter)

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
	CommitsContainer.SetMinSize(fyne.NewSize(0, 210))

	InstBtnUp := widget.NewButton("Ʌ", func() {})
	InstBtnDown := widget.NewButton("V", func() {})

	InstArrows := container.NewHBox(InstBtnUp, InstBtnDown)

	// Create the header for the instances list.
	InstsLabel := widget.NewLabel("Instances")

	// Create the button for the instances list.
	InstBtnCopy := widget.NewButton("Copy", func() {})
	InstBtnStar := widget.NewButton("Star", func() {})
	InstBtnName := widget.NewButton("Name", func() {})
	InstBtnDel := widget.NewButton("Del", func() {})

	// Combine the button in a horizontal box layout.
	InstButts := container.NewHBox(InstBtnStar, InstBtnCopy, InstBtnName, InstBtnDel)

	// For Instances, combine the label and buttons into a 'header'.
	InstsHead := container.NewHBox(InstArrows, InstsLabel, InstButts)

	// Make a list of all local instances. We'll eventually traverse an install dir.
	// TODO: Also need to extend this list too. Plan for it to take a majority of the
	// free space in the sidebar.
	InstsList := widget.NewList(
		func() int { return 7 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(fmt.Sprintf("Grabbed Instance %d", i))
		},
	)
	// For Instances, combine the header with the list.
	Insts := container.NewVBox(InstsHead, InstsList)
	InstsContainer := container.NewVScroll(Insts)
	InstsContainer.SetMinSize(fyne.NewSize(0, 420))

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
