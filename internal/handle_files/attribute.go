package handle_files

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"strings"
)

// Attribute holds the breakdown of an attribute found under a Resource
type Attribute struct {
	Name  string
	Value string
}

// Init
/*
@brief:
	Init initializes the Attribute object in context
@params:
	attribute - *hclwrite.Attribute - the attribute to initialize
	name - string - the attribute name
*/
func (a *Attribute) Init(attribute *hclwrite.Attribute, name string) {
	a.Name = name
	a.getAttributeValue(attribute)
}

// getAttributeValue
/*
@brief:
	getAttributeValue retrieves the attribute value from the *hclwrite.Attribute object and updates the attribute in context in-place
@params:
	attribute - *hclwrite.Attribute - the attribute to inspect
*/
func (a *Attribute) getAttributeValue(attribute *hclwrite.Attribute) {
	var attrExpr string
	for _, token := range attribute.BuildTokens(nil) {
		attrExpr += string(token.Bytes)
	}

	// split in equality rune
	i := strings.IndexRune(attrExpr, '=')
	attrExpr = strings.ReplaceAll(attrExpr[i+1:], "\n", "")

	a.Value = attrExpr
}
