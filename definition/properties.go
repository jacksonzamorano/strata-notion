package definition

import "time"

type NotionProperty struct {
	Title    *[]NotionRichText     `json:"title,omitempty"`
	RichText *[]NotionRichText     `json:"rich_text,omitempty"`
	Email    string                `json:"email,omitempty"`
	Checkbox *bool                 `json:"checkbox,omitempty"`
	Date     *NotionDateProperty   `json:"date,omitempty"`
	Select   *NotionSelectProperty `json:"select,omitempty"`
	Number   *int                  `json:"number,omitempty"`
}

func PropertyNumber(n int) *NotionProperty {
	return &NotionProperty{
		Number: &n,
	}
}

type NotionSelectProperty struct {
	Name string `json:"name"`
}

func PropertySelect(opt string) *NotionProperty {
	return &NotionProperty{
		Select: &NotionSelectProperty{
			Name: opt,
		},
	}
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
