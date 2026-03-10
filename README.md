# gorTTY

An array sorting screensaver written in go

## Running

```sh
go run .

```

![](https://github.com/suprtrtl/gortty/gallery/gortty.gif)

## TODO

- [x] Implement ArrayGraph interface
- [x] Implement BarGraph type
- [ ] Implement BraileGraph type
- [x] SortingMethod Interface
    - [ ] ~Come up with a way to store changes within an array (preferably without sorting the array beforehand)~
    - [x] Using goroutine setup async array that sends messages to tui
- [ ] Write Bubble Tea Tui Code
    - [x] Foundational program
    - [x] Fullscreen Capabilities
    - [x] Dynamic window Resizing
- [ ] Sorting algorithms
    - [x] Bubble Sort
    - [ ] Quick Sort
    - [x] Merge Sort
    - [x] Selection Sort
    - [x] Bogo Sort
    - [ ] Shell Sort
    - [ ] ...
- [ ] Some way of configuring whether through a file or flags
