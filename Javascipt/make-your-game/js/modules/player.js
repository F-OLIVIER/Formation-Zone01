import { caseSubPlayer } from "./wall.js";
import { bomb } from './bomb.js';
import { gameStopped, modalGame } from './modal.js';
import { player } from './map.js';

// --------------------------------------------------
// ------------ Création des animations -------------
// --------------------------------------------------

// Fonction qui gère le sens de l'animation.
// up    : movement_direction = 0
// Right : movement_direction = 1
// down  : movement_direction = 2
// left  : movement_direction = 3

// gestion des vies du player
let lives = 3; // nombre de vies
export function handleLives(option) {
    if (option == 0) { // re-initialisation entre chaque lvl
        lives = 3;
        document.getElementById('heart1').style.opacity = '100%';
        document.getElementById('heart2').style.opacity = '100%';
        document.getElementById('heart3').style.opacity = '100%';
        currentLeft = 0;
        currentTop = 0;
    } else if (lives > 1) { // gestion de la perte d'une vie si option vide
        document.getElementById('heart' + lives).style.opacity = '0%';
        lifePlayer();
        lives--;
        calculScore(0, 0, lives);
    } else { // plus de vie, mort du player
        deadPlayer();
    }
}

// animation perte d'une vie du player
export function lifePlayer() {
    player.classList.add('damage');
    setTimeout(function () {
        player.classList.remove('damage');
    }, 2 * 1000);

}

export function deadPlayer() {
    player.classList.add('dead');
    setTimeout(function () {
        player.remove();
        modalGame(1);
    }, 0.3 * 1000);
}

// --------------------------------------------------
// ------------ Gestion des événements --------------
// --------------------------------------------------
let eventType;
let move = false;

document.addEventListener('keydown', (event) => {
    move = true;
    eventType = event;
    if (!gameStopped) {
        switch (event.key) {
            case ' ': // pose d'une bombe
                bomb();
                break;
            case 'Escape': // jeux en pause
                modalGame(0);
                break;
            case 'F5':
                location.reload();
                break;
        }
    }
});

document.addEventListener('keyup', (event) => {
    if (!gameStopped) {
        player.classList.remove('animate');
    }
    if (event != undefined && event.key === eventType.key) {
        move = false;
    }
});

// --------------------------------------------------
// ------------ Gestion du déplacement --------------
// --------------------------------------------------
let currentTop = 0; // = player.offsetTop;
let currentLeft = 0; // = player.offsetLeft;
let hasCollided = false;
let lastInput;

export function updateCharacterState() {
    const playerSpeed = 1.7; // Vitesse du joueur
    const now = new Date();

    if (move && !gameStopped) {
        player.classList.add('animate');
        switch (eventType.key) {
            case 'ArrowUp': // 0
                document.documentElement.style.setProperty('--player-direction', '-0px');
                if (!caseSubPlayer(30, 30) || !hasCollided && lastInput !== eventType.key) {
                    currentTop -= playerSpeed;
                    lastInput = eventType.key;
                    hasCollided = true;
                } else {
                    hasCollided = false;
                }
                break;
            case 'ArrowRight': // 1
                document.documentElement.style.setProperty('--player-direction', '-32px');
                if (!caseSubPlayer(50, 50) || !hasCollided && lastInput !== eventType.key) {
                    currentLeft += playerSpeed;
                    lastInput = eventType.key;
                    hasCollided = true;
                } else {
                    hasCollided = false;
                }
                break;
            case 'ArrowDown': // 2
                document.documentElement.style.setProperty('--player-direction', '-64px');
                if (!caseSubPlayer(30, 60) || !hasCollided && lastInput !== eventType.key) {
                    currentTop += playerSpeed;
                    lastInput = eventType.key;
                    hasCollided = true;
                } else {
                    hasCollided = false;
                }
                break;
            case 'ArrowLeft': // 3
                document.documentElement.style.setProperty('--player-direction', '-96px');
                if (!caseSubPlayer(10, 50) || !hasCollided && lastInput !== eventType.key) {
                    currentLeft -= playerSpeed;
                    lastInput = eventType.key;
                    hasCollided = true;
                } else {
                    hasCollided = false;
                }
                break;
        }

        // Mettre à jour la position du joueur

        player.style.transform = `scale(1.9) translate(${currentLeft}px, ${currentTop}px)`;
    }
}

// --------------------------------------------------
// --------------- Gestion du score -----------------
// --------------------------------------------------
export let score = 0;
document.getElementById('numberscore').textContent = score;

export function calculScore(nbWallDestroy, nbennemyDestroy, life, bonusPoint) {
    // point pris sur les murs
    switch (nbWallDestroy) {
        case 1:
            score += 5;
            break;
        case 2:
            score += 12;
            break;
        case 3:
            score += 20;
            break;
        case 4:
            score += 50;
            break;
    }

    // point pris sur les ennemies
    switch (nbennemyDestroy) {
        case 1:
            score += 50;
            break;
        case 2:
            score += 120;
            break;
        case 3:
            score += 200;
            break;
        case 4:
            score += 500;
            break;
    }

    // point perdu avec les vies
    if (life === 2) {
        score -= 250;
    } else if (life === 1) {
        score -= 600;
    }

    // miniboss = 400 points
    // bigboss = 1000 points
    if (bonusPoint > 0) { 
        score += bonusPoint
    } else if (bonusPoint === -1) {
        score = 0;
    }

    // mise à jour du score sur la page
    document.getElementById('numberscore').textContent = score;
}