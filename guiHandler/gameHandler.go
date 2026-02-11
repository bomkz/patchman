package guihandler

func buildGameListWindowsSteam() {
	sr := global.Assure(steamutils.NewSteamReader(steamutils.SteamReaderConfig{FormatSteamPath: true}))

	global.AssureNoReturn(json.Unmarshal(global.IndexData, &index))

	for _, x := range index {
		gameOptions = append(gameOptions, x.AppName)
	}

	steamPathText := "Steam Path: " + global.SteamPath
	steamPathTextWidget := widget.NewLabel(steamPathText)

	gamePathTextPreset := "Game Path: "
	gamePathText := gamePathTextPreset + "None Selected"
	gamePathTextWidget := widget.NewLabel(gamePathText)

	buildIdTextPreset := "Build ID: "
	buildIdText := buildIdTextPreset + "None Selected"
	buildIdTextWidget := widget.NewLabel(buildIdText)

	gameTextWidget := widget.NewLabel("Select game to modify.")
	gameSelect := widget.NewSelect(gameOptions, func(s string) {
		for _, x := range index {
			if x.AppName == s {
				gamePath := global.Assure(sr.FindAppIDPath(x.AppID))
				gamePathText = gamePathTextPreset + gamePath
				gamePathTextWidget.SetText(gamePathText)

				buildId := global.Assure(sr.FindAppIDBuildID(x.AppID))
				buildIdText = buildIdTextPreset + buildId
				buildIdTextWidget.SetText(buildIdText)
			}
		}
	})

	global.MainWindow.SetContent(container.NewVBox(
		steamPathTextWidget,
		gamePathTextWidget,
		buildIdTextWidget,
		gameTextWidget,
		gameSelect,
		container.NewHBox(
			widget.NewButton("Next", func() { os.Exit(0) }),
			widget.NewButton("Cancel", func() { os.Exit(0) }),
		),
	))
}
