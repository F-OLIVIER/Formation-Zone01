import toast from "react-hot-toast";
import { fetchNotification } from "@/services/useFetchNotif";



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

export var currentOpenChatId = null;
export var currentOpenChatType = null; // "user" or "group"

export function startWS(currId, setNotifications, router) {
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
        conn.onmessage = async function (evt) {
            var data = JSON.parse(evt.data);
            var currentURL = window.location.href;
            console.log("Data websocket", data);
            //console.log(data.receiver_id)

            await fetchNotification(currId, setNotifications);
            if (data.msg_type === "post") {
                console.log("new post")
                console.log(data)
                if (data.targets && data.targets.includes(currId)) {
                    toast(
                        <span>
                            New post! Click <a className="custom-link" onClick={() => router.push('/home')}>here</a>
                        </span>,
                        {
                            duration: 4000,
                            position: 'top-center',
                            icon: '👏',
                        }
                    );
                }
            } else if (data.msg_type === "join_request") {
                if (data.targets && data.targets.includes(currId)) {
                    toast(
                        <span>
                            New join request! Click <a className="custom-link" onClick={() => router.push('/notif')}>here</a>
                        </span>,
                        {
                            duration: 4000,
                            position: 'top-center',
                            icon: '❓',
                        }
                    );
                }
            } else if (data.msg_type === "group") {
                //console.log("new group")
                if (data.targets && data.targets.includes(currId)) {
                    toast(
                        <span>
                            Your are invited to a new group ! Click <a className="custom-link"
                                onClick={() => router.push('/notif')}>here</a>
                        </span>,
                        {
                            duration: 4000,
                            position: 'top-center',
                            icon: '🫂',
                        }
                    );
                }
            } else if (data.msg_type === "follow") {
                console.log("new follow")
                //await Target(currId);
                if (data.targets && data.targets.includes(currId)) {

                    toast(
                        <span>
                            You have a new follow ! Click <a className="custom-link"
                                onClick={() => router.push('/notif')}>here</a>
                        </span>,
                        {
                            duration: 4000,
                            position: 'top-center',
                            icon: '🫵🏻',
                        }
                    );
                }
            } else if (data.msg_type === "stop_follow") {
                console.log("stop follow")
                if (data.targets && data.targets.includes(currId)) {
                    toast(
                        <span>
                            One follow stopped ! Click <a className="custom-link"
                                onClick={() => router.push('/user')}>here</a>
                        </span>,
                        {
                            duration: 4000,
                            position: 'top-center',
                            icon: '👀',
                        }
                    );
                }
            } else if (data.msg_type === "cancel_follow") {
                console.log("cancel follow")
                if (data.targets && data.targets.includes(currId)) {
                    toast(
                        <span>
                            One follow canceled ! Click <a className="custom-link"
                                onClick={() => router.push('/user')}>here</a>
                        </span>,
                        {
                            duration: 4000,
                            position: 'top-center',
                            icon: '❌',
                        }
                    );
                }

            } else if (data.msg_type === "msg" || data.msg_type === "groupmsg") {
                if (data.msg_type === "msg" && currentOpenChatType === "user" && currentOpenChatId === data.sender_id) {
                    displayChatMessage(data, currId);
                } else if (data.msg_type === "groupmsg" && currentOpenChatType === "group" && currentOpenChatId === data.receiver_id) {
                    displayChatMessage(data, currId);
                } else {
                    toast(
                        <span>
                            You have a new message ! Click <a className="custom-link" onClick={() => router.push('/chat')}>here</a>
                        </span>,
                        {
                            duration: 4000,
                            position: 'top-center',
                            icon: '📬',
                        }
                    );
                }


            } else if (data.msg_type === "online") {
                // Connexion d'un utilisateur, on met à jour des liste des contacts, et les statuts.
                online = data.user_ids;
                getUsers()
            } else if (data.msg_type === "comment") {
                // Nouveau commentaire, on notifie les autres utilisateurs d'un nouveau commentaire.
                console.log("comment now")
                if (data.targets && data.targets.includes(currId)) {
                    toast(
                        <span>
                            New Comment on your post !Click <a className="custom-link"
                                onClick={() => router.push('/home')}>here</a>
                        </span>,
                        {
                            duration: 4000,
                            position: 'top-center',
                            icon: '📩',
                        }
                    );
                }
            } else if (data.msg_type === "") {

            }
        }
            ;
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
}

