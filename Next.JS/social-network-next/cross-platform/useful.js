async function fetchSendServer(option, dataToSend) { //getallmessage
    try {
        const response = await fetch("http://localhost:8080/crossplatform/" + option, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(dataToSend),
        });

        if (!response.ok) {
            throw new Error(`Erreur réseau: ${response.status}`);
        }

        const data = await response.json();
        // console.log('Data received (/' + option + '):', data);
        return data;

    } catch (error) {
        console.error('Data recovery error in fetchSendServer :\n', error);
        return;
    }
}

module.exports = { fetchSendServer };