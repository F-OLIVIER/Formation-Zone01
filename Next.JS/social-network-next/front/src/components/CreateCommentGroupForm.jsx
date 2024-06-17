import React from 'react';
import { TextField, Button } from '@mui/material';

const CreateCommentGroupForm = ({ handleCreateGroupComment, Post_id }) => {
  const [content, setContent] = React.useState('');


  const handleSubmit = (e) => {
    e.preventDefault();
    const formData = new FormData();
    formData.append('content', content);
    formData.append('post_id', Post_id)
    console.log('Form Data:', formData);
    handleCreateGroupComment(formData);
    setContent('');
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
      <Button type="submit" variant="contained">
        Create Comment
      </Button>
    </form>
  );
};

export default CreateCommentGroupForm;
