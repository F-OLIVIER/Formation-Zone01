import React, { useState } from 'react';import DatePicker from 'react-datepicker';
import { useRouter } from 'next/router';
import 'react-datepicker/dist/react-datepicker.css';


const GroupForm = ({ setForm,form,formErrors, onRegisterClick,data,setInvite,setInviteErrors,formInvite,onInviteClick,groups,groupsWhereIamNotIn,handleJoinRequest }) => {
    const [inviteError, setInviteError] = useState(null); // New state variable for storing the error message
    const router = useRouter();

    const handleGroupClick = (group) => {
        console.log("Group clicked: ", group);
        router.push('/detail_group?id=' + group.IdGroup);
    };

    const handleInviteClick = async () => {
        const result = await onInviteClick();
        if (result.success) {
            resetInviteForm();
        }
    };
    const resetInviteForm = () => {
        setInvite({
            nameOfGroup: '',
            nameOfThePerson: '',
        });
    };

    const handleCreateClick = async () => {
        const result = await onRegisterClick();
        if (result.success) {
            resetForm();
        }
    };

    const resetForm = () => {
        setForm({
            title: '',
            aboutGroup: '',
        });
    };

    return (
        <div className={'mainContainer'}>

            <div className={'allform'}>

            <div className={'creategroupContainer'}>

                <div className={'titleContainer'}>
                    <div>Create Group</div>
                </div>
                <br/>
                <div className={'inputContainer'}>
                    <input
                        value={form.title}
                        placeholder="Title of the Group"
                        onChange={(ev) => setForm(prevForm => ({...prevForm, title: ev.target.value}))}
                        className={'inputBox'}
                    />
                    <label className="errorLabel">{formErrors.title}</label>
                </div>

                <br/>


                <br/>
                <div className={'inputContainer'}>
                <textarea
                    value={form.aboutGroup}
                    placeholder="Enter a description of the group"
                    onChange={(ev) => setForm(prevForm => ({...prevForm, aboutGroup: ev.target.value}))}
                    className={'inputBox'}
                />
                    <label className="errorLabel">{formErrors.aboutGroup}</label>
                </div>


                <br/>

                <div className={'inputContainer'}>
                    <input className={'inputButton'} type="button" onClick={handleCreateClick} value={'Create it'}/>
                </div>
            </div>


            <br/>

            <div className={'inviteinmygroup'}>

                <div className={'titleContainer'}>
                    <div>Invite In My Group</div>
                </div>

                <br/>
                <div className={'d'}>
                    <input
                        value={formInvite.nameOfGroup}
                        placeholder="Name of the Group"
                        onChange={(ev) => setInvite(prevForm => ({...prevForm, nameOfGroup: ev.target.value}))}
                        className={'inputBox'}
                    />
                    <label className="errorLabel">{formErrors.nameOfGroup}</label>

                </div>

                <br/>


                <br/>
                <div className={'inputContainer'}>
                    <input
                        value={formInvite.nameOfThePerson}
                        placeholder="Name of the person you want to invite"
                        onChange={(ev) => setInvite(prevForm => ({...prevForm, nameOfThePerson: ev.target.value}))}
                    className={'inputBox'}
                />
                    <label className="errorLabel">{formErrors.nameOfThePerson}</label>
                </div>


                <br/>

                <div className={'inputContainer'}>
                    <input className={'inputButton'} type="button" onClick={handleInviteClick} value={'Invite'}/>
                </div>
            </div>
            </div>


            <div className={'group-list'}>
                <h1>Own Group</h1>
                <hr/>

                {groups && groups.map((group, index) => (
                    <div key={index} id={`group-${index}`}>
                        <button onClick={() => handleGroupClick(group)}>
                            {group.Title && <h2>{group.Title}</h2>}
                            {group.AboutGroup && <p>{group.AboutGroup}</p>}
                        </button>
                    </div>
                    ))}

                <h1>Groups Where I am Not In</h1>
                <hr/>

                {groupsWhereIamNotIn && groupsWhereIamNotIn.map((group, index) => (
                    <div key={index} id={`groupsWhereIamNotIn-${index}`}>
                        {group.Title && <h2>{group.Title}</h2>}
                        {group.AboutGroup && <p>{group.AboutGroup}</p>}
                        <div className="divaskforjoin">
                            <button  className="asktojoinButton" onClick={() => handleJoinRequest (group,index)}>Demander à rejoindre le groupe</button>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default GroupForm;