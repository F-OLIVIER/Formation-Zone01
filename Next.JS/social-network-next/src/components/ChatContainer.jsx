import React from 'react';

const ChatContainer = () => {
    return (
        <div className="chat-wrapper">
          <div className="chat-username-wrapper">
            <div className="chat-user">
              <div className="chat-user-username"></div>
              <span id="typing-indicator" style={{ display: 'none' }}>
                <span id="typing-text"></span>
              </span>
            </div>
            <div className="close-chat link">X</div>
          </div>
          <div className="chat"></div>
          <div className="send-wrapper">
            <input type="text" id="chat-input" placeholder="Type a message" />
            <button id="send-btn">Send</button>
          </div>
        </div>
      );
};

export default ChatContainer;