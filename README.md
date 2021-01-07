# gitu

Switch git user fastly.

## Install

```sh
GO111MODULES=on go get -u github.com/haunt98/gitu
```

## Usage

Add new git user:

```sh
gitu add
```

Switch to saved git user:

```sh
gitu switch
```

Show current git user:

```sh
gitu status
```

List all saved git user:

```sh
gitu list
```

Delete saved git user:

```sh
gitu delete
```

## Thanks

- [fatih/color](https://github.com/fatih/color)
- [go-git/go-git](https://github.com/go-git/go-git/)
- [urfave/cli](https://github.com/urfave/cli)
