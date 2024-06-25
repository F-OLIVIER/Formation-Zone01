import React, { useState, useEffect } from 'react';
import { Button, Select, MenuItem, Checkbox, FormControlLabel } from '@mui/material';

const CreatePostForm = ({ handleCreatePost, fetchUsers, id, users }) => {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [privacy, setPrivacy] = useState(0);
  const [image, setImage] = useState(null);
  const [selectedFollowers, setSelectedFollowers] = useState([]);
  const [loading, setLoading] = useState(true);

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    setImage(file);
  };

  const handleRemoveImage = () => {
    setImage(null);
  };

  function openCreatePost() {
    const createpostFormContainer = document.querySelector('.createpost-form-container')
    console.log("createpostFormContainer:", createpostFormContainer)
    createpostFormContainer.classList.toggle('open')
  }

  const handleSubmit = (e) => {
    e.preventDefault();
    if (title.trim() === "" || content.trim() === "") {
      return
    }
    const formData = new FormData();
    formData.append('title', title);
    formData.append('content', content);
    formData.append('privacy', privacy);
    if (image) {
      formData.append('image', image);
    }
    if (selectedFollowers) {
      formData.append('selectedFollowers', selectedFollowers);
    }
    handleCreatePost(formData);
    setTitle('');
    setContent('');
    setPrivacy(0);
    setImage(null);
  };

  useEffect(() => {
    if (privacy === 2) {
      fetchUsers(id)
        .then(() => setLoading(false))
        .catch(error => console.error('Error fetching users:', error));
    } else {
      setSelectedFollowers([]);
      setLoading(false);
    }
  }, [privacy]);

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
          <h2>Create a new post</h2>
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
            <Select
              className='select-menu'
              label="Privacy"
              value={privacy}
              onChange={(e) => setPrivacy(parseInt(e.target.value))}
            >
              <MenuItem value={0}>Private</MenuItem>
              <MenuItem value={1}>Public</MenuItem>
              <MenuItem value={2}>Semi-Private</MenuItem>
            </Select>
            {privacy === 2 && users && users.listfollowers && (
              <div>
                {users.listfollowers.map(follower => (
                  <div key={follower.id}>
                    <FormControlLabel
                      control={
                        <Checkbox
                          checked={selectedFollowers && selectedFollowers.includes(follower.id)}
                          onChange={(e) => {
                            const isChecked = e.target.checked;
                            setSelectedFollowers(prevSelectedFollowers => {
                              // Ensure prevSelectedFollowers is initialized as an array
                              if (!Array.isArray(prevSelectedFollowers)) {
                                prevSelectedFollowers = [];
                              }
                              if (isChecked) {
                                return [...prevSelectedFollowers, follower.id];
                              } else {
                                return prevSelectedFollowers.filter(id => id !== follower.id);
                              }
                            });
                          }}
                        />
                      }
                      label={`${follower.firstname} ${follower.lastname}`}
                    />
                  </div>
                ))}
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

export default CreatePostForm;
