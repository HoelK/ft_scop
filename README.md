# Introduction

ft_scop is a small graphical project with the goal of parsing a Wavefront .obj file and show it on screen using a graphics API

The parser part will be made in Golang and the rendering in Rust using OpenGL or Vulkan<br>
The parser and rendering will be 2 different processes communicating by an FFI<br>

## Objective
Parse a Wavefront .obj file with mtl file
Show the object in 3D on screen

## Goals
- Learn basic of Go and Rust<br>
- Learn basic of Graphics programming<br>

# Setup

Compile
```bash
make
```

Use
```bash
./parser <file>
```
