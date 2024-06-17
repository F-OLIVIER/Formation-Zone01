import React, { useState } from 'react';
import Link from 'next/link';
import useComments from '../services/useComments';
import { Button, Typography } from '@mui/material';
import { useRouter } from 'next/router';

const CommentGroupContainer = ({ Post_id, NbComments }) => {
  const router = useRouter();
  const [fetching, setFetching] = useState(false);
  const { commentsGroup, fetchCommentsGroup } = useComments(Post_id)
  const handleFetchComments = () => {
    console.log(Post_id)
    setFetching(true);
    fetchCommentsGroup(Post_id)
      .then(() => setFetching(false))
      .catch(() => setFetching(false));
  };

  return (
    <div>
      <h2>Comments</h2>
      <Button
        variant="contained"
        onClick={handleFetchComments}
        disabled={fetching}
      >
        {NbComments}
      </Button>
      {commentsGroup && (
        <ul>
          {commentsGroup.map(comment => (
            <li key={comment.id}>
              <p>{comment.content}</p>
              <Typography variant="body2" color="textSecondary" component="p">
                Commented by{' '}
                <Typography variant="body2" color="primary" component="a" onClick={() => router.push(`/user?id=${comment.user_id}`)}>
                  User ID: {comment.user_id}
                </Typography>
              </Typography>
              <p>Date: {comment.date}</p>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default CommentGroupContainer;