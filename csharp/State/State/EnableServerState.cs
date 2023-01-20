namespace State {
    public class EnableServerState : ServerState {
        public override void Response() {
            Console.WriteLine("Response 200 Ok!");
        }
    }
}