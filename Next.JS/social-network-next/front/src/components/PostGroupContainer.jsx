import React from 'react';
import { useEffect, useState } from 'react';
import CommentGroupContainer from './CommentGroupContainer';
import CreateCommentGroupForm from './CreateCommentGroupForm';
import { useRouter } from 'next/router';
import { parse, formatDistanceToNow } from 'date-fns';


const formatDate = (dateStr) => {
  const parsedDate = parse(dateStr, 'MM-dd-yyyy HH:mm:ss', new Date());
  return formatDistanceToNow(parsedDate);
};



const PostGroupContainer = ({ posts, handleCreateGroupComment }) => {
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

  return (
    <div>
      <h2 className='pagetitle'>Posts</h2>
      <div className='postscontainer commentposts-container'>

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
                        {userDetails[post.user_id] ? userDetails[post.user_id].nickname : `User ID: ${post.user_id}`}
                        </p>
                      </div>
                      <p>
                        {formatDate(post.date)}
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
                    <div className='commentgroup-section'></div>
                    <div>
                      <CommentGroupContainer Post_id={post.id} NbComments={post.nbcomments} />
                      <CreateCommentGroupForm handleCreateGroupComment={handleCreateGroupComment} Post_id={post.id} />
                    </div>
                  </div>
                </div>
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
};

export default PostGroupContainer;