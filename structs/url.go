package structs

type AddUrlRequest struct {
	Short string `json:"short" binding:"required"`
	Url   string `json:"url" binding:"required"`
}
