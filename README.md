# Mini Me

Inspired by the
old [minecraft skin generator](https://web.archive.org/web/20170504033631/http://minecraftskinavatar.com/) website, you
can generate little mini-mes from a Minecraft skin. This program supports both 64x64 and 128x128 skins, but nothing
else.

# Usage

## Standalone program

You can use `go install github.com/twistedasylummc/minime` to install a standalone version of the program. You can then
use it in the command line using the command `minime <input> <output>`.

## Using in a project

First run `go get github.com/twistedasylummc/minime` to install the dependency. You can then use one of the following
methods to generate a mini-me:

- `minime.Skin64(src image.Image) (dst image.Image)`
- `minime.Skin128(src image.Image) (dst image.Image)`