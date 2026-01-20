# Go syntax revision notes

Some basic notes for myself to go over the Go syntax and generic setup steps.

## Installing go

- On Windows go [here](https://go.dev/doc/install) and download and install the `.msi` file
- On Mac get [homebrwew](https://brew.sh/) and `brew install go` or install the `.pkg` from the same link as windows
- On Linux use your package manager apt/dnf/pacman to install `go` or install the `.tar.gz` from the same link as windows

## Making a project

1. Create a directory and navigate to it
2. Initialise the project with:

```bash
go mod init my-project-name
```

3. Create a `main.go` file structured like:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world! 👋")
}
```

4. Run it with:

```bash
go run .
```

5. To create an executable simply run:

```bash
go build .
```

> Go uses your directory in its build system, all `.go` files with `package main` at the top will be pulled in and imported, don't have to have specific import statements for other local files -> the file system is the build system, in a way.

## Adding libraries

1. Add whatever you need in your import statement in `main.go` (can be names or github link):

```go
import (
    "fmt"
    "rsc.io/quote" // 👈 Import the external module
)
```

2. Tidy up your mod file

```bash
go mod tidy
```

3. All done, simply run your program again and sorted

```bash
go run .
```

## Build options

### Shrinking the binary

You can save over 30% of binary size by removing debugging information

```bash
go build -ldflags="-s -w" .
```

- `s`: Omits the symbol table and debug information
- `w`: Omits the DWARF debug information.

### Cross compilation

You can compile for other systems and architectures from the same machine.

- In `bash`:

```bash
GOOS=linux GOARCH=arm64 go build .
```

- In `powershell`

```powershell
$env:GOOS = "linux"
$env:GOARCH = "arm64"
go build .
Remove-Item env:GOOS
Remove-Item env:GOARCH

```

Allowed `GOOS` and `GOARCH` values:

<table>
<tr>
<td valign="top">

| `GOOS` | Target Platform |
| -- | -- |
| `aix` | IBM AIX |
| `android` | Android |
| `darwin` | macOS (and iOS) |
| `dragonfly` | Dragonfly BSD |
| `freebsd` | FreeBSD |
| `illumos` | Illumos |
| `ios` | iOS |
| `js` | JavaScript/WebAssembly |
| `linux` | Linux |
| `netbsd` | NetBSD |
| `openbsd` | OpenBSD |
| `plan9` | Plan 9 |
| `solaris` | Solaris |
| `windows` | Microsoft Windows |

</td>
<td valign="top">

| `GOARCH`	| Target Architecture |
| -- | -- |
| `386` |	32-bit x86 |
| `amd64` |	64-bit x86 |
| `arm` |	32-bit ARM | 
| `arm64` |	64-bit ARM |
| `loong64` |	64-bit LoongArch |
| `mips` | 32-bit MIPS |
| `mipsle` | 32-bit MIPS (little-endian) |
| `mips64` |	64-bit MIPS |
| `mips64le` |	64-bit MIPS (little-endian) |
| `ppc64` |	64-bit PowerPC | 
| `ppc64le` |	64-bit PowerPC (little-endian) |
| `riscv64` |	64-bit RISC-V |
| `s390x` |	64-bit IBM z/Architecture |
| `wasm` |	WebAssembly |

</td>
</tr>
</table>


### Setting variables at build time

Go supports injecting values for variables at compile time - you define them in the code and supply them in the build process.

1. Create the variable in the code:

```go
package main

import "fmt"

var Version string // This variable will be set by the build command

func main() {
    fmt.Printf("Hello from version: %s\n", Version)
}
```

2. Provide it as part of the build

```bash
go build -ldflags="-X 'main.Version=1.0.1-beta'" .
```

### Conditional builds - omitting files based on flags

1. Prepare your `some-name.go` file with a tag:

```go
//go:build production

package main

// This file is only included when the "production" tag is used.
var AppMode = "Production"
```

2. Use the tag in the build:

```bash
go build -tags="production" .
```

### Race condition detection

You can use Go's built in race condition detection tool to debug concurrency:

```bash
go build -race .
```
