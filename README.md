UHHH
Idk what to put here
Hi, this is a mod installer for my VTOL VR Mods, mostly sound asset replacements and some texture replacements.

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

patch.zip/
    |- assets1.resources
    |- assets2.resources
    |- assets3.resources
    \- patchscript.json 

For self reference on how to set up for compilation:

Modify versioning on: versioninfo.json modinstaller.exe.manifest 

go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo
go generate
go build

signtool sign /tr http://timestamp.digicert.com /td sha256 /fd sha256 /f <certificate> "modinstaller.exe"