function displayChatMessage(data, currId) {
    var senderContainer = document.createElement("div");
    senderContainer.className = (data.sender_id == currId) ? "sender-container" : "receiver-container";
    var sender = document.createElement("div");
    sender.className = (data.sender_id == currId) ? "sender" : "receiver";
    sender.innerText = data.content;
    var date = document.createElement("div");
    date.className = "chat-time";
    date.innerText = data.date.slice(0, -3);
    appendLog(senderContainer, sender, date);

    updateUsers(currId);
}

let target = [];
// Fonction d'envoi d'un message via le WS
export function sendMsg(conn, currId, msg, msg_type, target = undefined) {
    // Check si WS ouverte.
    if (!conn) {
        return false;
    }

    // Check message vide.
    if (!msg.value) {
        return false;
    }

    let rid = (currentOpenChatType === "user") ? currentOpenChatId : (currentOpenChatType === "group") ? currentOpenChatId : null;

    if (!rid) {
        return false;
    }

    console.log(rid)


    let msgData = {
        id: 0,
        sender_id: currId,
        receiver_id: rid,
        content: msg.value,
        targets: target,
        date: '',
        msg_type: msg_type,
        is_typing: false,
    }

    console.log(msgData)

    conn.send(JSON.stringify(msgData))
    msg.value = "";
    return false;
};


// Création de la liste de contacts
export async function createUsers(userdata, conn, currId) {
    await sleep(1000);
    const offlineUsers = document.querySelector('.offline-users');
    const groupList = document.querySelector('.group-list');

    if (offlineUsers != null) {
        offlineUsers.innerHTML = ""
    }
    if (groupList != null) {
        groupList.innerHTML = ""
    }
    const list = await fetchUsersFromAPI(currId);

    if (userdata == null) {
        return
    }

    userdata.map(({ id, nickname, privateprofile }) => {

        // Pour ne pas s'afficher soit même
        if (id == currId) {
            return
        }
        if (privateprofile == 0) {
            if (!list.listfollowers && !list.listfollowings) {
                return
            }
            if (list.listfollowers) {
                if (list.listfollowers.some(follower => follower.id === id)) {
                } else {
                    return
                }
            }
            if (list.listfollowings) {
                if (list.listfollowings.some(follower => follower.id === id)) {
                } else {
                    return
                }
            }
        }

        var user = document.createElement("div");
        user.className = "user"
        user.setAttribute("id", ('id' + id))

        // Répartition des users selon leur statut.


        var chatusername = document.createElement("p");
        chatusername.innerText = nickname
        user.appendChild(chatusername)
        if (offlineUsers != null) {
            offlineUsers.appendChild(user)
        }
        user.addEventListener("click", function (e) {
            GetElementToOpenChat(id, user, currId)
        });
    })

    const group = await getGroupForChat(currId);

    if (group.groups == null) {
        return
    }

    group.groups.map(({ IdGroup, Title }) => {


        var user = document.createElement("div");
        user.className = "user"
        user.setAttribute("id", ('id' + IdGroup))


        var chatusername = document.createElement("p");
        chatusername.innerText = Title
        user.appendChild(chatusername)
        if (groupList != null) {
            groupList.appendChild(user)
        }
        user.addEventListener("click", function (e) {
            GetElementToOpenChatGroup(IdGroup, currId, Title)
        });
    })

}

