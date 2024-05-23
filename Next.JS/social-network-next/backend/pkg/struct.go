package pkg

import "net/http"

var Sessions = map[string]SessionUser{}

type SessionUser struct {
	email  string
	Cookie *http.Cookie
}

type User struct { //Sert a Register et le profil
	Id              int    `json:"id"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	DateOfBirth     string `json:"dateofbirth"`
	Avatar          []byte `json:"avatar"` // A revoir
	Nickname        string `json:"nickname"`
	AboutMe         string `json:"aboutme"`
	PrivateProfile  string `json:"privateprofile"`
	PointOfInterest string `json:"pointofinterest"` //Join variables

	ListFollowings          []User `json:"listfollowings"`
	ListFollowers           []User `json:"listfollowers"`
	ListFollowersToValidate []User `json:"ListFollowersToValidate"`
}
type Session struct {
	Session_uuid string
	User_id      int
}
type Post struct {
	Id         int    `json:"id"`
	User_id    int    `json:"user_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Date       string `json:"date"`
	Likes      int    `json:"likes"`
	Private    int    `json:"privacy"`
	Image      []byte `json:"image"`
	NbComments int    `json:"nbcomments"`
}
type Comment struct {
	Id      int    `json:"id"`
	Post_id int    `json:"post_id"`
	User_id int    `json:"user_id"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Image   []byte `json:"image"`
}
type Login struct {
	Data     string `json:"data"`
	Password string `json:"password"`
}
type Followers struct {
	UserId_Following int    `json:"userid_following"`
	UserId_Follower  int    `json:"userid_follower"`
	DateFollow       string `json:"datefollow"`
	ValidateFollow   bool   `json:"validatefollow"`
	Action           string `json:"action"`
}

type Resp struct {
	Msg string `json:"msg"`
}

type Like struct {
	ID      int  `json:"id"`
	User_id int  `json:"user_id"`
	Post_Id int  `json:"post_id"`
	Like    bool `json:"liked"`
}

type Message struct {
	Id          int    `json:"id"`
	Sender_id   int    `json:"sender_id"`
	Receiver_id int    `json:"receiver_id"`
	Content     string `json:"content"`
	Date        string `json:"date"`
	Msg_type    string `json:"msg_type"`
	User_id     int    `json:"user_id"`
	ImageData   string `json:"image_data"`
}

type Chat struct {
	User_one int
	User_two int
	Time     int
}

type OnlineUsers struct {
	UserIds  []int  `json:"user_ids"`
	Msg_type string `json:"msg_type"`
}

type Group struct {
	IdGroup        int
	NameGroup      string
	Description    string
	UserID_Creator int
	MemberGroup    []User
	Event          EventGroup
	Image          []byte `json:"image"`
}
type EventGroup struct {
	UserIDCreatorEvent int
	Title              string
	Date               string
	Description        string
}

type InviteInTheGroup struct {
	NameOfGroup     string `json:"nameOfGroup"`
	NameOfThePerson string `json:"nameOfThePerson"`
}
