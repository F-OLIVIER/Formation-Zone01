import React from 'react';
import { useRouter } from 'next/router';
import PostContainer from '../components/PostContainer';

const ProfileContainer = ({ users, togglePrivacy, userPosts, follow, validatefollow, id, handleCreateComment, handlePostLike }) => {
    //console.log('users in ProfileContainer :', users)
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
        <div className='follower-list-container'>
            {list.map(follower => (
                <div key={follower.id} className='follower-container'>
                    <div className='follower-pp pp'></div>
                    <a className='follower-name' onClick={() => router.push(`/user?id=${follower.id}`)}>{follower.firstname} {follower.lastname}</a>
                    <p className='follower-name nickname'>{follower.nickname}</p>
                </div>
            ))}
        </div>
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
    if (userPosts !== null) {
        filteredPosts = userPosts.filter(post => post.user_id === users.id);

    }

    // Test si l'utilisateur courant et un folowers du profils qu'il regarde pour savoir en fonction de la privacity si on l'affiche ou non
    let isFollowers = false
    if (users.listfollowers) {
        isFollowers = users.listfollowers.some(follower => follower.id === id)
    }

    // Déterminer si le profil doit être affiché
    // Si profils Privée on affiche uniquement si l'utilisateur et dans les followers de l'utilisateur
    // Si profils Public on affiche tout
    const displayProfile = (users.privateprofile == 0 && isFollowers) || users.privateprofile == 1 || users.id === id;

    return (
        <>
            <h2 className='pagetitle'>Profile</h2>


            <div className='profilecontainer'>
                <p className='pagetitle white profilename'>{users.nickname}</p>
                <div className='privacy'>

                    <p className='white profileis'>Profile :</p>
                    <button className='toggleprivacy' onClick={togglePrivacy}>{getPrivacyLabel(users.privateprofile)}</button>
                </div>
                {displayProfile &&
                    <>
                        {users.avatar && (
                            <img
                                className='pp-profile'
                                src={`data:image/jpeg;base64,${users.avatar}`}
                                alt="Selected Image"
                            />
                        )}
                        {/*                         <p className='white'>ID: {users.id}</p> */}
                        <p className='white profileusername'>{users.firstname} {users.lastname}</p>
                        <p className='white'>Email: {users.email}</p>
                        <p className='white'>Date of Birth: {users.dateofbirth}</p>
                        <p className='white aboutme'>{users.aboutme}</p>
                        {/*                         <p className='white'>Point of Interest: {users.pointofinterest}</p> */}
                    </>
                }
                {users.id === id ? (
                    <div className='follow-section'>
                        <p className='white followtitle'>Followers</p>
                        {users.listfollowers && (
                            <div>{followList("List of followers", users.listfollowers)}</div>
                        )}
                        <p className='white followtitle'>Followings</p>

                        {users.listfollowings && (
                            <div>{followList("List of followings", users.listfollowings)}</div>
                        )}

                        {users.ListFollowersToValidate && (
                            <div>
                                <p className='white followtitle'>Accept/Reject Demands</p>
                                <div className='white followtitle'>{followRequestList(users.ListFollowersToValidate)}</div>
                            </div>
                        )}
                    </div>
                ) : (
                    <>
                        {followButton()}

                        {/* {users.privateprofile == 1 && users.listfollowers && ( */}
                        {displayProfile && (
                            <>
                            <p className='white followtitle'>Followers</p>

                            <div>{followList("List of followers", users.listfollowers)}</div>
                            </>
                        )}


                        {displayProfile && (
                            <>
                                                    <p className='white followtitle'>Followings</p>
                            <div>{followList("List of followings", users.listfollowings)}</div>
                            </>
                        )}
                    </>
                )}


                {displayProfile && filteredPosts !== null &&
                    <>
                        {filteredPosts && (
                            <PostContainer posts={filteredPosts} handleCreateComment={handleCreateComment} handlePostLike={handlePostLike} />
                        )}
                    </>
                }
            </div>
        </>

    );
};

export default ProfileContainer;