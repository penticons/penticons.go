package octicons

import (
 "fmt"
 "encoding/base64"
 "github.com/pravj/octicons.go/octicon"
)

func Generate(arg string) string {
 o := octicon.New(arg)

 return o.Svg_string()
}

func Base64_string(arg string) string {
 svg_str := Generate(arg)
 base64_str := base64.StdEncoding.EncodeToString([]byte(svg_str))

 return base64_str
}

func Uri_image(arg string) string {
 base64_str := Base64_string(arg)

 return fmt.Sprintf("url(data:image/svg+xml;base64,%s);", base64_str)
}
