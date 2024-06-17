export var TargetUsers = [];

export const Target = async (id) => {
    try {
        //console.log("id", id)
        const response = await fetch('http://localhost:8080/target', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({user_id: id}),
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const text = await response.text(); // Lire la réponse comme du texte

        const data = JSON.parse(text); // Analyser le texte en JSON
        //console.log("notif", data);

        if (data.success) {
            if (data.listTarget) {
                data.listTarget.forEach((user) => {
                    console.log("target", user.id);
                    TargetUsers.push(user.id);
                });
            }
        } else {
            console.error('Failed to fetch target');
        }
    } catch (error) {
        console.error('Error TARGET:', error.message);
    }
};