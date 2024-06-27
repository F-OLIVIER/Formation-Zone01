# MINI-FRAMEWORK

## 📝 Introduction

Bienvenue dans ce mini-framework !!! L'objectif de celui-ci est de répondre aux attentes de l'exercice suivant :</br>
<a href="https://github.com/01-edu/public/tree/master/subjects/mini-framework#instructions">https://github.com/01-edu/public/tree/master/subjects/mini-framework#instructions</a>.

Pour cela, nous avons créé des fonctions qui facilitent la création d'applications et d'autres projets tout en s'abstenant d'utiliser le DOM à outrance.

## ⚙️ Usage

### Script 

Script est le dossier du framework qui contient les 2 fichiers les plus important du projects car ils vont contenir les deux aspect du DOM virtuel qui sont :

- Les functions de manipulations du DOM virtuel ━🢂 virtualDOM.js
- Le `virtualElement` et autres lists ━🢂 virtualElement.js

🗒️ __NOTE:__</br>
utils.js sert juste pour le moment à stocker la function  `multipleVirtualElement` pour créer plusieurs copies de `virtualElement` à la fois. Néanmoins il peut servir pour vous à stocker des functions qui pourraient vous facilitez la manipulation de celles contenues dans virtualDOM.js
</br></br>

### 💻 VirtualElement (script/virtualElement.js)
___
**Basic Configuration**</br>
VirtualElement est un object/struct qui vous servira de première étape dans la configuration de votre élèment HTML. Pour l'utiliser il vous suffit de créer une copie de celui-ci par exemple :

```javascript 
const virtualE  = {...virtualElement}
```

ou en utilisant `multipleVirtualElement` dans script/utils.js pour avoir plusieur copie à la fois

```javascript 
const numb = 2
const arrayVirtuale  = multipleVirtualElement{virtualElement, numb}
const div = arrayVirtuale[0]
const button = arrayVirtuale[1]
```

Les paramêtres essentielles à son utilisation sont les suivantes:

- virtualElement.element ━🢂 `element` est une clé de l'object qui définit la nature de votre élèment. Si je lui dis qu'il est égale à `div` alors cette élèment sera une div mais vous pouvez également lui donner comme élèment celui d'`input`, de `form`, de `button` et bien d'autres encore...
- virtualElement.id ━🢂 permet de donner un id à l'élèment
- virtualElement.class ━🢂 lui donne une class
- virtualElement.type ━🢂 donne un type à un input (penser à `virtualElement.element="input"`)

Et ainsi de suite si vous voulez savoir le nombre attributs que vous pouvez configurer n'hésitez pas à regarder l'object virtualElement dans le chemin suivant `script/virtualElement.js`. Je précise aussi que presque tout les attributs se configurent génèralement grace à une string.

**Childs**</br>
Avant de finir sur cette partie, il est important de préciser que `virtualElement` ne contient pas seulement des key pour acceuillir des attribus(string) mais aussi des property qui vous serviront à ajouter des enfants à votre élèment : appendChild et removeChild.

- virtualElement.appendChild(otherVirtualElement) ━🢂 Cette property sert à acceuilir d'autres virtualElement au sein de votre élèment (dans virtualElement.childs plus précisèment).

✍️ __EXEMPLE:__
```javascript
const arrayVirtuale  = multipleVirtualElement{virtualElement, numb}
const div = arrayVirtuale[0]
const button = arrayVirtuale[1]

div.appenChild(button)
```

- virtualElement.removeChild(elementVirtualElement, idVirtualElement) ━🢂 Et au contraire celle-ci retire un élèment à condition de fournir l'élèment et l'id du virtualElement que vous souhaitez supprimé.

✍️ __EXEMPLE:__

```javascript
const arrayVirtuale  = multipleVirtualElement{virtualElement, numb}
const div = arrayVirtuale[0]
const button = arrayVirtuale[1]
button.id = "button1"
button.element="button"

div.appendChild(button)

div.removeChild("button", "button1")
```
</br>

📋 **Lists**

Si vous ne vous souvenez pas de certains noms de tags ou d'input n'hésitez pas à utiliser les functions `listsTags()` et `listInputType()` pour vous aider. Elles contiennent les noms sous format string, d'une majorité des tags et des types d'inputs.
</br></br>

### 💻 VirtualDOM (manipulation du document avec script/virtualDOM.js)
___
**Insert VirtualElement**

Maintenant passons aux functions qui vous permettrons de transformer ce virtualElement en véritable élèment html et de lui attacher des events!!!

Pour commencer nous allons voir comment l'insérer dans votre page html. Pour cela il vous faudra la function `insertToHtml()` qui sélèctionnera le body du document et rajouteras votre virtualElement tou en bas de celui ci. Voici comment l'utiliser :

✍️ __EXEMPLE:__
```javascript
const virtualE = {...virtualElement}

virtualE.element = "div"
virtualE.id = "test"

insertToHtml(virtualE)
```

Pour vous expliquez simplement son functionnement, celle-ci va créer une string à partir des keys et des values de l'object `virtualElement` pour ensuite l'introduire dans le corps de l'html. 

🗒️ __NOTE:__</br>
Pour visualiser votre virtualElement en ligne html, vous pouvez appellez les functions `createElementTags()`, `createElementInput()` et `createElementForm()`. 

Si vous souhaitez rajoutez un élèment dans un élèment déjà existant de votre page, utilisez la function `insertToElement()`, avec comme arguments l'id de l'élèment à qui vous souhaitez rajoutez un enfant et le virtualElement que vous souhaitez ajouter.

✍️ __EXEMPLE :__
```javascript
insertToElement(idElementCible, virtualElement)   //Tout simplement !!!!
```

**addEventListenerById**

Vous savez créer un virtualElement, le configurer et l'insérer dans une page html. Il ne vous manque plus qu'à lui ajouter un event !!!

Pour cela utilisez la function `addEventListenerById()` ou la property `visualElement.addEventListener`.

Elles nécessitent trois arguments : l'id de l'élèment, l'event (`click`, `dblclick` ou `keydown`) ainsi que la function qui se déclenchera au moment de l'event. Il y a un 4ème argument pour préciser qu'elle sera la touche qui déclanchera l'event si vous avez choisi comme event `keydown` (mais pas besoin de s'en soucier si l'event est un `click`).

✍️ __EXEMPLE:__
```javascript
function printEventMark(){
    console.log("it's work")
}

const virtualE = {...virtualElement}
virtualE.id="test"
virtualE.element= "button"
virtualE.addEventListener.event= "click";
virtualE.addEventListener.callback = printEventMark

insertToHtml(virtualE)
```

✍️ __EXEMPLE:__
```javascript
// Oh j'ai oublié d'utiliser la property addEventListener!!! 
// Pas de panique ;=)

// Pour rajouter un event simplement sans recréer un visualElement
addEventListener("test", "click", printEventMark()) // id, event, function 
```

**Others functions in virtualDOM.js**
</br>
- replaceTextContentFromElement(): Préciser le texteContent de l'élèment avec son id et remplacer le par ce que vous voulez
- getElementValueById(): Renvoie la value de l'élèment cible (id requis).
- removeChildsElement(): Retire tout les enfants d'un élèment ciblé (id requis).
</br></br>

## 🧑‍💻 Authors
- 🤙🏻 Armand AUVRAY
- 👨‍💼 Fabien OLIVIER

