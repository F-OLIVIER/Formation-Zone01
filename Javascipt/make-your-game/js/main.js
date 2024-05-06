import { updateCharacterState } from './modules/player.js';
import { initGrille } from './modules/map.js';
import { enemyHandler } from './modules/enemy.js';
import { modalGame, gameStopped } from './modules/modal.js';

// *******************************************************************
// ********************* Chronomètre de la partie ********************
// *******************************************************************

export let remainingTime = 180; // temps sur le chronomètre
export let TotalTimer = 0;      // temps de jeu du joueur tout level compris

let lastTime = null;

export function timer(timestamp) {
    if (!lastTime) {
        lastTime = timestamp;
    }

    const delta = timestamp - lastTime;

    if (delta >= 1000) { // 1000 ms = 1 s
        const minutes = Math.floor(remainingTime / 60);
        const secondes = remainingTime % 60;
        const formatChrono = `${minutes}:${secondes < 10 ? '0' : ''}${secondes}`;

        document.getElementById('timer').textContent = formatChrono;

        if (remainingTime > 0 && !gameStopped) {
            remainingTime--;
        } else if (!gameStopped) {
            // Le temps est écoulé, fin du jeu
            modalGame(1);
        }

        lastTime = timestamp;
    }

    if (!gameStopped) {
        requestAnimationFrame(timer);
    }
}

requestAnimationFrame(timer);

// *******************************************************************
// **************** Fonction de gestion globale main *****************
// *******************************************************************
let frameCount = 0;
let lastime = 0;
let lastUpdateTime = performance.now();
let fpsElement = document.getElementById('numberfps');
const FPS = (60 * 2);
const FRAME_TIME = 1000 / FPS;
// let lastFrameTime = null;

export function gameLoop(timestamp) {
    let now = performance.now();
    let delta = now - lastUpdateTime;
    if (delta >= 1000) {
        fpsElement.textContent = `${frameCount}`;
        frameCount = 0;
        lastUpdateTime = now;
    }
    const deltaTime = timestamp - lastime;
    if (deltaTime >= FRAME_TIME) {
        lastime = timestamp;
        updateCharacterState();
        frameCount++;
    }
    
    requestAnimationFrame(gameLoop);
}

modalGame(3);

export function demarrage() {
    initGrille();   // initialisation de la grille
    timer();        // Commencer le chronomètre
    enemyHandler(); // initialise les ennemies
    gameLoop();     // lance la gameloop du jeux
}

export function nextMap() {
    document.getElementById('map').innerHTML = '';  // suppression de l'ancienne map
    initGrille();                                   // initialisation de la nouvelle map
    TotalTimer += (180 - remainingTime);            // temps de jeu du joueur
    remainingTime = 180;                            // re-initialise le chronomètre
    timer();                                        // Commencer le chronomètre
    enemyHandler();                                 // initialise les ennemies
}

