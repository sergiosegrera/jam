# Jam
![Jam logo](./assets/icon128.png)   
A basic ebiten boilerplate for future game jam use.  
Inspired by pico-8.

## Stack
* Aseprite for sprites
* Audacity for SFX recording
* Use JSON for tilemaps
* `/audio/music.go` for mp3 infinite bg music
* `/audio/sound.go` for wav sfx
* `/utils/spritesheet.go` for spritesheet "slicing"

## Installation
```
$ git clone https://github.com/sergiosegrera/jam {Game Name}
$ cd {Game Name}
$ ./jam init
```

## Usage
* `$ ./jam init` -- Creates new go.mod and changes imports using gomove  
* `$ ./jam run` -- Builds and runs binary  
* `$ ./jam build` -- Builds binary  
* `$ ./jam windows` -- Creates windows release  
* `$ ./jam todo` -- Shows all "TODO:" comments  
* `$ ./jam clean` -- Cleans .exe, .o, .zip files  

## TODO
* Particle example
* Asset pre-loading example
* GameObject slice example?
* GameObject interaction
* `audio/music.go:21` Fix hard coded music length 
