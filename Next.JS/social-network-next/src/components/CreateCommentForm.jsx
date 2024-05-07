import React from 'react';
import { TextField, Button } from '@mui/material';

const CreateCommentForm = ({ handleCreateComment, Post_id }) => {
  const [content, setContent] = React.useState('');
  const [image, setImage] = React.useState(null);

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
    formData.append('content', content);
    formData.append('post_id', Post_id)
    if (image) {
      formData.append('image', image);
    }
    console.log('Form Data:', formData);
    handleCreateComment(formData);
    setContent('');
    setImage(null);
  };

  return (
    <form onSubmit={handleSubmit} encType="multipart/form-data">
      <h2>Create a New Comment</h2>
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
        Create Comment
      </Button>
    </form>
  );
};

export default CreateCommentForm;
