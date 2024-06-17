import React, { useState } from 'react';
import useComments from '../services/useComments';
import { useRouter } from 'next/router';
import { parse, formatDistanceToNow } from 'date-fns';
const formatDate = (dateStr) => {
  const parsedDate = parse(dateStr, 'MM-dd-yyyy HH:mm:ss', new Date());
  return formatDistanceToNow(parsedDate, { addSuffix: true });
};
const CommentContainer = ({ Post_id, NbComments }) => {
  const router = useRouter();
  const [fetching, setFetching] = useState(false);
  const { comments, fetchComments } = useComments(Post_id)
  const handleFetchComments = () => {
    setFetching(true);
    fetchComments(Post_id)
      .then(() => setFetching(false))
      .catch(() => setFetching(false));
  };

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
                  {comment.user_id}PASTEQUITOS:
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