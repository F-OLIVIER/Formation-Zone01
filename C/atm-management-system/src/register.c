#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <termios.h>
#include "header.h"

// Fonction qui vérifie le mdp lors de la connexion
void registerMenu(char a[50], char pass[50])
{
    struct termios oflags, nflags;

wrongUser:
    system("clear");
    printf("\n\n\n\t\t\t\t   Bank Management System\n\n\t\t\t\t enter new Login:");
    scanf("%s", a); // nom utilisateur

    // désactivation de l'écho
    tcgetattr(fileno(stdin), &oflags);
    nflags = oflags;
    nflags.c_lflag &= ~ECHO;
    nflags.c_lflag |= ECHONL;

    if (tcsetattr(fileno(stdin), TCSANOW, &nflags) != 0)
    {
        perror("tcsetattr");
        return exit(1);
    }

    printf("\n\n\t\t\t\tEnter password:");
    scanf("%s", pass); // mot de passe utilisateur

    // restaurer le terminal
    if (tcsetattr(fileno(stdin), TCSANOW, &oflags) != 0)
    {
        perror("tcsetattr");
        return exit(1);
    }

    // création de l'utilisateur
    if (strlen(a) > 0 && strlen(pass) > 0)
    {
        FILE *fp;
        struct User userChecker;

        if ((fp = fopen("./data/users.txt", "r+")) == NULL)
        {
            printf("Error! opening file");
            exit(1);
        }

        // récupération dernier id utilisateur
        int last_id = -1;
        while (fscanf(fp, "%d %s %s", &userChecker.id, userChecker.name, userChecker.password) != EOF)
        {
            if (userChecker.id > last_id)
            {
                last_id = userChecker.id;
                printf("\nlogin : %d, %d\n", last_id, userChecker.id);
            }

            // vérification si le nom d'utilisateur existe déjà
            if (strcmp(a, userChecker.name) == 0)
            {
                printf("\n✖ Existing user name");
                goto wrongUser;
            }
        }
        last_id++;
        fseek(fp, 0, SEEK_END); // Déplacer le curseur à la fin du fichier
        // Ecriture du nouvel utilisateur dans le fichier user.txt
        fprintf(fp, "%d %s %s\n", last_id, a, pass);

        fclose(fp);
        printf("\nnew user create");
    }
    else
    { // login ou mot de passe vide
        printf("\n✖ Empty password!! or Empty Name\n");
        goto wrongUser;
    }
};
