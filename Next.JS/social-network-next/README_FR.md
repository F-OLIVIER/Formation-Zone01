# SOCIAL NETWORK

<div align="center" markdown>

# Langue du Readme
🌏
[**English**](https://github.com/F-OLIVIER/Formation-Zone01/blob/main/Next.JS/social-network-next/README.md) | 
Français

</div>

## 📝 Descriptif

Le but du projet est de recréer un réseau social tels que Facebook ou Linkedin.

Le projets est décomposé en 2 parties, un serveur en <a href="https://go.dev">golang</a> et un front en <a href="https://nextjs.org">Next.js</a>.

<li>Un système de <a href="https://github.com/golang-migrate/migrate">migration de la base de données SQL</a></li>
<li>Un système d'authentification des utilisateurs</li>
<li>Création est gestion des posts</li>
<li>Gestion d'affichage des posts en fonction du type de profil (public, privée, semi-privée)</li>
<li>Un chat de discussion privée avec des émojis</li>
<li>Gestion de groupes d'utilisateurs</li>
<li>Gestion d'événement via les groupes</li>
<li>Gestion de suivi d'autres membres (follower & followings)</li>
<li>Gestion des notifications (en ligne et hors ligne)</li>

Image du projet
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


Une application cross platform optionnelle a été réalisée et se trouve dans le dossier `/cross-platform`.

___
## ⚙️ Installation & usage

Le projet est exécuté via <a href="https://www.docker.com">Docker</a> avec un `Docker Compose` qui permet de lancer les 2 programmes (la base de données SQL étant dans le backend), assurez-vous <a href="https://docs.docker.com/compose/install/">d'installer Docker</a> avant d'exécuter la commande de lancement du site.

```sh
sh ./launch.sh
```

Pour effectuer un lancement manuel, voici les commandes :
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
## 🔗 Dépendences

**Le serveur golang** <br>
Le serveur utilise la `version 22` de [golang](https://go.dev) et les librairies suivantes :<br>
- [go-sqlite3](https://github.com/mattn/go-sqlite3)
- [uuid](https://github.com/gofrs/uuid)
- [gorilla/websocket](https://pkg.go.dev/github.com/gorilla/websocket)

**Le front** <br>
Le front utilise la `version 14` de [Next.js](https://nextjs.org), `html` et `css` et les dépendances [npm](https://www.npmjs.com) suivantes :<br>
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
