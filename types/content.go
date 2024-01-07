package types

type Block interface{}

type Content[T Block] struct {
	ContentType string
	Block       T
}

func NewContent() Content[Block] {
	return Content[Block]{
		ContentType: "",
		Block:       nil,
	}
}

func (c *Content[Block]) SetContentType(contentType string) {
	c.ContentType = contentType
}

func (c *Content[Block]) SetBlock(block Block) {
	c.Block = block
}

func (c *Content[Block]) GetContentType() string {
	return c.ContentType
}

func (c *Content[Block]) GetBlock() Block {
	return c.Block
}
