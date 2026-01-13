# PatchManager 

Is a utility to quickly replace sound and texture assets in VTOL VR with the goal of being simple to use both as a user and a patch maker.

## Creating your own patch

PatchManager uses its own JSON-based patch"script", where you define the actions needed to patch the game once, with the ability to work across future versions of the game provided the asset names have not changed.

To start, you should have your own assets compiled in unity, as the resource file unity produces will be needed.

### How it works

When loading resources, Unity looks in the Data folder, even if reading a DLC file, so we can point Unity to a resource file 'example1.resource' that we copied to Data from anywhere and it'll be able to find and load our custom resources from there.

This is where your resource file comes in handy. In your patchscript.json file, you will need to tell PatchMan the name of your resource file, as well as where to copy it, the name can be anything, however, unless you have a niche situation, you should copy your resource file to your Data folder.

Do note, you only need to tell PatchMan where the assets and bundle files in relation to the root of the game path, PatchMan automatically find where a user has the Unity game installed and fills the rest of the path, assuming its a steam game.

This is valid: `VTOLVR_Data\resources.assets`

This is not valid: `C:\Program Files(x86)\steam\steamapps\common\VTOL VR\VTOLVR_Data\resources.assets`


### Defining patches

A valid patchscript file needs patchScriptVersion, motd, and data. 
Patchscript is the version of patchscript your patch is written in, mainly to maintain backwards compatibility with future releases. Currently only patchScriptVersion 1 is valid. Data contains the actions needed to patch the game.

```
{
    "patchScriptVersion":1,
    "data": []
}
```

### Patching Assets

Patchscript currently only has two valid patch types: `AudioClip` and `Texture2D`.
To define a patch you need to know the following: the name of the asset, the asset type (found above), and the resource asset file that contains it. You also need the size, length (for AudioClips), and offset, you can find this in Unity Asset Bundle Extractor. The following is an example of this:

![Example 1](example/1.png?raw=true)
![Example 2](example/2.png?raw=true)
![Example 3](example/3.png?raw=true)

```
        {
            "action": "importasset",
            "actionData": {
                "originalFilePath": "VTOLVR_Data\\resources.assets",
                "operations": [
                    {
                        "type": "import",
                        "assetType":"AudioClip",
                        "assetName": "radarLocked",
                        "assetPath": "example1.resource",
                        "offset": 1600,
                        "size": 864,
                        "length": 0.189375
                    },
                    {
                        "type": "import",
                        "assetType":"AudioClip",
                        "assetName": "ttsw_pullUp",
                        "assetPath": "example2.resource",
                        "offset": 0,
                        "size": 31104,
                        "length": 4.105465 
                    },
                                        {
                        "type": "import",
                        "assetType":"Texture2D",
                        "assetName": "rwrTexture",
                        "assetPath": "example3.assets.resS",
                        "offset": 691040,
                        "size": 691040
                    }
                ]
            }
        },
        
```


### Patching bundles

Due to DLCs being packaged into compressed bundle files, you need to modify them with a separate action called `importbundle`, within this action, you can do a batch of modifications at the same time on a single DLC as you would with a normal asset file like resources.assets.

```
        {
            "action": "importbundle",
            "actionData": {
                "originalFilePath": "DLC\\1770480\\1770480",
                "operations": [
                    {
                        "type": "import",
                        "assetType":"AudioClip",
                        "assetName": "radarLocked",
                        "assetPath": "example1.resource",
                        "offset": 1600,
                        "size": 864,
                        "length": 0.189375
                    },
                    {
                        "type": "import",
                        "assetType":"AudioClip",
                        "assetName": "ttsw_pullUp",
                        "assetPath": "example2.resource",
                        "offset": 0,
                        "size": 31104,
                        "length": 4.105465 
                    },
                                        {
                        "type": "import",
                        "assetType":"Texture2D",
                        "assetName": "rwrTexture",
                        "assetPath": "example3.assets.resS",
                        "offset": 691040,
                        "size": 691040
                    }
                ]
            }
        },
        
```

### Copying resource files

Since resource files need to be copied to VTOLVR_Data folder, we need to define a copy action to do this for us, for each resource file we plan to use, we need to define a copy action so it resembles as follows:

```
        {
            "action": "copy",
            "actionData": {
                "fileName": "assets1.resources",
                "destination": "VTOLVR_Data\\example1.resource"
            }
        },
```

