package facade

var DB = map[string]string{
	"a@a.com": "a",
	"b@b.com": "b",
}

type database struct {
}

func (db *database) getNameByMail(mail string) string {
	return DB[mail]
}

type mdWriter struct {
}

func (mw *mdWriter) title(title string) string {
	return "# Welcome to " + title + "'s page!"
}

type PageMaker struct {
}

func (pm *PageMaker) MakeWelcomePage(mail string) string {
	database := database{}
	writer := mdWriter{}

	name := database.getNameByMail(mail)
	page := writer.title(name)

	return page
}
