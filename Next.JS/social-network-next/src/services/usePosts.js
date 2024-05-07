import { useState, useEffect } from 'react';

const usePosts = () => {
  const [posts, setPosts] = useState([]);

  const fetchPosts = async (id) => {
    try {
      console.log("fetch for id:", id)
      const response = await fetch(`http://localhost:8080/post?id=${id}`);
      if (response.ok) {
        const data = await response.json();
        setPosts(data);
      } else {
        console.error('Failed to fetch posts:', response.statusText);
      }
    } catch (error) {
      console.error('Error fetching posts:', error);
    }
  };

  const createPost = async (formData) => {
    console.log(formData)
    try {
      const response = await fetch('http://localhost:8080/post', {
        method: 'POST',
        body: formData,
        credentials: 'include'
      });

      if (response.ok) {
        const data = await response.json();
        console.log(data);
      } else {
        console.error('Create a post failed:', response.statusText);
      }
    } catch (error) {
      console.error('Error during creation of a post:', error);
    }
  };

  return { posts, fetchPosts, createPost };
};

export default usePosts;