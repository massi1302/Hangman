# 🎮 Jeu du Pendu (Hangman Game)

Un jeu du pendu en ligne de commande développé en Go, avec plusieurs niveaux de difficulté et une interface utilisateur interactive.

## 📋 Table des matières
- [Fonctionnalités](#fonctionnalités)
- [Prérequis](#prérequis)
- [Installation](#installation)
- [Utilisation](#utilisation)
- [Structure du projet](#structure-du-projet)
- [Règles du jeu](#règles-du-jeu)
- [Contribution](#contribution)

## ✨ Fonctionnalités

- 🎯 Trois niveaux de difficulté (Facile, Moyen, Difficile)
- 🎨 Interface colorée et animations en ASCII art
- 💖 Système de vies avec affichage visuel
- 🔤 Option de révélation initiale de lettres
- ⌨️ Navigation intuitive dans les menus

## 🛠️ Prérequis

- Go 1.16 ou supérieur
- Les dépendances suivantes :
  ```bash
  github.com/fatih/color
  github.com/eiannone/keyboard
  ```

## 📥 Installation

1. Clonez le répertoire :
   ```bash
   git clone https://github.com/votre-username/jeu-de-pendu.git
   cd jeu-de-pendu
   ```

2. Installez les dépendances :
   ```bash
   go mod download
   ```

## 🎮 Utilisation

Il existe deux façons de lancer le jeu :

1. Mode menu interactif :
   ```bash
   go run main.go
   ```

2. Mode direct avec fichier de mots et lettres révélées :
   ```bash
   go run main.go [fichier-mots] [nombre-lettres]
   ```
   
   Exemples :
   ```bash
   go run main.go easy-words.txt 2    # Mode facile avec 2 lettres révélées
   go run main.go medium-words.txt 1   # Mode moyen avec 1 lettre révélée
   go run main.go hard-words.txt 0     # Mode difficile sans lettre révélée
   ```

## 📁 Structure du projet

```
jeu-de-pendu/
├── main.go            # Point d'entrée du programme
├── game/
│   ├── hangman.go     # Logique principale du jeu
│   ├── menu.go        # Gestion du menu
│   ├── affichage.go   # Fonctions d'affichage et couleurs
│   ├── wordsutil.go   # Utilitaires de gestion des mots
│   ├── asciiart.go    # Art ASCII pour le pendu
│   └── clearconsole.go # Utilitaire console
├── data/
│   ├── easy-words.txt  # Liste de mots faciles
│   ├── medium-words.txt # Liste de mots moyens
│   └── hard-words.txt  # Liste de mots difficiles
└── README.md
```

## 📌 Règles du jeu

1. Un mot est choisi aléatoirement selon le niveau de difficulté
2. Le joueur commence avec 10 vies (❤️)
3. À chaque tour, le joueur peut :
   - Proposer une lettre
   - Deviner le mot complet
4. Le joueur perd une vie (💔) pour chaque :
   - Lettre incorrecte
   - Mot incorrect (2 vies)
5. La partie est gagnée si le mot est trouvé avant de perdre toutes les vies

## 🎯 Niveaux de difficulté

- **Facile** : Mots courts (4-5 lettres)
- **Moyen** : Mots de longueur moyenne
- **Difficile** : Mots longs et complexes

## 🤝 Contribution

Les contributions sont les bienvenues ! Pour contribuer :

1. Forkez le projet
2. Créez une branche pour votre fonctionnalité
3. Committez vos changements
4. Poussez vers la branche
5. Ouvrez une Pull Request



Développé avec ❤️ et Go
