import { mapLevel1, mapLevel2, mapLevel3, mapBoss } from '../data/bomberman.data.js';
import { randomize } from './map.js';
import { deadPlayer, calculScore } from './player.js';
import { levelCounter, gameStopped } from './modal.js';
import { player } from './map.js';

let enemies_list = document.getElementsByClassName('enemy');
let enemySpeed = 2000;
let lifeMiniBoss = 3;
let lifeBigBoss = 5;

export function initEnemies() {
    let case_id = 0;
    let enemy_id_counter = 0;
    let grille = mapLevel1;
    switch (levelCounter) {
        case 1:
            grille = mapLevel1;
            break;
        case 2:
            grille = mapLevel2;
            enemySpeed = 1500;
            break;
        case 3:
            grille = mapLevel3;
            enemySpeed = 1000;
            break;
        case 4:
            grille = mapBoss;
            enemySpeed = 1500;
            break;
    }

    for (let x = 0; x < grille.length; x++) {
        for (let y = 0; y < grille[x].length; y++) {
            if (grille[x][y] === 't' && levelCounter < 4) { // ennemie basique (levels 1 à 3)
                const elem = document.getElementById(case_id);
                if (randomize() == 13 && !elem.classList.contains('wall') && !elem.classList.contains('block')) {
                    const container = document.getElementsByClassName("map")[0]
                    const enemy = document.createElement('div');
                    enemy.id = `enemy_${enemy_id_counter}`;
                    enemy.classList.add('enemy');
                    enemy.style.top = x * 50 + 16 + 'px';
                    enemy.style.left = y * 50 + 16 + 'px';
                    enemy.style.backgroundImage = "url('../../sprites/level_" + levelCounter + "/enemy.png')";
                    enemy.setAttribute('data-pos', case_id + ',' + 0 + ',' + 0);
                    container.appendChild(enemy);
                    enemy_id_counter++;
                }
            } else if (grille[x][y] === 's') { // spectateur map boss
                const container = document.getElementsByClassName("map")[0]
                const enemy = document.createElement('div');
                enemy.classList.add('viewer');
                enemy.style.top = x * 50 + 16 + 'px';
                enemy.style.left = y * 50 + 16 + 'px';
                container.appendChild(enemy);
            } else if (grille[x][y] === 'h' || grille[x][y] === 'H') { // miniboss level 3 et boss level 4
                const container = document.getElementsByClassName("map")[0]
                const enemy = document.createElement('div');
                enemy.id = `boss`;
                enemy.classList.add('enemy');
                if (grille[x][y] === 'h') {
                    enemy.classList.add('miniboss');
                } else {
                    enemy.classList.add('bigboss');
                }

                enemy.style.top = x * 50 + 16 + 'px';
                enemy.style.left = y * 50 + 16 + 'px';
                if (grille[x][y] === 'h') {
                    enemy.style.backgroundImage = "url('../../sprites/level_3/boss.png')";
                } else if (grille[x][y] === 'H') {
                    enemy.style.backgroundImage = "url('../../sprites/level_4/boss_1.png')";

                }
                enemy.setAttribute('data-pos', case_id + ',' + 0 + ',' + 0);
                container.appendChild(enemy);
                enemy_id_counter++;
            }
            case_id++;
        }
    }
}

let lastFrameTime = null;

export async function enemyHandler(timestamp) {
    if (!gameStopped) {
        if (!lastFrameTime) {
            lastFrameTime = timestamp;
        }

        const delta = timestamp - lastFrameTime;

        if (delta > enemySpeed) { // petit mob = 2000 // miniboss = 1000 // boss = 800
            moveEnemies();
            lastFrameTime = timestamp;
        }
        enemyCollision();

    }

    requestAnimationFrame(enemyHandler);
}

let timerThrottleDead = 0;
export async function enemyCollision() {
    const playerPosition = player.getBoundingClientRect();
    for (let enemy of enemies_list) {
        const enemyPosition = enemy.getBoundingClientRect();
        if (playerPosition.left + 20 < enemyPosition.left + enemyPosition.width &&
            playerPosition.left + playerPosition.width - 20 > enemyPosition.left &&
            playerPosition.top + 40 < enemyPosition.top + enemyPosition.height &&
            playerPosition.top + playerPosition.height - 20 > enemyPosition.top) {
            const now = new Date();
            if (now - timerThrottleDead > 1000) { // Evite les appels en boucle de deadplayer()
                timerThrottleDead = now;
                deadPlayer();
            }
            break;
        }
    }
}

