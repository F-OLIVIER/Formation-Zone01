const { ipcRenderer } = require('electron');
const WebSocket = require('ws');

function startWS(userstorage, win) {
    const conn = new WebSocket("ws://localhost:8080/ws");

    // gestion des erreurs websocket.
    conn.onerror = function (err) {
        console.error('Error WebSocket : \n', err);
    };

    // Ouverture connexion websocket.
    conn.onopen = function () {
        console.log("WebSocket connection is open");
        conn.send(JSON.stringify(userstorage));
    };

    // Fermeture connexion websocket.
    conn.onclose = function () {
        console.log("WebSocket connection is closed");
    };

    // En fonction du type de message on exécute une opération différente.
    conn.onmessage = async function (evt) {
        try {
            var data = JSON.parse(evt.data);

            switch (data.msg_type) {
                case "online":
                    // console.log("new user online in websocket : ", data)
                    win.webContents.send('online');
                    break;

                case "userleave":
                    // console.log("user disconect to websocket : ", data)
                    win.webContents.send('userleave', data);
                    break;

                case "msg":
                    // console.log("new message websocket : ", data)
                    win.webContents.send('new_msg', data);
                    break;

                default:
                    console.log("Other websocket (non traité) : ", data)
                    break;
            }

        } catch (err) {
            console.error('Error onmessage WebSocket : \n', err);
        }
    };

    return conn
}

module.exports = { startWS };
