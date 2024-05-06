import { timer, demarrage, nextMap, remainingTime, TotalTimer } from "../main.js";
import { bomb, animed_1_bombStopped, animed_2_bombStopped } from "./bomb.js";
import { calculScore, handleLives, score } from "./player.js";
import { SendInWebSocket } from "./webSocket.js";

export let gameStopped = false;
export let levelCounter = Number(localStorage.getItem('levelCounter')) || 0;
const modalElement = document.getElementById('modal');

// option 0 = pause, 1 = game over, 2 = win
export function modalGame(option = 0) {
    gameStopped = true;

    // affichage de la fenetre modal
    modalElement.style.opacity = '100%';

    const gameMessage = document.getElementById('typeMenu');
    if (option == 3) { // page de demarrage
        // TODO: Choix du level de depart ici
        getScores();
        levelCounter = 0;
        localStorage.setItem('levelCounter', levelCounter);
        calculScore(0, 0, 0, -1);
        modalElement.style.background = 'url("../../sprites/all_level/bg_start.jpg")';
        modalElement.style.backgroundSize = 'cover';
        gameMessage.innerHTML = "Notre histoire commence dans la petite ville de Boomville ou notre héros que vous incarnez mène une vie normale.<br>Vous travaillez dans une boutique de feux d'artifice, aimant les couleurs et les explosions.<br><br>Un jour, en rentrant du travail, vous trouvez un chat parlant nommé Sparky qui vous prend pour l'héritier du Roi ultime Galacto, gardien de la galaxie et du monde réel.<br>Sparky vous demande d'affronter la garde royale pour être fin prêt pour votre nouvelle vie.";
        scoreButton();
        StartButton();
    } else if (option == 1) { // Game over
        gameMessage.innerHTML = "GAME OVER<br><br>Ce qui devait arriver arriva... La force obscure gagne en puissance et commence à envahir le monde réel. Le monde est perdu... Les étoiles dans le ciel de Boomville s'assombrissent, reflétant votre échec. Sparky vous encourage à ne pas abandonner et vous rappelle que chaque échec est une leçon pour devenir plus fort. Espérons qu'un miracle se produise, car même en s'attendant au pire personne n'y est jamais prêt...<br><br>Votre score : " + score;

        if (lastScore < score) {
            tableScore()
        } else {
            restart();
        }

    } else if (option == 2) { // level win
        let textehistory = [
            "Ce premier combat ne vous laisse pas indifférent... Vous découvrez ce que vous allez réellement devoir affronter. Bien qu'il n'y ait pas eu de grosse difficulté lors de ce premier combat, ce n'est pas le moment de se relacher... La bataille ne fais que commencer, en effet avec Sparky vous découvrez un secret inattendu : la porte de sortie a disparu ! Où est-elle ?",
            "Après avoir réussi à vaincre les membres de la garde. Cette victoire est un premier pas dans votre quête pour sauver la galaxie et le monde réel. Le fidèle serviteur de Galacto junior (petit frère du roi ultime Galacto) Bubbleman est ici, il est certain qu’il va vous barrer la route et rendre votre quête difficile.",
            "Ca y est, après avoir affronté les différents groupes d'élite de la garde Royale, vous pouvez enfin vous reposer et démarrer votre nouveau role d'héritier !<br><br>Soudain, une atmosphère pesante se fit sentir.<br>Que se passe-t-il ? Le sol tremble, Sparky crie ! Vous vous retrouvez dans l'obsurité totale.<br><br>Vous vous re-saissisez et vous rendez compte que vous êtes dans la salle du trône et réalisez que ce n'était pas fini, mais au contraire que le vrai combat allait commencer... Contre Galacto junior le frère du Roi.",
            "FELICITATIONS, aidé par Sparky, vous avez affronté la force obscure dans une bataille épique. Vous avez réussit à vaincre tous les ennemis et gagné votre place au trone.<br>Sans prévenir, le Roi ultime Galacto apparaît devant vous vous félicitant pour votre courage et votre détermination. Il vous adoube et vous confère officiellement le titre d'Héritier de la Galaxie.<br>La petite ville de Boomville célèbre son héros avec un grand spectacle de feux d'artifice.<br><br>Votre score : " + score
        ];
        gameMessage.innerHTML = textehistory[levelCounter - 1];

        if (levelCounter < 4) {
            nextLevel();
        } else if (levelCounter == 4) {
            if (lastScore < score) {
                tableScore()
            } else {
                let tableScoreTable = document.getElementById("tableScoreTable");
                if (!tableScoreTable.classList.contains('display')) {
                    tableScoreTable.classList.add('display');
                    document.getElementById('scoringPage').classList.add('display');
                }
                restart();
            }
        }

        if (levelCounter >= 4) {
            levelCounter = 0;
        } else {
            levelCounter++;
        }
        localStorage.setItem('levelCounter', levelCounter);
    } else { // jeux en pause (touche echap)
        gameMessage.innerHTML = 'PAUSE';
        resume();
        restart();
    }
}

