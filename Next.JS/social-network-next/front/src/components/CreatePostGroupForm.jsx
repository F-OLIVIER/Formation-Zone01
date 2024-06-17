import React, { useState } from 'react';
import { TextField, Button } from '@mui/material';

const CreatePostGroupForm = ({ handleCreatePost }) => {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [image, setImage] = useState(null);

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    setImage(file);
  };

  const handleRemoveImage = () => {
    setImage(null);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    const formData = new FormData();
    formData.append('title', title);
    formData.append('content', content);
    if (image) {
      formData.append('image', image);
    }
    handleCreatePost(formData);
    setTitle('');
    setContent('');
    setImage(null);
  };


  return (
    <div>
      <form onSubmit={handleSubmit} encType="multipart/form-data">
        <h2>Create a New Post</h2>
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
        <h2>Add Image:</h2>
        <input type="file" onChange={handleImageChange} />
        {image && (
          <div>
            <img src={URL.createObjectURL(image)} alt="Uploaded" />
            <Button variant="contained" onClick={handleRemoveImage}>
              Remove Image
            </Button>
          </div>
        )}
        <Button type="submit" variant="contained">
          Create Post
        </Button>
      </form>
    </div>
  );
};

export default CreatePostGroupForm;
