import * as React from 'react';
import { useRouter } from 'next/router';
import { styled, useTheme } from '@mui/material/styles';
import toast from 'react-hot-toast';
import Link from 'next/link';
import Box from '@mui/material/Box';
import MuiDrawer from '@mui/material/Drawer';
import MuiAppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import List from '@mui/material/List';
import CssBaseline from '@mui/material/CssBaseline';
import Typography from '@mui/material/Typography';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import HomeIcon from '@mui/icons-material/Home';
import LoginIcon from '@mui/icons-material/Login';
import RegisterIcon from '@mui/icons-material/HowToReg';
import AccountBoxIcon from '@mui/icons-material/AccountBox';
import LogoutIcon from '@mui/icons-material/Logout';
import ChatIcon from '@mui/icons-material/Chat';
import { cookie } from '../services/useCookie';
import Groups2Icon from '@mui/icons-material/Groups2';


const drawerWidth = 240;
const openedMixin = (theme) => ({
  width: drawerWidth,
  transition: theme.transitions.create('width', {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.enteringScreen,
  }),
  overflowX: 'hidden',
});
const closedMixin = (theme) => ({
  transition: theme.transitions.create('width', {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
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
const AppBar = styled(MuiAppBar, {
  shouldForwardProp: (prop) => prop !== 'open',
})(({ theme, open }) => ({
  zIndex: theme.zIndex.drawer + 1,
  transition: theme.transitions.create(['width', 'margin'], {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  ...(open && {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
}));
const DrawerContainer = styled(Box)({
  zIndex: 1, // Higher z-index to ensure content appears above the drawer
  position: 'relative', // Make sure position is relative for z-index to work
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
  const theme = useTheme();
  const [open, setOpen] = React.useState(false);
  const router = useRouter(); // Using Next.js's useRouter hook
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
            style: {backgroundColor: 'rgba(0,255,34,0.5)', color: 'white'},
            icon: '👏',
          });
          setLoggedIn(false);
          router.push('/');
        } else {
          toast.error('Logout failed.Error: ' + responseData.message, {
            duration: 4000,
            position: 'top-center',
            style: {backgroundColor: 'rgba(255,0,0,0.5)', color: 'white'},
          });
        }
      } catch (error) {
        console.error(error);
        toast.error('Error during logout: ' + error.message, {
          duration: 4000,
          position: 'top-center',
          style: {backgroundColor: 'rgba(255,0,0,0.5)', color: 'white'},
        });
      }
    } else {
      router.push('/login');
    }
  };
  const handleDrawerOpen = () => {
    setOpen(true);
  };
  const handleDrawerClose = () => {
    setOpen(false);
  };
  return (
    <DrawerContainer>
      <CssBaseline />
      <AppBar position="fixed" style={{ background: '#2E3B55' }} open={open}>
        <Toolbar>
          <IconButton
            color="inherit"
            aria-label="open drawer"
            onClick={handleDrawerOpen}
            edge="start"
            sx={{
              marginRight: 5,
              ...(open && { display: 'none' }),
            }}
          >
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" noWrap component="div">
            Social Network
          </Typography>
        </Toolbar>
      </AppBar>
      <Drawer variant="permanent" open={open}>
        <DrawerHeader>
          <IconButton onClick={handleDrawerClose}>
            {theme.direction === 'rtl' ? <ChevronRightIcon /> : <ChevronLeftIcon />}
          </IconButton>
        </DrawerHeader>
        <List>
          {loggedIn ? ( // Afficher les boutons de menu en fonction de l'état de connexion
            <>
            <ListItemButton selected={router.pathname === '/'} component={Link} to="/">
                <ListItemIcon>
                  <HomeIcon />
                </ListItemIcon>
                <ListItemText primary="Home" />
              </ListItemButton>
              <ListItemButton 
    selected={router.pathname === '/user?id=' + id.toString()} 
    component={Link} 
    href={'/user?id=' + id.toString()}
>                <ListItemIcon>
                  <AccountBoxIcon />
                </ListItemIcon>
                <ListItemText primary="Profile" />
              </ListItemButton>
              <ListItemButton 
    selected={router.pathname === '/chat?id=' + id.toString()} 
    component={Link} 
    href={'/chat?id=' + id.toString()}
>                <ListItemIcon>
                  <ChatIcon />
                </ListItemIcon>
                <ListItemText primary="Chat" />
              </ListItemButton>

              <ListItemButton selected={router.pathname === '/'} component={Link} to="/group">
                <ListItemIcon>
                  <Groups2Icon/>
                </ListItemIcon>
                <ListItemText primary="Group" />
              </ListItemButton>

              <ListItemButton onClick={onButtonClick}>
                <ListItemIcon>
                  <LogoutIcon />
                </ListItemIcon>
                <ListItemText primary="Logout" />
              </ListItemButton>


            </>
          ) : (
            <>
              <ListItemButton selected={router.pathname === '/'} component={Link} to="/">
                <ListItemIcon>
                  <HomeIcon />
                </ListItemIcon>
                <ListItemText primary="Home" />
              </ListItemButton>
              <ListItemButton selected={router.pathname === '/login'} component={Link} href="/login">
                <ListItemIcon>
                  <LoginIcon />
                </ListItemIcon>
                <ListItemText primary="Login" />
              </ListItemButton>
              <ListItemButton selected={router.pathname === '/register'} component={Link} href="/register">
                <ListItemIcon>
                  <RegisterIcon />
                </ListItemIcon>
                <ListItemText primary="Register" />
              </ListItemButton>
            </>
          )}
        </List>
      </Drawer>
    </DrawerContainer>
  );
}