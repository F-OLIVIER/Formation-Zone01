export async function cookie(data) {
  try {
      const response = await fetch('http://localhost:8080/logout', {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json'
          },
          body: JSON.stringify(data),
          credentials: 'include',
      });
      if (!response.ok) {
          throw new Error('Network response was not ok');
      }
      return await response.json();
  } catch (error) {
      throw new Error('Error during delete cookie: ' + error.message);
  }
}

export async function session() {
  try {
      const response = await fetch('http://localhost:8080/session', {
          method: 'GET',
          headers: {
              'Content-Type': 'application/json'
          },
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
