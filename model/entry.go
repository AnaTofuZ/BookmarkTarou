package model

// Entry ブックマーク対象のURLなどの情報があるモデル
type Entry struct {
	ID    uint64 `db:"id"`
	URL   string `db:"url"`
	Title string `db:titile"`
}
