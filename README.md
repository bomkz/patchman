# PatchManager 

Is a utility to quickly replace sound and texture assets in VTOL VR with the goal of being simple to use both as a user and a patch maker.

## Creating your own patch

PatchManager uses its own JSON-based patch"script", where you define the actions needed to patch the game once, with the ability to work across future versions of the game provided the asset names have not changed.

To start, you should have your own assets compiled in unity, as the resource file unity produces will be needed.

### How it works

When loading resources, Unity looks in the VTOLVR_Data folder, even if reading a DLC file, so we can point VTOL VR to a resource file 'assets1.resources' that we copied to VTOLVR_Data from anywhere and it'll be able to find and load our custom resources from there.

This is where your resource file comes in handy. In your patchscript.json file, you will need to tell PatchMan the name of your resource file, as well as where to copy it, the name can be anything, however, unless you have a niche situation, you should copy your resource file to your VTOLVR_Data folder.

Do note, you only need to tell PatchMan where the assets and bundle files in relation to the root of VTOL VR path, PatchMan automatically find where a user has VTOL VR installed and fills the rest of the path.

This is valid: `VTOLVR_Data\resources.assets`

This is not valid: `C:\Program Files(x86)\steam\steamapps\common\VTOL VR\VTOLVR_Data\resources.assets`


### Defining patches

A valid patchscript file needs patchScriptVersion, motd, and data. 
Patchscript is the version of patchscript your patch is written in, mainly to maintain backwards compatibility with future releases. Currently only patchScriptVersion 1 is valid. motd is a message or description you may want to display when installing your patch, and data contains the actions needed to patch the game.

```
{
    "patchScriptVersion":1,
    "motd": "Example Message",
    "data": []
}
```

### Patching Assets

Patchscript currently only has two valid patch types: `AudioClip` and `Texture2D`.
To define a patch you need to know the following: the name of the asset, the asset type (found above), and the resource asset file that contains it. The following is an example of this:

```
{
    "action": "importasset",
    "actionData": {
        "originalFilePath": "VTOLVR_Data\\resources.assets",
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
}
        
```


### Patching bundles

Due to DLCs being packaged into compressed bundle files, you need to modify them with a separate action called `importbundle`, within this action, you can do a batch of modifications at the same time on a single DLC as you would with a normal asset file like resources.assets.

```
{
    "action": "importbundle",
    "actionData": {
        "originalFilePath": "DLC\\1770480\1770480",
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
}
        
```

### Copying resource files

Since resource files need to be copied to VTOLVR_Data folder, we need to define a copy action to do this for us, for each resource file we plan to use, we need to define a copy action so it resembles as follows:

```
{
    "action": "copy",
    "actionData": {
        "fileName": "examplemodified.resources",
        "destination": "VTOLVR_Data\\examplemodified.resources"
    }
}
```

### Packaging your patch

Once you have created your own patchscript.json, you can then package it and test it. 
To package it properly, you need to put it in a ZIP file along with your custom resource files so that the ZIP file structure resembles the following:

```
 patch.zip___
             |- assets1.resources
             |- assets2.resources
             |- assets3.resources
              \- patchscript.json 
```

### Installing your patch

Once you are ready to test your patch, you can open PatchMan, and click on the custom button in the lower area, there, you will input the path to your zip file, as an example: `C:\Users\bomkz\Documents\example.zip`, and then click on install. Assuming your patchscript is valid, it should finish without error.

PatchManager is game version agnostic, so long asset names do not change, however, periodically it is expected for the game to break when updating, so you should Verify game files and repatch the game after new VTOL VR updates.

## Example patchscript.json implementation:

```
{
    "patchScriptVersion":1,
    "motd": "Example Message",
    "data": [
        {
            "action": "importbundle",
            "actionData": {
                "originalFilePath": "DLC\\1770480\\1770480",
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
                "originalFilePath": "VTOLVR_Data\\resources.assets",
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
                "destination": "VTOLVR_Data\\assets1.resources"
            }
        },
        {
            "action": "copy",
            "actionData": {
                "fileName": "assets2.resources",
                "destination": "VTOLVR_Data\\assets2.resources"
            }
        },
        {
            "action": "copy",
            "actionData": {
                "fileName": "assets3.resources",
                "destination": "VTOLVR_Data\\assets3.resources"
            }
        }
    ]
}
```

## Adding your patch to the index

You can contact me on Discord @f45a to vet your patch and add it into the index.

## Compiling 

Modify information in: versioninfo.json && modinstaller.exe.manifest 

```

go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo
go generate
go build

signtool sign /tr http://timestamp.digicert.com /td sha256 /fd sha256 /f <certificate> "modinstaller.exe"
```