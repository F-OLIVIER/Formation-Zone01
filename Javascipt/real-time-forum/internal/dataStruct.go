package data

type DataTemplate struct {
	Data        map[string]string `json:"Data"`
	Msgerr      string            `json:"Msgerr"`
	Codeinit    bool              `json:"Codeinit"`
	Codeconfirm bool              `json:"Codeconfirm"`
}

type User struct {
	Photo            string   `json:"Photo"`
	Email            string   `json:"Email"`
	Username         string   `json:"Username"`
	Msgerr           string   `json:"Msgerr"`
	TabCategories    []string `json:"TabCategories"`
	StructCategories []StructCategorie
	Logged           bool   `json:"Logged"`
	Redirect         string `json:"Redirect"`
	Age              string `json:"Age"`
	Gender           string `json:"Gender"`
	FirstName        string `json:"FirstName"`
	LastName         string `json:"LastName"`
}
type StructCategorie struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

type Post struct {
	ID           int      `json:"ID"` // string
	Title        string   `json:"Title"`
	Date         string   `json:"Date"`
	Author       string   `json:"Author"`
	PhotoAuthor  string   `json:"PhotoAuthor"`
	Categorie    []string `json:"Categorie"`
	Content      string   `json:"Content"`
	ContentPhoto string   `json:"ContentPhoto"`
	ContentFile  string   `json:"ContentFile"`

	Comments []*Post
}
type GlobalData struct {
	UserData     User
	PostData     Post
	PostListData []*Post
	History      HistoryMsg
}

type GitData struct {
	Login     string `json:"login"`
	NodeID    string `json:"node_id"`
	AvatarURL string `json:"avatar_url"`
}
type GoogleUserResult struct {
	Id          string
	Name        string
	Given_name  string
	Family_name string
	Picture     string
}
type GoogleOauthToken struct {
	Access_token string
	Id_token     string
}

// Sert uniquement de stockages pour une insertion plus "propre" (selon Antoine) des données et éviter
// de se retrouver avec plein de vars dans les ().
type InsertC struct {
	FichB      bool
	ImageB     bool
	Titre      string
	Content    string
	User_id    string
	Image      string
	Fichier    string
	Categories []string
}

// ----------------------------------------------------------------------------
// -------------------- Struc de reception de donnée du js --------------------
// ----------------------------------------------------------------------------

type ViewPost struct {
	PostId string `json:"postId"`
}

type Register struct {
	// register et login
	Username string `json:"usernameValue"`
	Mail     string `json:"mailValue"`
	Pass     string `json:"passValue"`

	// register
	FirstNameregister string `json:"FirstNameregister"`
	LastNameregister  string `json:"LastNameregister"`
	Ageregister       string `json:"Ageregister"`
	Genderregister    string `json:"Genderregister"`
	ConfirmPass       string `json:"confirmPassregister"`
}

type Createpost struct {
	// pour post un recette
	NameRecette string   `json:"nameRecette"`
	Recette     string   `json:"recette"`
	Category    []string `json:"Category"`

	// pour post un commentaire d'un recette
	PostID  string `json:"postId"`
	Comment string `json:"comment"`
}

// ----------------------------------------------------------------------------
// --------------------- Struc pour le chat de discussion ---------------------
// ----------------------------------------------------------------------------

type ChatMessage struct {
	Username           string   `json:"username"`
	Horaire            string   `json:"horaire"`
	Message            string   `json:"message"`
	MajlistConnected   bool     `json:"majlistConnected"`
	UserPrivateMessage string   `json:"userPrivateMessage"`
	ListUser           []string `json:"listUser"`
	ListUserNewMsg     []bool   `json:"listUserNewMsg"`
	ListUserConnected  []string `json:"listUserConnected"`
	TypingProgress     bool     `json:"typingProgress"`
}

type ChatMsg struct {
	SenderMsg string `json:"senderMsg"`
	Timestamp string `json:"timestampMsg"`
	Content   string `json:"contentMsg"`
}

type HistoryMsg struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	ListMsg  []*ChatMsg
}
