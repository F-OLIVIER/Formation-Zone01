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
	Category        string `json:"category"`

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

type PostGroup struct {
	Id         int    `json:"id"`
	User_id    int    `json:"user_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Date       string `json:"date"`
	Likes      int    `json:"likes"`
	Group_id   int    `json:"group_id"`
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
	Id               int    `json:"id"`
	UserId_Following int    `json:"userid_following"`
	UserId_Follower  int    `json:"userid_follower"`
	DateFollow       string `json:"datefollow"`
	ValidateFollow   bool   `json:"validatefollow"`
	Action           string `json:"action"`
}

type Resp struct {
	Msg      string `json:"msg"`
	Receiver int    `json:"receiver"`
}

type Like struct {
	ID      int  `json:"id"`
	User_id int  `json:"user_id"`
	Post_Id int  `json:"post_id"`
	Like    bool `json:"liked"`
}

type Message struct {
	Id              int    `json:"id"`
	Sender_id       int    `json:"sender_id"`
	Sender_nickname string `json:"sender_nickname"`
	Receiver_id     int    `json:"receiver_id"`
	Content         string `json:"content"`
	Date            string `json:"date"`
	Msg_type        string `json:"msg_type"`
	User_id         int    `json:"user_id"`
	ImageData       string `json:"image_data"`
	Targets         []int  `json:"targets"`
}

// type GroupMessage struct {
// 	Id          int    `json:"id"`
// 	Group_id    int    `json:"group_id"`
// 	Sender_id   int    `json:"sender_id"`
// 	Receiver_id int    `json:"receiver_id"`
// 	Content     string `json:"content"`
// 	Date        string `json:"date"`
// 	Msg_type    string `json:"msg_type"`
// 	User_id     int    `json:"user_id"`
// 	ImageData   string `json:"image_data"`
// 	Targets     []int  `json:"targets"`
// }

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
	ID             int `json:"id"`
	IdWhoIsInvited int `json:"idwhoisinvited"`
	IdGroup        int
	Title          string
	AboutGroup     string
	UserID_Creator int
	MemberGroup    []User
	Event          EventGroup
	Image          []byte `json:"image"`
}
type EventGroup struct {
	IDEvent            int    `json:"id"`
	IDGroup            int    `json:"id_group"`
	UserIDCreatorEvent int    `json:"id_user"`
	Title              string `json:"title"`
	Date               string `json:"date"`
	Description        string `json:"description"`
	Coming             int    `json:"coming"`
	NotComing          int    `json:"notcoming"`
}

type InviteInTheGroup struct {
	ID int `json:"id"`

	NameOfGroup     string `json:"nameOfGroup"`
	NameOfThePerson string `json:"nameOfThePerson"`
}

type Coming struct {
	ID       int  `json:"id"`
	User_id  int  `json:"user_id"`
	Event_id int  `json:"event_id"`
	Come     bool `json:"Come"`
	NotCome  bool `json:"NotCome"`
}
