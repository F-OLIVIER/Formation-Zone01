import React, { useState, useEffect } from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import PasswordStrengthMeter from '../services/usePasswordStrength';

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
        setForm(prevForm => ({
            ...prevForm,
            dateofbirth: date
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

        <div className={'titleContainer'}>
                <div>Register</div>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    value={form.email}
                    placeholder="Enter your email here"
                    onChange={(ev) => handleDataChange('email', ev.target.value)}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.email}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    value={form.password}
                    placeholder="Enter your password here"
                    onChange={(ev) => handleDataChange('password', ev.target.value)}
                    className={'inputBox'}
                    type="password"
                />
                <PasswordStrengthMeter password={form.password} />
                <label className="errorLabel">{formErrors.password}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    value={form.firstname}
                    placeholder="Enter your first name here"
                    onChange={(ev) => handleDataChange('firstname', ev.target.value)}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.firstname}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    value={form.lastname}
                    placeholder="Enter your last name here"
                    onChange={(ev) => handleDataChange('lastname', ev.target.value)}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.lastname}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <DatePicker
                    selected={form.dateofbirth}
                    placeholderText="Select your date of birth"
                    onChange={handleDateChange}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.dateofbirth}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    type="file"
                    accept="image/*"
                    onChange={(ev) => handleDataChange('avatar', ev.target.value[0])}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.avatar}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    value={form.nickname}
                    placeholder="Enter your nickname (optional)"
                    onChange={(ev) => handleDataChange('nickname', ev.target.value)}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.nickname}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <textarea
                    value={form.aboutme}
                    placeholder="Enter your bio (optional)"
                    onChange={(ev) => handleDataChange('aboutme', ev.target.value)}
                    className={'inputBox'}
                    maxLength={200} // Example max length of 200 characters
                    // Empeche etirement horizontal mais vertical oui
                    style={{ resize: 'vertical'}}
                />
                <label className="errorLabel">{formErrors.aboutme}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input className={'inputButton'} type="button" onClick={onRegisterClick} value={'Register'} />
            </div>
        </div>
    );
};
export default RegisterForm;
