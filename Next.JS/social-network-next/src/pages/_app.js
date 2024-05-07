import '../styles/globals.css'
import { useEffect, useState, useRef } from 'react'
import { ToastProvider } from 'react-toast-notifications';
import { session } from '../services/useCookie'
import MiniDrawer from '../components/MiniDrawer'
import { startWS } from '@/services/useWebsocket';
import { useRouter } from 'next/router';
import usePosts from '../services/usePosts';



function MyApp({ Component, pageProps }) {
  const [loggedIn, setLoggedIn] = useState(false)
  const [id, setId] = useState('')
  const router = useRouter();
  const isMounted = useRef(false);
  const { fetchPosts } = usePosts();

  const fetchData = async () => {
    try {
      const responseData = await session();
      if (responseData.success === true) {
        setLoggedIn(true);
        setId(responseData.id)
        console.log(typeof responseData.id, responseData.id);
        startWS(responseData.id);
        router.push('/');
      } else {
        console.log('La demande de session n\'a pas réussi.')
      }
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    if (!isMounted.current) {
      fetchData();
      isMounted.current = true;
    }
  }, []);

  return (
    <ToastProvider>
      <MiniDrawer loggedIn={loggedIn} setLoggedIn={setLoggedIn} id={id} />
        <Component {...pageProps} loggedIn={loggedIn} setLoggedIn={setLoggedIn} setId={setId} id={id} />
    </ToastProvider>
  )
}

export default MyApp