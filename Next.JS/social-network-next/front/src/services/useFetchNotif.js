export const fetchNotification = async (id, setNotifications) => {
    try {
        //console.log("id", id)
        const response = await fetch('http://localhost:8080/notif', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({user_id_receiver: id}),
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const text = await response.text(); // Lire la réponse comme du texte

        const data = JSON.parse(text); // Analyser le texte en JSON
        //console.log("notif", data);

        if (data.success) {
            const allNotifications = [];
            let idCounter = 0;
            if (data.listGroupMp) {
                data.listGroupMp.forEach(element => {
                    const notification = {
                        id: idCounter++, // Assign a unique ID
                        category: 'GroupMP',
                        user: element[0],
                        message: element[1],
                        group: element[2]
                    };
                    //console.log("Notification Group: ", notification);
                    allNotifications.push(notification);
                });
            }

            if (data.listEvent) {
                data.listEvent.forEach(element => {
                    const notification = {
                        id: idCounter++, // Assign a unique ID
                        category: 'Event',
                        user: element[0],
                        group: element[1],
                        event: element[2]
                    };
                    //console.log("Notification Group: ", notification);
                    allNotifications.push(notification);
                });
            }

            if (data.listAskGroup) {
                data.listAskGroup.forEach(element => {
                    const notification = {
                        id: idCounter++, // Assign a unique ID
                        category: 'AskGroup',
                        user: element[0],
                        group: element[1]
                    };
                    //console.log("Notification Group: ", notification);
                    allNotifications.push(notification);
                });
            }

            if (data.listPost) {
                data.listPost.forEach(element => {
                    const notification = {
                        id: idCounter++, // Assign a unique ID
                        category: 'Post',
                        user: element[0],
                        post: element[1]
                    };
                    //console.log("Notification Follow: ", notification);
                    allNotifications.push(notification);
                });
            }

            if (data.listFollowers) {
                data.listFollowers.forEach(element => {
                    const notification = {
                        id: idCounter++, // Assign a unique ID
                        category: 'Follow',
                        user: element[0],
                        follow: element[1]
                    };
                    //console.log("Notification Follow: ", notification);
                    allNotifications.push(notification);
                });
            }

            if (data.listMP) {
                data.listMP.forEach(element => {
                    const notification = {
                        id: idCounter++,
                        category: 'MP',
                        user: element[0],
                        message: element[1]
                    };
                    //console.log("Notification MP: ", notification);
                    allNotifications.push(notification);
                });
            }

            if (data.listComment) {
                data.listComment.forEach(element => {
                    if (Array.isArray(element)) {
                        const notification = {
                            id: idCounter++,
                            category: 'Comment',
                            user: element[0],
                            comment: element[1],
                            post: element[2]
                        };
                        //console.log("Notification: ", notification);
                        allNotifications.push(notification);
                    }
                });
            }
            if (data.listGroup) {
                data.listGroup.forEach(element => {
                    if (Array.isArray(element)) {
                        const notification = {
                            id: idCounter++,
                            category: 'Group',
                            user: element[0],
                            group: element[1]
                        };
                        //console.log("Notification: ", notification);
                        allNotifications.push(notification);
                    }
                });
            }

            setNotifications(allNotifications);
            //const PastilleNotif = allNotifications.length
        } else {
            console.error('Failed to fetch notifications');
        }
    } catch (error) {
        console.error('Error NOTIF:', error.message);
    }
};
