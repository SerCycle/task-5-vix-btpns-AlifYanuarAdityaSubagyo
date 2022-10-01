package photo

type PhotoResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photourl"`
	UserID   int    `json:"userid"`
}
