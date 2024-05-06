import React, { useEffect, useState } from 'react';
import { loadSmartContract, loadAccountUser, fetchDataBlockchain, addPost, validTache, deletePost } from './smartContract';
import { Form, Entete, Todolist, UserMessage, HashMessage, UserConnexion } from './web';
import logo from './logo.svg';
import './App.css';

// ---------------------------------------------------------
// ----------- Interaction avec le smart contrat -----------
// ---------------------------------------------------------
// possibilité de deployer gratuitement sur vercel.app

// instance du smart contract utilisé
const ContractInstance = await loadSmartContract();
// // Information du compte (générateur des transactions)
// const currentAccount = await loadAccountUser();

// ---------------------------------------------------------
// ------------------ application todolist -----------------
// ---------------------------------------------------------

function App() {
  const [currentAccount, setConnectedAccount] = useState(null);
  const handleConnect = async () => {
    let userAccount = await loadAccountUser();
    setConnectedAccount(userAccount);
    fetchDataBlockchain(userAccount, ContractInstance, setPosts, changeVisibilityWaitingMsg);
  };

  // Gestion des etats
  const [posts, setPosts] = useState([]); // liste des posts

  const [isVisibleWaiting, setWaiting] = useState(false); // message d'attente 'Transaction en cours'
  const changeVisibilityWaitingMsg = (newVisibility) => {
    setWaiting(newVisibility);
  };

  const [isVisibleConfirm, setConfirm] = useState(false); // Gére l'affichage du message de confirmation
  const [hashValue, setHashValue] = useState(''); // Gére le hash du message de confirmation
  const displayVisibilityHashMsg = (transactionHash) => {
    setWaiting(false);
    setHashValue(transactionHash);
    fetchDataBlockchain(currentAccount, ContractInstance, setPosts, changeVisibilityWaitingMsg);
    setConfirm(true);
  };
  const removeVisibilityHashMsg = () => {
    setConfirm(false);
  };

  // Mise à jour de la liste des posts au demarrage
  useEffect(() => {
    fetchDataBlockchain(currentAccount, ContractInstance, setPosts, changeVisibilityWaitingMsg);
  }, []);

  // ajout d'une nouvelle tache
  const handleAddPost = async (content) => {
    changeVisibilityWaitingMsg(true);
    const transactionHash = await addPost(content, ContractInstance, currentAccount);
    if (transactionHash != "") {
      displayVisibilityHashMsg(transactionHash);
    } else {
      changeVisibilityWaitingMsg(false);
    }
  };

  // Changement du status de la tache
  const handleValidTache = async (post) => {
    changeVisibilityWaitingMsg(true);
    const transactionHash = await validTache(post, ContractInstance, currentAccount);
    if (transactionHash != "") {
      displayVisibilityHashMsg(transactionHash);
    } else {
      changeVisibilityWaitingMsg(false);
    }
  };


  // Suppression d'un post
  const handleDeletePost = async (postToDelete) => {
    changeVisibilityWaitingMsg(true);
    const transactionHash = await deletePost(postToDelete, ContractInstance, currentAccount);
    if (transactionHash != "") {
      displayVisibilityHashMsg(transactionHash);
    } else {
      changeVisibilityWaitingMsg(false);
    }
  };

  // Fonction pour changer le filtre
  const [filter, setFilter] = useState('all'); // filtre actuel
  const changeFilter = (newFilter) => {
    setFilter(newFilter);
  };
  const filteredPosts = posts.filter((post) => {
    if (filter === 'all') {
      return true;
    } else {
      if (filter === 'completed') {
        return post.completed === true;
      } else { // si 'notCompleted'
        return post.completed === false;
      }
    }
  });

  // Retour html effectué par l'application
  return (
    <div className="App">
      {/* fenêtre modal d'informations */}
      <UserMessage isVisibleWaiting={isVisibleWaiting} />
      <HashMessage
        isVisibleConfirm={isVisibleConfirm}
        hash={hashValue}
        removeVisibilityHashMsg={() => setConfirm(false)}
      />

      {/* Encart d'en tête*/}
      <div className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <h1>To-Do List</h1>
      </div>

      {/* Zone d'affichage des taches */}
      <div className="container">
        <div className='displaytache'>
          <h2>Liste des tâches existante</h2>
          <div id='listTache' className='listTache'>
            {Entete()}
            <Todolist
              posts={filteredPosts}
              handleValidTache={(post) => handleValidTache(post)}
              handleDeletePost={(post) => handleDeletePost(post)}
            />
          </div>
        </div>

        <div className='barreLateral'>
          {/* Zone d'information compte utilisateur */}
          <div className='user'>
            <UserConnexion
              account={currentAccount}
              onConnect={handleConnect}
            />
          </div>
          {/* Zone de création d'une tache */}
          <div className='createtache'>
            <h2>Création d'une tache</h2>
            <Form posts={filteredPosts} />
          </div>
          {/* Zone des filtres */}
          <div className='buttonFilter'>
            <div className='titleFilter'>filtrer par : </div>
            <button className='buttonAll' onClick={() => changeFilter('all')}>Toutes les tâches</button>
            <button className='buttonafaire' onClick={() => changeFilter('notCompleted')}>Tâches à faire</button>
            <button className='buttontermine' onClick={() => changeFilter('completed')}>Tâches terminées</button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
