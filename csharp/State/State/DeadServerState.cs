namespace State {
    public class DeadServerState : ServerState {
        public override void Response() {
            Console.WriteLine("Response 503 dead");
        }
    }
}