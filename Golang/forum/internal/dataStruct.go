package data

type DataTemplate struct {
	Data        map[string]string
	Msgerr      string
	Codeinit    bool
	Codeconfirm bool
}

type User struct {
	Photo            string
	Email            string
	Username         string
	Msgerr           string
	TabCategories    []string
	StructCategories []StructCategorie
	Logged           bool
}
type StructCategorie struct {
	Id   string
	Name string
}

type Post struct {
	ID           int // string
	Title        string
	Date         string
	Author       string
	PhotoAuthor  string
	Categorie    []string
	Content      string
	ContentPhoto string
	ContentFile  string

	Comments []*Post

	LikeCounter    []int // pour la home
	DislikeCounter []int
	Likes          int // pour le viewPost
	Dislikes       int
}
type GlobalData struct {
	UserData     User
	PostData     Post
	PostListData []*Post
}

// Sert uniquement de stockages pour une insertion plus "propre" (selon Antoine, bisous je t'aime) des données et éviter
// de se retrouver avec plein de vars dans les ().
type InsertC struct {
	FichB   bool
	ImageB  bool
	Titre   string
	Content string
	// Username   string
	User_id    string
	Image      string
	Fichier    string
	Categories []string
}
