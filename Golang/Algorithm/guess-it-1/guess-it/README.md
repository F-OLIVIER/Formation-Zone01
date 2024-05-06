# guess-it

## Usage

Vous devrez d'abord copier le dossier `student/` (fourni par l'étudiant) dans lequel vous trouverez le programme guessing de l'étudiant ainsi qu'un fichier appelé `script.sh`. Ce fichier doit être un script shell exécutable qui lance le programme de l'étudiant si vous êtes dans le dossier racine `guess-it/`. Le système de fichiers devrait ressembler à quelque chose comme ceci :

```console
─ guess-it/
├── ai/
│   ├── big-range
│   └── ...
├── index.html
├── index.js
└── ...
└── student/
    ├── ...
    └── script.sh

```

In order to test the student program, these commands should be ran to have the dependencies needed and to start the webpage on the port 3000:

```console
npm install
node server.js
```

You will need to run also this command inside the `ai/` directory to make the programs executable:

```console
chmod +x *
```

After opening your browser of preference in the port [3000](http://localhost:3000/), if you try clicking on any of the `Test Data` buttons, you will notice that in the Dev Tool/ Console there is a message which tells you that you need another guesser besides the student.

Adding a guesser is simple. You just need to add in the URL a guesser, in other words, the name of one of the files present in the `ai/` folder:

```console
?guesser=<name_of_guesser>
```

For example:

```console
?guesser=big-range
```

After that, choose which of the random data set to test. After that you can wait for the program to test all of the values (boooooring), or you can click `Quick` in order to skip the waiting and be presented with the results.

Since the website uses big data sets, we advise you to clear the displays clicking on the `Clean` button after each test.
