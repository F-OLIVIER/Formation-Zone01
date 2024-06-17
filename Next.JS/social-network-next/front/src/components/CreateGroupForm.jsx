import React, { useState, useEffect } from 'react';
import { TextField, Button, Select, MenuItem } from '@mui/material';

const CreateGroupForm = ({ handleCreateGroup }) => {
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [users, setUsers] = useState([]);
    const [selectedUser, setSelectedUser] = useState('');

    const fetchUsers = async () => {
        try {
            const response = await fetch('/listofusers', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            if (!response.ok) {
                throw new Error('Error fetching users');
            }

            const data = await response.json();
            return data;
        } catch (error) {
            throw new Error('Error fetching users:', error.message);
        }
    };

    useEffect(() => {
        fetchUsers()
            .then((data) => {
                setUsers(data);
            })
            .catch((error) => {
                console.error('Error fetching users:', error);
            });
    }, []);

    const handleSubmit = (e) => {
        e.preventDefault();
        const formData = new FormData();
        formData.append('title', title);
        formData.append('description', description);
        formData.append('selectedUser', selectedUser);
        handleCreateGroup(formData);
        setTitle('');
        setDescription('');
        setSelectedUser('');
    };

    return (
        <form onSubmit={handleSubmit} encType="multipart/form-data">
            <h2>Create a New Group</h2>
            <TextField
                label="Title"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
            />
            <TextField
                label="Description"
                value={description}
                onChange={(e) => setTitle(e.target.value)}
                multiline
            />
            <Select
                value={selectedUser}
                onChange={(e) => setSelectedUser(e.target.value)}
                label="Select User"
            >
                {users.map((user) => (
                    <MenuItem key={user.id} value={user.id}>
                        {user.nickname || `${user.firstName} ${user.lastName}`}
                    </MenuItem>
                ))}
            </Select>
            <Button type="submit" variant="contained">
                Create Group
            </Button>
        </form>
    );
};

export default CreateGroupForm;
