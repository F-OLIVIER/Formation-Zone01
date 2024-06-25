import React from "react";

const OneGroup = ({ SingleForm }) => {
    if (!SingleForm) {
        return <div>No data available</div>;
    }


    return (
        <>
            <h1 className="pagetitle">{SingleForm.Title}</h1>
            <div className="grouptitle">

            <p className="o5 white">{SingleForm.AboutGroup}</p>
            <hr/>
            </div>
            </>
    );
}

export default OneGroup;