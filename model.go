package ytdl

type Output struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Formats []Format
}

type Format struct {
	Url         string `json:"url"`
	Format_note string `format_note:"format_note"`
}
