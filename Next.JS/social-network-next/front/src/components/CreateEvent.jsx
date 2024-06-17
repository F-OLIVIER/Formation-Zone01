import React, { useState } from 'react';
import { TextField, Button } from '@mui/material';

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
    handleCreateEvent(formData);
    setTitle('');
    setContent('');
    setDate('');
  };


  return (
    <div>
      <form onSubmit={handleSubmit} encType="multipart/form-data">
        <h2>Create a New Event</h2>
        <TextField
          label="Title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <TextField
          label="Content"
          value={content}
          onChange={(e) => setContent(e.target.value)}
          multiline
        />
        <TextField
          label="Date"
          value={date}
          onChange={(e) => setDate(e.target.value)}
          multiline
        />

        <Button type="submit" variant="contained">
          Create Event
        </Button>
      </form>
    </div>
  );
};

export default CreateEvent;
