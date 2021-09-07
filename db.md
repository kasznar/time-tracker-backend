## funkciok:
* felvenni szabikat
* felhasznalo lista
* felhasznalo jogok: basic, admin, hr, csoportvezeto
* aznapi feladatok (string)
* munkanap valami vagy sem
* kotelezo szabik: szabi amit valaki mas vett fel neked
* szabi keres, engedelyezes
* hany szabid van
* betegszabi
* hozzadni embereket
* nem egesz ev kezelese

## extra funckiok:
* esemenyek
* ertesitesek

## adatbazis:
### felhasznalok:
* id
* tipus
* napi munkaora
* szokasos kezdes, szokasos vegzes
* email
* kezdes idopontja
* jogviszony

* number_of_vacations: object? [2021: 22, 2022: 23], joint table?

### szabi/ev:
* id
* ev
* user id

### munkanap:
* id
* tipus: tipus id (nap tipus)
* mivel foglalkoztal: string
* datum
* napok az evben kapcsolat id
* betegszabi?
* felhasznalo azonosito
* elfogadottsagi allapot

### munkanap tipus:
* id
* megnevezes

### kerelmek:
* id
* felado:
* cimzett:
* nap id: 
* allapot
* feladva
* elbiralva

### naptari nap
* id
* tipus: naptari nap tipus

### naptari nap tipus
* id
* tipus (munkanap, szabadnap, kotelezo szabi)




