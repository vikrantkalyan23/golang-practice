/*
GO INTRODUCTION: SETUP, TOOLS, EXECUTION, AND PROGRAM DESIGN

Read the numbered sections, then run this file to see the example below.
Terminal commands in this comment are instructions, not Go statements.

1. DOWNLOAD AND INSTALLATION

Official downloads: https://go.dev/dl/
Official installation instructions: https://go.dev/doc/install

Choose a supported stable release for your operating system and CPU:
  - Apple Silicon Macs: arm64; Intel Macs: amd64.
  - Windows/Linux: select the architecture matching your machine.

macOS:

	Download and run the official .pkg installer, follow its prompts, and reopen
	your terminal. The standard installation places Go in /usr/local/go.

Windows:

	Download and run the official .msi installer. Follow its prompts and reopen
	PowerShell or Command Prompt so it sees the updated PATH.

Linux:

	Download the matching archive and follow the official extraction instructions.
	When upgrading, follow the guide's replacement procedure rather than extracting
	a new release over an existing Go tree. Add /usr/local/go/bin to PATH if needed:

	    export PATH="$PATH:/usr/local/go/bin"

	Put that line in the appropriate shell configuration file for persistence.
	Installing into a system directory may require administrator permissions.

Verify installation in a new terminal:

	  go version
	  go env GOROOT GOPATH GOMOD

	GOROOT: the Go toolchain installation directory; normally set automatically.
	GOPATH: workspace/cache location used by tooling; projects can live elsewhere.
	GOMOD: the active module file, or a special value when no module is active.

If "go" is not found, check installation and PATH, then restart the terminal/editor.
You do not normally need to set GOROOT manually or place projects inside GOPATH.

2. EDITORS, VS CODE EXTENSIONS, AND DEDICATED IDEs

An editor is optional: the Go toolchain can build programs from a terminal.
Go does not ship an official dedicated IDE. The Go project lists editor choices:
https://go.dev/doc/editors

VS Code setup:
 1. Install VS Code: https://code.visualstudio.com/
 2. Install the official "Go" extension published by the Go Team at Google.
    Extension ID: golang.go
    Source and setup: https://github.com/golang/vscode-go
 3. Open your project folder and a .go file.
 4. Follow the extension's prompts to install its supporting tools.
    Use the Command Palette command "Go: Install/Update Tools" when needed.

If the VS Code command-line launcher is available:

	code --install-extension golang.go

The extension integrates:
  - gopls: the official Go language server for completion, navigation, and analysis.
  - Formatting and import organization.
  - Test execution and diagnostics.
  - Delve (dlv): the Go debugger used for breakpoints and inspecting variables.

These tools support the editor; the Go installation supplies the compiler.
No separate Code Runner extension is required.

Dedicated IDE alternative:

	GoLand from JetBrains is a dedicated third-party Go IDE, not an official
	Go-project IDE. It provides Go editing, debugging, refactoring, and testing.
	Product details: https://www.jetbrains.com/go/

Other editors can integrate gopls; see the official editor list above.

3. TERMINAL COMMANDS: RUN, BUILD, FORMAT, AND CHECK

Run commands from the directory containing this file.

For this repository's independent lessons:

	go run 0_introduction.go
	go run 0_introduction.go Vikrant
	go build -o introduction 0_introduction.go
	./introduction Vikrant

"go run" compiles and runs a temporary executable. "go build" creates an
executable you can run later without invoking the Go compiler again.
Arguments after the filename are passed to this program through os.Args.

Windows PowerShell build/run:

	go build -o introduction.exe 0_introduction.go
	./introduction.exe Vikrant

Format and inspect this lesson:

	gofmt -w 0_introduction.go
	go vet 0_introduction.go
	go doc fmt.Println
	go help build

IMPORTANT FOR THIS PRACTICE REPOSITORY:

	1_hello-go.go also declares func main. Run or build one lesson filename at
	a time. Combining both files in one package produces a duplicate-main error.
	The numeric filename prefix orders lessons for readers; it does not control
	which function executes. No go.mod is needed for this standard-library-only,
	explicitly named single-file example.

For a NEW application in its OWN directory, use a module:

	mkdir hello-app
	cd hello-app
	go mod init example.com/hello-app

Create main.go there, then use these commands from its module root:

	go run .                  # Compile and run the current package.
	go build -o hello .        # Build the current executable package.
	go fmt ./...              # Format Go packages throughout the module.
	go vet ./...              # Report suspicious constructs.
	go test ./...             # Run tests in files ending in _test.go.
	go test -race ./...        # Test with race detection on supported systems.
	go test -cover ./...       # Report statement coverage during tests.
	go mod tidy               # Reconcile imports and module dependencies.
	go get example.com/lib    # Example syntax; replace with a real module path.
	go install ./...          # Install executable packages into GOBIN or GOPATH/bin.

Use your repository's real module path for a shared/published project.
"./..." is a Go package pattern covering the directory tree within module rules.
Module commands above describe a separate application, not a request to combine
the standalone lessons in this repository into one executable.

4. PROGRAM STRUCTURE

A small executable usually has these parts:

	  package main              // Package declaration.
	  import "fmt"              // Dependency declarations.
	  const message = "Hello"   // Optional package-level declarations.
	  func main() { ... }       // Executable entry point.
	  func helper() { ... }     // Supporting behavior.

	- package main identifies an executable package. Reusable packages use other
	  names and do not need an executable entry point.
	- Imports belong after the package clause and before other declarations.
	  Unused imports and unused local variables cause compilation errors.
	- func main() has no parameters and no return values. Read command-line input
	  through os.Args or a flag parser instead of declaring main parameters.
	- Top-level constants, types, variables, and functions can appear in different
	  orders after imports. A helper can be declared after the function calling it.
	- Braces delimit blocks. Go normally inserts semicolons automatically.
	- Names starting with an uppercase letter are exported from their package.
	- Files in one package share package-level declarations; file boundaries do
	  not provide namespaces for duplicate functions.

5. FLOW OF A GO PROGRAM

Build time:

	Source files -> parse/type-check -> compile -> link executable.
	Syntax/type errors stop the build before program execution begins.

Runtime:

	Go runtime starts
	  -> imported packages initialize before the packages that import them
	  -> package-level variables initialize in dependency order
	  -> that package's init functions run, if present
	  -> main.main runs after initialization completes
	  -> statements execute in order, with calls/branches/loops changing the flow
	  -> deferred calls run as their surrounding function returns
	  -> returning from main ends the program.

Do not assume that every imported package initializes in textual import order.
Prefer explicit setup in ordinary functions over hidden init side effects.
If you start goroutines, coordinate their completion: returning from main does
not wait for them. os.Exit ends the process without running deferred calls.

The example below follows this path:

	main -> run -> read optional name -> greeting -> write output -> return.
	If validation or writing fails, the error returns to main for reporting.

6. HOW TO DESIGN A PROGRAM IN AN IDIOMATIC WAY

There is no single best folder structure for every application. Begin with
the simplest design that makes the required behavior clear.

 1. Define the problem: inputs, outputs, constraints, and expected failures.
    Example: accept an optional name and print one greeting; reject blank input.
 2. Sketch the flow before adding packages or frameworks.
 3. Keep main focused on startup, connecting dependencies, and reporting errors.
    Put behavior in functions that can return errors and be tested directly.
 4. Separate core decisions from external input/output where that helps.
    Here, greeting computes a message; run handles arguments and output.
 5. Use meaningful names, short cohesive functions, and useful zero values.
    Apply gofmt and document the intent of exported APIs.
 6. Handle errors explicitly. Add context when it helps callers; use %w when
    wrapping an error that callers may need to inspect. Avoid panic for normal
    invalid input, network failures, or missing files.
 7. Pass dependencies explicitly. Use small interfaces at useful boundaries;
    io.Writer lets this example write to a terminal, buffer, or file.
    Do not introduce an interface for every type without a concrete need.
 8. Keep resource ownership clear. Close resources you open, usually with defer
    after checking the open error. Add context and timeouts to blocking work.
 9. Add concurrency only for a reason; bound it and plan cancellation/shutdown.
 10. Test behavior and failure paths in _test.go files. Run format, vet, and tests;
    measure performance before optimizing.
 11. Start in one package. Extract cohesive packages as responsibilities grow;
    avoid generic "utils" packages and unnecessary abstraction layers.
 12. Validate external input and keep secrets out of source code and logs.

A possible layout once an application grows:

	hello-app/
	  go.mod
	  main.go                 # Enough for an initially small command.
	  main_test.go

For multiple commands or a larger service, a possible later layout is:

	hello-app/
	  go.mod
	  cmd/hello/main.go       # Executable startup and wiring.
	  internal/greeting/      # Application-internal behavior and tests.

cmd is a convention, not a required directory. internal has an enforced Go
import restriction. Do not create empty layers merely to match a template.

Further reading:

	https://go.dev/doc/modules/layout
	https://go.dev/doc/effective_go
	https://go.dev/wiki/CodeReviewComments
*/
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// main connects process input/output to the application and reports failures.
func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "introduction:", err)
		os.Exit(1)
	}
}

// run accepts dependencies explicitly so callers can supply arguments and output.
func run(args []string, out io.Writer) error {
	if len(args) > 1 {
		return fmt.Errorf("expected at most one name; put multiword names in quotes")
	}

	name := "Go learner"
	if len(args) == 1 {
		name = args[0]
	}

	message, err := greeting(name)
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintln(out, message); err != nil {
		return fmt.Errorf("write greeting: %w", err)
	}
	return nil
}

// greeting contains the input validation and message construction without I/O.
func greeting(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", fmt.Errorf("name must not be blank")
	}
	return fmt.Sprintf("Hello, %s! Welcome to Go.", name), nil
}
