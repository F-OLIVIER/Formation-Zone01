import { useState } from 'react';

const Profile = ({ account }) => {
    //TODO Get a few blockchain imgs to set for the profile
    //TODO Get balance from account and get token acronyme

    //TODO To Only get MINTED NFTs
    let username = ""

    return (
        <div id="profileDropDown">
            <div id="profileTitle">
                <img src="/pp.svg" alt="" />
                <p id="username">{username}</p>
            </div>
            <div id="currentWalletInfo">
                <img src="" alt="wallet current blockchain"/>
                <div>
                    {/* Wallet Name */}
                    <p>Main</p>
                    {/* Wallet Balance */}
                    <p>0 ETH</p>
                </div>
                <div>
                    {/* Wallet Adresse and button to copy adresse */}
                    <p>{account}</p>
                    {/* Button to add more funds */}
                </div>
            </div>
        </div>
    )
}

const Connexion = () => {
    const [currentAccount, setConnectedAccount] = useState(null);

    const connectWallet = async () => {
        if (window.ethereum) { // si un wallet et present sur le navigateur client
          const account = await window.ethereum.request({ method: 'eth_accounts' }); // initialisation de l'account
            if (account.length === 0) {
            // Si absence de compte, demande d'autorisation d'accées au compte puis rechargement du compte
                await window.ethereum.request({ method: 'eth_requestAccounts' }); // Demande au wallet l'autorisation d'accéder aux comptes
                return await connectWallet();
            }
            return account[0]
        } else {
            alert("Veuillez installer un wallet");
            return ""
        }
    }
    const handleConnect = async () => {
        try {
            let userAccount = await connectWallet();
            setConnectedAccount(userAccount);
        } catch (error) {
            alert("Couldn't connect user, try later.")
            console.error('Error connecting user:', error);
        }
    };

    return (
        <div id="profileSection">
            <input id="profileBtn" type="button" value="Connect Wallet" onClick={handleConnect}></input>
            <Profile 
                account={currentAccount}
            />
        </div>
    )
}

export default Connexion
