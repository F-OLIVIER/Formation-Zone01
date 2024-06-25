import { useEffect } from 'react';
import { useRouter } from 'next/router';
import CreatePostForm from '../components/CreatePostForm';
import PostContainer from '../components/PostContainer';
import usePosts from '../services/usePosts';
import useComments from '../services/useComments';
import useUsers from '../services/useUsers';
import { conn, sendMsg } from '../services/useWebsocket';
import { TargetUsers } from '@/services/useTarget';

const Home = ({ loggedIn, id }) => {
  const router = useRouter();
  const { posts, createPost, fetchPosts } = usePosts();
  const { comments, createComment } = useComments();
  const { users, fetchUsers } = useUsers();

  const onButtonClick = async () => {
    if (!loggedIn) {
      router.push('/login')
    }
  };

  useEffect(() => {
    if (loggedIn) {
      fetchPosts(id);
    }
  }, [loggedIn]);

  const handleCreatePost = async (formData) => {
    await createPost(formData);
    fetchPosts(id);

    let List = [];

    const selectedFollowers = formData.get('selectedFollowers');

    if (selectedFollowers) {
      List = selectedFollowers.split(',').map(Number);
    }

    console.log("privacy", formData.get('privacy'))
    switch (formData.get('privacy')) {
      case '0': //PRIVATE
        sendMsg(conn, 0, { value: "New Post" }, 'post', TargetUsers)
        break
      case '1': //PUBLIC
        sendMsg(conn, 0, { value: "New Post" }, 'post', TargetUsers)
        break
      case '2': //SEMI-PRIVATE
        sendMsg(conn, 0, { value: "New Post" }, 'post', List)
        break
      default:
        break
    }
  };

  const handleCreateComment = async (formData) => {
    await createComment(formData);
    fetchPosts(id);
    sendMsg(conn, 0, { value: "New Comment" }, 'comment')
  };




  const handlePostLike = async (postId) => {
    try {
      const response = await fetch('http://localhost:8080/like', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ post_id: postId }),
        credentials: 'include'
      });

      if (response.ok) {
        fetchPosts(id);
      } else {
        console.error('Failed to like the post:', response.statusText);
      }
    } catch (error) {
      console.error('Error while liking the post:', error);
    }
  };


  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}>

      {loggedIn ? (
        <>
          <div className='test'>
            <CreatePostForm handleCreatePost={handleCreatePost} fetchUsers={fetchUsers} id={id} users={users} />
            <PostContainer posts={posts} handleCreateComment={handleCreateComment} handlePostLike={handlePostLike} />
          </div>
        </>
      ) :
        <div className="mainContainer">
          <div className={'titleContainer usn'}>
            <p className='title'>Welcome to</p>
            <div className='logoContainer'>
              <div className='bee'></div>
              <svg className='hexagon' xmlns="http://www.w3.org/2000/svg" width="50%" height="50%" fill="none" viewBox="0 0 868 970">
                <path stroke="#F7F9FF" stroke-width="20" d="m479 22.528 333.013 192.265a90 90 0 0 1 45 77.942v384.53a90 90 0 0 1-45 77.942L479 947.472a90 90 0 0 1-90 0L55.987 755.207a90 90 0 0 1-45-77.942v-384.53a90 90 0 0 1 45-77.942L389 22.528a90 90 0 0 1 90 0Z" />
              </svg>
              <p className='logotext'>HIVE</p>
            </div>
          </div>
          <div className={'buttonContainer'}>
            <input
              className={'inputButton'}
              type="button"
              onClick={onButtonClick}
              value={'Log in'}
            />
            <p className='white minitext'>You are not logged in</p>
          </div>
        </div>
      }

    </div>
  );
}

export default Home;
