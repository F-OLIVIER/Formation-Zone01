import { useState, useEffect } from 'react';
import ChatContainer from '../components/ChatContainer';
import GroupUserContainer from '../components/GroupUserContainer';
import { getUsers, updateUsers } from '../services/useWebsocket';

const ChatGroup = ({ loggedIn, id }) => {

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
        <GroupUserContainer />
    </div>
    
      )

}

export default ChatGroup