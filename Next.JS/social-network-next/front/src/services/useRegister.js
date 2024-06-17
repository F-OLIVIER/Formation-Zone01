//login_auth.js

export async function RegisterUser(formData) {
    try {
        const response = await fetch('http://localhost:8080/register', {
            method: 'POST',
            body: formData,
            credentials: 'include'
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        return await response.json();
    } catch (error) {
        throw new Error('Error during registration: ' + error.message);
    }
}
