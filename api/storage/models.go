package storage

type Validate interface {
	Validate() bool
}

type Post struct {
	ID          int    `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (p *Post) Validate() bool {
	if len(p.Title) != 0 && len(p.Description) != 0 && p.ID != 0 {
		return true
	}

	return false
}
