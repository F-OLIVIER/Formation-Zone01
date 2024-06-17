import React, {useEffect, useState} from 'react';
import {fetchNotification} from "@/services/useFetchNotif";
import toast from "react-hot-toast";
import useUsers from "@/services/useUsers";
//import {acceptGroupNotification} from "@/services/useCreateGroup";

const NotificationFetcher = (props) => {
    const {id, setNotifications, notifications} = props;
    const [deletedNotifications, setDeletedNotifications] = useState([]);
    const {users, userPosts, fetchUsers, fetchUserPosts} = useUsers();


    const fetchFollow = async (dataToSend, notifId, followId) => {
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
                toast.success(
                    'Follow accepted',
                    {}
                )
                removeNotification(notifId, "follow", followId);
            }
            if (!response.ok) {
                throw new Error('Failed to update privacy setting');
            }
        } catch (error) {
            console.error('Error updating privacy setting:', error.message);
        }
    }
    const validatefollow = async (validated, idnewfollower, notifId, followId) => {
        const UserId_Following = props.id;
        const UserId_Follower = idnewfollower;
        const ValidateFollow = validated;
        const Action = "validatefollow";
        const dataToSend = {UserId_Following, UserId_Follower, ValidateFollow, Action};
        console.log(validated)
        console.log(idnewfollower)
        console.log(dataToSend)
        fetchFollow(dataToSend, notifId, followId);
    };

    const acceptGroupNotification = async (idNotif, idGroup, idUser, category = undefined) => {
        console.log(idGroup)
        console.log(idUser)
        try {
            const data = {
                id: idGroup,
                idwhoisinvited: idUser
            };
            const response = await fetch('http://localhost:8080/accept-group-notification', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });

            if (!response.ok) {
                throw new Error('Failed to delete notification');
            }

            toast.success(
                'Group invitation accepted',
                {}
            )
            if (category === "askgroup"){
                removeNotification(idNotif, "askgroup", idGroup);
            }else{
                removeNotification(idNotif, "group", idGroup);
            }

        } catch (error) {
            console.error('Error deleting notification:', error);
        }
    }

    const deleteNotification = async (notifIdCategory, category) => {
        try {
            const response = await fetch('http://localhost:8080/delete-notification', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    notification_id: notifIdCategory,
                    category: category
                }),
            });

            if (!response.ok) {
                throw new Error('Failed to delete notification');
            }

            // Ajoutez ici toute logique supplémentaire à effectuer après la suppression réussie de la notification

        } catch (error) {
            console.error('Error deleting notification:', error);
        }
    };


    const removeNotification = async (notificationId, category, notifIdCategory) => {
        try {
            setDeletedNotifications([...deletedNotifications, notificationId]);
            await deleteNotification(notifIdCategory, category);
        } catch (error) {
            console.error('Error deleting notification:', error);
        }
    };


    // Sort notifications by time
    const sortedNotifications = notifications.slice().reverse().filter(notif => !deletedNotifications.includes(notif.id));

    return (
        <div className="notification-container">
            <h1 className="notification-headers">Notification Fetcher</h1>
            <div id="notifications">
                {sortedNotifications && sortedNotifications.length > 0 ? sortedNotifications.map((notif) => (
                    <div key={notif.id} className="notification">
                        <div>Id : {notif.id}</div>
                        <div className="emoji-container">
                            {notif.category === 'Follow' && (
                                <>
                                    <button className="close-button"
                                            onClick={() => removeNotification(notif.id, "follow", notif.follow.id,)}>×
                                    </button>
                                    <span role="img" aria-label="Follow Emoji">➕</span>
                                </>
                            )}
                            {notif.category === 'AskGroup' && (
                                <>
                                    <button className="close-button"
                                            onClick={() => removeNotification(notif.id, "askgroup", notif.group.IdGroup)}>×
                                    </button>
                                    <span role="img" aria-label="AskGroup Emoji">❓</span>
                                </>
                            )}
                            {notif.category === 'GroupMP' && (
                                <>
                                    <button className="close-button"
                                            onClick={() => removeNotification(notif.id, "groupmp", notif.message.ID)}>×
                                    </button>
                                    <span role="img" aria-label="GroupMp Emoji">👩‍❤️‍👨</span>
                                </>
                            )}
                            {notif.category === 'Post' && (
                                <>
                                    <button className="close-button"
                                            onClick={() => removeNotification(notif.id, "post", notif.post.id)}>×
                                    </button>
                                    <span role="img" aria-label="Follow Emoji">📑</span>
                                </>
                            )}
                            {notif.category === 'Group' && (
                                <>
                                    <button className="close-button"
                                            onClick={() => removeNotification(notif.id, "group", notif.group.id)}>×
                                    </button>
                                    <span role="img" aria-label="Group Emoji">👥</span>
                                </>
                            )}
                            {notif.category === 'MP' && (
                                <>
                                    <button className="close-button"
                                            onClick={() => removeNotification(notif.id, "mp", notif.message.id)}>×
                                    </button>
                                    <span role="img" aria-label="Message Emoji">💬</span>
                                </>
                            )}
                            {notif.category === 'Comment' && (
                                <>
                                    <button className="close-button"
                                            onClick={() => removeNotification(notif.id, "comment", notif.comment.id)}>×
                                    </button>
                                    <span role="img" aria-label="Comment Emoji">📥</span>
                                </>
                            )}
                        </div>
                        <div className="notification-details">
                            {notif.category === 'GroupMP' && (
                                <>
                                    <p className="notification-group">Message ID: {notif.message.ID}</p>
                                    <div className="notification-date">{notif.message.Date}</div>
                                    <p className="notification-title">{notif.user.firstname} {notif.user.lastname} a
                                        envoyé un message dans le groupe "<b>{notif.group.Title}</b>"
                                        : <b>{notif.message.Content}</b> !</p>
                                </>
                            )}
                            {notif.category === 'AskGroup' && (
                                <>
                                    <p className="notification-group">Group ID: {notif.group.IdGroup}</p>
                                    <div className="notification-date">{notif.group.date}</div>
                                    <p className="notification-title">{notif.user.firstname} {notif.user.lastname} a
                                        demandé à rejoindre votre groupe <b>{notif.group.Title}</b> !</p>
                                </>
                            )}
                            {notif.category === 'Follow' && (
                                <>
                                    <p className="notification-group">Follow ID: {notif.follow.id}</p>
                                    <div className="notification-date">{notif.follow.datefollow}</div>
                                    {notif.follow.validatefollow ? (
                                        <p className="notification-title">
                                            <b>{notif.user.firstname} {notif.user.lastname} </b>demande à vous suivre !
                                        </p>
                                    ) : (
                                        <p className="notification-title">
                                            <b>{notif.user.firstname} {notif.user.lastname} </b>vous a
                                            suivis !</p>)}
                                </>
                            )}
                            {notif.category === 'Post' && (
                                <>
                                    <p className="notification-group">Post ID: {notif.post.id}</p>
                                    <div className="notification-date">{notif.post.date}</div>
                                    <p className="notification-title">{notif.user.firstname} {notif.user.lastname} a
                                        crée un post <b>{notif.post.title}</b> !</p>
                                </>
                            )}
                            {notif.category === 'Group' && (
                                <>
                                    <p className="notification-group">Group ID: {notif.group.IdGroup}</p>
                                    <div className="notification-date">{notif.group.date}</div>
                                    <p className="notification-title">{notif.user.firstname} {notif.user.lastname} vous
                                        a invité au groupe <b>{notif.group.Title}</b> !</p>
                                </>
                            )}
                            {notif.category === 'MP' && (
                                <>
                                    <p className="notification-group">Message ID: {notif.message.id}</p>
                                    <div className="notification-date">{notif.message.date}</div>
                                    <p className="notification-title">{notif.user.firstname} {notif.user.lastname} vous
                                        a envoyé un message privé :</p>
                                    <div className="comment-content">
                                        <b><p>{notif.message.content}</p></b>
                                    </div>
                                </>
                            )}
                            {notif.category === 'Comment' && (
                                <>
                                    <p className="notification-group">Post ID: {notif.comment.post_id}</p>
                                    <div className="notification-date">{notif.comment.date}</div>
                                    <p className="notification-title">{notif.user.firstname} {notif.user.lastname} a
                                        commenté :</p>
                                    <div className="comment-content">
                                        <b><p>{notif.comment.content}</p></b>
                                    </div>
                                </>
                            )}
                        </div>
                        <div className="group-buttons">
                            {notif.category === 'Follow' && notif.follow.validatefollow && (
                                <button className="accept-button"
                                        onClick={() => validatefollow(true, notif.user.id, notif.id, notif.follow.id)}>Accepter</button>
                            )}
                            {notif.category === 'Group' && (
                                <>
                                    <button className="accept-button"
                                            onClick={() => acceptGroupNotification(notif.id, notif.group.IdGroup, props.id)}>Accept
                                    </button>
                                    <button className="reject-button"
                                            onClick={() => removeNotification(notif.id, "group", notif.group.IdGroup)}>Reject
                                    </button>
                                </>
                            )}
                            {notif.category === 'AskGroup' && (
                                <>
                                    <button className="accept-button"
                                            onClick={() => acceptGroupNotification(notif.id, notif.group.IdGroup, notif.user.id, "askgroup")}>Accept
                                    </button>
                                    <button className="reject-button"
                                            onClick={() => removeNotification(notif.id, "askgroup", notif.group.IdGroup)}>Reject
                                    </button>
                                </>
                            )}
                        </div>
                    </div>
                )) : (
                    <p className="no-notifications">No notifications available</p>
                )}
            </div>
        </div>
    );
};

export default NotificationFetcher;
