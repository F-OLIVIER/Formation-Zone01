import { commonBlock } from "./useful.js";
import { cookieName, userNotConnected } from "./main.js";
import { home } from "./home.js";

export function viewPost(postId) {
    // Check la precence du cookie, si il n'existe pas ou si le fragment est absent redirection automatique vers "/"
    if (!document.cookie.split(";").some((item) => item.trim().startsWith(cookieName + "=")) || postId === "") {
        userNotConnected();
        return;
    }

    const dataToSend = { postId };
    console.log("dataToSend : ", dataToSend)

    fetch('http://localhost:8080/api/ViewPost', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(dataToSend),
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`Erreur de réseau: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            if (typeof data === 'object') {
                console.log('Data received (view post):', data);
                if (!data.UserData.Logged) { // si le cookie n'ai pas valide, suppression + redirection
                    document.cookie = cookieName + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
                    userNotConnected();
                } else {
                    containerViewPost(data, postId);
                }
            } else {
                throw new Error('Réponse invalide du serveur (non-JSON)');
            }
        })
        .catch(error => {
            console.error('Erreur lors de la récupération des données:', error);
        });
}

function containerViewPost(data, postId) {
    commonBlock(data);

    let Container = document.getElementById('Container');

    // ------------------------------------------------------------------
    // ------------------------- Contenu du post ------------------------
    // ------------------------------------------------------------------
    let post = document.createElement('div');
    post.className = 'post';

    // div de l'utilisateur qui a crée la post
    let user = document.createElement('div');
    user.className = 'user';
    let userImg = document.createElement('img');
    userImg.src = data.PostData.PhotoAuthor;
    userImg.textContent = 'Auteur : ' + data.PostData.Author;
    user.appendChild(userImg);
    post.appendChild(user);

    // div de contenu du post
    let contenu = document.createElement('div');
    contenu.className = 'contenu';

    // titre du post
    let title = document.createElement('div');
    title.className = 'title';
    title.textContent = data.PostData.Title;
    contenu.appendChild(title);

    // message du post
    let message = document.createElement('div');
    message.className = 'message';
    message.innerHTML = data.PostData.Content;
    if (data.PostData.ContentPhoto != "") {
        let postPhoto = document.createElement('div');
        postPhoto.className = 'postPhoto';
        let postImg = document.createElement('img');
        postImg.src = data.PostData.ContentPhoto;
        message.appendChild(postImg);
    }
    contenu.appendChild(message);

    let Infomsg = document.createElement('div');
    Infomsg.className = 'Infomsg';
    // date du post
    let InfomsgDate = document.createElement('div');
    InfomsgDate.textContent = 'Date de publication : ' + data.PostData.Date;
    InfomsgDate.className = 'InfomsgDate';
    Infomsg.appendChild(InfomsgDate);
    // catégorie du post
    let InfomsgCategory = document.createElement('div');
    InfomsgCategory.className = 'InfomsgCategory';
    let categories = '';
    for (let index = 0; index < data.PostData.Categorie.length; index++) {
        const categorie = data.PostData.Categorie[index];
        if (index === 0) {
            categories = categorie;
        } else {
            categories += ', ' + categorie;
        }

    }
    InfomsgCategory.innerHTML = 'Catégorie du post : ' + categories;
    Infomsg.appendChild(InfomsgCategory);
    contenu.appendChild(Infomsg);

    // -------------------------------------------------------------------------
    // ------------------------- Pour commenter le post ------------------------
    // -------------------------------------------------------------------------
    let comment = document.createElement('div');
    comment.className = 'comment';

    // bouton pour afficher l'éditeur de commentaire
    let commentPost = document.createElement('button');
    commentPost.id = 'displayButtonComment';
    commentPost.classList.add('buttonconnexion');
    commentPost.classList.add('buttoncomment');
    commentPost.textContent = 'Commenter le Post';
    commentPost.type = 'submit';
    commentPost.onclick = function () {
        document.getElementById('displayButtonComment').style.display = 'none';
        document.getElementById('postComment').style.display = 'block';
        document.getElementById('noneButtonComment').style.display = 'block';
    };
    comment.appendChild(commentPost);

    // bouton pour masquer l'éditeur de commentaire
    let displayPostEditor = document.createElement('button')
    displayPostEditor.id = 'noneButtonComment';
    displayPostEditor.style.display = 'none';
    displayPostEditor.classList.add('buttonconnexion');
    displayPostEditor.classList.add('buttoncomment');
    displayPostEditor.textContent = "Masquer l'éditeur de commentaire";
    displayPostEditor.type = 'submit';
    displayPostEditor.onclick = function () {
        document.getElementById('displayButtonComment').style.display = 'block';
        document.getElementById('postComment').style.display = 'none';
        document.getElementById('noneButtonComment').style.display = 'none';
    };
    comment.appendChild(displayPostEditor);
    contenu.appendChild(comment);

    // Container pour l'éditeur de commentaire
    let postComment = document.createElement('div');
    postComment.id = 'postComment';
    postComment.style.display = 'none';

    let formPost = document.createElement('form');
    formPost.method = 'POST';

    // titre zone de texte
    let ZoneTextEditor = document.createElement('div');
    ZoneTextEditor.className = 'ZoneTextEditor';
    let hTextEditor = document.createElement('h1');
    hTextEditor.innerHTML = 'Commenter la recette<br>' + data.PostData.Title;
    ZoneTextEditor.appendChild(hTextEditor);
    formPost.appendChild(ZoneTextEditor);

    // zone de saisie du texte
    let textarea = document.createElement('div');
    textarea.className = 'textarea';
    let textarea2 = document.createElement('textarea');
    textarea2.id = 'mytextarea';
    textarea2.name = 'newComment';
    textarea.appendChild(textarea2)
    formPost.appendChild(textarea);

    // input destiné à la récupération de l'id du post pour l'ajout dans la db
    let inputmasque = document.createElement('input');
    inputmasque.name = "post_id"
    inputmasque.value = postId;
    inputmasque.style.display = 'none';
    formPost.appendChild(inputmasque);


    let endPost = document.createElement('div');
    endPost.className = 'endPost';
    // Bouton d'envoi d'un nouveau commentaire
    let EnvoiePost = document.createElement('button');
    EnvoiePost.id = 'postNewComment';
    EnvoiePost.className = 'EnvoiePost';
    EnvoiePost.type = 'button';
    EnvoiePost.textContent = 'Publier le commentaire';
    endPost.appendChild(EnvoiePost)
    formPost.appendChild(endPost);
    postComment.appendChild(formPost);
    contenu.appendChild(postComment)
    post.appendChild(contenu);
    Container.appendChild(post);

    // -------------------------------------------------------------------------
    // ------------------- Affichage des commentaires du post ------------------
    // -------------------------------------------------------------------------

    if (data.PostData.Comments && data.PostData.Comments.length > 0) { // si un commentaire est présent
        let titlePostComment = document.createElement('div');
        titlePostComment.className = 'titlePostComment';
        titlePostComment.textContent = 'Commentaire(s) du Post';
        Container.appendChild(titlePostComment);

        // range de l'array pour crée toutes les div de commentaire
        for (let index = 0; index < data.PostData.Comments.length; index++) {
            const CurrentComment = data.PostData.Comments[index];

            let displayComment = document.createElement('div');
            displayComment.className = 'displayComment';

            // information sur la personne ayant posté le commentaire
            let userComment = document.createElement('div');
            let imguserComment = document.createElement('img');
            imguserComment.src = CurrentComment.PhotoAuthor;
            userComment.appendChild(imguserComment);
            userComment.textContent = 'Auteur : ' + CurrentComment.Author;
            displayComment.appendChild(userComment);

            // Boite pour le contenu du commentaire
            let contentComment = document.createElement('div');
            contentComment.className = 'contentComment';
            // message du commentaire
            let messageComment = document.createElement('div');
            messageComment.className = 'messageComment';
            messageComment.innerHTML = CurrentComment.Content;
            contentComment.appendChild(messageComment);
            // date du commentaire
            let infomsgComment = document.createElement('div');
            infomsgComment.className = 'infomsgComment';
            infomsgComment.textContent = 'Date du commentaire : ' + CurrentComment.Date;
            contentComment.appendChild(infomsgComment);
            displayComment.appendChild(contentComment);
            Container.appendChild(displayComment);
        }
    }

    // Charger la bibliothèque TinyMCE depuis le CDN
    if (typeof (tinyMCE) != "undefined") {
        tinymce.remove("#mytextarea");
    }
    tinymce.init({
        selector: '#mytextarea',
    });

    document.getElementById("postNewComment").addEventListener('click', postNewCommentButtonClick);
}

// -------------------------------------------------------------------------
// --------------- Button pour ajouter un nouveau commentaire --------------
// -------------------------------------------------------------------------

// Fonction gestionnaire d'événements lors du post d'un nouveau commentaire
let timerThrottlebutton = 0;
function postNewCommentButtonClick() {
    const now = new Date();
    if (now - timerThrottlebutton > 500) {
        timerThrottlebutton = now;
        sendFormNewComment()
    }
}

function sendFormNewComment() {
    // récupération des information saisie deans le formulaire
    var postId = document.getElementsByName('post_id')[0].value;
    let comment = tinymce.activeEditor.getContent(); // utilisation du getContent de tiny pour récupérerc le contenu

    const dataToSend = { postId, comment };
    console.log("dataToSend : ", dataToSend);

    fetch('http://localhost:8080/api/NewComment', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(dataToSend),
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`Erreur de réseau: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            if (typeof data === 'object') {
                console.log('Data received (New Comment):', data);
                if (data.UserData.Redirect !== "") {
                    home();
                } else {
                    msgError(data);
                    addNewComment();
                }
            } else {
                throw new Error('Réponse invalide du serveur (non-JSON)');
            }
        })
        .catch(error => {
            console.error('Erreur lors de la récupération des données:', error);
        });
}

function addNewComment() {
    console.log("ajout du nouveau commentaire a faire !!!!")
}