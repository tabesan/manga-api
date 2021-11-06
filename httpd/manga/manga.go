package manga

type Manga struct {
	ID          int    `json:id`
	Title       string `json:title`
	ReleaseDate string `json:release_date`
}
