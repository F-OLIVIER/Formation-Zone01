const endpoint = "https://zone01normandie.org/api/graphql-engine/v1/graphql";
let headers = {
    "content-type": "application/json",
};
let options = {
    "method": "POST",
    "headers": headers,
};

function home() {
    let container = document.getElementById('container');
    while (container.firstChild) {
        container.removeChild(container.firstChild);
    }
    container.innerHTML = '';
    container.classList.add('home');

    let divError = document.createElement('div');
    divError.id = 'divError';
    divError.className = 'divError';
    divError.style.display = 'none';
    container.appendChild(divError);

    let divConnexion = document.createElement('div');
    divConnexion.className = 'divConnexion';

    let titledivConnexion = document.createElement('div');
    titledivConnexion.className = 'titledivConnexion';
    titledivConnexion.textContent = "Formulaire de connexion";
    divConnexion.appendChild(titledivConnexion);

    let formConnexion = document.createElement('form');
    formConnexion.id = 'formConnexion';
    formConnexion.className = 'formConnexion';

    let inputuser = document.createElement('input');
    inputuser.id = 'inputuser';
    inputuser.type = 'text';
    inputuser.placeholder = 'Username / email';
    divConnexion.appendChild(inputuser);

    let inputpassword = document.createElement('input');
    inputpassword.id = 'inputpassword';
    inputpassword.type = 'password';
    inputpassword.placeholder = 'Mot de passe';
    divConnexion.appendChild(inputpassword);

    let buttonConnexion = document.createElement('button');
    buttonConnexion.id = 'buttonConnexion';
    buttonConnexion.className = 'buttonConnexion';
    buttonConnexion.textContent = 'Se connecter'
    divConnexion.appendChild(buttonConnexion);

    divConnexion.appendChild(formConnexion);
    container.appendChild(divConnexion);

    buttonConnexion.addEventListener('click', connexion)
}

home();

