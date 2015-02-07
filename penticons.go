// Package penticons generates identicons similar to
// GitHub's contribution activity calendar
package penticons

import (
	"encoding/base64"
	"fmt"
	"github.com/penticons/penticons.go/penticon"
)

// Returns identicons's SVG string representation
func Generate(arg string) string {
	p := penticon.New(arg)

	return p.Svg_string()
}

// Returns Base64 encoded value of identicon's SVG string representation
func Base64_string(arg string) string {
	svg_str := Generate(arg)
	base64_str := base64.StdEncoding.EncodeToString([]byte(svg_str))

	return base64_str
}

// Returns URI Image string for identicon
func Uri_image(arg string) string {
	base64_str := Base64_string(arg)

	return fmt.Sprintf("url(data:image/svg+xml;base64,%s);", base64_str)
}
