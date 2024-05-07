import React from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import PasswordStrengthMeter from '../services/usePasswordStrength';

const RegisterForm = ({ setForm,form,formErrors, onRegisterClick }) => {
    return (
        <div className={'mainContainer'} style={{ marginTop: '50px' }}>
            <div className={'titleContainer'}>
                <div>Register</div>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    value={form.email}
                    placeholder="Enter your email here"
                    onChange={(ev) => setForm(prevForm => ({...prevForm,email :ev.target.value }))}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.email}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    value={form.password}
                    placeholder="Enter your password here"
                    onChange={(ev) => setForm(prevForm => ({...prevForm,password :ev.target.value}))}
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
                    onChange={(ev) => setForm(prevForm => ({...prevForm,firstname :ev.target.value}))}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.firstname}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    value={form.lastname}
                    placeholder="Enter your last name here"
                    onChange={(ev) => setForm(prevForm => ({...prevForm,lastname :ev.target.value}))}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.lastname}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <DatePicker
                    selected={form.dateofbirth}
                    placeholderText="Select your date of birth"
                    onChange={(date) => setForm(prevForm => ({...prevForm, dateofbirth: date}))}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.dateofbirth}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    type="file"
                    accept="image/*"
                    onChange={(ev) => setForm(prevForm => ({...prevForm,avatar :ev.target.files[0]}))}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.avatar}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    value={form.nickname}
                    placeholder="Enter your nickname (optional)"
                    onChange={(ev) => setForm(prevForm => ({...prevForm,nickname :ev.target.value}))}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.nickname}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <textarea
                    value={form.aboutme}
                    placeholder="Enter your bio (optional)"
                    onChange={(ev) => setForm(prevForm => ({...prevForm,aboutme :ev.target.value}))}
                    className={'inputBox'}
                    maxLength={200} // Example max length of 200 characters
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
