import { useState, useEffect } from 'react';

const useEvents = () => {
  const [events, setEvents] = useState([]);

  const fetchEvents = async (id) => {
    try {
      //console.log("fetch for id:", id)
      const response = await fetch(`http://localhost:8080/event?id=${id}`);
      if (response.ok) {
        const data = await response.json();
        setEvents(data);
      } else {
        console.error('Failed to fetch events:', response.statusText);
      }
    } catch (error) {
      console.error('Error fetching events:', error);
    }
  };

  const createEvent = async (formData, id) => {
    //console.log(formData)
    try {
      const response = await fetch(`http://localhost:8080/event?id=${id}`, {
        method: 'POST',
        body: formData,
        credentials: 'include'
      });

      if (response.ok) {
        const data = await response.json();
        console.log(data);
        return data;
      } else {
        console.error('Create a event failed:', response.statusText);
      }
    } catch (error) {
      console.error('Error during creation of a event:', error);
    }
  };


  return { events, createEvent, fetchEvents};
};

export default useEvents;