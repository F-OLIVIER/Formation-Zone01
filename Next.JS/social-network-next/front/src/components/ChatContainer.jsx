import React, { useState } from 'react';
import Picker from "emoji-picker-react";
import { useRouter } from 'next/router';

const ChatContainer = () => {
  const [inputStr, setInputStr] = useState("");
  const [showPicker, setShowPicker] = useState(false);

  const handleInputChange = (e) => {
    setInputStr(e.target.value);
  };

  const onEmojiClick = (emojiObject, event) => {
    setInputStr((prevInput) => prevInput + emojiObject.emoji);
    setShowPicker(false);
  };

  const sendMessage = () => {
    console.log("Sending message:", inputStr);
    setInputStr("");
  };
  
  const router = useRouter();

  return (
    <div className="chat-wrapper">
      <div className="chat-username-wrapper">
        <div className="chat-user">
          <a className="chat-user-username" onClick={() => router.push('/user?id=' + document.querySelector('.chat-user-username').id)}></a>
          <span id="typing-indicator" style={{ display: 'none' }}>
            <span id="typing-text"></span>
          </span>
        </div>
        <div className="close-chat link">X</div>
      </div>
      <div className="chat"></div>
      <div className="send-wrapper">
      </div>
      <div className="app">
        <div className="picker-container">
        {showPicker && (
          <Picker className='emoji-picker' pickerStyle={{ width: "100%", backgroundColor: "transparent" }} onEmojiClick={onEmojiClick} />
                    )}
            </div>
          <input
            className="chat-input"
            id="chat-input"
            value={inputStr}
            onChange={handleInputChange}
            aria-label="Type a message"
          />
          <img
            className="emoji-icon"
            src="https://icons.getbootstrap.com/assets/icons/emoji-smile.svg"
            alt="Emoji Picker"
            onClick={() => setShowPicker((val) => !val)}
          />
          <button id="send-btn" aria-label="Send" onClick={sendMessage}></button>
      </div>
    </div>
  );
};

export default ChatContainer;