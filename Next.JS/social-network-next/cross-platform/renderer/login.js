// fichier annexe
import { dislaynoneerror, displayblockerror, fetchSendServer } from "./useful.js";

// module npm

window.addEventListener('DOMContentLoaded', () => {

    window.Electron.ipcRenderer.on('connwebsocket', async (websocket) => {
        if (websocket) { // connexion réussi
            dislaynoneerror();
        } else { // erreur de connexion (pas d'internet par exemple)
            displayblockerror('Connexion error, please check internet');
        }
    });

    document.getElementById('loginButton').addEventListener('click', async (event) => {
        event.preventDefault();
        let loginuser = {};
        loginuser.email = document.getElementById('email').value;
        loginuser.password = document.getElementById('password').value;

        let data = await fetchSendServer('login', loginuser);
        if (data == false) {
            displayblockerror('Connexion error, please check internet');
        } else {
            dislaynoneerror();
            console.log('data : ', data.userData);
            if (data.userData.id !== 0) {
                window.Electron.login(data.userData);
            } else {
                displayblockerror('Wrong email, username or password');
            }
        }
    });

    document.getElementById('linkregister').addEventListener('click', (event) => {
        console.log('enter linkregister');
        event.preventDefault();
        // Emmission du message 'linkregister'
        window.Electron.linkregister();
    })
});
