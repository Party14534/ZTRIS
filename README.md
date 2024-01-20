# ZTRIS

### About ZTRIS
ZTRIS is a TUI tetris clone created in GO using the [tview](https://github.com/rivo/tview) API to render the TUI 

## How to run ZTRIS
1. Clone this repo
2. Ensure you have the most recent version of go installed
3. In your terminal move into ZTRIS' directory
4. Use the go get command to install the required dependencies 
```console
go get
```
5. Build the project using the go build command
```console
go build ztris.go
```
6. Run the program from the project's directory
```console
./ztris
```
7. Have fun!

## Controls
- Move the tetronimo left and right using the arrow keys [<-, ->]
- Drop tetronimo using the Spacebar
- Rotate the tetronimo left using the z key

## TODO
- Add a player score to track progress
- Add a level system to make the blocks fall faster every level
- Add a fastfall system to allow blocks to fall if the user hold the down arrow
