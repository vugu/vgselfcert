# vgselfcert

This is a simple tool which will create a self signed certificate for use in local development.

The main reason this exists is because installing a Go program can be easier than installing `openssl` on some systems.

## Installation

Install with regular Go tooling:

```
go install github.com/vugu/vgselfcert@latest
```

Make sure `go/bin` inside your home directory is in your path.  On Linux/Mac this is typically done by adding a line like `export PATH=$PATH:~/go/bin` to `~/.profile` or wherever the "profile" script is for the shell that you use (e.g. `~/.zprofile` or `~/.bash_profile` are common alternatives).  On Windows you can edit your PATH environment variable through the Control Panel (search for "Environment").

You can verify that the path is correct using `echo $PATH` (Linux/Mac) or `echo %PATH%` (Windows).

## Usage

Once install you normally just want to type `vgselfcert` while being in the directory where you want the output to be.  It will generate `localhost.crt` and `localhost.key` by default.

To see more options run `vgselfcert -h`
