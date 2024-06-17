import useUsers from '../services/useUsers';
import ProfileContainer from '../components/ProfileContainer';
import useComments from '../services/useComments';
import {conn, sendMsg} from '../services/useWebsocket';


// import { useRouter } from 'next/router';


const User = ({id}) => {
    const {users, userPosts, fetchUsers, fetchUserPosts} = useUsers();
    const {createComment} = useComments();


    const togglePrivacy = async () => {
        try {
            const response = await fetch('http://localhost:8080/user', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                credentials: "include",
            });
            if (response.ok) {
                fetchUsers(response.ID);
            }
            if (!response.ok) {
                throw new Error('Failed to update privacy setting');
            }
        } catch (error) {
            console.error('Error updating privacy setting:', error.message);
        }
    };

    const fetchFollow = async (dataToSend) => {
        try {
            const response = await fetch('http://localhost:8080/follow', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                credentials: "include",
                body: JSON.stringify(dataToSend),
            });
            if (response.ok) {
                fetchUsers(response.ID);
            }
            if (!response.ok) {
                throw new Error('Failed to update privacy setting');
            }
        } catch (error) {
            console.error('Error updating privacy setting:', error.message);
        }
    }

    const follow = async () => {
        const UserId_Following = users.id;
        const UserId_Follower = id;
        const Action = "follow";
        const dataToSend = {UserId_Following, UserId_Follower, Action};
        fetchFollow(dataToSend);
        users.ListFollowersToValidate && users.ListFollowersToValidate.some(user => user.id === id);

        let Target = [];
        Target.push(users.id);

        if (users.listfollowers && users.listfollowers.some(user => user.id === id)) {
            sendMsg(conn, 0, {value: "New Follow"}, 'stop_follow', Target);
        } else if (users.ListFollowersToValidate && users.ListFollowersToValidate.some(user => user.id === id)){
            sendMsg(conn, 0, {value: "Cancel Follow"}, 'cancel_follow', Target);
        } else {
            sendMsg(conn, 0, {value: "Stop Follow"}, 'follow', Target);
        }

    };

    const validatefollow = async (validated, idnewfollower) => {
        const UserId_Following = users.id;
        const UserId_Follower = idnewfollower;
        const ValidateFollow = validated;
        const Action = "validatefollow";
        const dataToSend = {UserId_Following, UserId_Follower, ValidateFollow, Action};
        fetchFollow(dataToSend);
    };

    const handleCreateComment = async (formData) => {
        await createComment(formData);
        fetchUserPosts()
        console.log(formData)
    };

    const handlePostLike = async (postId) => {
        try {
            const response = await fetch('http://localhost:8080/like', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({post_id: postId}),
                credentials: 'include'
            });

            if (response.ok) {
                fetchUserPosts()
            } else {
                console.error('Failed to like the post:', response.statusText);
            }
        } catch (error) {
            console.error('Error while liking the post:', error);
        }
    };

  return (
    <div className='test'>
      <ProfileContainer users={users} userPosts={userPosts} follow={follow} validatefollow={validatefollow} togglePrivacy={togglePrivacy} id={id} handleCreateComment={handleCreateComment} handlePostLike={handlePostLike} />
    </div>
    );
};

export default User;