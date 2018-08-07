/*
Package version sets version information for the binary where it is imported.
The version can be retrieved either from the -version command line argument,
or from the /debug/version/ http endpoint.

To include in a project simply import the package and call version.Init().

The version and compile date is stored in version and date variables and
are supposed to be set during compile time. Typically this is done by the
Makefile.

To set these manually use -ldflags together with -X, like in this example:

	go build -ldflags "-X backend/version.version v1.2.3"

*/

package version

import (
	"bytes"
	"flag"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
)

var showVersion = flag.Bool("version", false, "Print version of this binary (only valid if compiled with make)")

var (
	version     string
	date        string
	commit      string
	serviceName string
)

func init() {
	http.Handle("/debug/version/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		printVersion(w, html.EscapeString(version), html.EscapeString(commit), html.EscapeString(date))
	}))
}

// Plz define serviceName in tantan-backend-common/service_name.go
func Init(name string) {
	serviceName = name
	if !flag.Parsed() {
		flag.Parse()
	}

	if showVersion != nil && *showVersion {
		printVersion(os.Stdout, version, commit, date)
		os.Exit(0)
	}
}

func Version() string {
	return version
}
func CompileDate() string {
	return date
}

func ServiceName() string {
	return serviceName
}

func Info() string {
	buf := &bytes.Buffer{}
	printVersion(buf, version, commit, date)
	return buf.String()
}

func printVersion(w io.Writer, version string, commit string, date string) {
	fmt.Fprintf(w, "Version: %s\n", version)
	fmt.Fprintf(w, "CommitID: %s\n", commit)
	fmt.Fprintf(w, "Binary: %s\n", os.Args[0])
	fmt.Fprintf(w, "Compile date: %s\n", date)
	fmt.Fprintf(w, "(version and date only valid if compiled with make)\n")
}
