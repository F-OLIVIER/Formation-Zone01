export function Searchbar(allUsersMessages, enter) {

    const resultSearch = document.getElementById('resultsearch');
    const divresultsearch = document.getElementById('divresultsearch');
    const closeButton = document.getElementById('closeButton');
    const inputText = document.getElementById('inputsearch');

    const clickHandler = () => {
        resultSearch.classList.remove('active');
        divresultsearch.innerHTML = '';
        closeButton.removeEventListener('click', clickHandler);
    };

    closeButton.addEventListener('click', clickHandler);

    let result = []

    for (const userId in allUsersMessages) {
        if (allUsersMessages.hasOwnProperty(userId)) {
            console.log('allUsersMessages[userId] : ', allUsersMessages[userId])
            const user = allUsersMessages[userId].User;
            // console.log(`User ID: ${user.id}, Name: ${user.firstname} ${user.lastname}, Nickname: ${user.nickname}`);
            let nameuser = '';
            if (user.nickname !== '') {
                nameuser = user.nickname;
            } else {
                nameuser = user.firstname + ' ' + user.lastname;
            }

            // Parcourir les messages de chaque utilisateur
            for (const userId in allUsersMessages) {
                if (Object.hasOwnProperty.call(allUsersMessages, userId)) {
                    const messages = allUsersMessages[userId].Messages;
                    if (messages !== null) {
                        for (let i = 0; i < messages.length; i++) {
                            const currentmessage = messages[i];
                            if (currentmessage.content.toLowerCase().includes(inputText.value.toLowerCase())) {
                                if (currentmessage.sender_id === user.id) {
                                    result.push('<b>On ' + currentmessage.date + ', ' + nameuser + ' wrote : </b>' + currentmessage.content);
                                } else {
                                    result.push('<b>On ' + currentmessage.date + ', you wrote to ' + nameuser + ' : </b>' + currentmessage.content);
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    if (result.length === 0) {
        divresultsearch.innerHTML = 'No result';
    } else {
        divresultsearch.innerHTML = '<span>' + result.join('</span><span>') + '</span>';
    }
    resultSearch.classList.add('active');

    // si enter saisie, vider la recherche
    if (enter) {
        inputText.value = '';
    }
}