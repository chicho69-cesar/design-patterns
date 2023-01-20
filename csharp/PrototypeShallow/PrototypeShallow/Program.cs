/*
Prototype Shallow.- El patron prototype shallow o superficial, lo que busca es
romper con la referencia de los objetos al momento de asignarlos de una variable 
a otra, para esta hacemos uso del metodo Clone que forma parte de la interfaz
ICloneable, el cual nos permite regresar una nueva instancia del objeto con los
mismos datos pero otra referencia
*/

namespace PrototypeShallow {
    class Program {
        static void Main(string[] args) {
            Animal animal1 = new() {
                Name = "Oveja Dolly",
                Feets = 4
            };

            Animal animal2 = animal1.Clone() as Animal ?? new();
            animal2.Feets = 3;

            Console.WriteLine(animal1.Feets);
            Console.WriteLine(animal2.Feets);
        }
    }
}