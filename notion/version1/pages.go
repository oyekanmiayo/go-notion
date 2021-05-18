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
// Page can either be *DatabaseParent, *PageParent or *WorkspaceParent
// Properties will contain a key and value. The key is dependent on what the User sets in Notion,
// so we cannoting predetermine the json tag :)
type Page struct {
	Object         string      `json:"object,omitempty"`
	ID             string      `json:"id,omitempty"`
	CreatedTime    string      `json:"created_time,omitempty"`
	LastEditedTime string      `json:"last_edited_time,omitempty"`
	Parent         interface{} `json:"parent,omitempty"`
	Archived       bool        `json:"archived,omitempty"`
	Properties     interface{} `json:"properties,omitempty"`
}

// Type is always "database_id"
type DatabaseParent struct {
	Type       string `json:"type"`
	DatabaseID string `json:"database_id,omitempty"`
}

// Type is always "page_id
type PageParent struct {
	Type   string `json:"type"`
	PageID string `json:"page_id,omitempty"`
}

// Type is always "workspace"
type WorkspaceParent struct {
	Type string `json:"type"`
}

// This contains an extra field corresponding with the value of Type
// https://developers.notion.com/reference/page#all-property-values
// Use []RichText for Title, RichText
// Use *Number for Number
// Use *Select for Select
// Use *MultiSelect for MultiSelect
// Use *Date for Date
// Use *Formula for Formula
// Use []PageReference for Relation
// Use *Rollup for Rollup
// Use []User for People
// Use []FileReference for Files
// Use *User for CreatedBy, LastEditedBy
type PageProperty struct {
	ID             string      `json:"id"`
	Type           string      `json:"type"`
	Title          interface{} `json:"title,omitempty"`
	RichText       interface{} `json:"rich_text,omitempty"`
	Number         interface{} `json:"number,omitempty"`
	Select         interface{} `json:"select,omitempty"`
	URL            string      `json:"url,omitempty"`
	Email          string      `json:"email,omitempty"`
	Phone          interface{} `json:"phone,omitempty"`
	Checkbox       bool        `json:"checkbox,omitempty"`
	MultiSelect    interface{} `json:"multi_select,omitempty"`
	CreatedTime    string      `json:"created_time,omitempty"`
	LastEditedTime string      `json:"last_edited_time,omitempty"`
	Date           interface{} `json:"date,omitempty"`
	CreatedBy      interface{} `json:"created_by,omitempty"`
	LastEditedBy   interface{} `json:"last_edited_by,omitempty"`
	Files          interface{} `json:"files,omitempty"`
	Relation       interface{} `json:"relation,omitempty"`
	Formula        interface{} `json:"formula,omitempty"`
	Rollup         interface{} `json:"rollup,omitempty"`
	People         interface{} `json:"people,omitempty"`
	PhoneNumber    string      `json:"phone_number,omitempty"`
}

type Number struct {
	Number int32 `json:"number,omitempty"`
}

// Color can only be "default", "gray", "brown", "red", "orange", "yellow", "green",
// "blue", "purple" or "pink"
type Select struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

type MultiSelect struct {
	MultiSelect []MultiSelectOptions `json:"multi_select,omitempty"`
}

// When updating a multi-select property, you can use either name or id.
// Setting a multi-select option by name will only update the page if the multi-select database
// property has an option by that name.
// Color can only be "default", "gray", "brown", "red", "orange", "yellow", "green",
// "blue", "purple" or "pink"
// https://developers.notion.com/reference/page#multi-select-property-values
type MultiSelectOptions struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

// If End is not defined, this Date struct isn't considered a range
// https://developers.notion.com/reference/page#date-property-values
type Date struct {
	Start string `json:"star,omitempty"`
	End   string `json:"end,omitempty"`
}

// Type must be one of "string", "number", "boolean", and "date".
// https://developers.notion.com/reference/page#formula-property-values
// Use *Date for Date
type Formula struct {
	Type    string      `json:"type"`
	String  string      `json:"string,omitempty"`
	Number  int32       `json:"number,omitempty"`
	Boolean bool        `json:"boolean,omitempty"`
	Date    interface{} `json:"date,omitempty"`
}

// https://developers.notion.com/reference/page#relation-property-values
type PageReference struct {
	ID string `json:"id"`
}

// https://developers.notion.com/reference/page#rollup-property-values
// Use *Date for Date
// Use []RollupProperty for Array
type Rollup struct {
	Type   string      `json:"type"`
	Number int32       `json:"number,omitempty"`
	Date   interface{} `json:"date,omitempty"`
	Array  interface{} `json:"array,omitempty"`
}

// https://developers.notion.com/reference/page#rollup-property-value-element
// Use []RichText for Title, RichText
// Use *Number for Number
// Use *Select for Select
// Use *MultiSelect for MultiSelect
// Use *Date for Date
// Use *Formula for Formula
// Use []PageReference for Relation
// Use *Rollup for Rollup
// Use []User for People
// Use []FileReference for Files
// Use *User for CreatedBy, LastEditedBy
type RollupProperty struct {
	Type           string      `json:"type"`
	Title          interface{} `json:"title,omitempty"`
	RichText       interface{} `json:"rich_text,omitempty"`
	Number         interface{} `json:"number,omitempty"`
	Select         interface{} `json:"select,omitempty"`
	URL            interface{} `json:"url,omitempty"`
	Email          interface{} `json:"email,omitempty"`
	Phone          interface{} `json:"phone,omitempty"`
	Checkbox       interface{} `json:"checkbox,omitempty"`
	MultiSelect    interface{} `json:"multi_select,omitempty"`
	CreatedTime    interface{} `json:"created_time,omitempty"`
	LastEditedTime interface{} `json:"last_edited_time,omitempty"`
	Date           interface{} `json:"date,omitempty"`
	CreatedBy      interface{} `json:"created_by,omitempty"`
	LastEditedBy   interface{} `json:"last_edited_by,omitempty"`
	Files          interface{} `json:"files,omitempty"`
	Relation       interface{} `json:"relation,omitempty"`
	Formula        interface{} `json:"formula,omitempty"`
	Rollup         interface{} `json:"rollup,omitempty"`
	People         interface{} `json:"people,omitempty"`
	PhoneNumber    interface{} `json:"phone_number,omitempty"`
}

// https://developers.notion.com/reference/page#files-property-values
type FileReference struct {
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
	Properties interface{} `json:"properties,omitempty"`
}

// https://developers.notion.com/reference/patch-page
func (p *PageService) UpdatePageProperties(pageID string,
	params *UpdatePagePropertiesBodyParams) (*Page, *http.Response, error) {
	page := new(Page)
	apiError := new(APIError)
	resp, err := p.sling.New().Patch(pageID).BodyJSON(params).Receive(page, apiError)

	return page, resp, relevantError(err, *apiError)
}