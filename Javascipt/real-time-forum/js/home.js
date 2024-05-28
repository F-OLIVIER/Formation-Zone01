import { commonBlock } from "./useful.js";
import { cookieName, userNotConnected } from "./main.js";
import { viewPost } from "./viewpost.js";
import { createPost } from "./createPost.js";

export function home() {
  // Si le cookie est present, fetch des données. Le back fera une vérification de la validité du cookie
  fetch('http://localhost:8080/api/home')
    .then(response => {
      // Vérifier si la requête a réussi (status code 200)
      if (!response.ok) {
        throw new Error(`Erreur de réseau: ${response.status}`);
      }

      // Convertir la réponse en JSON
      return response.json();
    })
    .then(data => {
      // Traiter les données récupérées
      console.log('Data received (home):', data);
      if (!data.UserData.Logged) { // si le cookie n'ai pas valide, suppression + redirection
        document.cookie = cookieName + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
        userNotConnected();
        return;
      }
      containerhome(data);
    })
    .catch(error => {
      // Gérer les erreurs
      console.error('Data recovery error:', error);
    });
}


function containerhome(data) {
  if (data.UserData.Logged) {
    // génére les div commune à toutes les pages HTML
    commonBlock(data);

    let Container = document.getElementById('Container');

    // création du bouton de publication d'un post
    let buttonCreatePost = document.createElement('button');
    buttonCreatePost.classList.add('buttonconnexion');
    buttonCreatePost.classList.add('buttonCreatePost');
    buttonCreatePost.id = 'newPost';
    buttonCreatePost.textContent = 'Crée un nouveau post';
    Container.appendChild(buttonCreatePost);
    activateButtonNewPost();

    let listpost = document.createElement('div');
    listpost.id = 'listPost';
    listpost.className = 'listPost';

    // affichage les posts de la page
    if (data.PostListData !== null) {
      for (let indexlistpost = 0; indexlistpost < data.PostListData.length; indexlistpost++) {
        const elementlistpost = data.PostListData[indexlistpost];

        let button = document.createElement('button');
        button.id = "idPost";
        button.value = elementlistpost.ID;
        button.classList = 'post';


        let title = document.createElement('div');
        title.className = 'title';
        title.textContent = elementlistpost.Title;
        button.appendChild(title);

        let postGauche = document.createElement('div');
        postGauche.className = 'postGauche';
        let Categories = '';
        for (let postGauche = 0; postGauche < elementlistpost.Categorie.length; postGauche++) {
          if (postGauche === 0) {
            Categories = elementlistpost.Categorie[postGauche];
          } else {
            Categories += " " + elementlistpost.Categorie[postGauche];
          }
        }
        postGauche.innerHTML = "Date de parution : " + elementlistpost.Date + '<br>Categories :' + Categories;
        button.appendChild(postGauche);

        let postDroite = document.createElement('div');
        postDroite.className = 'postDroite';
        postDroite.innerHTML = "Auteur : " + elementlistpost.Author;
        button.appendChild(postDroite);

        // link.appendChild(button);
        listpost.appendChild(button);
        Container.appendChild(listpost);
      }

      activateAllButtonViewpost();
    }
  } else {
    userNotConnected();
  }
}

function activateAllButtonViewpost() {
  // Création de écouteur d'événements pour chaque bouton
  var postButtons = document.querySelectorAll('[id^="idPost"]');
  postButtons.forEach(function (button) {
    button.addEventListener('click', handleClick);
  });
}
// Fonction gestionnaire d'événements
function handleClick() {
  var buttonValue = this.value;

  // Désactive tous les boutons en supprimant les écouteurs d'événements
  var postButtons = document.querySelectorAll('[id^="idPost"]');
  postButtons.forEach(function (btn) {
    btn.removeEventListener('click', handleClick);
  });

  // Traitement de l'information pour afficher le post pour le button cliqué
  viewPost(buttonValue);
}

let timerThrottlebutton = 0;
export function activateButtonNewPost() {
  var NewPost = document.getElementById("newPost");
  function NewPostClick() {
    const now = new Date();
    if (now - timerThrottlebutton > 500) {
      timerThrottlebutton = now;
      console.log("Page newPost demandé");
      NewPost.removeEventListener('click', NewPostClick);

      // Désactivation les boutons permettant d'afficher un post spécifique
      var postButtons = document.querySelectorAll('[id^="idPost"]');
      postButtons.forEach(function (btn) {
        btn.removeEventListener('click', handleClick);
      });

      createPost();
    }
  }
  NewPost.addEventListener('click', NewPostClick);
}