import React, { useState } from 'react';
import { useRouter } from 'next/router'; 
//import { useToasts } from 'react-toast-notifications';
import toast from 'react-hot-toast';
import { loginUser } from '../services/useLogin';
import LoginForm from '../components/LoginForm';
import { startWS } from '@/services/useWebsocket';

const Login = (props) => {
  const [form, setForm] = useState({
    data: '',
    password: '',
  });
  const [formErrors, setFormErrors] = useState({});
  //const { addToast } = useToasts();
  const router = useRouter();
  const onButtonClick = async () => {
    let valid = true;
    if ('' === form.data) {
      setFormErrors(prevErrors => ({
        ...prevErrors,
        data: 'Please enter your email or username',
      }));
      valid = false;
    }else if ((form.password).length < 8) {
      setFormErrors(prevErrors => ({
        ...prevErrors,
        password: 'The password must be 8 characters or longer'
      }));
      valid = false;
    }
    /* if (!form.email.includes('@') || !form.email.includes('.')) {
      setFormErrors ( prevErrors => ({
        ...prevErrors,
        email: "Please enter a valid email address",
      }));
      valid = false;
    } */
    if ('' === form.password) {
      setFormErrors(prevErrors => ({
        ...prevErrors,
        password: "Please enter a password",
      }));
      valid = false;
    }

    if (valid === true) {
      try {
        const responseData = await loginUser(form);
        if (responseData.success === true) {
          toast.success('Authentication successful!', {
            duration: 4000,
            position: 'top-center',
            style: {backgroundColor: 'rgba(0,255,34,0.5)', color: 'white'},
            icon: '👏',
          });
          props.setLoggedIn(true);
          props.setId(responseData.id)
          //console.log(typeof responseData.id, responseData.id);
          startWS(responseData.id);
          router.push('/');
        } else {
          toast.error('Authentication failed. Please check your credentials. Error: ' + responseData.message, {
            duration: 4000,
            position: 'top-center',
            style: {backgroundColor: 'rgba(255,0,0,0.5)', color: 'white'},
          });
        }
      } catch (error) {
        console.error(error);
        toast.error('Error during authentication: ' + error.message, {
          duration: 4000,
          position: 'top-center',
          style: {backgroundColor: 'rgba(255,0,0,0.5)', color: 'white'},
        });
      }
    }
  };
  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
      <LoginForm
        form={form}
        setForm={setForm}
        formErrors={formErrors}
        setFormErrors={setFormErrors}
        onButtonClick={onButtonClick}
      />
    </div>
  );
};
export default Login;

