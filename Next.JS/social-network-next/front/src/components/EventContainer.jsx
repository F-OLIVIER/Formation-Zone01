import React from 'react';
import { Button, Card, CardContent, CardActions, Typography } from '@mui/material';
import { useRouter } from 'next/router';



const EventContainer = ({ events, handleEventLike, handleEventDisLike }) => {
  const router = useRouter();

  return (
    <div>
      <h2>Events</h2>
      {events && (
        <ul>
          {events.map(event => (
            <li key={event.id}>
              <Card>
                <CardContent>
                  <Typography variant="h5" component="h2">
                    {event.title}
                  </Typography>
                  <Typography variant="body2" color="textSecondary" component="p">
                    {event.description}
                  </Typography>
                  <Typography variant="body2" color="textSecondary" component="p">
                    Date: {event.date}
                  </Typography>
                  <CardActions>
                  <Typography variant="body2" color="textSecondary" component="p">
                    Are you going ?
                  </Typography>
                  <Button
                      variant="contained"
                      onClick={() => handleEventLike(event.id)}
                    >
                      Yes
                    </Button>
                    <Button
                      variant="contained"
                      onClick={() => handleEventDisLike(event.id)}
                    >
                      No
                    </Button>
                  </CardActions>
                  <Typography variant="body2" color="textSecondary" component="p">
                    Coming: {event.coming}
                  </Typography>
                  <Typography variant="body2" color="textSecondary" component="p">
                    NotComing: {event.notcoming}
                  </Typography>
                </CardContent>
              </Card>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default EventContainer;