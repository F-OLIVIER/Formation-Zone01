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
    if (content.trim() === "") {
      return
    }
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
      <div className='commentInputContainer'>
        <input
          placeholder='Write a comment...'
          className='inputBox commentinput'
          value={content}
          onChange={(e) => setContent(e.target.value)}
        />
        {image && (
          <div className='comment-image-container'>
            <img className='comment-image-input' src={URL.createObjectURL(image)} alt="Uploaded" />
            <div className='remove-image' variant="contained" onClick={handleRemoveImage}></div>
          </div>
        )}
        <input id='file' className={'inputBox file-input'} type="file" onChange={handleImageChange} />
        <label className="input-file-label" for="file"></label>
        <Button className='commentbtn' type="submit" variant="contained">
        </Button>
      </div>
    </form>
  );
};

export default CreateCommentForm;
