package usecase

type Project struct {
	Title        string `toml:"title"`
	Description  string `toml:"description"`
	ImagePath    string `toml:"image_path"`
	ImageCaption string `toml:"image_caption"`
}