### Packaging your patch

Once you have created your own patchscript.json, you can then package it and test it. 
To package it properly, you need to put it in a ZIP file along with your custom resource files so that the ZIP file structure resembles the following (Note, may have less or more resources/resS files depending on your patch):

![Example 4](example/4.png?raw=true)


```
 patch.zip___
             |- example1.resource
             |- example2.resource
             |- example3.assets.resS
              \- patchscript.json 
```

### Installing your patch

Once you are ready to test your patch, you can open PatchMan, and click on the custom button in the lower area, there, you will input the path to your zip file, as an example: `C:\Users\bomkz\Documents\example.zip`, and then click on install. Assuming your patchscript is valid, it should finish without error.

PatchManager is game version agnostic, so long asset names do not change, however, periodically it is expected for the game to break when updating, so you should Verify game files and repatch the game after new updates.

### Mix and matching patches

You can mix and match different patches by disabling certain content and/or assets, installing the patch, and then rerunning patchman, and excluding everything except the stuff you excluded last time, and then installing another patch. This can be done however many times and combinations. You can also save and load asset/content presets to apply before installing a patch.

## Example patchscript.json implementation:

```
{
    "patchScriptVersion":1,
    "data": [
        {
            "action": "importbundle",
            "actionData": {
                "originalFilePath": "DLC\\1770480\\1770480",
                "operations": [
                    {
                        "type": "import",
                        "assetType":"AudioClip",
                        "assetName": "radarLocked",
                        "assetPath": "example1.resource",
                        "offset": 1600,
                        "size": 864,
                        "length": 0.189375
                    },
                    {
                        "type": "import",
                        "assetType":"AudioClip",
                        "assetName": "ttsw_pullUp",
                        "assetPath": "example2.resource",
                        "offset": 0,
                        "size": 31104,
                        "length": 4.105465 
                    },
                                        {
                        "type": "import",
                        "assetType":"Texture2D",
                        "assetName": "rwrTexture",
                        "assetPath": "example3.assets.resS",
                        "offset": 691040,
                        "size": 691040
                    }
                ]
            }
        },
        {
            "action": "importbundle",
            "actionData": {
                "originalFilePath": "DLC\\2531290\\2531290",
                "operations": [
                    {
                        "type": "import",
                        "assetType":"AudioClip",
                        "assetName": "radarLocked",
                        "assetPath": "example1.resource",
                        "offset": 1600,
                        "size": 864,
                        "length": 0.189375
                    },
                    {
                        "type": "import",
                        "assetType":"AudioClip",
                        "assetName": "ttsw_pullUp",
                        "assetPath": "example2.resource",
                        "offset": 0,
                        "size": 31104,
                        "length": 4.105465 
                    },
                                        {
                        "type": "import",
                        "assetType":"Texture2D",
                        "assetName": "rwrTexture",
                        "assetPath": "example3.assets.resS",
                        "offset": 691040,
                        "size": 691040
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
                        "assetType":"AudioClip",
                        "assetName": "radarLocked",
                        "assetPath": "example1.resource",
                        "offset": 1600,
                        "size": 864,
                        "length": 0.189375
                    },
                    {
                        "type": "import",
                        "assetType":"AudioClip",
                        "assetName": "ttsw_pullUp",
                        "assetPath": "example2.resource",
                        "offset": 0,
                        "size": 31104,
                        "length": 4.105465 
                    },
                                        {
                        "type": "import",
                        "assetType":"Texture2D",
                        "assetName": "rwrTexture",
                        "assetPath": "example3.assets.resS",
                        "offset": 691040,
                        "size": 691040
                    }
                ]
            }
        },
        {
            "action": "copy",
            "actionData": {
                "fileName": "assets1.resources",
                "destination": "VTOLVR_Data\\example1.resources"
            }
        },
                {
            "action": "copy",
            "actionData": {
                "fileName": "assets2.resources",
                "destination": "VTOLVR_Data\\example2.resources"
            }
        },
                {
            "action": "copy",
            "actionData": {
                "fileName": "assets3.assets.resS",
                "destination": "VTOLVR_Data\\example3.assets.resS"
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
go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo
go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo
go generate
go build

signtool sign /tr http://timestamp.digicert.com /td sha256 /fd sha256 /f <certificate> "modinstaller.exe"
```
