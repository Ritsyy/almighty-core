//************************************************************************//
// API "alm": Application Resource Href Factories
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/almighty/almighty-core/design
// --out=$(GOPATH)/src/github.com/almighty/almighty-core
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// VersionHref returns the resource href.
func VersionHref() string {
	return fmt.Sprintf("/api/version")
}
