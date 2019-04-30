# Actividad 1 - Ejercicio 5 - Item 1

programar en cualquier lenguaje de programación que permita dejar un ejecutable donde se generaran variables aleatorias basadas en un generado random U[0; 1]. en todos los casos debera mostrar el random y a su lado los valores obtenidos y mostrar la formula que utilizo (solo la formula para poder verificar los datos, no el calculo).

El programa debe tener 7 opciones y en cada una de ellas tendra una pausa cada 5 numeros generados:

1- generar números aleatorios con distribución uniforme continua entre 13 y 17 (decilitros de champagne a servir en cada copa en una fiesta con miles de invitados)

2- generar números aleatorios con distribución uniforme discreta entre 20 y 25 (numero de copas a llevar en la bandeja del mozo en cada viaje)

3- generar números aleatorios con distribución exponencial con media ( 1/lambda)= 10. (numero de segundos que transcurre entre que un invitado solicita una copa de una bandeja del mozo y otro invitado es el siguiente en solicitar otra copa

4-mostrar los items 1, 2 y 3 en columnas pero utilizando el mismo numero aleatorio para los tres.  

5- generar números aleatorios con distribución empirica donde x representa los segundos que tarda una persona en extraer una copa de la bandeja.

    f(x)= x si 0<=x<=1;   

    f(x)= 2-x si 1<=x<=2;   

    f(x)= 0   para cualquier otro caso;   

6- generar números aleatorios con distribución poisson con media 1,8 personas llegan en cada auto a la fiesta

7- generar números aleatorios con distribución normal con media 3 minutos t desvio 0,5 minutos que es el tiempo entre la llegada de un auto y el siguiente a la fiesta a partir de las 21:00. Ademas mostrar la hora de llegada del auto.



*****

Aclaraciones Adicionales:

Realizar un programa, preferentemente un ejecutable que no necesite instalar librerias adicionales en la maquina que se ejecute, que tenga un menu de al menos 7 opciones que podian ser identificadas por ejemplo con el nombre de la variable a generar:

- 1.dl champagne
- 2.copas champagne
- 3.segundos entre pedido de copas
- 4.opcion 1.2.3 juntas
- 5.Segundos en extraer copa
- 6.personas por auto
- 7.minutos entre autos
- 0.Salir
- h. help (estaria genial con una opcion asi pero no es necesaria)
- p.Imprimir a pdf (estaria genial con una opcion asi pero no es necesaria)
por ejemplo la opcion 1 deberia mostrar:

- titulo: (nombre de la variable)
- nombre de la distribucion a usar
- Formula utilizada (escrita o una imagen)
encabezado de columnas:  /// NºSimulacion /// Nº Aleatorio usado rnd() /// Nº Aleatorio Generado en la variable
ejemplo:
Variable: decilitros por copa

distribucion U[13;17]

formula: x = 13 + (17-3) * rnd()

/// NºSimulacion /// Nº Aleatorio usado rnd() /// Nº Aleatorio Generado en la variable

/// 1                      ///   0.35       (***)                  ///   x1= 13 + (17-13) * 0.35 = 14.40 dl  (****)

/// 2                     ///    0.81                                ///    x2= 16.24

/// 3                     ///    0.50                                ///    x3= 15.00    

/// 4                     ///    0.81                                ///    x2= 16.24

/// 5                     ///    0.50                                ///    x3= 15.00    

*********** acá, luego de 5 números vienen dos opciones:  /// Continuar  /// Salir  /// ****  si pulsa continuar genera 5 mas y salir va al menú principal

(***) los números aleatorio deberían ser generados con un formula del software

(****) (en este espacio no es necesario todo el texto, podría ser x1=14.40 dl  o solo 14.40
