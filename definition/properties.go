package definition

import "time"

type NotionProperty struct {
	Title    *[]NotionRichText   `json:"title,omitempty"`
	RichText *[]NotionRichText   `json:"rich_text,omitempty"`
	Email    string              `json:"email,omitempty"`
	Checkbox *bool               `json:"checkbox,omitempty"`
	Date     *NotionDateProperty `json:"date,omitempty"`
}
type NotionDateProperty struct {
	Start time.Time  `json:"start"`
	End   *time.Time `json:"end,omitempty"`
}

func PropertyTitle(rt ...NotionRichText) *NotionProperty {
	return &NotionProperty{
		Title: &rt,
	}
}

func PropertyRichText(rt ...NotionRichText) *NotionProperty {
	return &NotionProperty{
		RichText: &rt,
	}
}

func PropertyEmail(email string) *NotionProperty {
	return &NotionProperty{
		Email: email,
	}
}

func PropertyCheckbox(check bool) *NotionProperty {
	return &NotionProperty{
		Checkbox: &check,
	}
}

func PropertyDate(d time.Time) *NotionProperty {
	return &NotionProperty{
		Date: &NotionDateProperty{
			Start: d,
		},
	}
}

func PropertyDateRange(s time.Time, end time.Time) *NotionProperty {
	return &NotionProperty{
		Date: &NotionDateProperty{
			Start: s,
			End:   &end,
		},
	}
}
