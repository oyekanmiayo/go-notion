package version1

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

type DatabaseService struct {
	sling *sling.Sling
}

// newDatabaseService returns a new DatabaseService.
func newDatabaseService(sling *sling.Sling) *DatabaseService {
	return &DatabaseService{
		sling: sling.Path("databases/"),
	}
}

type Database struct {
	Object         string      `json:"object,omitempty"`
	ID             string      `json:"id,omitempty"`
	CreatedTime    string      `json:"created_time,omitempty"`
	LastEditedTime string      `json:"last_edited_time,omitempty"`
	Title          []RichText  `json:"title,omitempty"`
	Properties     interface{} `json:"properties,omitempty"`
}

func (d *DatabaseService) RetrieveDatabase(databaseID string) (*Database, *http.Response, error) {
	database := new(Database)
	apiError := new(APIError)
	resp, err := d.sling.New().Get(databaseID).Receive(database, apiError)

	return database, resp, relevantError(err, *apiError)
}

type QueryDatabaseBodyParams struct {
	Filter      interface{}   `json:"filter,omitempty"`
	Sorts       []SortDetails `json:"sorts,omitempty"`
	StartCursor string        `json:"start_cursor,omitempty"`
	PageSize    int32         `json:"page_size,omitempty"`
}

// Timestamp must either be "created_time" or "last_edited_time"
// Direction must either be "ascending" or "descending"
type SortDetails struct {
	Property  string `json:"property,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Direction string `json:"direction,omitempty"`
}

// Use []SingleFilter or []CompoundFilter for OR
// Use []SingleFilter or []CompoundFilter for AND
type CompoundFilter struct {
	OR  interface{} `json:"or,omitempty"`
	AND interface{} `json:"and,omitempty"`
}

// We use interfaces because many of the condition structs contain the same **json tags**.
// Setting them explicity here will throw an error. It is extremely important
// to fulfil each interface appropriately, so please look carefully at the guide below:
// Use *TextCondition for Title, RichText, URL, Email, Phone
// Use *NumberCondition for Number
// Use *CheckboxCondition for Checkbox
// Use *SelectCondition for Select
// Use *MultiSelectCondition for MultiSelect
// Use *DateCondition for CreatedTime, LastEditedTime
// Use *DateCondition or PeopleCondition for Date
// Use *PeopleCondition for CreatedBy, LastEditedBy
// Use *FileCondition for Files
// Use *RelationCondition for Relation
type SingleFilter struct {
	Property       string      `json:"property,omitempty"`
	Title          interface{} `json:"title,omitempty"`
	RichText       interface{} `json:"rich_text,omitempty"`
	URL            interface{} `json:"url,omitempty"`
	Email          interface{} `json:"email,omitempty"`
	Phone          interface{} `json:"phone,omitempty"`
	Number         interface{} `json:"number,omitempty"`
	Checkbox       interface{} `json:"checkbox,omitempty"`
	Select         interface{} `json:"select,omitempty"`
	MultiSelect    interface{} `json:"multi_select,omitempty"`
	CreatedTime    interface{} `json:"created_time,omitempty"`
	LastEditedTime interface{} `json:"last_edited_time,omitempty"`
	Date           interface{} `json:"date,omitempty"`
	CreatedBy      interface{} `json:"created_by,omitempty"`
	LastEditedBy   interface{} `json:"last_edited_by,omitempty"`
	Files          interface{} `json:"files,omitempty"`
	Relation       interface{} `json:"relation,omitempty"`
	Formula        interface{} `json:"formula,omitempty"`
}

type TextCondition struct {
	Equals         string `json:"equals,omitempty"`
	DoesNotEqual   string `json:"does_not_equal,omitempty"`
	Contains       string `json:"contains,omitempty"`
	DoesNotContain string `json:"does_not_contain,omitempty"`
	StartsWith     string `json:"starts_with,omitempty"`
	EndsWith       string `json:"ends_with,omitempty"`
	IsEmpty        bool   `json:"is_empty,omitempty"`
	isNotEmpty     bool   `json:"is_not_empty,omitempty"`
}

type NumberCondition struct {
	Equals               int32 `json:"equals,omitempty"`
	DoesNotEqual         int32 `json:"does_not_equal,omitempty"`
	GreaterThan          int32 `json:greater_than, omitempty`
	LessThan             int32 `json:"less_than,omitempty"`
	GreaterThanOrEqualTo int32 `json:"greater_than_or_equal_to,omitempty"`
	LessThanOrEqualTo    int32 `json:"less_than_or_equal_to,omitempty"`
	IsEmpty              bool  `json:"is_empty,omitempty"`
	isNotEmpty           bool  `json:"is_not_empty,omitempty"`
}

type CheckboxCondition struct {
	Equals       bool `json:"equals,omitempty"`
	DoesNotEqual bool `json:"does_not_equal,omitempty"`
}

type SelectCondition struct {
	Equals       string `json:"equals,omitempty"`
	DoesNotEqual string `json:"does_not_equal,omitempty"`
	IsEmpty      bool   `json:"is_empty,omitempty"`
	isNotEmpty   bool   `json:"is_not_empty,omitempty"`
}

type MultiSelectCondition struct {
	Contains       string `json:"contains,omitempty"`
	DoesNotContain string `json:"does_not_contain,omitempty"`
	IsEmpty        bool   `json:"is_empty,omitempty"`
	isNotEmpty     bool   `json:"is_not_empty,omitempty"`
}

type DateCondition struct {
	Equals     string      `json:"equals,omitempty"`
	Before     string      `json:"before,omitempty"`
	After      string      `json:"after,omitempty"`
	OnOrBefore string      `json:"on_or_before,omitempty"`
	IsEmpty    bool        `json:"is_empty,omitempty"`
	IsNotEmpty bool        `json:"is_not_empty,omitempty"`
	OnOrAfter  string      `json:"on_or_after,omitempty"`
	PastWeek   interface{} `json:"past_week,omitempty"`
	PastMonth  interface{} `json:"past_month,omitempty"`
	PastYear   interface{} `json:"past_year,omitempty"`
	NextWeek   interface{} `json:"next_week,omitempty"`
	NextMonth  interface{} `json:"next_month,omitempty"`
	NextYear   interface{} `json:"next_year,omitempty"`
}

type PeopleCondition struct {
	Contains       string `json:"contains,omitempty"`
	DoesNotContain string `json:"does_not_contain,omitempty"`
	IsEmpty        bool   `json:"is_empty,omitempty"`
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
}

type FileCondition struct {
	IsEmpty    bool `json:"is_empty,omitempty"`
	IsNotEmpty bool `json:"is_not_empty,omitempty"`
}

type RelationCondition struct {
	Contains        string `json:"contains,omitempty"`
	DoesNotContains string `json:"does_not_contains,omitempty"`
	IsEmpty         string `json:"is_empty,omitempty"`
	IsNotEmpty      string `json:"is_not_empty,omitempty"`
}

// Use TextCondition for Text
// Use CheckboxCondition for Checkbox
// Use NumberCondition for Number
// Use DateCondition for Date
type FormulaCondition struct {
	Text     interface{} `json:"text,omitempty"`
	Checkbox interface{} `json:"checkbox,omitempty"`
	Number   interface{} `json:"number,omitempty"`
	Date     interface{} `json:"date,omitempty"`
}

type QueryDatabaseResponse struct {
	Object     string `json:"object,omitempty"`
	Results    []Page `json:"results,omitempty"`
	NextCursor string `json:"next_cursor,omitempty"`
	HasMore    bool   `json:"has_more,omitempty"`
}

// https://developers.notion.com/reference/post-database-query
func (d *DatabaseService) QueryDatabase(databaseID string,
	params *QueryDatabaseBodyParams) (*QueryDatabaseResponse, *http.Response, error) {

	response := new(QueryDatabaseResponse)
	apiError := new(APIError)
	resp, err := d.sling.New().Post(databaseID+"/query").BodyJSON(params).Receive(response, apiError)

	return response, resp, relevantError(err, *apiError)
}

type ListDatabasesQueryParams struct {
	StartCursor string `url:"start_cursor,omitempty"`
	PageSize    int32  `url:"page_size,omitempty"`
}

type ListDatabasesResponse struct {
	Results    []Database `json:"results,omitempty"`
	NextCursor string     `json:"next_cursor,omitempty"`
	HasMore    bool       `json:"has_more,omitempty"`
}

// https://developers.notion.com/reference/get-databases
func (d *DatabaseService) ListDatabases(params *ListDatabasesQueryParams) (*ListDatabasesResponse, *http.Response, error) {
	response := new(ListDatabasesResponse)
	apiError := new(APIError)
	resp, err := d.sling.New().Get("").QueryStruct(params).Receive(response, apiError)

	x, _ := d.sling.New().Get("").QueryStruct(params).Request()
	fmt.Print(x.URL)
	fmt.Println()

	return response, resp, relevantError(err, *apiError)
}