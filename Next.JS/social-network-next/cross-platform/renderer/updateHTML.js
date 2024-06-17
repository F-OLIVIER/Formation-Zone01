import { online } from "./onlineofflinemode.js";
import { deleteDiscussionContents, fetchSendServer } from "./useful.js";

export async function updateUserList(userDataStorage, currentID = 0, uuid) {
    let listuser, currentUserOnline;
    if (online) {
        const data = await fetchSendServer('listuser', { id: currentID });
        console.log('data updateUserList : ', data);
        listuser = data.listuser
        currentUserOnline = data.online
    }

    // vider la liste
    const onlineUser = document.getElementById('onlineuser');
    while (onlineUser.firstChild) {
        onlineUser.removeChild(onlineUser.firstChild);
    }
    const offlineUser = document.getElementById('offlineuser');
    while (offlineUser.firstChild) {
        offlineUser.removeChild(offlineUser.firstChild);
    }

    // ajout titres Offline
    const titleofflineUser = document.createElement('p');
    titleofflineUser.textContent = 'Offline User';
    offlineUser.appendChild(titleofflineUser);

    if (online && listuser != undefined && currentUserOnline != undefined) { // Online et résultats des fetch ok
        // ajout titres Online
        const titleonlineUser = document.createElement('p');
        titleonlineUser.textContent = 'Online User';
        onlineUser.appendChild(titleonlineUser);

        for (let i = 0; i < listuser.length; i++) {
            const currentUser = listuser[i];
            if (currentUser.id !== currentID) {
                const divUser = document.createElement('div');
                divUser.id = 'userid' + currentUser.id;
                if (currentUser.nickname !== '') {
                    divUser.textContent = currentUser.nickname;
                } else {
                    divUser.textContent = currentUser.firstname + ' ' + currentUser.lastname;
                }

                if (currentUserOnline.includes(currentUser.id)) { // utilisateur en ligne
                    divUser.className = 'divOnlineUser';
                    onlineUser.appendChild(divUser);
                } else { // utilisateur hors ligne
                    divUser.className = 'divOfflineUser';
                    offlineUser.appendChild(divUser);
                }


                // Gestionnaire d'event si l'utilisateur clique sur cet utilisateur
                UserEventListener(divUser, currentUser, currentID, uuid);
            }
        }
    } else { // Offline
        for (const key in userDataStorage) {
            if (Object.hasOwnProperty.call(userDataStorage, key) && key != currentID) {
                const currentUser = userDataStorage[key].User;
                const divUser = document.createElement('div');
                divUser.id = 'userid' + currentUser.id;
                if (currentUser.nickname !== '') {
                    divUser.textContent = currentUser.nickname;
                } else {
                    divUser.textContent = currentUser.firstname + ' ' + currentUser.lastname;
                }
                divUser.className = 'divOfflineUser';
                offlineUser.appendChild(divUser);

                UserEventListener(divUser, currentUser, currentID, uuid);
            }
        }
    }
}

// Gestionnaire d'event si l'utilisateur clique sur cet utilisateur
function UserEventListener(divUser, currentUser, currentID, uuid) {
    divUser.addEventListener('click', async () => {
        // suppression du css de l'ancien click
        const olddeivselect = document.getElementById('userid').value
        if (olddeivselect != '') {
            document.getElementById('userid' + olddeivselect).classList.remove('active');
        }

        // mise à jour du button cliqué pour retirer 'New'
        if (currentUser.nickname !== '') {
            divUser.textContent = currentUser.nickname;
            document.getElementById('userid' + currentUser.id).textContent = currentUser.nickname;
        } else {
            document.getElementById('userid' + currentUser.id).textContent = currentUser.firstname + ' ' + currentUser.lastname;
        }

        // Gestion nouvelle div à mettre en avant
        document.getElementById('userid').value = currentUser.id;
        divUser.classList.add('active');
        if (online) {
            document.getElementById('inputform').style.display = 'flex';
        }
        // envoie du message d'update
        window.Electron.updateDiscussion(currentID, uuid);
    });
}

// permet d'afficher les messages dans le container message
export async function updateDiscussion(currentuserID, listMessages) {
    console.log('-> updateDiscussion');

    deleteDiscussionContents();

    if (listMessages.Messages !== undefined && listMessages.Messages !== null) {
        for (let i = 0; i < listMessages.Messages.length; i++) {
            const message = listMessages.Messages[i];

            let p = document.createElement('p');
            p.textContent = message.content;
            let span = document.createElement('span');
            span.textContent = message.date;
            p.appendChild(span);
            if (message.sender_id === currentuserID) {
                p.className = "sender";
            } else {
                p.className = "receiver";
            }
            document.getElementById('discussioncontent').appendChild(p);
        }
        document.getElementById('messagecontainer').scrollTop = messagecontainer.scrollHeight;
    }
}