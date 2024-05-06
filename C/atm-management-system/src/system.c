#include "header.h"

char *USERS = "./data/users.txt";
char *RECORDS = "./data/records.txt";

void stayOrReturn(struct User u)
{
    int option;
invalid:
    printf("Enter 1 to go to the main menu and 0 to exit!\n");
    scanf("%d", &option);
    system("clear");
    if (option == 1)
    {
        mainMenu(u);
    }
    else if (option == 0)
    {
        exit(1);
    }
    else
    {
        printf("✖ Insert a valid operation!\n\n");
        goto invalid;
    }
}

void createNewAcc(struct User u)
{
    struct Record r;
    struct Record cr;
    char userName[50];
    FILE *pf = fopen(RECORDS, "a+");

noAccount:
    system("clear");
    printf("\t\t\t===== New record =====\n");

    printf("\nEnter today's date(mm/dd/yyyy):");
    scanf("%d/%d/%d", &r.deposit.month, &r.deposit.day, &r.deposit.year);
    printf("\nEnter the account number:");
    getchar(); // Vide le tampon d'entrée
    scanf("%d", &r.accountNbr);

    // Vérification si le n° de compte existe déja
    while (getAccountFromFile(pf, userName, &cr))
    {
        // printf("%s vs %s et %d == %d\n", userName, u.name, cr.accountNbr == r.accountNbr);
        if (strcmp(userName, u.name) == 0 && cr.accountNbr == r.accountNbr)
        {
            printf("✖ This Account already exists for this user\n\n");
            goto noAccount;
        }
    }
    printf("\nEnter the country:");
    scanf("%s", r.country);
    printf("\nEnter the phone number:");
    scanf("%d", &r.phone);
    printf("\nEnter amount to deposit: $");
    scanf("%lf", &r.amount);
wrongType:
    printf("\nChoose the type of account:\n\t1 -> saving\n\t2 -> current\n\t3 -> fixed01(for 1 year)\n\t4 -> fixed02(for 2 years)\n\t5 -> fixed03(for 3 years)\n\n\tEnter number of your choice:");
    int NbrAccountType;
    scanf("%d", &NbrAccountType);

    if (NbrAccountType == 1)
    {
        strcpy(r.accountType, "saving");
    }
    else if (NbrAccountType == 2)
    {
        strcpy(r.accountType, "current");
    }
    else if (NbrAccountType == 3)
    {
        strcpy(r.accountType, "fixed01");
    }
    else if (NbrAccountType == 4)
    {
        strcpy(r.accountType, "fixed02");
    }
    else if (NbrAccountType == 5)
    {
        strcpy(r.accountType, "fixed03");
    }
    else
    {
        goto wrongType;
    }

    saveAccountToFile(pf, u, r);

    fclose(pf);
    printf("\n✔ Success!\n\n");
    stayOrReturn(u);
}

int getUsersFromFile(FILE *ptr, struct User *u)
{
    return fscanf(ptr, "%d %s %s",
                  &u->id,
                  u->name,
                  u->password) != EOF;
}

int getAccountFromFile(FILE *ptr, char name[50], struct Record *r)
{
    return fscanf(ptr, "%d %d %s %d %d/%d/%d %s %d %lf %s",
                  &r->id,
                  &r->userId,
                  name,
                  &r->accountNbr,
                  &r->deposit.month,
                  &r->deposit.day,
                  &r->deposit.year,
                  r->country,
                  &r->phone,
                  &r->amount,
                  r->accountType) != EOF;
}

void saveAccountToFile(FILE *ptr, struct User u, struct Record r)
{
    // Récupération du dernier id dans record
    FILE *fpID;
    struct Record recordChecker;
    int last_id_record = -1;

    if ((fpID = fopen("./data/records.txt", "r")) == NULL)
    {
        printf("Error! opening file");
        exit(1);
    }
    while (fscanf(fpID, "%d %d %s %d %d/%d/%d %s %d %lf %s", &recordChecker.id, &recordChecker.userId, recordChecker.name, &recordChecker.accountNbr, &recordChecker.deposit.month, &recordChecker.deposit.day, &recordChecker.deposit.year, recordChecker.country, &recordChecker.phone, &recordChecker.amount, recordChecker.accountType) != EOF)
    {
        if (recordChecker.id > last_id_record)
        {
            last_id_record = recordChecker.id;
        }
    }
    fclose(fpID);
    last_id_record++;

    // Ecriture dans le fichier des valeurs
    fprintf(ptr, "%d %d %s %d %d/%d/%d %s %d %.2lf %s\n\n",
            // records.txt (id, user_id, user name, account id, date of creation, country, phone nº, balance, type of account)
            last_id_record,  // int %d
            u.id,            // int %d
            u.name,          // char %s
            r.accountNbr,    // int %d
            r.deposit.month, // int %d
            r.deposit.day,   // int %d
            r.deposit.year,  // int %d
            r.country,       // char %s
            r.phone,         // int %d
            r.amount,        // long (ou double) float à 2 décimale %.2lf ou %2lf
            r.accountType);  // char %s
}

void checkAllAccounts(struct User u)
{
    char userName[100];
    struct Record r;

    FILE *pf = fopen(RECORDS, "r");

    system("clear");
    printf("\t\t====== All accounts from user, %s =====\n\n", u.name);
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
    fclose(pf);
    printf("\n✔ Success!\n\n");
    stayOrReturn(u);
}
