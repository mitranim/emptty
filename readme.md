## Overview

Clears the terminal, for real. (But probably only on my system.)

## Why

Googling brings up bizarre advice:

* Calling a subprocess, when all you needed was to print a few bytes.
* Scrolling the viewport without clearing the scrollback.

Enough! `emptty` provides several functions including a "hard clear" that should empty the TTY for real.

## Usage

```go
import "github.com/mitranim/emptty"

emptty.ClearHard()
```

See https://pkg.go.dev/github.com/mitranim/emptty for the full API.

## Compatibility

Works on everything I tested:

* MacOS Terminal.
* iTerm2.
* Microsoft Terminal running Ubuntu under WSL.
* Terminus (Sublime Text plugin).

May or may not work in other emulators.

## License

https://unlicense.org

## Misc

I'm receptive to suggestions. If this library _almost_ satisfies you but needs changes, open an issue or chat me up. Contacts: https://mitranim.com/#contacts
