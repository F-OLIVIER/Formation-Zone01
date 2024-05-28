import { home } from "./home.js";
import { commonBlock } from "./useful.js";

export const cookieName = "user_token";
userNotConnected();

// Fonction gestionnaire d'événements pour le boutton Home
// let timerThrottlebutton = 0;
// function HomeButtonClick() {
//     const now = new Date();
//     if (now - timerThrottlebutton > 500) {
//         timerThrottlebutton = now;
//         // console.log("Page home demandé");
//         home();
//     }
// }

// Bouton Couillére Burgond
document.getElementById("bandeauGauche").addEventListener('click', home);

export function userNotConnected() {
    // Affichage des boutons utilisateur non connecté
    let UserData = { Logged: false, Msgerr: "" };
    commonBlock({ UserData });
    document.getElementById('Container').innerHTML = '';

    if (document.cookie.split(";").some((item) => item.trim().startsWith(cookieName + "="))) { // check la precence du cookie
        // si le cookie existe redirection automatique vers "/home"
        home();
    } else {
        console.log("absence de cookie");
        let UserData = { Logged: false, Msgerr: "Vous devez être connecté pour pouvoir intéragir avec le forum" };
        commonBlock({ UserData });
    }
}

function badRequest() {
    if (windows.location.pathname != "") {
        windows.location.href = "/";
    }
}


// ----------------------------------------------------
// ----------- méthode de lecture du cookie -----------
// ----------------------------------------------------

function getCookieValue(cookieName) {
    // Liste de tous les cookies
    console.log("document.cookie : ", document.cookie);

    // Test 1
    if (document.cookie.split(";").some((item) => item.trim().startsWith(cookieName + "="))) {
        console.log("Test 1 : TROUVEEEEEEEEEEEEEE !!!!!!!!!!!!");
    }

    // Test 2
    const cookieValue = document.cookie
        .split("; ")
        .find((row) => row.startsWith(cookieName + '='))
        ?.split("=")[1];
    console.log("Test 2 :", cookieValue);

    // Test 3
    console.log("Test 3 :", lireCookie(cookieName));

    // Test 4
    const cookies = document.cookie.split(';');
    for (let i = 0; i < cookies.length; i++) {
        const cookie = cookies[i].trim();
        if (cookie.startsWith(cookieName + '=')) {
            return cookie.substring(cookieName.length + 1);
        }
    }

    return null;
}

// Fonction pour lire un cookie en JavaScript
function lireCookie(cookieName) {
    var nom = cookieName + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var cookies = decodedCookie.split(';');
    for (var i = 0; i < cookies.length; i++) {
        var cookie = cookies[i].trim();
        if (cookie.indexOf(nom) == 0) {
            return cookie.substring(nom.length);
        }
    }
    return "";
}