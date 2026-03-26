# Introduction

ft_scop is a small graphical project with the goal of parsing a Wavefront .obj file and show it on screen using a graphics API

The parser part will be made in Golang and the rendering in Rust using OpenGL or Vulkan
The parser and rendering will be 2 different processes communicating by using Unix sockets

The main goal of this project is just to learn the basic of Go, Rust and Graphics programming

# Setup

Type make at the root of the repository then launch the parser with ./parser <file>
