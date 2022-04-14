package model

type SearchTweetsModel struct {
	Data []struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
	Meta struct {
		NewestID    string `json:"newest_id"`
		OldestID    string `json:"oldest_id"`
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	} `json:"meta"`
}

type FollowingAccounts struct {
	Data []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		UserName string `json:"username"`
	} `json:"data"`
}
