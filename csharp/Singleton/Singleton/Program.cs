/*
Singleton: Este patron de diseño hace referencia a crear una clase, la cual
tiene un atributo privado estatico que es una instancia de la misma clase, y 
para acceder a el creamos una propiedad tambien estatica que lo que va a hacer
es que si la instancia del singleton es null entonces le va a asignar un valor
y la va a regresar, de esta forma las siguientes veces que queramos acceder
a esta instancia, ya no se va a volver a crear porque ya tiene un valor
*/

namespace Singleton {
    class Program {
        static void Main(string[] args) {
            Console.WriteLine(Singleton.Instance?.message);
            Console.WriteLine(Singleton.Instance!.message);

            Singleton.Instance.message = "I'm Cesar Villalobos Olmos";

            Console.WriteLine(Singleton.Instance!.message);
        }
    }
}