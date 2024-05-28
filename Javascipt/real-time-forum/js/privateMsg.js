import { socket } from "./chat.js";
import { chatusername } from "./useful.js";

export function historyMsg(sender, receiver) {
    const dataToSend = { sender, receiver };
    // console.log("sender : ", sender, "\nreceiver : ", receiver)
    // console.log("dataToSend : ", dataToSend);

    fetch('http://localhost:8080/api/historyMsg', {
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
                // console.log('Data received (historyMsg):', data);
                if (data.History != null) {
                    generateHistory(data.History, receiver);
                }
            } else {
                throw new Error('Réponse invalide du serveur (non-JSON)');
            }
        })
        .catch(error => {
            console.error('Erreur lors de la récupération des données:', error);
        });
}

let nbCharger = 0;
let logElement = document.getElementById("log");
function generateHistory(data, receiver) {
    // nettoyage de la div
    logElement.innerHTML = '';
    while (logElement.firstChild) {
        logElement.removeChild(logElement.firstChild);
    }

    nbCharger = 0;
    // génére l'historique dans la div
    if (data.ListMsg !== null) {
        if (data.ListMsg.length > 10) {
            nbCharger = 10;
        } else {
            nbCharger = data.ListMsg.length;
        }
        loadmsg(data, receiver, nbCharger);
        nbCharger += 1;

        logElement.addEventListener('scroll', function () {
            if (logElement.scrollTop === 0 && nbCharger < data.ListMsg.length) {
                console.log(nbCharger);
                loadoldmsg(data, receiver, nbCharger);
                nbCharger += 10;
            }
        });
        // for (let i = data.ListMsg.length - nbCharger; i < data.ListMsg.length; i++) {
        //     const structMessage = data.ListMsg[i];

        //     let p = document.createElement('p');
        //     p.textContent = structMessage.senderMsg + ": " + structMessage.contentMsg;
        //     let span = document.createElement('span');
        //     span.textContent = structMessage.timestampMsg;
        //     p.appendChild(span);
        //     logElement.appendChild(p);
        //     if (structMessage.senderMsg === receiver) {
        //         p.className = "receiver";
        //     } else {
        //         p.className = "sender";
        //     }
        // }
        logElement.scrollTop = logElement.scrollHeight; // Défilement automatique vers le bas à chaque nouveau message
    }

    // mise à jour des div 
    socket.send(JSON.stringify({ username: chatusername, MajlistConnected: true }));
}

function eventscroll(data, receiver, nbCharger) {
    if (logElement.scrollTop === 0 && nbCharger < data.ListMsg.length) {
        console.log(nbCharger);
        loadoldmsg(data, receiver, nbCharger);
        nbCharger += 10;
    }
};

function loadmsg(data, receiver, nbCharger) {
    for (let i = data.ListMsg.length - nbCharger; i < data.ListMsg.length; i++) {
        const structMessage = data.ListMsg[i];

        let p = document.createElement('p');
        p.textContent = structMessage.senderMsg + ": " + structMessage.contentMsg;
        let span = document.createElement('span');
        span.textContent = structMessage.timestampMsg;
        p.appendChild(span);
        logElement.appendChild(p);
        if (structMessage.senderMsg === receiver) {
            p.className = "receiver";
        } else {
            p.className = "sender";
        }
    }
}

function loadoldmsg(data, receiver, nbCharger) {
    console.log('loadoldmsg 1 :', data.ListMsg.length - nbCharger)
    console.log('loadoldmsg 2 :', data.ListMsg.length - nbCharger - 10);
    for (let i = data.ListMsg.length - nbCharger; i > data.ListMsg.length - nbCharger - 10; i--) {

        console.log('HERE');
        if (i < 0) {
            break;
        }
        console.log('HERE OKKKKKKKKKKKKKKKKKKKK');

        const structMessage = data.ListMsg[i];

        let p = document.createElement('p');
        p.textContent = structMessage.senderMsg + ": " + structMessage.contentMsg;
        let span = document.createElement('span');
        span.textContent = structMessage.timestampMsg;
        p.appendChild(span);
        // logElement.appendChild(p);
        console.log('p : ', p);
        logElement.insertAdjacentElement('afterbegin', p);
        logElement.scrollTop += 1;
        if (structMessage.senderMsg === receiver) {
            p.className = "receiver";
        } else {
            p.className = "sender";
        }
    }
}