/*
Mediador.- El patron de diseño mediador funcion cuando tenemos objetos que son
de diferente tipo y queremos hacer una especie de comunicacion entre ellos, 
para evitar tener que hacer esta comunicacion a mano en cada uno de los objetos
lo que hacemos es que creamos un objeto mediador el cual en su nucleo tenga
una lista de objetos con los elementos que queremos comunicar y creamos un metodo
el cual se va a encargar de dicha comunicacion, realizando el envio de informacion
de un objeto a otro a traves de si mismo, pero para esto los tipos de datos
de datos que queremos comunicar deben de heredar de un tipo de dato en comun.
*/

namespace Mediador {
    class Program {
        static void Main(string[] args) {
            Mediator mediator = new();

            Colleague user = new User(mediator);
            Colleague userAdmin1 = new UserAdmin(mediator);
            Colleague userAdmin2 = new UserAdmin(mediator);

            mediator.Add(user);
            mediator.Add(userAdmin1);
            mediator.Add(userAdmin2);

            user.Communicate("Oye admin tengo un problema :(");
        }
    }
}