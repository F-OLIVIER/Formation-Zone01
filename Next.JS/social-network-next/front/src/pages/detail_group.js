import { useRouter } from 'next/router';
import { useState, useEffect } from 'react';
import { getOneGroup } from '@/services/useCreateGroup';
import OneGroup from '@/components/OneGroup';
import PostGroupContainer from '@/components/PostGroupContainer';
import CreatePostGroupForm from '@/components/CreatePostGroupForm';
import EventContainer from '@/components/EventContainer';
import CreateEvent from '@/components/CreateEvent';
import usePosts from '@/services/usePosts';
import useEvents from '@/services/useEvents';
import { conn, sendMsg } from '../services/useWebsocket';
import useComments from '@/services/useComments';


const DetailGroup = (props) => {
    const router = useRouter();
    const { id } = router.query;
    const { postsGroup, createPostGroup, fetchPostsGroup } = usePosts();
    const { createCommentGroup } = useComments();
    const [SingleForm, setSingleForm] = useState(null);
    const {events, createEvent, fetchEvents} = useEvents();
    const [user_id, setUser_id] = useState(props.id)


    useEffect(() => {
        console.log("id: ",id)
        console.log("user_id: ", user_id)
        
        if (!id) return; // Exit early if id is not available
        const fetchData = async () => {
            try {
                const result = await getOneGroup(id);
                if (result.success) {
                    setSingleForm(result.data);
                } else {
                    setError('Failed to get group data: ' + result.message);
                }
            } catch (error) {
                setError('Error during fetching of group data: ' + error.message);
            } finally {
            }
            fetchPostsGroup(id);
            fetchEvents(id)
            console.log(events)
        };

        fetchData();
    }, [id]);

    const handleCreatePostGroup = async (formData) => {
        await createPostGroup(formData, id);
        fetchPostsGroup(id);
        sendMsg(conn, 0, { value: "New PostGroup" }, 'post')
      };

      const handleCreateCommentGroup = async (formData) => {
        await createCommentGroup(formData);
        fetchPostsGroup(id);
        sendMsg(conn, 0, { value: "New Comment" }, 'comment')
      };
    
      const handleEventLike = async (eventId) => {
        try {
          const response = await fetch('http://localhost:8080/comingevent', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ event_id: eventId }),
            credentials: 'include'
          });
    
          if (response.ok) {
            fetchEvents(eventId);
          } else {
            console.error('Failed to like the post:', response.statusText);
          }
        } catch (error) {
          console.error('Error while liking the post:', error);
        }
      };

      const handleEventDisLike = async (eventId) => {
        try {
          const response = await fetch('http://localhost:8080/notcomingevent', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ event_id: eventId }),
            credentials: 'include'
          });
    
          if (response.ok) {
            fetchEvents(eventId);
          } else {
            console.error('Failed to dislike the post:', response.statusText);
          }
        } catch (error) {
          console.error('Error while disliking the post:', error);
        }
      };

      const handleCreateEvent = async (formData) => {
        await createEvent(formData, id);
        fetchEvents(id);
        sendMsg(conn, 0, { value: "New Event" }, 'event')
      };




    return (
        <>
        <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
        <div className='test'>
        <OneGroup SingleForm={SingleForm} />
        <CreatePostGroupForm handleCreatePost={handleCreatePostGroup}/>
        <PostGroupContainer posts={postsGroup} handleCreateGroupComment={handleCreateCommentGroup} />
        <CreateEvent handleCreateEvent={handleCreateEvent}/>
        <EventContainer events={events} handleEventLike={handleEventLike} handleEventDisLike={handleEventDisLike} />
        </div>
        </div>
        </>
    );
};

export default DetailGroup;