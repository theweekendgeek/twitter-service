package data

type Profile struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type Tweet struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type Meta struct {
	NextToken   string `json:"next_token"`
	ResultCount int    `json:"result_count"`
}

type ProfileResponse struct {
	Data Profile `json:"data"`
}

type Tweets []Tweet

type TimelineResponse struct {
	Tweets   Tweets `json:"data"`
	MetaData Meta   `json:"meta"`
}
