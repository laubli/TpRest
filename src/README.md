# Projet d'API en GO
Ce projet consiste en la création d'une API REST en go avec diverses opérations concernant des langues (languages) et des étudiants (students).

## Mise en marche du projet
Afin d'executer le projet, il faut lancer le fichier cmd/restserveur/main.go
Afin de faciliter cela, j'ai créé un fichier gobuild.bat situé à la raçine du projet qui build et execute le projet.

Toutes les opérations se réalisent sur le port 8080 à l'adresse : http://localhost:8080

## Liste des opérations
Malheuresement je n'ai pas réussis à faire fonctionner Swagger qui aurait dû se situer sur l'adresse : http://localhost:8080/swagger/index.html

À la place, voici une liste des opérations :
Languages
* GET | GetOneLanguage | Retourne un langage | http://localhost:8080/rest/languages/[code]
* GET | GetAllLanguage | Retourne tous les langages | http://localhost:8080/rest/languages
* POST | CreateLanguageHandler | Crée un language | http://localhost:8080/rest/languages
* PUT | UpdateLanguageHandler | Mise à jour d'un langage | http://localhost:8080/rest/languages
* DELETE | DeleteLanguageByIdHandler | Supprime un language selon son code | http://localhost:8080/rest/languages/[code]

Students
* GET | GetOneStudent | Retourne un étudiant | http://localhost:8080/rest/languages/[id]
* GET | GetAllStudent | Retourne tous les étudiants | http://localhost:8080/rest/languages
* POST | CreateStudentHandler | Crée un étudiant | http://localhost:8080/rest/languages
* PUT | UpdateStudentHandler | Mise à jour d'un étudiant | http://localhost:8080/rest/languages
* DELETE | DeleteStudentByIdHandler | Supprime un étudiant selon son ID | http://localhost:8080/rest/languages/[id]

### Comment tester facilement ces opérations ?
Je conseil d'utiliser Postman pour cela. Pour se faire j'ai mis dans le dossiers "tests" des json qui peuvent facilement être importé dans ce logiciel afin d'avoir toutes les commandes déjà faites pour tester l'API.

## Base de donnée en mémoire et sur Bolt
Les deux types de base de donnée en mémoire et sur Bolt fonctionnent.

## Swagger
Je n'ai pas réussi à mettre en place Swagger pour la documentation.
J'ai voulu essayé avec la librairie https://github.com/swaggo/swag et cela n'a pas marché.
