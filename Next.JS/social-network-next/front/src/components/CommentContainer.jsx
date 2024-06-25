import React, { useEffect, useState } from 'react';
import useComments from '../services/useComments';
import { useRouter } from 'next/router';
import { parse, formatDistanceToNow } from 'date-fns';

const formatDate = (dateStr) => {
  const parsedDate = parse(dateStr, 'MM-dd-yyyy HH:mm:ss', new Date());
  return formatDistanceToNow(parsedDate, { addSuffix: true });
};

const CommentContainer = ({ Post_id, NbComments }) => {
  const router = useRouter();
  const [userDetails, setUserDetails] = useState({});
  const [fetching, setFetching] = useState(false);
  const { comments, fetchComments } = useComments(Post_id)

  const handleFetchComments = () => {
    setFetching(true);
    fetchComments(Post_id)
      .then(() => setFetching(false))
      .catch(() => setFetching(false));
  };

  const fetchUsers = async (userIds) => {
    try {
      const userResponses = await Promise.all(
        userIds.map(id =>
          fetch(`http://localhost:8080/user?id=${id}`, {
            credentials: 'include'
          }).then(response => response.json())
        )
      );
      const usersData = userResponses.reduce((acc, userData) => {
        acc[userData.id] = userData;
        return acc;
      }, {});
      setUserDetails(usersData);
    } catch (error) {
      console.error('Error fetching users:', error);
    }
  };
  useEffect(() => {
    if (comments && comments.length > 0) {
      const userIds = comments.map(comment => comment.user_id);
      const uniqueUserIds = [...new Set(userIds)];
      fetchUsers(uniqueUserIds);
    }
  }, [comments]);

  return (
    <div>
      <div className='post-comments'>
        <span
          className='commentsbtn'
          variant="contained"
          onClick={handleFetchComments}
          disabled={fetching}
        >
        </span>
        <span className='commentscount'>
          {NbComments}
        </span>
      </div>
      {comments && (
        <ul>
          {comments.map(comment => (
            <li key={comment.id}>
              <div className='comment-container'>
                <p className='username' onClick={() => router.push(`/user?id=${comment.user_id}`)}>
                {userDetails[comment.user_id] ? userDetails[comment.user_id].nickname : `User ID: ${comment.user_id}`}
                </p>
                <p className='comment-content'>{comment.content}</p>
              </div>
              <p className='time'>{formatDate(comment.date)}</p>
              {comment.image && (
                <img
                  className='comment-image'
                  src={`data:image/jpeg;base64,${comment.image}`}
                  alt="Selected Image"
                />
              )}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};
export default CommentContainer;