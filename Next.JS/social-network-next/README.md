# SOCIAL NETWORK

<div align="center" markdown>

# Readme language
🌏
English | 
[**Français**](https://github.com/F-OLIVIER/Formation-Zone01/blob/main/Next.JS/social-network-next/README_FR.md)

</div>

## 📝 Description

The aim of the project is to recreate a social network like Facebook or Linkedin.

The project is split into 2 parts, a server in <a href="https://go.dev">golang</a> and a front end in <a href="https://nextjs.org">Next.js</a>.

<li>An <a href="https://github.com/golang-migrate/migrate">SQL database migration system</a></li>
<li>A user authentication system</li>
<li>Creating and managing posts</li>
<li>Post display management according to profile type (public, private, semi-private)</li>
<li>A private chat room with emojis</li>
<li>User group management</li>
<li>Event management via groups</li>
<li>Follow-up management for other members (followers & followings)</li>
<li>Notification management (online and offline)</li>

Project image
<table align= "center" width="95%">
    <tbody>
        <tr>
            <td><img src="./readme_img/notconnected.png"></td>
            <td><img src="./readme_img/post.png"></td>
        </tr>
        <tr>
            <td><img src="./readme_img/group.png"></td>
            <td><img src="./readme_img/chat.png"></td>
        </tr>
        <tr>
            <td><img src="./readme_img/notif.png"></td>
            <td><img src="./readme_img/navbar.png"></td>
        </tr>
    </tbody>
</table>


An optional cross platform application has been created and can be found in the folder

___
## ⚙️ Installation & usage

The project run via <a href="https://www.docker.com">Docker</a> with a `Docker Compose` which allows the 2 programs to be launched (the SQL database being in the backend), so make sure you <a href="https://docs.docker.com/compose/install/">install Docker</a> before running the site launch command.

```sh
sh ./launch.sh
```

To launch manually, use the following commands:
```sh
# console 1 (serveur golang)
cd back
go run .
```
```sh
# console 2 (front)
cd front
npm i
npm run dev
```
___
## 🔗 Dependencies

**Golang server** <br>
The server uses `version 22` of [golang](https://go.dev) and the following libraries:<br>
- [go-sqlite3](https://github.com/mattn/go-sqlite3)
- [uuid](https://github.com/gofrs/uuid)
- [gorilla/websocket](https://pkg.go.dev/github.com/gorilla/websocket)

**Le front** <br>
The frontend uses version 14 of [Next.js](https://nextjs.org), `html` and `css` and the following [npm](https://www.npmjs.com) dependencies:<br>
- [@emotion/react](https://www.npmjs.com/package/@emotion/react)
- [@emotion/styled](https://www.npmjs.com/package/@emotion/styled)
- [@mui/icons-material](https://www.npmjs.com/package/@mui/icons-material)
- [@mui/material](https://www.npmjs.com/package/@mui/material)
- [emoji-mart](https://www.npmjs.com/package/emoji-mart)
- [emoji-picker-react](https://www.npmjs.com/package/emoji-picker-react)
- [next.js](https://www.npmjs.com/package/next)
- [react](https://www.npmjs.com/package/react)
- [react-datepicker](https://www.npmjs.com/package/react-datepicker)
- [react-dom](https://www.npmjs.com/package/react-dom)
- [react-hot-toast](https://www.npmjs.com/package/react-hot-toast)
- [react-router-dom](https://www.npmjs.com/package/react-router-dom)
- [zxcvbn](https://www.npmjs.com/package/zxcvbn)
___
## 🧑‍💻 Authors

+ Enzo FEMENIA
+ Fabien OLIVIER
+ Maxime DUBOIS
+ Mathieu CATELAIN
+ Romain CLATOT
