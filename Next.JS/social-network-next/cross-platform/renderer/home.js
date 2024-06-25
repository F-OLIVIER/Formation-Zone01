// fichier annexe
import { fetchSendServer, getCurrentDateTimeString } from "./useful.js";
import { updateDiscussion, updateUserList } from "./updateHTML.js";
import { OnlineOfflineMode, online } from "./onlineofflinemode.js";
import { chatEmojis } from "./emojis.js";
import { Searchbar } from "./searchbar.js";

let timerThrottle = 0;

function home() {

    OnlineOfflineMode();

    // Ajout des emojis
    chatEmojis();

    // récupération de l'utilisateur stocker
    window.Electron.ipcRenderer.on('userDataStorage', async (DataStorage) => {
        console.log('DataStorage : ', DataStorage);
        const userDataStorage = DataStorage.userData;
        // console.log('userDataStorage : ', userDataStorage);
        const usersmessages = DataStorage.usersmessages;
        // console.log('usersmessages : ', usersmessages);
        // console.log('online : ', online)

        // Button logout
        document.getElementById('logout').addEventListener('click', async () => {
            const msgSocket = {
                sender_id: userDataStorage.id,
                msg_type: "userleave",
                uuid: userDataStorage.uuid,
            }
            await window.Electron.chat(msgSocket);
            await window.Electron.logout();
        });

        document.getElementById('userconnected').innerHTML = `${userDataStorage.firstname} ${userDataStorage.lastname}</br> (${userDataStorage.nickname})`;

        // mise à jour des data lorsque internet reviens
        window.Electron.ipcRenderer.on('internetonline', async () => {
            window.Electron.updateDiscussion(userDataStorage.id, userDataStorage.uuid);
            window.Electron.updateUserList(usersmessages, userDataStorage.id, userDataStorage.uuid);
        });

        // Utilisateur qui se connecte
        window.Electron.ipcRenderer.on('online', async () => {
            console.log('new online user (home.js)')
            notification(userDataStorage.id, userDataStorage.uuid, 'online', userDataStorage.id);
            window.Electron.updateUserList(usersmessages, userDataStorage.id, userDataStorage.uuid);
        });
        // Utilisateur qui se déconnecte
        window.Electron.ipcRenderer.on('userleave', async (msgdata) => {
            if (online) {
                // console.log('offline user (home.js): ', msgdata)
                notification(userDataStorage.id, userDataStorage.uuid, 'offline', userDataStorage.id);
                window.Electron.updateUserList(usersmessages, userDataStorage.id, userDataStorage.uuid);
            }
        });

        setInterval(() => {
            // mise a jour du stockage pour delog l'utilisateur si besoin
            window.Electron.updateDiscussion(userDataStorage.id, userDataStorage.uuid);
            // mise à jour de la liste des utilisateurs
            window.Electron.updateUserList(usersmessages, userDataStorage.id, userDataStorage.uuid)
        }, 8000);

        // gestion de l'input
        const inputmessage = document.getElementById('inputmessage');
        // Détection quand 'enter' et saisie au clavier
        inputmessage.addEventListener('keypress', async function (event) {
            if (online && event.key === 'Enter' && event.target.value != '') {
                const now = new Date();
                if (now - timerThrottle > 500) {
                    timerThrottle = now;

                    const msgSocket = {
                        sender_id: userDataStorage.id,
                        receiver_id: parseInt(document.getElementById('userid').value, 10),
                        content: event.target.value,
                        date: getCurrentDateTimeString(),
                        msg_type: "msg",
                        uuid: userDataStorage.uuid,
                    }
                    await window.Electron.chat(msgSocket);
                    inputmessage.value = '';
                    await window.Electron.updateDiscussion(userDataStorage.id, userDataStorage.uuid);
                }
            }
        });

        // réception d'un nouveau message
        window.Electron.ipcRenderer.on('new_msg', async (msgdata) => {
            // console.log('new_msg : ', msgdata);
            notification(msgdata.sender_id, userDataStorage.uuid, ' sent you a message', userDataStorage.id);
            window.Electron.updateDiscussion(userDataStorage.id, userDataStorage.uuid);
        });

        window.Electron.ipcRenderer.on('updateUserList', async () => {
            await updateUserList(usersmessages, userDataStorage.id, userDataStorage.uuid);
        });

        if (!online) {
            await updateUserList(usersmessages, userDataStorage.id, userDataStorage.uuid);
        }
    });

    // réception d'une demande d'update de l'encart de discussion
    window.Electron.ipcRenderer.on('updateDiscussion', async (dataStorageupdate) => {
        // vérification de l'utilisateur
        console.log('dataStorageupdate :\n', dataStorageupdate.userData.uuid)
        if (dataStorageupdate.userData.uuid.trim() == "") {
            await window.Electron.logout();
            return
        }
        // console.log('-> updateDiscussion valid');
        const usersmessages = dataStorageupdate.usersmessages
        const id = parseInt(document.getElementById('userid').value, 10);
        if (usersmessages[id] && usersmessages[id].Messages !== undefined) {
            await updateDiscussion(dataStorageupdate.userData.id, usersmessages[id]);
        }
    });

    // moteur de rechercher
    document.getElementById('inputsearch').addEventListener('input', async function (event) {
        // Vérifie si l'événement keypress ou supp ou delete text a été déclenché
        if (event.inputType === 'insertText' || event.inputType === 'deleteContentBackward' || event.inputType === 'deleteContentForward') {
            if (event.inputType === 'insertText' && event.data === '\n') { // recherche avec enter
                window.Electron.search(true);
            } else { // recherche sans enter
                window.Electron.search(false);
            }
        }
    });
    window.Electron.ipcRenderer.on('search', async (DataStorage, enter) => {
        Searchbar(DataStorage.usersmessages, enter);
    });
}

home();

async function notification(useridtofetch, uuid, notifMessage, currentID = 0) {
    const data = await fetchSendServer('getuser', { id: useridtofetch, uuid: uuid });
    console.log('data notification : ', useridtofetch, data.userData)
    if (currentID !== data.userData.id) {
        const currentUserIdOpen = parseInt(document.getElementById('userid').value, 10);
        // modification du bouton utilisateur
        if (currentUserIdOpen != data.userData.id) {
            document.getElementById('userid' + data.userData.id).textContent += " 🆕"
        }

        // affichage de la notification
        const notif = document.getElementById('notif');
        if (data.userData.nickname != '') {
            notif.textContent = data.userData.nickname + notifMessage;
        } else {
            notif.textContent = data.userData.firstname + " " + data.userData.lastname + notifMessage;
        }
        notif.classList.add('active');
        // Désactivation de la notification apres x secondes
        setTimeout(async function () {
            notif.classList.remove('active');
        }, 3000);
    }
}

