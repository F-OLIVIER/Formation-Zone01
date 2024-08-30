# SMART-ROAD

## 📝 Descriptif
Ce projet consiste à créer en `Rust` un programme avec la dépendance `sdl2`. Ce programme vise à voir une intersection à laquelle des véhicules vont se croiser, mais où il n'y aura jamais de collision. Il s'agit d'une version plus complexe d'un autre projet intitulé _Road-Intersection_.

__Voici les touches pour faire apparaitre des véhicules :__
- `↑` Up: un véhicule apparait au sud.
- `↓` Down: un véhicule apparait au nord.
- `→` Right: un véhicule apparait à l'ouest.
- `←` Left: un véhicule apparait à l'est.
- `r` : un véhicule apparait de façon aléatoire (nord, sud, ouest ou est).
- `s` : Affiche/cache les statistiques.
- `Esc` Escape: Affiche les statistiques (Cliquez sur le bouton 'Exit' pour quitter le programme).

**Image du projet**
<table align=center>
    <tr>
        <td><img src="src/assets/Capture d’écran du 2024-08-26 17-16-50.png"></td>
        <td><img src="src/assets/Capture d’écran du 2024-08-26 17-16-38.png"></td>
        <td><img src="src/assets/Capture d’écran du 2024-08-26 16-23-22.png"></td>
    </tr>
</table>

___
## ⚙️ Installation & usage

**Installation des dépendence pour linux**
```sh
sudo apt install libsdl2-dev 
sudo apt install libsdl2-image-dev
sudo apt install libsdl2-ttf-dev
```

**Installation de la bibliothèque SDL2 pour Windows**
1. Téléchargez la version MSVC de SDL2 depuis [http://www.libsdl.org/](http://www.libsdl.org/) (généralement nommée quelque chose comme `SDL2-devel-2.x.x-VC.zip`).
2. Décompressez `SDL2-devel-2.x.x-VC.zip`.
3. Copiez tous les fichiers `.lib` de `SDL2-2.x.x\lib\x64\` vers `%userprofile%\.rustup\toolchains\stable-x86_64-pc-allwindows-msvc\lib\rustlib\x86_64-pc-windows-msvc\lib\`.
4. Ajoutez `SDL2-2.x.x\lib\x64\SDL2.dll` dans vos variables d'environnement.

**Mise en route du programme**
```sh
cargo run
```

___
## 🔗 Dépendences

Le programme utilise `Rust` avec la `version 1.79.0` de [cargo](https://www.rust-lang.org/fr) et les imports suivants :<br>
- [rand](https://docs.rs/rand/latest/rand/) version `0.8.5`
- [sdl2](https://docs.rs/sdl2/0.37.0/sdl2/index.html) version `0.37.0` | features = ["image", "ttf"]
- [tokio](https://docs.rs/tokio/latest/tokio/) version `1` | features = ["full"]
- [chrono](https://docs.rs/chrono/latest/chrono/) version `0.4`

___
## 🧑‍💻 Authors
+ Fabien OLIVIER
+ Raphaël LOVERGNE 
+ Axelle FOUQUEMBERG
