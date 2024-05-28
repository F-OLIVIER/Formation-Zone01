import { cookieName } from "./main.js";
import { commonBlock, msgError } from "./useful.js";


export function account() {
    fetch('http://localhost:8080/api/Compte', {
        method: 'POST',
    })
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
            console.log('Data received (Compte):', data);
            if (!data.UserData.Logged) { // si le cookie n'ai pas valide, suppression + redirection
                document.cookie = cookieName + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
                userNotConnected();
                return;
            }
            accountElement(data);
        })
        .catch(error => {
            // Gérer les erreurs
            console.error('Data recovery error:', error);
        });

}
function accountElement(data) {
    commonBlock(data);
    let Container = document.getElementById('Container');
    let box = document.createElement('div');
    box.className = "box account-box";

    let topheader = document.createElement('div');
    topheader.className = 'top-header';
    let Title = document.createElement('header');
    Title.textContent = "Mes Informations";
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
    let username = document.createElement('div');
    username.className = 'input';
    username.textContent = data.UserData.Username;
    divinput.appendChild(spanUsername);
    spanUsername.appendChild(username);

    // First Name
    let spanFirstName = document.createElement('span');
    spanFirstName.className = 'spanfirstname';
    let FirstName = document.createElement('div');
    FirstName.className = 'input';
    FirstName.textContent = data.UserData.FirstName;
    divinput.appendChild(spanFirstName);
    spanFirstName.appendChild(FirstName);


    // Last Name
    let spanLastName = document.createElement('span');
    spanLastName.className = 'spanlastname';
    let LastName = document.createElement('div');
    LastName.className = 'input';
    LastName.textContent = data.UserData.LastName;
    divinput.appendChild(spanLastName);
    spanLastName.appendChild(LastName);

    // Age
    let spanAge = document.createElement('span');
    spanAge.className = 'spanage';
    let Age = document.createElement('div');
    Age.className = 'input';
    Age.textContent = data.UserData.Age;
    divinput.appendChild(spanAge);
    spanAge.appendChild(Age);

    // Gender
    let spanGender = document.createElement('span');
    spanGender.className = 'spangender';
    let Gender = document.createElement('div');
    Gender.className = 'input';
    Gender.textContent = data.UserData.Gender;
    divinput.appendChild(spanGender);
    spanGender.appendChild(Gender);

    // E-mail
    let spanEmail = document.createElement('span');
    spanEmail.className = 'spanemail';
    let email = document.createElement('div');
    email.className = 'input';
    email.textContent = data.UserData.Email;
    divinput.appendChild(spanEmail);
    spanEmail.appendChild(email);
    

    // Button de validation
    // let register = document.createElement('button');
    // register.type = 'button';
    // register.id = 'Register'
    // register.className = 'submit';
    // register.textContent = "S'enregistrer";
    // divinput.appendChild(register);
    form.appendChild(divinput);
    box.appendChild(form);
    Container.appendChild(box);
}
function containerCompte() {
    // let UserData = { Logged: false, Msgerr: "" };
    // commonBlock({ UserData });

}