package main

import (
	"fmt"
	"net/url"

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
	window := app.NewWindow("ΔLOFT")

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	/// The Main Content.
	///////////////////////////////////////////////////////////////////////////
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

	// Not Happy with this naming scheme ...
	linkURLWin, err := url.Parse("https://chocolatey.org/")
	if err != nil {
		panic(err)
	}

	linkChocolatey := widget.NewHyperlink("On Windows, We Need To Install Chocolatey.", linkURLWin)

	linkURLMac, err := url.Parse("https://brew.sh/")
	if err != nil {
		panic(err)
	}
	linkHomebrew := widget.NewHyperlink("On MacOSX, We Need To Install Homebrew.", linkURLMac)

	linkURLLinux, err := url.Parse("https://distrobox.privatedns.org/")
	if err != nil {
		panic(err)
	}
	linkDistrobox := widget.NewHyperlink("On Linux, We Need To Install Distrobox.", linkURLLinux)

	labelWizard := widget.NewLabel("TODO: Have The Wizard Do The Rest!")

	// Create the pages for the tab bar.
	homePage := container.NewVBox(showcaseImage, loremEntry)
	newsPage := container.NewVBox(widget.NewLabel("WIP: News And Updates will be mirrored here."), loremEntry)
	refsPage := container.NewVBox(widget.NewLabel("WIP: The Reference Documentation will mirrored here."), loremEntry)
	setPage := container.NewVBox(widget.NewLabel("SETUP:"), linkChocolatey, linkHomebrew, linkDistrobox, labelWizard, widget.NewLabel("SETTINGS:"), checkboxes)

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

	////////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	/// The Sidebar UI.
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	VersLabel := widget.NewLabelWithStyle("VERSION MANAGER", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	///////////////////////////////////////////////////////////////////////////
	/// The Commits Pane.
	///////////////////////////////////////////////////////////////////////////

	refreshIconResource, _ := fyne.LoadResourceFromPath("resource/fontawesome/rotate-solid.svg")
	CommitBtnRefresh := widget.NewButtonWithIcon("", refreshIconResource, func() {
		// ...
	})

	grabIconResource, _ := fyne.LoadResourceFromPath("resource/fontawesome/download-solid.svg")
	CommitBtnGrab := widget.NewButtonWithIcon("", grabIconResource, func() {
		// ...
	})

	infoIconResource, _ := fyne.LoadResourceFromPath("resource/fontawesome/circle-question-regular.svg")
	CommitBtnInfo := widget.NewButtonWithIcon("", infoIconResource, func() {
		// ...
	})

	// Create the dropdown box for the commit filter.
	CommitFilter := widget.NewSelect([]string{"All Vers", "Upstream", "Stable"}, func(selected string) {
		fmt.Println("Selected commit filter:", selected)
	})
	CommitFilter.SetSelected("All")

	// Combine the buttons in a horizontal box layout.
	CommitButts := container.NewHBox(CommitBtnGrab, CommitBtnInfo)

	CommitLabel := widget.NewLabelWithStyle("Commits", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// For Commits, create the header.
	CommitHead := container.NewHBox(CommitBtnRefresh, CommitButts, CommitFilter, CommitLabel)

	// For Commits, create the list.
	CommitList := widget.NewList(
		func() int { return 21 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(fmt.Sprintf("Commit Hash %d", i))
		},
	)
	CommitListContainer := container.NewVScroll(CommitList)
	CommitListContainer.SetMinSize(fyne.NewSize(0, 140))

	// For Commits, combine the header with the list.
	Commits := container.NewVBox(CommitHead, CommitListContainer)
	CommitsContainer := container.NewVScroll(Commits)
	CommitsContainer.SetMinSize(fyne.NewSize(0, 0))

	///////////////////////////////////////////////////////////////////////////
	/// The Instances Pane.
	///////////////////////////////////////////////////////////////////////////

	upIconResource, _ := fyne.LoadResourceFromPath("resource/fontawesome/chevron-up-solid.svg")
	InstBtnUp := widget.NewButtonWithIcon("", upIconResource, func() {
		// ...
	})

	downIconResource, _ := fyne.LoadResourceFromPath("resource/fontawesome/chevron-down-solid.svg")
	InstBtnDown := widget.NewButtonWithIcon("", downIconResource, func() {
		// ...
	})

	InstArrows := container.NewHBox(InstBtnUp, InstBtnDown)

	copyIconResource, _ := fyne.LoadResourceFromPath("resource/fontawesome/copy-regular.svg")
	InstBtnCopy := widget.NewButtonWithIcon("", copyIconResource, func() {
		// ...
	})

	starIconResource, _ := fyne.LoadResourceFromPath("resource/fontawesome/star-regular.svg")
	InstBtnStar := widget.NewButtonWithIcon("", starIconResource, func() {
		// ...
	})

	nameIconResource, _ := fyne.LoadResourceFromPath("resource/fontawesome/pen-to-square-regular.svg")
	InstBtnName := widget.NewButtonWithIcon("", nameIconResource, func() {
		// ...
	})

	delIconResource, _ := fyne.LoadResourceFromPath("resource/fontawesome/trash-can-regular.svg")
	InstBtnDel := widget.NewButtonWithIcon("", delIconResource, func() {
		// ...
	})

	// Combine the button in a horizontal box layout.
	InstButts := container.NewHBox(InstBtnStar, InstBtnCopy, InstBtnName, InstBtnDel)

	InstLabel := widget.NewLabelWithStyle("Instances", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// For Instances, create the header.
	InstsHead := container.NewHBox(InstArrows, InstButts, InstLabel)

	// For Instances, create the list.
	InstsList := widget.NewList(
		func() int { return 14 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(fmt.Sprintf("Local Instance %d", i))
		},
	)
	InstsListContainer := container.NewVScroll(InstsList)
	InstsListContainer.SetMinSize(fyne.NewSize(0, 380))

	// For Instances, combine the header with the list.
	Insts := container.NewVBox(InstsHead, InstsListContainer)
	InstsContainer := container.NewVScroll(Insts)
	InstsContainer.SetMinSize(fyne.NewSize(0, 0))

	// Create a separator for the commits and the insts.
	Separator := widget.NewSeparator()

	launchIconResource, _ := fyne.LoadResourceFromPath("resource/fontawesome/rocket-solid.svg")
	BtnLaunch := widget.NewButtonWithIcon("", launchIconResource, func() {
		// ...
	})

	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////
	/// Presentation And Layout.
	///////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////

	// Create a sidebar with the logo and some sample content.

	// Combine the sidebar content with separators.
	sidebar := container.NewVBox(
		VersLabel,
		Commits,
		Separator,
		Insts,
		Separator,
		BtnLaunch,
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
