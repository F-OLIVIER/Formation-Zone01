import { useState, useEffect } from 'react';
import { useRouter } from 'next/router';

const useUsers = () => {
    const [users, setUsers] = useState([]);
    const [userPosts, setUserPosts] = useState([]);
    const router = useRouter();

  const fetchUsers = async (ID) => {
    try {
      let id = router.query.id;
      if (!id) {
        id = ID;
      }; 
      const response = await fetch(`http://localhost:8080/user?id=${id}`, {
        credentials: 'include'
      });
      if (response.ok) {
        const data = await response.json();
        console.log("test: data in fetchUsers : ", data)
        setUsers(data);
      } else {
        console.error('Failed to fetch users:', response.statusText);
      }
    } catch (error) {
      console.error('Error fetching users:', error);
    }
  };

    const fetchUserPosts = async () => {
      try {
        const response = await fetch(`http://localhost:8080/post`, {
          credentials: 'include'
        });
        if (response.ok) {
          const data = await response.json();
          setUserPosts(data);
        } else {
          console.error('Failed to fetch user posts:', response.statusText);
        }
      } catch (error) {
        console.error('Error fetching user posts:', error);
      }
    };

    useEffect(() => {
      fetchUsers();
      fetchUserPosts();// Fetch users when component mounts or when id changes
    }, [router.query.id]);

    return { users, userPosts, fetchUsers, fetchUserPosts };
};

export default useUsers;