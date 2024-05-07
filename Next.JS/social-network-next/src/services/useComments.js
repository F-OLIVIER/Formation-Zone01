import { useState, useEffect } from 'react';

const useComments = (post_id) => {
  const [comments, setComments] = useState([]);
  const fetchComments = async () => {
    try {
      const response = await fetch('http://localhost:8080/comment?param=post_id&data='+post_id);
      if (response.ok) {
        const data = await response.json();
        console.log("fetching comments")
        setComments(data);
      } else {
        console.error('Failed to fetch comments:', response.statusText);
      }
    } catch (error) {
      console.error('Error fetching comments:', error);
    }
  };


  const createComment = async (formData) => {
    console.log(formData)
    try {
      const response = await fetch('http://localhost:8080/comment', {
        method: 'POST',
        body: formData,
        credentials: 'include'
      });

      if (response.ok) {
        const data = await response.json();
        console.log(data);

      } else {
        console.error('Create a comment failed:', response.statusText);
      }
    } catch (error) {
      console.error('Error during creation of a comment:', error);
    }
  };

  return { comments, fetchComments, createComment };
};

export default useComments;