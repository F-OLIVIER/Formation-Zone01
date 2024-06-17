export async function fetchSendServer(option, dataToSend) {
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
        dislaynoneerror();
        return data;

    } catch (error) {
        console.error('Data recovery error in fetchSendServer :\n', error);
        displayblockerror('Internal server error (code 500)');
        return false;
    }
}

export async function fetchServer(option) {
    try {
        const response = await fetch("http://localhost:8080/" + option, {
            method: 'GET',
            credentials: 'include'
        });

        if (!response.ok) {
            throw new Error(`Erreur réseau: ${response.status}`);
        }

        const data = await response.json();
        console.log('Data received (/' + option + '):', data);
        dislaynoneerror();
        return data;

    } catch (error) {
        displayblockerror('Internal server error (code 500)');
        return false;
    }
}

export function deleteDiscussionContents() {
    let discussioncontent = document.getElementById('discussioncontent');
    while (discussioncontent.firstChild) {
        discussioncontent.removeChild(discussioncontent.firstChild);
    }
}

export function displayblockerror(message) {
    const error = document.getElementById('error');
    error.textContent = message;
    error.style.display = "block";
}

export function dislaynoneerror() {
    const error = document.getElementById('error');
    error.textContent = '';
    error.style.display = "none";
}

// msg_type : 'listUser','getHistoryChat', 'chat', 'typing'
export function sendMsg(type, receiver_id, content, msg_type, target = undefined) {
    console.log("target send msg", target);

    let msgData = {
        id: 0,
        sender_id: 0,
        receiver_id: receiver_id,
        content: content,
        targets: target,
        date: '',
        msg_type: msg_type,
        is_typing: false,
    }

    switch (type) {
        case 'chat':
            window.Electron.chat(msgData);
            break;

        case 'typing':
            msgData.is_typing = true;
            window.Electron.typing(msgData);
            break;
        default:
            break;
    }
};

export function getCurrentDateTimeString() {
    const now = new Date();

    const month = String(now.getMonth() + 1).padStart(2, '0');
    const day = String(now.getDate()).padStart(2, '0');
    const year = now.getFullYear();

    const hours = String(now.getHours()).padStart(2, '0');
    const minutes = String(now.getMinutes()).padStart(2, '0');
    const seconds = String(now.getSeconds()).padStart(2, '0');

    const dateString = `${month}-${day}-${year} ${hours}:${minutes}:${seconds}`;

    return dateString;
}

