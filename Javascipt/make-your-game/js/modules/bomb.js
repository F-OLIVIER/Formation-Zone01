import { handleLives, calculScore } from "./player.js";
import { id_div_sub_player, destructionWall } from "./wall.js";
import { enemyDeath } from "./enemy.js";
import { gameStopped } from "./modal.js";

// *******************************************************************
// ***************** Fonction de gestion de la bomb ******************
// *******************************************************************
let timerThrottleBomb = 0;
let tmp_id_case_bomb = 0;
const MILLE = 1000;
export let animed_1_bombStopped = false; // mise en pause du jeu
export let animed_2_bombStopped = false; // mise en pause du jeu


export function bomb() {
    const now = new Date();
    let id_case_bomb = id_div_sub_player;
    if (now - timerThrottleBomb > MILLE || animed_1_bombStopped || animed_2_bombStopped) {
        let nbWallDestroy = 0;
        let nbEnnemieDestroy = 0;
        let bomb;
        if (animed_1_bombStopped || animed_2_bombStopped) {
            // récupération des informations necessaire suite a la mise en pause
            bomb = document.getElementById('bomb');
            id_case_bomb = tmp_id_case_bomb;
        } else {
            timerThrottleBomb = now;

            // div sous le bonhomme
            const subDiv = document.getElementById(id_case_bomb);

            // création de la div de la bomb
            bomb = document.createElement('div');
            bomb.classList.add('bomb');
            bomb.style.top = subDiv.offsetTop + 16 + 'px';
            bomb.style.left = subDiv.offsetLeft + 18 + 'px';

            const container = document.getElementsByClassName("map")[0];
            container.appendChild(bomb);
        }

        const enemies_list = document.getElementsByClassName('enemy');
        // Disparition de la div bomb
        setTimeout(function () {
            if (!gameStopped) {

                // gestion de l'explosion centrale de la bomb avec la div existante
                bomb.classList.remove('bomb');
                bomb.id = 'explosion';
                bomb.classList.add('expCentral');
                setTimeout(function () {
                    bomb.remove();
                }, 0.3 * MILLE);

                for (let enemy of enemies_list) {
                    if (enemy.getAttribute('data-pos').split(',')[0] == id_case_bomb) {
                        enemyDeath(enemy.id);
                        nbEnnemieDestroy++;
                    };
                }

                const divList = [-1, -2, 1, 2, -17, -34, 17, 34]; // liste des id des cases à checker
                const explosionList = [[64, 48], [0, 48], [64, 48], [0, 32], [64, 32], [64, 0], [64, 32], [64, 16]]; // liste des positions X et Y de l'explosion dans la sprite sheet

                for (let i = 0; i < divList.length; i++) {
                    let div_explosion = document.getElementById(id_case_bomb + divList[i]);
                    if (!div_explosion.classList.contains('wall') && !div_explosion.classList.contains('block')) {
                        switch (divList[i]) {
                            case -1: // case 1 a gauche
                                explosionBomb(id_case_bomb + divList[i], explosionList[i][0], explosionList[i][1], 'expLeft1');
                                break;
                            case -2: // case 2 à gauche
                                explosionBomb(id_case_bomb + divList[i], explosionList[i][0], explosionList[i][1], 'expLeft2');
                                break;
                            case 1:
                                explosionBomb(id_case_bomb + divList[i], explosionList[i][0], explosionList[i][1], 'expRight1');
                                break;
                            case 2:
                                explosionBomb(id_case_bomb + divList[i], explosionList[i][0], explosionList[i][1], 'expRight2');
                                break;
                            case -17:
                                explosionBomb(id_case_bomb + divList[i], explosionList[i][0], explosionList[i][1], 'expUp1');
                                break;
                            case -34:
                                explosionBomb(id_case_bomb + divList[i], explosionList[i][0], explosionList[i][1], 'expUp2');
                                break;
                            case 17:
                                explosionBomb(id_case_bomb + divList[i], explosionList[i][0], explosionList[i][1], 'expDown1');
                                break;
                            case 34:
                                explosionBomb(id_case_bomb + divList[i], explosionList[i][0], explosionList[i][1], 'expDown2');
                                break;
                        }

                        if (id_div_sub_player == id_case_bomb + divList[i]) { // check du player
                            handleLives()
                        }

                        for (let enemy of enemies_list) { // check des ennemies
                            if (enemy.getAttribute('data-pos').split(',')[0] == id_case_bomb + divList[i]) {
                                enemyDeath(enemy.id);
                                nbEnnemieDestroy++;
                            };
                        }

                    } else if (div_explosion.classList.contains('wall')) { // gestion suppression du mur à 2 cases
                        if (div_explosion.classList.contains('door')) {
                            div_explosion.classList.remove('wall');
                            div_explosion.style.backgroundImage = "url('../../sprites/all_level/door.png')";
                        } else {
                            destructionWall(div_explosion); // destruction du mur 1 cassable apres l'explosion à 2 cases
                        }
                        nbWallDestroy++;
                        if (i % 2 == 0) { // Si un mur cassable est détecté et qu'il n'est qu'à une case (indice pair : 0, 2, 4 etc), on skip la case suivante
                            i++;
                        }
                    } else {
                        if (i % 2 == 0) { // Si un mur non cassable est détecté et qu'il n'est qu'à une case (indice pair : 0, 2, 4 etc), on skip la case suivante
                            i++;
                        }
                    }
                };

                calculScore(nbWallDestroy, nbEnnemieDestroy, 0);

                // check si le player et sur la case de la bombe, si c'est le cas, mort du player
                if (id_div_sub_player == id_case_bomb) {
                    handleLives()
                }
            } else {
                animed_2_bombStopped = true;
                tmp_id_case_bomb = id_case_bomb;
            }
        }, MILLE);
    }
    animed_1_bombStopped = false;
    animed_2_bombStopped = false;
}

// *******************************************************************
// **************** Fonction d'animation d'explosion *****************
// *******************************************************************

// génére l'animation de l'explosion sur la div mis en entrée
function explosionBomb(id_case_bomb, x_image, y_image, directionExplosion) {
    const container = document.getElementsByClassName("map")[0];
    const subDiv = document.getElementById(id_case_bomb);
    // création de la div d'explosion
    let div_explosion = document.createElement('div');
    div_explosion.id = 'explosion';
    div_explosion.style.top = subDiv.offsetTop + 16 + 'px';
    div_explosion.style.left = subDiv.offsetLeft + 18 + 'px';
    div_explosion.style.backgroundImage = "url('../../sprites/all_level/bomb.png')";
    // ajout de la bomb à l'HTML
    div_explosion.style.backgroundPosition = `-${x_image}px -${y_image}px`;
    div_explosion.classList.add('explosion');
    container.appendChild(div_explosion);
    div_explosion.classList.add(directionExplosion);

    setTimeout(function () {
        div_explosion.remove(); // Disparition de la div d'explosion
    }, 0.3 * MILLE);
}



