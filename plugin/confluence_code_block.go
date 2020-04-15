//
// DISCLAIMER
//
// Copyright 2020 ArangoDB GmbH, Cologne, Germany
//
// Author Gergely Brautigam
//

package plugin

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	md "github.com/arangodb-managed/html-to-markdown"
)

// ConfluenceCodeBlock converts `<ac:structured-macro>` elements
func ConfluenceCodeBlock() md.Plugin {
	return func(c *md.Converter) []md.Rule {
		character := "```"
		return []md.Rule{
			md.Rule{
				Filter: []string{"ac:structured-macro"},
				Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
					for _, node := range selec.Nodes {
						if node.Data == "ac:structured-macro" {
							// node's last child -> <ac:plain-text-body>. We don't want to filter on that
							// because we would end up with structured-macro around us.
							// ac:plain-text-body's last child is [CDATA which has the actual content we are looking for.
							content := strings.TrimPrefix(node.LastChild.LastChild.Data, "[CDATA[")
							content = strings.TrimSuffix(content, "]]")
							formatted := fmt.Sprintf("%s\n%s\n%s", character, content, character)
							return md.String(formatted)
						}
					}
					return md.String(character + content + character)
				},
			},
		}
	}
}
