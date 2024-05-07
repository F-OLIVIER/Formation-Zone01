import React, { useState } from 'react';
import { useRouter } from 'next/router';
import { useToasts } from 'react-toast-notifications';
import RegisterForm from '../components/RegisterForm';
import { RegisterUser } from '../services/useRegister';

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

    let valid = true;
    const { addToast } = useToasts();
    const onRegisterClick = async () => {
        if ('' === form.email) {
            setFormErrors(prevErrors => ({
                ...prevErrors,
                email: 'Please enter your email',
            }));
            valid = false;
        }
        if (!form.email.includes('@') || !form.email.includes('.')) {
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
        }
        if ('' === form.firstname) {
            setFormErrors(prevErrors => ({
                firstname : 'Please enter your first name'
            }));
            valid = false;
        }
        if ('' === form.lastname) {
            setFormErrors(prevErrors => ({
                lastname : 'Please enter your last name',
            }));
            valid = false;
        }
        if ('' === form.dateofbirth) {
            setFormErrors(prevErrors => ({
                dateofbirth : 'Please enter your date of birth',
            }));
            valid = false;
        }
        if ((form.password).length < 7) {
            setFormErrors(prevErrors =>({
                password : 'The password must be 8 characters or longer',
            }));
            valid = false;
        }
        if (/[!@#$%^&*(),.?":{}|<>]/.test(form.nickname)) {
            setFormErrors( prevErrors => ({
                nickname : 'Nickname should not contain special characters.',
            }));
            valid = false;
        }
        if ((form.aboutme).length > 200) {
            setFormErrors(prevErrors => ({
                aboutme : 'Bio should not exceed 200 characters.',
            }));
            valid = false;
        }
        if (valid === true) {
        try {
            const responseData = await RegisterUser(form);
            if (responseData.success === true) {
                addToast('Registration successful!', {
                    appearance: 'success',
                    autoDismiss: true,
                });
                router.push('/login');
            } else {
                addToast('Registration failed. Please check your credentials. Error: ' + responseData.message, {
                    appearance: 'error',
                    autoDismiss: true,
                });
            }
        } catch (error) {
            console.error(error);
            addToast('Error during registration: ' + error.message, {
                appearance: 'error',
                autoDismiss: true,
            });
        }
    };
    }
    return (
        <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>

        <RegisterForm
            form={form}
            formErrors={formErrors}
            onRegisterClick={onRegisterClick}
            setForm={setForm}
        />
        </div>
    );
};
export default Register;
