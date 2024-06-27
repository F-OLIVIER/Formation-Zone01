//créer le tag ou l'input de ton choix en fonction de sa configuration
export const virtualElement = {
    childs: [],           //Définit les balises et les inputs qui se trouveront à l'intérieur de cette element
    appendChild: function (child) { //Met un élèment virtuel à l'intérieur de childs
        let array = []
        for (let i = 0; i < this.childs.length; i++) {
            array.push(this.childs[i]);
        }
        array.push(child)
        this.childs = array;
    },
    removeChild: function (childElementTarget, childID) {//Retire un élèment virtuel à l'intérieur de childs
        let array = []
        for (let i = 0; i < this.childs.length; i++) {
            if (this.childs[i].element === childElementTarget) {
                if (childID) {
                    if (this.childs[i].id != childID) {
                        array.push(this.childs[i])
                    }
                }
            } else {
                array.push(this.childs[i])
            }
        }
        this.childs = array;
    },
    element: "",        // Le type de l'élément (pour les <input> seulement)
    textContent: "",    // Le contenu texte de l'élément
    class: "",          // La classe CSS de l'élément
    id: "",             // L'identifiant unique de l'élément
    placeholder: "",    // Le texte de substitution pour les éléments de formulaire
    value: "",          // Valeur saisie pour les éléments de formulaire
    name: "",           // Le nom de l'élément (utilisé lors de l'envoi de formulaires)
    style: "",          // Les styles CSS en ligne pour l'élément
    autocomplete: "",   // Indique si l'autocomplétion doit être activée pour l'élément
    disabled: "",       // Indique si l'élément est désactivé ou non
    readonly: "",       // Indique si l'élément est en lecture seule ou non
    required: "",       // Indique si l'élément est requis ou non
    min: "",            // Définit la valeur minimale acceptée pour les éléments numériques
    max: "",            // Définit la valeur maximale acceptée pour les éléments numériques
    step: "",           // Définit l'incrément entre les valeurs numériques pour les éléments de type "number"
    pattern: "",        // Définit une expression régulière pour valider la valeur de l'élément
    hidden: false,      // Définit l'élément comme non visible
    onClick: "",         // Définit un event comme l'appelle d'une function js par exemple
    type: "",            // Donne le type de l'input (text, email...)
    action: "",          // Définit l'action à realiser lors du submit du formulaire
    method: "",          // Définit la méthod lorsque il est submit
    checked: false,
    addEventListener: {
        event: "",
        callback: () => {
            // Corps de la fonction de rappel
        },
        key: ""
    },
    href: "",
};


//list des tags possible à mettre comme "element" pour virtualElement.element
export const listTag = {
    div: "div",
    p: "p",
    span: "span",
    a: "a",
    img: "img",
    input: "input",
    button: "button",
    ul: "ul",
    ol: "ol",
    li: "li",
    h1: "h1",
    h2: "h2",
    h3: "h3",
    h4: "h4",
    h5: "h5",
    h6: "h6",
    table: "table",
    tr: "tr",
    td: "td",
    th: "th",
    form: "form",
    label: "label",
    textarea: "textarea",
    select: "select",
    option: "option",
    iframe: "iframe",
    video: "video",
    audio: "audio",
    nav: "nav",
    section: "section",
    header: "header",
    footer: "footer",
    main: "main",
    article: "article",
    aside: "aside",
    figure: "figure",
    figcaption: "figcaption",
    details: "details",
    summary: "summary",
    menu: "menu",
    submenu: "menu",
    legend: "legend",
    fieldset: "fieldset",
    iframe: "iframe",
    canvas: "canvas",
    svg: "svg",
    path: "path",
    circle: "circle",
    rect: "rect",
    ellipse: "ellipse",
    line: "line",
    polygon: "polygon",
    polyline: "polyline",
    defs: "defs",
    symbol: "symbol",
    use: "use",
    marker: "marker",
    linearGradient: "linearGradient",
    radialGradient: "radialGradient",
    stop: "stop",
    mask: "mask",
    pattern: "pattern",
    clipPath: "clipPath",
    text: "text",
    tspan: "tspan",
    textPath: "textPath",
    foreignObject: "foreignObject"
};

//list de tout les types d'inputs pour l'attribution dans virtualElement.type
export const listInputType = {
    text: "text",
    password: "password",
    email: "email",
    number: "number",
    checkbox: "checkbox",
    radio: "radio",
    file: "file",
    submit: "submit",
    reset: "reset",
    button: "button",
    date: "date",
    time: "time",
    datetime: "datetime",
    month: "month",
    week: "week",
    color: "color",
    range: "range",
    tel: "tel",
    url: "url",
    search: "search"
};
