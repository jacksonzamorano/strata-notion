package definition

import (
	"github.com/jacksonzamorano/strata/component"
)

var Manifest = component.ComponentManifest{
	Name:    "notion",
	Version: "1.1.0",
}

var Setup = component.Define[SetupInput, SetupOutput](Manifest, "setup")
var FindPage = component.Define[FindPageInput, PagesOutput](Manifest, "find-page")
var QueryDataSource = component.Define[QueryDataSourceInput, PagesOutput](Manifest, "query-datasource")
var CreatePage = component.Define[CreatePageInput, CreatePageOutput](Manifest, "create-page")
var EditPage = component.Define[EditPageInput, CreatePageOutput](Manifest, "edit-page")
var Append = component.Define[AppendToBlockInput, struct{}](Manifest, "append")
var LaunchNotion = component.Define[struct{}, struct{}](Manifest, "launch")

type SetupInput struct {
	APIKey string
	Debug  bool
}
type SetupOutput struct {
	Ok bool
}

type FindPageInput struct {
	Query string
}
type PagesOutput struct {
	Items []NotionPage
}
type FindPageOutputItem struct {
	Id string
}

type QueryDataSourceInput struct {
	DataSourceId string
	Filter       map[string]any
}

type AppendToBlockInput struct {
	BlockId    string
	Components []any
}

type EditPageInput struct {
	PageId     string
	Properties map[string]any
}

type CreatePageInput struct {
	ParentDataSourceId string
	ParentDatabaseId   string
	ParentPageId       string
	Properties         map[string]any
	Markdown           string
	Children           []any
}

type CreatePageOutput struct {
	PageId string
}

// Notion types
type NotionPage struct {
	Id             string                    `json:"id"`
	CreatedDate    string                    `json:"created_time"`
	LastEditedTime string                    `json:"last_edited_time"`
	Url            string                    `json:"url"`
	Properties     map[string]NotionProperty `json:"properties"`
}
