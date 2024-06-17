import * as React from 'react';
import { useRouter } from 'next/router';
import { styled, useTheme } from '@mui/material/styles';
import toast from 'react-hot-toast';
import Link from 'next/link';
import Box from '@mui/material/Box';
import MuiDrawer from '@mui/material/Drawer';
import List from '@mui/material/List';
import CssBaseline from '@mui/material/CssBaseline';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import { cookie } from '../services/useCookie';
import setTransiName from '@/services/setTransiName';
let hoveredmenu = false;

const drawerWidth = 240;
const openedMixin = (theme) => ({
  width: drawerWidth,
  transition: theme.transitions.create('width', {
    easing: 'cubic-bezier(0.7, 0, 0.3, 1)',
    duration: '400ms',
  }),
  overflowX: 'hidden',
});
const closedMixin = (theme) => ({
  transition: theme.transitions.create('width', {
    easing: 'cubic-bezier(0.7, 0, 0.3, 1)',
    duration: '400ms',
  }),
  overflowX: 'hidden',
  width: `calc(${theme.spacing(7)} + 1px)`,
  [theme.breakpoints.up('sm')]: {
    width: `calc(${theme.spacing(8)} + 1px)`,
  },
});
const DrawerHeader = styled('div')(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'flex-end',
  padding: theme.spacing(0, 1),
  ...theme.mixins.toolbar,
}));

const Logo = () => (
  <StyledLogo className="headerLogo">
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="-30 -50 935 1090">
      <path stroke="#F7F9FF" strokeWidth="50" d="m479 22.528 333.013 192.265a90 90 0 0 1 45 77.942v384.53a90 90 0 0 1-45 77.942L479 947.472a90 90 0 0 1-90 0L55.987 755.207a90 90 0 0 1-45-77.942v-384.53a90 90 0 0 1 45-77.942L389 22.528a90 90 0 0 1 90 0Z" />
    </svg>
  </StyledLogo>
);

const StyledLogo = styled('div')`
  display: flex;
  justify-content: center;
  align-items: center;
`;

const DrawerContainer = styled(Box)({
  zIndex: 1,
  position: 'relative',
});
const Drawer = styled(MuiDrawer, { shouldForwardProp: (prop) => prop !== 'open' })(
  ({ theme, open }) => ({
    width: drawerWidth,
    flexShrink: 0,
    whiteSpace: 'nowrap',
    boxSizing: 'border-box',
    ...(open && {
      ...openedMixin(theme),
      '& .MuiDrawer-paper': openedMixin(theme),
    }),
    ...(!open && {
      ...closedMixin(theme),
      '& .MuiDrawer-paper': closedMixin(theme),
    }),
  }),
);
export default function MiniDrawer({ loggedIn, setLoggedIn, id }) {
  const router = useRouter();
  const onButtonClick = async () => {
    if (loggedIn) {
      try {
        const data = {
          test: 1,
        };
        const responseData = await cookie(data);
        if (responseData.success === true) {
          toast.success('Logout successful!', {
            duration: 4000,
            position: 'top-center',
            style: { backgroundColor: 'rgba(0,255,34,0.5)', color: 'white' },
            icon: '👏',
          });
          setLoggedIn(false);
          router.push('/');
        } else {
          toast.error('Logout failed.Error: ' + responseData.message, {
            duration: 4000,
            position: 'top-center',
            style: { backgroundColor: 'rgba(255,0,0,0.5)', color: 'white' },
          });
        }
      } catch (error) {
        console.error(error);
        toast.error('Error during logout: ' + error.message, {
          duration: 4000,
          position: 'top-center',
          style: { backgroundColor: 'rgba(255,0,0,0.5)', color: 'white' },
        });
      }
    } else {
      router.push('/login');
    }
  };

  const handleHovreMenu = () => {
    if (hoveredmenu) {
      hoveredmenu = false;
      document.querySelector('.headerLogo').classList.remove('active');
      document.querySelector('.headertext').classList.remove('active');
      document.querySelectorAll('.Menutxt').forEach((item) => {
        item.classList.remove('active');
      })
    } else {
      hoveredmenu = true;
      document.querySelector('.headerLogo').classList.add('active');
      document.querySelector('.headertext').classList.add('active');
      document.querySelectorAll('.Menutxt').forEach((item) => {
        item.classList.add('active');
      });
    }
  }

  const handleDrawerSelected = (event) => {
    const listItems = document.querySelectorAll('.Menu');
    listItems.forEach((item) => {
      item.classList.remove('selected');
    });
    event.currentTarget.classList.add('selected');
  };


  return (
    <DrawerContainer>
      <CssBaseline />
      <Drawer classes={{ paper: 'drawer' }} variant="permanent" onMouseEnter={() => {
        handleHovreMenu();
        hoveredmenu = true;
      }}
        onMouseLeave={() => {
          handleHovreMenu();
          hoveredmenu = false;
        }}>
        <DrawerHeader className='drawerHeader'>
          <Logo />
          <p className='headertext logotext'>HIVE</p>
        </DrawerHeader>
        <span className='headerSeparation'></span>
        <List className='listMenu'>
          {loggedIn ? (
            <>
              <ListItemButton className='Menu' component={Link} href="/" onClick={setTransiName()}>
                <ListItemIcon>
                  <div className='HomeIcon MenuIcon'>
                  </div>
                </ListItemIcon>
                <span className='Menutxt'>Home</span>
              </ListItemButton>

              <ListItemButton className='Menu' component={Link} href={'/user?id=' + id.toString()}>
                <ListItemIcon>
                  <div className='HomeIcon MenuIcon'>
                  </div>
                </ListItemIcon>
                <span className='Menutxt'>Profile</span>
              </ListItemButton>

              <ListItemButton className='Menu' component={Link} href={'/chat?id=' + id.toString()}>
                <ListItemIcon>
                  <div className='HomeIcon MenuIcon'>
                  </div>
                </ListItemIcon>
                <span className='Menutxt'>Chat</span>
              </ListItemButton>

               <ListItemButton className='Menu' component={Link} href={'/group?id='}>
                <ListItemIcon>
                  <div className='HomeIcon MenuIcon'>
                  </div>
                </ListItemIcon>
                <span className='Menutxt'>Group</span>
              </ListItemButton>
              
              <ListItemButton className='Menu' onClick={onButtonClick}>
                <ListItemIcon>
                  <div className='HomeIcon MenuIcon'>
                  </div>
                </ListItemIcon>
                <span className='Menutxt'>Log out</span>
              </ListItemButton>


            </>
          ) : (
            <>
              <ListItemButton className='Menu' component={Link} href="/" onClick={setTransiName()}>
                <ListItemIcon>
                  <div className='HomeIcon MenuIcon'>
                  </div>
                </ListItemIcon>
                <span className='Menutxt'>Home</span>
              </ListItemButton>
              
              <ListItemButton className='Menu' component={Link} href="/login">
                <ListItemIcon>
                  <div className='LoginIcon MenuIcon'>
                    <span className='LoginIcon LoginArrow'></span>
                  </div>
                </ListItemIcon>
                <span className='Menutxt'>Login</span>
              </ListItemButton>

              <ListItemButton className='Menu' component={Link} href="/register">
                <ListItemIcon>
                  <div className='RegisterIcon MenuIcon'>
                    <span className='RegisterIcon RegisterBeeIcon'></span>
                  </div>
                </ListItemIcon>
                <span className='Menutxt'>Sign up</span>
              </ListItemButton>
            </>
          )}
        </List>
      </Drawer>
    </DrawerContainer>
  );
}