package usecase

type Project struct {
	Title       string
	Description string
	ImagePath   string
}

func ListProjects() []Project {
	projects := []Project{
		{
			Title:       "BDR Projekt",
			Description: "insert text here",
			ImagePath:   "static/bdr-image.jpeg",
		},
		{
			Title:       "Data Science Projekt",
			Description: "NLP Analyse von Artikeldaten",
			ImagePath:   "static/cait-image.jpg",
		},
		{
			Title:       "Skyblock AH History",
			Description: "Ein Videospiel, wow",
			ImagePath:   "static/sky-image.jpg",
		},
	}

	return projects
}
