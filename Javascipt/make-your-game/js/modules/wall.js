import { modalGame, levelCounter } from "./modal.js";
import { player } from "./map.js";

export let id_div_sub_player = 18;

export function caseSubPlayer(correction_left, correction_top) {
    const playerPosition = player.getBoundingClientRect();
    const tabDivs = document.querySelectorAll('div.floor');
    const enemies_list = document.getElementsByClassName('enemy');
    for (let i = 1; i < tabDivs.length; i++) {
        const obstaclePosition = tabDivs[i].getBoundingClientRect();

        if (playerPosition.left + correction_left < obstaclePosition.left + obstaclePosition.width &&
            playerPosition.left + playerPosition.width > obstaclePosition.left &&
            playerPosition.top + correction_top < obstaclePosition.top + obstaclePosition.height &&
            playerPosition.top + playerPosition.height > obstaclePosition.top) {
            if (tabDivs[i].classList.contains('wall') || tabDivs[i].classList.contains('block')) {
                return true
            }
            if (tabDivs[i].classList.contains('door') && enemies_list.length == 0) {
                modalGame(2);
            }
            if (id_div_sub_player != tabDivs[i].id) {
                id_div_sub_player = Number(tabDivs[i].id);
            }
            return false;
        }
    }

}

export function destructionWall(div_explosion) {
    div_explosion.classList.add("destroyed");
    setTimeout(function () {
        div_explosion.classList.remove("destroyed");
        div_explosion.classList.remove("wall");
        div_explosion.style.backgroundImage = "url('../../sprites/level_" + levelCounter + "/ground.png')";
    }, 250);
}
