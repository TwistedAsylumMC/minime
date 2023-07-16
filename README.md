# Mini Me

Inspired by the
old [minecraft skin generator](https://web.archive.org/web/20170504033631/http://minecraftskinavatar.com/) website, you
can generate little mini-mes from a Minecraft skin. This program supports both 64x64 and 128x128 skins, but nothing
else.

# Usage

## Standalone program

You can run the example program by running `go run cmd/main.go` and providing an input and output file. For example:

```shell
go run cmd/main.go steve.png steve-mini.png
```

## Using in a project

First run `go get github.com/twistedasylummc/minime` to install the dependency. You can then use one of the following
methods to generate a mini-me:

- `minime.Skin64(src image.Image) (dst image.Image)`
- `minime.Skin128(src image.Image) (dst image.Image)`