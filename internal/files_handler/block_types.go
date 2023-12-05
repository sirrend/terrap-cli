package files_handler

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// Block represents a configuration block, which can contain attributes or nested blocks
type Block struct {
	Type       string
	Attributes []*Attribute
	Blocks     []*Block
}

// Init initializes the Block object with the given hclwrite.Block and block type
func (b *Block) Init(block *hclwrite.Block, blockType string) {
	b.Type = blockType
	b.Attributes = []*Attribute{}
	b.Blocks = []*Block{}

	for name, attr := range block.Body().Attributes() {
		a := &Attribute{}
		a.Init(attr, name)
		b.Attributes = append(b.Attributes, a)
	}

	for _, nestedBlock := range block.Body().Blocks() {
		nested := &Block{}
		nested.Init(nestedBlock, nestedBlock.Type())
		b.Blocks = append(b.Blocks, nested)
	}
}
