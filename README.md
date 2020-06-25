# mkstruct

## Excute
`$ go run mkstruct.go -n event_option -f "id event_id event_date_id name num price event_type limit place organize created_at updated_at"`

## Output
```
type EventOption struct {
	Id          string `json:"id"`
	EventId     string `json:"event_id"`
	EventDateId string `json:"event_date_id"`
	Name        string `json:"name"`
	Num         string `json:"num"`
	Price       string `json:"price"`
	EventType   string `json:"event_type"`
	Limit       string `json:"limit"`
	Place       string `json:"place"`
	Organize    string `json:"organize"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
```
