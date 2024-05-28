export let socket;
import { historyMsg } from "./privateMsg.js";
import { chatusername } from "./useful.js";

let username_to_send = '';
let typingTimeout;
let timerThrottlebutton = new Date();

export async function initChat(username) {
    socket = new WebSocket("ws://localhost:8080/ws");
    socket.onmessage = function (event) { // Réception d'un message
        // console.log("username_to_send : ", username_to_send)
        const logElement = document.getElementById("log");
        const structMessage = JSON.parse(event.data);
        let typinprogress = document.getElementById('typinprogress');
        // console.log("Receve structMessage : ", structMessage);
        // si message n'ai pas vide && ce n'ai pas un message de mise a jour de la liste && Si l'utilisateur à la discussion ouverte avec le sender du message
        if (structMessage.message != "" && !structMessage.majlistConnected && !structMessage.TypingProgress) {
            if (((structMessage.userPrivateMessage === username_to_send && structMessage.username === username) ||
                (structMessage.username === username_to_send && structMessage.userPrivateMessage === username)) &&
                username_to_send != '') {
                // logElement.innerHTML += "<p><span>" + structMessage.horaire + "</span>" + structMessage.username + ": " + structMessage.message + "</p>";
                let p = document.createElement('p');
                p.textContent = structMessage.username + ": " + structMessage.message;
                let span = document.createElement('span');
                span.textContent = structMessage.horaire;
                p.appendChild(span);
                if (structMessage.username === username) {
                    p.className = "sender";
                } else {
                    p.className = "receiver";
                }
                logElement.appendChild(p);

                typinprogress.textContent = '';
                clearTimeout(typingTimeout);
            }
        } else if (structMessage.typingProgress &&
            structMessage.userPrivateMessage === username &&
            structMessage.username === username_to_send) {
            typinprogress.textContent = "l'autre gland de " + structMessage.username + " est en train d'écrire...";
            // Annuler le timeout précédent (s'il existe)
            if (typingTimeout) {
                clearTimeout(typingTimeout);
            }
            // Créer un nouveau timeout
            typingTimeout = setTimeout(function () {
                typinprogress.textContent = '';
            }, 3 * 1000); // 3 secondes
        }
        updateListUserInChat(structMessage, username)

        logElement.scrollTop = logElement.scrollHeight; // Défilement automatique vers le bas à chaque nouveau message
    };

    // gestionnaire d'événement pour pouvoir envoyer des message via 'enter' sans recharger la page
    document.getElementById("formWebSocket").addEventListener("submit", function (event) {
        console.log("username_to_send : ", username_to_send, timerThrottlebutton);
        const now = new Date();
        if (username_to_send === '') { // gestion d'erreur, si pas de destinataire
            alert('Merci de préciser à qui le message doit être envoyé !')
        } else if (now - timerThrottlebutton > 2000) {
            timerThrottlebutton = now;
            event.preventDefault(); // Empêche la soumission par défaut du formulaire (evite le rechargement de page)
            sendMessage(); // Appelle la fonction sendMessage lors de la soumission du formulaire
        } else {
            alert('Trop rapide pour le piaf voyageur !!!!!')
        }
    });


    socket.onopen = function (event) {
        const newUser = { username: username, message: "-->New_User_In_Chat<--" };
        socket.send(JSON.stringify(newUser));
    };

    socket.onclose = function (event) {
        const newUser = { username: username, message: "-->User_Leave_Chat<--" };
        socket.send(JSON.stringify(newUser));
    };
    // gestionnaire d'événement pour pouvoir le "typing in progress"
    document.getElementById("msg").addEventListener("input", function () {
        console.log("Le champ 'msg' a été modifié !");
        // const newUser = { username: username, UserPrivateMessage: username_to_send, TypingProgress: true, message: '' };
        // socket.send(JSON.stringify(newUser));
    });

    // gestionnaire d'événement pour pouvoir le "typing in progress"
    document.getElementById("msg").addEventListener("input", function () {
        // console.log("Le champ 'msg' a été modifié !");
        const typingProgress = { username: username, UserPrivateMessage: username_to_send, TypingProgress: true, message: '' };
        socket.send(JSON.stringify(typingProgress));
    });
}

export function sendMessage() { // envoi d'un message
    const msgInput = document.getElementById("msg");
    const message = msgInput.value;
    // console.log("Send message : ", message);

    if (message.trim() !== "") {
        const chatMessage = { username: chatusername, message: message, Horaire: horodatage(), UserPrivateMessage: username_to_send };
        socket.send(JSON.stringify(chatMessage));
        msgInput.value = "";
    }
}

