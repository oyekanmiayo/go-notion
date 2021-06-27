package version1

import (
	"github.com/dghubble/sling"
	"net/http"
)

type PageService struct {
	sling *sling.Sling
}

// newPageService returns a new PageService.
func newPageService(sling *sling.Sling) *PageService {
	return &PageService{
		sling: sling.Path("pages/"),
	}
}

// https://developers.notion.com/reference/page#all-pages
// Object must always be "page"
// For Properties,
// - If parent.type is "page_id" or "workspace", then the only valid key is title.
// - If parent.type is "database_id", then the keys and values of this field are determined by the properties
//   of the database this page belongs to.
// Page can be *DatabaseParent, *PageParent or *WorkspaceParent
type Page struct {
	Object         string                  `json:"object,omitempty"`
	ID             string                  `json:"id,omitempty"`
	CreatedTime    string                  `json:"created_time,omitempty"`
	LastEditedTime string                  `json:"last_edited_time,omitempty"`
	Parent         interface{}             `json:"parent,omitempty"`
	Archived       bool                    `json:"archived,omitempty"`
	Properties     map[string]PageProperty `json:"properties,omitempty"`
}

// Type is always "database_id" for Database Parent
type DatabaseParent struct {
	Type       string `json:"type,omitempty"`
	DatabaseID string `json:"database_id,omitempty"`
}

// Type is always "page_id
type PageParent struct {
	Type   string `json:"type,omitempty"`
	PageID string `json:"page_id,omitempty"`
}

// Type is always "workspace"
type WorkspaceParent struct {
	Type string `json:"type,omitempty"`
}

// Type can be only one of "rich_text", "number", "select", "multi_select", "date", "formula",
// "relation", "rollup", "title", "people", "files", "checkbox", "url", "email", "phone_number",
// "created_time", "created_by", "last_edited_time", and "last_edited_by".
// https://developers.notion.com/reference/page#all-property-values
type PageProperty struct {
	ID             string                    `json:"id,omitempty"`
	Type           string                    `json:"type,omitempty"`
	Title          []RichText                `json:"title,omitempty"`
	RichText       []RichText                `json:"rich_text,omitempty"`
	Number         int64                     `json:"number,omitempty"`
	Select         *SelectProperty           `json:"select,omitempty"`
	URL            string                    `json:"url,omitempty"`
	Email          string                    `json:"email,omitempty"`
	Phone          interface{}               `json:"phone,omitempty"`
	Checkbox       bool                      `json:"checkbox,omitempty"`
	MultiSelect    []MultiSelectPropertyOpts `json:"multi_select,omitempty"`
	CreatedTime    string                    `json:"created_time,omitempty"`
	LastEditedTime string                    `json:"last_edited_time,omitempty"`
	Date           *DateProperty             `json:"date,omitempty"`
	CreatedBy      *User                     `json:"created_by,omitempty"`
	LastEditedBy   *User                     `json:"last_edited_by,omitempty"`
	Files          []FileReferenceProperty   `json:"files,omitempty"`
	Relation       []PageReferenceProperty   `json:"relation,omitempty"`
	Formula        *FormulaProperty          `json:"formula,omitempty"`
	Rollup         *RollupProperty           `json:"rollup,omitempty"`
	People         []User                    `json:"people,omitempty"`
	PhoneNumber    string                    `json:"phone_number,omitempty"`
}

type NumberProperty struct {
	Number int32 `json:"number,omitempty"`
}

// Color can only be "default", "gray", "brown", "red", "orange", "yellow", "green",
// "blue", "purple" or "pink"
type SelectProperty struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

// When updating a multi-select property, you can use either name or id.
// Setting a multi-select option by name will only update the page if the multi-select database
// property has an option by that name.
// Color can only be "default", "gray", "brown", "red", "orange", "yellow", "green",
// "blue", "purple" or "pink"
// https://developers.notion.com/reference/page#multi-select-property-values
type MultiSelectPropertyOpts struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

// If End is not defined, this Date struct isn't considered a range
// https://developers.notion.com/reference/page#date-property-values
type DateProperty struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

