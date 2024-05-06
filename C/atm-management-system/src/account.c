#include "header.h"
#include <stdbool.h>

void checkDetailAcc(struct User u)
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

    printf("\n\n\t\tYou want checking the details of account number:");
    int accountNbrToDetail;
    scanf("%d", &accountNbrToDetail);

    while (getAccountFromFile(pf, userName, &r))
    {
        // printf("r.accountNbr : %d\n", r.accountNbr);
        if (strcmp(userName, u.name) == 0 && accountNbrToDetail == r.accountNbr)
        {
            // printf("r.accountNbr dans le if : %d %s %d\n", r.accountNbr, r.accountType, r.deposit.month);
            if (strcmp(r.accountType, "current") == 0) // absence d'interet
            {
                printf("\nYou will not get interests because the account is of type current\n");
            }
            else if (strcmp(r.accountType, "saving") == 0) // 7%
            {
                printf("\nYou will get $%.2lf as interest on day %d of every month\n", (r.amount * 0.07) / 12, r.deposit.month);
            }
            else if (strcmp(r.accountType, "fixed01") == 0) // 4% sur 1 ans
            {
                printf("\nYou will get $%.2lf as interest on %d/%d/%d\n", (r.amount * 0.04), r.deposit.month, r.deposit.day, r.deposit.year+1);
            }
            else if (strcmp(r.accountType, "fixed02") == 0) // 5% sur 2 ans
            {
                printf("\nYou will get $%.2lf as interest on %d/%d/%d\n", (r.amount * 0.05) *2, r.deposit.month, r.deposit.day, r.deposit.year+2);
            }
            else if (strcmp(r.accountType, "fixed03") == 0) // 8% sur 3 ans
            {
                printf("\nYou will get $%.2lf as interest on %d/%d/%d\n", (r.amount * 0.08) *3, r.deposit.month, r.deposit.day, r.deposit.year+3);
            }
            else
            {
                printf("\n✖ Error in account type, %s is not valid\n", r.accountType);
                stayOrReturn(u);
            }
            break;
        }
    }
    fclose(pf);
    printf("\n✔ Success!\n\n");
    stayOrReturn(u);
}

void updateAcc(struct User u)
{
    char userName[100];
    struct Record r;

    FILE *pf = fopen(RECORDS, "r");
wrongUpdate:
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

    printf("\n\n\t\tYou want update account number:");
    int accountNbrToUpdate;
    scanf("%d", &accountNbrToUpdate);

    bool existingAccount = false;
    while (getAccountFromFile(pf, userName, &r))
    {
        if (strcmp(userName, u.name) == 0 && accountNbrToUpdate == r.accountNbr)
        {
            existingAccount = true;
        }
    }
    fseek(pf, 0, SEEK_SET); // re-initialise le curseur au début du fichier

    if (!existingAccount)
    {
        printf("\n✖ Account number does not exist");
        goto wrongUpdate;
    }

    system("clear");
    printf("\n\n\t\t====== Update account %d from user, %s =====\n\n", accountNbrToUpdate, u.name);
    printf("\n\t\t-->> Feel free to choose one of the options below <<--\n");
    printf("\n\t\t[1]- Update country\n");
    printf("\n\t\t[2]- Update the phone number\n");
    printf("\n\t\t[3]- Exit to the main menu\n");
    printf("\n\t\t[4]- Exit\n");
    printf("\nYour selection : ");
    int option;
    scanf("%d", &option);

    char newCountry[100];
    int newPhone;
    switch (option)
    {
    case 1:
        printf("\nEnter the new country:");
        scanf("%s", newCountry);
        break;
    case 2:
        printf("\nEnter the new phone number:");
        scanf("%d", &newPhone);
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

        while (getAccountFromFile(pf, userName, &r))
        {
            if (strcmp(userName, u.name) == 0 && accountNbrToUpdate == r.accountNbr)
            {
                if (option == 1)
                {
                    snprintf(line, sizeof(line), "%d %d %s %d %d/%d/%d %s %d %.2lf %s\n\n", r.id, r.userId, userName, r.accountNbr, r.deposit.month, r.deposit.day, r.deposit.year, newCountry, r.phone, r.amount, r.accountType);
                }
                else if (option == 2)
                {
                    snprintf(line, sizeof(line), "%d %d %s %d %d/%d/%d %s %d %.2lf %s\n\n", r.id, r.userId, userName, r.accountNbr, r.deposit.month, r.deposit.day, r.deposit.year, r.country, newPhone, r.amount, r.accountType);
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
}

void removeAcc(struct User u)
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

wrongDelete:
    printf("\n\n\t\tYou want delete account number:");
    int accountNbrToDelete;
    scanf("%d", &accountNbrToDelete);

    bool existingAccount = false;
    while (getAccountFromFile(pf, userName, &r))
    {
        if (strcmp(userName, u.name) == 0 && accountNbrToDelete == r.accountNbr)
        {
            existingAccount = true;
        }
    }
    fseek(pf, 0, SEEK_SET); // re-initialise le curseur au début du fichier

    if (!existingAccount)
    {
        printf("\n✖ Account number does not exist\n\n");
        goto wrongDelete;
    }

    char *newFile;
    newFile = (char *)malloc(100 * sizeof(char));
    char line[300];

    // printf("\n\naccountNbrToDelete : %d", accountNbrToDelete);
    while (getAccountFromFile(pf, userName, &r))
    {
        if (strcmp(userName, u.name) == 0 && accountNbrToDelete == r.accountNbr)
        {
            // ligne à supprimer, ne rien faire
        }
        else
        {
            snprintf(line, sizeof(line), "%d %d %s %d %d/%d/%d %s %d %.2lf %s\n\n",
                     // records.txt (id, user_id, user name, account id, date of creation, country, phone nº, balance, type of account)
                     r.id, r.userId, userName, r.accountNbr, r.deposit.month, r.deposit.day, r.deposit.year, r.country, r.phone, r.amount, r.accountType);
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

    // printf("\n\n====== newFile, %s =====\n\n", newFile);

    // Ecriture par ecrasement dans le fichier des nouvelles valeurs
    FILE *pfw = fopen(RECORDS, "w");
    fprintf(pfw, "%s", newFile);
    fclose(pfw);

    // libération de la mémoire allouée à "newFile"
    free(newFile);

    printf("\n✔ Success!\n\n");
    stayOrReturn(u);
}
