# onsave

An unsophisticated tool to watch a directory and run a command when a file is saved.

## Usage

```
onsave <directory-to-watch> <command-to-run>
```

## Example

```
onsave ./src "npm run format"
```

This will run `npm run format` whenever a file in the `./src` directory is saved.

## Installation

```
go install onsave.go
```

Verify that the binary is in your `$PATH` by running `which onsave`.

```bash
which onsave
```

If it's not, you can add it to your `$PATH` by running `export PATH="$PATH:$(go env GOPATH)/bin"` or by 
adding this to your shell config file (`.zshrc`, `.bashrc`, etc.):

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```
