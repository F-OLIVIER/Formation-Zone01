import toast from "react-hot-toast";


export var firstId = 512;
export var offset = 0;
export let counter = 0
export var unread = []
export var currUsername = ""
export var allUsers = []
export var online = []
export var currComments = []
export let currentScrollPos;
export let lastFetchedId;
export let value;
export let closeChat;
export let debouncedScrollHandler;


export var conn;

async function getData(url = '') {
    const response = await fetch(url, {
        method: 'GET',
        credentials: 'include'
    })

    return response.json()
}


export async function getUsers() {
    await getData('http://localhost:8080/user')
        .then(value => {
            allUsers = value
        }).catch(err => {
            console.log(err)
        })
}

export async function updateUsers(currId) {
    await getData('http://localhost:8080/chat?user_id=' + currId)
        .then(value => {
            var newUsers = []
            if (value.user_ids != null) {
                newUsers = value.user_ids.map((i) => {
                    return allUsers.filter(u => u && u.id == i)[0]
                })
            }
            let otherUsers = allUsers.filter(x => !newUsers.includes(x));
            allUsers = newUsers.concat(otherUsers);
            createUsers(allUsers, conn, currId)
        }).catch(err => {
            console.log(err)
        })
}


export function startWS(currId) {
    console.log("call startWS JS")
    console.log(currId)
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://localhost:8080/ws");

        conn.onopen = function () {
            // Ouverture connexion websocket.
            console.log("WebSocket connection is open");
        };

        conn.onclose = function (evt) {
            // Fermeture connexion websocket.
            console.log("WebSocket connection is closed");
        };


        // En fonction du type de message on exécute une opération différente.
        conn.onmessage = function (evt) {
            var data = JSON.parse(evt.data);
            console.log(data);
            if (data.msg_type === "post") {
                // Nouveau post, on notifie les autres utilisateurs d'un nouveau post, pas géré pour les commentaires.
                console.log("new post")
                toast(
                    <span>
                        New post !Click <a href="/" onClick={() => window.location.reload()}>here</a>
                    </span>,
                    {
                        duration: 4000,
                        position: 'top-center',
                        icon: '👏',
                    }
                );


            } else if (data.msg_type === "msg") {
                // Message chat, on regarde si on est l'émetteur ou le destinataire, on créé la div correspondante, et on "append" le message.
                var senderContainer = document.createElement("div");
                senderContainer.className = (data.sender_id == currId) ? "sender-container" : "receiver-container";
                var sender = document.createElement("div");
                sender.className = (data.sender_id == currId) ? "sender" : "receiver";
                sender.innerText = data.content;
                var date = document.createElement("div");
                date.className = "chat-time";
                date.innerText = data.date.slice(0, -3);
                appendLog(senderContainer, sender, date);

                if (data.sender_id == currId) {
                    return;
                }

                let unreadMsgs = unread.filter((u) => {
                    id = data.sender_id;
                    return u[0] == id;
                });


                const chatWrapper = document.querySelector('.chat-wrapper');
                if (!chatWrapper || chatWrapper.style.display !== 'flex') {
                    toast(
                        <span>
                        New message from {data.sender_id}!
                    </span>,
                        {
                            duration: 4000,
                            position: 'top-center',
                            icon: '💻',
                            style: {
                                backgroundColor: 'rgba(157,161,157,0.5)', color: 'white',
                            },
                        }
                    );
                }

                if (chatWrapper && chatWrapper.style.display === 'none') {
                    if (unreadMsgs.length == 0) {
                        unread.push([data.sender_id, 1]);
                    } else {
                        unreadMsgs[0][1] += 1;
                    }
                }
                updateUsers(currId);

            } else if (data.msg_type === "online") {
                // Connexion d'un utilisateur, on met à jour des liste des contacts, et les statuts.
                online = data.user_ids;
                getUsers()
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
}

// Fonction d'envoi d'un message via le WS
export function sendMsg(conn, rid, msg, msg_type) {

    // Check si WS ouverte.
    if (!conn) {
        return false;
    }

    // Check message vide.
    if (!msg.value) {
        return false;
    }

    let msgData = {
        id: 0,
        sender_id: 0,
        receiver_id: rid,
        content: msg.value,
        date: '',
        msg_type: msg_type,
        is_typing: false
    }

    conn.send(JSON.stringify(msgData))
    msg.value = "";
    return false;
};


// Création de la liste de contacts
export async function createUsers(userdata, conn, currId) {
    await sleep(1000);
    const offlineUsers = document.querySelector('.offline-users');
    if (offlineUsers) {
        offlineUsers.innerHTML = ""
    }
    if (userdata == null) {
        return
    }

    userdata.map(({id, nickname}) => {

        // Pour ne pas s'afficher soit même
        if (id == currId) {
            return
        }
        var user = document.createElement("div");
        user.className = "user"
        user.setAttribute("id", ('id' + id))

        // Répartition des users selon leur statut.


        var chatusername = document.createElement("p");
        chatusername.innerText = nickname
        user.appendChild(chatusername)
        if (offlineUsers) {
            offlineUsers.appendChild(user)
        }
        /*         var msgNotification = document.createElement("div");
                msgNotification.className = "msg-notification"
                msgNotification.innerText = 1
                user.appendChild(msgNotification)
        
                let unreadMsgs = unread.filter((u) => {
                    return u[0] == id
                })
        
                if (unreadMsgs.length != 0 && unreadMsgs[0][1] != 0) {
                    const msgNotif =  document.getElementById('id'+id).querySelector('.msg-notification');
                    msgNotif.style.opacity = "1"
                    msgNotif.innerText = unreadMsgs[0][1]
                    document.getElementById('id'+id).style.fontWeight = "900"
                }  */

        // En cas de click sur un utilisateur, on check la DB message et on ouvre une fenêtre de chat.
        user.addEventListener("click", function (e) {
            GetElementToOpenChat(id, user, currId)
        });
    })
}

export function GetElementToOpenChat(id, user, currId) {
    resetScroll();
    if (typeof conn === "undefined") {
        // Protection si problème de WS.
        return;
    }
    // On récupère les logs si ils existent.
    let resp = getData('http://localhost:8080/message?receiver=' + id + '&firstId=' + firstId);
    resp.then(value => {
        let rIdStr = user.getAttribute("id");
        const regex = /id/i;
        const rId = parseInt(rIdStr.replace(regex, ''));
        if (value && value.length > 0) {
            const lastIndex = value.length - 1;
            firstId = value[lastIndex].id;
            lastFetchedId = firstId;
        }
        counter = 0;
        // Ouverture d'une fenêtre de chat.
        OpenChat(rId, conn, value, currId, firstId);
    }).catch();
}

let log;

if (typeof document !== 'undefined') {
    // Access DOM elements only if document is defined

    log = document.querySelector(".chat")
}

export {log};


export function appendLog(container, msg, date, prepend = false) {
    // Retrieve the log element
    var log = document.querySelector(".chat");
    if (!log) {
        // Chat log element not found, cannot append log
        return;
    }

    // Check if document is defined
    if (!document) {
        // Document is not defined, cannot append log
        return;
    }

    var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
    if (prepend) {
        if (log.firstChild) {
            log.insertBefore(container, log.firstChild);
        } else {
            log.appendChild(container);
        }
    } else {
        log.appendChild(container);
    }
    container.appendChild(msg);
    msg.appendChild(date);

    // Scroll to bottom regardless of prepend or append
    log.scrollTop = log.scrollHeight;
}

export function resetScroll() {
    var chatBox = document.querySelector(".chat");
    if (chatBox && typeof debouncedScrollHandler === "function") {
        chatBox.removeEventListener("scroll", debouncedScrollHandler);
    }
    firstId = 512;
    offset = 0;
    currentScrollPos = 0;
    lastFetchedId = null;
    value = null;
    closeChat = true;
}


// Fonction active si ouverture d'un chat.
export function OpenChat(rid, conn, data, currId, firstId) {
    document.getElementById('id' + rid).style.fontWeight = "400";
    var chatBox = document.querySelector('.chat');

    chatBox.innerHTML = '';

    for (var i = 0; i < unread.length; i++) {
        if (unread[i][0] == rid) {
            unread[i][1] = 0;
        }
    }

    let oldElem = document.querySelector(".send-wrapper");
    let newElem = oldElem.cloneNode(true);
    oldElem.parentNode.replaceChild(newElem, oldElem);

    document.querySelector(".chat-user-username").innerText = allUsers.filter(u => {
        return u.id == rid;
    })[0].nickname;

    document.querySelector(".chat-wrapper").style.display = "flex";
    var msg = document.getElementById("chat-input");

    // Fermeture du chat.
    document.querySelector(".close-chat").addEventListener("click", function () {
        document.querySelector(".chat-wrapper").style.display = "none";
        resetScroll();
    });


    // Envoi du message si click.
    document.querySelector("#send-btn").addEventListener("click", function () {
        // On fait un envoi msg via WS avec le type "msg" (donc message chat)

        sendMsg(conn, rid, msg, 'msg');

        offset = null;
        firstId = firstId + 10;
        let resp = getData('http://localhost:8080/message?receiver=' + rid + '&firstId=512' + '&offset=10');
        resp.then(value => {
            if (value && value.length > 0) {
                const lastIndex = value.length - 1;
                firstId = value[lastIndex].id;
                lastFetchedId = firstId;
            }
            OpenChat(rid, conn, value, currId, firstId)
            chatBox.scrollTop = chatBox.scrollHeight;
        }).catch();
    });

    document.querySelector(".emoji-icon").addEventListener("click", function () {
        var chatBox = document.querySelector('.chat');
        chatBox.scrollTop = chatBox.scrollHeight;
    });

    // Même chose pour "entrée".
    /*     document.querySelector("#chat-input").addEventListener("keydown", function (event) {
            if (event.keyCode === 13) {

                offset = 0;
                sendMsg(conn, rid, msg, 'msg');
                firstId = firstId + 10;
                let resp = getData('http://localhost:8080/message?receiver=' + rid + '&firstId=512' + '&offset=10');
                resp.then(value => {
                    if (value && value.length > 0) {
                        const lastIndex = value.length - 1;
                        firstId = value[lastIndex].id;
                        lastFetchedId = firstId;
                    }
                    OpenChat(rid, conn, value, currId, firstId)
                    chatBox.scrollTop = chatBox.scrollHeight;
                }).catch();
            }
        }); */

    var offset = 10;
    var lastFetchedId = null;

    // Fonction pour charger les messages 10/10.
    debouncedScrollHandler = debounce(function () {

        if (chatBox.scrollTop === 0) {
            let resp = getData('http://localhost:8080/message?receiver=' + rid + '&firstId=' + firstId + '&offset=' + offset);
            resp.then(value => {
                if (value == null) {
                    return;
                }
                if (value.length == 0) {
                    return;

                }
                value = value.filter(message => message.id !== lastFetchedId);


                if (value.length > 0) {
                    const lastIndex = value.length - 1;
                    firstId = value[lastIndex].id;
                    lastFetchedId = firstId;
                }
                currentScrollPos = chatBox.scrollHeight - chatBox.scrollTop;
                CreateMessages(value, currId);
                var newScrollPos = chatBox.scrollHeight - currentScrollPos;
                chatBox.scrollTop = newScrollPos;

                offset += value.length;
            }).catch();
        }
    }, 300);

    chatBox.addEventListener("scroll", debouncedScrollHandler);

    function debounce(func, delay) {
        let timer;
        return function () {
            clearTimeout(timer);
            timer = setTimeout(func, delay);
        };
    }

    if (data == null) {
        return;
    }
    CreateMessages(data, currId);
}

export function CreateMessages(data, currId) {
    const chatBox = document.querySelector('.chat');

    // Iterate over the data array from the beginning
    for (let i = 0; i < data.length; i++) {
        const {id, sender_id, content, date} = data[i];

        // Check for duplicates
        if (document.getElementById(`message-${id}`)) {
            continue;
        }
        // Create message elements
        const messageContainer = document.createElement("div");
        messageContainer.className = sender_id === currId ? "sender-container" : "receiver-container";

        const message = document.createElement("div");
        message.className = sender_id === currId ? "sender" : "receiver";
        message.innerText = content;

        const messageDate = document.createElement("div");
        messageDate.className = sender_id === currId ? "chat-time-left" : "chat-time";
        messageDate.innerText = date.slice(0, -3);

        messageContainer.id = `message-${id}`;

        // Append message elements to the chat box
        appendLog(messageContainer, message, messageDate, true);
    }

    // Scroll to the bottom of the chat box
    chatBox.scrollTop = chatBox.scrollHeight;
}


function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}