// *******************************************************************
// ******************* Fonction pour chaque bouton *******************
// *******************************************************************
let timerThrottlebutton = 0;

function StartButton() {
    var startButton = document.getElementById("startButton");
    startButton.classList.remove('display');

    // Fonction gestionnaire d'événements
    function StartButtonClick() {
        const now = new Date();
        if (now - timerThrottlebutton > 500) {
            timerThrottlebutton = now;
            levelCounter++;
            localStorage.setItem('levelCounter', levelCounter);
            modalElement.style.background = 'rgba(0, 0, 0, 0.8)';
            modalElement.style.opacity = '0%';
            startButton.classList.add('display');
            gameStopped = false;

            document.getElementById("scoreTable").classList.add('display');
            let tableScoreTable = document.getElementById("tableScoreTable");
            if (!tableScoreTable.classList.contains('display')) {
                tableScoreTable.classList.add('display');
                document.getElementById('scoringPage').classList.add('display');
            }
            demarrage();

            // Supprime les gestionnaires d'événements devenu inutile
            startButton.removeEventListener('click', StartButtonClick);
            document.getElementById("scoreTable").removeEventListener('click', scoreButtonClick);
            document.getElementById('page1').removeEventListener("click", previewPageButtonClick);
            document.getElementById('page2').removeEventListener("click", nextPageButtonClick);
        }
    }

    startButton.addEventListener('click', StartButtonClick);
}


function restart() {
    var boutonRestart = document.getElementById("Restart");
    boutonRestart.classList.remove('display');

    boutonRestart.addEventListener("click", function () {
        const now = new Date();
        if (now - timerThrottlebutton > 500) {
            timerThrottlebutton = now;
            location.reload();
        }
    });
}

function scoreButton() {
    let scoreTable = document.getElementById("scoreTable");
    scoreTable.classList.remove('display');
    scoreTable.addEventListener('click', scoreButtonClick);
}

// Fonction gestionnaire d'événements
function scoreButtonClick() {
    const now = new Date();
    if (now - timerThrottlebutton > 500) {
        timerThrottlebutton = now;
        document.querySelector('#tableScoreTable').classList.toggle('display');
        document.querySelector('#scoringPage').classList.toggle('display');
    }
}

function nextLevel() {
    var boutonNextLevel = document.getElementById('NextLevel');
    boutonNextLevel.classList.remove('display');

    // Fonction gestionnaire d'événements
    function boutonNextLevelClick() {
        const now = new Date();
        if (now - timerThrottlebutton > 500) {
            timerThrottlebutton = now;

            modalElement.style.opacity = '0%';
            handleLives(0); // re-initialisationd des vies
            gameStopped = false;
            nextMap();

            boutonNextLevel.classList.add('display');
            boutonNextLevel.removeEventListener("click", boutonNextLevelClick);
        }
    }

    boutonNextLevel.addEventListener("click", boutonNextLevelClick);
}

function resume() {
    var boutonResume = document.getElementById("Resume");
    boutonResume.classList.remove('display');
    var boutonRestart = document.getElementById("Restart");
    boutonRestart.classList.remove('display');

    // Fonction gestionnaire d'événements
    function resumeNextLevelClick() {
        const now = new Date();
        if (now - timerThrottlebutton > 500) {
            timerThrottlebutton = now;

            // masquer le menu
            modalElement.style.opacity = '0%';
            boutonResume.classList.add('display');
            boutonRestart.classList.add('display');

            // relancer le jeux
            gameStopped = false;

            // relancer le timer
            timer();

            // relancer l'animation de la bomb
            if (animed_1_bombStopped || animed_2_bombStopped) {
                bomb();
            }

            boutonResume.removeEventListener("click", resumeNextLevelClick);
        }
    }

    boutonResume.addEventListener("click", resumeNextLevelClick);
}

