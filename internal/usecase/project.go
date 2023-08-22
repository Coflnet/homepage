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
			Description: "Wir haben in einem Forschungsprojekt für die Regierung ein neues System entwickelt, welches die Kommunikation zwischen den Bürgern und der Regierung vereinfachen soll.",
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
