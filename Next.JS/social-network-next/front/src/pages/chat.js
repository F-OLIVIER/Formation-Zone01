import { useState, useEffect } from 'react';
import ChatContainer from '../components/ChatContainer';
import UserContainer from '../components/UserContainer';
import { getUsers, updateUsers } from '../services/useWebsocket';

const Chat = ({ loggedIn, id }) => {
    useEffect(() => {
        if (loggedIn) {
        //console.log(id)
        getUsers().then(function () {
          updateUsers(id);
      });
    }
      }, [loggedIn]);

      return (

    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
        <ChatContainer />
        <UserContainer />
    </div>
    
      )

}

export default Chat