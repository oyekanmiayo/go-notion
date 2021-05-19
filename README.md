TODO:

* **Talk about interfaces and type-specific objects etc.**

# go-notion

go-notion is a minimal Go client library for [Notion's v1 API](https://developers.notion.com/). Check the [usage] or
examples to see how to access Notion's v1 API.

***NB**: Notion's v1 API is still in beta. This integration may change as they update their endpoints*

## Table of Contents

* [Installation]()
* [Authentication]()
* [Usage]()
    * [Databases]()
    * [Pages]()
    * [Blocks]()
    * [Users]()
    * [Search]()
* [Roadmap]()
* [Contributing]()
* [License]()

## Installation

```
go get github.com/oyekanmiayo/go-notion/notion/version1
```

## Authentication

HTTP requests to Notion's API must contain a bearer token in the `Authorization` header to be successful. Currently,
Notion allows two types of integrations: Internal and Public. Getting the bearer token is different for each case.

### Internal Integration

When you create an internal integration on Notion's developer portal, you are given something called an "Internal
Integration Token" - this is the bearer token. All you need to do is simply copy it out and store it as an environment
variable (preferred) or use it directly in your code.

Show image here

NB: Internal Integration Token, Bearer Token and Access Token mean the same things here.

```go
import (
	notion "github.com/oyekanmiayo/go-notion/notion/version1"
	"net/http"
	"os"
) 

func main() {
	accessToken := os.GETENV("NOTION_BEARER_TOKEN")
	
	// Notion Client
	client := notion.NewClient(http.DefaultClient, accessToken)
}
```

### Public Integration

This integration receives bearer tokens each time a user completes the [OAuth flow]().

## Usage

### Databases

Read more about the Database endpoints [here](https://developers.notion.com/reference/database).

#### Retrieve a database

This retrieves a Notion database based on a specified ID.

```go
client := notion.NewClient(http.DefaultClient, *accessToken)

// Retrieve DB
db, _, err := client.Databases.RetrieveDatabase(databaseID)
if err != nil {
    fmt.Printf("Err %v\n", err)
}
```

See full code example [here]().

#### Query a database

This gets a list of [Pages]() contained in a database, filtered and ordered according to the filter conditions and sort
criteria provided in the request. Filters can [single filters]() or [compound filters]() - read more [here]().

```go
client := notion.NewClient(http.DefaultClient, *accessToken)

// Query DB with SingleFilter
params := &notion.QueryDatabaseBodyParams{
    Filter: &notion.SingleFilter{
        Property: "Tags",
        MultiSelect: &notion.MultiSelectCondition{
            Contains: "Tag1",
        },
    },
}
resp, _, err := client.Databases.QueryDatabase(*databaseID, params)
if err != nil {
    fmt.Printf("Err %v\n", err)
}
```

See full code example [here]() - it also contains a compound filter example :)

#### List databases

More than one database can be shared with a Notion integration. This endpoint lists all the databases shared with an
authenticated integration.

```go
client := notion.NewClient(http.DefaultClient, *accessToken)

// List all DBs in a workspace
params := &notion.ListDatabasesQueryParams{
    PageSize: 20,
}
resp, _, err := client.Databases.ListDatabases(params)
if err != nil {
    fmt.Printf("Err %v\n", err)
}
```

See full code example [here]().

### Pages

Read more about Page endpoints [here](https://developers.notion.com/reference/page).

#### Retrieve a page

This retrieves a Notion page based on a specified ID.

```go
client := notion.NewClient(http.DefaultClient, *accessToken)

// Retrieve a page using its pageID
db, _, err := client.Pages.RetrievePage(*pageID)
if err != nil {
    fmt.Printf("Err %v\n", err)
}
```

See full code example [here]().

#### Create a page

Pages in Notion can be created within a database or within another page. This endpoint creates a page as a child of the
parent (database or page) specified.

```go
client := notion.NewClient(http.DefaultClient, *accessToken)

// Create a page in a database
params := &notion.CreatePageBodyParams{
    Parent: &notion.DatabaseParent{
        DatabaseID: *databaseID,
    },
    Properties: map[string]interface{}{
        "Name": &notion.PageProperty{
            Title: []notion.RichText{
                {
                    Text: &notion.Text{
                        Content: "Creating Page Sample",
                    },
                },
            },
        },
        "Tags": &notion.PageProperty{
            MultiSelect: []notion.MultiSelectPropertyOptions{
                {
                    Name: "Tag1",
                },
                {
                    Name: "Tag3",
                },
            },
        },
        "Recommended": &notion.PageProperty{
            Checkbox: true,
        },
    },
}
resp, _, err := client.Pages.CreatePage(params)
if err != nil {
    fmt.Printf("Err %v\n", err)
}
```

The example above creates a page within a database. The body of the request is determined by the structure of the
database. See full code example [here]() along with an example to create a page within another page.

#### Update page properties

Updates page property values for the specified page. Properties that are not set via the "properties" parameter will
remain unchanged.

```go
client := notion.NewClient(http.DefaultClient, *accessToken)

// Update the title of a page
params := &notion.UpdatePagePropertiesBodyParams{
    Properties: map[string]map[string]interface{}{
        "Name": {
            "title": []notion.RichText{
                {
                    Type: "text",
                    Text: &notion.Text{
                        Content: "Jamaican Cuisines II",
                    },
                },
            },
        },
        "Recommended": {
            "checkbox": false,
        },
    },
}
resp, _, err := client.Pages.UpdatePageProperties(*pageID, params)
if err != nil {
    fmt.Printf("Err %v\n", err)
}
```

See full code example [here]().

### Blocks

A block object represents content within Notion. Blocks can be text, lists, media, and more. A page is a type of block,
too! Read more about Block endpoints [here](https://developers.notion.com/reference/block).

#### Retrieve block children

Returns a paginated array of child block objects contained in the block using the ID specified. In order to receive a
complete representation of a block, you may need to recursively retrieve the block children of child blocks.

```go
client := notion.NewClient(http.DefaultClient, *accessToken)

// Retrieve the block children for a block
params := &notion.RetrieveBlockChildrenParams{}
db, _, err := client.Blocks.RetrieveBlockChildren(*blockID, params)
if err != nil {
    fmt.Printf("Err %v\n", err)
}
```

See full code example [here]().

#### Append block children

Creates and appends new children blocks to the block using the ID specified. Returns the Block object which contains the
new children.

```go
client := notion.NewClient(http.DefaultClient, *accessToken)

// Append block children (Header Two & Paragraph) to a block
params := &notion.AppendBlockChildrenBodyParams{
    Children: []notion.Block{
        {
            Object: "block",
            Type:   "heading_2",
            HeadingTwo: &notion.HeadingTwo{
                Text: []notion.RichText{
                    {
                        Type: "text",
                        Text: &notion.Text{
                            Content: "Header Two Test",
                        },
                    },
                },
            },
        },
        {
            Object: "block",
            Type:   "paragraph",
            Paragraph: &notion.Paragraph{
                Text: []notion.RichText{
                    {
                        Type: "text",
                        Text: &notion.Text{
                            Content: "Paragraph Test",
                        },
                    },
                },
            },
        },
    },
}

db, _, err := client.Blocks.AppendBlockChildren(*blockID, params)
if err != nil {
    fmt.Printf("Err %v\n", err)
}
```

See full code example [here]().

### Users

The User object represents a user in a Notion workspace. Users include guests, full workspace members, and bots. Read
more about User endpoints [here](https://developers.notion.com/reference/user).

### Retrieve a user

This retrieves a Notion user based on a specified ID.

```go
client := notion.NewClient(http.DefaultClient, *accessToken)

// Retrieve a user by userID
db, _, err := client.Users.RetrieveUser(*userID)
if err != nil {
    fmt.Printf("Err %v\n", err)
}
```

See full code example [here]().

### List all users

Returns a paginated list of users for the workspace.

```go
client := notion.NewClient(http.DefaultClient, *accessToken)

// List all users in workspace
params := &notion.ListUsersQueryParams{
    PageSize: 20,
}
db, _, err := client.Users.ListUsers(params)
if err != nil {
    fmt.Printf("Err %v\n", err)
}
```

See full code example [here]().

### Search

Searches all pages and child pages that are shared with the integration and returns a list of databases and pages which
have titles that contain the `query` parameter. Other parameters like `sort` and `filter` also affect the output. Read
more about it [here](https://developers.notion.com/reference/post-search)

```go
client := notion.NewClient(http.DefaultClient, *accessToken)

// Search the workspace and return pages that have titles containing "Yurts"
// The result should be sorted in descending order of "last_edited_time"
params := &notion.SearchBodyParams{
    Query: "Yurts",
    Sort: &notion.Sort{
        Direction: "descending",
        Timestamp: "last_edited_time",
    },
    Filter: &notion.SearchFilter{
        Value:    "page",
        Property: "object",
    },
}
db, _, err := client.Search.Search(params)
if err != nil {
    fmt.Printf("Err %v\n", err)
}
```

See full code example [here]().

### Contributing

Contributions won't be accepted until Notion's v1 API is out of beta. This may change in the future, if needed.

### License

[Apache 2.0 License]()