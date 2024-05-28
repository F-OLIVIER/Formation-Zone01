import { login } from "./login.js";
import { commonBlock, msgError } from "./useful.js";


export function register() {
    // Affichage des boutons utilisateur non connecté
    let UserData = { Logged: false, Msgerr: "" };
    commonBlock({ UserData });

    let Container = document.getElementById('Container');
    let box = document.createElement('div');
    box.className = "box register-box";

    let topheader = document.createElement('div');
    topheader.className = 'top-header';
    let Title = document.createElement('header');
    Title.textContent = "Formulaire d'enregistrement";
    topheader.appendChild(Title);
    box.appendChild(topheader);

    let form = document.createElement('form');
    form.id = 'registerForm';
    form.className = 'input-field';
    form.method = 'POST';
    let divinput = document.createElement('div');
    divinput.className = 'divinput';

    // Nickname
    let spanUsername = document.createElement('span');
    spanUsername.className = 'spanusername';
    let username = document.createElement('input');
    username.type = 'text';
    username.className = 'input';
    username.placeholder = 'Pseudo';
    username.id = 'usernameregister';
    username.name = 'usernameregister';
    username.required;
    divinput.appendChild(spanUsername);
    spanUsername.appendChild(username);
    

    // First Name
    let spanFirstName = document.createElement('span');
    spanFirstName.className = 'spanfirstname';
    let FirstName = document.createElement('input');
    FirstName.type = 'text';
    FirstName.className = 'input';
    FirstName.placeholder = 'Prénom';
    FirstName.id = 'FirstNameregister';
    FirstName.name = 'FirstNameregister';
    FirstName.required;
    divinput.appendChild(spanFirstName);
    spanFirstName.appendChild(FirstName);
    

    // Last Name
    let spanLastName = document.createElement('span');
    spanLastName.className = 'spanlastname';
    let LastName = document.createElement('input');
    LastName.type = 'text';
    LastName.className = 'input';
    LastName.placeholder = 'Nom de famille';
    LastName.id = 'LastNameregister';
    LastName.name = 'LastNameregister';
    LastName.required;
    divinput.appendChild(spanLastName);
    spanLastName.appendChild(LastName);

    // Age
    let spanAge = document.createElement('span');
    spanAge.className = 'spanage';
    let Age = document.createElement('input');
    Age.type = 'number';
    Age.className = 'input';
    Age.placeholder = 'Age';
    Age.id = 'ageregister';
    Age.name = 'ageregister';
    Age.required;
    divinput.appendChild(spanAge);
    spanAge.appendChild(Age);

    // Gender
    let spanGender = document.createElement('span');
    spanGender.className = 'spangender';
    let labelGender = document.createElement('label');
    labelGender.setAttribute("for", "labelGender");
    labelGender.textContent = "Choississez votre genre :";
    let Gender = document.createElement('select');
    Gender.className = 'input';
    Gender.id = "Genderregister";
    Gender.name = "Genderregister";
    Gender.required;
    // Ajout de l'option par défaut
    const defaultOption = document.createElement("option");
    defaultOption.value = "";
    // defaultOption.className = 'input';
    defaultOption.textContent = "--Choisissez une option--";
    Gender.appendChild(defaultOption);
    // Ajout des options individuelles
    const option = ["Homme", "Femme", "Non genré"];
    for (const element of option) {
        const currentoption = document.createElement("option");
        currentoption.value = element;
        currentoption.textContent = element;
        Gender.appendChild(currentoption);
    }
    divinput.appendChild(spanGender);
    spanGender.appendChild(labelGender);
    spanGender.appendChild(Gender);


    // E-mail
    let spanEmail = document.createElement('span');
    spanEmail.className = 'spanemail';
    let email = document.createElement('input');
    email.type = 'email';
    email.size = '30';
    email.className = 'input';
    email.placeholder = 'Adresse e-mail';
    email.id = 'mailregister';
    email.name = 'mailregister';
    email.required;
    divinput.appendChild(spanEmail);
    spanEmail.appendChild(email);
    

    // password
    let spanPassword = document.createElement('span');
    spanPassword.className = 'spanpassword';
    let password = document.createElement('input');
    password.type = 'password';
    password.className = 'input';
    password.placeholder = 'Mot de passe';
    password.id = 'passregister';
    password.name = 'passregister';
    password.required;
    divinput.appendChild(spanPassword);
    spanPassword.appendChild(password);

    // Confirmation du password
    let spanConfirmPass = document.createElement('span');
    spanConfirmPass.className = 'spanconfirmpass';
    let confirmPass = document.createElement('input');
    confirmPass.type = 'password';
    confirmPass.className = 'input';
    confirmPass.placeholder = 'Confirmer le mot de passe';
    confirmPass.id = 'confirmPassregister';
    confirmPass.name = 'confirmPassregister';
    confirmPass.required;
    divinput.appendChild(spanConfirmPass);
    spanConfirmPass.appendChild(confirmPass);

    // Button de validation
    let register = document.createElement('button');
    register.type = 'button';
    register.id = 'Register'
    register.className = 'submit';
    register.textContent = "S'enregistrer";
    divinput.appendChild(register);
    form.appendChild(divinput);
    box.appendChild(form);
    Container.appendChild(box);

    document.getElementById('Register').addEventListener('click', ButtonClickRegister);
}

function validForm() {
    var inputs = ['usernameregister', 'FirstNameregister', 'LastNameregister', 'ageregister', 'Genderregister', 'mailregister', 'passregister', 'confirmPassregister'];
    for (var i = 0; i < inputs.length; i++) {
        // Vérifier si le champ est vide
        if (document.getElementById(inputs[i]).value.trim() === '') {
            alert('Veuillez remplir correctement tous les champs.');
            return false; // Arrêter la soumission du formulaire
        }
    }
    return true;
}

let timerThrottlebutton = 0;
// Fonction gestionnaire d'événements
function ButtonClickRegister() {
    const now = new Date();
    if (now - timerThrottlebutton > 500) {
        timerThrottlebutton = now;
        if (document.getElementById('registerForm').onsubmit = validForm()) {
            sendformRegister();
        }
    }
}

function sendformRegister() {
    // récupération des information saisie deans le formulaire
    var usernameValue = document.getElementsByName('usernameregister')[0].value;
    var FirstNameregister = document.getElementsByName('FirstNameregister')[0].value;
    var LastNameregister = document.getElementsByName('LastNameregister')[0].value;
    var Ageregister = document.getElementsByName('ageregister')[0].value;
    var Genderregister = document.getElementsByName('Genderregister')[0].value;
    var mailValue = document.getElementsByName('mailregister')[0].value;
    var passValue = document.getElementsByName('passregister')[0].value;
    var confirmPassregister = document.getElementsByName('confirmPassregister')[0].value;
    const dataToSend = { usernameValue, FirstNameregister, LastNameregister, Ageregister, Genderregister, mailValue, passValue, confirmPassregister };
    console.log("dataToSend : ", dataToSend);

    fetch('http://localhost:8080/api/Register', {
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
                console.log('Data received (Register):', data);

                if (data.UserData.Msgerr === "") {
                    document.getElementById('Register').removeEventListener('click', ButtonClickRegister);
                    login();
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

