// Fonction pour créer une représentation HTML à partir d'un élément du Virtual DOM
export function createElementTag(virtualE) {

    //Au cas où qu'on oublie de désigner un élèment
    let html = `<div`
    if (virtualE.element) {
        html = `<${virtualE.element}`;
    }

    //On rajoute les attrs de la balise qui été configuré au préalable
    const attr = attrToGive(virtualE);
    if (attr.length > 0) {
        html += ` ${attr.join(" ")}`;
    }

    html += '>';

    // Si textContent existe alors on le rajoute
    if (virtualE.textContent && virtualE.element != "form") {
        html += virtualE.textContent;
    }

    //Si la balise contient d'autres élèments on les rajoutes après le textContent
    if (virtualE.childs) {
        for (let i = 0; i < virtualE.childs.length; i++) {
            let child = "";
            if (virtualE.childs[i].element === "input") {
                child = createElementInput(virtualE.childs[i])
            } else if (virtualE.childs[i].element === "form") {
                child = createElementForm(virtualE.childs[i])
            } else {
                child = createElementTag(virtualE.childs[i])
            }
            html += child
        }
    }
    if (virtualE.element) {
        html += `</${virtualE.element}>`;
    } else {
        html += `</div>`;
    }

    return html;
}

//Fonction pour créer un input avec ces inputs si ceci existent
export function createElementInput(virtualE) {
    let html = `<input`;


    // On rajoute les attrs de la balise qui ont été configurés au préalable
    const attr = attrToGive(virtualE);
    if (attr.length > 0) {
        html += ` ${attr.join(" ")}`;
    }

    if (virtualE.checked === true) {
        html += `checked`
    }

    html += ` >`

    if (virtualE.childs) {
        for (let i = 0; i < virtualE.childs.length; i++) {
            let child = "";
            if (virtualE.childs[i].element === "input") {
                child = createElementInput(virtualE.childs[i])
            } else if (virtualE.childs[i].element === "form") {
                child = createElementForm(virtualE.childs[i])
            } else {
                child = createElementTag(virtualE.childs[i])
            }
            html += child
        }
    }

    return html;
}

//Fonction pour créer un input avec ces inputs si ceci existent
export function createElementForm(virtualE) {
    let html = `<form`;


    // On rajoute les attrs de la balise qui ont été configurés au préalable
    const attr = attrToGive(virtualE);
    if (attr.length > 0) {
        html += ` ${attr.join(" ")}`;
    }

    html += `>`

    if (virtualE.childs) {
        for (let i = 0; i < virtualE.childs.length; i++) {
            let child = "";
            if (virtualE.childs[i].element === "input") {
                child = createElementInput(virtualE.childs[i])
            } else if (virtualE.childs[i].element === "form") {
                child = createElementForm(virtualE.childs[i])
            } else {
                child = createElementTag(virtualE.childs[i])
            }
            html += child
        }
    }

    return html;
}


//Return tout les attribues qui ont recu une value
function attrToGive(virtualE) {
    let array = [];
    for (let i = 0; i < Object.keys(virtualE).length; i++) {
        const key = Object.keys(virtualE)[i];
        const value = virtualE[key];
        if (value && key != "element" && key != "textContent" && key != "childs" && key != "appendChild" && key != "removeChild" && key != "addEventListener" && key != "events") {
            array.push(`${key}="${value}"`);
        };
    };
    return array;
};

//insert les élèments htmls dans le document html souhaiter
export function insertToHtml(html) {
    const element = html.element
    let htmlString = ""


    switch (element) {
        case ("input"):
            htmlString = createElementInput(html)
            break
        case ("form"):
            htmlString = createElementForm(html)
            break
        default:
            htmlString = createElementTag(html)
            break
    }
    // console.log('htmlString : ', htmlString)
    document.body.insertAdjacentHTML("beforeend", htmlString);


    if (html.addEventListener.event != "") {
        addEventListenerById(html.id, html.addEventListener.event, html.addEventListener.callback, html.addEventListener.key)
    }

}

//Permet d'inserer un child dans un élèment à partir de son id
export function insertToElement(id, html) {
    const element = document.querySelector(`#${id}`);
    if (element) {
        const elementHTML = html.element
        let htmlString = ""
        switch (elementHTML) {
            case ("input"):
                htmlString = createElementInput(html)
                break
            case ("form"):
                htmlString = createElementForm(html)
                break
            default:
                htmlString = createElementTag(html)
                break
        }
        // console.log('htmlString ' + id, htmlString)
        document.getElementById(id).insertAdjacentHTML("beforeend", htmlString);
    } else {
        console.error(`Element with ID '${id}' not found.`);
    }
}

export function addEventListenerById(id, event, callback, key) {
    const element = document.getElementById(id);
    if (element) {
        if (event === "click") {
            element.onclick = callback;
        } else if (event === "dblclick") {
            element.ondblclick = callback;
        } else if (event === "keydown" && key) {
            element.onkeydown = function (event) {
                if (event.key === key) {
                    callback();
                }
            };
        } else {
            console.error(`Event '${event}' is not supported or no key specified.`);
        }
    } else {
        console.error(`Element with ID '${id}' not found.`);
    }
}


// Fonction pour supprimer un attribut spécifique d'un élément à partir de son id
export function removeAttrFromElement(id, attributeName) {
    const element = document.querySelector(`#${id}`);
    if (element) {
        element.removeAttribute(attributeName);
    } else {
        console.error(`Element with ID '${id}' not found.`);
    }
}

// remplace une chaîne de caractères spécifique de l'élément à partir de son id
export function replaceTextContentFromElement(id, stringToRemove, stringToplace) {
    const element = document.querySelector(`#${id}`);
    if (element) {
        element.innerHTML = element.innerHTML.replace(String(stringToRemove), String(stringToplace));
    } else {
        console.error(`Element with ID '${id}' not found.`);
    }
}


//Retourne la value d'un élèment
export function getElementValueById(id) {
    return document.getElementById(id).value
}

//Enlève les enfants de l'élèment
export function removeChildsElement(id) {
    const container = document.getElementById(id);
    container.innerHTML = ""; // Effacer le contenu existant
}