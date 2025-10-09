package steamutils

type AppState struct {
	// Basic metadata
	AppID        string `json:"appid"`
	Universe     string `json:"universe"`
	LauncherPath string `json:"LauncherPath"`
	Name         string `json:"name"`
	StateFlags   string `json:"StateFlags"`
	Installdir   string `json:"installdir"`
	LastUpdated  string `json:"LastUpdated"`
	LastPlayed   string `json:"LastPlayed"`
	SizeOnDisk   string `json:"SizeOnDisk"`
	StagingSize  string `json:"StagingSize"`
	Buildid      string `json:"buildid"`
	LastOwner    string `json:"LastOwner"`

	// Update & download information.
	DownloadType                    string `json:"DownloadType,omitempty"` // Only present for some apps.
	UpdateResult                    string `json:"UpdateResult"`
	BytesToDownload                 string `json:"BytesToDownload"`
	BytesDownloaded                 string `json:"BytesDownloaded"`
	BytesToStage                    string `json:"BytesToStage"`
	BytesStaged                     string `json:"BytesStaged"`
	TargetBuildID                   string `json:"TargetBuildID"`
	AutoUpdateBehavior              string `json:"AutoUpdateBehavior"`
	AllowOtherDownloadsWhileRunning string `json:"AllowOtherDownloadsWhileRunning"`
	ScheduledAutoUpdate             string `json:"ScheduledAutoUpdate"`

	// Nested objects.
	InstalledDepots map[string]DepotInfo `json:"InstalledDepots"`
	// SharedDepots is typically a simple map from one depot id to another.
	SharedDepots map[string]string `json:"SharedDepots,omitempty"`
	// InstallScripts maps depot ids to a path string.
	InstallScripts map[string]string `json:"InstallScripts,omitempty"`

	// Configuration sections.
	UserConfig    map[string]string `json:"UserConfig"`
	MountedConfig map[string]string `json:"MountedConfig"`
}

// DepotInfo represents a depot entry within "InstalledDepots".
type DepotInfo struct {
	Manifest string `json:"manifest"`
	Size     string `json:"size"`
	// dlcappid is optional and only appears for some depots.
	DLCAppid string `json:"dlcappid,omitempty"`
}

// The code in this file was made by ChatGPT, use in production is highly discouraged as unexpected results may occur. The code in this file is not vetted for stability or edge cases.
