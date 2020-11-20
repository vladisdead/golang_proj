package model

type Author struct {
	Id          int
	FirstName 	string `json:"firstName"`
	LastName 	string `json:"lastName"`
	MiddleName	string `json:"middleName"`
}

type Book struct {
	Id		int
	Tittle  string `json:"tittle"`
	Place   string `json:"place"`
	Edition string `json:"edition"`
	Year    string `json:"year"`
	NumberOfPage int `json:"numberOfPage"`
	AuthorId     int `json:"authorId"`
}

type Responsible struct {
	EditorID int `json:"editorId"`
	BookID int `json:"bookId"`
}

type Editor struct {
	Id			int
	FirstName 	string `json:"firstName"`
	LastName 	string `json:"lastName"`
	MiddleName	string `json:"middleName"`
}

type Translator struct {
	Id			int
	FirstName 	string `json:"firstName"`
	LastName 	string `json:"lastName"`
	MiddleName	string `json:"middleName"`
}

type Translation struct {
	TranslatorID int `json:"translatorId"`
	BookID int `json:"bookId"`
}

