namespace Mediador {
    public abstract class Colleague {
        private IMediator _mediator;

        public Colleague(IMediator mediator) {
            _mediator = mediator;
        }

        public IMediator Mediator {
            get {
                return this._mediator;
            }

            set {
                this._mediator = value;
            }
        }

        public void Communicate(string message) {
            this._mediator.Send(message, this);
        }

        public abstract void Receive(string message);
    }
}