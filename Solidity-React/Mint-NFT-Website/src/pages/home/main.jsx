import ABI from '../../contract/abi.json'
import Card from './card';
import { useState } from 'react';
const { Web3 } = require('web3');
let web3 = new Web3(window.ethereum);

const NewMintedNFT = ({ content, setContent }) => {
    const w = web3.eth;

    //TODO setting up account, not necessary at this stage
    let ctx = '0xa870AB2744E3aC5FCbc64AFd976c936aB3dB9Cb6';
    let contract = new w.Contract(
        ABI,
        ctx
    );

    const fetchMinted = () => {
        fetch('https://ipfs.io/ipfs/QmYdjCkVuQ981Lw6GLgiDbe79GNtburPQXThKZS5RGkHaK/_metadata.json')
            .then(resp => resp.json())
            .then(async data => {
                let totalSupply = Number(await contract.methods.totalSupply().call())

                data.slice(0, totalSupply).forEach(token => {
                    setContent(f => [...f, token])
                })
            })
    }

    return (
        <>
            <button onClick={fetchMinted}>NFT COLLECTION #1</button>
            <div className="collection">
                {console.log("Content = ", content)}
                {content.map((element, index) =>
                    <Card
                        key={index}
                        name={element.name}
                        image={element.image.replace("ipfs://", "https://ipfs.io/ipfs/")}
                        description={element.description}
                    />
                )}
            </div>
        </>
    )


}

const MintedNFT = ({ content, setContent }) => {
    {/* <div className="listCollection">
        <NFTCollectionList />
        </div> */}
    const w = web3.eth;

    // const blockPerDay = 7150
    const blockPerDay = 40000
    let tokensList = []


    //TODO setting up account, not necessary at this stage
    let ctx = '0xa870AB2744E3aC5FCbc64AFd976c936aB3dB9Cb6';
    let contract = new w.Contract(
        ABI,
        ctx
    );

    const [all, setAll] = useState()

    const GetLatestBlockNUmber = async () => { return Number(await web3.eth.getBlock("latest").then(b => b.number)) }

    const fetchMinted = async () => {
        let latest = await GetLatestBlockNUmber()
        let actual = latest - (blockPerDay * 7)
        const metadata = await fetch('https://ipfs.io/ipfs/QmYdjCkVuQ981Lw6GLgiDbe79GNtburPQXThKZS5RGkHaK/_metadata.json')
            .then(resp => resp.json())

        while (actual <= latest) {
            const events = await contract.getPastEvents('Transfer', {
                filter: { from: '0x0000000000000000000000000000000000000000' },
                fromBlock: actual,
                toBlock: (actual + 1000 <= latest) ? actual + 1000 : latest,
            })

            if (events.length !== 0) {
                console.log("Events", events);

                events.forEach(token => {
                    let data = metadata[Number(token.returnValues.tokenId) - 1]
                    // setContent(() => [ ...content, metadata[Number(token.returnValues.tokenId)]])
                    // setContent([ ...content, metadata[Number(token.returnValues.tokenId)]])
                    tokensList.push(data)
                })

                console.log("Token List = ", tokensList.length);

                setContent([...tokensList])
            }

            actual += 1000
        }
    }

    fetchMinted()

    return (
        <>
            {/* <button onClick={fetchMinted}>NFT COLLECTION #1</button> */}
            <button onClick={setAll(false)}>NFT COLLECTION #1 MINTED</button>
            <button onClick={setAll(true)}>NFT COLLECTION #1 ALL</button>
            <div className="collection">
                {content.map((element, index) =>
                    <Card
                        key={index}
                        name={element.name}
                        image={element.image.replace("ipfs://", "https://ipfs.io/ipfs/")}
                        description={element.description}
                    />
                )}
            </div>
        </>
    )
}

const Main = () => {
    const [content, setContent] = useState([])
    const provider = window.ethereum

    if (provider === undefined) {
        alert("Please add a metamask wallet to your browser.");
    }

    return (
        <main>
            {provider
                ? <NewMintedNFT content={content} setContent={setContent} />
                : <p> You need to download metamask to be able to access this page</p>
            }

        </main>
    )
}

export default Main;
