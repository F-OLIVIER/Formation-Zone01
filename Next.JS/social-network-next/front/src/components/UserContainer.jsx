import React from 'react';

const UserContainer = () => {
  return (
    <div className='test'>
      <div className="users">
        <p className="title-users"></p>
        <p className="online-list"></p>
        <div className="online-users"></div>
        <p className="online-list pagetitle" id="offline-list">Contacts</p>
        <div className="offline-users"></div>
      </div>
    </div>
  );
}

export default UserContainer;