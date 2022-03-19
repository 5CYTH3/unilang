# Contribute
To contribute to the project, I recommend you to have chocolatey (windows) or any package manager of your 
linux distribution to install the packages `nasm` and `gcc` (TDM-GCC for windows), required for the 
development of uni (i.e. main.go:65).

# Getting Started

* [MacOS / Linux](#macos-/-linux)
* [Windows](#windows)

## MacOS / Linux
> Make sure you have [go](https://go.dev/) and `nasm` installed on your machine (you can installe them using your default package manager such as **pacman** or **apt**)

To install all the dependencies required, just follow these steps (we will use pacman as an example)

```
cd yourdirectory
pacman -S nasm
git clone git clone https://github.com/5CYTH3/unilang.git

```

## Windows

> Please make sure that you have [chocolatey](https://chocolatey.org/) and [go **1.18**](https://go.dev/) installed on your computer.
To install all the dependencies required, just follow these steps :

```
cd yourdirectory
choco install make
choco install nasm
choco install mingw
git clone https://github.com/5CYTH3/unilang.git
cd unilang
```
*(You can skip the "mingw" package if you already have any C/C++ compiler installed on your computer)*

## Build part
Building the project would'nt have been easier !

```
make build-go -C src
```
