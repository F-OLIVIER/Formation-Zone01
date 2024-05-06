#include "header.h"

void maketransactionAcc(struct User u)
{
    char userName[100];
    struct Record r;

    FILE *pf = fopen(RECORDS, "r");

    system("clear");
wrongUpdate:
    printf("\n\t\t====== All accounts from user, %s =====\n\n", u.name);
    while (getAccountFromFile(pf, userName, &r))
    {
        if (strcmp(userName, u.name) == 0)
        {
            printf("_____________________\n");
            printf("\nAccount number: %d\nDeposit Date: %d/%d/%d \ncountry: %s \nPhone number: %d \nAmount deposited: $%.2f \nType Of Account: %s\n",
                   r.accountNbr,
                   r.deposit.day,
                   r.deposit.month,
                   r.deposit.year,
                   r.country,
                   r.phone,
                   r.amount,
                   r.accountType);
        }
    }
    fseek(pf, 0, SEEK_SET); // re-initialise le curseur au début du fichier

    printf("\n\n\t\tNumber of the account where you wish to make a transaction:");
    int accountNbrToTransaction;
    scanf("%d", &accountNbrToTransaction);

    while (getAccountFromFile(pf, userName, &r))
    {
        // printf("r.accountNbr : %d\n", r.accountNbr);
        if (strcmp(userName, u.name) == 0 && accountNbrToTransaction == r.accountNbr)
        {
            // printf("r.accountNbr dans le if : %d %s %d\n", r.accountNbr, r.accountType, r.deposit.month);
            if (strcmp(r.accountType, "current") == 0 || strcmp(r.accountType, "saving") == 0) // transaction autorisé
            {
                system("clear");
                printf("\n\n\t\t====== Update account %d from user, %s =====\n\n", accountNbrToTransaction, u.name);
                printf("\n\t\t-->> Feel free to choose one of the options below <<--\n");
                printf("\n\t\t[1]- withdrawal\n");
                printf("\n\t\t[2]- deposit\n");
                printf("\n\t\t[3]- Exit to the main menu\n");
                printf("\n\t\t[4]- Exit\n");
                printf("\nYour selection : ");
                int option;
                scanf("%d", &option);

                double trading; // variable à virgule
                switch (option)
                {
                case 1:
                    printf("\nHow much do you want to withdraw:");
                    scanf("%lf", &trading);
                    break;
                case 2:
                    printf("\nHow much do you want to deposit:");
                    scanf("%lf", &trading);
                    break;
                case 3:
                    mainMenu(u);
                    break;
                case 4:
                    exit(1);
                    break;
                default:
                    printf("\n✖ Invalid operation!\n\n");
                    // stayOrReturn(u);
                }

                if (option == 1 || option == 2)
                {
                    char *newFile; // Déclaration du pointeur vers la chaîne de caractères
                    // Allocation dynamique de mémoire pour la chaîne (par exemple, pour 100 caractères)
                    newFile = (char *)malloc(100 * sizeof(char));
                    char line[300];

                    fseek(pf, 0, SEEK_SET); // re-initialise le curseur au début du fichier avant la boucle while pour re-écrire le fichier avec les modifications
                    while (getAccountFromFile(pf, userName, &r))
                    {
                        if (strcmp(userName, u.name) == 0 && accountNbrToTransaction == r.accountNbr)
                        {
                            // printf("%s %d : %d", userName, accountNbrToTransaction, r.amount);
                            if (option == 1)
                            {
                                if (r.amount - trading > 0)
                                {
                                    snprintf(line, sizeof(line), "%d %d %s %d %d/%d/%d %s %d %.2lf %s\n\n", r.id, r.userId, userName, r.accountNbr, r.deposit.month, r.deposit.day, r.deposit.year, r.country, r.phone, r.amount - trading, r.accountType);
                                }
                                else
                                {
                                    printf("\n✖ The amount requested is greater than the amount available (possible)\n\n");
                                    goto wrongUpdate;
                                }
                            }
                            else if (option == 2)
                            {
                                snprintf(line, sizeof(line), "%d %d %s %d %d/%d/%d %s %d %.2lf %s\n\n", r.id, r.userId, userName, r.accountNbr, r.deposit.month, r.deposit.day, r.deposit.year, r.country, r.phone, r.amount + trading, r.accountType);
                            }
                        }
                        else
                        {
                            snprintf(line, sizeof(line), "%d %d %s %d %d/%d/%d %s %d %.2lf %s\n\n", r.id, r.userId, userName, r.accountNbr, r.deposit.month, r.deposit.day, r.deposit.year, r.country, r.phone, r.amount, r.accountType);
                        }

                        // Augmentation de la taille de newFile si line != vide
                        if (strlen(line) > 0)
                        {
                            // Réallocation pour ajouter la taille de line à newFile
                            char *temp = realloc(newFile, (strlen(newFile) + strlen(line) + 1) * sizeof(char));
                            if (temp == NULL)
                            {
                                printf("Reallocation de mémoire a échoué.");
                                free(newFile); // Libérer la mémoire précédemment allouée
                                exit(1);
                            }
                            else
                            {
                                newFile = temp;
                            }

                            strcat(newFile, line);
                            memset(line, 0, sizeof(line));
                        }
                    }
                    fclose(pf);

                    // Ecriture par ecrasement dans le fichier des nouvelles valeurs
                    FILE *pfw = fopen(RECORDS, "w");
                    fprintf(pfw, "%s", newFile);
                    fclose(pfw);

                    // libération de la mémoire allouée à "newFile"
                    free(newFile);

                    printf("\n✔ Success!\n\n");
                    stayOrReturn(u);
                    break;
                }
            }
            else if (strcmp(r.accountType, "fixed01") == 0 || strcmp(r.accountType, "fixed02") == 0 || strcmp(r.accountType, "fixed03") == 0)
            {
                printf("\n✖ %s account, no trading allowed\n\n", r.accountType);
                stayOrReturn(u);
            }
            else
            {
                printf("\n✖ Error in account type, %s is not valid\n\n", r.accountType);
                stayOrReturn(u);
            }
            break;
        }
    }
    fclose(pf);
    // printf("\n✔ Success!\n\n");
    stayOrReturn(u);
}

