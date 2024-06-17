const emojis = [
    "😀", "😃", "😄", "😁", "😆", "😅", "😂", "😊", "😇", "🙃", "😉", "😌", "😍", "😘",
    "😗", "😙", "😚", "😋", "😛", "😜", "😝", "😐", "😑", "😶", "😏", "😒", "😌", "😔",
    "😪", "😴", "😷"
];

export function chatEmojis() {
    const emojiPicker = document.getElementById('emoji-picker');
    const emojiButton = document.getElementById('emojis');
    const inputMessage = document.getElementById('inputmessage');

    // Ajouter les emojis dans le picker
    emojis.forEach(emoji => {
        const emojiDiv = document.createElement('div');
        emojiDiv.textContent = emoji;
        // action quand l'utilisateur clique sur un emoji
        emojiDiv.addEventListener('click', () => {
            inputMessage.value += emoji;
            emojiPicker.classList.remove('active');
        });
        emojiPicker.appendChild(emojiDiv);
    });

    // Afficher ou masquer le picker d'emojis lors du clic sur le bouton emoji
    emojiButton.addEventListener('click', (event) => {
        event.stopPropagation();
        if (emojiPicker.classList.contains('active')) {
            emojiPicker.classList.remove('active');
        } else {
            emojiPicker.classList.add('active');
        }
    });

    // Masquer le picker d'emojis lorsque l'utilisateur clique en dehors de la fenetre emoji
    document.addEventListener('click', (event) => {
        if (!emojiPicker.contains(event.target)) {
            emojiPicker.classList.remove('active');
        }
    });
}