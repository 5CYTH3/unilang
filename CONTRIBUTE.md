# Contribute
To contribute to the project, I recommend you to have chocolatey (windows) or any package manager of your 
linux distribution to install the packages `nasm` and `gcc` (TDM-GCC for windows), required for the 
development of uni (i.e. main.go:65).

## Getting Started

* [MacOS / Linux](#macos-/-linux)
* [Windows](#windows)

### MacOS / Linux

### Windows

Please make sure that you have [scoop](https://scoop.sh/) and [go](https://go.dev/) **1.18.1** installed on your computer (you can use ```install_contribution_req.bat``` to install them).
To install the project, just do these following steps :

```
cd yourdirectory
scoop install make
git clone https://github.com/5CYTH3/unilang.git
cd unilang
make install-windows
```

Building the project would'nt have been easier ! :

```
make build
```
