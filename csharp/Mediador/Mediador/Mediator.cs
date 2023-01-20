namespace Mediador {
    public class Mediator : IMediator {
        private List<Colleague> colleagues;

        public Mediator() {
            colleagues = new List<Colleague>();
        }

        public void Add(Colleague colleague) {
            this.colleagues.Add(colleague);
        }

        public void Send(string message, Colleague colleague) {
            foreach (Colleague c in colleagues) {
                if (colleague != c) {
                    c.Receive(message);
                }
            }
        }
    }
}