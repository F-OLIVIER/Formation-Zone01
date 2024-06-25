import React from 'react';
import { Button, CardActions } from '@mui/material';

const EventContainer = ({ events, handleEventLike, handleEventDisLike }) => {

  return (
    <div>
      <h2 className='pagetitle'>Events</h2>
      <div className='postscontainer eventsContainer'>
        {events && (
          <ul>
            {events.map(event => (
              <li key={event.id}>
                <div className='post'>
                  <div>
                    <div className='post-header'>
                      <div className='post-username-container'>
                        <p className='post-username'>
                          {event.title}
                        </p>
                      </div>
                    </div>
                    <p className="post-content">
                      {event.description}
                    </p>
                    <p className="eventdate">
                      Date: {event.date}
                    </p>
                    <CardActions>
                      <p>
                        Interested ?
                      </p>
                      <Button
                      className='createEventbtn yesnobtn'
                        variant="contained"
                        onClick={() => handleEventLike(event.id)}
                      >
                        ✅
                      </Button>
                      <Button
                      className='createEventbtn yesnobtn'

                        variant="contained"
                        onClick={() => handleEventDisLike(event.id)}
                      >
                        ❌
                      </Button>
                    </CardActions>
                    <div className='commingornot'>
                    <p>
                      Come : {event.coming}
                    </p>
                    <p>
                      Don't come : {event.notcoming}
                    </p>
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

export default EventContainer;