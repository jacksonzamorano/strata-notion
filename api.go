package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/jacksonzamorano/strata-notion/definition"
	"github.com/jacksonzamorano/strata/component"
)

var client http.Client

type NotionUserMeResponse struct {
	Id   string                   `json:"id"`
	Name string                   `json:"name"`
	Type string                   `json:"type"`
	Bot  *NotionUserMeResponseBot `json:"bot"`
}

type NotionUserMeResponseBot struct {
	WorkspaceName string `json:"workspace_name"`
}

type NotionParent struct {
	PageId       string `json:"page_id,omitempty"`
	DataSourceId string `json:"data_source_id,omitempty"`
	DatabaseId   string `json:"database_id,omitempty"`
	Type         string `json:"type"`
}
type NotionEditPageRequest struct {
	Properties map[string]any `json:"properties"`
}
type NotionCreatePageRequest struct {
	Parent     *NotionParent  `json:"parent,omitempty"`
	Properties map[string]any `json:"properties"`
	Markdown   string         `json:"markdown,omitempty"`
	Children   []any          `json:"children,omitempty"`
}
type NotionCreatePageResponse struct {
	Id string `json:"id"`
}

type NotionSearchRequest struct {
	Query string `json:"query"`
}
type NotionSearchResponse struct {
	Results []definition.NotionPage `json:"results"`
}
type NotionFilterRequest struct {
	Filter map[string]any `json:"filter"`
}

type NotionAppendRequest struct {
	Children []any `json:"children"`
}
type NotionAppendResponse struct {
	Results []any `json:"results"`
}

func notion[T any](method, url string, body any, container *component.ComponentContainer) (T, error) {
	var output T
	var sendBody io.ReadCloser
	if body != nil {
		d, _ := json.Marshal(body)
		if notionDebug {
			container.Logger.Log("Sent to Notion: '%s'", string(d))
		}
		sendBody = io.NopCloser(bytes.NewReader(d))
	}

	req, err := http.NewRequest(method, "https://api.notion.com"+url, sendBody)
	if err != nil {
		return output, err
	}
	req.Header.Add("Authorization", notionKey)
	req.Header.Add("Notion-Version", "2025-09-03")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return output, err
	}
	defer res.Body.Close()
	rb, _ := io.ReadAll(res.Body)
	if notionDebug {
		container.Logger.Log("Received from Notion: '%s'", string(rb))
	}

	err = json.Unmarshal(rb, &output)
	if err != nil {
		return output, err
	}
	return output, nil
}
