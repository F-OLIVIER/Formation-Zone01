export const createGroup = async (form) => {
    try {
        console.log(form)

        const response = await fetch('http://localhost:8080/creategroup', {
            method: 'POST',
            body: JSON.stringify(form),
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include'
        });

        if (response.ok) {
            const data = await response.json();
            if (data && data.success) {
                console.log(data);
                return { success: true, data };
            } else {
                console.error('Create group failed:', data ? data.message : 'No response body');
                return { success: false, message: data ? data.message : 'No response body' };
            }
        } else {
            console.error('Create group failed:', response.statusText);
            return { success: false, message: response.statusText };
        }
    } catch (error) {
        console.error('Error during creation of group:', error);
        return { success: false, message: error.message };
    }
};

export const getGroup = async () => {
    try {
        const response = await fetch('http://localhost:8080/getallgroups', {
            method: 'GET',
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
        return { success: true, data }; // Retourne l'objet JavaScript
    } catch (error) {
        const errorMessage = error.message ? error.message : 'An error occurred';

        console.error('Error:', errorMessage);

        return {success: false, message: errorMessage};
    }
}

export const InviteInMyGroup = async (formInvite) => {
    if (!formInvite.nameOfGroup || !formInvite.nameOfThePerson) {
        console.error('Error: nameOfGroup or nameOfThePerson is undefined or empty');
        return { success: false, message: 'nameOfGroup or nameOfThePerson is undefined or empty' };
    }
    try {
        const response = await fetch('http://localhost:8080/inviteinmygroup', {
            method: 'POST',
            body: JSON.stringify(formInvite),
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include'
        });

        if (response.ok) {
            const data = await response.json();
            if (data && data.success) {
                console.log(data);
                return { success: true, data };
            } else {
                console.error('Invite group failed:', data ? data.message : 'No response body');
                return { success: false, message: data ? data.message : 'No response body' };
            }
        } else {
            console.error('Invite group failed:', response.statusText);
            return { success: false, message: response.statusText };
        }
    } catch (error) {
        console.error('Error during invitation of group:', error);
        return { success: false, message: error.message };
    }
}
