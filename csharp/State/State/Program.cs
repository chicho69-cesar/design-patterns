/*
State.- State es un patron de diseño el cual los permite manipular una serie
de objetos a traves de un contexto dependiendo del estado de un objeto, es decir,
podemos tener una clase abstracta o una interfaz de la cual heredan diferentes
clases las cuales van a tener una funcionalidad similar, pero la cual va a 
actuar de diferente forma entre cada clase dependiendo de un estado determinado,
por lo cual el contexto lo que hara sera usar la clase abstracta para invocar 
la accion que debemos ejecutar, pero dependiendo del estado esta accion
sera hecho por una clase o por otra.
*/

namespace State {
    class Program {
        static void Main(string[] args) {
            ServerContext server = new();
            
            server.ServerState = new EnableServerState();
            server.AttendRequest();

            server.ServerState = new SaturedServerState();
            server.AttendRequest();
            server.AttendRequest();

            server.ServerState = new SuperSaturedServerState();
            server.AttendRequest();
            server.AttendRequest();

            server.ServerState = new DeadServerState();
            server.AttendRequest();
            server.AttendRequest();

            server.ServerState = new EnableServerState();
            server.AttendRequest();
        }
    }
}