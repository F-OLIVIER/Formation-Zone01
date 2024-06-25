import '../styles/globals.css';
import {useEffect, useRef, useState} from 'react';
import {session} from '@/services/useCookie';
import MiniDrawer from '../components/MiniDrawer';
import {startWS} from '@/services/useWebsocket';
import {useRouter} from 'next/router';
import usePosts from '../services/usePosts';
import { Toaster } from 'react-hot-toast';
import { motion, AnimatePresence } from 'framer-motion';

function MyApp({Component, pageProps}) {
    const [loggedIn, setLoggedIn] = useState(false);
    const [id, setId] = useState('');
    const router = useRouter();
    const isMounted = useRef(false);
    const {fetchPosts} = usePosts();
    const [notifications, setNotifications] = useState([]);


    useEffect(() => {
        const fetchData = async () => {
            try {
                const responseData = await session();
                if (responseData.success === true) {

                    startWS(responseData.id, setNotifications, router);
                    setLoggedIn(true);
                    setId(responseData.id);
                    //console.log(typeof responseData.id, responseData.id)
                    //router.push('/');
                } else {
                    console.log('La demande de session n\'a pas réussi.');
                }
            } catch (error) {
                console.error(error);
            }
        };
        fetchData();
    }, []);

  return (
    <AnimatePresence mode='wait'>
      <motion.div key={router.pathname}>
        <motion.div className='transitext' initial={{ top: "-100vh" }} animate={{ top: "-100vh" }} exit={{ top: "100vh" }} transition={{ duration: 2, delay: 0.5, ease: [0, 1, 1, 0] }}>
        </motion.div>
        <motion.div className='slide-in'
          initial={{ borderRadius: "0 0 50% 50% / 50% 50% 30% 30%", transform: "translateY(-150vh)" }} animate={{ borderRadius: "0 0 50% 50% / 50% 50% 30% 30%", transform: "translateY(-150vh)" }} exit={{ borderRadius: "0 0 0 0", transform: "translateY(0vh)" }} transition={{ duration: 1, ease: [0.7, 0, 0.3, 1] }}>
        </motion.div>
        <motion.div className='slide-out'
          initial={{ borderRadius: "0 0 0 0", transform: "translateY(-50vh)" }} animate={{ borderRadius: "50% 50% 0 0 / 20% 20% 50% 50%", transform: "translateY(100vh)" }} exit={{ borderRadius: "50% 50% 0 0 / 20% 20% 50% 50%", transform: "translateY(100vh)" }} transition={{ duration: 1, ease: [0.7, 0, 0.3, 1] }}>
        </motion.div>
        <MiniDrawer loggedIn={loggedIn} setLoggedIn={setLoggedIn} id={id} />
        <Component {...pageProps} loggedIn={loggedIn} setLoggedIn={setLoggedIn} setId={setId} id={id} />
        <Toaster />
      </motion.div>
    </AnimatePresence>
  )
}

export default MyApp;
