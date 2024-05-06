import { mapLevel1, mapLevel2, mapLevel3, mapBoss } from '../data/bomberman.data.js';
import { initEnemies } from './enemy.js';
import { levelCounter } from './modal.js';
export let player;


export function initGrille() {
    const container = document.getElementsByClassName("map")[0]
    let doorAdded = false;
    let box_Count = 0;
    let box_Depart = 18;
    let box_Trone = 0;
    let grille = mapLevel1;
    switch (levelCounter) {
        case 1:
            grille = mapLevel1;
            break;
        case 2:
            grille = mapLevel2;
            break;
        case 3:
            grille = mapLevel3;
            break;
        case 4:
            grille = mapBoss;
            break;
    }
    // boucle qui permet de créer les div pour initialiser le terrain
    for (let x = 0; x < grille.length; x++) {
        const line = document.createElement('div');
        line.className = "line";
        for (let y = 0; y < grille[x].length; y++) {
            const elem = document.createElement('div');
            elem.classList.add('floor');
            if (grille[x][y] === 'B') {
                elem.classList.add('block');
                elem.style.backgroundImage = "url('../../sprites/level_" + levelCounter + "/wall.png')";
            } else {
                let objectType = randomize();
                if (levelCounter < 4) {
                    objectType = randomize();
                } else {
                    objectType = 0;
                }
                // création des murs cassables
                if (objectType > 13 && grille[x][y] === 't') {
                    elem.classList.add('wall');
                    elem.style.backgroundImage = "url('../../sprites/level_" + levelCounter + "/wall_animed.png')";
                } else {
                    elem.style.backgroundImage = "url('../../sprites/level_" + levelCounter + "/ground.png')";
                }
                if (!doorAdded && randomize() === 15 && grille[x][y] === 't' && levelCounter < 3) { // placement aléatoire de la porte pour les levels 1 et 2
                    elem.classList.add('door');
                    if (levelCounter == 1) {
                        if (elem.classList.contains('wall')) {
                            elem.classList.remove('wall');
                        }
                        elem.style.backgroundImage = "url('../../sprites/all_level/door.png')";
                    } else {
                        if (!elem.classList.contains('wall')) {
                            elem.classList.add('wall');
                        }
                        elem.style.backgroundImage = "url('../../sprites/level_" + levelCounter + "/wall_animed.png')";
                    }
                    doorAdded = true; // Set the flag to true so that the door is not added multiple times
                } else if (levelCounter > 2 && grille[x][y] === 'O') { // placement fixe de la porte pour les levels 3 et 4
                    elem.classList.add('door');
                    elem.style.backgroundImage = "url('../../sprites/all_level/door.png')";
                }
            }

            if (grille[x][y] === 'T') { // position du trone
                box_Trone = box_Count;
                elem.classList.add('block');
            }

            if (grille[x][y] === 'P') { // position de depart du player
                box_Depart = box_Count
            }
            elem.id = box_Count;
            box_Count++;
            elem.classList.add('column');
            line.appendChild(elem);
        }
        container.appendChild(line);
    }

    // récupération des informations de depart du bonhomme
    const depart = document.getElementById(box_Depart);
    const x_top = depart.offsetTop; // Coordonnée Y par rapport à la fenêtre
    const y_left = depart.offsetLeft; // Coordonnée x par rapport à la fenêtre

    // création du player
    const caseplayer = document.createElement('div');
    caseplayer.className = 'player';
    caseplayer.style.top = x_top + (50 - 32) / 2 + 1 + 'px'; // (taille case - taille personnage) / 2 + bordure de la case
    caseplayer.style.left = y_left + (50 - 32) / 2 + 1 + 'px';
    container.appendChild(caseplayer);
    player = document.getElementsByClassName('player')[0];

    // positionnement du trone
    if (box_Trone != 0) {
        const depart_trone = document.getElementById(box_Trone);
        const trone = document.createElement('div');
        trone.className = 'trone';
        trone.style.top = depart_trone.offsetTop + (50 - 230) / 2 + 1 + 'px'; // (taille case - taille personnage) / 2 + bordure de la case
        trone.style.left = depart_trone.offsetLeft + (50 - 124) / 2 + 1 + 'px';
        container.appendChild(trone);
    }

    initEnemies();
}

// Fonction qui permet de générer des murs cassables de façon aléatoire
export function randomize() {
    const random = Math.floor(Math.random() * 17);
    return random;
}