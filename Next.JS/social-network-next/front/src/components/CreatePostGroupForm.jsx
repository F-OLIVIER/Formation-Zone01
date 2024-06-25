import React, { useState } from 'react';
import { Button } from '@mui/material';

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

  function openCreatePost() {
    const createpostFormContainer = document.querySelector('.createpost-form-container')
    console.log("createpostFormContainer:", createpostFormContainer)
    createpostFormContainer.classList.toggle('open')
  }


  return (
    <>

    <div className='createpost'>
    <div className='createpostbtn' onClick={openCreatePost}>
      <span className='addpostplus'></span>
      <svg className='addpost' xmlns="http://www.w3.org/2000/svg" width="30%" height="50%" fill="#FFC700" viewBox="0 0 868 970">
        <path stroke="#FFC700" stroke-width="20" d="m479 22.528 333.013 192.265a90 90 0 0 1 45 77.942v384.53a90 90 0 0 1-45 77.942L479 947.472a90 90 0 0 1-90 0L55.987 755.207a90 90 0 0 1-45-77.942v-384.53a90 90 0 0 1 45-77.942L389 22.528a90 90 0 0 1 90 0Z" />
      </svg>
    </div>
  </div>

    <div className='createpost-form-container'>
        <form onSubmit={handleSubmit} encType="multipart/form-data" className='createpost-form'>
          <h2>Create a New Post</h2>
          <input placeholder='Title' className='createpost-form-input title-input' type="text" value={title} onChange={(e) => setTitle(e.target.value)} />
          <textarea placeholder='Content' className='createpost-form-input' type="text" value={content} onChange={(e) => setContent(e.target.value)} />
          <div className='displayflex'>
            {!image && (
              <div>
                <input id='file' className={'inputBox file-input'} type="file" onChange={handleImageChange} />
                <label className="input-file-label" for="file"></label>
              </div>
            )}
            {image && (
              <div className='flexcolumn'>
                <img className='img-preview' src={URL.createObjectURL(image)} alt="Uploaded" />
                <Button className='removeimgbtn' variant="contained" onClick={handleRemoveImage}>
                  Remove Image
                </Button>
              </div>
            )}
          </div>
          <Button className='submitbtn' type="submit" variant="contained">
            Create Post
          </Button>
        </form>
      </div>
    </>
  );
};

export default CreatePostGroupForm;
