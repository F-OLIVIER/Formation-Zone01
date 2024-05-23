import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/router';
import GroupForm from "../components/GroupForm";


import {createGroup, getGroup, InviteInMyGroup} from '../services/useCreateGroup';

const Group = (props) => {
    const [data, setData] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const router = useRouter();
    const [form, setForm] = useState({
        title: '',
        aboutGroup: '',
    });

    const [formInvite, setInvite] = useState({
        nameOfGroup: '',
        nameOfThePerson: '',
    });

    const [formErrors, setFormErrors] = useState({});
    const [inviteErrors, setInviteErrors] = useState({});


    //pour fetcher tout les groupes dont l'utilisateur est chef
    useEffect(() => {
        const fetchData = async () => {
            try {
                const result = await getGroup();
                if (result.success) {
                    setData(result.data);
                } else {
                    console.error('Failed to get group data:', result.message);
                }
            } catch (error) {
                console.error('Error during fetching of group data:', error);
            }
            setLoading(false);
        };

        fetchData();
    }, []);



    const onRegisterClick = async () => {
        let valid = true;

        if ('' === form.title) {
            setFormErrors(prevErrors => ({
                ...prevErrors,
                title: 'Please enter a title for the group',
            }));
            valid = false;
        }

        if ('' === form.aboutGroup) {
            setFormErrors(prevErrors => ({
                ...prevErrors,
                aboutGroup: "Please enter a desciption of the group",
            }));
            valid = false;
        }

        if ((form.aboutGroup).length > 200) {
            setFormErrors(prevErrors => ({
                aboutGroup : 'Bio should not exceed 200 characters.',
            }));
            valid = false;
        }

        if (valid === true) {
            if (Object.values(form).some(value => value === '')) {
                return;
            }

            try {
                console.log(form)
                const responseData = await createGroup(form);
                if (responseData.success) {

                } else {
                    console.error('Create group failed:', responseData.message);
                  /*  addToast(responseData.message, {
                        appearance: 'error',
                    });*/
                }
            } catch (error) {
                console.error('Error during creation of group:', error);

            }
        }
    }


    const onInviteClick = async () => {
        console.log("test validation invitation")
        console.log(formInvite)

            const reponseData = await InviteInMyGroup(formInvite);
            if (reponseData.success) {

            } else {
                console.error('Invitation failed:', reponseData.message);

            }


    }



        if (loading) return 'Loading...';
    if (error) return 'An error occurred';

    return (
        <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
            <GroupForm
                form={form}
                formErrors={formErrors}
                onRegisterClick={onRegisterClick}
                setForm={setForm}
                data={data}
                setInvite={setInvite}
                setInviteErrors={setInviteErrors}
                formInvite={formInvite}
                onInviteClick={onInviteClick}
            />
        </div>
    );
};

export default Group;