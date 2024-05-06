// Récupération de l'account client
export async function loadAccountUser() {
  // console.log("abi :", abi);
  if (window.ethereum) { // si un wallet et present sur le navigateur client
    const account = await window.ethereum.request({ method: 'eth_accounts' }); // initialisation de l'account
    if (account.length === 0) {
      // Si absence de compte, demande d'autorisation d'accées au compte puis rechargement du compte
      await window.ethereum.request({ method: 'eth_requestAccounts' }); // Demande au wallet l'autorisation d'accéder aux comptes
      return await loadAccountUser();
    }
    return account[0]
  } else {
    alert("Veuillez installer un wallet");
    return ""
  }
}

export function UserConnexion({ account, onConnect }) {
  if (!account) {
    return (
      <div className='userInfo'>
        <button onClick={onConnect}>
          Connecter votre wallet
        </button>
      </div>
    );
  } else {
    const startSubstring = account.slice(0, 6);
    const endSubstring = account.slice(-4);
    return (
      <div className='userInfo'>
        <p>
          <strong>
            Connecté avec le wallet n°:
            <br />
            {startSubstring}...{endSubstring}
          </strong>
        </p>
      </div>
    );
  }
}
