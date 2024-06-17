export let online = false;

export function OnlineOfflineMode() {

    // mode online / offline au demarrage de l'application
    window.Electron.ipcRenderer.on('internetonline', async () => {
        online = true;
        document.getElementById('offline').classList.remove('active');
        if (!isNaN(parseInt(document.getElementById('userid').value, 10))) {
            document.getElementById('inputform').style.display = 'flex';
        }
    });
    window.Electron.ipcRenderer.on('internetoffline', async () => {
        online = false;
        document.getElementById('offline').classList.add('active');
        document.getElementById('inputform').style.display = 'none';
    });

    // Changement du mode online / offline
    window.Electron.notifyOnlineStatus((status) => {
        online = status;
        if (online) { // internetonline
            document.getElementById('offline').classList.remove('active');
            if (!isNaN(parseInt(document.getElementById('userid').value, 10))) {
                document.getElementById('inputform').style.display = 'flex';
            }
            window.Electron.internetonline();
        } else { // internetoffline
            document.getElementById('offline').classList.add('active');
            document.getElementById('inputform').style.display = 'none';
        }
    });
}