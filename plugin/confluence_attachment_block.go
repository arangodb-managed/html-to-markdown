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

	"github.com/PuerkitoBio/goquery"
	md "github.com/arangodb-managed/html-to-markdown"
)

// ConfluenceAttachments converts `<ri:attachment ri:filename=""/>` elements
func ConfluenceAttachments() md.Plugin {
	return func(c *md.Converter) []md.Rule {
		return []md.Rule{
			md.Rule{
				Filter: []string{"ri:attachment"},
				Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
					if v, ok := selec.Attr("ri:filename"); ok {
						formatted := fmt.Sprintf("![][%s]", v)
						return md.String(formatted)
					}
					return md.String("")
				},
			},
		}
	}
}
