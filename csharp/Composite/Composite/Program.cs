/*
Composite.- El patron de diseño composite sirve para tener clases las cuales 
hereden de una clase component, la cual tiene los atributos o propiedades necesarias
para que los hijos que hereden de ella funcionen correctamente, hablando de los 
hijos, tenemos dos tipos, los que son hojas o leaf que son simples componentes
que vamos a usar para crear componentes mas grandes y complejos, despues los
componentes mas grandes o complejos pueden contener dentro de si mismos una lista
con objetos de tipo component que pueden ser tanto hojas como componentes grandes.
Y mediante esto podemos crear objetos de estos componentes a traves de otros 
componentes.
*/

namespace Composite {
    class Program {
        static void Main(string[] args) {
            Ingredient ingredient1 = new("Harina", 100, 200, "gramos");
            Ingredient ingredient2 = new("Leche", 20, 1, "litro");
            Ingredient ingredient3 = new("Huevo", 45, 1, "kilogramo");

            CakeComposite cakeOfMilk = new("Pastel de leche", 200);
            cakeOfMilk.Add(ingredient1);
            cakeOfMilk.Add(ingredient2);
            cakeOfMilk.Add(ingredient3);

            Console.WriteLine(cakeOfMilk.Cost);

            Ingredient ingredient4 = new("Chocolate", 100, 1, "kilogramo");

            CakeComposite cakeOfChocoAndMilk = new("Pastel de chocolate y leche", 400);
            cakeOfChocoAndMilk.Add(ingredient4);
            cakeOfChocoAndMilk.Add(cakeOfMilk);

            Console.WriteLine(cakeOfChocoAndMilk.Cost);
        }
    }
}