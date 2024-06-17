import React, { useState, useEffect } from 'react';
import CreateCommentForm from '../components/CreateCommentForm';
import CommentContainer from '../components/CommentContainer';
import { CardActions } from '@mui/material';
import { useRouter } from 'next/router';
const PostContainer = ({ posts = [], handleCreateComment, handlePostLike }) => {
  const [userDetails, setUserDetails] = useState({});
  const router = useRouter();
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
    if (Array.isArray(posts) && posts.length > 0) {
      const userIds = posts.map(post => post.user_id);
      const uniqueUserIds = [...new Set(userIds)];
      fetchUsers(uniqueUserIds);
    }
  }, [posts]);
  const getPrivacyLabel = (privacy) => {
    switch (privacy) {
      case 0:
        return 'Private';
      case 1:
        return 'Public';
      case 2:
        return 'Semi-Private';
      default:
        return 'Unknown';
    }
  };

  return (
    <>
      <h2 className='pagetitle'>Posts</h2>

      <div className='postscontainer'>
        {posts && (
          <ul>
            {posts.map(post => (
              <li key={post.id}>
                <div className='post'>
                  <div>
                    <div className='post-header'>
                      <div className='post-username-container'>
                        <span className='pp'></span>
                        <p className='post-username' onClick={() => router.push(`/user?id=${post.user_id}`)}>
                          {userDetails[post.user_id] ? userDetails[post.user_id].nickname : `User ID: ${post.user_id}`}                        </p>
                      </div>
                      <p>
                        Date: {post.date}
                      </p>
                    </div>
                    {post.image && (
                      <img
                        src={`data:image/jpeg;base64,${post.image}`}
                        alt="Selected Image"
                        className='post-image'
                      />
                    )}
                    <p className="post-title">
                      {post.title}
                    </p>
                    <p className="post-content">
                      {post.content}
                    </p>
                    <CardActions className='post-likes'>
                      {post.likes && (
                        <>
                          <div
                            className="like-button liked"
                            variant="contained"
                            onClick={() => handlePostLike(post.id)}
                          ></div>
                        </>
                      )}
                      {!post.likes && (
                        <>
                          <div
                            className="like-button"
                            variant="contained"
                            onClick={() => handlePostLike(post.id)}
                          ></div>
                        </>
                      )}
                      <p className="post-likes-count">{post.likes}</p>
                    </CardActions>


                    <div>
                      <CommentContainer Post_id={post.id} NbComments={post.nbcomments} />
                      <CreateCommentForm handleCreateComment={handleCreateComment} Post_id={post.id} />
                    </div>
                  </div>
                </div>
              </li>
            ))}
          </ul>
        )}
      </div>
    </>
  );
};

export default PostContainer;