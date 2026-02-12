package global

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/bomkz/steamutils"
)

func GetSteamPath() string {
	sr := Assure(steamutils.NewSteamReader(steamutils.SteamReaderConfig{}))
	steamPath := sr.GetSteamPath()
	return steamPath
}
func GetGameByAppId(appID string) {
	//Section 1
	sr, err := steamutils.NewSteamReader(steamutils.SteamReaderConfig{})
	if err != nil {
		FatalErrorPriviledged(err, "Exception encountered at at: GetGameByAppId:S1")
	}

	//Section 2
	game, err := sr.GetInstalledAppByID(appID)
	if err != nil {
		FatalErrorPriviledged(err, "Exception encountered at at: GetGameByAppId:S2")
	}

	//Section 3
	gameByte, err := json.Marshal(game)
	if err != nil {
		FatalErrorPriviledged(err, "Exception encountered at at: GetGameByAppId:S3")
	}

	fmt.Println(gameByte)
}

func FatalErrorPriviledged(err error, simplified string) {

	logfile, err1 := os.OpenFile("./patchman.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		fmt.Println("Fatal Error: Helper could not create or write to log file, simplified error: " + simplified)
		os.Exit(1)
	}
	timestamp := time.Now().Truncate(time.Second).String()

	if _, err1 := logfile.WriteString(timestamp + ": " + err.Error() + "\n"); err1 != nil {
		fmt.Println("Fatal Error: Helper could not create or write to log file, simplified error: " + simplified)
	}

	fmt.Println("Fatal Error: Helper encountered a problem, error dumped into ./patchman.log. Simplified Error: " + simplified)

	logfile.Close()

	os.Exit(1)

}
