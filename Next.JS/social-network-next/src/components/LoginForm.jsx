import React, {useState} from 'react';

const LoginForm = ({form, setForm, formErrors, setFormErrors, onButtonClick}) => {

    const [showPassword, setShowPassword] = useState(false);

    const toggleShowPassword = () => {
        setShowPassword(prevState => !prevState);
    };
    const handlePasswordChange = (ev) => {
        const newPassword = ev.target.value;
        setForm(prevForm => ({...prevForm, password: newPassword}))

        if (formErrors.password && newPassword !== '') {
            setFormErrors(prevErrors => ({
                ...prevErrors,
                password: ''
            }));
        }

    };
    const handleDataChange = (ev) => {
        const newData = ev.target.value;
        setForm(prevForm => ({...prevForm, data: newData}))

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
                <div>Login</div>
            </div>
            <br/>
            <div className={'inputContainer'}>
                <input
                    value={form.data}
                    placeholder="Enter your email or username here"
                    onChange={handleDataChange}
                    className={'inputBox'}
                />
                <label className="errorLabel">{formErrors.data}</label>
            </div>
            <br/>
            <div className={'inputContainer'}>
                <div className="password-input">
                    <input
                        value={form.password}
                        placeholder="Enter your password here"
                        onChange={handlePasswordChange}
                        className={'inputBox'}
                        type={showPassword ? "text" : "password"}
                    />
                    <button className="password-toggle" onClick={toggleShowPassword}>
                        <span style={{fontSize: '1.5em'}}>{showPassword ? "👁️" : "🫣"}</span>
                    </button>
                </div>
                <label className="errorLabel">{formErrors.password}</label>
            </div>
            <br/>
            <div className={'inputContainer'}>
                <input className={'inputButton'} type="button" onClick={onButtonClick} value={'Log in'}/>
            </div>
            <p>Pas encore de compte? <a href="/register">Créez-en un ici</a>.</p>
        </div>

    );
};
export default LoginForm;