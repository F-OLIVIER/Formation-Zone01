
export function SendInWebSocket(dataToSendServerGO) {
    // Établissement de la connexion WebSocket
    let socket = new WebSocket('ws://localhost:5501/ws-endpoint');

    // Gestionnaire d'événements pour l'ouverture de la connexion
    socket.addEventListener('open', function () {
        socket.send(JSON.stringify(dataToSendServerGO)); // Envoi des données au serveur go
    });
    // Gestionnaire d'événements pour les erreurs
    socket.addEventListener('error', function (event) {
        console.error('Erreur WebSocket :', event);
    });

    // // Gestionnaire d'événements pour la réception des messages
    // socket.addEventListener('message', function (message) {
    //     console.log('message :', message.data);
    // });

    // Gestionnaire d'événements pour la fermeture de la connexion
    socket.addEventListener('close', function () {
        console.log('WebSocket connexion close');
    });
}