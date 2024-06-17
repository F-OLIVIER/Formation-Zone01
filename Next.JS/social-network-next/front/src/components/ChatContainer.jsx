import React, { useState } from 'react';
import Picker from "emoji-picker-react";

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
      </div>
      <div className="app">
        <div className="picker-container">
        {showPicker && (
          <Picker pickerStyle={{ width: "100%", backgroundColor: "transparent" }} onEmojiClick={onEmojiClick} />
                    )}
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
          <button id="send-btn" aria-label="Send" onClick={sendMessage}>Send</button>


        </div>
      </div>
    </div>
  );
};

export default ChatContainer;