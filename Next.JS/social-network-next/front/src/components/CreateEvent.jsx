import React, { useState } from 'react';
import { Button } from '@mui/material';

const CreateEvent = ({ handleCreateEvent }) => {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [date, setDate] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    const formData = new FormData();
    formData.append('title', title);
    formData.append('content', content);
    formData.append('date', date)
    if (title.trim() != "" && content.trim() != "" && date.trim() != "") {
      handleCreateEvent(formData);
      setTitle('');
      setContent('');
      setDate('');
    } else {
      alert('Please complete all fields');
    }
  };


  return (
    <div>
      <form onSubmit={handleSubmit} encType="multipart/form-data">
        <h2 className='pagetitle'>Create a New Event</h2>
        <div className='createEvent-form'>

        <div className='inputContainer'>
          <input
            className={'inputBox autosize'}
            placeholder="Title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            />
        </div>
        <div className='inputContainer'>
        <input
          className={'inputBox autosize'}
          placeholder="Content"
          value={content}
          onChange={(e) => setContent(e.target.value)}
          multiline
          />
        </div>
        <div className='inputContainer'>
        <input
          className={'inputBox autosize'}
          placeholder="Date"
          value={date}
          onChange={(e) => setDate(e.target.value)}
          multiline
          />
        </div>

        <Button className='createEventbtn' type="submit" variant="contained">
          +
        </Button>
          </div>
      </form>
    </div>
  );
};

export default CreateEvent;
