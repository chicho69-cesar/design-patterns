namespace State {
    public class SuperSaturedServerState : ServerState {
        public override void Response() {
            Task.Delay(1000);
            Console.WriteLine("Response 200 Ok! super satured");
        }
    }
}