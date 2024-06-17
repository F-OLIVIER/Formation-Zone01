import React, { useState, useEffect } from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import PasswordStrengthMeter from '../services/usePasswordStrength';
import {Button} from '@mui/material';

const RegisterForm = ({ setForm,form,formErrors, setFormErrors,onRegisterClick }) => {

    const [isFixed, setIsFixed] = useState(false);
    const [fixedHeight, setFixedHeight] = useState(0);
 


    useEffect(() => {
        const hasErrors = Object.values(formErrors).some(error => error !== '');
        setIsFixed(hasErrors);

        // Calculer la hauteur de la div fixe
        const mainContainer = document.querySelector('.mainContainer');
        if (mainContainer) {
            setFixedHeight(mainContainer.offsetHeight);
        }
    }, [formErrors]);

    const handleImageChange = (e) => {
        const file = e.target.files[0];
        if (file) {
            setForm(prevForm => ({
                ...prevForm,
                avatar: file
            }));

            if (formErrors.avatar) {
                setFormErrors(prevErrors => ({
                    ...prevErrors,
                    avatar: ''
                }));
            }
        }
    };

    const handleDataChange = (field, value) => {
        setForm(prevForm => ({
            ...prevForm,
            [field]: value
        }));

        if (formErrors[field] && value !== '') {
            setFormErrors(prevErrors => ({
                ...prevErrors,
                [field]: ''
            }));
        }
    };
    const handleDateChange = (date) => {
        const formattedDate = date.toISOString(); // Converts Date object to ISO 8601 string
        setForm(prevForm => ({
            ...prevForm,
            dateofbirth: formattedDate
        }));
    
        if (formErrors.dateofbirth && date !== '') {
            setFormErrors(prevErrors => ({
                ...prevErrors,
                dateofbirth: ''
            }));
        }
    };

    const marginTopThreshold = 50; // Seuil de la marge supérieure

    const marginTop = marginTopThreshold > 0 ? Math.min(marginTopThreshold, 50) : 0;


    return (
        <div className={`mainContainer ${isFixed ? 'fixed' : ''}`} style={{ marginTop: '50px', maxHeight: '1000px' }}>
        <form onSubmit={onRegisterClick} encType="multipart/form-data">

        <div className={'titleContainer'}>
                <div className='title'>Register</div>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    value={form.email}
                    placeholder="Email"
                    onChange={(ev) => handleDataChange('email', ev.target.value)}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.email}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <PasswordStrengthMeter password={form.password} />
                <input
                    value={form.password}
                    placeholder="Password"
                    onChange={(ev) => handleDataChange('password', ev.target.value)}
                    className={'inputBox'}
                    type="password"
                />
                <label className="errorLabel">{formErrors.password}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    value={form.firstname}
                    placeholder="First name"
                    onChange={(ev) => handleDataChange('firstname', ev.target.value)}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.firstname}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    value={form.lastname}
                    placeholder="Last name"
                    onChange={(ev) => handleDataChange('lastname', ev.target.value)}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.lastname}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <DatePicker
                    selected={form.dateofbirth}
                    placeholderText="Birthdate"
                    onChange={handleDateChange}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.dateofbirth}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input type="file" onChange={handleImageChange} />
                <label className="errorLabel">{formErrors.avatar}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    value={form.nickname}
                    placeholder="Username (optional)"
                    onChange={(ev) => handleDataChange('nickname', ev.target.value)}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.nickname}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <textarea
                    value={form.aboutme}
                    placeholder="Tell us about yourself (optional)"
                    onChange={(ev) => handleDataChange('aboutme', ev.target.value)}
                    className={'inputBox'}
                    maxLength={200} // Example max length of 200 characters
                    // Empeche etirement horizontal mais vertical oui
                    style={{ resize: 'vertical'}}
                />
                <label className="errorLabel">{formErrors.aboutme}</label>
            </div>
            <br />
        <Button type="submit" variant="contained">
          Register
        </Button>
      </form>
        </div>
    );
};
export default RegisterForm;
