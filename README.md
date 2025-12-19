PatchManager for VTOL VR Mods, mostly sound asset replacements and some texture replacements.

When loading resources, Unity looks in the VTOLVR_Data folder, even if reading a DLC file, so we can point VTOL VR to a resource file 'assets1.resources' that we copied to VTOLVR_Data from anywhere and it'll be able to find and load our custom resources from there.

Currently, only AudioClip and Texture2D is supported for replacement in Asset files and Bundle files. 

DLCs are bundle files, which you can modify with the importbundle action, within this action, you can do a batch of modifications at the same time on a single DLC, you'll have to define these for each DLC.

To replace stuff in the non DLC aircraft, you need to use importasset and point it to VTOL VR's resource file. 

You use Unity to import your custom audios and textures, then compile it to extract your custom resource file that you can then copy into VTOLVR_Data, or in this case, you will save it along the patchscript.json into a ZIP file. You can then use the copy action that the program will use to copy the resource file when patching the game.

Once you have your ZIP file ready, you can click the custom button, and point PatchManager to your zip file (e.g. "C:\Users\bomkz\file.zip") and click on patch to install it.

This implementation is game version agnostic, so long asset names do not change, however, periodically it is expected for the game to break when updating, so you should Verify game files and repatch the game after new VTOL VR updates.

You can contact me in Discord @f45a to get your patch vetted and added to the integrated patch index.

Example patchscript.json implementation:

```
{
    "patchScriptVersion":1,
    "motd": "Null",
    "data": [
        {
            "action": "importbundle",
            "actionData": {
                "originalFilePath": "C:\\Program Files (x86)\\Steam\\steamapps\\common\\VTOL VR\\DLC\\1770480\\1770480",
                "operations": [
                    {
                        "type": "import",
                        "assetType":"AudioClip",
                        "assetName": "ttsw_autopilotOff",
                        "assetPath": "assets1.resources"
                    },
                    {
                        "type": "import",
                        "assetType":"AudioClip",
                        "assetName": "ttsw_pullUp",
                        "assetPath": "assets1.resources"
                    }
                ]
            }
        },
        {
            "action": "importasset",
            "actionData": {
                "originalFilePath": "C:\\Program Files (x86)\\Steam\\steamapps\\common\\VTOL VR\\VTOLVR_Data\\resources.assets",
                "operations": [
                    {
                        "type": "import",
                        "assetType": "AudioClip",
                        "assetName": "ttsw_autopilotOff",
                        "assetPath": "assets2.resources"
                    },
                    {
                        "type": "import",
                        "assetType": "AudioClip",
                        "assetName": "ttsw_pullUp",
                        "assetPath": "assets3.resources"
                    }
                ]
            }
        },
        {
            "action": "copy",
            "actionData": {
                "fileName": "assets1.resources",
                "destination": "assets1.resources"
            }
        },
                {
            "action": "copy",
            "actionData": {
                "fileName": "assets2.resources",
                "destination": "assets2.resources"
            }
        },
                {
            "action": "copy",
            "actionData": {
                "fileName": "assets3.resources",
                "destination": "assets3.resources"
            }
        }
    ]
}
```

Example ZIP File Structure:

```
 patch.zip___
             |- assets1.resources
             |- assets2.resources
             |- assets3.resources
              \- patchscript.json 
```

For self reference on how to set up for compilation:

Modify versioning on: versioninfo.json modinstaller.exe.manifest 

```
go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo
go generate
go build

signtool sign /tr http://timestamp.digicert.com /td sha256 /fd sha256 /f <certificate> "modinstaller.exe"
```