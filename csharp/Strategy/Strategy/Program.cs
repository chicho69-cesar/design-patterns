/*
Strategy.- El patron de diseño strategy lo que nos conlleva es que cuando tenemos
un objeto el cual puede tener diferentes comportamientos, para esto creamos
diferentes clases que van a representar cada uno de estos comportamientos
y ademas creamos un contexto para la estrategia, donde vamos a definir una instancia
de la interfaz que comparten cada una de las diferentes estrategias que 
definimos y mediante metodos le asignamos una instancia de un comportamiento
diferente o bien tambie lo podemos hacer mediante un enum, por ultimo creamos
dentro del contexto un metodo que va a disparar los comportamientos de la
estrategia, podiendo ejecutar comportamientos diferentes dependiendo
de la instancia que se tenga
*/

namespace Strategy {
    class Program {
        static void Main(string[] args) {
            var alcoholic1 = new StrategyContext();

            alcoholic1.BeRomantic();
            alcoholic1.SetInvite();
            alcoholic1.BeRomantic();
            alcoholic1.SetPrince();
            alcoholic1.BeRomantic();

            Console.WriteLine("\n\n");

            var alcoholic2 = new StrategyContext();
            alcoholic2.BeRomantic(StrategyContext.Behaviour.Eyes);
            alcoholic2.BeRomantic(StrategyContext.Behaviour.Invite);
            alcoholic2.BeRomantic(StrategyContext.Behaviour.Prince);
        }
    }
}