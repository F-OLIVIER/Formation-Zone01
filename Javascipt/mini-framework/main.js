import { toDolist, newTask, isInArray } from "./toDoList.js";
import { addEventListenerById, getElementValueById, removeChildsElement, replaceTextContentFromElement } from "./script/virtualDOM.js";
import { startFuncOnthisUrl } from "./script/utils.js";

let toDoTaskList = []; // All tasks

// Récupère les données de localStorage
let ancienValueCountTasks = 0;
let buttonActive = "";
let task = {
    check: false,
    name: "",
};

export function main() {
    // Mise à jour des task avec les eventuels données sont stockées
    const storedToDoTaskList = localStorage.getItem('toDoTaskList');
    if (storedToDoTaskList) {
        // Convertit les données JSON en objet JavaScript
        toDoTaskList = JSON.parse(storedToDoTaskList);
    }
    // ré-initialisation variables
    ancienValueCountTasks = 0

    // On remplace le classique addEventListener("DOMContentLoaded") par un timer
    setTimeout(() => {
        toDolist();
        updateCountTaskComplete();

        startFuncOnthisUrl("http://127.0.0.1:5500/html/ToDoList.html", HomePage);
        startFuncOnthisUrl("http://127.0.0.1:5500/html/CompleteTasks.html", CompletedPage);
        startFuncOnthisUrl("http://127.0.0.1:5500/html/ToComplete.html", toDoTasksPage);

        addEventListenerById("button", "click", () => {
            addTask();
        });

        addEventListenerById("button-complete", "click", () => {
            displayCompleteTasks();
        });

        addEventListenerById("button-toDo", "click", () => {
            displayToDOTasks();
        });

        addEventListenerById("button-AllTasks", "click", () => {
            displayAllTasks();
        });

        addEventListenerById("button-check-all", "click", () => {
            checkAllTasks();
        });

        addEventListenerById("button-remove-complete-task", "click", () => {
            removeAllTasks();
        });

    }, 200);

    displayAllTasks();
}

main();

// -----------------------------------------------------------------------
// ------------------------- Affichage des tasks -------------------------
// -----------------------------------------------------------------------

function displayAllTasks() {
    window.history.pushState({}, "/html/ToDoList.html", "/html/ToDoList.html");
    buttonActive = "--- All Tasks ---"
    newTask(toDoTaskList, buttonActive)
    eventCheckTaskList()
    eventSuppTaskList()
}

function displayCompleteTasks() {
    let listTasksChangeCheck = checkedTasksOnly();
    // Met à jour l'historique du navigateur sans recharger la page
    window.history.pushState({}, "/html/CompleteTasks.html", "/html/CompleteTasks.html");
    buttonActive = "--- Tasks Completes ---";
    newTask(listTasksChangeCheck, buttonActive);
    eventCheckTaskList(listTasksChangeCheck);
    eventSuppTaskList(listTasksChangeCheck);
}

function displayToDOTasks() {
    window.history.pushState({}, "/html/ToComplete.html", "/html/ToComplete.html");
    let listTasksChangeCheck = [];
    for (let i = 0; i < toDoTaskList.length; i++) {
        if (toDoTaskList[i].check === false) {
            listTasksChangeCheck.push(toDoTaskList[i]);
        }
    }
    buttonActive = "--- Tasks To Do ---";
    newTask(listTasksChangeCheck, buttonActive);
    eventCheckTaskList(listTasksChangeCheck);
    eventSuppTaskList(listTasksChangeCheck);
}

// -----------------------------------------------------------------------
// -------------- Gestion button ajout/modification task(s) --------------
// -----------------------------------------------------------------------

function addTask() {
    window.history.pushState({}, "/html/ToDoList.html", "/html/ToDoList.html");
    buttonActive = "--- All Tasks ---";
    let newtask = { ...task }
    // console.log(toDoTaskList)
    newtask.name = getElementValueById("input_todolist");
    // console.log(newtask)
    if (!isInArray(toDoTaskList, newtask.name) && newtask.name !== "") {
        toDoTaskList.push(newtask);
        newTask(toDoTaskList, buttonActive);
        eventCheckTaskList();
        eventSuppTaskList();
        localStorage.setItem('toDoTaskList', JSON.stringify(toDoTaskList));
        updateCountTaskComplete();
    };
}

function checkAllTasks() {
    window.history.pushState({}, "/html/ToDoList.html", "/html/ToDoList.html");
    buttonActive = "--- All Tasks ---";
    for (let i = 0; i < toDoTaskList.length; i++) {
        toDoTaskList[i].check = true;
    }
    newTask(toDoTaskList, buttonActive);
    eventCheckTaskList();
    eventSuppTaskList();
    updateCountTaskComplete();
}

function removeAllTasks() {
    window.history.pushState({}, "/html/ToDoList.html", "/html/ToDoList.html");
    buttonActive = "--- All Tasks ---";
    toDoTaskList = NocheckedTasksOnly();
    localStorage.setItem('toDoTaskList', JSON.stringify(toDoTaskList));
    newTask(toDoTaskList, buttonActive);
    eventCheckTaskList();
    eventSuppTaskList();
    updateCountTaskComplete();
}