// Fonction gestionnaire d'événements
function nextPageButtonClick() {
    const now = new Date();
    if (now - timerThrottlebutton > 500) {
        timerThrottlebutton = now;
        document.getElementById('page2').classList.add('display');
        document.getElementById('page1').classList.remove('display');

        let tableScoreTable = document.getElementById("tableScoreTable");
        let rows = tableScoreTable.getElementsByTagName("tr");

        for (let i = 0; i < rows.length; i++) {
            let element = rows[i];
            if (!element.classList.contains('entete')) {
                if (element.classList.contains('page2')) {
                    element.classList.remove('display');
                } else {
                    element.classList.add('display');
                }
            }
        }
    }
}

// Fonction gestionnaire d'événements
function previewPageButtonClick() {
    const now = new Date();
    if (now - timerThrottlebutton > 500) {
        timerThrottlebutton = now;
        document.getElementById('page1').classList.add('display');
        document.getElementById('page2').classList.remove('display');

        let tableScoreTable = document.getElementById("tableScoreTable");
        let rows = tableScoreTable.getElementsByTagName("tr");
        for (let i = 0; i < rows.length; i++) {
            let element = rows[i];
            if (!element.classList.contains('entete')) {
                if (element.classList.contains('page2')) {
                    element.classList.add('display');
                } else {
                    element.classList.remove('display');
                }
            }
        }
    }
}

// ----------------------------------------------
// ------------- Gestion du scoring -------------
// ----------------------------------------------

let lastScore;

function getScores() {
    fetch('../data/scores.json')
        .then(resp => resp.json())
        .then(data => {
            let table = document.getElementById("tableScoreTable");
            let page = false;
            let tbody = document.createElement("tbody");

            for (let score of data) {
                let row = document.createElement("tr");
                for (let property in score) {
                    let cell = document.createElement("td");
                    cell.textContent = score[property];
                    row.appendChild(cell);

                    // gestion du numero de page
                    if (score.playerRank <= 5) {
                        page = true;
                    }
                }

                if (page) {
                    row.classList.add("page1");
                } else {
                    row.classList.add("page2");
                    row.classList.add("display");
                }
                page = false;

                tbody.appendChild(row);
            }
            table.appendChild(tbody);
            lastScore = data[data.length - 1].playerScore;
            document.getElementById('page1').addEventListener("click", previewPageButtonClick);
            document.getElementById('page2').addEventListener("click", nextPageButtonClick);
        });
}

function tableScore() {
    document.getElementById("scoreTable").classList.add('display');
    // affichage de la div de saisir pseudo
    let divNameScore = document.getElementById("nameScore");
    divNameScore.classList.remove('display');
    const scoreMsg = document.getElementById('scoreMsg');
    scoreMsg.innerHTML = 'NOUVEAU RECORD !<br>Votre score : ' + score;

    // gestion du bouton de l'événement du score
    var buttonScore = document.getElementById("buttonScore");
    buttonScore.addEventListener("click", function () {
        const now = new Date();
        if (now - timerThrottlebutton > 0.5 * 1000) { // 0.5 sec
            timerThrottlebutton = now;
            divNameScore.classList.add('display');

            // --------------------------------------------------------------------------------------------------------
            // ----------- envoie de la valeur du joueur Via Web-Socket au serveur go pour traitement json ------------
            // --------------------------------------------------------------------------------------------------------
            const minutes = Math.floor((TotalTimer + 180 - remainingTime) / 60);
            const secondes = (TotalTimer + 180 - remainingTime) % 60;
            const formatChrono = `${minutes}:${secondes < 10 ? '0' : ''}${secondes}`;

            // Création de la variable 'dataToSendServerGO' qui sera envoyé au serveur go
            var dataToSendServerGO = {
                playerName: document.getElementById('inputScore').value,
                playerLevel: levelCounter,
                playerScore: score,
                playerTimer: formatChrono
            };
            SendInWebSocket(dataToSendServerGO);

            // Mise a jour et affichage du tableau des scores
            let tableScoreTable = document.getElementById("tableScoreTable");
            tableScoreTable.classList.remove('display');
            document.getElementById('scoringPage').classList.remove('display');
            // suppression de tout les éléments du tableau
            let rows = tableScoreTable.getElementsByTagName("tr");
            for (let i = rows.length - 1; i > 0; i--) {
                tableScoreTable.deleteRow(1);
            }

            // mise à jour du tableau
            setTimeout(() => {
                let previous = document.getElementById('page1');
                if (!previous.classList.contains('display')) {
                    previous.classList.add('display');
                }
                let next = document.getElementById('page2');
                if (next.classList.contains('display')) {
                    next.classList.remove('display');

                }
                getScores();
            }, 40);

            restart();
        }
    });
}
