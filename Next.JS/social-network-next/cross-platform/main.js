// module npm
const { app, BrowserWindow, ipcMain, shell } = require('electron');
const path = require('path');
const { startWS } = require('./websocket.js');
const { fetchSendServer } = require('./useful.js');
let win;
let conn;
let online = false;

(async () => {
  const Store = (await import('electron-store')).default;
  const store = new Store();

  const createWindow = () => {
    win = new BrowserWindow({
      width: 1000,
      height: 600,
      webPreferences: {
        nodeIntegration: true,
        preload: path.join(__dirname, 'preload.js'),
      },
    });

    if (store && store.get('isLog') && store.get('dataStorage')) {
      const dataStorage = store.get('dataStorage');
      // console.log('dataStorage on isLog : ', dataStorage);
      win.loadFile(path.join(__dirname, 'renderer/home.html')).then(async () => {
        conn = startWS(dataStorage.userData, win);
        if (conn.onerror) {
          win.webContents.send('internetoffline');
          online = false;
        } else {
          win.webContents.send('internetonline');
          online = true;
        }
        win.webContents.send('userDataStorage', dataStorage);
      });
    } else {
      win.loadFile(path.join(__dirname, 'renderer/login.html'));
    }

    // chargement du status au demarrage de l'application (apres chargement complet de la fenêtre intérieur)
    win.webContents.on('did-finish-load', async () => {
      online = await win.webContents.executeJavaScript('navigator.onLine');
      if (online) {
        win.webContents.send('internetonline');
      } else {
        win.webContents.send('internetoffline');
      }
    });
  };

  app.whenReady().then(() => {
    createWindow();
    app.on('activate', () => {
      if (BrowserWindow.getAllWindows().length === 0) {
        createWindow();
      }
    });
  });

  app.on('window-all-closed', async () => {
    if (process.platform !== 'darwin') {
      const DataStorage = await store.get('dataStorage');
      if (DataStorage) {
        const msgSocket = {
          msg_type: "userleave",
          uuid: DataStorage.userData.uuid,
        }
        conn.send(JSON.stringify(msgSocket));
      }
      app.quit();
    }
  });

  // Réception du message 'linkregister'
  ipcMain.on('linkregister', () =>
    shell.openExternal('http://localhost:3000/register')
  );

  // Réception du message 'login'
  ipcMain.on('login', async (event, userData) => {
    store.set('isLog', true);
    const dataStorage = await fetchSendServer("getallmessage", { id: userData.id, uuid: userData.uuid });
    store.set('dataStorage', dataStorage);
    // console.log('dataStorage on login : ', dataStorage);

    await win.loadFile(path.join(__dirname, 'renderer/home.html')).then(() => {
      win.webContents.send('userDataStorage', dataStorage);
      conn = startWS(dataStorage.userData, win);
      if (conn.onerror) {
        win.webContents.send('internetoffline');
        online = false;
      } else {
        win.webContents.send('internetonline');
        online = true
      }
    });
  });

  // Réception du message 'logout'
  ipcMain.on('logout', () => {
    store.delete('islog');
    store.delete('dataStorage');
    win.loadFile(path.join(__dirname, 'renderer/login.html'));
  });

  // Réception de message 'chat' pour le Websocket
  ipcMain.on('chat', (currentEvent, msgData) => {
    conn.send(JSON.stringify(msgData));
  });

  // Réception de message 'update' pour l'actualisation de la liste des messages dans le stockage
  ipcMain.on('updateDiscussion', async (currentEvent, id, uuid) => {
    let dataStorage;
    if (online) {
      dataStorage = await fetchSendServer("getallmessage", { id: id, uuid: uuid });
      store.set('dataStorage', dataStorage);
    } else {
      dataStorage = await store.get('dataStorage');
    }
    await win.webContents.send('updateDiscussion', dataStorage);
  });

  // update de la liste des utilisateurs (message 'updateUserList')
  ipcMain.on('updateUserList', async () => {
    win.webContents.send('updateUserList');
  });

  // update de la liste des utilisateurs au retour d'internet (message 'internetonline')
  ipcMain.on('internetonline', async () => {
    win.webContents.send('internetonline');
  });

  // Moteur de recherche
  ipcMain.on('search', async (event, enter) => {
    const DataStorage = await store.get('dataStorage');
    win.webContents.send('search', DataStorage, enter);
  });

  // gestion du mode offline or online
  ipcMain.on('online-status-changed', (event, status) => {
    if (status) {
      win.webContents.send('internetonline');
      online = true;
    } else {
      win.webContents.send('internetoffline');
      online = false;
    }
  });

})();


// Fenetre Electron : code js lancer par npm start (execute le package.json) en type "communjs"
// Intérieur de la fenetre Electron : code js lancer par home.html ou login.html (pas d'execusion du package.json) en type "module"
// Chemin des données : intérieur de la fenêtre -> pont -> main.js -> emission websocket -> serveur -> reception websocket --> main.js --> pont -> intérieur de la fenêtre
// Les fetch peuvent être effectuer en direct : intérieur de la fenêtre -> serveur -> intérieur de la fenêtre


