let log, chatBox;
export let currentScrollPos;
export let lastFetchedId;

// Check if document is defined (i.e., if the code is running in the browser)
if (typeof document !== 'undefined') {
    // Access DOM elements only if document is defined
    log = document.querySelector('.chat');
    offlineUsers = document.querySelector('.offline-users');
}

export { log };

// Fonction pour "append" l'historique des messages d'un chat sélectionné.
export function appendLog(container, msg, date, prepend = false) {
    var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
    if (prepend) {
      log.prepend(container); 
    } else {
      log.appendChild(container); 
    }
    container.appendChild(msg);
    msg.appendChild(date);
    
    // Pour se placer en bas de la conversation
    if (doScroll) {
      log.scrollTop = log.scrollHeight - log.clientHeight; 
    }
  }
  

// Création de l'historique des messages d'un chat sélectionné.
export function CreateMessages(data, currId) {
    // On scrute en descendant pour écrire le plus ancien en premier.
    data.reverse();
    for (let i = data.length - 1; i >= 0; i--) {
      const { id, sender_id, content, date } = data[i];
  
      // Check doublon
      if (document.getElementById(`message-${id}`)) {
        continue;
      }
  
      // Création du message avec la mise en forme correspondate si "receiver" ou "sender".
      const receiverContainer = document.createElement("div");
      receiverContainer.className = sender_id == currId ? "sender-container" : "receiver-container";
      
      const receiver = document.createElement("div");
      receiver.className = sender_id == currId ? "sender" : "receiver";
      receiver.innerText = content;
  
      const messagedate = document.createElement("div");
      messagedate.className = sender_id == currId ? "chat-time-left" : "chat-time";
      messagedate.innerText = date.slice(0, -3);
  
      receiverContainer.id = `message-${id}`;
  
      // On "append" avec la fonction ci-dessus dans notre container "chat".
      appendLog(receiverContainer, receiver, messagedate, true);
      chatBox.scrollTop = chatBox.scrollHeight;
    }
  }


  if (typeof document !== 'undefined') {
    // Access DOM elements only if document is defined
    chatBox = document.querySelector('.chat');
}

export { chatBox };

// Fonction active si ouverture d'un chat.
export function OpenChat(rid, conn, data, currId, firstId) {
    document.getElementById('id' + rid).style.fontWeight = "400";

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
    })[0].username;

    document.querySelector(".chat-wrapper").style.display = "flex";
    var msg = document.getElementById("chat-input");
    log.innerHTML = "";
    var typingTimer;
    var typingInterval = 3000;
    var isTyping = false;

    // Typing notification.
    document.querySelector("#chat-input").addEventListener("keydown", function(event) {

        if (event.keyCode !== 13 && !isTyping) { 
            console.log("User started typing");
            isTyping = true;
    
            clearTimeout(typingTimer);
            typingTimer = setTimeout(function() {
                if (isTyping) {
                    console.log("User stopped typing");
                    sendTypingStatus(conn, rid, false);
                    isTyping = false;
                }
            }, typingInterval);
    
            sendTypingStatus(conn, rid, true);
        }
    });
    
    // Envoi du message si click.
    document.querySelector("#send-btn").addEventListener("click", function() {
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
            CreateMessages(value, currId);
            chatBox.scrollTop = chatBox.scrollHeight;
        }).catch();
        log.innerHTML = "";
        clearTimeout(typingTimer); 
        isTyping = false;
    });
    
    // Même chose pour "entrée".
    document.querySelector("#chat-input").addEventListener("keydown", function(event) {
        if (event.keyCode === 13) { 
            offset = 0;
            sendMsg(conn, rid, msg, 'msg');
            firstId = firstId + 10;
            let resp = getData('http://localhost:8080/message?receiver=' + rid + '&firstId=512'  + '&offset=10');
            resp.then(value => {
                if (value && value.length > 0) {
                    const lastIndex = value.length - 1;
                    firstId = value[lastIndex].id;
                    lastFetchedId = firstId;
                  }
                CreateMessages(value, currId);
                chatBox.scrollTop = chatBox.scrollHeight;
            }).catch();
            log.innerHTML = "";
            clearTimeout(typingTimer);
            isTyping = false; 
            offset += value.length;
        }
    });

      var offset = 10; 
      var lastFetchedId = null;
     
      // Fonction pour charger les messages 10/10.
      debouncedScrollHandler = debounce(function() {

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
            return function() {
            clearTimeout(timer);
            timer = setTimeout(func, delay);
                };
      }
            
    if (data == null) {
        return;
    }
    CreateMessages(data, currId);
}
// Fermeture du chat.
document.querySelector(".close-chat").addEventListener("click", function() {
    document.querySelector(".chat-wrapper").style.display = "none";
    resetScroll();
});

// RAZ
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


