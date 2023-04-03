package handle_files

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/sirrend/terrap-cli/internal/rules_api"
	"github.com/sirrend/terrap-cli/internal/utils"
	"github.com/spf13/cast"
)

// Resource holds all the data scraped from user files for a specific resource
type Resource struct {
	Type             string
	Name             string
	Alias            string
	FullNameSequence string
	Provider         tfconfig.ProviderRef
	Pos              tfconfig.SourcePos
	Attributes       []Attribute
}

// Init
/*
@brief:
	Init initializes the Resource object in context
@params:
	block - *hclwrite.Block - the resource to initialize with
	metadata - *tfconfig.Resource - the resource metadata
*/
func (r *Resource) Init(block *hclwrite.Block, metadata *tfconfig.Resource) {
	r.Type = block.Type()
	r.Name = block.Labels()[0]
	r.Alias = block.Labels()[1]
	r.Pos = metadata.Pos
	r.Provider = metadata.Provider
	r.FullNameSequence = r.Type + "." + r.Name + "." + r.Alias
	r.analyzeAttributes(block) // fill the attributes slice
}

// GetAttributesKeys
/*
@brief:
	GetAttributesKeys returns all attributes keys in the Resource in context
@returns:
	keys - []string - the keys found
*/
func (r Resource) GetAttributesKeys() (keys []string) {
	for _, details := range r.Attributes {
		keys = append(keys, details.Name)
	}

	return keys
}

// analyzeAttributes
/*
@brief:
	analyzeAttributes initializes all the attributes found under the given block
@params:
	block - *hclwrite.Block - the block to inspect
*/
func (r *Resource) analyzeAttributes(block *hclwrite.Block) {
	attributeHolder := Attribute{}

	for name, attr := range block.Body().Attributes() {
		attributeHolder.Init(attr, name)
		r.Attributes = append(r.Attributes, attributeHolder)
	}
}

// IsDataSource
/*
@brief:
	IsDataSource checks if the resource in context is of type data
@returns:
	bool
*/
func (r Resource) IsDataSource() bool {
	if r.Type == "data" {
		return true
	}

	return false
}

// IsResource
/*
@brief:
	IsResource checks if the resource in context is of type resource
@returns:
	bool
*/
func (r Resource) IsResource() bool {
	if r.Type == "resource" {
		return true
	}

	return false
}

// GetRuleset
/*
@brief:
	GetRuleset checks if the resource in context is inside the given rulebook and returns it if it does.
	If no RuleSet is found, will return an empty object and no error.
@params:
	rulebook rules_api.Rulebook - the rulebook to search for the ruleset in
@returns:
	*gabs.Container - the ruleset to execute
	error - if exists
*/
func (r Resource) GetRuleset(rulebook rules_api.Rulebook) (rulesetObj rules_api.RuleSet, err error) {
	var rules []rules_api.Rule

	if ruleset, err := rulebook.GetRuleSetByResource(r.Name); err == nil {
		if ruleset != nil {
			if components, err := ruleset.Children(); err == nil {
				for _, component := range components {
					if rulesSlice, err := component.Children(); err == nil {
						for _, rule := range rulesSlice {
							rules = append(rules, rules_api.Rule{
								Path:          utils.MustUnquote(rule.Path("HumanReadablePath").String()),
								Operation:     utils.MustUnquote(rule.Path("Operation").String()),
								ComponentName: utils.MustUnquote(rule.Path("AttributeKey").String()),
								ComponentType: utils.MustUnquote(rule.Path("ResourceComponent").String()),
								Required:      cast.ToBool(rule.Path("IsRequired").String()),
								Notification:  utils.MustUnquote(rule.Path("Notification").String()),
								URL:           utils.MustUnquote(rule.Path("URL").String()),
							})
						}
					}
				}
			}

			rulesetObj = rules_api.RuleSet{
				ResourceName: r.Name,
				Rules:        rules,
			}

			return rulesetObj, nil
		}
	}

	return rulesetObj, err
}