export function GetElementToOpenChat(id, user, currId) {
    resetScroll();
    document.querySelector(".chat-wrapper").style.display = "none";

    if (typeof conn === "undefined") {
        // Protection si problème de WS.
        return;
    }

    currentOpenChatId = id;
    currentOpenChatType = "user";



    // On récupère les logs si ils existent.
    let resp = getData('http://localhost:8080/message?receiver=' + id + '&firstId=' + firstId);
    resp.then(value => {
        if (value && value.length > 0) {
            const lastIndex = value.length - 1;
            firstId = value[lastIndex].id;
            lastFetchedId = firstId;
        }
        counter = 0;
        console.log(id)
        // Ouverture d'une fenêtre de chat.
        OpenChat(id, conn, value, currId, firstId);
    }).catch();
}

export function GetElementToOpenChatGroup(id, currId, Title) {
    resetScroll();
    if (typeof conn === "undefined") {
        // Protection si problème de WS.
        return;
    }

    currentOpenChatId = id;
    currentOpenChatType = "group";

    // On récupère les logs si ils existent.
    let resp = getData('http://localhost:8080/messagegroup?group=' + id + '&firstId=' + firstId);
    resp.then(value => {


        // Ouverture d'une fenêtre de chat.
        OpenChatGroup(id, conn, value, currId, firstId, Title);
    }).catch();
}

let log;

if (typeof document !== 'undefined') {
    // Access DOM elements only if document is defined

    log = document.querySelector(".chat")
}

export { log };


export function appendLog(container, msg, date) {
    const log = document.querySelector(".chat");
    if (!log) {
        return;
    }

    const newMessageDate = parseDate(date.innerText);
    const existingMessages = [...log.children];

    let inserted = false;
    for (let i = 0; i < existingMessages.length; i++) {
        const existingMessage = existingMessages[i];
        const existingMessageDate = parseDate(existingMessage.querySelector('.chat-time, .chat-time-left').innerText);

        if (newMessageDate < existingMessageDate) {
            log.insertBefore(container, existingMessage);
            inserted = true;
            break;
        }
    }

    if (!inserted) {
        log.appendChild(container);
    }

    container.appendChild(msg);
    msg.appendChild(date);

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

    console.log("rid", rid)
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
        if (msg.value.trim() === "") {
            return
        }
        sendMsg(conn, currId, msg, 'msg');

        offset = null;
        firstId = firstId + 10;
        console.log("rid", rid)
        let resp = getData('http://localhost:8080/message?receiver=' + rid + '&firstId=512' + '&offset=10');
        resp.then(value => {
            if (value && value.length > 0) {
                const lastIndex = value.length - 1;
                firstId = value[lastIndex].id;
                lastFetchedId = firstId;
            }
            console.log(value)
            CreateMessages(value, currId);
            chatBox.scrollTop = chatBox.scrollHeight;
        }).catch();
    });

    document.querySelector(".emoji-icon").addEventListener("click", function () {
        var chatBox = document.querySelector('.chat');
        chatBox.scrollTop = chatBox.scrollHeight;
    });


    var offset = 10;
    var lastFetchedId = null;

    // Fonction pour charger les messages 10/10.
    /*     debouncedScrollHandler = debounce(function () {
    
            if (chatBox.scrollTop === 0) {
                let resp = getData('http://localhost:8080/message?receiver=' + rid + '&firstId=' + firstId + '&offset=' + offset);
                resp.then(value => {
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
        } */

    if (data == null) {
        return;
    }
    CreateMessages(data, currId);
}
function parseDate(dateString) {
    return new Date(dateString.replace(/\./g, '/')); // Replace dots with slashes for compatibility
}

function compareMessages(a, b) {
    return parseDate(a.date) - parseDate(b.date);
}

