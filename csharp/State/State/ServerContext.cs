namespace State {
    public class ServerContext {
        private ServerState serverState;

        public ServerState ServerState {
            get {
                return serverState;
            }

            set {
                serverState = value;
            }
        }

        public void AttendRequest() {
            serverState.Response();
        }
    }
}