function updateListUserInChat(structMessage, username) {
    const listUser = structMessage.listUser;
    const listUserConnected = structMessage.listUserConnected;
    const listnewmsg = structMessage.listUserNewMsg;

    let divListUser = document.getElementById('userConnected')
    divListUser.innerHTML = `<p>Connecté en tant que : ${username}</p><p></p>`;

    if (listUserConnected.length > 1) {
        divListUser.innerHTML += `<p> Utilisateur connecté au chat</p>`;
        // utilisateur avec un nouveau message
        majList(listUser, listnewmsg, listUserConnected, username, true, true);
        // utilisateur sans nouveau messages
        majList(listUser, listnewmsg, listUserConnected, username, true, false);
    } else {
        divListUser.innerHTML += '<p>Aucun utilisateur en ligne</p>';
    }
    divListUser.innerHTML += '<p></p><p>Utilisateur hors ligne</p>';

    // utilisateur hors ligne avec un nouveau message
    majList(listUser, listnewmsg, listUserConnected, username, false, true);
    // utilisateur hors ligne sans un nouveau message
    majList(listUser, listnewmsg, listUserConnected, username, false, false);

    // gestionnaire d'événements des bouttons
    var buttons = document.querySelectorAll('button[name="user_ID"]');
    buttons.forEach(function (button) {
        button.addEventListener('click', function () {
            username_to_send = button.value;
            historyMsg(username, button.value);
        });
    });
}

function majList(listUser, listnewmsg, listUserConnected, username, inline, displayUser) {
    let divListUser = document.getElementById('userConnected');
    for (let i = 0; i < listUser.length; i++) {
        const nameUser = listUser[i];
        const newMsg = listnewmsg[i];
        let button = document.createElement('button');
        button.name = "user_ID";
        button.value = nameUser;
        let addbutton = false;

        if (nameUser != username) {
            if (inline && checkUser(nameUser, listUserConnected)) { // Utilisateur  en ligne
                if (displayUser && newMsg) { // boucle 1 : affichage des utilisateurs avec nouveaux message
                    button.textContent = "😺 " + nameUser + " (nouveau message)";
                    addbutton = true;
                } else if (!displayUser && !newMsg) { // boucle 2 : affichage des utilisateurs sans nouveaux message
                    button.textContent = "😺 " + nameUser;
                    addbutton = true;
                }
            } else if (!inline && !checkUser(nameUser, listUserConnected)) { // Utilisateur hors ligne
                if (displayUser && newMsg) { // boucle 1 : affichage des utilisateurs avec nouveaux message
                    button.textContent = "😺 " + nameUser + " (nouveau message)";
                    addbutton = true;
                } else if (!displayUser && !newMsg) { // boucle 2 : affichage des utilisateurs sans nouveaux message
                    button.textContent = "😺 " + nameUser;
                    addbutton = true;
                }
            }
        }

        if (addbutton) {
            if (nameUser === username_to_send) {
                button.classList.add('clicked');
            }
            divListUser.appendChild(button);
        }
    }
}

function checkUser(user, listUserConnected) {
    for (let i = 0; i < listUserConnected.length; i++) {
        if (user === listUserConnected[i]) {
            return true
        }
    }
    return false
}

function horodatage() {
    const now = new Date();
    return '[' +
        String(now.getDate()).padStart(2, '0') + '-' +
        String(now.getMonth() + 1).padStart(2, '0') + '-' +
        now.getFullYear() + ' ' +
        String(now.getHours()).padStart(2, '0') + ':' +
        String(now.getMinutes()).padStart(2, '0') + ':' +
        String(now.getSeconds()).padStart(2, '0') +
        '] ';
}

// Position fenêtre de chat
function positionChat() {
    let chat = document.getElementById('chat');
    let btnChat = document.createElement('div');
    btnChat.className = 'btn-chat';

    let btn = document.createElement('div');
    btn.className = 'btn-icon';
    btnChat.appendChild(btn);

    chat.appendChild(btnChat);
    let btnClick = document.querySelector('.btn-chat');
    let icon = document.querySelector('.btn-icon');

    btnClick.addEventListener('click', function () {
        let close = document.getElementById('bodyContainer');
        // console.log("close : ", close);
        if (chat.style.right === '0px') { // demande fermeture chat
            chat.style.right = '-750px';
            icon.style.transform = 'rotate(0)';
        } else {  // demande ouverture chat
            chat.style.right = 0;
            icon.style.transform = 'rotate(180deg)';
        }
    });

}

addEventListener('load', positionChat);




