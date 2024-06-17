import React from 'react';
import CommentGroupContainer from './CommentGroupContainer';
import CreateCommentGroupForm from './CreateCommentGroupForm';
import { Button, Card, CardContent, CardActions, Typography } from '@mui/material';
import { useRouter } from 'next/router';



const PostGroupContainer = ({ posts, handleCreateGroupComment }) => {
  const router = useRouter();

  return (
    <div>
      <h2>Posts</h2>
      {posts && (
        <ul>
          {posts.map(post => (
            <li key={post.id}>
              <Card>
                <CardContent>
                  <Typography variant="h5" component="h2">
                    {post.title}
                  </Typography>
                  <Typography variant="body2" color="textSecondary" component="p">
                    {post.content}
                  </Typography>
                  <Typography variant="body2" color="textSecondary" component="p">
                    Posted by{' '}
                    <Typography variant="body2" color="primary" component="a" onClick={() => router.push(`/user?id=${post.user_id}`)}>
                      User ID: {post.user_id}
                    </Typography>
                  </Typography>
                  <Typography variant="body2" color="textSecondary" component="p">
                    Date: {post.date}
                  </Typography>
                  {post.image && (
                    <img
                      src={`data:image/jpeg;base64,${post.image}`}
                      alt="Selected Image"
                    />
                  )}


                  <div>
                    Comments :
                    <CreateCommentGroupForm handleCreateGroupComment={handleCreateGroupComment} Post_id={post.id} />
                    <CommentGroupContainer Post_id={post.id} NbComments={post.nbcomments} />
                  </div>
                </CardContent>
              </Card>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default PostGroupContainer;