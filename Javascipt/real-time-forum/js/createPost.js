import { commonBlock } from "./useful.js";
import { cookieName, userNotConnected } from "./main.js";
import { home } from "./home.js";

const listCat = ['Boissons', 'Sauces', 'Entrées', 'Poissons', 'Viandes', 'Desserts', 'Autres'];
let tinymceLoad = false;

export function createPost() {
    // Check la precence du cookie, si il n'existe pas redirection automatique vers "/"
    if (!document.cookie.split(";").some((item) => item.trim().startsWith(cookieName + "="))) {
        userNotConnected();
        return;
    }

    let UserData = { Logged: true, Msgerr: "" };
    commonBlock({ UserData });

    let Container = document.getElementById('Container');
    let box = document.createElement('div');
    box.className = "box";

    let form = document.createElement('form');
    form.method = 'POST';
    form.enctype = "multipart/form-data";

    // En tête
    let ZoneTextEditor = document.createElement('div');
    ZoneTextEditor.className = 'ZoneTextEditor';
    let h1 = document.createElement('h1');
    h1.textContent = "Publiez une recette";
    ZoneTextEditor.appendChild(h1);
    let labelZoneTextEditor = document.createElement('label');
    labelZoneTextEditor.className = "titlerecette";
    labelZoneTextEditor.textContent = "Nom de la recette :";
    ZoneTextEditor.appendChild(labelZoneTextEditor);
    let inputNameRecette = document.createElement('input');
    inputNameRecette.type = "text";
    inputNameRecette.className = 'inputRecette';
    inputNameRecette.name = 'NomRecette';
    ZoneTextEditor.appendChild(inputNameRecette);
    form.appendChild(ZoneTextEditor);

    // zone de texte
    let divTextarea = document.createElement('div');
    divTextarea.className = 'textarea';
    let textarea = document.createElement('textarea');
    textarea.id = 'mytextarea';
    textarea.placeholder = 'Saisissez votre recette';
    divTextarea.appendChild(textarea);
    form.appendChild(divTextarea);

    let endPost = document.createElement('div');
    endPost.className = 'endPost';
    // affichage des catégories du post
    let categories = document.createElement('div');
    categories.id = 'categories';
    let labelCategorie = document.createElement('label');
    labelCategorie.id = 'ChoixDesTags';
    labelCategorie.textContent = 'Choisissez votre ou vos catégorie(s) (au minimum 1 catégorie et au maximum 3 catégories) :';
    categories.appendChild(labelCategorie);
    // cat 1 :
    for (const element of listCat) {
        const currentCat = document.createElement('input');
        currentCat.type = 'checkbox';
        currentCat.name = 'categoriePost' + element;
        currentCat.id = 'checkbox_categories' + element;
        currentCat.value = element.toLowerCase();
        // label
        const labelCat = document.createElement('label');
        labelCat.setAttribute("for", 'checkbox_categories' + element);
        labelCat.textContent = element + ' ';
        categories.appendChild(currentCat);
        categories.appendChild(labelCat);
    }
    endPost.appendChild(categories);
    Container.appendChild(endPost);

    // button de publication
    let createpost = document.createElement('button');
    createpost.type = 'button';
    createpost.id = 'createpost'
    createpost.className = 'submit';
    createpost.classList.add('buttonconnexion');
    createpost.textContent = "Publier";
    endPost.appendChild(createpost);
    form.appendChild(endPost);
    box.appendChild(form);
    Container.appendChild(box);

    // Charger la bibliothèque TinyMCE depuis le CDN
    if (typeof (tinyMCE) != "undefined") {
        tinymce.remove("#mytextarea");
    }
    tinymce.init({
        selector: '#mytextarea',
    });
    
    document.getElementById('createpost').addEventListener('click', ButtonClickCreatepost);
}

// Fonction gestionnaire d'événements
let timerThrottlebutton = 0;
function ButtonClickCreatepost() {
    const now = new Date();
    if (now - timerThrottlebutton > 500) {
        timerThrottlebutton = now;
        sendformCreatepost();
    }
}

function sendformCreatepost() {
    // récupération des information saisie deans le formulaire
    let nameRecette = document.getElementsByName('NomRecette')[0].value;
    let recette = tinymce.activeEditor.getContent(); // utilisation du getContent de tiny pour récupérerc le contenu
    let Category = [];
    for (const cat of listCat) {
        const checkbox = document.getElementById('checkbox_categories' + cat);
        if (checkbox.checked) {
            let nameCat = checkbox.value
            Category.push(nameCat.charAt(0).toUpperCase() + nameCat.slice(1));

        }
    }
    const dataToSend = { nameRecette, recette, Category };
    console.log("dataToSend : ", dataToSend);

    fetch('http://localhost:8080/api/PostEditor', {
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
                console.log('Data received (createPost):', data);
                if (data.UserData.Redirect === "/home") {
                    document.getElementById('createpost').removeEventListener('click', ButtonClickCreatepost);
                    home();
                } else {
                    let divError = document.getElementById('error');
                    divError.innerHTML = data.UserData.Msgerr;
                    divError.style.display = "block";
                }
            } else {
                throw new Error('Réponse invalide du serveur (non-JSON)');
            }
        })
        .catch(error => {
            console.error('Erreur lors de la récupération des données:', error);
        });
}

