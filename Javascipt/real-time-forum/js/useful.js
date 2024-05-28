import { account } from "./account.js";
import { initChat, socket } from "./chat.js";
import { login } from "./login.js";
import { userNotConnected } from "./main.js";
import { register } from "./register.js";
export let chatusername = '';

export function commonBlock(data, confirmcomment = false) {
    badRequest();
    if (data.UserData.Logged) {
        document.getElementById('notConnected').style.display = 'none';
        document.getElementById('connected').style.display = 'block';
        document.getElementById('chat').style.display = 'flex';
        DisconnectedButton();
        AccountButton();
        if (chatusername === '') { // gestion du bug F5
            if (socket !== undefined) {
                socket.close();
            }
            initChat(data.UserData.Username);
        }
        chatusername = data.UserData.Username;
    } else {
        document.getElementById('notConnected').style.display = 'block';
        document.getElementById('connected').style.display = 'none';
        document.getElementById('chat').style.display = 'none';
        RegisterButton();
        LoginButton();

    }
    msgError(data, confirmcomment)

    var container = document.getElementById('Container');
    // Supprime tous les éléments enfants de la div
    while (container.firstChild) {
        container.removeChild(container.firstChild);
    }
    // Vide le contenu de la div
    container.innerHTML = '';
}

export function msgError(data, confirmcomment) {
    let divError = document.getElementById('error');
    if (data.UserData.Msgerr !== "") {
        divError.innerHTML = data.UserData.Msgerr;
        divError.style.display = "block";
    } else if (confirmcomment) {
        divError.textContent = "Votre commenntaire a bien été enregistré";
        divError.style.display = "block";
        window.scrollTo(0, 0);
    } else {
        divError.style.display = "none";
    }
}

function badRequest(){
    if (window.location.pathname !== "/"){
        window.location.href = "/";
    }
}

export function notFound() {
    document.getElementById('connected').style.display = 'none';
    document.getElementById('notConnected').style.display = 'none';
    let divError = document.getElementById('error');
    divError.textContent = "Erreur 404 : Page not Found";
    divError.style.color = 'red';
    divError.style.display = "block";
}

let timerThrottlebutton = 0;
function RegisterButton() {
    var RegisterButton = document.getElementById("RegisterButton");

    // Fonction gestionnaire d'événements
    function RegisterButtonClick() {
        const now = new Date();
        if (now - timerThrottlebutton > 500) {
            timerThrottlebutton = now;

            console.log("Page register demandé");
            register();

            // Supprime les gestionnaires d'événements devenu inutile
            RegisterButton.removeEventListener('click', RegisterButtonClick);
        }
    }

    RegisterButton.addEventListener('click', RegisterButtonClick);
}

function LoginButton() {
    var LoginButton = document.getElementById("LoginButton");

    // Fonction gestionnaire d'événements
    function LoginButtonClick() {
        const now = new Date();
        if (now - timerThrottlebutton > 500) {
            timerThrottlebutton = now;
            console.log("Page login demandé");
            login();
            LoginButton.removeEventListener('click', LoginButtonClick);
        }
    }

    LoginButton.addEventListener('click', LoginButtonClick);
}

export function AccountButton() {
    var AccountButton = document.getElementById("AccountButton");
    function AccountButtonClick() {
        const now = new Date();
        if (now - timerThrottlebutton > 500) {
            timerThrottlebutton = now;
            console.log("Page account demandé");
            account();
            AccountButton.removeEventListener('click', AccountButtonClick);
        }
    }
    AccountButton.addEventListener('click', AccountButtonClick);
}

export function DisconnectedButton() {
    var DisconnectedButton = document.getElementById("DisconnectButton");
    function DisconnectedButtonClick() {
        const now = new Date();
        if (now - timerThrottlebutton > 500) {
            timerThrottlebutton = now;
            // suppression du cookie
            document.cookie = "user_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
            // deconnexion du chat et vidage des éléments
            socket.close();
            chatusername = '';
            let logElement = document.getElementById("log");
            logElement.innerHTML = '';
            while (logElement.firstChild) {
                logElement.removeChild(logElement.firstChild);
            }

            // renvoie vers la home
            userNotConnected();
            socket.close();
        }
    }
    DisconnectedButton.addEventListener('click', DisconnectedButtonClick);
}
