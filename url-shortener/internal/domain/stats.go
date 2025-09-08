package domain

type LinkStats struct {
	Link        `json:",inline"`
	AccessCount int `json:"access_count"`
}

func NewLinkStats(link Link, accessCount int) *LinkStats {
	return &LinkStats{Link: link, AccessCount: accessCount}
}
