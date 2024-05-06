import React, { useState } from 'react';

// ----------------------------------------------------------
// ------------------- Création des taches ------------------
// ----------------------------------------------------------

export function Form({ handleAddPost }) {
  const [content, setContent] = useState('');

  function handleSubmit() {
    handleAddPost(content);
    setContent('');
  }

  return (
    <form>
      <textarea
        className='postContent'
        value={content}
        onChange={(e) => setContent(e.target.value)}
        placeholder='Description de la tâche'
        required
      />
      <button onClick={() => { handleSubmit() }}>
        Créer une tâche
      </button>
    </form>
  );
}

// ----------------------------------------------------------
// ------------------ affichage des taches ------------------
// ----------------------------------------------------------

export function Entete() {
  return (
    <div className='divlistTache'>
      <div className='buttonetatlistTache'></div>
      <div className='etatlistTache'>Status</div>
      <div className='contentlistTache'>Contenu</div>
    </div>
  )
}

export function Todolist({ posts, handleValidTache, handleDeletePost }) {
  return (posts.map((post, idx) => (
    <div className={`post ${post.completed ? 'Terminé' : 'À faire'}`} key={idx}>
      <div className='checkEtatPost'>
        <button onClick={() => handleValidTache(post)}>
          Changer le status
        </button>
      </div>
      <div className='etatPost'>{post.completed ? 'Terminé' : 'À faire'}</div>
      <div className='contentpost'>&nbsp;{post.content}</div>
      <button onClick={() => { handleDeletePost(post) }}>Suprimer</button>
    </div>
  )))
}

// ------------------------------------------------------------------
// ------------------- information à l'utilisateur ------------------
// ------------------------------------------------------------------

export function UserConnexion({ account, onConnect }) {
  if (!account) {
    return (
      <div className='userInfo'>
        <button onClick={onConnect}>
          Connecter le wallet pour intéragir
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

// Message de traitement en cours
export function UserMessage({ isVisibleWaiting }) {
  return (
    isVisibleWaiting ? (
      <div className='msg'>
        <div className="Waiting-logo"></div>
        &nbsp;Transaction en cours
      </div>
    ) : null
  );
}

// Message de confirmation de transaction
export function HashMessage({ isVisibleConfirm, hash, removeVisibilityHashMsg }) {
  return (
    isVisibleConfirm ? (
      <div className='msg'>
        <div className='contenuMsg'>
          <p>Tâche réalisée avec succès !!!</p>
          {/* <br /> */}
          <a href={'https://mumbai.polygonscan.com/tx/' + hash} target='blank'>
            <button>
              Detail de la transaction
            </button>
          </a>
        </div>
        <br />
        <button onClick={() => removeVisibilityHashMsg()}>
          Fermer la confirmation
        </button>
      </div>
    ) : null
  );
}