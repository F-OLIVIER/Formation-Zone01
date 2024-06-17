import React from "react";

const OneGroup = ({ SingleForm }) => {
    if (!SingleForm) {
        return <div>No data available</div>;
    }


    return (
        <>
            <h1> Title : {SingleForm.Title}</h1>
            <hr/>
            <p> about US : {SingleForm.AboutGroup}</p>
            </>
    );
}

export default OneGroup;