void transferAcc(struct User u)
{
    char userName[100];
    struct Record r;
    FILE *pf = fopen(RECORDS, "r");

    system("clear");
    printf("\n\t\t====== All accounts from user, %s =====\n\n", u.name);
    while (getAccountFromFile(pf, userName, &r))
    {
        if (strcmp(userName, u.name) == 0)
        {
            printf("_____________________\n");
            printf("\nAccount number: %d\nDeposit Date: %d/%d/%d \ncountry: %s \nPhone number: %d \nAmount deposited: $%.2f \nType Of Account: %s\n",
                   r.accountNbr,
                   r.deposit.day,
                   r.deposit.month,
                   r.deposit.year,
                   r.country,
                   r.phone,
                   r.amount,
                   r.accountType);
        }
    }
    fseek(pf, 0, SEEK_SET); // re-initialise le curseur au début du fichier

    printf("\n\n\t\tNumber of the account you want to transfer:");
    int accountNbrToTransfer;
    scanf("%d", &accountNbrToTransfer);

    struct User userChecker;
    FILE *us = fopen(USERS, "r");

    printf("\n\t\t====== List of user =====\n\n");
    while (getUsersFromFile(us, &userChecker))
    {
        printf("_____________________\n");
        printf("\nUser number: %d\nUser name: %s\n",
               userChecker.id,
               userChecker.name);
    }
    fseek(us, 0, SEEK_SET); // re-initialise le curseur au début du fichier
    printf("\n\n\t\tNumber of user you want to transfer it:");
    int idUserToTransaction;
    scanf("%d", &idUserToTransaction);

    char userNameToTransfer[50];
    while (getUsersFromFile(us, &userChecker))
    {
        if (idUserToTransaction == userChecker.id)
        {
            strcpy(userNameToTransfer, userChecker.name);
        }
    }
    fclose(us);

numberAccountExist:
    int newaccountNbr;
    printf("\nEnter the new account number:");
    scanf("%d", &newaccountNbr);

    // Vérification si le n° de compte existe déja
    while (getAccountFromFile(pf, userName, &r))
    {
        // printf("%s vs %s et %d == %d\n", userName, userNameToTransfer, newaccountNbr == r.accountNbr);
        if (strcmp(userName, userNameToTransfer) == 0 && newaccountNbr == r.accountNbr)
        {
            printf("✖ This Account already exists for this user\n\n");
            goto numberAccountExist;
        }
    }
    fseek(pf, 0, SEEK_SET); // re-initialise le curseur au début du fichier

    char *newFile;
    newFile = (char *)malloc(100 * sizeof(char));
    // Vider le contenu de newFile (bug, sans il contiens le contenu des 2 fichiers)
    memset(newFile, 0, 100 * sizeof(char));
    char line[300];

    while (getAccountFromFile(pf, userName, &r))
    {
        if (strcmp(userName, u.name) == 0 && accountNbrToTransfer == r.accountNbr)
        {
            // Transfert
            snprintf(line, sizeof(line), "%d %d %s %d %d/%d/%d %s %d %.2lf %s\n\n", r.id, idUserToTransaction, userNameToTransfer, newaccountNbr, r.deposit.month, r.deposit.day, r.deposit.year, r.country, r.phone, r.amount, r.accountType);
        }
        else
        {
            snprintf(line, sizeof(line), "%d %d %s %d %d/%d/%d %s %d %.2lf %s\n\n", r.id, r.userId, userName, r.accountNbr, r.deposit.month, r.deposit.day, r.deposit.year, r.country, r.phone, r.amount, r.accountType);
        }

        // Augmentation de la taille de newFile si line != vide
        if (strlen(line) > 0)
        {

            // Réallocation pour ajouter la taille de line à newFile
            char *temp = realloc(newFile, (strlen(newFile) + strlen(line) + 1) * sizeof(char));
            if (temp == NULL)
            {
                printf("Reallocation de mémoire a échoué.");
                free(newFile); // Libérer la mémoire précédemment allouée
                exit(1);
            }
            else
            {
                newFile = temp;
            }

            // Ajout de line à newFile
            strcat(newFile, line);

            // Vider le tableau line
            memset(line, 0, sizeof(line));
        }
    }
    fclose(pf);

    // Ecriture par ecrasement dans le fichier des nouvelles valeurs
    FILE *pfw = fopen(RECORDS, "w");
    fprintf(pfw, "%s", newFile);
    fclose(pfw);

    // libération de la mémoire allouée à "newFile"
    free(newFile);
    printf("\n✔ Success!\n\n");
    stayOrReturn(u);
}