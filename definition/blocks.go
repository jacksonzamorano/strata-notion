package definition

type NotionBlock struct {
	Type      string           `json:"type"`
	Paragraph *NotionParagraph `json:"paragraph,omitempty"`
	Callout   *NotionCallout   `json:"callout,omitempty"`
}

func BlockParagraph(rt ...NotionRichText) *NotionBlock {
	return &NotionBlock{
		Type: "paragraph",
		Paragraph: &NotionParagraph{
			RichText: rt,
		},
	}
}

type IconNotion struct {
	Emoji string `json:"emoji,omitempty"`
}

type NotionCallout struct {
	RichText []NotionRichText `json:"rich_text"`
	Icon     *IconNotion      `json:"icon,omitempty"`
	Color    string           `json:"color"`
}

func BlockCallout(color string, emoji string, rt ...NotionRichText) *NotionBlock {
	return &NotionBlock{
		Type: "callout",
		Callout: &NotionCallout{
			RichText: rt,
			Icon: &IconNotion{
				Emoji: emoji,
			},
			Color: color,
		},
	}
}
