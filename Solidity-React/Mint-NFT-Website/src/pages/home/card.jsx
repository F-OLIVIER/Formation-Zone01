const Card = (props)=>{
    if (document.getElementById(props.name)){
        document.querySelectorAll('.card').forEach(e=>e.remove())
    }
    return(
        <div className={`card ${props.name}`} onClick={props.details}>
            <div className="card-header">
                <img src={props.image} className="card-img" alt="nft"/>
                <p id={props.name}>
                    <strong>{props.name}</strong>
                </p>
            </div>
            <hr />
            <div className="card-main">
                <p id="cardDescription">
                    <i>{props.description}</i>
                </p>               
            </div>
        </div>
    )
}


export default Card;