package definition

type NotionParagraph struct {
	RichText []NotionRichText `json:"rich_text"`
}

func Paragraph(rt ...NotionRichText) NotionBlock {
	return NotionBlock{
		Type: "paragraph",
		Paragraph: &NotionParagraph{
			RichText: rt,
		},
	}
}

type NotionRichText struct {
	Type        string                   `json:"type"`
	Text        *NotionRichTextText      `json:"text,omitempty"`
	Annotations NotionRichTextAnnotation `json:"annotations"`
}

func PlainText(text string) NotionRichText {
	return NotionRichText{
		Type: "text",
		Text: &NotionRichTextText{
			Content: text,
		},
		Annotations: RichTextDefaultAnnotation(),
	}
}

func CodeText(text string) NotionRichText {
	return NotionRichText{
		Type: "text",
		Text: &NotionRichTextText{
			Content: text,
		},
		Annotations: NotionRichTextAnnotation{
			Code:  true,
			Color: "default",
		},
	}
}

type NotionRichTextText struct {
	Content string  `json:"content"`
	Link    *string `json:"link"`
}
type NotionRichTextAnnotation struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

func RichTextDefaultAnnotation() NotionRichTextAnnotation {
	return NotionRichTextAnnotation{
		Color: "default",
	}
}
