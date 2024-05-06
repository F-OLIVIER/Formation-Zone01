import abi from './abiToDolist.json' // 2eme smart Contract avec delete message
import web3 from 'web3';

// https://mumbai.polygonscan.com
//  ___________________________________________________________________________________
// |                             Fonction du smart contract                            |
// |-----------------------------------------------------------------------------------|
// | createTask      --> Création d'une tache                                          |
// | toggleCompleted --> Compléter la tache (change son status)                        |
// | deleteTask      --> Suppression d'une tache                                       |
// | getMyTasks      --> Call la liste des id des taches existante                     |
// | getTask         --> Appel la tache à partir de son id pour connaitre sont status  |
// |___________________________________________________________________________________|

// Création de l'instance du smart contract
const contractAddress = '0xfFA7a3f2895b7316457d101411fe4038443B887a'; // 2eme smart Contract avec delete message
const Web3 = new web3(window.ethereum);
export async function loadSmartContract() {
  const ContractInstance = new Web3.eth.Contract(abi, contractAddress);
  return ContractInstance
}

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
    alert("Veuillez installer un wallet comme MetaMask");
    return ""
  }
}

// ------------------------------------------------------------------------
// ------------------ Interaction avec le Smart Contract ------------------
// ------------------------------------------------------------------------

// Récupération des taches
export async function fetchDataBlockchain(currentAccount, ContractInstance, setPosts, changeVisibilityWaitingMsg) {
  changeVisibilityWaitingMsg(false);
  // Récupération des id des taches existante sur la blockchain
  if (currentAccount != "") {
    setPosts([]); // vider l'array de post
    // Récupération de la liste des id des taches
    let getMyTasks = await ContractInstance.methods.getMyTasks().call({ from: currentAccount });
    // Récupération du contenu de chaque tache à partir des id récupéré
    for (let i = getMyTasks.length - 1; i >= 0; i--) {
      const getTask = await ContractInstance.methods.getTask(getMyTasks[i]).call({ from: currentAccount });
      // console.log("getTask : ", getTask)
      if (getTask.content != "") {
        setPosts((prevPosts) => [...prevPosts, getTask]);
      }
    }
  }
}

export async function addPost(content, ContractInstance, currentAccount) {
  if (ContractInstance && currentAccount != "") {
    // Création de la nouvelle tache
    const transactionCreate = await ContractInstance.methods.createTask(content).send({ from: currentAccount });
    // Récupération du hash de la transaction
    const transactionHash = transactionCreate.events.TaskCreated.transactionHash;
    // Gestion de la confirmation de la transaction
    if (transactionCreate) {
      const receipt = await Web3.eth.getTransactionReceipt(transactionHash);
      if (receipt.status) {
        return transactionHash
      } else {
        alert(`Probléme avec la transaction !`);
        return ""
      }
    }
  }
};

export async function validTache(post, ContractInstance, currentAccount) {
  if (ContractInstance && currentAccount != "") {
    const transactionValidTache = await ContractInstance.methods.toggleCompleted(post.id).send({ from: currentAccount });
    const transactionHash = transactionValidTache.events.TaskCompleted.transactionHash;
    if (transactionValidTache) {
      const receipt = await Web3.eth.getTransactionReceipt(transactionHash);
      if (receipt.status) {
        return transactionHash
      } else {
        alert(`Probléme avec la transaction !`);
        return ""
      }
    }
  }
};

export async function deletePost(postToDelete, ContractInstance, currentAccount) {
  const isConfirmed = window.confirm("Êtes-vous sûr de vouloir supprimer cette tâche ?");
  if (isConfirmed && currentAccount != "" && ContractInstance) {
    const transactionValidTache = await ContractInstance.methods.deleteTask(postToDelete.id).send({ from: currentAccount });
    const transactionHash = transactionValidTache.transactionHash;
    if (transactionValidTache) {
      const receipt = await Web3.eth.getTransactionReceipt(transactionHash);
      if (receipt.status) {
        return transactionHash
      } else {
        alert(`Probléme avec la transaction !`);
        return ""
      }
    }
  }
};

