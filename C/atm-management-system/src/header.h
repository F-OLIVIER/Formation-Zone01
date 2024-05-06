#include <stdio.h>
#include <stdlib.h>
#include <string.h>

extern char *RECORDS; // Déclarer RECORDS comme une variable externe
extern char *USERS; // Déclarer USERS comme une variable externe

struct Date
{
    int month, day, year;
};

// tous les champs pour chaque enregistrement d'un compte
struct Record
{
    int id;
    int userId;
    char name[100];
    char country[100];
    int phone;
    char accountType[10];
    int accountNbr;
    double amount;
    struct Date deposit;
    struct Date withdraw;
};

struct User
{
    int id;
    char name[50];
    char password[50];
};

// Fonctions d'authentification
void loginMenu(char a[50], char pass[50]);
const char *getPassword(struct User u);

// Fonctions d'enregistrement
void registerMenu(char a[50], char pass[50]);

// Fonction système
void mainMenu(struct User u);
void stayOrReturn(struct User u);

// Fonction d'accées fichier
int getUsersFromFile(FILE *ptr, struct User *u);
int getAccountFromFile(FILE *ptr, char name[50], struct Record *r);
void saveAccountToFile(FILE *ptr, struct User u, struct Record r);

// Fonction pour le compte
void createNewAcc(struct User u);
void checkDetailAcc(struct User u);
void checkAllAccounts(struct User u);
void updateAcc(struct User u);
void removeAcc(struct User u);

// Fonctions pour les transactions
void maketransactionAcc(struct User u);
void transferAcc(struct User u);
