//return un array de plusieurs virtualElement, mutliplier par le nombre qu'on veut
//Sert de banque pour utiliser des virtualElement sans le recréer à chaque fois 
export function multipleVirtualElement(virtualElement, count) {
    const result = [];
    for (let i = 0; i < count; i++) {
        // Crée une copie du virtualElement pour chaque itération
        const copy = { ...virtualElement };
        result.push(copy);
    }
    return result;
}

export function startFuncOnthisUrl(url, func) {
    const currentURL = window.location.href;

    if (currentURL === url) {
        func();
    }
}
