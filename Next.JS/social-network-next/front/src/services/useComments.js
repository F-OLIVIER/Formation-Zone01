import { useState, useEffect } from 'react';
import {conn, sendMsg} from "@/services/useWebsocket";

const useComments = (post_id) => {
  const [comments, setComments] = useState([]);
  const [commentsGroup, setCommentsGroup] = useState([]);

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
        let Target = [];
        Target.push(data.receiver);
        sendMsg(conn, 0, {value: "New Comment"}, 'comment', Target )
      } else {
        console.error('Create a comment failed:', response.statusText);
      }
    } catch (error) {
      console.error('Error during creation of a comment:', error);
    }
  };

    const fetchCommentsGroup = async () => {
      try {
        const response = await fetch('http://localhost:8080/commentgroup?param=post_id&data='+post_id);
        if (response.ok) {
          const data = await response.json();
          console.log("fetching comments")
          setCommentsGroup(data);
        } else {
          console.error('Failed to fetch comments group:', response.statusText);
        }
      } catch (error) {
        console.error('Error fetching comments group:', error);
      }
    };
  
  
    const createCommentGroup = async (formData) => {
      console.log(formData)
      try {
        const response = await fetch('http://localhost:8080/commentgroup', {
          method: 'POST',
          body: formData,
          credentials: 'include'
        });
  
        if (response.ok) {
          const data = await response.json();
          console.log(data);
  
        } else {
          console.error('Create a group comment failed:', response.statusText);
        }
      } catch (error) {
        console.error('Error during creation of a group comment:', error);
      }
    };

  return { comments, commentsGroup, fetchComments, createComment, fetchCommentsGroup, createCommentGroup };
};

export default useComments;