/*
Factory.- El patron de diseño factory lo que hace es que mediante una clase
abstracta de la cual heredan otros tipos de datos, lo que hacemos es que 
creamos una clase la cual nos va a servir de intermediario para crear objetos
ya sea mediante un enum o mediante constantes, podemos identificar que
objeto es el que queremos crear y este metodo nos va a regresar una instancia
de dicho objeto
*/

namespace Factory {
    class Program {
        static void Main(string[] args) {
            AlcoholicDrink? drink1 = Creator.DrinkCreator(Creator.BEER);
            AlcoholicDrink? drink2 = Creator.DrinkCreator(Creator.TINT_WINE);
            
            Console.WriteLine(drink1!.HowMuchAlcoholizeMePerHour());
            Console.WriteLine(drink2!.HowMuchAlcoholizeMePerHour());
        }
    }
}