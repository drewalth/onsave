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
