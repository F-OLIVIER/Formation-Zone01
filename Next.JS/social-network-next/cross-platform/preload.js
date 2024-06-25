// https://www.electronjs.org/docs/latest/tutorial/process-model#preload-scripts

// module npm
const { contextBridge, ipcRenderer } = require('electron');

// pont avec l'intérieur de la fenetre electron
contextBridge.exposeInMainWorld('Electron', {
    // Transmission du message 'linkregister'
    linkregister: () => {
        ipcRenderer.send('linkregister');
    },

    // Transmission du message 'login'
    login: (userinfo) => {
        ipcRenderer.send('login', userinfo);
    },

    // Transmission du message 'logout'
    logout: () => {
        ipcRenderer.send('logout');
    },

    // Transmission du message pour le websocket 'chat'
    chat: (msgdata) => {
        ipcRenderer.send('chat', msgdata);
    },

    // Transmission du message 'updateDiscussion' pour l'actualisation de la liste des messages dans le stockage
    updateDiscussion: (id, uuid) => {
        ipcRenderer.send('updateDiscussion', id, uuid);
    },
    
    updateUserList: () => {
        ipcRenderer.send('updateUserList');
    },
    
    internetonline: () => {
        ipcRenderer.send('internetonline');
    },

    // Transmission de data à l'intérieur de la fenétre (emission dans main.js et reception dans la fenêtre intérieur)
    ipcRenderer: {
        // permet d'envoyer des data dans la fenetre
        send: (channel, data) => ipcRenderer.send(channel, data),
        // permet de recevoir des data de la fenetre
        on: (channel, callback) => ipcRenderer.on(channel, (event, ...args) => callback(...args))
    },

    // gestion des mode online et offline
    notifyOnlineStatus: (callback) => {
        window.addEventListener('online', () => {
            callback(navigator.onLine);
        });
        window.addEventListener('offline', () => {
            callback(navigator.onLine);
        });
    },
    sendOnlineStatus: (status) => {
        ipcRenderer.send('online-status-changed', status);
    },

    // moteur de recherche
    search: (enter) => {
        ipcRenderer.send('search', enter);
    },
});