function connexion() {
    // Requete connexion :
    // https://zone01normandie.org/api/auth/signin
    // username:password base64 encoding
    var dataEncodedStringBtoA = btoa(document.getElementById('inputuser').value + ':' + document.getElementById('inputpassword').value)
    fetch(('https://zone01normandie.org/api/auth/signin'), {
        method: 'POST',
        headers: {
            'Authorization': 'Bearer' + dataEncodedStringBtoA
        },
    })
        .then(response => {
            if (!response.ok) {
                console.log('response.status : ', response.status)
                let divError = document.getElementById('divError');
                if (response.status === 401) {
                    divError.textContent = 'Utilisateur non autorisé';
                } else {
                    divError.textContent = 'Code erreur : ' + response.status;
                }
                divError.style.display = 'block';
                throw new Error(`Erreur de réseau: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            headers.Authorization = 'Bearer ' + data;
            fetchData();
        })
        .catch(error => {
            console.error('Erreur lors de la récupération des données:', error);
        });
}


// Requete API :
// https://zone01normandie.org/api/graphql-engine/v1/graphql
async function fetchData() {
    let user = {};

    // requeté d'information sur l'utilisateur (table user)
    const queryuserInfo = {
        "query": `
        {
            user {
            lastName
            firstName
          }
        }`
    };
    options.body = JSON.stringify(queryuserInfo);
    const responseUserInfo = await fetch(endpoint, options);
    const dataUserInfo = await responseUserInfo.json();
    if (dataUserInfo.errors !== undefined) {
        console.log("Error : ", dataUserInfo.errors); // error
    }
    user.firstName = dataUserInfo.data.user[0].firstName;
    user.lastName = dataUserInfo.data.user[0].lastName;

    // requéte du level 
    const queryuserLevel = {
        "query": `
        {
            user {
            events(where: {event: {path: {_ilike: "/rouen/div-01"}}})  {
              level
            }
          }
        }`
    };
    options.body = JSON.stringify(queryuserLevel);
    const responseUserLevel = await fetch(endpoint, options);
    const dataUserLevel = await responseUserLevel.json();
    if (dataUserLevel.errors !== undefined) {
        console.log("Error : ", dataUserLevel.errors); // error
    }
    // console.log('dataUserLevel : ', dataUserLevel)
    user.lvl = dataUserLevel.data.user[0].events[0].level;

    // requeté d'xp utilisateur (table transaction)
    const queryuserXP = {
        "query": `
        {
            transaction(where: {type: {_eq: "xp"} event: {path: {_ilike: "/rouen/div-01"}}}, order_by: {id: asc}) {
                amount
                createdAt
            }
        }`
    };
    options.body = JSON.stringify(queryuserXP);
    const responseUserXP = await fetch(endpoint, options);
    const dataUserXP = await responseUserXP.json();
    if (dataUserXP.errors !== undefined) {
        console.log("Error : ", dataUserXP.errors); // error
    }

    user.listTransaction = dataUserXP.data.transaction;
    // console.log("Liste des transaction : ", listTransaction);

    let sum = 0;
    for (let i = 0; i < user.listTransaction.length; i++) {
        sum += user.listTransaction[i].amount
    }
    user.maxXP = sum;

    // requéte de la somme d'audit reçu (Received)
    const queryXPdown = {
        "query": `
        {
            transaction_aggregate(where: {type: {_eq: "down"} event: {path: {_ilike: "/rouen/div-01"}}}, order_by: {id: asc}) {
            	aggregate {sum {amount}}
            }
        }`
    };
    options.body = JSON.stringify(queryXPdown);
    const responseXPdown = await fetch(endpoint, options);
    const dataXPdown = await responseXPdown.json();
    if (dataXPdown.errors !== undefined) {
        console.log("Error : ", dataXPdown.errors); // error
    }
    user.XPdown = dataXPdown.data.transaction_aggregate.aggregate.sum.amount;

    // requéte de la somme d'audit envoyé (Done)
    const queryXPup = {
        "query": `
        {
            transaction_aggregate(where: {type: {_eq: "up"} event: {path: {_ilike: "/rouen/div-01"}}}, order_by: {id: asc}) {
            	aggregate {sum {amount}}
            }
        }`
    };
    options.body = JSON.stringify(queryXPup);
    const responseXPup = await fetch(endpoint, options);
    const dataXPup = await responseXPup.json();
    if (dataXPup.errors !== undefined) {
        console.log("Error : ", dataXPup.errors); // error
    }
    user.XPup = dataXPup.data.transaction_aggregate.aggregate.sum.amount;
    MAJpage(user);
}


function MAJpage(user) {
    // console.log("User : ", user);
    let container = document.getElementById('container');
    container.classList.remove('home');
    while (container.firstChild) {
        container.removeChild(container.firstChild);
    }
    container.innerHTML = '';

    // Information utilisateur
    let divUser = document.createElement('div');
    divUser.className = 'divUser';
    let divInfo = document.createElement('div');
    divInfo.className = 'divInfo';

    // logout
    let logout = document.createElement('div')
    logout.id = 'logout';
    logout.className = 'logout';
    logout.textContent = 'Se déconnecter';
    divInfo.appendChild(logout);

    let divName = document.createElement('div');
    divName.className = 'divName';
    let name = document.createElement('h3');
    name.textContent = user.firstName + ' ' + user.lastName;
    divName.appendChild(name);
    divInfo.appendChild(divName);

    // utilisateur + xp max
    let divxp = document.createElement('div');
    divxp.className = 'divxp';
    let lvl = document.createElement('div');
    lvl.textContent = 'Level : ' + user.lvl;
    divxp.appendChild(lvl);
    let totalXP = document.createElement('div');
    if (user.maxXP >= 1000000) {
        totalXP.textContent = 'XP total : ' + parseFloat(user.maxXP / 1000000).toFixed(2) + ' Mb';
    } else {
        totalXP.textContent = 'XP total : ' + Math.round(user.maxXP / 1000) + ' kb';
    }
    divxp.appendChild(totalXP);
    divInfo.appendChild(divxp);
    divUser.appendChild(divInfo);


    // info ratio d'audit
    let divaudit = document.createElement('div');
    divaudit.className = 'divaudit';
    let titleaudit = document.createElement('h4');
    titleaudit.textContent = "Information d'audit";
    divaudit.appendChild(titleaudit);
    let XPdone = document.createElement('div');
    XPdone.classList = 'XPdone';
    XPdone.innerHTML = '⬆️ Done &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; ' + Math.round(user.XPup / 10000) / 100 + ' MB';
    divaudit.appendChild(XPdone);

    // graph svg des ratio d'audit
    // calcul du % de la barre
    let traceXPup = 0;
    let traceXPdown = 0;
    if (user.XPup > user.XPdown) {
        traceXPdown = user.XPdown / user.XPup
        traceXPup = 1;
    } else {
        traceXPup = user.XPup / user.XPdown;
        traceXPdown = 1;
    }

    let svgAudit = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
    let lineDone = document.createElementNS('http://www.w3.org/2000/svg', 'line');
    lineDone.setAttribute('x1', 0);
    lineDone.setAttribute('y1', 5);
    lineDone.setAttribute('x2', `${traceXPup * 100}%`); //user.XPup);
    lineDone.setAttribute('y2', 5);
    lineDone.setAttribute('stroke-width', 20);
    lineDone.setAttribute('stroke', 'yellow');
    svgAudit.appendChild(lineDone);
    let lineReceived = document.createElementNS('http://www.w3.org/2000/svg', 'line');
    lineReceived.setAttribute('x1', 0);
    lineReceived.setAttribute('y1', 30);
    lineReceived.setAttribute('x2', `${traceXPdown * 100}%`); //user.XPdown);
    lineReceived.setAttribute('y2', 30);
    lineReceived.setAttribute('stroke-width', 20);
    lineReceived.setAttribute('stroke', 'white');
    svgAudit.appendChild(lineReceived);

    divaudit.appendChild(svgAudit);

    let XPreceived = document.createElement('div');
    XPreceived.className = 'XPreceived';
    XPreceived.innerHTML = '⬇️ Received &nbsp;&nbsp; ' + Math.round(user.XPdown / 10000) / 100 + ' MB';
    divaudit.appendChild(XPreceived);
    let ratio = document.createElement('div');
    console.log(user.XPup, ' / ', user.XPdown, ' = ', user.XPup / user.XPdown)
    ratio.innerHTML = 'ℹ️ Audit ratio &nbsp;' + parseFloat((user.XPup / user.XPdown)).toFixed(1);
    divaudit.appendChild(ratio);
    divUser.appendChild(divaudit);

    container.appendChild(divUser);

    // Graph d'xp de l'utilisateur
    let divgraph = document.createElement('div');
    divgraph.className = 'divgraph';
    let title = document.createElement('h1');
    title.textContent = "Graphique de l'XP";
    divgraph.appendChild(title);

    let svg = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
    svg.setAttribute('width', 80);
    svg.setAttribute('height', 50);
    svg.setAttribute('viewBox', "0 0 100 100");
    let rect = document.createElementNS('http://www.w3.org/2000/svg', 'rect');
    rect.setAttribute('width', 100);
    rect.setAttribute('height', 100);
    rect.setAttribute('fill', "black");
    svg.appendChild(rect);

    let firstDate = Date.parse(user.listTransaction[0].createdAt);
    let lastDate = Date.parse(user.listTransaction[user.listTransaction.length - 1].createdAt);
    let amplitudeDate = new Date(lastDate - firstDate);
    amplitudeDate.setMonth(amplitudeDate.getMonth() + 1); // décalage de 1 mois sur l'amplitude de la date

    let sum = 0;
    let maxgraph = (user.maxXP + 100000);
    for (let i = 0; i < user.listTransaction.length; i++) {
        sum += user.listTransaction[i].amount;
        // produit en croix pour ramener à 100 <rect width="100%" height="100%" fill="back" />
        //              user.maxXP + 1000 <==> 100%
        // user.listTransaction[i].amount <==> x%

        // amplitudeDate + 7 <==> 100%
        //  date - fisrtDate <==> x%
        let line = document.createElementNS('http://www.w3.org/2000/svg', 'line');
        if (i != 0) {
            line.setAttribute('x1', ((Date.parse(user.listTransaction[i].createdAt) - firstDate) * 100) / amplitudeDate);
            let y = (sum * 100) / maxgraph;
            line.setAttribute('y1', 100 - y);
            line.setAttribute('x2', ((Date.parse(user.listTransaction[i - 1].createdAt) - firstDate) * 100) / amplitudeDate);
            let ypreview = ((sum - user.listTransaction[i].amount) * 100) / maxgraph;
            line.setAttribute('y2', 100 - ypreview);
            line.setAttribute('stroke', 'blue');
            svg.appendChild(line);
        }
    }

    let datafirstDate = new Date(firstDate);
    let datalastDate = new Date(lastDate);
    let titlelineLegendex = document.createElementNS('http://www.w3.org/2000/svg', 'text');
    titlelineLegendex.setAttribute('x', 15);
    titlelineLegendex.setAttribute('y', 99);
    titlelineLegendex.setAttribute('fill', 'white');
    titlelineLegendex.setAttribute('font-size', 4);
    titlelineLegendex.textContent = 'Axe des dates : du ' + String(datafirstDate.getDate()).padStart(2, '0') + '/' + String(datafirstDate.getMonth()).padStart(2, '0') + '/' + datafirstDate.getFullYear() + ' au ' + String(datalastDate.getDate()).padStart(2, '0') + '/' + String(datalastDate.getMonth()).padStart(2, '0') + '/' + datalastDate.getFullYear();
    svg.appendChild(titlelineLegendex);

    let titlelineLegendey = document.createElementNS('http://www.w3.org/2000/svg', 'text');
    titlelineLegendey.setAttribute('x', 1);
    titlelineLegendey.setAttribute('y', 55);
    titlelineLegendey.setAttribute('fill', 'white');
    titlelineLegendey.setAttribute('font-size', 4);
    titlelineLegendey.setAttribute('transform', "rotate(-90 4,55)");
    titlelineLegendey.textContent = "Axe de l'XP reçu";
    svg.appendChild(titlelineLegendey);

    divgraph.appendChild(svg);
    container.appendChild(divgraph);

    logout.addEventListener('click', () => {
        home();
    });
}

// 
// https://pages.github.com
