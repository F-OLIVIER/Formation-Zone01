import React, { useState } from 'react';

const LoginForm = ({ form, setForm, formErrors, setFormErrors, onButtonClick }) => {

    const [showPassword, setShowPassword] = useState(false);

    const toggleShowPassword = () => {
        setShowPassword(prevState => !prevState);
    };
    const handlePasswordChange = (ev) => {
        const newPassword = ev.target.value;
        setForm(prevForm => ({ ...prevForm, password: newPassword }))

        if (formErrors.password && newPassword !== '') {
            setFormErrors(prevErrors => ({
                ...prevErrors,
                password: ''
            }));
        }

    };
    const handleDataChange = (ev) => {
        const newData = ev.target.value;
        setForm(prevForm => ({ ...prevForm, data: newData }))

        if (formErrors.data && newData !== '') {
            setFormErrors(prevErrors => ({
                ...prevErrors,
                data: ''
            }));
        }
    };
    return (
        <div className={'mainContainer'}>
            <div className={'titleContainer'}>
                <p className='title'>Login</p>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                placeholder='Username or Email'
                    value={form.data}
                    onChange={handleDataChange}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.data}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <div className="password-input">
                    <input
                        value={form.password}
                        onChange={handlePasswordChange}
                        className={'inputBox'}
                        type={showPassword ? "text" : "password"}
                        placeholder='Password'
                    />
                    <button className="password-toggle" onClick={toggleShowPassword}>
                        <span style={{ fontSize: '1.5em' }}>{showPassword ? "👁️" : "🫣"}</span>
                    </button>
                </div>
                <label className="errorLabel">{formErrors.password}</label>
            </div>
            <br />
            <div className={'inputContainer confirmbtn'}>
                <input className={'inputButton'} type="button" onClick={onButtonClick} value={'Log in'} />
            </div>
            <p className='minitext'>Don't have an account? <a className='clicable-link' href="/register">Create one.</a>.</p>
        </div>

    );
};
export default LoginForm;