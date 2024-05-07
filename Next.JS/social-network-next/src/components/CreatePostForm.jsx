import React, { useState, useEffect } from 'react';
import { TextField, Button, Select, MenuItem, Checkbox, FormControlLabel } from '@mui/material';

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

  const handleSubmit = (e) => {
    e.preventDefault();
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
        <Select
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

export default CreatePostForm;
