import Connexion from "./profile/profile";

const Header = ()=> {


    return(
        <header>
            <div id="mainNav">
                {/*TODO Change title from simple P to an IMG*/}
                <p id="mainTitle">Site de Mint de NFT</p>
            </div>
            <div id="profileNav">
                <input id="mainSearchBtn" type="text" placeholder="Search for an NFT collection"></input>
                <Connexion />
            </div>
        </header>
    )
}

export default Header;