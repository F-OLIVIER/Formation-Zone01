import React from 'react';
import { useRouter } from 'next/router';
import PostContainer from '../components/PostContainer';

const ProfileContainer = ({ users, togglePrivacy, userPosts, follow, validatefollow, id, handleCreateComment, handlePostLike }) => {
    // console.log('users in ProfileContainer :', users)
    const router = useRouter();
    const getPrivacyLabel = (privateprofile) => {
        privateprofile = parseInt(privateprofile);
        switch (privateprofile) {
            case 0:
                return 'Private';
            case 1:
                return 'Public';
            default:
                return 'Unknown';
        }
    };

    const followButton = () => {
        const isFollowing = users.listfollowers && users.listfollowers.some(user => user.id === id);
        const isPendingValidation = users.ListFollowersToValidate && users.ListFollowersToValidate.some(user => user.id === id);
        console.log('isPendingValidation : ', isPendingValidation)
        let buttonText;
        if (isPendingValidation) {
            buttonText = 'Cancel follow request';
        } else {
            buttonText = isFollowing ? 'Stop follow' : 'Follow';
        }

        return (
            <div>
                {isPendingValidation && <p>Waiting to validation</p>}
                <button onClick={follow}>{buttonText}</button>
            </div>
        );
    };

    const followList = (title, list) => (
        <details>
            <summary>{title}</summary>
            {list.map(follower => (
                <div key={follower.id}>
                    <a onClick={() => router.push(`/user?id=${follower.id}`)}>• {follower.firstname} {follower.lastname} {follower.nickname ? `(${follower.nickname})` : ''}</a>
                </div>
            ))}
        </details>
    );

    const followRequestList = (list) => (
        <div>
            {list.map(follower => (
                <div key={follower.id}>
                    <p>• {follower.firstname} {follower.lastname} {follower.nickname ? `(${follower.nickname})` : ''}</p>
                    <button onClick={() => validatefollow(true, follower.id)}>Accept follower</button>
                    <button onClick={() => validatefollow(false, follower.id)}>Reject follower</button>
                </div>
            ))}
        </div>
    );

    let filteredPosts
        if (userPosts !== null )  {
           filteredPosts =  userPosts.filter(post => post.user_id === id);

        }



    return (
        <div>

            <h2>Infos</h2>

            <p>Nickname: {users.nickname}</p>
            <p>Private Profile: {getPrivacyLabel(users.privateprofile)}</p>
            {users.privateprofile === 1 || users.id === id &&
                <>
                    <p>ID: {users.id}</p>
                    <p>Email: {users.email}</p>
                    <p>Firstname: {users.firstname}</p>
                    <p>Lastname: {users.lastname}</p>
                    <p>Date of Birth: {users.dateofbirth}</p>
                    <p>Avatar: {users.avatar}</p>
                    <p>About Me: {users.aboutme}</p>
                    <p>Point of Interest: {users.pointofinterest}</p>
                </>
            }
            {users.id === id ? (
                <>
                    <button onClick={togglePrivacy}>Toggle Privacy</button>
                    <p>Follow Section</p>
                    {users.listfollowers && (
                        <div>{followList("List of followers", users.listfollowers)}</div>
                    )}
                    {users.listfollowings && (
                        <div>{followList("List of followings", users.listfollowings)}</div>
                    )}

                    {users.ListFollowersToValidate && (
                        <div>
                            <p>Accept/Reject Demands</p>
                            <div>{followRequestList(users.ListFollowersToValidate)}</div>
                        </div>
                    )}
                </>
            ) : (
                <>
                    {followButton()}

                    {users.privateprofile == 1 && users.listfollowers && (
                        <div>{followList("List of followers", users.listfollowers)}</div>
                    )}

                    {users.privateprofile == 1 && users.listfollowings && (
                        <div>{followList("List of followings", users.listfollowings)}</div>
                    )}
                </>
            )}


            {users.privateprofile === 1 || users.id === id && filteredPosts !== null &&
                <>
                    {filteredPosts && (
                        <PostContainer posts={filteredPosts} handleCreateComment={handleCreateComment} handlePostLike={handlePostLike} />
                    )}
                </>
            }
        </div>
    );
};

export default ProfileContainer;