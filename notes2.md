# Notes Go
## Suite des variables
- Il n'y a pas de changement de type possible, une fois qu'une variable est déclarée, elle ne peut pas changer de type.
- Les variables ne peuvent pas être undifined ou null, elles ont une valeur par défaut (0 pour les int, "" pour les string, etc.)
- Il y a trois manières de déclarer une variable :
  - `var x int = 10`
  - `var x = 10` (type inféré)
  - `x := 10` (type inféré)
  - Il est possible de déclarer plusieurs variables en même temps :
    - `var x, y int = 10, 20`
    - `x, y := 10, 20`
- On peut échanger les valeurs de deux variables sans utiliser de variable temporaire :
  - `x, y = y, x`
- Go refuse les variables non utilisées, ce qui est une bonne pratique pour éviter les bugs liés à des variables oubliées.