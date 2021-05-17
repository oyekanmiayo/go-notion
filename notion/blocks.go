package notion

import (
	"github.com/dghubble/sling"
	"net/http"
)

type BlockService struct {
	sling *sling.Sling
}

// newBlockService returns a new BlockService.
func newBlockService(sling *sling.Sling) *BlockService {
	return &BlockService{
		sling: sling.Path("blocks/"),
	}
}

// https://developers.notion.com/reference/block
// Object is always "block"
// Type must be one of e "paragraph", "heading_1", "heading_2", "heading_3", "bulleted_list_item",
// "numbered_list_item", "to_do", "toggle", "child_page", and "unsupported"
// Use *Paragraph for Paragraph
// Use *HeadingOne for HeadingOne
// Use *HeadingTwo for HeadingTwo
// Use *HeadingThree for HeadingThree
// Use *BulletedListItem for BulletedListItem
// Use *ToDo for ToDo
// Use *Toggle for Toggle
type Block struct {
	Object           string      `json:"object,omitempty"`
	ID               string      `json:"id,omitempty"`
	Type             string      `json:"type,omitempty"`
	CreatedTime      string      `json:"created_time,omitempty"`
	LastEditedTime   string      `json:"last_edited_time,omitempty"`
	HasChildren      bool        `json:"has_children,omitempty"`
	Paragraph        interface{} `json:"paragraph,omitempty"`
	HeadingOne       interface{} `json:"heading_1,omitempty"`
	HeadingTwo       interface{} `json:"heading_2,omitempty"`
	HeadingThree     interface{} `json:"heading_3,omitempty"`
	BulletedListItem interface{} `json:"bulleted_list_item,omitempty"`
	NumberedListItem interface{} `json:"numbered_list_item,omitempty"`
	ToDo             interface{} `json:"to_do,omitempty"`
	Toggle           interface{} `json:"toggle,omitempty"`
	ChildPage        interface{} `json:"child_page,omitempty"`
}

// https://developers.notion.com/reference/block#paragraph-blocks
// Use []RichText for Text
// Use []Block for Children
type Paragraph struct {
	Text     interface{} `json:"text,omitempty"`
	Children interface{} `json:"children,omitempty"`
}

// https://developers.notion.com/reference/block#heading-one-blocks
// Use []RichText for Text
type HeadingOne struct {
	Text interface{} `json:"text,omitempty"`
}

// https://developers.notion.com/reference/block#heading-two-blocks
// Use []RichText for Text
type HeadingTwo struct {
	Text interface{} `json:"text,omitempty"`
}

// https://developers.notion.com/reference/block#heading-three-blocks
// Use []RichText for Text
type HeadingThree struct {
	Text interface{} `json:"text,omitempty"`
}

// https://developers.notion.com/reference/block#bulleted-list-item-blocks
// Use []RichText for Text
// Use []Block for Children
type BulletedListItem struct {
	Text     interface{} `json:"text,omitempty"`
	Children interface{} `json:"children,omitempty"`
}

// https://developers.notion.com/reference/block#numbered-list-item-blocks
// Use []RichText for Text
// Use []Block for Children
type NumberedListItem struct {
	Text     interface{} `json:"text,omitempty"`
	Children interface{} `json:"children,omitempty"`
}

// https://developers.notion.com/reference/block#to-do-blocks
// Use []RichText for Text
// Use []Block for Children
type ToDo struct {
	Text     interface{} `json:"text,omitempty"`
	Checked  bool        `json:"checked,omitempty"`
	Children interface{} `json:"children,omitempty"`
}

// https://developers.notion.com/reference/block#toggle-blocks
// Use []RichText for Text
// Use []Block for Children
type Toggle struct {
	Text     interface{} `json:"text,omitempty"`
	Children interface{} `json:"children,omitempty"`
}

// https://developers.notion.com/reference/block#child-page-blocks
type ChildPage struct {
	Title string `json:"title,omitempty"`
}

type RetrieveBlockChildrenParams struct {
	StartCursor string `url:"start_cursor,omitempty"`
	PageSize    int32  `url:"page_size,omitempty"`
}

type RetrieveBlockChildrenResponse struct {
	Object     string  `json:"object"`
	Results    []Block `json:"results"`
	NextCursor string  `json:"next_cursor"`
	HasMore    bool    `json:"has_more"`
}

// https://developers.notion.com/reference/get-block-children
func (b *BlockService) RetrieveBlockChildren(blockID string,
	params *RetrieveBlockChildrenParams) (*RetrieveBlockChildrenResponse, *http.Response, error) {
	response := new(RetrieveBlockChildrenResponse)
	apiError := new(APIError)
	resp, err := b.sling.New().Get(blockID+"/children").QueryStruct(params).Receive(response, apiError)

	return response, resp, relevantError(err, *apiError)
}

type AppendBlockChildrenBodyParams struct {
	Children []Block `json:"children"`
}

// https://developers.notion.com/reference/patch-block-children
// NB: Blocks cannot be modified currently. Once a block is appended as a child of another block,
// it cannot be updated or deleted.
func (b *BlockService) AppendBlockChildren(blockID string, params *AppendBlockChildrenBodyParams) (*Block, *http.Response, error) {
	block := new(Block)
	apiError := new(APIError)
	resp, err := b.sling.New().Patch(blockID+"/children").BodyJSON(params).Receive(block, apiError)

	return block, resp, relevantError(err, *apiError)
}
