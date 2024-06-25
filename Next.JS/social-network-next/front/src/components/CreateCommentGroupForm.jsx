import React from 'react';
import { Button } from '@mui/material';

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
      <div className='commentInputContainer'>
      <input
          placeholder='Write a comment...'
          className='inputBox commentinput'
          value={content}
          onChange={(e) => setContent(e.target.value)}
        />
        <Button className='commentbtn' type="submit" variant="contained">
        </Button>
        </div>
    </form>
  );
};

export default CreateCommentGroupForm;
