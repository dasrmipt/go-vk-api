package obj

type Chat struct {
	ID           int64             `json:"id"`
	Type         string            `json:"type"`
	Title        string            `json:"title"`
	AdminId      int64             `json:"admin_id"`
	Users        []int64           `json:"users"`
	PushSettings map[string]string `json:"push_settings"`
	Photo50      string            `json:"photo_50"`
	Photo100     string            `json:"photo_100"`
	Photo200     string            `json:"photo_200"`
	Left         int               `json:"left"`
	Kicked       int               `json:"kicked"`
}