// -----------------------------------------------------------------------
// --------------------------- Fonction d'event --------------------------
// -----------------------------------------------------------------------

// On met à jour le check
function eventCheckTaskList(listTasksChangeCheck) {
    for (let i = 0; i < toDoTaskList.length; i++) {
        if (!listTasksChangeCheck) {
            addEventListenerById(toDoTaskList[i].name, "click", () => {
                toDoTaskList[i].check = !toDoTaskList[i].check;
                localStorage.setItem('toDoTaskList', JSON.stringify(toDoTaskList));
                updateCountTaskComplete();
            });
        } else {
            for (let a = 0; a < listTasksChangeCheck.length; a++) {
                if (toDoTaskList[i] === listTasksChangeCheck[a]) {
                    addEventListenerById(toDoTaskList[i].name, "click", () => {
                        toDoTaskList[i].check = !toDoTaskList[i].check;
                        removeChildsElement("tasks");
                        listTasksChangeCheck = listTasksChangeCheck.filter((task) => task !== listTasksChangeCheck[a]);
                        newTask(listTasksChangeCheck, buttonActive);
                        eventCheckTaskList(listTasksChangeCheck);
                        eventSuppTaskList(listTasksChangeCheck);
                        localStorage.setItem('toDoTaskList', JSON.stringify(toDoTaskList));
                        updateCountTaskComplete();
                    });
                }
            }
        }
    }

}

function eventSuppTaskList(listTasksChangeCheck) {
    for (let i = 0; i < toDoTaskList.length; i++) {
        if (!listTasksChangeCheck) {
            addEventListenerById(`${toDoTaskList[i].name}-remove`, "click", () => {
                // Filtrer la liste pour enlever la tâche cible
                let updatedTaskList = toDoTaskList.filter(function (task) {
                    return task !== toDoTaskList[i];
                });
                toDoTaskList = updatedTaskList;

                newTask(toDoTaskList, buttonActive);
                eventCheckTaskList();
                eventSuppTaskList();
                // Afficher la liste mise à jour
                localStorage.setItem('toDoTaskList', JSON.stringify(toDoTaskList));
                updateCountTaskComplete();
            });

        } else {
            for (let a = 0; a < listTasksChangeCheck.length; a++) {
                if (listTasksChangeCheck[a] === toDoTaskList[i]) {
                    addEventListenerById(`${toDoTaskList[i].name}-remove`, "click", () => {
                        // Filtrer la liste pour enlever la tâche cible
                        let updatedTaskList = toDoTaskList.filter(function (task) {
                            return task !== toDoTaskList[i];
                        });
                        toDoTaskList = updatedTaskList;
                        let updatedTaskList1 = listTasksChangeCheck.filter(function (task) {
                            return task !== listTasksChangeCheck[a];
                        });
                        listTasksChangeCheck = updatedTaskList1;

                        newTask(listTasksChangeCheck, buttonActive);
                        eventCheckTaskList();
                        eventSuppTaskList();
                        localStorage.setItem('toDoTaskList', JSON.stringify(toDoTaskList));
                        updateCountTaskComplete();
                    });
                }
            }
        }
    }
}

// -----------------------------------------------------------------------
// -------------------------- Gestion des pages --------------------------
// -----------------------------------------------------------------------

function HomePage() {
    buttonActive = "--- All Tasks ---";
    newTask(toDoTaskList, buttonActive);
    eventCheckTaskList();
    eventSuppTaskList();
}

function CompletedPage() {
    let listTasksChangeCheck = [];
    for (let i = 0; i < toDoTaskList.length; i++) {
        if (toDoTaskList[i].check === true) {
            listTasksChangeCheck.push(toDoTaskList[i]);
        }
    }
    buttonActive = "--- Tasks Completes ---";
    newTask(listTasksChangeCheck, buttonActive);
    eventCheckTaskList(listTasksChangeCheck);
    eventSuppTaskList(listTasksChangeCheck);
}

function toDoTasksPage() {
    let listTasksChangeCheck = [];
    for (let i = 0; i < toDoTaskList.length; i++) {
        if (toDoTaskList[i].check === false) {
            listTasksChangeCheck.push(toDoTaskList[i]);
        }
    }
    buttonActive = "--- Tasks To Do ---";
    newTask(listTasksChangeCheck, buttonActive);
    eventCheckTaskList(listTasksChangeCheck);
    eventSuppTaskList(listTasksChangeCheck);
}

// -----------------------------------------------------------------------
// --------------------------- Fonction annexe ---------------------------
// -----------------------------------------------------------------------

function updateCountTaskComplete() {
    const array = checkedTasksOnly(toDoTaskList);
    const len = array.length;
    replaceTextContentFromElement("count-tasks-complete", ancienValueCountTasks, len);
    ancienValueCountTasks = len;
}

function checkedTasksOnly() {
    let array = [];
    for (let i = 0; i < toDoTaskList.length; i++) {
        if (toDoTaskList[i].check === true) {
            array.push(toDoTaskList[i]);
        }
    }
    return array
}

function NocheckedTasksOnly() {
    let array = [];
    for (let i = 0; i < toDoTaskList.length; i++) {
        if (toDoTaskList[i].check === false) {
            array.push(toDoTaskList[i]);
        }
    }
    return array
}