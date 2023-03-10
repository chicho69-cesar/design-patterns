namespace Mediador {
    public class User : Colleague {
        public User(IMediator mediator) : base(mediator) { }

        public override void Receive(string message) {
            Console.WriteLine("User receive: " + message);
        }
    }
}