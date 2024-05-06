import abi from '../../contract/abi.json'
import { useState, useEffect } from 'react';
import Web3 from 'web3';
import GridLoader from "react-spinners/GridLoader";

const contractAddress = "0xa870AB2744E3aC5FCbc64AFd976c936aB3dB9Cb6";
const w3 = new Web3(window.ethereum);
const W3 = w3.eth;

function Mint() {
    const [error, setError] = useState('');
    const [data, setData] = useState({})
    const [waiting, setwaiting] = useState(false)
    const [transactionValid, setTransactionValid] = useState(false)
    const [hash, sethash] = useState('');

    useEffect(() => {
        fetchData();
    }, [])

    const contract = new W3.Contract(abi, contractAddress);

    async function fetchData() {
        if (typeof window.ethereum !== 'undefined') {
            // const provider = new Web3.providers.Web3Provider(window.ethereum);
            try {
                const cost = await contract.methods.cost().call();
                const totalSupply = await contract.methods.totalSupply().call();
                const object = { "cost": String(cost), "totalSupply": String(totalSupply) }
                setData(object);
            }
            catch (err) {
                setError(err.message);
                setwaiting(false);
            }
        }
    }

    async function mint() {
        setwaiting(false);
        setTransactionValid(false);
        sethash('');
        
        if (typeof window.ethereum !== 'undefined') {
            let accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
            try {
                let dataToSend = {
                    from: accounts[0],
                    value: data.cost
                }
                setwaiting(true);
                const transaction = await contract.methods.mint(accounts[0], 1).send(dataToSend);
                console.log('transction : ', transaction)

                const transactionHash = transaction.transactionHash;
                sethash(transactionHash);

                setwaiting(false);
                setTransactionValid(true);
                fetchData();
            }
            catch (err) {
                setError(err.message);
                setwaiting(false);
            }
        }
    }

    return (
        <div className="App">
            <div className="container">
                <h1>Mint one NFT !</h1>
                {error && <p>{error}</p>}
                {waiting && <p><GridLoader color="#36d7b7" /> Transaction in process</p>}
                {transactionValid && <p>Succes transaction !!! <a href={`https://mumbai.polygonscan.com/tx/${hash}`} target='_blank'>Transaction details</a></p>}

                <p className="count">{data.totalSupply} / 100</p>
                <p className="cost">NFT costs {data.cost / 10 ** 18} eth (excluding gas fees)</p>
                <button onClick={mint}>BUY one NFT</button>
            </div>
        </div>
    );
}

export default Mint;
