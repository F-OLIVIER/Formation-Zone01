import { commonBlock, msgError } from "./useful.js";
import { home } from "./home.js";

export function login() {
    // Affichage des boutons utilisateur non connecté
    let UserData = { Logged: false, Msgerr: "" };
    commonBlock({ UserData });

    let Container = document.getElementById('Container');

    let box = document.createElement('div');
    box.className = "box login-box";

    // titre
    let divTitle = document.createElement('div');
    divTitle.className = 'top-header';
    let Title = document.createElement('header');
    Title.textContent = "Formulaire de connexion";
    divTitle.appendChild(Title);
    box.appendChild(divTitle);

    // Formulaire
    let form = document.createElement('form');
    form.className = 'input-field';
    form.method = 'POST';
    let divinput = document.createElement('div');
    divinput.className = 'divinput';

    // Pseudo ou Email
    let spanEmail = document.createElement('span');
    spanEmail.className = 'spanemail';
    let email = document.createElement('input');
    email.type = 'email';
    email.size = '30';
    email.className = 'input';
    email.placeholder = 'Pseudo ou adresse e-mail';
    email.name = 'maillogin';
    email.required;
    divinput.appendChild(spanEmail);
    spanEmail.appendChild(email);

    // Password
    let spanPassword = document.createElement('span');
    spanPassword.className = 'spanpassword';
    let password = document.createElement('input');
    password.type = 'password';
    password.className = 'input';
    password.placeholder = 'Mot de passe';
    password.name = 'passlogin';
    password.required;
    divinput.appendChild(spanPassword);
    spanPassword.appendChild(password);

    // Button
    let connexion = document.createElement('button');
    connexion.type = 'button';
    connexion.id = 'SendLoginButton';
    connexion.className = 'submit';
    connexion.textContent = "Se connecter";
    divinput.appendChild(connexion);
    form.appendChild(divinput);
    box.appendChild(form);
    Container.appendChild(box);

    document.getElementById('SendLoginButton').addEventListener('click', ButtonClickLogin);
}

// Fonction gestionnaire d'événements
let timerThrottlebutton = 0;
function ButtonClickLogin() {
    const now = new Date();
    if (now - timerThrottlebutton > 500) {
        timerThrottlebutton = now;
        sendformLogin();
    }
}

function sendformLogin() {
    // récupération des information saisie deans le formulaire
    var mailValue = document.getElementsByName('maillogin')[0].value;
    var passValue = document.getElementsByName('passlogin')[0].value;
    const dataToSend = { mailValue, passValue };
    console.log("dataToSend : ", dataToSend);

    fetch('http://localhost:8080/api/Login', {
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
                console.log('Data received (Login):', data);
                if (data.UserData.Redirect !== "") {
                    document.getElementById('LoginButton').removeEventListener('click', ButtonClickLogin);
                    console.log("data.UserData.Username : ", data.UserData.Username);
                    home();
                } else {
                    msgError(data);
                }
            } else {
                throw new Error('Réponse invalide du serveur (non-JSON)');
            }
        })
        .catch(error => {
            console.error('Erreur lors de la récupération des données:', error);
        });
}