// Type must be one of "string", "number", "boolean", and "date".
// https://developers.notion.com/reference/page#formula-property-values
// Use *Date for Date
type FormulaProperty struct {
	Type    string        `json:"type"`
	String  string        `json:"string,omitempty"`
	Number  int32         `json:"number,omitempty"`
	Boolean bool          `json:"boolean,omitempty"`
	Date    *DateProperty `json:"date,omitempty"`
}

// https://developers.notion.com/reference/page#relation-property-values
type PageReferenceProperty struct {
	ID string `json:"id,omitempty"`
}

// https://developers.notion.com/reference/page#rollup-property-values
// Use *DateProperty for Date
// Use  for Array
type RollupProperty struct {
	Type   string                  `json:"type,omitempty"`
	Number int32                   `json:"number,omitempty"`
	Date   *DateProperty           `json:"date,omitempty"`
	Array  []RollupPropertyElement `json:"array,omitempty"`
}

// https://developers.notion.com/reference/page#rollup-property-value-element
// Type is one of rich_text", "number", "select", "multi_select", "date", "formula", "relation",
// "rollup", "title", "people", "files", "checkbox", "url", "email", "phone_number", "created_time",
// "created_by", "last_edited_time", or "last_edited_by"
type RollupPropertyElement struct {
	Type           string                    `json:"type"`
	Title          []RichText                `json:"title,omitempty"`
	RichText       []RichText                `json:"rich_text,omitempty"`
	Number         int64                     `json:"number,omitempty"`
	Select         *SelectProperty           `json:"select,omitempty"`
	URL            string                    `json:"url,omitempty"`
	Email          string                    `json:"email,omitempty"`
	Phone          interface{}               `json:"phone,omitempty"`
	Checkbox       bool                      `json:"checkbox,omitempty"`
	MultiSelect    []MultiSelectPropertyOpts `json:"multi_select,omitempty"`
	CreatedTime    string                    `json:"created_time,omitempty"`
	LastEditedTime string                    `json:"last_edited_time,omitempty"`
	Date           *DateProperty             `json:"date,omitempty"`
	CreatedBy      *User                     `json:"created_by,omitempty"`
	LastEditedBy   *User                     `json:"last_edited_by,omitempty"`
	Files          []FileReferenceProperty   `json:"files,omitempty"`
	Relation       interface{}               `json:"relation,omitempty"`
	Formula        *FormulaProperty          `json:"formula,omitempty"`
	Rollup         interface{}               `json:"rollup,omitempty"`
	People         []User                    `json:"people,omitempty"`
	PhoneNumber    string                    `json:"phone_number,omitempty"`
}

// https://developers.notion.com/reference/page#files-property-values
type FileReferenceProperty struct {
	Name string `json:"name,omitempty"`
}

// https://developers.notion.com/reference/get-page
func (p *PageService) RetrievePage(pageID string) (*Page, *http.Response, error) {
	page := new(Page)
	apiError := new(APIError)
	resp, err := p.sling.New().Get(pageID).Receive(page, apiError)

	return page, resp, relevantError(err, *apiError)
}

// Use *DatabaseParent or *PageParent for Parent
// Read https://developers.notion.com/reference/post-page before using this
type CreatePageBodyParams struct {
	Parent     interface{} `json:"parent,omitempty"`
	Properties interface{} `json:"properties,omitempty"`
	Children   []Block     `json:"children,omitempty"`
}

// https://developers.notion.com/reference/post-page
func (p *PageService) CreatePage(params *CreatePageBodyParams) (*Page, *http.Response, error) {
	page := new(Page)
	apiError := new(APIError)
	resp, err := p.sling.New().Post("").BodyJSON(params).Receive(page, apiError)

	return page, resp, relevantError(err, *apiError)
}

type UpdatePagePropertiesBodyParams struct {
	Properties map[string]PageProperty `json:"properties,omitempty"`
}

// https://developers.notion.com/reference/patch-page
func (p *PageService) UpdatePageProperties(pageID string,
	params *UpdatePagePropertiesBodyParams) (*Page, *http.Response, error) {
	page := new(Page)
	apiError := new(APIError)
	resp, err := p.sling.New().Patch(pageID).BodyJSON(params).Receive(page, apiError)

	return page, resp, relevantError(err, *apiError)
}
