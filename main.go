package main

import (
	"sync"

	d "github.com/jacksonzamorano/strata-notion/definition"
	"github.com/jacksonzamorano/strata/component"
)

var notionLock sync.RWMutex
var notionKey string
var notionUser NotionUserMeResponse
var notionDebug bool

func setup(in *component.ComponentInput[d.SetupInput, d.SetupOutput], container *component.ComponentContainer) *component.ComponentReturn[d.SetupOutput] {
	notionLock.Lock()
	notionKey = in.Body.APIKey
	notionDebug = in.Body.Debug
	notionLock.Unlock()
	res, err := notion[NotionUserMeResponse]("GET", "/v1/users/me", nil, container)
	if err != nil {
		container.Logger.Log("Error when authenticating with Notion: %s", err.Error())
		return in.Return(d.SetupOutput{
			Ok: false,
		})
	}
	notionLock.Lock()
	notionUser = res
	notionLock.Unlock()

	if notionUser.Bot != nil {
		container.Logger.Log("Logged in as bot '%s' in workspace '%s'", notionUser.Name, notionUser.Bot.WorkspaceName)
	} else {
		container.Logger.Log("Logged in as user '%s'.", notionUser.Name)
	}
	return in.Return(d.SetupOutput{
		Ok: true,
	})
}

func createPage(in *component.ComponentInput[d.CreatePageInput, d.CreatePageOutput], container *component.ComponentContainer) *component.ComponentReturn[d.CreatePageOutput] {
	var parent *NotionParent
	if len(in.Body.ParentDatabaseId) > 0 {
		parent = &NotionParent{
			DatabaseId: in.Body.ParentDatabaseId,
			Type:       "database_id",
		}
	} else if len(in.Body.ParentPageId) > 0 {
		parent = &NotionParent{
			PageId: in.Body.ParentPageId,
			Type:   "page_id",
		}
	} else if len(in.Body.ParentDataSourceId) > 0 {
		parent = &NotionParent{
			DataSourceId: in.Body.ParentDataSourceId,
			Type:         "data_source_id",
		}
	}

	pageRequest := NotionCreatePageRequest{
		Parent:     parent,
		Properties: in.Body.Properties,
		Markdown:   in.Body.Markdown,
		Children:   in.Body.Children,
	}
	page, err := notion[NotionCreatePageResponse]("POST", "/v1/pages", pageRequest, container)
	if err != nil {
		container.Logger.Log("Error when creating page: %s", err.Error())
		return in.Error("Error when creating page.")
	}
	return in.Return(d.CreatePageOutput{
		PageId: page.Id,
	})
}

func editPage(in *component.ComponentInput[d.EditPageInput, d.CreatePageOutput], container *component.ComponentContainer) *component.ComponentReturn[d.CreatePageOutput] {
	pageRequest := NotionEditPageRequest{
		Properties: in.Body.Properties,
	}
	page, err := notion[NotionCreatePageResponse]("PATCH", "/v1/pages/"+in.Body.PageId, pageRequest, container)
	if err != nil {
		container.Logger.Log("Error when editing page: %s", err.Error())
		return in.Error("Error when editing page.")
	}
	return in.Return(d.CreatePageOutput{
		PageId: page.Id,
	})
}

func getPage(in *component.ComponentInput[d.FindPageInput, d.PagesOutput], container *component.ComponentContainer) *component.ComponentReturn[d.PagesOutput] {
	req := NotionSearchRequest{
		Query: in.Body.Query,
	}
	pages, err := notion[NotionSearchResponse]("POST", "/v1/search", req, container)
	if err != nil {
		container.Logger.Log("Error when searching pages: %s", err.Error())
		return in.Error("Error when searching pages.")
	}
	return in.Return(d.PagesOutput{
		Items: pages.Results,
	})
}

func queryDataSource(in *component.ComponentInput[d.QueryDataSourceInput, d.PagesOutput], container *component.ComponentContainer) *component.ComponentReturn[d.PagesOutput] {
	filter := NotionFilterRequest{
		Filter: in.Body.Filter,
	}
	u := "/v1/data_sources/" + in.Body.DataSourceId + "/query"
	pages, err := notion[NotionSearchResponse]("POST", u, filter, container)
	if err != nil {
		container.Logger.Log("Error when querying pages: %s", err.Error())
		return in.Error("Error when querying pages.")
	}
	return in.Return(d.PagesOutput{
		Items: pages.Results,
	})
}

func appendBlock(in *component.ComponentInput[d.AppendToBlockInput, struct{}], container *component.ComponentContainer) *component.ComponentReturn[struct{}] {
	ap := NotionAppendRequest{
		Children: in.Body.Components,
	}
	u := "/v1/blocks/" + in.Body.BlockId + "/children"
	_, err := notion[NotionAppendResponse]("PATCH", u, ap, container)
	if err != nil {
		container.Logger.Log("Error when appending to page: %s", err.Error())
		return in.Error("Error when appending to page.")
	}
	return in.Return(struct{}{})
}

func launch(in *component.ComponentInput[struct{}, struct{}], container *component.ComponentContainer) *component.ComponentReturn[struct{}] {
	container.OpenUrl("https://notion.com")
	return in.Return(struct{}{})
}

func main() {
	component.CreateComponent(d.Manifest,
		component.Mount(d.Setup, setup),
		component.Mount(d.FindPage, getPage),
		component.Mount(d.QueryDataSource, queryDataSource),
		component.Mount(d.Append, appendBlock),
		component.Mount(d.CreatePage, createPage),
		component.Mount(d.EditPage, editPage),
		component.Mount(d.LaunchNotion, launch),
	).Start()
}
