import { virtualElement, listInputType } from "./script/virtualElement.js";
import { addEventListenerById, getElementValueById, insertToElement, insertToHtml, removeChildsElement } from "./script/virtualDOM.js";
import { multipleVirtualElement } from "./script/utils.js";
import { main } from "./main.js";

// -----------------------------------------------------------------------
// ----------------------- création du DOM virtuel -----------------------
// -----------------------------------------------------------------------
export function toDolist() {

    // Création des objets virtuels
    const arrayVirtualElement = multipleVirtualElement(virtualElement, 11);
    const titreObj = titleDiv(arrayVirtualElement[0]);
    const input = inputToDolist(arrayVirtualElement[1]);
    const boutonAdd = buttonAddToDoList(arrayVirtualElement[2]);
    const boutonCompleteTasks = buttonCompleteTasks(arrayVirtualElement[3]);
    const boutonNonCompleteTasks = buttonNonCompleteTasks(arrayVirtualElement[4]);
    const divboutons = arrayVirtualElement[5];
    const boutonAllTasks = buttonAllTasks(arrayVirtualElement[6]);
    const buttonCheckAll = buttonCheckAllTask(arrayVirtualElement[7]);
    const buttonRemoveAllCompletedTasks = buttonRemoveAllTaskComplete(arrayVirtualElement[8])
    const numberOfCompletedTasks = countTaskCompleteDIV(arrayVirtualElement[9])
    const footer = footerTask(arrayVirtualElement[10])

    titreObj.appendChild(input)
    titreObj.appendChild(boutonAdd)
    divboutons.appendChild(boutonAllTasks)
    divboutons.appendChild(boutonNonCompleteTasks)
    divboutons.appendChild(boutonCompleteTasks)
    divboutons.appendChild(buttonCheckAll)
    divboutons.appendChild(buttonRemoveAllCompletedTasks)
    divboutons.id = "menu-buttons"

    const globalVirtualElement = { ...virtualElement }
    globalVirtualElement.id = "global"
    globalVirtualElement.appendChild(titreObj);
    globalVirtualElement.appendChild(divboutons);
    globalVirtualElement.appendChild(numberOfCompletedTasks);
    insertToHtml(globalVirtualElement);
    insertToHtml(footer);

    // console.log(titreObj)
}

function countTaskCompleteDIV(virtualE) {
    virtualE.textContent = "Number of tasks completes: 0"
    virtualE.id = "count-tasks-complete"
    return virtualE
}

function titleDiv(virtualE) {
    virtualE.textContent = "ToDoList: "
    virtualE.class = "title"
    return virtualE
}

function buttonCheckAllTask(virtualE) {
    virtualE.element = "button";
    virtualE.textContent = "Check all tasks";
    virtualE.id = "button-check-all";
    virtualE.href = "/AllTasks";
    return virtualE
}

function buttonRemoveAllTaskComplete(virtualE) {
    virtualE.element = "button";
    virtualE.textContent = "Remove completed tasks";
    virtualE.id = "button-remove-complete-task"
    return virtualE
}

function inputToDolist(virtualE) {
    virtualE.placeholder = "Give a new task to do";
    virtualE.id = "input_todolist";
    virtualE.type = listInputType.text;
    virtualE.value = "";
    virtualE.element = "input";
    return virtualE;
}

function buttonAddToDoList(virtualE) {
    virtualE.element = "button";
    virtualE.textContent = "Add task";
    virtualE.id = "button"

    return virtualE;
}

function buttonCompleteTasks(virtualE) {
    virtualE.element = "button";
    virtualE.textContent = "Complete tasks"
    virtualE.id = "button-complete"
    return virtualE
}

function buttonNonCompleteTasks(virtualE) {
    virtualE.element = "button";
    virtualE.textContent = "Tasks to complete"
    virtualE.id = "button-toDo"
    return virtualE
}

function buttonAllTasks(virtualE) {
    virtualE.element = "button";
    virtualE.textContent = "All tasks"
    virtualE.id = "button-AllTasks"
    return virtualE
}

function footerTask(virtualE) {
    virtualE.element = "footer";
    const arrayVirtualElement = multipleVirtualElement(virtualElement, 2);
    const explication1 = div(arrayVirtualElement[0], "Double-click to edit a todo when ✏️");
    virtualE.appendChild(explication1);
    const explication2 = div(arrayVirtualElement[1], "Create by Z01's team");
    virtualE.appendChild(explication2);
    return virtualE
}

