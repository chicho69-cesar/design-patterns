namespace Mediador {
    public class UserAdmin : Colleague {
        public UserAdmin(IMediator mediator) : base(mediator) { }

        public override void Receive(string message) {
            Console.WriteLine("Admin recieve: " + message);
        }
    }
}