export async function enemyDeath(enemy_id) {
    const enemy = document.getElementById(enemy_id);
    if (enemy.classList.contains('miniboss')) {
        console.log("lifeMiniBoss : ", lifeMiniBoss);
        if (lifeMiniBoss == 0) { // mort du miniboss
            enemy.style.backgroundImage = "url('../../sprites/level_3/miniboss_death.png')";
            enemy.classList.add('deadboss');
            calculScore(0, 0, 0, 400);
            setTimeout(function () {
                enemy.remove();
                enemies_list = document.getElementsByClassName('enemy');
            }, 720);
        } else { // damage du miniboss
            enemy.style.backgroundImage = "url('../../sprites/level_3/boss_damage.png')";
            enemy.classList.add('damageboss');
            setTimeout(function () {
                enemy.style.backgroundImage = "url('../../sprites/level_3/boss.png')";
                enemy.classList.remove('damageboss');
            }, 1000);
        }
        lifeMiniBoss--
    } else if (enemy.classList.contains('bigboss')) {
        if (lifeBigBoss == 0) { // mort du boss
            enemy.style.backgroundImage = "url('../../sprites/level_4/boss_death.png')";
            enemy.classList.add('deadboss');
            calculScore(0, 0, 0, 1000);
            setTimeout(function () {
                enemies_list = document.getElementsByClassName('enemy');
                enemy.remove();
            }, 720);
        } else if (lifeBigBoss === 3) { // damage du boss (apparence 1) et passage en mode méchant (apparence 2)
            enemy.style.backgroundImage = "url('../../sprites/level_4/boss_1_damage.png')";
            enemy.classList.add('damageboss');
            setTimeout(function () {
                enemy.style.backgroundImage = "url('../../sprites/level_4/boss_2.png')";
                enemy.classList.remove('damageboss');
            }, 1000);
            enemySpeed = 800;
        } else if (lifeBigBoss > 3) { // damage big boss (apparence 1)
            enemy.style.backgroundImage = "url('../../sprites/level_4/boss_1_damage.png')";
            enemy.classList.add('damageboss');
            setTimeout(function () {
                enemy.style.backgroundImage = "url('../../sprites/level_4/boss_1.png')";
                enemy.classList.remove('damageboss');
            }, 1000);
        } else if (lifeBigBoss < 3) { // damage big boss (apparence 2)
            enemy.style.backgroundImage = "url('../../sprites/level_4/boss_2_damage.png')";
            enemy.classList.add('damageboss');
            setTimeout(function () {
                enemy.style.backgroundImage = "url('../../sprites/level_4/boss_2.png')";
                enemy.classList.remove('damageboss');
            }, 1000);
        }
        lifeBigBoss--
    } else {
        enemy.style.backgroundImage = "url('../../sprites/all_level/enemy_death.png')";
        enemy.classList.add('dead');
        setTimeout(function () {
            enemy.remove();
            enemies_list = document.getElementsByClassName('enemy');
        }, 200);
    }
}

async function moveEnemies() {
    for (let enemy of enemies_list) {
        const random_direction = Math.floor(Math.random() * 4);
        let valeur_de_depladement = 16.5;
        let patchX = 0;
        let patchY = 0;
        let patchDiv = 0;

        switch (random_direction) {
            case 0: // up
                patchY = -valeur_de_depladement;
                patchDiv = -17;
                break;
            case 1: // right
                patchX = valeur_de_depladement;
                patchDiv = 1;
                break;
            case 2: // down
                patchY = valeur_de_depladement;
                patchDiv = 17;
                break;
            case 3: // left
                patchX = -valeur_de_depladement;
                patchDiv = -1;
                break;
        }

        let Attributes = enemy.getAttribute('data-pos').split(',');
        let div_ennemy = Number(Attributes[0]);
        if (!checkWallCollision(div_ennemy, patchDiv)) {
            enemy.setAttribute('data-pos', (div_ennemy + patchDiv) + ',' + (Number(Attributes[1]) + patchX) + ',' + (Number(Attributes[2]) + patchY));
            enemy.style.transform = `scale(3) translate3d(${Number(Attributes[1]) + patchX}px, ${Number(Attributes[2]) + patchY}px, 0)`;
        } else if (enemy.classList.contains('miniboss') || enemy.classList.contains('bigboss')) { // si c'est un boss on re-itére l'opération pour avoir un mouvement constant
            let direction = [1, 17, -1, -17];
            patchX = 0;
            patchY = 0;
            for (let index = 0; index < 5; index++) {
                if (!checkWallCollision(div_ennemy, direction[index])) {
                    switch (direction[index]) {
                        case -17: // up
                            patchY = -valeur_de_depladement;
                            patchDiv = -17;
                            break;
                        case 1: // right
                            patchX = valeur_de_depladement;
                            patchDiv = 1;
                            break;
                        case 17: // down
                            patchY = valeur_de_depladement;
                            patchDiv = 17;
                            break;
                        case -1: // left
                            patchX = -valeur_de_depladement;
                            patchDiv = -1;
                            break;
                    }

                    enemy.setAttribute('data-pos', (div_ennemy + patchDiv) + ',' + (Number(Attributes[1]) + patchX) + ',' + (Number(Attributes[2]) + patchY));
                    enemy.style.transform = `scale(3) translate3d(${Number(Attributes[1]) + patchX}px, ${Number(Attributes[2]) + patchY}px, 0)`;
                    break;
                }
            }
        }
    }
}

function checkWallCollision(enemy_pos_id, direction) {
    const wall = document.getElementById(enemy_pos_id + direction);
    return wall.classList.contains('wall') || wall.classList.contains('block');
}
