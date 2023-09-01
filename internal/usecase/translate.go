package usecase

import (
	"log/slog"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Translator struct {
	bundle *i18n.Bundle
}

func NewTranslator() *Translator {
	bundle := i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err := bundle.LoadMessageFile("active.en.toml")
	if err != nil {
		slog.Error("error loading language", "err", err)
	}

	_, err = bundle.LoadMessageFile("active.de.toml")
	if err != nil {
		slog.Error("error loading language", "err", err)
	}

	slog.Info("loaded all languages")

	return &Translator{
		bundle: bundle,
	}
}

func (t *Translator) Translate(msg, lang string) string {
	localizer := i18n.NewLocalizer(t.bundle, lang)

	val := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: msg,
	})

	return val
}

type WebsiteData struct {
	Projects           []Project
	CompanyName        string
	CompanyDescription string
	ThemeSwitchText    string
	ThemeSwitchAuto  string
	ThemeSwitchLight string
	ThemeSwitchBlack string
	BlogLinkText       string
	ImprintText        string
	ContactFormData    ContactFormData
}

type ContactFormData struct {
	Header      string
	Description string
	Firstname   string
	Lastname    string
	Email       string
	Message     string
	SendText    string
}

func (t *Translator) RetrieveWebsiteData(lang string) *WebsiteData {
	return &WebsiteData{
		CompanyName:        t.Translate("CompanyName", lang),
		CompanyDescription: t.Translate("CompanyDescription", lang),
		ThemeSwitchText:    t.Translate("ThemeSwitchText", lang),
        ThemeSwitchAuto:  t.Translate("ThemeSwitchAuto", lang),
        ThemeSwitchLight: t.Translate("ThemeSwitchLight", lang),
        ThemeSwitchBlack: t.Translate("ThemeSwitchBlack", lang),
		BlogLinkText:       t.Translate("BlogLinkText", lang),
		ImprintText:        t.Translate("ImprintText", lang),
		ContactFormData: ContactFormData{
			Header:      t.Translate("ContactFormHeader", lang),
			Description: t.Translate("ContactFormDescription", lang),
			Firstname:   t.Translate("ContactFormFirstname", lang),
			Lastname:    t.Translate("ContactFormLastname", lang),
			Email:       t.Translate("ContactFormEmail", lang),
			Message:     t.Translate("ContactFormMessage", lang),
			SendText:    t.Translate("ContactFormSendText", lang),
		},
	}
}

func (t *Translator) RetrieveWebsiteDataWithProjects(lang string, projects []Project) *WebsiteData {
	translatedProjects := make([]Project, len(projects))

	// translate projects
	for i := range projects {
		title := t.Translate(projects[i].Title, lang)
		description := t.Translate(projects[i].Description, lang)
		imageCaption := t.Translate(projects[i].ImageCaption, lang)

		translatedProject := Project{
			Title:        title,
			Description:  description,
			ImagePath:    projects[i].ImagePath,
			ImageCaption: imageCaption,
		}

		translatedProjects[i] = translatedProject
	}

    res := t.RetrieveWebsiteData(lang)
    res.Projects = translatedProjects
    return res
}
