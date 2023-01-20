/*
Prototype Deep.- El patron de diseño prototype deep hace referencia a lo mismo 
que el prototype shallow, pero aqui lo que se busca resolver es que cuando 
usamos la clonacion de los objetos las propiedades que son referencias, no se
rompen, por lo cual lo que hacemos es que en el metodo de clonacion creamos 
una instancia completamente nueva de ese objeto con los mismos valores 
asignados manualmente y regresamos dicho objeto, el cual ya va a tener
otras referencias asignadas en si mismo.
*/

namespace PrototypeDeep {
    class Program {
        static void Main(string[] args) {
            Animal animal1 = new Animal {
                Name = "Dolly",
                Feets = 4,
                Details = new Details {
                    Color = "Blanca",
                    Race = "Oveja"
                }
            };

            Animal animal2 = (Animal)animal1.Clone();

            animal1.Details.Color = "Negra";
            animal1.Name = "Oveja Negra";

            Console.WriteLine("Animal original: " + animal1.Details.Color);
            Console.WriteLine("Animal clonado: " + animal2.Details!.Color);

            Console.WriteLine("Animal original: " + animal1.Name);
            Console.WriteLine("Animal clonado: " + animal2.Name);
        }
    }
}