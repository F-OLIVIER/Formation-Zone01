import React from 'react';
import ReactDOM from 'react-dom/client';
import reportWebVitals from './reportWebVitals';
// import Main from './script/main';
import './static/app.css';
import Header from './pages/header';
import Main from './pages/home/main';
import './static/index.css';
import Mint from './script/mint/mint';
import Footer from './pages/footer';
const root = ReactDOM.createRoot(document.getElementById('container'));
//REMIND: NFT PUBLISHED ON pinata
root.render(
  <>
    <Header />
    <Mint />
    <Main />
    <Footer /> 
  </>
)
//If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
