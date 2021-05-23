package version1

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
// Type must be one of "paragraph", "heading_1", "heading_2", "heading_3", "bulleted_list_item",
// "numbered_list_item", "to_do", "toggle", "child_page", and "unsupported"
type Block struct {
	Object           string                 `json:"object,omitempty"`
	ID               string                 `json:"id,omitempty"`
	Type             string                 `json:"type,omitempty"`
	CreatedTime      string                 `json:"created_time,omitempty"`
	LastEditedTime   string                 `json:"last_edited_time,omitempty"`
	HasChildren      bool                   `json:"has_children,omitempty"`
	Paragraph        *ParagraphBlock        `json:"paragraph,omitempty"`
	HeadingOne       *HeadingOneBlock       `json:"heading_1,omitempty"`
	HeadingTwo       *HeadingTwoBlock       `json:"heading_2,omitempty"`
	HeadingThree     *HeadingThreeBlock     `json:"heading_3,omitempty"`
	BulletedListItem *BulletedListItemBlock `json:"bulleted_list_item,omitempty"`
	NumberedListItem *NumberedListItemBlock `json:"numbered_list_item,omitempty"`
	ToDo             *ToDoBlock             `json:"to_do,omitempty"`
	Toggle           *ToggleBlock           `json:"toggle,omitempty"`
	ChildPage        *ChildPageBlock        `json:"child_page,omitempty"`
}

// https://developers.notion.com/reference/block#paragraph-blocks
type ParagraphBlock struct {
	Text     []RichText `json:"text,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// https://developers.notion.com/reference/block#heading-one-blocks
type HeadingOneBlock struct {
	Text []RichText `json:"text,omitempty"`
}

// https://developers.notion.com/reference/block#heading-two-blocks
type HeadingTwoBlock struct {
	Text []RichText `json:"text,omitempty"`
}

// https://developers.notion.com/reference/block#heading-three-blocks
type HeadingThreeBlock struct {
	Text []RichText `json:"text,omitempty"`
}

// https://developers.notion.com/reference/block#bulleted-list-item-blocks
type BulletedListItemBlock struct {
	Text     []RichText `json:"text,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// https://developers.notion.com/reference/block#numbered-list-item-blocks
type NumberedListItemBlock struct {
	Text     []RichText `json:"text,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// https://developers.notion.com/reference/block#to-do-blocks
type ToDoBlock struct {
	Text     []RichText `json:"text,omitempty"`
	Checked  bool       `json:"checked,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// https://developers.notion.com/reference/block#toggle-blocks
type ToggleBlock struct {
	Text     []RichText `json:"text,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// https://developers.notion.com/reference/block#child-page-blocks
type ChildPageBlock struct {
	Title string `json:"title,omitempty"`
}

type RetrieveBlockChildrenParams struct {
	StartCursor string `url:"start_cursor,omitempty"`
	PageSize    int32  `url:"page_size,omitempty"`
}

type RetrieveBlockChildrenResponse struct {
	Object     string  `json:"object,omitempty"`
	Results    []Block `json:"results,omitempty"`
	NextCursor string  `json:"next_cursor,omitempty"`
	HasMore    bool    `json:"has_more,omitempty"`
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