function div(virtualE, text) {
    virtualE.textContent = text
    return virtualE
}

// -----------------------------------------------------------------------
// -------------------- Création div task dans le DOM --------------------
// -----------------------------------------------------------------------

export function newTask(toDoTaskList, buttonAction) {
    creationDivTask(buttonAction)

    const array = multipleVirtualElement(virtualElement, 1);
    const menu = array[0];
    menu.textContent = buttonAction;
    menu.id = "menu"
    insertToElement("tasks", menu);

    for (let i = 0; i < toDoTaskList.length; i++) {

        const ve = multipleVirtualElement(virtualElement, 4);
        const newtask = ve[0];
        const inputNewTask = ve[1];
        const removeTask = ve[2];
        const divcontent = ve[3];

        // newtask.textContent = toDoTaskList[i].name + ": ";
        divcontent.textContent = toDoTaskList[i].name;
        divcontent.id = "global-" + toDoTaskList[i].name;
        newtask.class = "newTask";
        inputNewTask.element = "input";
        inputNewTask.type = "checkbox";
        inputNewTask.id = toDoTaskList[i].name;
        removeTask.element = "button";
        removeTask.class = "removeTask"
        removeTask.textContent = "X";
        removeTask.id = toDoTaskList[i].name + "-remove"

        if (toDoTaskList[i].check) {
            inputNewTask.checked = true
        }

        newtask.appendChild(divcontent);
        newtask.appendChild(inputNewTask);
        newtask.appendChild(removeTask);

        insertToElement("tasks", newtask);

        addEventListenerById("global-" + toDoTaskList[i].name, "dblclick", () => {
            updatetask(toDoTaskList[i].name);
        })
    }
}

function creationDivTask() {
    let div;
    if (document.getElementById("tasks")) {
        removeChildsElement("tasks");
    } else {
        const veDIV = multipleVirtualElement(virtualElement, 1)
        div = veDIV[0];
        div.id = "tasks"
        insertToElement('global', div);
    }
}

// -----------------------------------------------------------------------
// ----------------- Div d'update d'une task dans le DOM -----------------
// -----------------------------------------------------------------------

function updatetask(id) {
    const ve = multipleVirtualElement(virtualElement, 3);
    const title = ve[0];
    title.textContent = "Edit task";
    title.class = "title";
    const textarea = ve[1];
    textarea.textContent = id;
    textarea.id = "contentedittask";
    textarea.element = "textarea";
    const button = ve[2];
    button.element = "button";
    button.textContent = "Update task";
    button.id = "buttonedittask";

    const divedittask = { ...virtualElement }
    divedittask.id = "divedittask"
    divedittask.appendChild(title);
    divedittask.appendChild(textarea);
    divedittask.appendChild(button);
    const globaldiv = { ...virtualElement }
    globaldiv.id = "globaldiv";
    globaldiv.appendChild(divedittask);
    insertToHtml(globaldiv);

    eventEditTask(id);
}

export function eventEditTask(oldvalue) {
    addEventListenerById("buttonedittask", "click", async () => {
        const newvalue = await getElementValueById('contentedittask');
        updateStorage(oldvalue, newvalue);
        // supression de la div d'edition
        removeChildsElement("body");
        // re-création du body
        main();
    })
}

// -----------------------------------------------------------------------
// --------------------------- Fonction annexe ---------------------------
// -----------------------------------------------------------------------

function updateStorage(oldvalue, newvalue) {
    const storedToDoTaskList = localStorage.getItem('toDoTaskList');
    let newToDoTaskList = [];
    if (storedToDoTaskList) {
        newToDoTaskList = JSON.parse(storedToDoTaskList);
    }

    for (let i = 0; i < newToDoTaskList.length; i++) {
        let element = newToDoTaskList[i];
        if (element.name === oldvalue) {
            element.name = newvalue
            break;
        }
    }
    localStorage.setItem('toDoTaskList', JSON.stringify(newToDoTaskList));
}

export function isInArray(array, cible) {
    for (let i = 0; i < array.length; i++) {
        if (array[i].name === cible) {
            return true
        }
    }
    return false
}