export function CreateMessages(data, currId, otherId) {
    // Check if data is null or undefined
    if (!data || !Array.isArray(data)) {
        console.error('Data is null or not an array:', data);
        return;
    }

    const chatBox = document.querySelector('.chat');

    // Clear the chat box to prevent duplication
    chatBox.innerHTML = '';

    const messageElements = [];

    console.log('data : ', data);

    // Create message elements
    data.forEach(({ id, sender_id, content, date, sender_nickname }) => {
        const messageContainer = document.createElement("div");
        messageContainer.className = sender_id === currId ? "sender-container" : "receiver-container";
        messageContainer.id = `message-${id}`;

        const message = document.createElement("div");
        message.className = sender_id === currId ? "sender" : "receiver";
        message.innerText = content;

        const messageDate = document.createElement("div");
        messageDate.className = sender_id === currId ? "chat-time-left" : "chat-time";
        if (sender_nickname != "") {
            messageDate.innerText = date.slice(0, -3) + ' FROM ' + sender_nickname;
        } else {
            messageDate.innerText = date.slice(0, -3);
        }

        messageContainer.appendChild(message);
        message.appendChild(messageDate);

        messageElements.push({ element: messageContainer, date: date });
    });

    // Sort messages by date
    messageElements.sort(compareMessages);

    // Append sorted messages to chat box
    messageElements.forEach(({ element }) => chatBox.appendChild(element));

    // Scroll to the bottom of the chat box
    chatBox.scrollTop = chatBox.scrollHeight;
}

// Fonction active si ouverture d'un chat.
export function OpenChatGroup(rid, conn, data, currId, firstId, Title) {
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

    document.querySelector(".chat-user-username").innerText = Title;

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
        if (msg.value.trim() === "") {
            return
        }
        sendMsg(conn, currId, msg, 'groupmsg');

        let resp = getData('http://localhost:8080/messagegroup?group=' + rid + '&firstId=' + firstId);
        resp.then(value => {


            // Ouverture d'une fenêtre de chat.
            CreateMessages(value, currId);
        }).catch();
    });

    document.querySelector(".emoji-icon").addEventListener("click", function () {
        var chatBox = document.querySelector('.chat');
        chatBox.scrollTop = chatBox.scrollHeight;
    });


    var offset = 10;
    var lastFetchedId = null;

    // Fonction pour charger les messages 10/10.
    /* debouncedScrollHandler = debounce(function () {

        if (chatBox.scrollTop === 0) {
            let resp = getData('http://localhost:8080/message?receiver=' + rid + '&firstId=' + firstId + '&offset=' + offset);
            resp.then(value => {
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
    }, 300); */

    /*     chatBox.addEventListener("scroll", debouncedScrollHandler);
    
        function debounce(func, delay) {
            let timer;
            return function () {
                clearTimeout(timer);
                timer = setTimeout(func, delay);
            };
        } */

    if (data == null) {
        return;
    }
    CreateMessages(data, currId);
}


function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

export const fetchUsersFromAPI = async (id) => {
    try {
        const response = await fetch(`http://localhost:8080/user?id=${id}`, {
            credentials: 'include'
        });
        if (response.ok) {
            const data = await response.json();
            return data;
        } else {
            throw new Error('Failed to fetch users: ' + response.statusText);
        }
    } catch (error) {
        throw new Error('Error fetching users: ' + error.message);
    }
};


export const getGroupForChat = async (id) => {
    try {
        const response = await fetch('http://localhost:8080/getallgroups', {
            method: 'POST',
            body: JSON.stringify(id),
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: "include",
        });

        if (!response.ok) {
            console.error('Get group failed:', response.statusText);
            return { success: false, message: response.statusText };
        }

        const data = await response.json(); // Convertit la réponse JSON en un objet JavaScript
        const groups = data.groups;
        const groupsWhereIamNotIn = data.groupsWhereIamNotIn;

        return { success: true, groups, groupsWhereIamNotIn }; // Retourne l'objet JavaScript
    } catch (error) {
        const errorMessage = error.message ? error.message : 'An error occurred';

        console.error('Error:', errorMessage);

        return { success: false, message: errorMessage };
    }
}

