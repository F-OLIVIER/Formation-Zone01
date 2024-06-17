//cookie.js

export async function session() {
    try {
        const response = await fetch('http://localhost:8080/session', {
            credentials: 'include',
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        return await response.json();
    } catch (error) {
        throw new Error('Error during re-login: ' + error.message);
    }
}
