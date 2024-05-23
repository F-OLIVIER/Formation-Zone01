import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/router';
import RegisterForm from '../components/RegisterForm';
import { RegisterUser } from '../services/useRegister';
import toast from 'react-hot-toast';

const Register = (props) => {
    const [form, setForm] = useState({
        email: '',
        password: '',
        firstname: '',
        lastname: '',
        dateofbirth: '',
        avatar: '',
        nickname: '',
        aboutme: '',
        privateprofile: '',
    });
    const router = useRouter();
    const [formErrors, setFormErrors] = useState({});
    const onRegisterClick = async () => {
        let valid = true;

        if ('' === form.email) {
            setFormErrors(prevErrors => ({
                ...prevErrors,
                email: 'Please enter your email',
            }));
            valid = false;
        }else if (!form.email.includes('@') || !form.email.includes('.')) {
            setFormErrors ( prevErrors => ({
                ...prevErrors,
                email: "Please enter a valid email address",
            }));
            valid = false;
        }
        if ('' === form.password) {
            setFormErrors(prevErrors => ({
                ...prevErrors,
                password: "Please enter a password",
            }));
            valid = false;
        }else if ((form.password).length < 7) {
            setFormErrors(prevErrors =>({
                ...prevErrors,
                password : 'The password must be 8 characters or longer',
            }));
            valid = false;
        }
        if ('' === form.firstname) {
            setFormErrors(prevErrors => ({
                ...prevErrors,
                firstname : 'Please enter your first name',
            }));
            valid = false;
        }
        if ('' === form.lastname) {
            setFormErrors(prevErrors => ({
                ...prevErrors,
                lastname : 'Please enter your last name',
            }));
            valid = false;
        }
        if ('' === form.dateofbirth) {
            setFormErrors(prevErrors => ({
                ...prevErrors,
                dateofbirth : 'Please enter your date of birth',
            }));
            valid = false;
        }
        if (/[!@#$%^&*(),.?":{}|<>]/.test(form.nickname)) {
            setFormErrors( prevErrors => ({
                ...prevErrors,
                nickname : 'Nickname should not contain special characters.',
            }));
            valid = false;
        }
        if ((form.aboutme).length > 200) {
            setFormErrors(prevErrors => ({
                ...prevErrors,
                aboutme : 'Bio should not exceed 200 characters.',
            }));
            valid = false;
        }
        if (valid === true) {
            try {
                const responseData = await RegisterUser(form);
                if (responseData.success === true) {
                    toast.success('Registration successful!', {
                        duration: 4000,
                        position: 'top-center',
                        icon: '👏',
                        style: {backgroundColor: 'rgba(0,255,34,0.5)', color: 'white'},
                    });
                    router.push('/login');
                } else {
                    toast.error('Registration failed. Please check your credentials. Error: ' + responseData.message, {
                        duration: 4000,
                        position: 'top-center',
                        style: {backgroundColor: 'rgba(255,0,0,0.5)', color: 'white'},
                    });
                }
            } catch (error) {
                console.error(error);
                toast.error('Error during registration: ' + error.message, {
                    duration: 4000,
                    position: 'top-center',
                    style: {backgroundColor: 'rgba(255,0,0,0.5)', color: 'white'},
                });
            }
        };
    }
    return (
        <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>

        <RegisterForm
            form={form}
            setForm={setForm}
            formErrors={formErrors}
            setFormErrors={setFormErrors}
            onRegisterClick={onRegisterClick}
        />
        </div>
    );
};
export default Register;
