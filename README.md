# Tetris in Golang

Thank you to L-Gad for the Tetrimino board, mino and tetrimino [sprites](https://l-gad.itch.io/tetriminos-asset-pack)

This is a personal project that aimed to recreate Tetris in Go. I used the [Ebitengine API](https://ebitengine.org/) for simple 2D graphics and rendering. All code is original.

The aim of the project **practice and refresh myself on OO principles mainly on classes and interfaces**. I did not find a suitable use for Go's concurrency.

## Features
- Holding and Tetrimino queue using the [7-bag randomiser](https://harddrop.com/wiki/Random_Generator) system
- [Lock Delay](https://harddrop.com/wiki/Lock_delay) to allow for spins and slides
- [Delay Auto Shift](https://tetris.fandom.com/wiki/DAS) (DAS) for left and right movement
- Kicks and Spins using the [Super Rotation System](https://harddrop.com/wiki/SRS) guidelines

## Things that I did not end up implementing
- Sound effects from Tetr.io
- Menu and settings for hotkey reassignment
- Text labels for Queue and Hold areas
- General UI improvements with graphics
- Preset levels/Gamemodes for 40 line sprint

[![Gameplay](https://raw.githubusercontent.com/Three6ty1/tetrigo/main/thumbnail.png)](https://raw.githubusercontent.com/Three6ty1/tetrigo/main/gameplay.mp4)