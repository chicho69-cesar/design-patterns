namespace State {
    public class SaturedServerState : ServerState {
        public override void Response() {
            Task.Delay(500);
            Console.WriteLine("Response 200 Ok! satured");
        }
    }
}