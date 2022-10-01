package photo

type AddPhoto struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption" binding:"required"`
	PhotoUrl string `json:"photourl" binding:"required"`
	UserID   int    `json:"userid" binding:"required"`
}
